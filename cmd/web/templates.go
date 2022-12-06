package main

import (
	"html/template"
	"path/filepath"

	"github.com/LachlanStephan/ls_server/internal/models"
)

// list of templates used across the site
// can compose multiple sets of data here

type templateData struct {
	Blog      *models.Blog
	BlogLinks []*models.BlogLink
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

		ts, err := template.ParseFiles("./ui/html/base.tmpl.html")
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
