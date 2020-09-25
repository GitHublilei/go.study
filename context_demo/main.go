package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func f(ctx context.Context) {
	defer wg.Done()
	go f1(ctx)
LOOP:
	for {
		fmt.Println("悟空")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-ctx.Done():
			break LOOP
		default:
		}
	}
}

func f1(ctx context.Context) {
LOOP:
	for {
		fmt.Println("八戒")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-ctx.Done():
			break LOOP
		default:
		}
	}
}

func withDeadLine() {
	d := time.Now().Add(2000 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	// 尽管ctx会过期，但是在任何情况下调用它的cancel函数都是很好的实践
	// 如果不这样做，可能会使上下文及其父类存活的时间超过必要的时间
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("over sleep")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

func main() {
	// ctx, cancel := context.WithCancel(context.Background())
	// wg.Add(1)
	// go f(ctx)
	// time.Sleep(time.Second * 5)
	// cancel()
	// wg.Wait()
	withDeadLine()
}
