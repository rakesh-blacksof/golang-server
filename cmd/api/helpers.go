package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) readIDParams(r *http.Request) (int64, error) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("Invalid ID Parameter")
	}
	return id, nil
}

type envelope map[string]interface{}

func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	resBytes, err := json.Marshal(data)
	if err != nil {
		return errors.New("Error Encoding Data. Please check the data passed.")
	}
	resBytes = append(resBytes, '\n')
	// if we have reached here without any errors, We are good to go withs sending the data
	for key, value := range headers {
		w.Header()[key] = value
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(resBytes)
	return nil
}

func (app *application) readJSON(w http.ResponseWriter, r *http.Request, destination interface{}) error {
	// features to add :
	// limiting size of json , ensuring that only json object is being sent, only known keys are provided.

	decode := json.NewDecoder(r.Body) // decoder fields.
	decode.DisallowUnknownFields()    // disallow unknown fields.
	err := decode.Decode(destination)
	if err != nil {
		// start the triage here. We will start looping
		// through the error types and pushing the error.

		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError

		switch {
		case errors.As(err, &syntaxError):
			return fmt.Errorf("Body contains badly fomatted JSON (at character %d)", syntaxError.Offset)
		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("Body contains badly formatted JSON") // generic error to catch any JSON fomatting error.
		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf("Body contains incorrect JSON type for field %s ", unmarshalTypeError.Field)
			}
			return fmt.Errorf("Body contains incorrect JSON type for field %d ", unmarshalTypeError.Offset)
		case errors.Is(err, io.EOF):
			return errors.New("body must not be empty")
		case errors.As(err, &invalidUnmarshalError):
			panic(err)
		default:
			return err
		}
	}
	return nil
}

func (app *application) handleContact(http.ResponseWriter, *http.Request) {

}
