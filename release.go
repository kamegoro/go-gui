//go:build release
// +build release

package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:generate sh -c "cd frontend; npm run build"
//go:embed frontend/.svelte-kit/build*
var content embed.FS

func init() {
	pub, err := fs.Sub(content, "frontend/.svelte-kit/build")
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/", http.FileServer(http.FS(pub)))
}