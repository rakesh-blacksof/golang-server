package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/rakesh-gupta29/movie-server/internal/data"
)

func (app *application) getAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("showing 10 latest movies"))
}

func (app *application) getMovieByID(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParams(r)
	if err != nil {
		app.errorNotFound(w, r)
		return
	}
	movie := data.Movie{
		PublicID:  id,
		PrivateID: 10000,
		Title:     "Mizrapur",
		CreatedAt: time.Now(),
		Runtime:   100,
		Year:      2018,
		Genres:    []string{"drama", "crime", "action"},
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		app.logger.Fatal("Error", err)
		app.errorInternalServer(w, r, err)
	}
}

func (app *application) addMovie(w http.ResponseWriter, r *http.Request) {
	// parsing and validaiton the JSON request.
	// anonymous struct to behave as a target for the request body.
	var reqData struct {
		Title   string   `json:"title"`
		Year    int32    `json:"year"`
		Runtime int32    `json:"runtime_in_mins"`
		Genres  []string `json:"genres"`
	} // other fields are to be generated by the sever itself.

	// what is to be done if  user sends anonymous data( xml, random bytes , malformed error,wrong keys or no body at all.)
	// since default errors thrown by the Decode package is not descriptive enough or can ecpose underlying API info therefore ,
	// it is quite inmortant for to target particular error, format it and send proper response.
	// this is quite important for Public API.
	// for a list of errors and what is to be done can be referenced from the doc files.
	err := app.readJSON(w, r, &reqData)
	if err != nil {
		app.errorBadRequest(w, r, err)
		return
	}
	fmt.Fprintf(w, "%+v\n", reqData)
}

func (app *application) contactFormSubmission(w http.ResponseWriter, r *http.Request) {
	fromEmail := os.Getenv("fromEmail")
	resString := "sending mail from " + fromEmail
	w.Write([]byte(resString))
}
