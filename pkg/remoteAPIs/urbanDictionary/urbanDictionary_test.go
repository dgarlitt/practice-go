package urbanDictionary

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var request = struct {
	path, query       string // request
	contenttype, body string // response
}{
	path:        "/search.json?",
	query:       "q=%23Kenya",
	contenttype: "application/json",
	body:        udResponse,
}

var (
	udResponse = `{"tags":["tag1","tag2"],"result_type":"exact","list":[{"defid":12345,"word":"fake1","author":"authorone","permalink":"http://wat.urbanup.com/3322419","definition":"Definition number one.","example":"example text one","thumbs_up":3000,"thumbs_down":2000,"current_vote":""},{"defid":67890,"word":"fake2","author":"authortwo","permalink":"http://wat.urbanup.com/1308686","definition":"Definition number two.","example":"example text two","thumbs_up":300,"thumbs_down":250,"current_vote":""}],"sounds":["http://media.urbandictionary.com/sound/fake-111.mp3","http://wav.urbandictionary.com/fake-222.wav"]}`
)

func TestLookupDefinition(t *testing.T) {
	assert := assert.New(t)
	expectedWord := "fake2"
	expectedText := "Definition number two."

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, udResponse)
	}))
	defer ts.Close()

	defURL = ts.URL
	params := &Params{Term: "fake", APIKey: "fakeAPIKey"}

	def, err := LookupDefinition(params)

	assert.Equal(nil, err)
	assert.Equal(expectedWord, def.Word)
	assert.Equal(expectedText, def.Text)

}
