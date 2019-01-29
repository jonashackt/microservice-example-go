package service

import "net/http"

// Defines a single route, e.g. a human readable name, HTTP method and the

// pattern the function that will execute when the route is called.

type Route struct {
	Name string

	Method string

	Pattern string

	HandlerFunc http.HandlerFunc
}

// Defines the type Routes which is just an array (slice) of Route structs.

type Routes []Route

// Initialize our routes

var routes = Routes{

	Route{

		"GetAccount", // Name

		"GET", // HTTP method

		"/accounts/{accountId}", // Route pattern

		func(writer http.ResponseWriter, request *http.Request) {

			writer.Header().Set("Content-Type", "application/json; charset=UTF-8")

			writer.Write([]byte("{\"result\":\"OK\"}"))

		},
	},
}
