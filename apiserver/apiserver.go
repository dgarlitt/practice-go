package main

import (
	"log"
	"net/http"
)

type Config struct {
	TemplatePath string
}

var config = Config{
	TemplatePath: "apiserver/templates",
}

func main() {
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8008", nil))
}
