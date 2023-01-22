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

	dynamic := alice.New(app.SessionManager.LoadAndSave, noSurf, app.authenticate)

	// Unprotected ########################################################
	//home
	router.Handler(http.MethodGet, "/", dynamic.ThenFunc(app.home))
	//blog
	router.Handler(http.MethodGet, "/blog", dynamic.ThenFunc(app.blog))
	router.Handler(http.MethodGet, "/blog/view/:id", dynamic.ThenFunc(app.blogView))
	router.Handler(http.MethodGet, "/user/signup", dynamic.ThenFunc(app.userSignUp))
	router.Handler(http.MethodPost, "/user/signup", dynamic.ThenFunc(app.userSignUpPost))
	router.Handler(http.MethodGet, "/user/login", dynamic.ThenFunc(app.userLogin))
	router.Handler(http.MethodPost, "/user/login", dynamic.ThenFunc(app.userLoginPost))

	// Protected ########################################################
	protected := dynamic.Append(app.requireAuthentication)
	// user
	router.Handler(http.MethodPost, "/user/logout", protected.ThenFunc(app.userLogoutPost))
	// blog
	router.Handler(http.MethodGet, "/blog/create", protected.ThenFunc(app.blogCreate))
	router.Handler(http.MethodPost, "/blog/create", protected.ThenFunc(app.blogCreatePost))

	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	return standard.Then(router)
}
