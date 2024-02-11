package main

import (
	"log"
	"net/http"

	"github.com/SaiHtetMyatHtut/potatoverse/user-api/handler"
)

func main() {
	http.HandleFunc("/user", handler.UserHandler)
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}
