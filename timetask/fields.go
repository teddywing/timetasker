package timetask

type Client struct {
	ID       uint
	Name     string
	Projects []Project
}

type Project struct {
	ID        uint
	Name      string
	Modules   []Module
	Tasks     []Task
	WorkTypes []WorkType `yaml:"work_types"`
}

type Module struct {
	ID   uint
	Name string
}
type Task struct {
	ID   uint
	Name string
}
type WorkType struct {
	ID   uint
	Name string
}

type Fields struct {
	PersonID uint `yaml:"person_id"`
	Clients  []Client
}
