package main

import (
	"net/http"

	"golang-apiserver/internal/response"
)

func (app *application) status(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"Status": "OK",
	}

	err := response.JSON(w, http.StatusOK, data)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) nextNumber(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"nextNumber": "50",
	}

	err := response.JSON(w, http.StatusOK, data)
	if err != nil {
		app.serverError(w, r, err)
	}
}
