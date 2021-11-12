package controllers

import (
	"encoding/json"
	"net/http"
	"training/tutorialGL/pkg/application/dtos"

	"github.com/go-chi/jwtauth/v5"
)

type AuthController struct {
	TokenAuth *jwtauth.JWTAuth
}

func (a AuthController) Authenticate(rw http.ResponseWriter, r *http.Request) {

	rw.Header().Add("content-type", "application/json")

	var body dtos.JWTAuth

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		http.Error(rw, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	if body.Data != "admin" {
		http.Error(rw, "Data is incorrect", http.StatusBadRequest)
	}

	_, tokenString, _ := a.TokenAuth.Encode(map[string]interface{}{"data": "admin"})

	if err != nil {
		http.Error(rw, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
	}

	json.NewEncoder(rw).Encode(dtos.BasicResponseDTO{
		Status:  http.StatusOK,
		Message: "Authenticated",
		Data:    tokenString,
	})
}
