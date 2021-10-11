package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	//with anonymous functiom
	router.GET("/images/*image", func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		image := p.ByName("image")
		text := "Image: " + image

		fmt.Fprintln(rw, text)
	})

	server := http.Server{
		Handler: router,
		Addr:    "localhost:3000",
	}

	server.ListenAndServe()
}
