package main

import (
	"extensao-api/config"
	"extensao-api/router"
	"log"
	"net/http"
)

func main() {
	config.LoadEnv()
	r := router.New()
	const addr = ":8080"
	log.Printf("Starting server on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
