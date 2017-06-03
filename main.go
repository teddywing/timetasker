package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/teddywing/timetasker/timetask"

	"gopkg.in/alecthomas/kingpin.v2"
)

var VERSION string = "0.1.0"

var config Config

func main() {
	var err error

	err = loadConfig()
	kingpin.FatalIfError(err, "Could not load config file")

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

	if *project_alias == "" && !*write_config {
		kingpin.Fatalf("required flag --project not provided, try --help")
	}

	if *write_config {
		err = maybeWriteConfig()
		kingpin.FatalIfError(err, "Could not write config file")

		os.Exit(0)
	}
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
			"Date '%s' could not be parsed. Example: -d 2017-01-31\n",
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

	resp, client, err := timetask.Login(
		config.Auth.Username,
		config.Auth.PasswordCmd,
	)
	kingpin.FatalIfError(err, "Login request failed")
	log.Printf("%+v\n", resp)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	log.Println(string(body))

	resp, err = timetask.SubmitTimeEntry(*client, time_entry)
	kingpin.FatalIfError(err, "Time entry submission request failed")
	log.Printf("%+v\n", resp)

	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	log.Println(string(body))
}
