package main

import (
	"fmt"

	"github.com/jonashackt/microservice-example-go/accountservice/service"
)

var appName = "accountservice"

func main() {

	fmt.Printf("Stooaaarrrrting %v\n", appName)
	service.StartWebServer("6767")
}
