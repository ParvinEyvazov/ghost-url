package url

import (
	"net/url"
	"strings"
)

const _SPLITTER = "/"

func Valid(text string) (isValid bool) {
	_, err := url.ParseRequestURI(text)

	if err == nil {
		isValid = true
	}

	return
}

func Create(host, encrypted string) string {

	return host + _SPLITTER + encrypted
}

func Parse(text string) (host, encrypted string) {

	texts := strings.Split(text, _SPLITTER)

	host = texts[0]
	encrypted = texts[1]

	return
}
