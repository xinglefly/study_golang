package model

import "fmt"

type Company struct {
	Id       int64
	Name     string
	Industry string
}

func (c *Company) String() string {
	return fmt.Sprintf("Company{Id=%d,Name=%s,Industry= %s}", c.Id, c.Name,c.Industry)
}

type Partment struct {
	Id   int64
	Cid  int64
	Name string
}

func (p *Partment) String() string {
	return fmt.Sprintf("Partment{Id=%d,Name=%s}", p.Id, p.Name)
}

type Person struct {
	Id         int64
	Pid        int64
	Name       string
	Age        int64
	JobContent string
	Manager    bool
}

func (p *Person) String() string {
	return fmt.Sprintf("Person{Id=%d,Name=%s,Age=%d}", p.Id, p.Name, p.Age)
}

type Project struct {
	Id   int64
	Name string
}

func (p *Project) String() string {
	return fmt.Sprintf("Project{Id=%d,Name=%s}", p.Id, p.Name)
}

//工厂模式
func CreateCompany(id int64, name, industry string) *Company {
	return &Company{
		Id:       id,
		Name:     name,
		Industry: industry,
	}
}

func CreatePartment(id, cid int64, name string) *Partment {
	return &Partment{
		Id:   id,
		Cid:  cid,
		Name: name,
	}
}

func CreateProject(id int64, name string) *Project {
	return &Project{
		Id:   id,
		Name: name,
	}
}



