package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func main() {
	config := sarama.NewConfig()
	// 等待回复ack 的模式 0，1， all 三种
	config.Producer.RequiredAcks = sarama.WaitForAll
	//
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	//
	config.Producer.Return.Successes = true
	msg := &sarama.ProducerMessage{}
	msg.Topic = "web_log"
	msg.Value = sarama.StringEncoder("this is test log~")
	client, err := sarama.NewSyncProducer([]string{"10.10.23.101:9092"}, config)
	if err != nil{
		fmt.Printf("producer closed error:%v\n", err.Error())
		return
	}
	defer client.Close()
	pid, offset, err := client.SendMessage(msg)
	if err != nil{
		fmt.Printf("send masage error:%s\n",err.Error())
		return
	}
	fmt.Printf("pid:%v, offset:%v\n", pid, offset)
}