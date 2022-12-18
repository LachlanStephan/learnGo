package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.notFound(w)
	})

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer))

	router.HandlerFunc(http.MethodGet, "/", app.home)
	router.HandlerFunc(http.MethodGet, "/blog", app.blog)
	router.HandlerFunc(http.MethodGet, "/blog/view/:id", app.blogView)
	router.HandlerFunc(http.MethodPost, "/blog/create", app.blogCreate)
	router.HandlerFunc(http.MethodGet, "/admin", app.admin)

	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	return standard.Then(router)
}
