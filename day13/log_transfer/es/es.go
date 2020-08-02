package es

import (
	"context"
	"fmt"
	"strings"

	"github.com/olivere/elastic"
)

// 初始化es， 准别接受kafka数据

var (
	client *elastic.Client
)

// Init ...
func Init(address string) error {

	if !strings.HasPrefix(address, "http") {
		address = "http://" + address
	}
	var err error
	client, err = elastic.NewClient(elastic.SetURL(address))
	if err != nil {
		fmt.Println("new elastic failed, err:", err.Error())
		return err
	}
	fmt.Println("init elastic success!")
	return nil
}

// SendToES ...
func SendToES(indexStr, typeStr string, data interface{}) error {
	// ?
	// p1 := student{Name: "rion", Age: 22, Married: false}
	put1, err := client.Index().
		Index(indexStr).
		Type(typeStr).
		BodyJson(data).
		Do(context.Background())
	if err != nil {
		return err
	}
	fmt.Printf("Indexed user %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
	return err
}
