package timetask

import (
	"flag"
	"io/ioutil"
	"testing"
)

var username, password string

func init() {
	flag.StringVar(&username, "username", "", "Username")
	flag.StringVar(&password, "password", "", "Password")
	flag.Parse()

}

func TestLogin(t *testing.T) {
	response, _ := Login(username, password)
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	t.Log(response)
	t.Logf("%s", body)
}
