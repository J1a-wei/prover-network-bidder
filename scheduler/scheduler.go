package scheduler

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/brevis-network/prover-network-bidder/client"
	"github.com/brevis-network/prover-network-bidder/config"
	"github.com/brevis-network/prover-network-bidder/dal"
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
	var ruleConfig *config.RuleConfig
	err := viper.UnmarshalKey(config.KeyRule, c)
	if err != nil {
		return nil, fmt.Errorf("UnmarshalKey err: %w", err)
	}
	return &Scheduler{db, c, p, ruleConfig}, nil
}

func (s *Scheduler) Start() {
	go s.scheduleAppRegister()
	go s.scheduleBid()
	go s.scheduleReveal()
	go s.scheduleProve()
}

func (s *Scheduler) scheduleAppRegister() {
	for {
		time.Sleep(5 * time.Second)
		apps, err := s.FindNotRegisteredApps(context.Background())
		if err != nil {
			log.Errorf("FindNotRegisteredApps err: %s", err)
			continue
		}
		var elf []byte
		// TODO: retrieve elf from img url
		for _, app := range apps {
			err = s.RegisterApp(app.AppID, "", elf)
			if err != nil {
				log.Errorf("RegisterApp %s err: %s", app.AppID, err)
				continue
			}

			err = s.UpdateAppAsRegistered(context.Background(), app.AppID)
			if err != nil {
				log.Errorf("UpdateAppAsRegistered %s err: %s", app.AppID, err)
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
		var inputs [][]byte
		// TODO: get inputs from req.InputData or req.InputUrl
		for _, req := range reqs {
			cost, pvDigest, err := s.EstimateCost(req.AppID, inputs)
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

			// compare cost to rule, demo logic, subject to be changed later
			// TODO
			myFee := big.NewInt(0).Mul(big.NewInt(1e18), big.NewInt(int64(cost)))
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
			nonce := big.NewInt(time.Now().Unix())
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
				log.Errorf("Bid err: %s", err)
				if strings.Contains(err.Error(), "bidding phase ended") || strings.Contains(err.Error(), "request does not exist") {
					err = s.UpdateRequestAsProcessed(context.Background(), req.ReqID)
					if err != nil {
						log.Errorf("UpdateRequestAsProcessed %s err: %s", req.ReqID, err)
					}
				}
				continue
			}
			log.Infof("Bid tx %s", tx.Hash().Hex())
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
				log.Errorf("Reveal err: %s", err)
				continue
			}
			log.Infof("Reveal tx %s", tx.Hash().Hex())
			err = s.UpdateBidAsRevealed(context.Background(), bid.ReqID)
			if err != nil {
				log.Errorf("UpdateBidAsRevealed %s err: %s", bid.ReqID, err)
			}
		}
	}
}

func (s *Scheduler) scheduleProve() {
	for {
		time.Sleep(5 * time.Second)
		// TODO
	}
}
