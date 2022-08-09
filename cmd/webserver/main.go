package main

import (
	"log"
	"net/http"
)

func main() {

	// text := `encrypt me please`

	// encr, err := core.Encrypt(text)
	// if err != nil {
	// 	fmt.Println("Error Encrypting!")
	// } else {
	// 	fmt.Println("BlowFish Encrypted String: " + encr)
	// }

	// decr, err := core.Decrypt(encr)
	// if err != nil {
	// 	fmt.Println("Error Decrypting!")
	// } else {
	// 	fmt.Println("BlowFish Decrypted String: " + decr)
	// }

	// return

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	log.Println("Listening on http://localhost:3000/index.html")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}

}
