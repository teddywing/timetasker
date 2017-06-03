package main

import (
	"io/ioutil"
	"log"
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
}

func loadConfig() {
	config = Config{}
	_, err := toml.DecodeFile("config2.toml", &config)
	if err != nil {
		log.Println(err)
	}
}
