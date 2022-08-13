package url

import (
	"net/url"
	"strings"
)

const _SPLITTER = "/"

var HOST string = "https://ghost-url.netlify.app"

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

// TODO: should be improved
func Parse(text string) (encrypted string) {

	texts := strings.Split(text, _SPLITTER)

	encrypted = texts[1]

	return
}
