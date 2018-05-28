package main

import (
	"fmt"
	_ "math/cmplx"
	"math"
	"math/cmplx"
	"io/ioutil"
	"strconv"
	"os"
	"bufio"
	"reflect"
	"runtime"
)

/**
1.Go语言里面的常量不用大写
 */

//强制类型转换
func TypeChange() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))

	fmt.Println(c)
}

//欧拉公式
func euler() {
	//复数类型， 实数和虚数
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))

	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)

	fmt.Printf("%.3f", cmplx.Exp(1i*math.Pi)+1)
}

func consts() {
	const filename = "abc.txt"
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(c, filename)
}

func enum() {
	const (
		cpp    = iota
		java
		python
		php
	)

	const (
		b  = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	println(cpp, java, python, php)
	println(b, kb, mb, gb, tb, pb)
}

func readFile() {
	const filename = "abc.txt"
	/*contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Printf("%s \n",contents)
	}*/

	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("%s", contents)
	}
}

//不需要每次都写break,默认就有
func mathOperation(a, b int, oper string) int {
	var result int
	switch oper {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("unsupported operate: " + oper)
	}
	fmt.Println(result)
	return result
}

//switch可以没有表达式
func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("wrong score: %d", score))
	case score < 60:
		g = "D"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}

func forever() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func converToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}

	return result
}

func foreach() {
	for {
		fmt.Println("abc")
	}
}

//另一种方式读取file文件
func printFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b) //不想要第二个参数,省略变量
		return q, nil
	default:
		return 0, fmt.Errorf("unsupported operation: %s", op)
	}
}

//函数式编程，代替上面的方法
func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Println("Calling function %s with args"+"(%d,%d)", opName, a, b)
	return op(a, b)
}

func div(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return q, r
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func swap(a, b int) (int, int) {
	return b, a
}

func main() {
	fmt.Println("Test int")
	//TypeChange()
	//euler()
	//consts()
	//enum()
	//readFile()
	//mathOperation(10, 20, "ww")

	//println(
	//	grade(59),
	//	grade(68),
	//	grade(77),
	//	grade(99),
	//	grade(88),
	//)

	//forever()
	fmt.Println(
		converToBin(3),  //011
		converToBin(13), //1011 --> 1101反转
		converToBin(342),
		converToBin(19),
	)

	//foreach()

	//printFile("abc.txt")

	q, r := div(10, 11)
	fmt.Println(q, r)
	fmt.Println(eval(3, 4, "X"))

	if result, err := eval(3, 4, "|"); err != nil {
		fmt.Println("Error")
	} else {
		fmt.Println(result)
	}

	//调用 重写方法pow
	fmt.Println(apply(pow, 3, 4))
	//使用匿名方法
	fmt.Println("匿名方法", apply(func(a int, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 3, 4))

	fmt.Println(sum(1, 2, 3, 4, 5))

	a, b := swap(3, 4)

	fmt.Println(a, b)

}
