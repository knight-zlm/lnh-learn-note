package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Printf("start up server failed~! %s\n", err.Error())
	}
	conn, err := listener.Accept()
	if err != nil {
		fmt.Printf("accept server failed~! %s\n", err.Error())
	}
	var acce [128]byte
	conn.Read(acce[:])
}
