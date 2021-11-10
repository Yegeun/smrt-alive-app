package main

import (
	"net/http"
	"github.com/bmizerany/pat" // New import
	"github.com/justinas/alice"
)
func (app *application) routes() http.Handler {
	// Use the http.NewServeMux() function to initialize a new servemux, then
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	mux := pat.New()
	mux.Get("/", http.HandlerFunc(app.home))
	mux.Get("/user/create", http.HandlerFunc(app.createSnippetForm))
	mux.Post("/user/create", http.HandlerFunc(app.createUser))
	mux.Get("/user/:id", http.HandlerFunc(app.showUser))
	// Moved down Use the http.ListenAndServe() function to start a new web server. We pas
	// two parameters: the TCP network address to listen on (in this case ":4000
	// and the servemux we just created. If http.ListenAndServe() returns an er
	// we use the log.Fatal() function to log the error message and exit.

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))
	// Return the 'standard' middleware chain followed by the servemux.
	return standardMiddleware.Then(mux)

	////create a file server for the static files this is the way without the library
	//fileServer := http.FileServer(http.Dir("./ui/static"))
	//
	//// use the mux handle function to register the filee server as the handler
	//mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	//
	//return app.recoverPanic(app.logRequest(secureHeaders(mux)))
}