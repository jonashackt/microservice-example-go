# microservice-example-go
Example project showing how to create a simple Microservice with Go


# Getting started

Inspired by

https://dzone.com/articles/go-microservices-blog-series-part-1
https://dzone.com/articles/go-microservices-part-2-building-our-first-service

## Setup a GO dev environment

### GO workspace

https://golang.org/doc/code.html

`mkdir goworkspace`

Now make this directory you 'central' go workspace on this machine:

```
export GOPATH=`pwd`
```

Also add this to your `~/.bash_profile`, e.g.:

`export GOPATH=/Users/jonashecht/dev/goworkspace`


### Install SDK

https://golang.org/doc/install

`brew install go` 

### IDE

IntelliJ with Go Plugin or Goland, both from Jetbrains (see https://www.jetbrains.com/go/)

### Create first project

```
$ cd goworkspace
mkdir microservice-example-go
cd microservice-example-go
mkdir accountservice
cd accountservice
touch main.go
```

Open the project in IntelliJ.

Now implement the `main.go`:

```
package main

import (

        "fmt"

        )

var appName = "accountservice"

func main() {

    fmt.Printf("Starting %v\n", appName)

}
```

### Run first go app

```
go run *.go
```

# HTTP endpoints with Go

see https://thenewstack.io/make-a-restful-json-api-go/

Create a new dir and file `service/webservice.go` inside `accountservice`:

```
package service

import (

	"net/http"

	"log"

)

func StartWebServer(port string) {

	log.Println("Starting HTTP service at " + port)

	err := http.ListenAndServe(":" + port, nil)    // Goroutine will block here

	if err != nil {

		log.Println("An error occured starting HTTP listener at port " + port)

		log.Println("Error: " + err.Error())

	}

}
```

# Links

https://www.golang-book.com/books/intro/4


