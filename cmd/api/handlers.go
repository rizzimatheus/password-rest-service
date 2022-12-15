package main

import (
	"fmt"
	"net/http"
)

func (app *Config) VerifyPassword(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		Password string  `json:"password"`
		Rules    []Rules `json:"rules"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	nomatch := []string{}
	for _, rule := range requestPayload.Rules {
		switch rule.Rule {
		case "minSize":
			if !minSize(requestPayload.Password, rule.Value) {
				nomatch = append(nomatch, "minSize")
			}
		case "minUppercase":
			if !minUppercase(requestPayload.Password, rule.Value) {
				nomatch = append(nomatch, "minUppercase")
			}
		case "minLowercase":
			if !minLowercase(requestPayload.Password, rule.Value) {
				nomatch = append(nomatch, "minLowercase")
			}
		case "minDigit":
			if !minDigit(requestPayload.Password, rule.Value) {
				nomatch = append(nomatch, "minDigit")
			}
		case "minSpecialChars":
			if !minSpecialChars(requestPayload.Password, rule.Value) {
				nomatch = append(nomatch, "minSpecialChars")
			}
		case "noRepeted":
			if !noRepeted(requestPayload.Password) {
				nomatch = append(nomatch, "noRepeted")
			}
		default:
			nomatch = append(nomatch, fmt.Sprintf("Error: '%s' is an invalid rule", rule.Rule))
		}
	}

	verify := len(nomatch) == 0

	payload := jsonResponse {
		Verify:  verify,
		NoMatch: nomatch,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

