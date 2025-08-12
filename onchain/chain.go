package monitor

import (
	"time"

	"github.com/brevis-network/prover-network-bidder/config"
	"github.com/brevis-network/prover-network-bidder/eth"
	"github.com/celer-network/goutils/eth/mon2"
	"github.com/celer-network/goutils/log"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var ZeroAddr common.Address

type ChainClient struct {
	*config.Config
	*ethclient.Client
	mon *mon2.Monitor
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
	go c.mon.MonAddr(mon2.PerAddrCfg{
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
