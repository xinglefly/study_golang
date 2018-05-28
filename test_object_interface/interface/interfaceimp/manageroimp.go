package interfaceimp

import "studygolang/test_object_interface/model"

type manager struct {
	Name        string
	ProjectName model.Project
}

func (m *manager) SetManager(name string, projectName model.Project) string {
	if len(name) == 0 && len(projectName.Name) == 0 {
		panic("Input name and Project is not empty!!")
	}
	*m = manager{Name: name, ProjectName: projectName}
	return "setting ok"
}
