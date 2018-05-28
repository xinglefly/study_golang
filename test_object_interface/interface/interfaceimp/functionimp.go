package interfaceimp

import "studygolang/test_object_interface/model"

type Function struct {
	JobContent *model.Person
}

func (f *Function) SetJobContent(person model.Person, jobContent string) string {
	*f = Function{
		JobContent: &model.Person{
			Id:person.Id,
			Pid:person.Pid,
			Name:person.Name,
			Age:person.Age,
			JobContent: jobContent,
			},
	}

	return "修改成功"
}
