package main

import (
	"fmt"
	"strconv"
	"sync"
)

var m = make(map[string]int)

func get(key string) int {
	return m[key]
}

func set(key string, value int) {
	m[key] = value
}
func main() {
	// var mp = make(map[string]int)
	wg := sync.WaitGroup{}
	for i := 0; i < 23; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			// k, ok := mp["ok"]
			// if !ok {
			// 	fmt.Println(ok)
			// }
			// fmt.Println("k:", k)
			key := strconv.Itoa(n)
			set(key, n)
			// mp[key] = k + 1
			fmt.Printf("k=%v, v=%v\n", key, get(key))
		}(i)
	}
	wg.Wait()
}
