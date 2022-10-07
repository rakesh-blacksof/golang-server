package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(app.errorNotFound)
	router.MethodNotAllowed = http.HandlerFunc(app.errorInvaidMethod)                    // allow header is set automatically
	router.HandlerFunc(http.MethodGet, "/v1/config", app.getAppConfig)                   // app config
	router.HandlerFunc(http.MethodGet, "/v1/movies", app.getAllMovies)                   // READ all movies
	router.HandlerFunc(http.MethodPost, "/v1/movies", app.addMovie)                      // CREATE
	router.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.getMovieByID)               // READ
	router.HandlerFunc(http.MethodGet, "/v1/form/submission", app.contactFormSubmission) // READ
	router.ServeFiles("/v1/static/*filepath", http.Dir("static"))                        // serve static assets
	// router.HandlerFunc(http.MethodPut, "/v1/movies/:id", app.editMovie)      // UPDATE
	// router.HandlerFunc(http.MethodDelete, "/v1/movies/:id", app.deleteMovie) // DELETE
	// there are some scenorios where simple plain text error responses will
	// be sent to the client e.g. using bad host , bad origin methods.
	// we need not worry about it as these are non-customizable.
	return app.middlewareRateLimiter(router)
}
