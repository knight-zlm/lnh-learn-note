package kafka

import (
	"fmt"
	"github/knight/learn-go/day13/log_transfer/es"

	"github.com/Shopify/sarama"
)

// LogData ...
type LogData struct {
	Data string `json:"data"`
}

// Init ...
func Init(addrs []string, topic string) (err error) {
	// 初始化消费者
	// var consumer sarama.Consumer
	consumer, err := sarama.NewConsumer(addrs, nil)
	if err != nil {
		fmt.Printf("init consumer failed, error:%v\n", err.Error())
		return err
	}
	fmt.Println("consumer:", consumer)
	partitionList, err := consumer.Partitions(topic)
	if err != nil {
		fmt.Println("get partitions failed, err:", err.Error())
		return err
	}
	// fmt.Println(partitionList)
	for partition := range partitionList {
		pc, err := consumer.ConsumePartition("web_log", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Println("consume partition failed, err:", err.Error())
			return err
		}
		// defer pc.AsyncClose()
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d Offset:%d Key:%v Value:%v\n", msg.Partition, msg.Offset, msg.Key, msg.Value)
				// 写给es
				ld := LogData{
					Data: string(msg.Value),
				}
				// err := json.Unmarshal(msg.Value, ld)
				if err != nil {
					fmt.Println("unmarshal data failed, err:", err.Error())
					continue
				}
				es.SendToES(topic, "GO", ld)
			}
		}(pc)
		// fmt.Println(partition)
	}
	return nil
}
