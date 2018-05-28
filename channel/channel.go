package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
	for {
		n := <-c
		fmt.Printf("Worker %d receiver %c\n", id, n)
	}
}

func workerClose(id int, c chan int) {
	/*for {
		n,ok := <- c
		if !ok {
			break
		}
		fmt.Printf("Worker %d receiver %d\n", id, n)
	}*/

	for n := range c {
		fmt.Printf("Worker %d receiver %d\n", id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	/*go func() {
		for {
			fmt.Printf("Worker %d receiver %c\n", id, <-c)
		}
	}()*/
	go worker(id, c)
	return c
}

func chanDemo() {
	//var c chan int //这里申明的channel 为 nil
	c := make(chan int)
	go worker(0, c) //可以作为方法提取出来
	c <- 1
	c <- 2

	time.Sleep(time.Millisecond)

	//TODO all goroutines are asleep - deadlock! err :造成的原因是因为实在goruntine中接受数据的，所以会造成死锁
	//n := <-c
	//fmt.Println(n)
}

func chanDemoList() {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		//channels[i] = make(chan int)
		//go worker(i, channels[i])
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}

	time.Sleep(time.Millisecond)
}

func bufferChannel() {
	c := make(chan int, 3)

	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'

	time.Sleep(time.Millisecond)
}

func channelClose() {
	c := make(chan int, 3)

	go workerClose(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)
	time.Sleep(time.Millisecond)
}

func main() {
	fmt.Println("Channel is first-class citizen")
	chanDemo()
	//chanDemoList()
	fmt.Println("Buffered channel")
	//bufferChannel()
	fmt.Println("Channel close as range")
	channelClose()
}

