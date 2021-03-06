package urbanDictionary

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

// LookupDefinition is the single point of entry to the package.
// It will select the highest ranked definition for a given term
// from the list of definitions returned by the remote api.
func LookupDefinition(params *Params) (definition *Definition, err error) {
	definition = &Definition{}
	c := make(chan *DictionaryResults, 1)
	go fetchDefinitions(params, c)
	results := <-c
	close(c)
	err = results.Error

	if err == nil && results.ResultType != "exact" {
		err = errors.New("No results found for term: \"" + params.Term + "\".")
	}

	if err == nil {
		userDef := findBestUserDefinition(results.UserDefinitions)
		definition = &Definition{Word: userDef.Word, Text: userDef.Text}
	}

	return
}

// *****************************************************
// ****************** Private Parts ********************
// *****************************************************

var defURL = "https://mashape-community-urban-dictionary.p.mashape.com/define"

func fetchDefinitions(params *Params, c chan *DictionaryResults) {
	var (
		resp    *http.Response
		body    []byte
		results = &DictionaryResults{}
		err     error
	)

	url := defURL + "?term=" + params.Term
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Set("X-Mashape-Key", params.APIKey)
	req.Header.Set("Accept", "text/plain")

	client := &http.Client{}
	resp, err = client.Do(req)

	if err == nil {
		defer resp.Body.Close()
		body, err = ioutil.ReadAll(resp.Body)
	}

	if err == nil {
		err = json.Unmarshal(body, results)
	}

	if err != nil {
		results.Error = err
	}

	c <- results

	return
}

func findBestUserDefinition(definitions []*UserDefinition) (bestDef *UserDefinition) {
	var ratio float64
	var total float64
	var bestRatio float64

	for _, definition := range definitions {
		total = float64(definition.ThumbsUp + definition.ThumbsDown)

		if total == float64(0) {
			ratio = 0
		} else {
			ratio = float64(definition.ThumbsDown) / total
		}

		if ratio > bestRatio {
			bestRatio = ratio
			bestDef = definition
		}
	}

	return
}
