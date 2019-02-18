package main

import (
	"fmt"
)

func main() {

	app := App{}
	app.Initialize()

	fmt.Printf("Stooaaarrrrting %v\n", appName)

	app.Run()
}
