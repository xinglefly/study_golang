package main

import (
	"studygolang/test_object_interface/model"
	"studygolang/test_object_interface/queue_person"
	"fmt"
)

type Content interface {
	SetJobContent(person model.Person, jobContent string) string
}

type Manager interface {
	SetManager(name string, projectName model.Project) string
}

type work interface {
	Content
	Manager
}

var (
	queue_company  = queue_person.Queue{}
	queue_partment = queue_person.Queue{}
	person         = queue_person.Queue{}
)

func queueCompany() {
	queue_company.Push(model.CreateCompany(1, "十二面体", "Internet"))
	queue_company.Push(model.CreateCompany(2, "梧桐树", "Internet"))
	fmt.Println("company:", queue_company.Empty())
	fmt.Println(queue_company.Traverse())
}

func queuePartment() {
	queue_partment.Push(model.CreatePartment(1, 1, "技术部"))
	queue_partment.Push(model.CreatePartment(2, 1, "行政"))
	queue_partment.Push(model.CreatePartment(1, 2, "技术部"))
	fmt.Println("partment:", queue_partment.Empty())
	fmt.Println(queue_partment.Traverse())
}

func AddProject() {
	p := model.CreateProject(1, "tokenHub")
	fmt.Printf("%T %v\n", p, p)
}

func AddPerson() {
	person.Push(model.Person{
		Id:   1,
		Pid:  1,
		Name: "xingle",
	})
	person.Push(model.Person{
		Id:   2,
		Pid:  1,
		Name: "John",
	})
	person.Push(model.Person{
		Id:   3,
		Pid:  1,
		Name: "Kerry",
	})

	println("person:", person.Empty())
	fmt.Println(person.Traverse())
}

func updateProfile(c Content, person model.Person, jobContent string) string {
	return c.SetJobContent(person, jobContent)
}

func main() {
	queueCompany()
	queuePartment()
	AddProject()
	AddPerson()

	//操作
	p := person.Peek()
	fmt.Println(p)

	//c := interfaceimp.Function{}

	//TODO 如何将interface强制转换成Person类型
	//updateProfile(&c, p, "android")
	//
	//person.Traverse()


	c := model.CreateCompany(1, "hello", "golang")
	fmt.Println(c)


}
