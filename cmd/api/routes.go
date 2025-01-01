package main

import (

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	// declare a new servermux and add a /v1/healthcheck route
	router := httprouter.New()
    router.GET("/v1/healthcheck", app.healthcheckHandler)
    router.POST("/v1/movies", app.createMovieHandler)
    router.GET("/v1/movies/:id", app.showMovieHandler)

	return router
}
