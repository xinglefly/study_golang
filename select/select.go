package main

import (
	"fmt"
	"time"
	"math/rand"
)

//TODO select调度

func general() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func worker(id int, c chan int) {
	for n := range c {
		time.Sleep(1 * time.Second) //TODO 接收时间变长发送时长随机，接收会将数据给冲掉，需要把发送的数据保存起来
		fmt.Printf("worder id %d receiver %d\n", id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func main() {
	var c1, c2 = general(), general()
	var work = createWorker(0)
	var values [] int //保存数据到队列中去
	//n := 0
	//hasValue := false
	tm := time.After(10 * time.Second) //倒计时
	tick := time.Tick(time.Second)
	for {
		var activeWork chan<- int
		var activeValue int
		if len(values) > 0 {
			activeWork = work
			activeValue = values[0]
		}
		//if hasValue {
		//	activeWork = work
		//}

		select {
		case n := <-c1:
			//fmt.Println("Receive c1:", n)
			//w <- n
			//hasValue = true
			values = append(values, n)
		case n := <-c2:
			//fmt.Println("Receive c2：", n)
			//w <- n
			//hasValue = true
			values = append(values, n)

		case activeWork <- activeValue:
			values = values[1:] //TODO 缓存起来，排队消耗
		case <-time.After(800 * time.Millisecond):
			fmt.Println("time out")
		case <-tick:
			fmt.Println("Queue len:", len(values))
		case <-tm:
			fmt.Println("bye bye！")
			return
		}
	}

}
