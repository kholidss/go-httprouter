package main

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//go:embed resources
var resources embed.FS

func main() {
	router := httprouter.New()
	//dibawah digunakan untuk menghilangkan penulisan folfer resources
	//yang perlu diubah adalah parameter ke-2, yaitu nama foldernya
	directory, _ := fs.Sub(resources, "resources")

	//with anonymous functiom
	router.ServeFiles("/files/*filepath", http.FS(directory))

	server := http.Server{
		Handler: router,
		Addr:    "localhost:3000",
	}

	server.ListenAndServe()
}
