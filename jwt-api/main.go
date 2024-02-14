package main

import (
	"log"
	"net/http"

	"github.com/SaiHtetMyatHtut/potatoverse/jwt-api/handler"
)

func main() {
	http.HandleFunc("/jwt", handler.JWTHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("> Failed to Start The Server. \n Reason : %s", err)
		panic(err)
	}
}
