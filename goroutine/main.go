package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func hello(i int) {
	fmt.Println("hello ", i)
}

func randInt() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		r1 := rand.Int()
		r2 := rand.Intn(10)
		fmt.Println(r1, r2)
	}
}

func randIntGo(i int) {
	defer wg.Done()
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(300)))
	fmt.Println(i)
}

var wg sync.WaitGroup

// GOMAXPROCS

func a() {
	defer wg.Done()
	for i := 0; i < 20; i++ {
		fmt.Printf("A:%d\n", i)
	}
}

func b() {
	defer wg.Done()
	for i := 0; i < 20; i++ {
		fmt.Printf("B:%d\n", i)
	}
}

// 程序启动之后会创建一个主goroutine去执行
func main() {
	for i := 0; i < 100; i++ {
		// go hello(i) // 开启一个单独的goroutine去执行hello函数（任务)
	}
	fmt.Println("main")

	// randInt()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go randIntGo(i)
	}

	wg.Wait() // 等待wg的计数器减为0

	// main函数结束了 由main函数启动的goroutine也都结束了
	runtime.GOMAXPROCS(4)
	fmt.Println(runtime.NumCPU(), "<---------12")
	wg.Add(2)
	go a()
	go b()
	wg.Wait()
}
