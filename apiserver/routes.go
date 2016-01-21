package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func setupRoutes() {
	router := mux.NewRouter()
	router.HandleFunc("/", indexHandler)
	router.HandleFunc("/rolliefingers", mustacheHandler)
	router.HandleFunc("/definition/{term}", urbanDictionaryHandler)
	router.HandleFunc("/v1/rectangle", rectangleHandler)
	router.HandleFunc("/v1/square", squareHandler)
	router.HandleFunc("/v1/circle", circleHandler)

	http.Handle("/", router)
}
