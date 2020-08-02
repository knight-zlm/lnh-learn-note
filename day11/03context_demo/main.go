package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// var notify bool

func f1(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
LOOP:
	for {
		fmt.Println("ok")
		select {
		case <-ctx.Done():
			break LOOP
		default:
			time.Sleep(time.Millisecond * 500)
		}
		// if notify {
		// 	break
		// }
	}
}

func fDeadline(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
LOOP:
	for {
		fmt.Println("ok")
		select {
		case <-ctx.Done():
			fmt.Println("func dealine")
			break LOOP
		case <-time.After(1 * time.Second):
			fmt.Println("overslept")
		default:
			time.Sleep(time.Millisecond * 500)
		}
		// if notify {
		// 	break
		// }
	}
}

// TraceCode ...
type TraceCode string

func work(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	key := TraceCode("TraceCode")
	traceCode, ok := ctx.Value(key).(string)
	if !ok {
		fmt.Println("invalid trace code~")
	}
LOOP:
	for {
		fmt.Printf("worker, trace code:%s\n", traceCode)
		time.Sleep(10 * time.Millisecond)
		select {
		case <-ctx.Done():
			break LOOP
		default:
		}
	}
}

func main() {
	wg := sync.WaitGroup{}
	// ctx, cancel := context.WithCancel(context.Background())
	// d := time.Now().Add(50 * time.Millisecond)
	// ctx, cancel := context.WithDeadline(context.Background(), d)
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	ctx = context.WithValue(ctx, TraceCode("TraceCode"), "9527")
	wg.Add(1)
	// go f1(ctx, &wg)
	// go fDeadline(ctx, &wg)
	go work(ctx, &wg)
	time.Sleep(time.Second * 5)
	cancel()
	wg.Wait()
}
