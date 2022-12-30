package main

import (
	"errors"
	"fmt"
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
	data := app.newTemplateData(r)
	app.render(w, http.StatusOK, "blog-create.tmpl.html", data)
}

func (app *application) blogCreatePost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	user_id, err := strconv.Atoi(r.PostForm.Get("user_id"))
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	title := r.PostForm.Get("title")
	content := r.PostForm.Get("content")

	id, err := app.blogs.Insert(user_id, title, content)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/blog/view/%d", id), http.StatusSeeOther)
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	blogs, err := app.blogs.Recent()
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := app.newTemplateData(r)
	data.BlogLinks = blogs

	app.render(w, http.StatusOK, "home.tmpl.html", data)
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

// func (app *application) admin(w http.ResponseWriter, r *http.Request) {
// if r.URL.Path != "/admin" {
// app.notFound(w)
// return
// }

// app.render(w, http.StatusOK, "admin.tmpl.html", nil)
// /*
// this will be an admin view
// just get the right templates/html and show it
// */
// }

func (app *application) adminLogin(password string) bool {
	return false
	/*
		hash password and see if matches the db for admin user
	*/
}
