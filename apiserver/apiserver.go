package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dgarlitt/practice-go/pkg/shapes"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", IndexHandler)
	router.HandleFunc("/v1/rectangle", RectangleHandler)
	router.HandleFunc("/v1/square", SquareHandler)
	router.HandleFunc("/v1/circle", CircleHandler)

	http.Handle("/", router)

	log.Fatal(http.ListenAndServe(":8008", nil))
}

// IndexHandler - handles default route
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("Give me a <a href=\"/v1/rectangle\">rectangle.</a>"))
}

// RectangleHandler - handles rectangle requests
func RectangleHandler(w http.ResponseWriter, r *http.Request) {
	rec := shapes.Rectangle{Length: 3, Width: 4, Name: "generatedRectangle"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rec)
}

//SquareHandler - handles square requests
func SquareHandler(w http.ResponseWriter, r *http.Request) {
	rec := pkg.Square{Length: 4, Width: 6, Name: "generatedSquare"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rec)
}

//CircleHandler - handles square requests
func CircleHandler(w http.ResponseWriter, r *http.Request) {
	rec := pkg.Square{Length: 3, Width: 2, Name: "generatedCircle"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rec)
}
