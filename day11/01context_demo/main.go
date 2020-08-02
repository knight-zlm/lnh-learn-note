package main

import (
	"fmt"
	"sync"
	"time"
)

var notify bool

func f1(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		fmt.Println("ok")
		time.Sleep(time.Millisecond * 500)
		if notify {
			break
		}
	}
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go f1(&wg)
	time.Sleep(time.Second * 5)
	notify = true
	wg.Wait()
}
