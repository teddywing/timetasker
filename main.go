package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/teddywing/timetasker/timetask"

	"gopkg.in/alecthomas/kingpin.v2"
)

var VERSION string = "0.1.0"

var config Config

func main() {
	var err error

	// Parse command line arguments
	project_alias := kingpin.Flag(
		"project",
		"Project alias defined in config.toml.",
	).
		Short('p').
		String()
	time_spent := kingpin.Flag("time", "Time spent working on project.").
		Short('t').
		Default("7").
		Float()
	date_str := kingpin.Flag("date", "Date when work was done (e.g. 2017-01-31)").
		Short('d').
		String()
	description := kingpin.Flag("description", "Description of work.").
		Short('m').
		String()
	write_config_description := fmt.Sprintf(
		"Initialise a new config file template at %s",
		configFile(),
	)
	write_config := kingpin.Flag("write-config", write_config_description).
		Bool()
	kingpin.Version(VERSION)
	kingpin.Parse()

	// Error if no --project unless --write-config was passed
	if *project_alias == "" && !*write_config {
		kingpin.Fatalf("required flag --project not provided, try --help")
	}

	if *write_config {
		err = maybeWriteConfig()
		kingpin.FatalIfError(err, "could not write config file")

		os.Exit(0)
	}

	err = loadConfig()
	kingpin.FatalIfError(err, "could not load config file, try --write-config")

	// Submit time entry
	project, ok := config.Projects[*project_alias]
	if !ok {
		kingpin.Errorf("project '%s' not found", *project_alias)
		os.Exit(1)
	}

	var date time.Time

	// If the date argument isn't sent, default to today
	if *date_str == "" {
		date = time.Now()
	} else {
		date, err = time.Parse("2006-01-02", *date_str)
		kingpin.FatalIfError(
			err,
			"date '%s' could not be parsed. Example: -d 2017-01-31\n",
			*date_str,
		)
	}

	time_entry := timetask.NewTimeEntry(
		config.Profile,
		project,
		date,
		*time_spent,
		*description,
	)

	password, err := passwordCmd(config.Auth.PasswordCmd)
	kingpin.FatalIfError(err, "password command failed")

	resp, client, err := timetask.Login(
		config.Auth.Username,
		password,
	)
	kingpin.FatalIfError(err, "login request failed")

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if strings.Contains(
		string(body),
		"The username and password don't appear to be valid.",
	) {
		kingpin.Errorf("TimeTask authentication failed")
		os.Exit(1)
	}

	resp, err = timetask.SubmitTimeEntry(*client, time_entry)
	kingpin.FatalIfError(err, "time entry submission request failed")

	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if strings.Contains(
		string(body),
		"No time entries were created.",
	) {
		kingpin.Errorf("time entry creation failed")
		os.Exit(1)
	}
}
