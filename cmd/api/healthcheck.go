package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request,_ httprouter.Params){
	js:= `{"status":"available","environment":%q,"version":%q}`
	js = fmt.Sprintf(js,app.config.env,version)

	w.Header().Set("Content-Type","application/json")
	w.Write([]byte(js))
}
