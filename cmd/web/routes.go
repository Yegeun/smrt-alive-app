package main

import (
	"net/http"

	"github.com/bmizerany/pat" // New import
	"github.com/justinas/alice"
)
func (app *application) routes() http.Handler {
	// Use the http.NewServeMux() function to initialize a new servemux, then
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	// Create a new middleware chain containing the middleware specific to
	// our dynamic application routes. For now, this chain will only contain
	// the session middleware but we'll add more to it later.
	dynamicMiddleware := alice.New(app.session.Enable)

	mux := pat.New()
	// Update these routes to use the new dynamic middleware chain followed
	// by the appropriate handler function.
	mux.Get("/", dynamicMiddleware.ThenFunc(app.home))
	mux.Get("/user/create", dynamicMiddleware.ThenFunc(app.createSnippetForm))
	mux.Post("/user/create", dynamicMiddleware.ThenFunc(app.createUser))
	mux.Get("/user/:id", dynamicMiddleware.ThenFunc(app.showUser))

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