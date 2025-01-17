package main

import (
	"net/http"


)

func (app *application) routes() http.Handler {
	// declare a new servermux and add a /v1/healthcheck route
	mux := http.NewServeMux()

	
    mux.HandleFunc("GET /v1/healthcheck", app.healthcheckHandler)

	mux.HandleFunc("GET /v1/movies",app.requirePermissions("movies:read",app.listMoviesHandler))
    mux.HandleFunc("POST /v1/movies", app.requirePermissions("movies:write",app.createMovieHandler))
    mux.HandleFunc("GET /v1/movies/{id}", app.requirePermissions("movies:read",app.showMovieHandler))
	mux.HandleFunc("PATCH /v1/movies/{id}",app.requirePermissions("movies:write",app.updateMovieHandler))
	mux.HandleFunc("DELETE /v1/movies/{id}", app.requirePermissions("movies:write",app.deleteMovieHandler))


	mux.HandleFunc("POST /v1/users",app.registerUserHandler)
	mux.HandleFunc( " PUT/v1/users/activated", app.activateUserHandler)

	mux.HandleFunc("POST /v1/tokens/authentication", app.createAuthenticationTokenHandler)

	
	return app.recoverPanic(app.rateLimit(app.authenticate(mux)))
	
}

