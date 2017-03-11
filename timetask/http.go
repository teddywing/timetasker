package timetask

import (
	// "mime/multipart"
	"net/http"
)

func Login() (resp *http.Response, err error) {
	resp, err = http.Get("https://duckduckgo.com")
	return resp, err
}
