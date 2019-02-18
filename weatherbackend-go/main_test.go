package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var app App

func TestMain(m *testing.M) {
	app = App{}
	app.Initialize()

	code := m.Run()

	os.Exit(code)
}

func TestGetSenseInThat(t *testing.T) {
	// Given
	name := "Jonas"
	request, _ := http.NewRequest("GET", "/"+name, nil)

	// When
	response := executeRequest(request)

	// Then
	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "Hello "+name+"! This is a RESTful HttpService written in Go. Try to use some other HTTP verbs (don´t say 'methods' :P )\n" {
		t.Errorf("Expected a specific string. Got %s", body)
	}
}

func executeRequest(request *http.Request) *httptest.ResponseRecorder {
	responseRecorder := httptest.NewRecorder()
	app.Router.ServeHTTP(responseRecorder, request)
	return responseRecorder
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
