package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	// Use the http.NewServeMux() function to initialize a new servemux, then
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/user", app.showUser)
	mux.HandleFunc("/user/create", app.createUser)
	// Use the http.ListenAndServe() function to start a new web server. We pas
	// two parameters: the TCP network address to listen on (in this case ":4000
	// and the servemux we just created. If http.ListenAndServe() returns an er
	// we use the log.Fatal() function to log the error message and exit.

	//create a file server for the static files
	fileServer := http.FileServer(http.Dir("./ui/static"))

	// use the mux handle function to register the filee server as the handler
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux
}