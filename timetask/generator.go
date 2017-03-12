package timetask

import (
	"io"
	"log"
	"text/template"
	"time"

	"github.com/olebedev/when"
	"github.com/olebedev/when/rules/common"
	"github.com/olebedev/when/rules/en"
)

func GenerateWeeklyTimesheet(wr io.Writer, defaults TimeEntry) {
	w := when.New(nil)
	w.Add(en.All...)
	w.Add(common.All...)

	monday, err := w.Parse("last monday", time.Now())
	if err != nil {
		log.Panic(err)
	}

	time_entries := []TimeEntry{}
	day := monday.Time
	for i := 1; i <= 5; i++ {
		time_entries = append(time_entries, defaults)
		time_entries[len(time_entries) - 1].Date = day
		day = day.AddDate(0, 0, 1) // Add 1 day
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
