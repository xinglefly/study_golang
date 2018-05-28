package main

import (
	"fmt"
	"time"
	"runtime"
)

//协程 Coroutine
//TODO 协程跟线程的比较： 1.轻量级"线程" 2.非抢占式多任务处理，由协程主动交出控制权 3.编译器、解释器、虚拟机层面的多任务 4.多个协程可能在一个或多个线程上运行
//手动交出控制权 runtime.Gosched()
//top 查看CPU的使用情况
//go run -race gorutine.go 检测程序错误
//子程序是协程的特例

func gorutineDemo() {
	//var a [10] int

	for i := 0; i < 1000; i++ {
		go func(ii int) {
			for {
				//IO的操作会切换，有等待的过程
				fmt.Println("hello,welcome to study golang!", ii)
				//a[ii]++
				//runtime.Gosched()
			}
		}(i)
	}

	time.Sleep(time.Minute)
	//fmt.Println(a)
}

func main() {
	//gorutineDemo()

	cpu := runtime.NumCPU()
	fmt.Println("cpu Number :",cpu)
}
