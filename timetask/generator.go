package timetask

import (
	"io"
	"log"
	"text/template"
)

func GenerateWeeklyTimesheet(wr io.Writer, defaults TimeEntry) {
	time_entries := []TimeEntry{}
	for i := 1; i <= 5; i++ {
		time_entries = append(time_entries, defaults)
	}

	t, err := template.ParseFiles(
		"templates/weekly_timesheet.yml.tmpl",
		"templates/timesheet.yml.tmpl",
	)
	if err != nil {
		log.Panic(err)
	}

	err = t.Execute(wr, time_entries)
	if err != nil {
		log.Panic(err)
	}
}
