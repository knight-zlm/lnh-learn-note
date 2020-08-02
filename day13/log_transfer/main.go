package main

import (
	"fmt"
	"github/knight/learn-go/day13/log_transfer/conf"
	"github/knight/learn-go/day13/log_transfer/es"
	"github/knight/learn-go/day13/log_transfer/kafka"

	"gopkg.in/ini.v1"
)

func main() {
	var cfg conf.LogTransferCfg
	err := ini.MapTo(&cfg, "./conf/cfg.ini")
	if err != nil {
		fmt.Println("load ini failed, err:", err.Error())
	}
	fmt.Printf("%v\n", cfg)
	// 初始化es
	// 初始化一个esclient
	err = es.Init(cfg.ESCfg.Address)
	if err != nil {
		fmt.Printf("ini es clinet failed, err:%v\n", err.Error())
		return
	}
	// 初始化kafka
	// 1.1 连接kafka，创建分区的消费者
	// 1.2 每个分区的消费者分别取出数据通过sendto将数据发往es
	err = kafka.Init([]string{cfg.KafkaCfg.Address}, cfg.KafkaCfg.Topic)
	if err != nil {
		fmt.Printf("ini kafka failed, err:%v\n", err.Error())
	}
	select {}
}
