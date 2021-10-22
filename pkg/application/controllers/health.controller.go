package controllers

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

type HealthController struct {
}

const (
	RESPONSE_MSG = "Hola Mundo"
)

func (HealthController) HealthChecks(rw http.ResponseWriter, r *http.Request) {
	json.NewEncoder(rw).Encode(Response{
		Status:  http.StatusOK,
		Message: RESPONSE_MSG,
	})
}
