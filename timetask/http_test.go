package timetask

import (
	"io/ioutil"
	"testing"
)

func TestLogin(t *testing.T) {
	response, _ := Login()
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	t.Log(response)
	t.Logf("%s", body)
}
