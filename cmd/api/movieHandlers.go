package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) getMovie(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.logger.Println(errors.New("invalid ID parameter"))
		app.errorJson(w, err)
		return
	}

	app.logger.Println(id)

	movie, err := app.models.DB.Get(id)
	if err != nil {
		app.logger.Println(err)
		app.errorJson(w, err)
		return
	}

	err = app.writeJson(w, http.StatusOK, movie, "data")
	if err != nil {
		app.logger.Println(err)
		app.errorJson(w, err)
		return
	}
}

func (app *application) getAllMovies(w http.ResponseWriter, r *http.Request) {

}
