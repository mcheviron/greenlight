package main

import (
	"net/http"
	"strings"
)

func (a *application) logError(r *http.Request, err error) {
	a.logger.Println(err)
}

func (a *application) errorResponse(
	w http.ResponseWriter,
	r *http.Request,
	message any,
	status int,
) {
	env := envelope{"error": message}
	err := a.writeJSON(w, status, env, nil)
	if err != nil {
		a.logError(r, err)
		http.Error(
			w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
	}
}

func (a *application) serverErrorResponse(
	w http.ResponseWriter,
	r *http.Request,
	err error,
) {
	a.logError(r, err)
	a.errorResponse(
		w,
		r,
		// Converting to lowercase isn't necessary but it's more JSON-like.
		strings.ToLower(http.StatusText(http.StatusInternalServerError)),
		http.StatusInternalServerError,
	)
}

func (a *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	a.errorResponse(
		w,
		r,
		strings.ToLower(http.StatusText(http.StatusNotFound)),
		http.StatusNotFound,
	)
}

func (a *application) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	a.errorResponse(
		w,
		r,
		strings.ToLower(http.StatusText(http.StatusMethodNotAllowed)),
		http.StatusMethodNotAllowed,
	)
}

func (a *application) badRequestResponse(
	w http.ResponseWriter,
	r *http.Request,
	err error,
) {
	a.errorResponse(
		w,
		r,
		err.Error(),
		http.StatusBadRequest,
	)
}

func (a *application) failedValidationResponse(
	w http.ResponseWriter,
	r *http.Request,
	errors map[string]string,
) {
	a.errorResponse(w, r, errors, http.StatusUnprocessableEntity)
}
