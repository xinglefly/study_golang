package main

import (
	"sort"
	"fmt"
	"strings"
	"time"
)

type bin int

func (b bin) String() string {
	return fmt.Sprintf("%b", b)
}

func sortInt() {
	intList := [] int{2, 4, 3, 5, 7, 6, 9, 8, 1, 0}
	float8List := [] float64{4.2, 5.9, 12.3, 10.0, 50.4, 99.9, 31.4, 27.81828, 3.14}
	stringList := [] string{"a", "c", "b", "d", "f", "i", "z", "x", "w", "y"}
	//sort.Ints(intList)
	sort.Float64s(float8List)
	sort.Strings(stringList)
	fmt.Println("升序排序")
	fmt.Println(intList)
	fmt.Println(float8List)
	fmt.Println(stringList)
	fmt.Println("降序排序")
	sort.Sort(sort.Reverse(sort.IntSlice(intList)))
	fmt.Printf("%v", intList)
}

func main() {
	//sortInt()
	s := []string{"ma", "c", "x"}
	fmt.Println(strings.Join(s, ","), s)

	fmt.Println(bin(42))

	t := time.Now()

	fmt.Println("now",t)
	fmt.Println(t.Unix())
	fmt.Println(t.UnixNano())
	fmt.Println(time.Now())

	var b = new(time.Time)
	fmt.Println(b)

	fmt.Println(time.Now().Unix())

	tt := time.Unix(time.Now().Unix(), 0)
	fmt.Printf("%d-%d-%d %d:%d:%d\n", tt.Year(), tt.Month(), tt.Day(), tt.Hour(), tt.Minute(), tt.Second())
	strDate := fmt.Sprintf("%d-%d-%d %d:%d:%d\n", tt.Year(), tt.Month(), tt.Day(), tt.Hour(), tt.Minute(), tt.Second())
	// output: 2016-7-27 8:38:19
	fmt.Println(strDate)


}
