package main

import (
	"errors"
	"html/template"
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

	// blog.Content -> deal with this being markdown and render it to html instead
	files := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
		"./ui/html/pages/view.tmpl.html",
		"./ui/html/partials/footer.tmpl.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// use a template to allow for multiple data sets to be executed - ts.ExecuteTemplate will only allow one dataset to be passed.
	data := &blogTemplate{
		Blog: blog,
	}

	err = ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		app.serverError(w, err)
	}
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

	app.clientResponse(w, 201)
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

	files := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
		"./ui/html/pages/home.tmpl.html",
		"./ui/html/partials/footer.tmpl.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := &blogRecentTemplate{
		Recent: blogs,
	}
	
	err = ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		app.serverError(w, err)
	}
}

// not in use
// func (app *application) userCreate(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
// 		w.Header().Set("Allow", http.MethodPost)
// 		app.clientError(w, http.StatusMethodNotAllowed)
// 		return
// 	}

// 	// we don't need id returned so assign _ instead && return 201 if no err
// 	_, err := app.users.Insert("hello", "there", true)
// 	if err != nil {
// 		app.serverError(w, err)
// 		return
// 	}
// 	app.clientResponse(w, 201)
// }

func (app *application) blog(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/blog" {
		app.notFound(w)
		return
	}

	files := []string {
		"./ui/html/base.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
		"./ui/html/partials/footer.tmpl.html",
		"./ui/html/pages/blog.tmpl.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, err)
		return
	}
}