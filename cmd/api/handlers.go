package main

import (
	
	"fmt"
	"net/http"
	"time"

	"github.com/half-blood-prince-2710/greenlight-GO-RestApi/internal/data"
)

// Add a createMovieHandler for the "POST /v1/movies" endpoint. For now we simply
// return a plain-text placeholder response.
func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title string `json="title"`
		Year int16 `json="year"`
		Runtime int16 `json="runtime"`
		Genres []string `json="genres"`
	}
	err := app.readJSON(w, r, &input)
	if err !=nil{
		app.badRequestResponse(w, r, err)
		return
	}
	fmt.Fprintf(w, "%+v\n", input)
}

// Add a showMovieHandler for the "GET /v1/movies/:id" endpoint. For now, we retrieve
// the interpolated "id" parameter from the current URL and include it in a placeholder
// response.
func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Casablanca",
		Runtime:   102,
		Genres:    []string{"drama", "romance", "war"},
		Version:   1,
	}
	err = app.writeJSON(w,http.StatusOK,envelope{"movie":movie},nil)
	if err !=nil {
		app.serverErrorResponse(w, r, err)
	}
}

// Otherwise, interpolate the movie ID in a placeholder response.
