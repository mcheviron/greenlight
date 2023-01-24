package main

import (
	"net/http"
)

func (a *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	data := envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": a.config.env,
			"version":     version,
		},
	}
	err := a.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		a.logger.Println(err)
		a.serverErrorResponse(w, r, err)
		return
	}
}
