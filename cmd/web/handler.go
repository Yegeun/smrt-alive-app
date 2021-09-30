package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// Define a home handler function which writes a byte slice containing
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// check if the current request URL path exactly matches
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}
	//Initialize a slice containing that paths of two files
	files := []string{
	    "./ui/html/home.page.tmpl",
	    "./ui/html/base.layout.tmpl",
	    "./ui/html/footer.partial.tmpl",
	}

	// relative to your current working directory,
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err) // Use the serverError() helper.
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, nil) // ts template set to write
	if err != nil {
		app.serverError(w, err) // Use the serverError() helper.
	}
}

func (app *application) showUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w) // Use the notFound() helper.
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func (app *application) createUser(w http.ResponseWriter, r *http.Request) {

	//r.method checks if its a post or not if its not it will write method not alloawed
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	// Create some variables holding dummy data. We'll remove these later on
	// during the build.
	title := "hello"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
	expires := "5"
	// ID of the new record back.
	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}
	// Redirect the user to the relevant page for the snippet.
	http.Redirect(w, r, fmt.Sprintf("/user?id=%d", id), http.StatusSeeOther)
}
