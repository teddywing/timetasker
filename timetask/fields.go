package timetask

type IDName struct {
	ID   uint
	Name string
}

type Client struct {
	IDName
	Projects []Project
}

type Project struct {
	IDName
	Modules   []Module
	Tasks     []Task
	WorkTypes []WorkType `yaml:"work_types"`
}

type Module IDName
type Task IDName
type WorkType IDName

type Fields struct {
	PersonID uint `yaml:"person_id"`
	Clients  []Client
}
