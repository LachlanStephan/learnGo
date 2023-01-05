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

	dynamic := alice.New(app.SessionManager.LoadAndSave)

	router.Handler(http.MethodGet, "/", dynamic.ThenFunc(app.home))
	router.Handler(http.MethodGet, "/blog", dynamic.ThenFunc(app.blog))
	router.Handler(http.MethodGet, "/blog/view/:id", dynamic.ThenFunc(app.blogView))
	router.Handler(http.MethodGet, "/blog/create", dynamic.ThenFunc(app.blogCreate))
	router.Handler(http.MethodPost, "/blog/create", dynamic.ThenFunc(app.blogCreatePost))
	// router.HandlerFunc(http.MethodGet, "/admin", app.admin)

	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	return standard.Then(router)
}
