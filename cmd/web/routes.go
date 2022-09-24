package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/blog/get", app.blogGet)
	mux.HandleFunc("/blog/post", app.blogPost)

	return mux
}
