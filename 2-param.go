package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	//with anonymous functiom
	router.GET("/products/:id", func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		id := p.ByName("id")
		text := "Product " + id
		fmt.Fprintln(rw, text)

		intId, _ := strconv.Atoi(id)
		if intId == 1 {
			fmt.Fprintln(rw, "Yes agree")
		}
	})

	server := http.Server{
		Handler: router,
		Addr:    "localhost:3000",
	}

	server.ListenAndServe()
}
