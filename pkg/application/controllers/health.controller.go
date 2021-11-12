package controllers

import (
	"encoding/json"
	"net/http"
	"training/tutorialGL/pkg/application/dtos"
)

type HealthController struct {
}

const (
	RESPONSE_MSG = "IS ALIVE"
)

func (HealthController) HealthChecks(rw http.ResponseWriter, r *http.Request) {
	json.NewEncoder(rw).Encode(dtos.BasicResponseDTO{
		Status:  http.StatusOK,
		Message: RESPONSE_MSG,
	})
}
