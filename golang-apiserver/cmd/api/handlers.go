package main

import (
	"log"
	"net/http"
	"strconv"

	"golang-apiserver/internal/number"
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
	numberStr := r.URL.Query().Get("number")
	if numberStr == "" {
		http.Error(w, "Missing 'number' query parameter", http.StatusBadRequest)
		return
	}

	numberValue, err := strconv.Atoi(numberStr)
	if err != nil {
		http.Error(w, "Invalid 'number' query parameter", http.StatusBadRequest)
		return
	}

	log.Println("Sleeping...")
	nextNumber := number.Sleep(numberValue)

	data := map[string]int{
		"nextNumber": nextNumber,
	}

	err = response.JSON(w, http.StatusOK, data)
	if err != nil {
		app.serverError(w, r, err)
	}
}
