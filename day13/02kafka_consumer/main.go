package main

import (
	"fmt"

	"github.com/Shopify/sarama"
)

func main() {
	consumer, err := sarama.NewConsumer([]string{"10.10.23.101:9092"}, nil)
	if err != nil {
		fmt.Println("new consumer failed, err:", err.Error())
		return
	}
	fmt.Println(consumer)
	partitionList, err := consumer.Partitions("web_log")
	if err != nil {
		fmt.Println("get partitions failed, err:", err.Error())
		return
	}
	fmt.Println(partitionList)
	for partition := range partitionList {
		pc, err := consumer.ConsumePartition("web_log", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Println("consume partition failed, err:", err.Error())
		}
		defer pc.AsyncClose()
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d Offset:%d Key:%v Value:%v\n", msg.Partition, msg.Offset, msg.Key, msg.Value)
			}
		}(pc)
		// fmt.Println(partition)
	}
	// 需要阻塞的地方
	select {}
}
