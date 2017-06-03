package timetask

// import "fmt"

// type Client struct {
// 	ID       int
// 	Name     string
// 	Projects []Project
// }

// type Project struct {
// 	ID        int
// 	Name      string
// 	Modules   []Module
// 	Tasks     []Task
// 	WorkTypes []WorkType `yaml:"work_types"`
// }

// type Module struct {
// 	ID   int
// 	Name string
// }
// type Task struct {
// 	ID   int
// 	Name string
// }
// type WorkType struct {
// 	ID   int
// 	Name string
// }
//
// type Fields struct {
// 	PersonID int `yaml:"person_id"`
// 	Clients  []Client
// }
//
// func (f *Fields) ClientByName(client_name string) (*Client, error) {
// 	for _, client := range f.Clients {
// 		if client.Name == client_name {
// 			return &client, nil
// 		}
// 	}
//
// 	return nil, fmt.Errorf("Client %s not found", client_name)
// }
//
// func (c *Client) ProjectByName(project_name string) (*Project, error) {
// 	for _, project := range c.Projects {
// 		if project.Name == project_name {
// 			return &project, nil
// 		}
// 	}
//
// 	return nil, fmt.Errorf("Project %s not found", project_name)
// }
//
// func (p *Project) ModuleByName(module_name string) (*Module, error) {
// 	for _, module := range p.Modules {
// 		if module.Name == module_name {
// 			return &module, nil
// 		}
// 	}
//
// 	return nil, fmt.Errorf("Module %s not found", module_name)
// }
//
// func (p *Project) TaskByName(task_name string) (*Task, error) {
// 	for _, task := range p.Tasks {
// 		if task.Name == task_name {
// 			return &task, nil
// 		}
// 	}
//
// 	return nil, fmt.Errorf("Task %s not found", task_name)
// }
//
// func (p *Project) WorkTypeByName(work_type_name string) (*WorkType, error) {
// 	for _, work_type := range p.WorkTypes {
// 		if work_type.Name == work_type_name {
// 			return &work_type, nil
// 		}
// 	}
//
// 	return nil, fmt.Errorf("Work type %s not found", work_type_name)
// }
