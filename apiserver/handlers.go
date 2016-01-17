package main

import (
	"encoding/json"
	"net/http"

	"github.com/dgarlitt/practice-go/pkg/shapes"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	outText := "Give me a <a href=\"/v1/rectangle\">rectangle.</a>"
	outText += "<br />"
	outText += "Give me a <a href=\"/v1/circle\">circle.</a>"
	w.Write([]byte(outText))
}

func rectangleHandler(w http.ResponseWriter, r *http.Request) {
	rec := shapes.Rectangle{Length: 3, Width: 4, Name: "generatedRectangle"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rec)
}

func squareHandler(w http.ResponseWriter, r *http.Request) {
	rec := shapes.Square{Width: 6, Name: "generatedSquare"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rec)
}

func circleHandler(w http.ResponseWriter, r *http.Request) {
	rec := shapes.Circle{Radius: 2, Name: "generatedCircle"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rec)
}
