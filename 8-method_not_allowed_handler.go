package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Hello(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprint(rw, "this is from hello")
}

func main() {
	router := httprouter.New()

	//handlerfunction buat Not Found
	router.NotFound = http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "tidak ketemu nih")
	})

	//handlerfunction buat method not allowed
	router.MethodNotAllowed = http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "method salah bosss")
	})

	//with anonymous functiom
	router.POST("/", func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(rw, "Hello World")
	})

	//with call the function
	router.GET("/hello", Hello)

	server := http.Server{
		Handler: router,
		Addr:    "localhost:3000",
	}

	server.ListenAndServe()
}
