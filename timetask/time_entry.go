package timetask

import "time"

type TimeEntry struct {
	PersonID    int
	Client      int
	Project     int
	Module      int
	Task        int
	WorkType    int
	Date        time.Time
	Time        int
	Billable    bool
	Description string
}

func NewTimeEntry(
	profile Profile,
	project Project,
	date time.Time,
	time int,
	description string,
) TimeEntry {
	return TimeEntry{
		PersonID:    profile.PersonID,
		Client:      project.Client,
		Project:     project.Project,
		Module:      project.Module,
		Task:        project.Task,
		WorkType:    project.WorkType,
		Date:        date,
		Time:        time,
		Billable:    project.Billable,
		Description: description,
	}
}
