package timetask

import "time"

type TimeEntry struct {
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
