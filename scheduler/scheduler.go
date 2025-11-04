package scheduler

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"math/rand"
	"net/http"
	"time"

	"github.com/brevis-network/prover-network-bidder/client"
	"github.com/brevis-network/prover-network-bidder/config"
	"github.com/brevis-network/prover-network-bidder/dal"
	"github.com/brevis-network/prover-network-bidder/eth"
	"github.com/brevis-network/prover-network-bidder/onchain"
	"github.com/celer-network/goutils/log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"github.com/spf13/viper"
)

type Scheduler struct {
	*dal.DAL
	*onchain.ChainClient
	*client.ProverNetworkClient
	ruleConfig *config.RuleConfig
}

func NewScheduler(db *dal.DAL, c *onchain.ChainClient, p *client.ProverNetworkClient) (*Scheduler, error) {
	var ruleConfig config.RuleConfig
	err := viper.UnmarshalKey(config.KeyRule, &ruleConfig)
	if err != nil {
		return nil, fmt.Errorf("UnmarshalKey err: %w", err)
	}
	return &Scheduler{db, c, p, &ruleConfig}, nil
}

func (s *Scheduler) Start() {
	go s.scheduleAppRegister()
	go s.scheduleBid()
	go s.scheduleReveal()
	go s.scheduleQueryBidResult()
	go s.scheduleProve()
	go s.scheduleQueryProvingResult()
	go s.scheduleSubmitProof()
}

func (s *Scheduler) scheduleAppRegister() {
	for {
		time.Sleep(5 * time.Second)
		apps, err := s.FindNotRegisteredApps(context.Background())
		if err != nil {
			log.Errorf("FindNotRegisteredApps err: %s", err)
			continue
		}
		for _, app := range apps {
			// TODO: retrieve elf from img url
			// assume HTTP GET to downlaod
			elf, err := DownloadFile(app.ImgUrl)
			if err != nil {
				log.Errorf("app %s, download %s err: %s", app.AppID, app.ImgUrl, err)
				continue
			}

			log.Infof("Start register app %s...", app.AppID)
			retAppId, err := s.RegisterApp("", elf)
			if err != nil {
				log.Errorf("RegisterApp %s err: %s", app.AppID, err)
				continue
			}
			log.Infof("End register app %s", app.AppID)

			if retAppId != app.AppID {
				errMsg := fmt.Sprintf("calc appId %s does not match request appId %s", retAppId, app.AppID)
				log.Errorln(errMsg)
				err = s.UpdateAppAsRegisterFailed(context.Background(), dal.UpdateAppAsRegisterFailedParams{
					AppID:         app.AppID,
					RegisterError: errMsg,
				})
				if err != nil {
					log.Errorf("UpdateAppAsRegisterFailed %s err: %s", app.AppID, err)
				}
				continue
			}

			err = s.UpdateAppAsRegisterSuccess(context.Background(), app.AppID)
			if err != nil {
				log.Errorf("UpdateAppAsRegisterSuccess %s err: %s", app.AppID, err)
				continue
			}
		}
	}
}

func (s *Scheduler) scheduleBid() {
	for {
		time.Sleep(5 * time.Second)
		reqs, err := s.FindNotProcessedProofRequests(context.Background())
		if err != nil {
			log.Errorf("FindNotProcessedProofRequests err: %s", err)
			continue
		}
		for _, req := range reqs {
			if req.InputData == "" {
				ok, err := checkInputSize(req.InputUrl, s.ruleConfig.MaxInputSize)
				if err != nil {
					log.Errorf("checkInputSize for req %s err: %s", req.ReqID, err)
					continue
				}
				if !ok { // input too large
					log.Infof("skip req %s due to input exceeds max size", req.ReqID)
					s.UpdateRequestAsProcessed(context.Background(), req.ReqID)
					continue
				}
			}
			// TODO: get inputs from req.InputData or req.InputUrl
			inputs, err := retrieveInputs(req.InputData, req.InputUrl)
			if err != nil {
				log.Errorf("retrieveInputs for req %s err: %s", req.ReqID, err)
				continue
			}

			proverGas, pvDigest, err := s.EstimateCost(req.AppID, inputs)
			if err != nil {
				log.Errorf("EstimateCost %s err: %s", req.ReqID, err)
				continue
			}
			if common.BytesToHash(pvDigest) != common.HexToHash(req.PublicValuesDigest) {
				log.Errorf("pvDigest not match, %x vs %s", pvDigest, req.PublicValuesDigest)
				err = s.UpdateRequestAsProcessed(context.Background(), req.ReqID)
				if err != nil {
					log.Errorf("UpdateRequestAsProcessed %s err: %s", req.ReqID, err)
				}
				continue
			}

			proverGasPrice, _ := big.NewInt(0).SetString(s.ruleConfig.ProverGasPrice, 0)
			myFee := big.NewInt(0).Mul(proverGasPrice, big.NewInt(int64(proverGas)))
			maxFee, _ := big.NewInt(0).SetString(s.ruleConfig.MaxFee, 0)

			if myFee.Cmp(maxFee) == 1 {
				log.Infof("fee exceeds my rule, myFee %s, maxFee in rule %s", myFee.String(), maxFee.String())
				err = s.UpdateRequestAsProcessed(context.Background(), req.ReqID)
				if err != nil {
					log.Errorf("UpdateRequestAsProcessed %s err: %s", req.ReqID, err)
				}
				continue
			}

			// submit bid if pass rule
			nonce := big.NewInt(int64(rand.Uint64()))
			bidHash := solsha3.SoliditySHA3(
				[]string{"uint256", "uint256"},
				[]interface{}{myFee, nonce},
			)
			tx, _, err := onchain.TransactWaitSuccess(
				s.Client, s.TransactOpts,
				func(transactor bind.ContractTransactor, opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
					return s.Bid(opts, common.HexToHash(req.ReqID), common.BytesToHash(bidHash))
				})
			if err != nil {
				errString := err.Error()
				var jsonErr JsonError
				errJson, _ := json.Marshal(err)
				json.Unmarshal(errJson, &jsonErr)
				if jsonErr.Data != "" {
					errName, _ := ParseSolCustomErrorName(eth.BrevisMarketABI, common.FromHex(jsonErr.Data))
					errString = errString + " - " + errName
				}
				log.Errorf("Bid req %s err: %s", req.ReqID, errString)

				abi, _ := eth.BrevisMarketMetaData.GetAbi()
				calldata, _ := abi.Pack("bid", common.HexToHash(req.ReqID), common.BytesToHash(bidHash))
				log.Infof("Bid req %s calldata: %x", req.ReqID, calldata)

				if jsonErr.Data != "" /*not satisfy contract requirement*/ {
					err = s.UpdateRequestAsProcessed(context.Background(), req.ReqID)
					if err != nil {
						log.Errorf("UpdateRequestAsProcessed %s err: %s", req.ReqID, err)
					}
				}
				continue
			}
			log.Infof("Bid req %s tx %s", req.ReqID, tx.Hash().Hex())
			err = s.UpdateRequestAsProcessed(context.Background(), req.ReqID)
			if err != nil {
				log.Errorf("UpdateRequestAsProcessed %s err: %s", req.ReqID, err)
			}

			shouldRevealAfter := req.CreatedAt + int64(s.BiddingPhaseDuration)
			shouldRevealBefore := shouldRevealAfter + int64(s.RevealPhaseDuration)
			err = s.AddBid(context.Background(), dal.AddBidParams{
				ReqID:              req.ReqID,
				MyFee:              myFee.String(),
				BidNonce:           nonce.String(),
				ShouldRevealAfter:  shouldRevealAfter,
				ShouldRevealBefore: shouldRevealBefore,
			})
			if err != nil {
				log.Errorf("AddBid %s err: %s", req.ReqID, err)
			}
		}
	}
}

// HEAD for content length, if larger than maxSize, return false
// otherwise return true. any http error will be false
// if maxSize is 0, return true, nil directly
func checkInputSize(inputUrl string, maxSize uint64) (bool, error) {
	if maxSize == 0 {
		return true, nil // 0 means no cap
	}
	resp, err := http.Head(inputUrl)
	if err != nil {
		return false, fmt.Errorf("http HEAD %s err: %w", inputUrl, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("http HEAD %s unexpected status code: %d", inputUrl, resp.StatusCode)
	}
	if resp.ContentLength > int64(maxSize) {
		return false, nil // too large
	}
	return true, nil
}

func retrieveInputs(inputData string, inputUrl string) ([]byte, error) {
	var inputs []byte
	if inputData != "" {
		inputs = common.Hex2Bytes(inputData)
	} else {
		// assume HTTP GET to downlaod
		inputFromFile, err := DownloadFile(inputUrl)
		if err != nil {
			return nil, err
		}
		inputs = inputFromFile
	}

	return inputs, nil
}

func (s *Scheduler) scheduleReveal() {
	for {
		time.Sleep(5 * time.Second)
		bids, err := s.FindToBeRevealedBid(context.Background(), time.Now().Unix())
		if err != nil {
			log.Errorf("FindToBeRevealedBid err: %s", err)
			continue
		}
		for _, bid := range bids {
			myFee, _ := big.NewInt(0).SetString(bid.MyFee, 0)
			nonce, _ := big.NewInt(0).SetString(bid.BidNonce, 0)
			tx, _, err := onchain.TransactWaitSuccess(
				s.Client, s.TransactOpts,
				func(transactor bind.ContractTransactor, opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
					return s.Reveal(opts, common.HexToHash(bid.ReqID), myFee, nonce)
				})
			if err != nil {
				errString := err.Error()
				var jsonErr JsonError
				errJson, _ := json.Marshal(err)
				json.Unmarshal(errJson, &jsonErr)
				if jsonErr.Data != "" {
					errName, _ := ParseSolCustomErrorName(eth.BrevisMarketABI, common.FromHex(jsonErr.Data))
					errString = errString + " - " + errName
				}
				log.Errorf("Reveal req %s err: %s", bid.ReqID, errString)
				continue
			}
			log.Infof("Reveal req %s tx %s", bid.ReqID, tx.Hash().Hex())
			err = s.UpdateBidAsRevealed(context.Background(), bid.ReqID)
			if err != nil {
				log.Errorf("UpdateBidAsRevealed %s err: %s", bid.ReqID, err)
			}
		}
	}
}

type BidResult string

const (
	Success BidResult = "success"
	Fail    BidResult = "fail"
)

func (s *Scheduler) scheduleQueryBidResult() {
	for {
		time.Sleep(5 * time.Second)
		bids, err := s.FindBidsWithoutResult(context.Background(), time.Now().Unix())
		if err != nil {
			log.Errorf("FindBidsWithoutResult err: %s", err)
			continue
		}
		for _, bid := range bids {
			reqState, err := s.Requests(nil, common.HexToHash(bid.ReqID))
			if err != nil {
				log.Errorf("Requests req %s err: %s", bid.ReqID, err)
				continue
			}

			result := Fail
			if reqState.Winner.Prover == common.HexToAddress(s.ProverEthAddr) {
				result = Success
			}
			err = s.UpdateBidResult(context.Background(), dal.UpdateBidResultParams{
				BidResult: string(result),
				ReqID:     bid.ReqID,
			})
			if err != nil {
				log.Errorf("UpdateBidResult %s err: %s", bid.ReqID, err)
			}
		}
	}
}

type ProofState string

const (
	Init      ProofState = "init"
	Generated ProofState = "generated"
	Submitted ProofState = "submitted"
)

func (s *Scheduler) scheduleProve() {
	for {
		time.Sleep(5 * time.Second)
		bids, err := s.FindToBeProvedBids(context.Background(), time.Now().Unix())
		if err != nil {
			log.Errorf("FindToBeProvedBids err: %s", err)
			continue
		}
		for _, bid := range bids {
			// TODO: get inputs from bid.InputData or bid.InputUrl
			inputs, err := retrieveInputs(bid.InputData, bid.InputUrl)
			if err != nil {
				log.Errorf("retrieveInputs for req %s err: %s", bid.ReqID, err)
				continue
			}

			err = s.ProveTask(bid.AppID, bid.ReqID, inputs)
			if err != nil {
				log.Errorf("ProveTask %s err: %s", bid.ReqID, err)
				continue
			}
			err = s.UpdateBidProofTaskId(context.Background(), dal.UpdateBidProofTaskIdParams{
				ProofTaskID: bid.ReqID,
				ReqID:       bid.ReqID,
			}) // use reqId as proofTaskId
			if err != nil {
				log.Errorf("UpdateBidProofTaskId %s err: %s", bid.ReqID, err)
			}
		}
	}
}

func (s *Scheduler) scheduleQueryProvingResult() {
	for {
		time.Sleep(5 * time.Second)
		bids, err := s.FindBidsToQueryProvingResult(context.Background())
		if err != nil {
			log.Errorf("FindBidsToQueryProvingResult err: %s", err)
			continue
		}
		for _, bid := range bids {
			proof, err := s.GetProvingResult(bid.AppID, bid.ProofTaskID)
			if err != nil {
				log.Errorf("GetProvingResult %s err: %s", bid.ReqID, err)
				continue
			}
			if len(proof) == 0 {
				log.Infof("GetProvingResult %s proof not generated yet", bid.ReqID)
				continue
			}
			err = s.UpdateBidWithProof(context.Background(), dal.UpdateBidWithProofParams{
				Proof: common.Bytes2Hex(proof),
				ReqID: bid.ReqID,
			})
			if err != nil {
				log.Errorf("UpdateBidWithProof %s err: %s", bid.ReqID, err)
			}
		}
	}
}

func (s *Scheduler) scheduleSubmitProof() {
	for {
		time.Sleep(5 * time.Second)
		bids, err := s.FindBidsToSubmitProof(context.Background())
		if err != nil {
			log.Errorf("FindBidsToSubmitProof err: %s", err)
			continue
		}
		for _, bid := range bids {
			var proof [8]*big.Int
			proofBytes := common.Hex2Bytes(bid.Proof)
			if len(proofBytes) != 10*32 /*first 8 as proof and last 2 as inputs*/ {
				log.Errorf("invalid proof length, reqId %s", bid.ReqID)
				continue
			}
			for i := 0; i < 8; i++ /*first 8 needed only*/ {
				proof[i] = big.NewInt(0).SetBytes(proofBytes[32*i : 32*(i+1)])
			}
			tx, _, err := onchain.TransactWaitSuccess(
				s.Client, s.TransactOpts,
				func(transactor bind.ContractTransactor, opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
					return s.SubmitProof(opts, common.HexToHash(bid.ReqID), proof)
				})
			if err != nil {
				errString := err.Error()
				var jsonErr JsonError
				errJson, _ := json.Marshal(err)
				json.Unmarshal(errJson, &jsonErr)
				if jsonErr.Data != "" {
					errName, _ := ParseSolCustomErrorName(eth.BrevisMarketABI, common.FromHex(jsonErr.Data))
					errString = errString + " - " + errName
				}
				log.Errorf("SubmitProof req %s err: %s", bid.ReqID, errString)
				continue
			}
			log.Infof("SubmitProof req %s tx %s", bid.ReqID, tx.Hash().Hex())
			err = s.UpdateBidAsProofSubmitted(context.Background(), dal.UpdateBidAsProofSubmittedParams{
				ProofSubmitTx: tx.Hash().Hex(),
				ReqID:         bid.ReqID,
			})
			if err != nil {
				log.Errorf("UpdateBidAsProofSubmitted %s err: %s", bid.ReqID, err)
			}
		}
	}
}
