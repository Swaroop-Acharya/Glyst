package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

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
	w.Write([]byte("Display a form for creating a specific Glyst"))
}

func (app *application) glystCreatePost(w http.ResponseWriter, r *http.Request) {
	title := "My Books"
	content := "I like to read books"
	expires := 10

	id, err := app.glysts.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/glyst/view/%d", id), http.StatusSeeOther)
}
