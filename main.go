package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/teddywing/timetasker/timetask"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Auth struct {
		Username    string
		PasswordCmd string `yaml:"password_cmd"`
	}
	Fields   timetask.Fields
	Defaults timetask.TimeEntry
}

var config Config

func main() {
	loadConfig()

	if len(os.Args) == 1 {
		fmt.Println("Not enough arguments")
		os.Exit(1)
	}

	file_path := os.Args[len(os.Args)-1]
	file, err := ioutil.ReadFile(file_path)
	if err != nil {
		log.Println(err)
	}

	time_entries := []timetask.TimeEntry{}
	err = yaml.Unmarshal(file, &time_entries)
	if err != nil {
		log.Println(err)
	}

	log.Printf("%+v", time_entries)

	// timetask.SubmitTimeEntries(config.Fields, time_entries)

	timetask.GenerateWeeklyTimesheet(os.Stdout, config.Defaults)
}

func loadConfig() {
	config_str, err := ioutil.ReadFile("config.yml")
	config = Config{}
	err = yaml.Unmarshal(config_str, &config)
	if err != nil {
		log.Println(err)
	}
}
