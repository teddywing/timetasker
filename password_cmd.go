package main

import (
	"os"
	"os/exec"
)

// Execute the given string as a shell command and return the resulting output
func passwordCmd(password_cmd string) (password string, err error) {
	shell := os.Getenv("SHELL")

	// `Command` requires us to pass shell arguments as parameters to the
	// function, but we don't know what the arguments are because
	// `password_cmd` is an arbitrary command. To get around this, we pass the
	// password command to the current shell to execute.
	output, err := exec.Command(shell, "-c", password_cmd).Output()
	if err != nil {
		return "", err
	}

	return string(output), nil
}
