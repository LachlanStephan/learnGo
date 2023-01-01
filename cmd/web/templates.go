package main

import (
	"html/template"
	"path/filepath"
	"time"

	"github.com/LachlanStephan/ls_server/internal/models"
)

// list of templates used across the site
// can compose multiple sets of data here
type templateData struct {
	Blog        *models.Blog
	BlogLinks   []*models.BlogLink
	CurrentYear int
	Form        any
}

func formatCreatedAt(t time.Time) string {
	if t.IsZero() {
		return "Unknown"
	}

	return t.Format("02 Jan 2006 at 15:04") + " (UTC)"
}

// creating an object to store helper functions that the templates may use
// these must return only 1 value || 1 value && err
var functions = template.FuncMap{
	"formatCreatedAt": formatCreatedAt,
}

// create in memory cache to store templates and prevent repeating ourselves in the handler funcs
func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./ui/html/pages/*.tmpl.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles("./ui/html/base.tmpl.html")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob("./ui/html/partials/*.tmpl.html")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}
