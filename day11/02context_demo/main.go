package main

import (
	"fmt"
	"sync"
	"time"
)

// var notify bool

func f1(wg *sync.WaitGroup, ch chan struct{}) {
	defer wg.Done()
LOOP:
	for {
		fmt.Println("ok")
		select {
		case <-ch:
			break LOOP
		default:
			time.Sleep(time.Millisecond * 500)
		}
		// if notify {
		// 	break
		// }
	}
}

func main() {
	wg := sync.WaitGroup{}
	ch := make(chan struct{}, 0)
	wg.Add(1)
	go f1(&wg, ch)
	time.Sleep(time.Second * 5)
	ch <- struct{}{}
	wg.Wait()
}
