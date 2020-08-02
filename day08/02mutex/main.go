package main

import (
	"fmt"
	"sync"
	"time"
)

// func write(i *int, l *sync.RWMutex){
func write(i *int, l *sync.Mutex){
	// defer l.Unlock()
	l.Lock()
	(*i) ++
	time.Sleep(time.Millisecond*5)
	l.Unlock()
}

// func read(i int, l *sync.RWMutex, wg *sync.WaitGroup){
func read(i int, l *sync.Mutex, wg *sync.WaitGroup){
	defer wg.Done()
	// defer l.RUnlock()
	l.Lock()
	// l.RLock()
	fmt.Println(i)
	time.Sleep(time.Millisecond)
	l.Unlock()
	// l.RUnlock()
}

func main() {
	golob := 0
	// lk := sync.RWMutex{}
	lk := sync.Mutex{}
	start := time.Now()
	for i:=0;i<50;i++{
		go write(&golob, &lk)
	}
	time.Sleep(time.Second)
	wg := sync.WaitGroup{}
	for i:= 0;i<1000;i++{
		wg.Add(1)
		go read(golob, &lk, &wg)
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start))
}