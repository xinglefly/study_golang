package main

import (
	"fmt"
	"studygolang/queue"
)

func main() {
	q := queue.Queue{}

	fmt.Println(q.IsEmpty())
	q.Push(10)
	q.Push(15)
	q.Push(1)
	q.Push(8)
	fmt.Println(q.IsEmpty())

	/*for !q.IsEmpty(){
		fmt.Print(q," ")
	}*/

	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())

	q.Push(5555)
	fmt.Println(q.Pop())
}
