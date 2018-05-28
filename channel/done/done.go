package main

import (
	"fmt"
	"sync"
)

//打印完毕通知我
//TODO 每次去channel中写入数据在读出数据，导致数据协程顺序执行了，如何解决这种问题呢？
//TODO 代码的重构，一步步的提升如何去包装，如何面向函数式编程
type worker struct {
	in chan int
	//done chan bool   //TODO 这一步是申明一个是否接收完的channel
	//wg *sync.WaitGroup //TODO 这一步是申明一个WaitGroup,指针的方式更改
	done func()		//TODO 把抽象出来
}

func doWorker(id int, worker worker) {
//func doWorker(id int, c chan int, wg *sync.WaitGroup) {
	for n := range worker.in {
		fmt.Printf("worker %d receiver %c\n", id, n)
		//go func() {done <- true}()
		//done <- true  //TODO 这个问题存在于先去写入10个小写字母后执行一次done,然后在执行写入10个大写字母的时候执行导致阻塞，在开一个协程去并发处理就Ok了
		worker.done()
	}
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	work := worker{
		in: make(chan int),
		//done: make(chan bool),
		//wg: wg,
		done: func() { wg.Done() },
	}
	//go doWorker(id, work.in, wg)
	go doWorker(id, work)

	return work
}

func channelDemo() {
	var workers [10]worker
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}

	wg.Add(20)
	//for i := 0; i < 10; i++ {
	for i, worker := range workers {
		worker.in <- 'a' + i
		//<-workers[i].done
	}

	for i, worker := range workers {
		worker.in <- 'A' + i
		//<-workers[i].done
	}

	//for _,worker := range workers{
	//	<-worker.done
	//	<-worker.done
	//}

	wg.Wait()
}

func main() {
	channelDemo()
}
