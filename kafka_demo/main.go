package main

import (
	"fmt"

	"github.com/Shopify/sarama"
)

// kafka demo

// 基于sarama第三方库开发的kafka client
func main() {
	config := sarama.NewConfig()
	// tailf包使用
	// 发送完数据需要leader和follow都确认
	config.Producer.RequiredAcks = sarama.WaitForAll
	// 新选出一个partition
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	// 成功交付的消息将在sucess_channel返回
	config.Producer.Return.Successes = true

	// 构造一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "web_log"
	msg.Value = sarama.StringEncoder("this is a test log")
	// 连接kafka
	client, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		fmt.Printf("producer close, err:%v\n", err)
		return
	}
	defer client.Close()

	// 发送消息
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Printf("send msg failed, err:%v\n", err)
		return
	}
	fmt.Printf("pid:%v offset:%v\n", pid, offset)
}
