package main

import (
	"fmt"
	"sort"
	"strconv"
	"math"
)

type Person struct {
	Name    string
	Age     int
	Created string
}

type PersonSlice [] Person

func (a PersonSlice) Len() int {
	return len(a)
}

func (a PersonSlice) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a PersonSlice) Less(i, j int) bool {
	return a[j].Age < a[i].Age
}

func firstSortAge() {
	p := [] Person{
		{"zhangsan", 12, "2018-5-26 15:3:55"},
		{"lisi", 25, "2018-5-26 15:3:55"},
		{"wangwu", 72, "2018-5-26 15:3:55"},
		{"xingle", 6, "2018-5-26 15:3:55"},
	}

	fmt.Println(p)

	sort.Sort(PersonSlice(p)) //按Age降序排序
	fmt.Println(p)

	sort.Sort(sort.Reverse(PersonSlice(p))) //按照Age 升序排序
	fmt.Println(p)
}

//-----------------抽取wrapper 对Name 和 age分别排序-------------------------------

type PersonWapper struct {
	people [] Person
	by     func(p, q *Person) bool
}

//进一步封装，不对外暴露sort.Sort
type SortBy func(p, q *Person) bool

func SortPerson(people []Person, by SortBy) {
	sort.Sort(PersonWapper{people, by})
}

func (pw PersonWapper) Len() int {
	return len(pw.people)
}

func (pw PersonWapper) Swap(i, j int) {
	pw.people[i], pw.people[j] = pw.people[j], pw.people[i]
}

func (pw PersonWapper) Less(i, j int) bool {
	return pw.by(&pw.people[i], &pw.people[j])
}

func refectoringSort() {
	people := [] Person{
		{"zhangsan", 12, "2018-5-26 15:3:55"},
		{"lisi", 25, "2018-5-26 15:3:55"},
		{"wangwu", 72, "2018-5-26 15:3:55"},
		{"xing", 28, "2018-5-26 15:2:55"},
		{"le", 28, "2018-5-26 15:3:55"},
		{"xingle", 6, "2018-5-26 15:3:55"},
	}

	/*fmt.Println(people)

	sort.Sort(PersonWapper{people, func(p, q *Person) bool {
		return q.Age < p.Age
	}})
	fmt.Println(people)

	sort.Sort(PersonWapper{people, func(p, q *Person) bool {
		return p.Name < q.Name
	}})
	fmt.Println(people)*/

	SortPerson(people, func(p, q *Person) bool {
		return q.Age < p.Age
	})

	SortPerson(people, func(p, q *Person) bool {
		if p.Age == q.Age {
			return p.Created < q.Created
		}
		return false
	})

	fmt.Println(people)

}

func main(){
	refectoringSort()
}

func main1() {
	var orders map[string]Person
	orders = make(map[string]Person)
	orders["test"] = Person{"zhangsan", 25, "2018-5-26 15:3:55"}
	orders["test1"] = Person{"zhang", 20, "2018-5-26 15:3:55"}
	orders["test2"] = Person{"san", 26, "2018-5-26 15:3:55"}
	orders["test3"] = Person{"li", 10, "2018-5-26 15:3:55"}
	orders["test4"] = Person{"si", 15, "2018-5-26 15:3:55"}

	for order := range orders {
		fmt.Println(order, "value ->", orders[order])
	}

	or := []Person{}
	or = append(or, Person{"zhangsan", 25, "2018-5-26 15:3:55"})
	or = append(or, Person{"zhang", 21, "2018-5-26 15:3:55"})
	or = append(or, Person{"san", 22, "2018-5-26 15:3:55"})

	fmt.Println(or)

	f1 := "0.35678"
	f3, _ := strconv.ParseFloat(f1, 64)
	fmt.Println(f3)
	f2, _ := fmt.Printf("%.1f", f3)
	fmt.Println(f2)

	//cost := (t.Price - m.Price)/m.Price
	var f float64
	var ff float64
	f = math.Trunc((81 - 47) / 47)
	ff = (float64(84) - float64(47)) / float64(47)

	trade := ff / float64(47)
	//ff,_ := strconv.ParseFloat(,64)
	fmt.Println(f, ff, trade)

	result := strconv.FormatFloat(ff, 'f', 2, 64)
	fmt.Println("resutlt: ", ff, result)

	fmt.Println(MathQ(81, 67))
}

func MathQ(t, m int) bool {
	//var bound float64 = 0.3
	q := (float64(t) - float64(m)) / float64(m)
	r, _ := strconv.ParseFloat(strconv.FormatFloat(q, 'f', 2, 64), 64)
	fmt.Println(r)
	return math.Min(r, 0.3) == r
}
