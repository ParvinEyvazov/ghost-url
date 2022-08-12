package main

import (
	"fmt"
	"log"
	"net/http"
)

type IServer interface {
	start()
}

type Server struct {
}

const HOST string = ""

func main() {

	server := NewServer()

	server.start()

	// r.Host

}

func (*Server) start() {
	fs := http.FileServer(http.Dir("./static"))
	fmt.Println()
	http.Handle("/ghost", fs)

	log.Println("Listening on http://localhost:3000/index.html")
	err := http.ListenAndServe("localhost:3000", nil)
	if err != nil {
		log.Fatal(err)
	}

}

func NewServer() IServer {
	return &Server{}
}
