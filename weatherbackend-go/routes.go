package main

import (
	"github.com/gorilla/mux"
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
			contentTypeTextPlain(writer)

			name := pathVariable("name", request)

			log.Println("Request for /{name} with GET, name was: " + name)

			writer.Write([]byte("Hello " + name + "! This is a RESTful HttpService written in Go. Try to use some other HTTP verbs (don´t say 'methods' :P )\n"))
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

/*@RequestMapping(value = "/{name}", method = RequestMethod.GET, produces = "text/plain")
public String whatsTheSenseInThat(@PathVariable("name") String name) {
LOG.info("Request for /{name} with GET");
return "Hello " + name + "! This is a RESTful HttpService written in Spring. Try to use some other HTTP verbs (don´t say 'methods' :P ) :)";
}*/
