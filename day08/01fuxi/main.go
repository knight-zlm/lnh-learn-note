package main

import (
	"fmt"
	"sync"
)

func worker(index int, job <-chan int, result chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range job {
		fmt.Printf("index:%d, finish:%d\n", index, i)
		result <- i * i
	}
}

func main() {
	job := make(chan int, 5)
	result := make(chan int, 5)
	wg := sync.WaitGroup{}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go worker(i, job, result, &wg)
	}
	for i := 0; i < 10; i++ {
		job <- i
	}
	close(job)
	go func() {
		wg.Wait()
		close(result)
	}()
	for r := range result {
		fmt.Println(r)
	}
}
