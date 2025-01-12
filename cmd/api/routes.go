package main

import (
	"net/http"


)

func (app *application) routes() http.Handler {
	// declare a new servermux and add a /v1/healthcheck route
	mux := http.NewServeMux()

	
    mux.HandleFunc("GET /v1/healthcheck", app.healthcheckHandler)

	mux.HandleFunc("GET /v1/movies",app.listMoviesHandler)
    mux.HandleFunc("POST /v1/movies", app.createMovieHandler)
    mux.HandleFunc("GET /v1/movies/{id}", app.showMovieHandler)
	mux.HandleFunc("PATCH /v1/movies/{id}",app.updateMovieHandler)
	mux.HandleFunc("DELETE /vq/movies/{id}", app.deleteMovieHandler)
	return app.recoverPanic(mux)
	
}
