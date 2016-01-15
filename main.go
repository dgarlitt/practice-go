package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", IndexHandler)
	router.HandleFunc("/rectangle", RectangleHandler)
	http.Handle("/", router)

	log.Fatal(http.ListenAndServe(":8008", nil))
}

// IndexHandler - handles default route
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("Give me a <a href=\"/rectangle\">rectangle.</a>"))
}

// RectangleHandler - handles rectangle requests
func RectangleHandler(w http.ResponseWriter, r *http.Request) {
	rec := Rectangle{3, 4, "generatedRectangle"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rec)
}
