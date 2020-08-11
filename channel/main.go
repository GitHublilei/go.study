package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var a []int
var b chan int // 需要指定通道中元素的类型
var wg sync.WaitGroup
var once sync.Once

func noBufChannel() {
	fmt.Println(b)     // 未初始化不生效
	b = make(chan int) // 通道的初始化
	wg.Add(1)
	go func() {
		defer wg.Done()
		x := <-b
		fmt.Println("-------->", x)
	}()
	b <- 10
	fmt.Println("10 in channel...")
	wg.Wait()
}

func bufChannel() {
	fmt.Println(b)         // 未初始化不生效
	b = make(chan int, 16) // 通道的初始化
	b <- 10
	fmt.Println("10 in channel...")
}

// channel练习
func expriChannel() {
	a := make(chan int, 100)
	b := make(chan int, 100)
	wg.Add(3)
	go finput(a)
	go fget(a, b)
	go fget(a, b)
	wg.Wait()
	for ret := range b {
		fmt.Println(ret)
	}
}

// chan<- 单向通道（参数）
func finput(ch1 chan<- int) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	close(ch1)
}

func fget(ch1, ch2 chan int) {
	defer wg.Done()
	// for x := range ch1 {
	// 	ch2 <- x * x
	// }
	for {
		x, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- x * x
	}

	// 确保某个函数只执行一次
	once.Do(func() { close(ch2) })
}

// ------------------------------

func workerMain() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	for w := 0; w < 3; w++ {
		go workerFunc(w, jobs, results)
	}
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)
	for a := 1; a <= 5; a++ {
		<-results
	}
}

func workerFunc(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("worker:%d start job:%d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("worker:%d end job:%d\n", id, j)
		results <- j * 2
	}
}

//---------------------------
// job ...
type calcJob struct {
	x int64
}

// result ...
type calcResult struct {
	calcJob *calcJob
	sum     int64
}

var calcJobChan = make(chan *calcJob, 100)
var calcResultChan = make(chan *calcResult, 100)

func getDataWorker(cb chan<- *calcJob) {
	defer wg.Done()
	// 循环生成int64类型的随机数，发送到calcJobChan
	for {
		x := rand.Int63()
		newJob := &calcJob{
			x: x,
		}
		cb <- newJob
		time.Sleep(time.Millisecond * 500)
	}
}

func doCalcWorker(cb <-chan *calcJob, cr chan<- *calcResult) {
	defer wg.Done()
	// 从calcJobChan中取出随机数计算各位数的和，将结果发送到calcResultChan
	for {
		calcJob := <-cb
		sum := int64(0)
		n := calcJob.x
		for n > 0 {
			sum += n % 10
			n = n / 10
		}
		newResult := &calcResult{
			calcJob: calcJob,
			sum:     sum,
		}

		cr <- newResult
	}
}

func calcSumMain() {
	wg.Add(1)
	go getDataWorker(calcJobChan)
	// 开启24个goroutine执行doCalcWorker执行
	wg.Add(24)
	for i := 0; i < 24; i++ {
		go doCalcWorker(calcJobChan, calcResultChan)
	}

	// 主goroutine从resultChan取出结果并打印到终端输出
	for result := range calcResultChan {
		fmt.Printf("value:%d sum:%d\n", result.calcJob.x, result.sum)
	}
	wg.Wait()
}

// ----------------------

func selectMain() {
	ch := make(chan int, 1)
	for i := 1; i <= 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		}
	}
}

func main() {
	// bufChannel()
	// expriChannel()
	// workerMain()
	// calcSumMain()
	selectMain()
}
