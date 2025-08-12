package config

type Config struct {
	ChainID          uint64 `mapstructure:"submit_chain_id"`
	ChainRpc         string `mapstructure:"submit_chain_rpc"`
	BlkInterval      uint64 `mapstructure:"blk_interval"`
	BrevisMarketAddr string `mapstructure:"brevis_market_addr"`

	BidderEthAddr    string `mapstructure:"bidder_eth_addr"`
	BidderKeystore   string `mapstructure:"bidder_keystore"`
	BidderPassphrase string `mapstructure:"bidder_passphrase"`

	MyBidRule MyBidRule `mapstructure:"my_bid_rule"`
}

// demo purpose, rule params to be determined by business
type MyBidRule struct {
	MaxFee string `mapstructure:"max_fee"`
}
