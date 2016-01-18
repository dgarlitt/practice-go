package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/dgarlitt/practice-go/pkg/remoteAPIs/urbanDictionary"
	"github.com/dgarlitt/practice-go/pkg/shapes"
	"github.com/gorilla/mux"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	outText := "Give me a <a href=\"/v1/rectangle\">rectangle.</a>"
	outText += "<br />"
	outText += "Give me a <a href=\"/v1/circle\">circle.</a>"
	w.Write([]byte(outText))
}

func urbanDictionaryHandler(w http.ResponseWriter, r *http.Request) {
	var jsonstr []byte

	pathVars := mux.Vars(r)
	params := &urbanDictionary.Params{
		Term:   strings.Replace(pathVars["term"], "-", "%20", -1),
		APIKey: r.URL.Query().Get("api-key"),
	}
	c := make(chan *urbanDictionary.Definition)

	go urbanDictionary.LookupDefinition(params, c)

	definition := <-c
	err := definition.Error

	if err == nil {
		jsonstr, err = json.Marshal(definition)
	}

	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(jsonstr)
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

func fileReaderHandler(w http.ResponseWriter, r *http.Request) {
	file, _ := ioutil.ReadFile("file.txt")
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("TEXT Generated\n\n"))
	w.Write([]byte(file))
}
