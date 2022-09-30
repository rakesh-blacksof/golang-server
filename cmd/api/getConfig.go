package main

import (
	"net/http"
)

func (app *application) getAppConfig(w http.ResponseWriter, r *http.Request) {
	resData := map[string]string{
		"status":      "online",
		"environment": app.config.env,
		"version":     version,
	}
	err := app.writeJSON(w, http.StatusOK, envelope{"app_config": resData}, nil)
	if err != nil {
		app.logger.Println("Error", err)
		app.errorInternalServer(w, r, err)
	}
}
