package logagent

import (
	"fmt"
	"time"

	"github.com/go.study/logagent/kafka"
	"github.com/go.study/logagent/taillog"
)

func run() {
	// 1.读取日志
	for {
		select {
		case line := <-taillog.ReadChan():
			kafka.SendToKafka("web_log", line.Text)
		default:
			time.Sleep(time.Second)
		}
	}
	// 2.发送到kafka
}

// logAgent 入口程序
func main() {
	// 1.初始化kafka连接
	err := kafka.Init([]string{"127.0.0.1:9092"})
	if err != nil {
		fmt.Printf("init kafka failed, err:%v\n", err)
		return
	}
	// 2. 打开日志文件准备收集日志
	err = taillog.Init("./my.log")
	if err != nil {
		fmt.Printf("Init taillog failed, err:%v\n", err)
		return
	}

	run()
}
