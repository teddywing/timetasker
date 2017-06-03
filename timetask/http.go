package timetask

import (
	// "fmt"
	// "log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	// "strconv"
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

// func SubmitTimeEntries(fields Fields, time_entries []TimeEntry) (resp *http.Response, err error) {
// 	v := buildSubmissionParams(fields, time_entries)
//
// 	v.Set("module", "time")
// 	v.Set("action", "submitmultipletime")
//
// 	return nil, nil
// }
//
// func buildSubmissionParams(fields Fields, time_entries []TimeEntry) url.Values {
// 	v := url.Values{}
// 	entry_indexes := []string{}
//
// 	for i, entry := range time_entries {
// 		entry_indexes = append(entry_indexes, strconv.Itoa(i))
//
// 		client, err := fields.ClientByName(entry.Client)
// 		if err != nil {
// 			log.Panic(err)
// 		}
//
// 		project, err := client.ProjectByName(entry.Project)
// 		if err != nil {
// 			log.Panic(err)
// 		}
//
// 		module, err := project.ModuleByName(entry.Module)
// 		if err != nil {
// 			log.Panic(err)
// 		}
//
// 		task, err := project.TaskByName(entry.Task)
// 		if err != nil {
// 			log.Panic(err)
// 		}
//
// 		work_type, err := project.WorkTypeByName(entry.WorkType)
// 		if err != nil {
// 			log.Panic(err)
// 		}
//
// 		var billable string
// 		if entry.Billable {
// 			billable = "t"
// 		} else {
// 			billable = "f"
// 		}
//
// 		v.Set(
// 			fmt.Sprintf("f_personID%d", i),
// 			strconv.Itoa(fields.PersonID),
// 		)
//
// 		v.Set(
// 			fmt.Sprintf("f_clientID%d", i),
// 			strconv.Itoa(client.ID),
// 		)
//
// 		v.Set(
// 			fmt.Sprintf("f_projectID%d", i),
// 			strconv.Itoa(project.ID),
// 		)
//
// 		v.Set(
// 			fmt.Sprintf("f_moduleID%d", i),
// 			strconv.Itoa(module.ID),
// 		)
//
// 		v.Set(
// 			fmt.Sprintf("f_taskID%d", i),
// 			strconv.Itoa(task.ID),
// 		)
//
// 		v.Set(
// 			fmt.Sprintf("f_worktypeID%d", i),
// 			strconv.Itoa(work_type.ID),
// 		)
//
// 		v.Set(
// 			fmt.Sprintf("f_date%d", i),
// 			entry.Date.Format("02/01/06"), // day/month/year
// 		)
//
// 		v.Set(
// 			fmt.Sprintf("f_time%d", i),
// 			strconv.Itoa(entry.Time),
// 		)
//
// 		v.Set(
// 			fmt.Sprintf("f_billable%d", i),
// 			billable,
// 		)
//
// 		v.Set(
// 			fmt.Sprintf("f_description%d", i),
// 			entry.Description,
// 		)
// 	}
//
// 	v.Set("f_entryIndexes", strings.Join(entry_indexes, ","))
//
// 	return v
// }
