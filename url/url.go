package url

import (
	"net/url"
	"strings"
)

type Routes map[string]string

var RoutesMap Routes = Routes{
	"encryption": "/",
	"decryption": "/d/",
}

var _SPLITTER string = RoutesMap["decryption"]

var HOST string = "http://localhost:3000"

// var HOST string = "https://ghost-url.netlify.app"

func Valid(text string) (isValid bool) {
	_, err := url.ParseRequestURI(text)

	if err == nil {
		isValid = true
	}

	return
}

func Create(encrypted string) string {

	return HOST + _SPLITTER + encrypted
}

func Parse(text string) (encrypted string) {

	texts := strings.Split(text, _SPLITTER)

	encrypted = texts[1]

	return
}
