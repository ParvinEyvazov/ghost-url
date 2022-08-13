package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ParvinEyvazov/ghost-url/url"
)

type IServer interface {
	start()
}

type Server struct {
}

func main() {

	server := NewServer()

	server.start()
}

func (*Server) start() {
	fmt.Println()
	http.Handle(url.RoutesMap["encryption"], EncryptionHandler())

	log.Println("Listening on http://localhost:3000")
	err := http.ListenAndServe("localhost:3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func EncryptionHandler() http.Handler {

	return http.FileServer(http.Dir("./static"))
}

func NewServer() IServer {
	return &Server{}
}
