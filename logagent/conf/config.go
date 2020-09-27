package conf

// AppConf AppConf
type AppConf struct {
	KafkaConf   `ini:"kafka"`
	TaillogConf `ini:"taillog"`
}

// KafkaConf ...
type KafkaConf struct {
	Address string `ini:"address"`
	Topic   string `ini:"topic"`
}

// TaillogConf ...
type TaillogConf struct {
	FileName string `ini:"path"`
}
