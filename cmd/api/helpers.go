package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	
)

// Retrieve the "id" URL parameter from the current request context, then convert it to
// an integer and return it. If the operation isn't successful, return 0 and an error.
func (app *application) readIDParam(r *http.Request) (int64, error) {


	val:= r.PathValue("id")
	id, err := strconv.ParseInt(val, 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}
	return id, nil
}

func (app *application) writeJSON(w http.ResponseWriter,status int, data any, headers http.Header) error{
	js , err := json.Marshal(data)
	if err !=nil {
		return err

	}

	js = append(js,'\n')

	for key, value := range headers {
		w.Header()[key] =  value
	}
	w.Header().Set("Content-type","application/json")
	w.WriteHeader(status)
	w.Write(js)
	return nil
}
