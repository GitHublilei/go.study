package kafka

import (
	"fmt"

	"github.com/Shopify/sarama"
)

// 向kafka写入日志模块

var (
	client sarama.SyncProducer // 声明一个全局的链接kafka的生产者client
)

// Init 初始化kafka
func Init(addrs []string) (err error) {
	config := sarama.NewConfig()
	// tailf包使用
	// 发送完数据需要leader和follow都确认
	config.Producer.RequiredAcks = sarama.WaitForAll
	// 新选出一个partition
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	// 成功交付的消息将在sucess_channel返回
	config.Producer.Return.Successes = true

	// 构造一个消息
	// msg := &sarama.ProducerMessage{}
	// msg.Topic = "web_log"
	// msg.Value = sarama.StringEncoder("this is a test log")
	// 连接kafka
	client, err = sarama.NewSyncProducer(addrs, config)
	if err != nil {
		fmt.Printf("producer close, err:%v\n", err)
		return
	}
	return
}

// SendToKafka 发送消息到kafka
func SendToKafka(topic, data string) {
	// 构造一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.StringEncoder(data)
	// 发送到kafka
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Printf("send msg failed, err:%v\n", err)
		return
	}
	fmt.Printf("pid:%v offset:%v\n", pid, offset)
}
