package main

import (
	// "fmt"
	// "io/ioutil"
	"log"
	// "os"

	"github.com/teddywing/timetasker/timetask"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Auth struct {
		Username    string
		PasswordCmd string
	}
	Projects map[string]timetask.Project
}

var config Config

func main() {
	loadConfig()

	// if len(os.Args) == 1 {
	// 	fmt.Println("Not enough arguments")
	// 	os.Exit(1)
	// }
	//
	// file_path := os.Args[len(os.Args)-1]
	// file, err := ioutil.ReadFile(file_path)
	// if err != nil {
	// 	log.Println(err)
	// }

	// time_entries := []timetask.TimeEntry{}
	// err = yaml.Unmarshal(file, &time_entries)
	// if err != nil {
	// 	log.Println(err)
	// }
	//
	// log.Printf("%+v", time_entries)

	// timetask.SubmitTimeEntries(config.Fields, time_entries)

	// timetask.GenerateWeeklyTimesheet(os.Stdout, config.Defaults)
}

func loadConfig() {
	config = Config{}
	_, err := toml.DecodeFile("config2.toml", &config)
	if err != nil {
		log.Println(err)
	}
}
