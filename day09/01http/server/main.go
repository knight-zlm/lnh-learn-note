package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request) {
	str := "hello 沙河"
	w.Write([]byte(str))
}

func f2(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Query())
	fmt.Println(r.Method)
	defer r.Body.Close()
	fmt.Println(ioutil.ReadAll(r.Body))
	w.Write([]byte("hello"))
}

func main() {
	http.HandleFunc("/posts/Go/", f1)
	http.HandleFunc("/hello/", f2)
	http.ListenAndServe("127.0.0.1:8080", nil)
}
