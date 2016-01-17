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
	outText := "Give me a <a href=\"/v1/rectangle\">rectangle.</a>"
	outText += "<br />"
	outText += "Give me a <a href=\"/v1/circle\">circle.</a>"
	w.Write([]byte(outText))
}

// RectangleHandler - handles rectangle requests
func RectangleHandler(w http.ResponseWriter, r *http.Request) {
	rec := shapes.Rectangle{Length: 3, Width: 4, Name: "generatedRectangle"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rec)
}

//SquareHandler - handles square requests
func SquareHandler(w http.ResponseWriter, r *http.Request) {
	rec := shapes.Square{Width: 6, Name: "generatedSquare"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rec)
}

//CircleHandler - handles circle requests
func CircleHandler(w http.ResponseWriter, r *http.Request) {
	rec := shapes.Circle{Radius: 2, Name: "generatedCircle"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rec)
}
