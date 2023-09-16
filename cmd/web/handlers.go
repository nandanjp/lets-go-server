package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
		"./ui/html/pages/home.tmpl.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}
	err = ts.ExecuteTemplate(w, "base", nil) //template body writes content to response body, nil is the dynamic data
	if err != nil {
		app.serverError(w, err)
	}
	// w.Write([]byte("Hello from Snippetbox"))
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d...\n", id)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Create a new snippet...."))
}

func (app *application) downloadHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./ui/static/js/main.js")
}

/*

technically speaking, a handler is something that implements the ServeHTTP method:

type home struct {}

func (h* home) (w http.ResponseWriter, r *http.Request) {
	...
}

mux.Handle("/asdasd", &home())

but saying
func home(w http.ResponseWriter, r *http.Request) {
	...
}
is not technically a handler: it's a function

hence, to make it into a handler, you would write the following:

mux.Handle("/", http.HandlerFunc(home)) // this automatically adds a ServeHTTP() method to home function, which contains the contents of the home function

but this is done for us using the HandleFunc method instead:
mux.HandleFunc("/", home)
*/
