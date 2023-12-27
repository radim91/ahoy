package main

import (
	"embed"
	"io/fs"
	"net/http"
	"radim91/api"
)

//go:embed ui/build
var content embed.FS

func clientHandler() http.Handler {
	fsys, err := fs.Sub(content, "ui/build")
    if err != nil {
        panic(err)
    }

	return http.FileServer(http.FS(fsys))
}

func main() {
	http.Handle("/", clientHandler())

	http.ListenAndServe(":3333", http.HandlerFunc(api.Serve))
}
