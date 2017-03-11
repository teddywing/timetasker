package timetask

import (
	"flag"
	"io/ioutil"
	"strings"
	"testing"
)

var username, password string

func init() {
	flag.StringVar(&username, "username", "", "Username")
	flag.StringVar(&password, "password", "", "Password")
	flag.Parse()

}

func TestLogin(t *testing.T) {
	response, err := Login(username, password)
	if err != nil {
		t.Fatal(err)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(string(body), "<title>Home :: af83</title>") {
		t.Error("Login failed, got ", body)
	}
}
