package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

var content embed.FS

func init() {
	pub, err := fs.Sub(content, "frontend/build")
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/", http.FileServer(http.FS(pub)))
}