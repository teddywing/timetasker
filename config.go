package main

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/teddywing/timetasker/timetask"

	"github.com/BurntSushi/toml"
	"github.com/goulash/xdg"
)

type Config struct {
	Auth struct {
		Username    string
		PasswordCmd string `toml:"password_cmd"`
	}
	Profile  timetask.Profile
	Projects map[string]timetask.Project
}

const emptyConfig = `[auth]
username = ""
password_cmd = ""


[profile]
person_id = # ADD PERSON ID


[projects.example]
client =    # ADD CLIENT ID
project =   # ADD PROJECT ID
module =    # ADD MODULE ID
task = 0
work_type = # ADD WORK TYPE ID
billable = true
`

func configDir() string {
	return filepath.Join(xdg.ConfigHome, "timetasker")
}

func configFile() string {
	return filepath.Join(configDir(), "config.toml")
}

func maybeWriteConfig() error {
	path := xdg.FindConfig("timetasker/config.toml")

	if path == "" {
		path = configDir()
		if _, err := os.Stat(path); os.IsNotExist(err) {
			os.Mkdir(path, 0700)
		}

		config_path := configFile()
		err := ioutil.WriteFile(config_path, []byte(emptyConfig), 0644)
		if err != nil {
			return err
		}
	}

	return nil
}

func loadConfig() error {
	config = Config{}
	_, err := toml.DecodeFile("config2.toml", &config)
	if err != nil {
		return err
	}

	return nil
}
