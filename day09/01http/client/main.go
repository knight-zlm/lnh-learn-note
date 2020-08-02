package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("http://127.0.0.1:8080/hello/?name=sb")
	if err != nil {
		fmt.Println("get error:", err.Error())
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read body error", err.Error())
	}
	fmt.Println(string(b))
}
