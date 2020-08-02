package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func product(jobChan chan<- int64, wg *sync.WaitGroup) {
	defer close(jobChan)
	defer wg.Done()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		jobChan <- rand.Int63()
	}
}

func consumer(jobChan <-chan int64, resultChan chan<- string, wg *sync.WaitGroup) {
	defer close(resultChan)
	defer wg.Done()
	for i := range jobChan {
		// 计算每一位的数的和
		var sum int64
		raw := strconv.Itoa(int(i))
		for i > 0 {
			mod := i % 10
			sum += mod
			i /= 10
		}
		resultChan <- fmt.Sprintf("%s,sum:%d", raw, sum)
	}
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	jobChan := make(chan int64, 100)
	resultChan := make(chan string, 100)
	go product(jobChan, &wg)
	go consumer(jobChan, resultChan, &wg)
	wg.Wait()
	for i := range resultChan {
		fmt.Println(i)
	}
}
