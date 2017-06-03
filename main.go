package main

import (
	"io/ioutil"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/teddywing/timetasker/timetask"

	"github.com/BurntSushi/toml"
	"gopkg.in/alecthomas/kingpin.v2"
)

type Config struct {
	Auth struct {
		Username    string
		PasswordCmd string `toml:"password_cmd"`
	}
	Profile  timetask.Profile
	Projects map[string]timetask.Project
}

var config Config

func main() {
	var err error

	loadConfig()

	// Parse command line arguments
	project_alias := kingpin.Flag(
		"project",
		"Project alias defined in config.toml.",
	).
		Short('p').
		Required().
		String()
	time_spent := kingpin.Flag("time", "Time spent working on project.").
		Short('t').
		Default("7").
		Int()
	date_str := kingpin.Flag("date", "Date when work was done (e.g. 2017-01-31)").
		Short('d').
		String()
	description := kingpin.Flag("description", "Description of work.").
		Short('m').
		String()
	kingpin.Version("0.1.0")
	kingpin.Parse()

	// Submit time entry
	project, ok := config.Projects[*project_alias]
	if !ok {
		fmt.Printf("Project '%s' not found\n", *project_alias)
		os.Exit(1)
	}

	var date time.Time

	// If the date argument isn't sent, default to today
	if *date_str == "" {
		date = time.Now()
	} else {
		date, err = time.Parse("2006-01-02", *date_str)
		if err != nil {
			fmt.Printf("Date '%s' could not be parsed. Example: -d 2017-01-31\n", *date_str)
			os.Exit(1)
		}
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
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%+v\n", resp)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	log.Println(string(body))

	resp, err = timetask.SubmitTimeEntry(*client, time_entry)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%+v\n", resp)

	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	log.Println(string(body))
}

func loadConfig() {
	config = Config{}
	_, err := toml.DecodeFile("config2.toml", &config)
	if err != nil {
		log.Println(err)
	}
}
