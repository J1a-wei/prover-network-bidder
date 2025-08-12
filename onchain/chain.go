package onchain

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/brevis-network/prover-network-bidder/config"
	"github.com/brevis-network/prover-network-bidder/dal"
	"github.com/brevis-network/prover-network-bidder/eth"
	ethutils "github.com/celer-network/goutils/eth"
	"github.com/celer-network/goutils/eth/mon2"
	"github.com/celer-network/goutils/log"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var ZeroAddr common.Address

type ChainClient struct {
	*config.ChainConfig
	*ethclient.Client
	*ethutils.Transactor
	*mon2.Monitor
	*dal.DAL
	*eth.BrevisMarket
}

func NewChainClient(c *config.ChainConfig, db *dal.DAL) (*ChainClient, error) {
	ec, err := ethclient.Dial(c.ChainRpc)
	if err != nil {
		return nil, fmt.Errorf("dial err: %w", err)
	}
	chid, err := ec.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("ChainID err: %w", err)
	}
	if chid.Uint64() != c.ChainID {
		return nil, fmt.Errorf("chainid mismatch! cfg has %d but onchain has %d", c.ChainID, chid.Uint64())
	}

	bidder, addr, err := createSigner(c.BidderEthAddr, c.BidderKeystore, chid)
	if err != nil {
		return nil, fmt.Errorf("CreateSigner err: %w", err)
	}
	if addr != common.HexToAddress(c.BidderEthAddr) {
		return nil, fmt.Errorf("bidder eth addr mismatch! cfg as %s but %s from keystore", c.BidderEthAddr, addr.String())
	}
	transactor := ethutils.NewTransactorByExternalSigner(
		addr,
		bidder,
		ec,
		big.NewInt(int64(c.ChainID)),
		ethutils.WithBlockDelay(c.BlkDelay),
		ethutils.WithPollingInterval(time.Duration(c.BlkInterval)*time.Second),
		ethutils.WithAddGasEstimateRatio(c.AddGasEstimateRatio),
		ethutils.WithMaxFeePerGasGwei(c.MaxFeePerGasGwei),
		ethutils.WithMaxPriorityFeePerGasGwei(float64(c.MaxPriorityFeePerGasGwei)),
	)
	mon, err := mon2.NewMonitor(ec, db, mon2.PerChainCfg{
		BlkIntv:         time.Duration(c.BlkInterval) * time.Second,
		BlkDelay:        c.BlkDelay,
		MaxBlkDelta:     c.MaxBlkDelta,
		ForwardBlkDelay: c.ForwardBlkDelay,
	})
	if err != nil {
		return nil, fmt.Errorf("NewMonitor err: %w", err)
	}
	brevisMarket, err := eth.NewBrevisMarket(common.HexToAddress(c.BrevisMarketAddr), ec)
	if err != nil {
		return nil, fmt.Errorf("NewBrevisMarket err: %w", err)
	}

	return &ChainClient{c, ec, transactor, mon, db, brevisMarket}, nil
}

// funcs for monitor brevis events
func (c *ChainClient) StartMon() {
	c.monMarket()
}

func (c *ChainClient) monMarket() {
	brevisMarketAddr := common.HexToAddress(c.BrevisMarketAddr)
	if brevisMarketAddr == ZeroAddr {
		return
	}
	go c.MonAddr(mon2.PerAddrCfg{
		Addr:    brevisMarketAddr,
		ChkIntv: time.Duration(c.BlkInterval) * time.Second,
		AbiStr:  eth.BrevisMarketABI,
	}, c.marketCallback)
}

func (c *ChainClient) marketCallback(evname string, elog ethtypes.Log) {
	switch evname {
	case "NewRequest":
		c.handleNewRequest(elog)
	default:
		log.Infoln("unsupported evname: ", evname)
		return
	}
}

func (c *ChainClient) handleNewRequest(eLog ethtypes.Log) {
	ev, err := c.ParseNewRequest(eLog)
	if err != nil {
		log.Errorf("ParseNewRequest err %s, data %+v", err, eLog)
		return
	}
	log.Infof("MonEv -  NewRequest: reqId %x, req %+v", ev.Reqid, ev.Req)

	err = c.SaveApp(context.Background(), dal.SaveAppParams{
		AppID:      common.Bytes2Hex(ev.Req.Vk[:]),
		ImgUrl:     ev.Req.ImgURL,
		Registered: false,
	})
	if err != nil {
		log.Errorf("SaveApp err %s, reqId %x, req %+v", err, ev.Reqid, ev.Req)
		// continue to save proof request, in case app can be saved by other req or manually
	}

	inputData := ""
	for _, i := range ev.Req.InputData {
		inputData += inputData + common.Bytes2Hex(i) + ","
	}

	header, err := c.HeaderByNumber(context.Background(), big.NewInt(int64(eLog.BlockNumber)))
	if err != nil {
		log.Errorf("HeaderByNumber err %s, reqId %x, req %+v", err, ev.Reqid, ev.Req)
		return
	}

	err = c.AddProofRequest(context.Background(), dal.AddProofRequestParams{
		ReqID:              common.Bytes2Hex(ev.Reqid[:]),
		AppID:              common.Bytes2Hex(ev.Req.Vk[:]),
		Nonce:              int64(ev.Req.Nonce),
		PublicValuesDigest: common.Bytes2Hex(ev.Req.PublicValuesDigest[:]),
		InputData:          inputData,
		InputUrl:           ev.Req.InputURL,
		MaxFee:             ev.Req.Fee.MaxFee.String(),
		MinStake:           ev.Req.Fee.MinStake.String(),
		Deadline:           int64(ev.Req.Fee.Deadline),
		CreatedAt:          int64(header.Time),
	})
	if err != nil {
		log.Errorf("AddProofRequest err %s, reqId %x, req %+v", err, ev.Reqid, ev.Req)
	}
}

const awskmsPre = "awskms"

func createSigner(ksfile, passphrase string, chainid *big.Int) (ethutils.Signer, common.Address, error) {
	if strings.HasPrefix(ksfile, awskmsPre) {
		kmskeyinfo := strings.SplitN(ksfile, ":", 3)
		if len(kmskeyinfo) != 3 {
			return nil, ZeroAddr, fmt.Errorf("%s has wrong format", ksfile)
		}
		awskeysec := []string{"", ""}
		if passphrase != "" {
			awskeysec = strings.SplitN(passphrase, ":", 2)
			if len(awskeysec) != 2 {
				return nil, ZeroAddr, fmt.Errorf("%s has wrong format", passphrase)
			}
		}
		kmsSigner, err := ethutils.NewKmsSigner(kmskeyinfo[1], kmskeyinfo[2], awskeysec[0], awskeysec[1], chainid)
		if err != nil {
			return nil, ZeroAddr, err
		}
		return kmsSigner, kmsSigner.Addr, nil
	}
	ksBytes, err := os.ReadFile(ksfile)
	if err != nil {
		return nil, ZeroAddr, err
	}
	key, err := keystore.DecryptKey(ksBytes, passphrase)
	if err != nil {
		return nil, ZeroAddr, err
	}
	signer, err := ethutils.NewSigner(hex.EncodeToString(crypto.FromECDSA(key.PrivateKey)), chainid)
	return signer, key.Address, err
}
