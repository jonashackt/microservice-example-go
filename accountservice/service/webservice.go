package service

import (
	"net/http"

	"log"
)

func StartWebServer(port string) {

	log.Println("Starting HTTP service at " + port)

	err := http.ListenAndServe(":"+port, nil) // Goroutine will block here

	if err != nil {

		log.Println("An error occured starting HTTP listener at port " + port)

		log.Println("Error: " + err.Error())

	}

}
