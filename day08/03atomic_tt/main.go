package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var i int64

func add(wg *sync.WaitGroup, lk *sync.Mutex) {
	defer wg.Done()
	lk.Lock()
	i++
	lk.Unlock()
}

func add2(wg *sync.WaitGroup) {
	defer wg.Done()
	// i6 := int64(i)
	atomic.AddInt64(&i, 1)
}

func main() {
	lk := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(50000)
	start := time.Now()
	for i := 0; i < 50000; i++ {
		go add(&wg, &lk)
		// go add2(&wg)
	}
	wg.Wait()
	fmt.Printf("sum:%d, take:%v\n", i, time.Now().Sub(start))
}
