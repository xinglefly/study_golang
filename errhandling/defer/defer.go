package main

import (
	"fmt"
	"os"
	"bufio"
	fib2 "studygolang/functions"
)

//多个defer存放在栈里面，先进后出
func print() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)

	fmt.Println("")
	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		if i >= 30 {
			panic("printed too many.")
		}
	}
}

func writerFile(name string) {
	file, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	write := bufio.NewWriter(file)
	defer write.Flush()

	f := fib2.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(write, f())
	}
}

func printPanicErr(name string) {
	file, err := os.Open(name)

	//err = errors.New("this is a customer error !") 自定义error
	if err != nil {
		if pathError, ok := err.(*os.PathError); ok {
			fmt.Printf("op=%s, pth=%s, err=%s", pathError.Op, pathError.Path, pathError.Err)
		} else {
			panic(err)
		}
		return
	}
	defer file.Close()
}

func main() {
	//print()
	writerFile("fib.txt")

	printPanicErr("abc.txt")
}
