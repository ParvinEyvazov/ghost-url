package main

import (
	"syscall/js"

	"github.com/ParvinEyvazov/ghost-url/core"
	"github.com/ParvinEyvazov/ghost-url/url"
)

func main() {
	done := make(chan struct{}, 0)

	js.Global().Set("ghostUrl", js.FuncOf(ghostUrl))
	js.Global().Set("recoverUrl", js.FuncOf(recoverUrl))

	<-done
}

func ghostUrl(this js.Value, args []js.Value) interface{} {

	urlText := (args[0]).String()

	isValid := url.Valid(urlText)

	if !isValid {
		return ""
	}

	encryptedText, err := core.Encrypt(urlText)
	if err != nil {
		return ""
	}

	finalURL := url.Create(encryptedText)

	return finalURL
}

func recoverUrl(this js.Value, args []js.Value) interface{} {

	return "test recover output"
}
