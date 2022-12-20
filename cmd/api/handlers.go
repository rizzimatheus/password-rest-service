package main

import (
	"net/http"
)

func (app *Config) verifyPassword(w http.ResponseWriter, r *http.Request) {
	var requestPayload Verifier

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	nomatch := requestPayload.verifyRules()

	verify := len(nomatch) == 0

	payload := jsonResponse{
		Verify:  verify,
		NoMatch: nomatch,
	}

	app.writeJSON(w, http.StatusOK, payload)
}
