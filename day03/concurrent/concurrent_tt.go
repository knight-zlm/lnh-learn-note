package main

import (
	"fmt"
	"sync"
)

func main() {
	// 并发处理
	ch := make(chan int, 1)
	wg := sync.WaitGroup{}
	wg.Add(2)
	for _, v := range [2]int{1, 2} {
		// , ch chan int, wg *sync.WaitGroup
		go func(index int) {
			defer func() {
				fmt.Printf("gorutine index:%d end\n", index)
				wg.Done()
			}()
			for {
				now, ok := <-ch
				if ok != true {
					fmt.Printf("gorutine index:%d, get now error\n", index)
				}
				fmt.Printf("gorutine index:%d, get now\n", index)
				if now >= 5 {
					fmt.Printf("gorutine index:%d, begin set end\n", index)
					ch <- now
					// runtime.Gosched()
					fmt.Printf("gorutine index:%d, end set end\n", index)
					break
				}
				now++
				fmt.Printf("gorutine index:%d, now is %d\n", index, now)
				ch <- now
				fmt.Printf("gorutine index:%d, set now\n", index)
				// runtime.Gosched()
			}
		}(v)
	}
	ch <- 0
	wg.Wait()
}
