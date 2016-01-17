package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexHandlerResponse(t *testing.T) {
	assert := assert.New(t)
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()

	expectedResponseCode := 200
	expectedResponseHead := "text/html"
	expectedResponseBody := "Give me a <a href=\"/v1/rectangle\">rectangle.</a><br />Give me a <a href=\"/v1/circle\">circle.</a>"

	IndexHandler(response, request)

	assert.Equal(expectedResponseCode, response.Code)
	assert.Equal(expectedResponseHead, response.HeaderMap.Get("Content-Type"))
	assert.Equal(expectedResponseBody, response.Body.String())
}

func TestRectangleHandlerResponse(t *testing.T) {
	assert := assert.New(t)
	request, _ := http.NewRequest("GET", "/rectangle", nil)
	response := httptest.NewRecorder()

	expectedResponseCode := 200
	expectedResponseHead := "application/json"
	expectedResponseBody := "{\"Length\":3,\"Width\":4,\"Name\":\"generatedRectangle\"}\n"

	RectangleHandler(response, request)

	assert.Equal(expectedResponseCode, response.Code)
	assert.Equal(expectedResponseHead, response.HeaderMap.Get("Content-Type"))
	assert.Equal(expectedResponseBody, response.Body.String())
}

func TestSquareHandlerResponse(t *testing.T) {
	assert := assert.New(t)
	request, _ := http.NewRequest("GET", "/rectangle", nil)
	response := httptest.NewRecorder()

	expectedResponseCode := 200
	expectedResponseHead := "application/json"
	expectedResponseBody := "{\"Width\":6,\"Name\":\"generatedSquare\"}\n"

	SquareHandler(response, request)

	assert.Equal(expectedResponseCode, response.Code)
	assert.Equal(expectedResponseHead, response.HeaderMap.Get("Content-Type"))
	assert.Equal(expectedResponseBody, response.Body.String())
}

func TestCircleHandlerResponse(t *testing.T) {
	assert := assert.New(t)
	request, _ := http.NewRequest("GET", "/circle", nil)
	response := httptest.NewRecorder()

	expectedResponseCode := 200
	expectedResponseHead := "application/json"
	expectedResponseBody := "{\"Radius\":2,\"Name\":\"generatedCircle\"}\n"

	CircleHandler(response, request)

	assert.Equal(expectedResponseCode, response.Code)
	assert.Equal(expectedResponseHead, response.HeaderMap.Get("Content-Type"))
	assert.Equal(expectedResponseBody, response.Body.String())
}
