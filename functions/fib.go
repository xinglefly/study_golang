package fib

import (
	"fmt"
	"io"
	"strings"
	"studygolang/functions/fib"
)

//斐波那契数列
//实现原理 向右移一位
// 1 1 2 3 5 8 13 21 ...
//   a b
//     a b
func Fibonacci() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

//优化能不能把fibnaci实现一种接口
type intGen func() int

//函数如何实现接口
func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)      //转换成string
	return strings.NewReader(s).Read(p) //代理reader
}



func main() {
	f := Fibonacci()
	fib.PrintFileContents(f)
	//f() //1
	//f() //1
	//f() //2
	//f() //3
	//f() //5
	//f() //8
	//f() //13
	//f() //21
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
}
