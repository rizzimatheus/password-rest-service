package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var app Config

func Test_verifyPassword(t *testing.T) {
	table := []struct {
		name string
		requestBody string
		responseBody string
	}{
		{
			"default test", 
			`{"password":"TesteSenhaForte!123&","rules":[{"rule": "minSize","value": 8},{"rule": "minSpecialChars","value": 2},{"rule": "noRepeted","value": 0},{"rule": "minDigit","value": 4}]}`, 
			`{"verify":false,"noMatch":["minDigit"]}`,
		},
		{
			"empty password", 
			`{"password":"","rules":[{"rule": "minSize","value": 21},{"rule": "minUppercase","value": 4},{"rule": "noRepeted","value": 0}]}`, 
			`{"verify":false,"noMatch":["minSize","minUppercase"]}`,
		},
		{
			"ok password", 
			`{"password":"SenhaForte!123&","rules":[{"rule": "minSize","value": 8},{"rule": "minSpecialChars","value": 2},{"rule": "noRepeted","value": 0}]}`, 
			`{"verify":true,"noMatch":[]}`,
		},
		{
			"invalid rule", 
			`{"password":"SenhaForte!123&","rules":[{"rule": "minSpecialChars","value": 2},{"rule": "rule123","value": 2},{"rule": "noRepeted","value": 0}]}`, 
			`{"verify":false,"noMatch":["Error: 'rule123' is an invalid rule"]}`,
		},
		{
			"invalid request", 
			`"rules":[{"rule": "minSpecialChars","value": 2}]}`, 
			`{"verify":false,"noMatch":["Error reading json"]}`,
		},
	}

	for _, expected := range table {
		var reader io.Reader
		reader = strings.NewReader(expected.requestBody)
		req, _ := http.NewRequest("POST", "/verify", reader)
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(app.verifyPassword)

		handler.ServeHTTP(rr, req)

		if expected.responseBody != rr.Body.String() {
			t.Errorf("%s: returned wrong response body\nExpected: %s\nGot: %s", expected.name, expected.responseBody, rr.Body)
		}
	}
}