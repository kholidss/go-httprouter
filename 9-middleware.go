package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//middleware
type LogMiddleware struct {
	http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("Receive request")
	middleware.Handler.ServeHTTP(rw, r)
}

func Hello(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprint(rw, "this is from hello")
}

func main() {
	router := httprouter.New()

	//with anonymous functiom
	router.GET("/", func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(rw, "middleware")
	})

	middleware := LogMiddleware{router}

	//with call the function
	router.GET("/hello", Hello)

	server := http.Server{
		Handler: &middleware,
		Addr:    "localhost:3000",
	}

	server.ListenAndServe()
}
