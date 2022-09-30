package main

import "net/http"

func (app *application) logError(r *http.Request, err error) {
	app.logger.Println(err)
	// log additional info of the error
}

func (app *application) errorWriteJson(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	env := envelope{"error": message}
	err := app.writeJSON(w, status, env, nil)
	if err != nil {
		app.logError(r, err)
		w.WriteHeader(500)
	}
}

func (app *application) errorInternalServer(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)
	// this is for logging the errors. otherwise no indicaiton is to be
	//given to client as to where the server went wrong.

	message := "Server encountered a problem while processing your request."
	app.errorWriteJson(w, r, http.StatusInternalServerError, message)
}

func (app *application) errorNotFound(w http.ResponseWriter, r *http.Request) {
	message := "Requested resource could not be found"
	app.errorWriteJson(w, r, http.StatusNotFound, message)
}
func (app *application) errorInvaidMethod(w http.ResponseWriter, r *http.Request) {
	message := "Requested method is not supported for this resource."
	app.errorWriteJson(w, r, http.StatusMethodNotAllowed, message)
}

func (app *application) errorBadRequest(w http.ResponseWriter, r *http.Request, err error) {
	app.errorWriteJson(w, r, http.StatusBadRequest, err.Error())
}

func (app *application) ErrorTooManyRequests(w http.ResponseWriter, r *http.Request) {
	app.errorWriteJson(w, r, http.StatusTooManyRequests, "You have exceeded your limit for requests.")
}
