package timetask

import (
	// "fmt"
	// "log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"
	// "strings"

	"golang.org/x/net/publicsuffix"
)

func Login(username, password string) (resp *http.Response, err error) {
	cookies, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		return nil, err
	}

	client := http.Client{Jar: cookies}
	resp, err = client.PostForm(
		"https://af83.timetask.com/index.php",
		url.Values{
			"module":     {"people"},
			"action":     {"loginsubmit"},
			"f_url":      {"/"},
			"f_username": {username},
			"f_password": {password},
		},
	)
	if err != nil {
		return resp, err
	}

	return resp, err
}

func SubmitTimeEntry(
	project Project,
	time_entry TimeEntry,
) (resp *http.Response, err error) {
}

func SubmitTimeEntries(fields Fields, time_entries []TimeEntry) (resp *http.Response, err error) {
	v := buildSubmissionParams(fields, time_entries)

	v.Set("module", "time")
	v.Set("action", "submitmultipletime")

	return nil, nil
}

func buildSubmissionParams(profile Profile, time_entry TimeEntry) url.Values {
	v := url.Values{}

	v.Set(
		"f_personID0",
		strconv.Itoa(profile.PersonID),
	)

	v.Set(
		"f_clientID0",
		strconv.Itoa(time_entry.Client),
	)

	v.Set(
		"f_projectID0",
		strconv.Itoa(time_entry.Project),
	)

	v.Set(
		"f_moduleID0",
		strconv.Itoa(time_entry.Module),
	)

	v.Set(
		"f_taskID0",
		strconv.Itoa(time_entry.Task),
	)

	v.Set(
		"f_worktypeID0",
		strconv.Itoa(time_entry.WorkType),
	)

	v.Set(
		"f_date0",
		time_entry.Date.Format("02/01/06"), // day/month/year
	)

	v.Set(
		"f_time0",
		strconv.Itoa(time_entry.Time),
	)

	var billable string
	if time_entry.Billable {
		billable = "t"
	} else {
		billable = "f"
	}

	v.Set(
		"f_billable0",
		billable,
	)

	v.Set(
		"f_description0",
		time_entry.Description,
	)

	return v
}
