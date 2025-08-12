package config

type ChainConfig struct {
	ChainID                  uint64  `mapstructure:"submit_chain_id"`
	ChainRpc                 string  `mapstructure:"submit_chain_rpc"`
	BlkInterval              uint64  `mapstructure:"blk_interval"`
	BlkDelay                 uint64  `mapstructure:"blk_delay"`
	MaxBlkDelta              uint64  `mapstructure:"max_blk_delta"`
	ForwardBlkDelay          uint64  `mapstructure:"forward_blk_delay"`
	AddGasEstimateRatio      float64 `mapstructure:"add_gas_estimate_ratio"`
	MaxFeePerGasGwei         uint64  `mapstructure:"max_fee_per_gas_gwei"`
	MaxPriorityFeePerGasGwei float64 `mapstructure:"max_priority_fee_per_gas_gwei"`
	BrevisMarketAddr         string  `mapstructure:"brevis_market_addr"`

	BidderEthAddr    string `mapstructure:"bidder_eth_addr"`
	BidderKeystore   string `mapstructure:"bidder_keystore"`
	BidderPassphrase string `mapstructure:"bidder_passphrase"`
}

// demo purpose, rule params to be determined by business
type RuleConfig struct {
	MaxFee string `mapstructure:"max_fee"`
}
