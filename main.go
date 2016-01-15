package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", RectangleHandler)
	http.Handle("/", router)

	log.Fatal(http.ListenAndServe(":8008", nil))
}

// Rectangle shape definition
type Rectangle struct {
	Length, Width int
	Name          string
}

func (r Rectangle) display() {
	fmt.Printf("Length = %d, Width=%d, name = %s", r.Length, r.Width, r.Name)
	return
}

// RectangleHandler - handles rectangle requests
func RectangleHandler(w http.ResponseWriter, r *http.Request) {
	rec := Rectangle{3, 4, "generatedRectangle"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rec)
}
