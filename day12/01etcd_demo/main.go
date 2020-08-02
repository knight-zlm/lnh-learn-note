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
	_, err = cli.Put(ctx, "foo", "ok")
	cancel()
	if err != nil {
		fmt.Printf("put to etcd failed, err:%v\n", err.Error())
		return
	}
	//get
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, "foo")
	if err != nil {
		fmt.Printf("get from etcd failed, err:%v\n", err.Error())
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%v:%v\n", ev.Key, ev.Value)
	}
	// watch 操作
	// rch := cli.Watch(context.Background(), "foo")
	// for wresp := range rch {
	// 	for _, ev := range wresp.Events {
	// 		fmt.Printf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
	// 	}
	// }
}
