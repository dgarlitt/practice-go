package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleRectangle() {
	rectangle := Rectangle{3, 4, "myRectangle"}
	rectangle.display()
	// Output: Length = 3, Width=4, name = myRectangle
}

func TestRectangleHandlerResponse(t *testing.T) {
	assert := assert.New(t)
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()

	expectedResponseCode := 200
	expectedResponseHead := "application/json"
	expectedResponseBody := "{\"Length\":3,\"Width\":4,\"Name\":\"generatedRectangle\"}\n"

	RectangleHandler(response, request)

	assert.Equal(expectedResponseCode, response.Code)
	assert.Equal(expectedResponseHead, response.HeaderMap.Get("Content-Type"))
	assert.Equal(expectedResponseBody, response.Body.String())
}
