package main

import (
	// "fmt"
	"io/ioutil"
	"log"
	// "os"
	"time"

	"github.com/teddywing/timetasker/timetask"

	"github.com/BurntSushi/toml"
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
	loadConfig()

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

	time_entry := timetask.NewTimeEntry(
		config.Profile,
		config.Projects["example"],
		time.Now(),
		7,
		"timetasker test",
	)
	resp, err = timetask.SubmitTimeEntry(*client, time_entry)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%+v\n", resp)

	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	log.Println(string(body))

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
