package main

import (
	"net/http"


)

func (app *application) routes() http.Handler {
	// declare a new servermux and add a /v1/healthcheck route
	mux := http.NewServeMux()

	
    mux.HandleFunc("GET /v1/healthcheck", app.healthcheckHandler)
    mux.HandleFunc("POST /v1/movies", app.createMovieHandler)
    mux.HandleFunc("GET /v1/movies/{id}", app.showMovieHandler)
	mux.HandleFunc("PUT /v1/movies/{id}",app.updateMovieHandler)

	return app.recoverPanic(mux)
	
}
