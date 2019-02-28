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

Very good basic concepts site:
https://gobyexample.com

Curated list of tools https://awesome-go.com/

Configuration / parameterization https://github.com/spf13/viper

HTTP & Dependency Injection https://github.com/mustafaakin/gongular

### Enums

https://blog.learngoprogramming.com/golang-const-type-enums-iota-bc4befd096d3

https://stackoverflow.com/questions/14426366/what-is-an-idiomatic-way-of-representing-enums-in-go

### Testing

https://jaxenter.de/testen-benchmarks-go-70936

execute

```
go test -v
```

If that doesn't work, it gives you something like this

```
weatherbackend-go jonashecht$ go test -v
=== RUN   TestGetSenseInThat
--- FAIL: TestGetSenseInThat (0.00s)
panic: runtime error: invalid memory address or nil pointer dereference [recovered]
	panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x1244687]

goroutine 19 [running]:
testing.tRunner.func1(0xc000116100)
	/usr/local/Cellar/go/1.11.5/libexec/src/testing/testing.go:792 +0x387
panic(0x128b6e0, 0x14e3340)
	/usr/local/Cellar/go/1.11.5/libexec/src/runtime/panic.go:513 +0x1b9
github.com/gorilla/mux.(*Router).ServeHTTP(0x0, 0x1324900, 0xc00009c800, 0xc000116200)
	/Users/jonashecht/dev/goworkspace/src/github.com/gorilla/mux/mux.go:176 +0x37
github.com/jonashackt/microservice-example-go/weatherbackend-go.executeRequest(0xc000116200, 0x3)
	/Users/jonashecht/dev/goworkspace/src/github.com/jonashackt/microservice-example-go/weatherbackend-go/main_test.go:33 +0xb0
github.com/jonashackt/microservice-example-go/weatherbackend-go.TestGetSenseInThat(0xc000116100)
	/Users/jonashecht/dev/goworkspace/src/github.com/jonashackt/microservice-example-go/weatherbackend-go/main_test.go:22 +0x65
testing.tRunner(0xc000116100, 0x12eef18)
	/usr/local/Cellar/go/1.11.5/libexec/src/testing/testing.go:827 +0xbf
created by testing.(*T).Run
	/usr/local/Cellar/go/1.11.5/libexec/src/testing/testing.go:878 +0x35c
exit status 2
FAIL	github.com/jonashackt/microservice-example-go/weatherbackend-go	0.018s
```

Now we should use Debugging to get more information

### Debugging Go with IntelliJ on Mac

You maybe get an error, that tells you `could not launch process: exec: "lldb-server": executable file not found in $PATH` or the like.

All you need to do is to install X-Code command line tools with 

```
$ xcode-select --install
```

For me, this didn't work because of "networking problems". Luckily [you can download them directly](https://stackoverflow.com/a/20243261/4964553) from Apple Developer Tools site: https://developer.apple.com/downloads/index.action (you only need your Apple ID user & pw in place).



### Vendoring

https://goenning.net/2017/02/23/packages-vendoring-in-go/

https://arslan.io/2018/08/26/using-go-modules-with-vendor-support-on-travis-ci/

### Webservices

https://jaxenter.de/restful-rest-api-go-golang-68845

https://kev.inburke.com/kevin/golang-json-http/

https://gist.github.com/reagent/043da4661d2984e9ecb1ccb5343bf438


#### Testing Webservices

httptest

Start the server again with `go run .` and curl 

Use Postman to access `localhost:6767/weather/general/outlook` via a POST with the following JSON body:

```

```



##### gorilla-mux

https://gowebexamples.com/routes-using-gorilla-mux/

https://semaphoreci.com/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql

https://medium.com/@kelvin_sp/building-and-testing-a-rest-api-in-golang-using-gorilla-mux-and-mysql-1f0518818ff6

##### Swagger Go

https://www.ribice.ba/swagger-golang/


### Microservice frameworks

https://medium.com/seek-blog/microservices-in-go-2fc1570f6800

[3 main choices](https://www.quora.com/What-are-some-of-the-most-well-known-Go-Golang-Microservice-frameworks-libraries-out-there):

* https://github.com/micro/go-micro
* https://github.com/nytimes/gizmo
* https://github.com/go-kit/kit (complex! https://www.reddit.com/r/golang/comments/5tesl9/why_i_recommend_to_avoid_using_the_gokit_library/)

or do it without a Microservice framework

https://outcrawl.com/go-microservices-cqrs-docker/




# Microservice(s) with go-micro

https://github.com/micro/micro

https://micro.mu/docs/toolkit.html

### Install micro CLI

__Don't do__ `brew install micro` - this will install an text editor.

Instead, install go-micro incl. CLI via

```
go get -u github.com/micro/micro
```

> This isn't everything! You'll get $ micro --version  -bash: micro: command not found errors, so please also do:

And then make sure to add the bin directory to your path inside your `.bash_profile`:

```
export GOPATH=/Users/jonashecht/dev/goworkspace
export PATH=$PATH:$GOPATH/bin
```

```
brew install protobuf
```

### Generate project skeleton

Inside your `GOPATH` (for me this is `/Users/jonashecht/dev/goworkspace`), execute:

```
micro new github.com/jonashackt/microservice-example-go/weatherservice
```

This will give something like:

```
$ micro new github.com/jonashackt/microservice-example-go/weatherservice
Creating service go.micro.srv.weatherservice in /Users/jonashecht/dev/goworkspace/src/github.com/jonashackt/microservice-example-go/weatherservice

.
├── main.go
├── plugin.go
├── handler
│   └── example.go
├── subscriber
│   └── example.go
├── proto/example
│   └── example.proto
├── Dockerfile
├── Makefile
└── README.md


download protobuf for micro:

brew install protobuf
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
go get -u github.com/micro/protoc-gen-micro

compile the proto file example.proto:

cd /Users/jonashecht/dev/goworkspace/src/github.com/jonashackt/microservice-example-go/weatherservice
protoc --proto_path=. --go_out=. --micro_out=. proto/example/example.proto
```