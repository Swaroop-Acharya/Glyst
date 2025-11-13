package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"

	"glyst/internal/models"
)

// Define a home handler function which writes a byte slice containing
// "Hello from Snippetbox" as the response body.
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	glysts, err := app.glysts.Latest()
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data := app.newTemplateData(r)
	data.Glysts = glysts
	app.render(w, r, http.StatusOK, "home.tmpl", data)
}

func (app *application) glystView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	// Use the SnippetModel's Get() method to retrieve the data for a
	// specific record based on its ID. If no matching record is found,
	// return a 404 Not Found response.
	glyst, err := app.glysts.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			http.NotFound(w, r)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	data := app.newTemplateData(r)
	data.Glyst = glyst

	app.render(w, r, http.StatusOK, "view.tmpl", data)
}

func (app *application) glystCreate(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	app.render(w, r, http.StatusOK, "create.tmpl", data)
}

func (app *application) glystCreatePost(w http.ResponseWriter, r *http.Request) {
	
	// First we call r.ParseForm() which adds any data in POST request bodies
	// to the r.PostForm map. This also works in the same way for PUT and PATCH
	// requests.

	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadGateway)
		return
	}

	// Use the r.PostForm.Get() method to retrieve the title and content
	// from the r.PostForm map.
	title := r.PostForm.Get("title")
	content := r.PostForm.Get("content")

	expires, err := strconv.Atoi(r.PostForm.Get("expires"))
	if err != nil {
		app.clientError(w, http.StatusBadGateway)
		return
	}

	fieldErrors:= make(map[string]string)


	if strings.TrimSpace(title) == "" {
		fieldErrors["title"]= "This field cannot be blank"
	}else if utf8.RuneCountInString(title) > 100{
		fieldErrors["title"]= "This field cannot be more than 100 characters long"
	}

	if strings.TrimSpace(content)==""{
		fieldErrors["content"]= "This field cannot be blank"
	}

	if expires!=1 && expires!=7 && expires!=365{
		fieldErrors["expires"]= "This field must be equal to 1, 7 or 365"
	}

	if len(fieldErrors) > 0 {
		fmt.Fprint(w,fieldErrors)	
		return
	}

	id, err := app.glysts.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/glyst/view/%d", id), http.StatusSeeOther)
}
