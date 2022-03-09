package main

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/andrewchurchill/go-tutorial/models"
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

	movie := models.Movie{
		Id:          id,
		Title:       "Some movie",
		Description: "Some description",
		Year:        2021,
		ReleaseDate: time.Date(2021, 1, 1, 1, 0, 0, 0, time.Local),
		Runtime:     100,
		Rating:      5,
		MpaaRating:  "PG-13",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
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
