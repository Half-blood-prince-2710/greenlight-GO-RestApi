package main

import (
	"net/http"


)

func (app *application) routes() *http.ServeMux {
	// declare a new servermux and add a /v1/healthcheck route
	mux := http.NewServeMux()
    mux.HandleFunc("/v1/healthcheck", app.healthcheckHandler)
    mux.HandleFunc("/v1/movies", app.createMovieHandler)
    mux.HandleFunc("/v1/movies/:id", app.showMovieHandler)

	return mux
}
