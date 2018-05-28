package main

import "fmt"

func printArr(arr *[5]int) {
	arr[0] = 100
	for _, value := range arr {
		fmt.Println(value)
	}
}

func printSlice(s []int) {
	fmt.Printf("%v,len=%d,cap=%d\n", s, len(s), cap(s))
}

func main() {
	//var arr1 [5]int
	//arr2 := [3]int{1, 3, 5}
	//arr3 := [...]int{2, 4, 6, 8, 10}
	//
	//fmt.Println(arr1, arr2, arr3)
	//
	//printArr(&arr1)
	//printArr(&arr3)
	//
	//fmt.Println(arr1, arr2, arr3)

	//arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8,}
	//
	//s1 := arr[2:6]
	//s2 := s1[3:6]
	//fmt.Println("s1 := ",s1)
	//fmt.Println("s2 := ",s2)
	//fmt.Println("s3 := ",arr[2:])
	//fmt.Println("s4 := ",arr[:])
	//
	//sAppend := append(append(s2, 10), 11)
	//
	//fmt.Println(sAppend,len(sAppend),cap(sAppend))

	//var s []int

	//for i := 0; i < 100; i++ {
	//	printSlice(s)
	//	s = append(s, i*2+1)
	//}
	//fmt.Println(s)

	s1 := []int{1, 3, 5, 7, 9}
	printSlice(s1)

	s2 := make([]int, 10)
	printSlice(s2)
	s3 := make([]int, 8, 15)
	printSlice(s3)

	fmt.Println("copy slice:")
	copy(s2, s1)
	printSlice(s2)

	fmt.Println("Delteing slice [3]")
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)

	fmt.Println("Poping from front:")
	front := s2[0]
	s2 = s2[1:]
	fmt.Println(front)
	printSlice(s2)

	fmt.Println("Poping from back")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	fmt.Println(tail)
	printSlice(s2)

}
