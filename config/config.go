package config

type ChainConfig struct {
	ChainID          uint64 `mapstructure:"chain_id"`
	ChainRpc         string `mapstructure:"chain_rpc"`
	BlkInterval      uint64 `mapstructure:"blk_interval"`
	BlkDelay         uint64 `mapstructure:"blk_delay"`
	MaxBlkDelta      uint64 `mapstructure:"max_blk_delta"`
	ForwardBlkDelay  uint64 `mapstructure:"forward_blk_delay"`
	BrevisMarketAddr string `mapstructure:"brevis_market_addr"`

	ProverEthAddr       string `mapstructure:"prover_eth_addr"`
	SubmitterKeystore   string `mapstructure:"submitter_keystore"`
	SubmitterPassphrase string `mapstructure:"submitter_passphrase"`
}

// demo purpose, rule params to be determined by business
type RuleConfig struct {
	// prove cycle * prover gas price = estimated cost (bid fee), with a default 1e12 denominator
	ProverGasPrice string `mapstructure:"prover_gas_price"`
	// skip the requests that the calculated bid fee exceeds the `max_fee`
	MaxFee string `mapstructure:"max_fee"`
	// max input size,  default 0 means no limit. if this value is non-zero, and request input is larger, skip request.
	MaxInputSize uint64 `mapstructure:"max_input_size"`
	// accepts only those requests that the duration from prove start time to deadline not less than `ProveMinDuration`
	ProveMinDuration uint64 `mapstructure:"prove_min_duration"`

	// 忽略预检测
	IgnoreEstimateCost bool `mapstructure:"ignore_estimate_cost"`

	// default empty means no limit, if not empty, only process those requests targeted to the whitelist apps
	VkWhitelist []string `mapstructure:"vk_whitelist"`
	// if not empty, skip the requests targeted to the blacklist apps
	VkBlacklist []string `mapstructure:"vk_blacklist"`
}
