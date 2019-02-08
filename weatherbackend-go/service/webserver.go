package service

import (
	"log"
	"net/http"
)

func StartWebServer(port string) {

	router := NewRouter()

	http.Handle("/", router)

	log.Println("Starting HTTP service at " + port)

	err := http.ListenAndServe(":"+port, nil) // Goroutine will block here

	if err != nil {
		log.Println("An error occured starting HTTP listener at port " + port)
		log.Println("Error: " + err.Error())
	}
}
