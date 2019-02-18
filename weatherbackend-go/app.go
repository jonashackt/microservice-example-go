package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var appName = "weatherbackend"

type App struct {
	Router *mux.Router
}

func (app *App) Run() {
	app.startWebServer("6768")
}

func (app *App) startWebServer(port string) {

	http.Handle("/", app.Router)

	log.Println("Starting HTTP service at " + port)

	err := http.ListenAndServe(":"+port, nil) // Goroutine will block here

	if err != nil {
		log.Println("An error occured starting HTTP listener at port " + port)
		log.Println("Error: " + err.Error())
	}
}

func (app *App) Initialize() {
	app.Router = NewRouter()
}
