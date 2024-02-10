package main

import (
	"log"
	"net/http"

	"github.com/SaiHtetMyatHtut/potatoverse/handler"
)

func main() {
	http.HandleFunc("/user", handler.AuthenticationHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
