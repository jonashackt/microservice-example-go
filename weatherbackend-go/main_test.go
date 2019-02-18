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
	request, _ := http.NewRequest("GET", "/Jonas", nil)
	response := executeRequest(request)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
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
