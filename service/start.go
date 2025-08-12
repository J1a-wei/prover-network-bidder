package service

import (
	"flag"
	"fmt"

	"github.com/brevis-network/prover-network-bidder/client"
	"github.com/brevis-network/prover-network-bidder/config"
	"github.com/brevis-network/prover-network-bidder/dal"
	"github.com/brevis-network/prover-network-bidder/onchain"
	"github.com/brevis-network/prover-network-bidder/scheduler"
	"github.com/spf13/viper"
)

var (
	flagConfig = flag.String("config", "config.toml", "config toml file")
)

func Start() error {
	flag.Parse()
	viper.SetConfigFile(*flagConfig)
	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("ReadInConfig err: %w", err)
	}

	// setup db
	dbName := viper.GetString(config.KeyDb)
	db, err := dal.NewDAL(dbName)
	if err != nil {
		return fmt.Errorf("NewDAL err: %w", err)
	}

	var c *config.ChainConfig
	err = viper.UnmarshalKey(config.KeyChain, c)
	if err != nil {
		return fmt.Errorf("UnmarshalKey err: %w", err)
	}

	chainClient, err := onchain.NewChainClient(c, db)
	if err != nil {
		return fmt.Errorf("NewChainClient err: %w", err)
	}
	chainClient.StartMon()

	proverUrl := viper.GetString(config.KeyProverUrl)
	proverClient, err := client.NewProverNetworkClient(proverUrl)
	if err != nil {
		return fmt.Errorf("NewProverNetworkClient err: %w", err)
	}
	scheduler := scheduler.NewScheduler(db, chainClient, proverClient)
	scheduler.Start()

	return nil
}
