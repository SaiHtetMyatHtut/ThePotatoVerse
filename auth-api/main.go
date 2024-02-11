package main

import (
	"log"
	"net/http"

	"github.com/SaiHtetMyatHtut/potatoverse/handler"
)

func main() {
	http.HandleFunc("/signin", handler.SignIn)
	http.HandleFunc("/signup", handler.SignUp)
	// http.HandleFunc("/user", handler.UserHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
