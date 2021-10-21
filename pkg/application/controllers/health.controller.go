package controllers

import (
	"encoding/json"
	"net/http"
)

const (
	response_message = "Hola Mundo"
)

func HealthChecks(rw http.ResponseWriter, r *http.Request) {
	json.NewEncoder(rw).Encode(response_message)
}
