package main

import (
	"fmt"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request)  {
	fmt.Print("hello there")
}

func (app *application) blogGet(w http.ResponseWriter, r *http.Request) {
	fmt.Print("get blogs here")
}

func (app *application) blogPost(w http.ResponseWriter, r *http.Request) {
	fmt.Print("post blogs here")
}