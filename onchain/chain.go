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

	return &ChainClient{c, ec, transactor, mon, db}, nil
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
	// TODO
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
