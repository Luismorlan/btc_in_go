package config

// This is the global app config for the blockchain.
type AppConfig struct {
	// How many leading 0s to form a valid hash.
	DIFFICULTY int `yaml:"DIFFICULTY"`
	// The default coinbase reward.
	COINBASE_REWARD float64 `yaml:"COINBASE_REWARD"`
	// How deep a block is confirmed. Aka how many block need to be after this block to confirm a block.
	CONFIRMATION int `yaml:"CONFIRMATION"`
}
