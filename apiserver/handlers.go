package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/dgarlitt/practice-go/pkg/remoteAPIs/urbanDictionary"
	"github.com/dgarlitt/practice-go/pkg/shapes"
	"github.com/gorilla/mux"
	"github.com/hoisie/mustache"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	outText := "Give me a <a href=\"/v1/rectangle\">rectangle.</a>"
	outText += "<br />"
	outText += "Give me a <a href=\"/v1/circle\">circle.</a>"
	w.Write([]byte(outText))
}

func fancyPantsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	filename := path.Join(path.Join(os.Getenv("PWD"), "apiserver/templates"), "index.mustache")
	data := mustache.RenderFile(filename, map[string]string{"title": "Index", "body": "laser"})
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(data))
}

func mustacheHandler(w http.ResponseWriter, r *http.Request) {
	filename := path.Join(path.Join(os.Getenv("PWD"), config.TemplatePath), "rolliefingers.mustache")
	data := mustache.RenderFile(filename, map[string]string{"title": "Rollie Fingers", "body": "Check out my 'stache!!!"})
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(data))
}

func urbanDictionaryHandler(w http.ResponseWriter, r *http.Request) {
	var jsonstr []byte

	pathVars := mux.Vars(r)
	params := &urbanDictionary.Params{
		Term:   strings.Replace(pathVars["term"], "-", "%20", -1),
		APIKey: r.URL.Query().Get("api-key"),
	}

	definition, err := urbanDictionary.LookupDefinition(params)

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

func fakeCoverallsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("############### Headers #################")

	for k, v := range r.Header {
		fmt.Println(k)
		fmt.Println(v)
		fmt.Println("-------------------------------------------")
	}

	fmt.Println("################ Body ###################")
	fmt.Println(r.Body)

	fmt.Println("################ Form ###################")
	fmt.Println(r.FormValue("json_file"))
	for k, v := range r.Form {
		fmt.Println(k)
		fmt.Println(v)
		fmt.Println("-------------------------------------------")
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{}"))
}
