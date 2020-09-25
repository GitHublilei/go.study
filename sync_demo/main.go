package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var x = 0
var wg sync.WaitGroup
var lock sync.Mutex
var rwLock sync.RWMutex

func add() {
	defer wg.Done()
	for i := 0; i < 500000; i++ {
		lock.Lock()
		x = x + 1
		lock.Unlock()
	}
}

func mutexLock() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}

// -----------------

var (
	y = 0
)

func read() {
	defer wg.Done()
	// lock.Lock()
	rwLock.RLock()
	fmt.Println(y)
	time.Sleep(time.Millisecond)
	// lock.Unlock()
	rwLock.RUnlock()
}

func write() {
	defer wg.Done()
	rwLock.Lock()
	y = y + 1
	time.Sleep(time.Millisecond * 5)
	rwLock.Unlock()
}

func lockMain() {
	start := time.Now()
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go write()
	}
	time.Sleep(time.Second)
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start))
}

// -------------------

// Go内置的map并发不是安全的

var m = make(map[string]int)
var m1 = sync.Map{}

func getM(key string) int {
	return m[key]
}

func setM(key string, value int) {
	m[key] = value
}

func mapMain() {
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m1.Store(key, n)
			value, _ := m1.Load(key)
			fmt.Printf("k=%v v=%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func main() {
	// mutexLock()
	// lockMain()
	mapMain()
}
