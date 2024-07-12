package main

import (
	"math/rand"
	"net/http"
	"strconv"

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
		"nextNumber": strconv.Itoa(rand.Intn(100)),
	}

	err := response.JSON(w, http.StatusOK, data)
	if err != nil {
		app.serverError(w, r, err)
	}
}
