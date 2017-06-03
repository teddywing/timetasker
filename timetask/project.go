package timetask

type Project struct {
	Client   int
	Project  int
	Module   int
	Task     int
	WorkType int `toml:"work_type"`
	Billable bool
}
