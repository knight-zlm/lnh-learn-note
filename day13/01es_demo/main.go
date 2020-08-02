package main

import (
	"context"
	"fmt"

	"github.com/olivere/elastic"
)

type student struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Married bool   `json:"married"`
}

func main() {
	// client, err := elastic.NewClient()
	// if err != nil {
	// 	// Handle error
	// 	panic(err)
	// }
	client, err := elastic.NewClient(elastic.SetURL("http://10.10.23.101:9200/"))
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Println("connect to es success")
	p1 := student{Name: "rion", Age: 22, Married: false}
	put1, err := client.Index().
		Index("student").
		Type("go").
		BodyJson(p1).
		Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Indexed user %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
}
