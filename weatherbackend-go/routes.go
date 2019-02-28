package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jonashackt/microservice-example-go/weatherbackend-go/domain"
	"log"
	"net/http"
)

// Defines a single route, e.g. a human readable name, HTTP method and the
// pattern the function that will execute when the route is called.
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Defines the type Routes which is just an array (slice) of Route structs.
type Routes []Route

// Initialize our routes
var routes = Routes{

	Route{
		"GetAccount",            // Name
		"GET",                   // HTTP method
		"/accounts/{accountId}", // Route pattern
		func(writer http.ResponseWriter, request *http.Request) {
			contentTypeApplicationJson(writer)

			log.Println("Request received with Method: " + request.Method)

			writer.Write([]byte("{\"result\":\"OK\"}"))
		},
	},
	Route{
		"GetTheSenseInThat", // Name
		"GET",               // HTTP method
		"/{name}",           // Route pattern
		func(writer http.ResponseWriter, request *http.Request) {

			name := pathVariable("name", request)

			log.Println("Request for /{name} with GET, name was: " + name)

			contentTypeTextPlain(writer)
			writer.Write([]byte("Hello " + name + "! This is a RESTful HttpService written in Go. Try to use some other HTTP verbs (don´t say 'methods' :P )\n"))
		},
	},
	Route{
		"GeneralOutlook",
		"POST",
		"/weather/general/outlook",
		func(writer http.ResponseWriter, request *http.Request) {

			var weather domain.Weather
			if request.Body == nil {
				http.Error(writer, "Please send a request body", 400)
				return
			}

			err := json.NewDecoder(request.Body).Decode(weather)
			if err != nil {
				http.Error(writer, err.Error(), 400)
				return
			}

			log.Println("Request for /weather/general/outlook with POST and weather with postalCode " + weather.PostalCode + " & flagColor " + weather.FlagColor)

			contentTypeApplicationJson(writer)
			writer.WriteHeader(http.StatusCreated)
			writer.Write([]byte("{\"result\":\"CREATED\"}"))
		},
	},
}

func pathVariable(pathVariableName string, request *http.Request) string {
	vars := mux.Vars(request)
	return vars[""+pathVariableName+""]
}

func contentTypeTextPlain(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "text/plain")
	return
}

func contentTypeApplicationJson(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	return
}
