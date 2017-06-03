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

var baseURL string = "https://af83.timetask.com/index.php"

func Login(username, password string) (
	resp *http.Response,
	client *http.Client,
	err error,
) {
	cookies, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		return nil, nil, err
	}

	client = &http.Client{Jar: cookies}
	resp, err = client.PostForm(
		baseURL,
		url.Values{
			"module":     {"people"},
			"action":     {"loginsubmit"},
			"f_url":      {"/"},
			"f_username": {username},
			"f_password": {password},
		},
	)
	if err != nil {
		return resp, client, err
	}

	return resp, client, err
}

func SubmitTimeEntry(
	client http.Client,
	time_entry TimeEntry,
) (resp *http.Response, err error) {
	values := buildSubmissionParams(time_entry)

	values.Set("module", "time")
	values.Set("action", "submitmultipletime")

	resp, err = client.PostForm(
		baseURL,
		values,
	)
	if err != nil {
		return resp, err
	}

	return resp, nil
}


func buildSubmissionParams(time_entry TimeEntry) url.Values {
	v := url.Values{}

	v.Set(
		"f_personID0",
		strconv.Itoa(time_entry.PersonID),
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

	v.Set("f_entryIndexes", "0")

	return v
}
