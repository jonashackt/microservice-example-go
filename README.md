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
mkdir src/github.com/YourUserNameHere/microservice-example-go
cd microservice-example-go
mkdir accountservice
cd accountservice
touch main.go
```

Open the project in IntelliJ.

Now implement the `main.go`:

```go
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


### Bootstrapping a HTTP server

Create a new dir and file `service/webservice.go` inside `accountservice`:

```go
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

Update [main.go](accountservice/main.go) to startup HTTP server:

```go
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
```

Now run the app again:

```
$ go run *.go
Stooaaarrrrting accountservice
2019/01/29 21:13:19 Starting HTTP service at 6767
```

Switch over to a new tab/bash and curl our new server:

```
$ curl localhost:6767
404 page not found
```

As we didn't declare any routes, this 404 should be great for now.


### Adding Routes

Let's use [gorillatoolkit](http://www.gorillatoolkit.org/) as the web toolkit for Go.

Therefore create `routes.go` file inside `service` directory:

```go
package service

import "net/http"

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

"GetAccount",                                     // Name

"GET",                                            // HTTP method

"/accounts/{accountId}",                          // Route pattern

func(w http.ResponseWriter, r *http.Request) {

            w.Header().Set("Content-Type", "application/json; charset=UTF-8")

            w.Write([]byte("{\"result\":\"OK\"}"))

        },

},

}
```

We're using [Go structs](https://gobyexample.com/structs) here. Structs can be used to define data structure/models in Go. Here the struct defines everything needed for our first route.

And __now - OMG! - there we are!__ We need to create some __boilderplate__ code to actually spin up the Gorilla Router (oh, I thought here's everything really better than in Spring Boot, hu?!!).

So let's create `router.go` also inside `service` directory:

```go
package service

import (

"github.com/gorilla/mux"

)

// Function that returns a pointer to a mux.Router we can use as a handler.

func NewRouter() *mux.Router {

    // Create an instance of the Gorilla router

router := mux.NewRouter().StrictSlash(true)

// Iterate over the routes we declared in routes.go and attach them to the router instance

for _, route := range routes {

    // Attach each route, uses a Builder-like pattern to set each route up.

router.Methods(route.Method).

                Path(route.Pattern).

                Name(route.Name).

                Handler(route.HandlerFunc)

}

return router

}
```

That's interesting. We can access `routes` directly from [routes.go](accountservice/service/routes.go) without importing the file - it's in the same package.

Furthermore we can use [range](https://gobyexample.com/range) and for (the `_` is just to ignore the index that range provides also) to iterate over our `routes` and set the route up with Gorilla's builder API.

Now head over to [webserver.go](accountservice/service/webservice.go) and add the two lines at the start of the `StartWebServer` function:

```go
func StartWebServer(port string) {

        r := NewRouter()             
        http.Handle("/", r) 
```

Start the server again with `go run *.go` and curl 

```
$ curl localhost:6767/accounts/10000 
{"result":"OK"}
```

# Links

https://www.golang-book.com/books/intro/4

https://gobyexample.com/range


