package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ParvinEyvazov/ghost-url/core"
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
	http.HandleFunc(url.RoutesMap["decryption"], DecryptionHandler)

	log.Println("Listening on http://localhost:3000")
	err := http.ListenAndServe("localhost:3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func EncryptionHandler() http.Handler {

	return http.FileServer(http.Dir("./static"))
}

func DecryptionHandler(w http.ResponseWriter, r *http.Request) {
	setupCORS(&w, r)

	encryptedText := url.Parse(r.URL.String())

	decryptedURL, err := core.Decrypt(encryptedText)

	if err != nil {
		log.Panic("ERROR: Decryption error", err)
	}

	fmt.Printf("go to %T", decryptedURL)

	http.Redirect(w, r, decryptedURL, http.StatusFound)
}

func NewServer() IServer {
	return &Server{}
}

func setupCORS(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
