package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/LachlanStephan/ls_server/internal/models"
	"github.com/julienschmidt/httprouter"
)

type blogCreateForm struct {
	Title      string
	Content    string
	User_id    int
	FormErrors map[string]string
}

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
	// TODO:
	// check that the user is logged in here
	// if not
	// redirect them to login page
	// this can be done once auth is implemented

	data := app.newTemplateData(r)
	data.Form = blogCreateForm{
		User_id: 1,
		Title:   "TITLE TEST",
		Content: "TEST",
	}

	app.render(w, http.StatusOK, "blog-create.tmpl.html", data)
}

func (app *application) blogCreatePost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form := blogCreateForm{
		FormErrors: map[string]string{},
	}

	// move this validation to
	// validation.go file later
	user_id, err := strconv.Atoi(r.PostForm.Get("user_id"))
	if err != nil {
		form.FormErrors["user_id"] = "Invalid user_id"
	} else {
		form.User_id = user_id
	}

	title := r.PostForm.Get("title")
	content := r.PostForm.Get("content")

	// move this validation to
	// validation.go file later
	if strings.TrimSpace(title) == "" || utf8.RuneCountInString(title) > 50 {
		form.FormErrors["title"] = "Invalid title"
	} else {
		form.Title = title
	}

	if strings.TrimSpace(content) == "" {
		form.FormErrors["content"] = "Invalid content"
	} else {
		form.Content = content
	}

	if len(form.FormErrors) > 0 {
		data := app.newTemplateData(r)
		data.Form = form
		// problem HERE
		fmt.Fprint(w)
		app.render(w, http.StatusUnprocessableEntity, "blog-create.tmpl.html", data)
		return
	}

	id, err := app.blogs.Insert(form.User_id, form.Title, form.Content)
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
