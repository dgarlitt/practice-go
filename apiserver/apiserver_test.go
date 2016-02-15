package main

import (
	"bytes"
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

	indexHandler(response, request)

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

	rectangleHandler(response, request)

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

	squareHandler(response, request)

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

	circleHandler(response, request)

	assert.Equal(expectedResponseCode, response.Code)
	assert.Equal(expectedResponseHead, response.HeaderMap.Get("Content-Type"))
	assert.Equal(expectedResponseBody, response.Body.String())
}

func TestFileReaderHandlerResponse(t *testing.T) {
	assert := assert.New(t)
	request, _ := http.NewRequest("GET", "/fileread", nil)
	response := httptest.NewRecorder()

	expectedResponseCode := 200
	expectedResponseHead := "text/plain"
	expectedResponseBody := "TEXT Generated\n\nthis\nis\nmy\nfile\nthat\nI\ncreated\ntoday\n"

	fileReaderHandler(response, request)

	assert.Equal(expectedResponseCode, response.Code)
	assert.Equal(expectedResponseHead, response.HeaderMap.Get("Content-Type"))
	assert.Equal(expectedResponseBody, response.Body.String())
}

func TestFakeCoverallsHandler(t *testing.T) {
	out = &bytes.Buffer{}
	assert := assert.New(t)
	request, _ := http.NewRequest("POST", "/api/v1/jobs", nil)
	response := httptest.NewRecorder()

	request.Header.Set("Content-Type", "application/json")

	expectedResponseCode := 200
	expectedResponseHead := "application/json"
	expectedResponseBody := "{}"

	fakeCoverallsHandler(response, request)

	assert.Equal(expectedResponseCode, response.Code)
	assert.Equal(expectedResponseHead, response.HeaderMap.Get("Content-Type"))
	assert.Equal(expectedResponseBody, response.Body.String())
}
