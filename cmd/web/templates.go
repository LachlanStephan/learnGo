package main

import (
	"github.com/LachlanStephan/ls_server/internal/models"
)

// list of templates used across the site
// can compose multiple sets of data here

type blogTemplate struct{
	Blog *models.Blog
}