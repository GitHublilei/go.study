package main

import (
	"fmt"
	"time"

	"github.com/go.study/logagent/conf"

	"github.com/go.study/logagent/kafka"
	"github.com/go.study/logagent/taillog"
	"gopkg.in/ini.v1"
)

var (
	cfg = new(conf.AppConf)
)

func run() {
	// 1.读取日志
	for {
		select {
		case line := <-taillog.ReadChan():
			kafka.SendToKafka(cfg.KafkaConf.Topic, line.Text)
		default:
			time.Sleep(time.Second)
		}
	}
	// 2.发送到kafka
}

// logAgent 入口程序
func main() {
	// 0.加载配置文件
	// cfg, err := ini.Load("./conf/config.ini")
	err := ini.MapTo(cfg, "./conf/config.ini")
	if err != nil {
		fmt.Printf("load ini failed, err:%v\n", err)
		return
	}
	// 1.初始化kafka连接
	err = kafka.Init([]string{cfg.KafkaConf.Address})
	if err != nil {
		fmt.Printf("init kafka failed, err:%v\n", err)
		return
	}
	fmt.Println("init kafka success")
	// 2. 打开日志文件准备收集日志
	err = taillog.Init(cfg.TaillogConf.FileName)
	if err != nil {
		fmt.Printf("Init taillog failed, err:%v\n", err)
		return
	}
	fmt.Println("init taillog success")

	run()
}

// ./kafka-console-consumer.sh --bootstrap-server=127.0.0.1:9092 --topic=web_log --from-beginning
