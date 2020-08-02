package main

import (
	"flag"
	"fmt"
)

func main() {
	var name string
	flag.StringVar(&name, "name", "zlm", "获得名字")
	flag.Parse()
	fmt.Printf("name:%v\n", name)
}
