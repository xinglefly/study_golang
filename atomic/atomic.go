package main

import (
	"fmt"
	"sync"
)

type atomicInt struct{
	Value int
	lock sync.Mutex
}

func (a *atomicInt) increment() {
	a.lock.Lock()
	defer a.lock.Unlock()
	a.Value++
}

func (a *atomicInt) get() int{
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.Value
}

func main(){
	var a atomicInt
	a.increment()
	go func() {
		a.increment()
	}()
	fmt.Println(a.get())
}
