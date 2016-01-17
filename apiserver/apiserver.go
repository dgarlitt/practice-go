package main

import (
	"log"
	"net/http"
)

func main() {
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8008", nil))
}
