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
func LookupDefinition(params *Params, c chan *Definition) (err error) {
	var (
		results    *DictionaryResults
		definition = &Definition{}
	)

	results, err = fetchDefinitions(params)

	if err == nil && results.ResultType != "exact" {
		err = errors.New("No results found for term: \"" + params.Term + "\".")
	}

	if err == nil {
		userDef := findBestUserDefinition(results)
		definition.Word = userDef.Word
		definition.Text = userDef.Text
	} else {
		definition.Error = err
	}

	c <- definition

	return
}

// *****************************************************
// ****************** Private Parts ********************
// *****************************************************

var defURL = "https://mashape-community-urban-dictionary.p.mashape.com/define"

func fetchDefinitions(params *Params) (results *DictionaryResults, err error) {
	var resp *http.Response
	var body []byte

	url := defURL + "?term=" + params.Term
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Set("X-Mashape-Key", params.APIKey)
	req.Header.Set("Accept", "text/plain")

	client := &http.Client{}
	resp, err = client.Do(req)
	defer resp.Body.Close()

	if err == nil {
		body, err = ioutil.ReadAll(resp.Body)
	}

	if err == nil {
		err = json.Unmarshal(body, &results)
	}

	return
}

func findBestUserDefinition(results *DictionaryResults) (bestDef *UserDefinition) {
	var ratio float64
	var total float64
	var bestRatio float64

	for _, definition := range results.UserDefinitions {
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
