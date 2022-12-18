package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/LachlanStephan/ls_server/internal/models"
	"github.com/julienschmidt/httprouter"
)

func (app *application) blogView(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	blog, err := app.blogs.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	app.render(w, http.StatusOK, "view.tmpl.html", &templateData{
		Blog: blog,
	})
}

func (app *application) blogCreate(w http.ResponseWriter, r *http.Request) {
	/*
		Plan:
			this will display a form to create the blog

			on this form we will have a username and password field as part of the form
			all fields are required
			we will check if the username matches db username
			we will hash password and see if it matches db password

			in future we would also have userCreate route so others could publish blogs -> this would require the admin login and some kind of approval setting/method
	*/
}

func (app *application) blogCreatePost(w http.ResponseWriter, r *http.Request) {
	user_id := 1
	title := "some title"
	content := "this is a blog post"

	_, err := app.blogs.Insert(user_id, title, content)
	if err != nil {
		app.serverError(w, err)
		return
	}
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	blogs, err := app.blogs.Recent()
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, http.StatusOK, "home.tmpl.html", &templateData{
		BlogLinks: blogs,
	})
}

func (app *application) blog(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/blog" {
		app.notFound(w)
		return
	}

	blogs, err := app.blogs.GetAll()
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, http.StatusOK, "blog.tmpl.html", &templateData{
		BlogLinks: blogs,
	})
}

func (app *application) admin(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/admin" {
		app.notFound(w)
		return
	}

	app.render(w, http.StatusOK, "admin.tmpl.html", nil)
	/*
		this will be an admin view
		just get the right templates/html and show it
	*/
}

func (app *application) adminLogin(password string) bool {
	return false
	/*
		hash password and see if matches the db for admin user
	*/
}
