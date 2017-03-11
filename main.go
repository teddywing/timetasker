package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v1"
)

type Config struct {
	Auth struct {
		Username    string
		PasswordCmd string `yaml:"password_cmd"`
	}
	Fields struct {
		PersonID uint `yaml:"person_id"`
		Clients  []struct {
			ID       uint
			Name     string
			Projects []struct {
				ID      uint
				Name    string
				Modules []struct {
					ID   uint
					Name string
				}
				Tasks []struct {
					ID   uint
					Name string
				}
				WorkTypes []struct {
					ID   uint
					Name string
				} `yaml:"work_types"`
			}
		}
	}
}

var config Config

func main() {
	loadConfig()
}

func loadConfig() {
	config_str, err := ioutil.ReadFile("config.yml")
	config = Config{}
	err = yaml.Unmarshal(config_str, &config)
	if err != nil {
		log.Println(err)
	}
}
