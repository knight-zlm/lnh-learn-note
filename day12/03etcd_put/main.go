package main

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		fmt.Printf("connect to etcd  failed, err:%v\n", err.Error())
		return
	}
	fmt.Println("connect to etcd success!")
	defer cli.Close()
	// put 操作
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = cli.Put(ctx, "/logagent/collect_config", `[{"path":"D:/temp/etcdlog/redis.log","topic":"db_log"},{"path":"D:/temp/etcdlog/access.log","topic":"web_log"}]`)
	cancel()
	if err != nil {
		fmt.Printf("put to etcd failed, err:%v\n", err.Error())
		return
	}
	fmt.Println("put to etcd sucess!")
}
