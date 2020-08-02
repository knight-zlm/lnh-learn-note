package conf

// LogTransferCfg ...
type LogTransferCfg struct {
	KafkaCfg `ini:"kafka"`
	ESCfg    `ini:"es"`
}

// KafkaCfg ...
type KafkaCfg struct {
	Address string `ini:"address"`
	Topic   string `ini:"topic"`
}

// ESCfg ...
type ESCfg struct {
	Address string `ini:"address"`
}
