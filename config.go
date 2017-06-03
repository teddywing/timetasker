package main

import (
	"os"
	"path/filepath"

	"github.com/goulash/xdg"
)

func MaybeWriteConfig() {
	path := xdg.FindConfig("timetasker/config.toml")

	if path == "" {
		path = filepath.Join(xdg.ConfigHome, "timetasker")
		if _, err := os.Stat(path); os.IsNotExist(err) {
			os.Mkdir(path, 0700)
		}
	}
}
