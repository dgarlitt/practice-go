package main

import (
	"encoding/json"
	"fmt"
	"io"
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

var out io.Writer = os.Stdout

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	outText := "Give me a <a href=\"/v1/rectangle\">rectangle.</a>"
	outText += "<br />"
	outText += "Give me a <a href=\"/v1/circle\">circle.</a>"
	w.Write([]byte(outText))
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
	fmt.Fprint(out, "############### Headers #################")

	for k, v := range r.Header {
		fmt.Fprint(out, k)
		fmt.Fprint(out, v)
		fmt.Fprint(out, "-------------------------------------------")
	}

	fmt.Fprint(out, "################ Body ###################")
	fmt.Fprint(out, r.Body)

	fmt.Fprint(out, "################ Form ###################")
	fmt.Fprint(out, r.FormValue("json_file"))
	for k, v := range r.Form {
		fmt.Fprint(out, k)
		fmt.Fprint(out, v)
		fmt.Fprint(out, "-------------------------------------------")
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{}"))
}
