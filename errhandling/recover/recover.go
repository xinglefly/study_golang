package main

import (
	"fmt"
	"strings"
)

func tryRecover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Printf("Error is : %s", err.Error())
		} else {
			panic(r)
		}
	}()

	//panic(errors.New("this is panic, procedure shut down."))
	b := 0
	a := 5 / b
	fmt.Println(a)
}

func main() {
	tryRecover()

	fmt.Println()

}
