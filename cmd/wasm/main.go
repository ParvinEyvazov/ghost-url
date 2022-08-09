package main

import (
	"syscall/js"
)

func main() {
	done := make(chan struct{}, 0)

	js.Global().Set("ghostUrl", js.FuncOf(ghostUrl))
	js.Global().Set("recoverUrl", js.FuncOf(recoverUrl))

	<-done
}

func ghostUrl(this js.Value, args []js.Value) interface{} {

	return "test output"
}

func recoverUrl(this js.Value, args []js.Value) interface{} {

	return "test recover output"
}
