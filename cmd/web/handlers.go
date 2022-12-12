package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/LachlanStephan/ls_server/internal/models"
)

func (app *application) blogView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
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
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	_, err := app.blogs.Insert(1, "some title", "this is an amazing blog post")
	if err != nil {
		app.serverError(w, err)
		return
	}
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

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
