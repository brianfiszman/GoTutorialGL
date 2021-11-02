package controllers

import (
	"encoding/json"
	"net/http"
	"training/tutorialGL/pkg/application/dtos"
	"training/tutorialGL/pkg/domain/services"
)

type UserController struct {
	Service services.UserService
}

func (u UserController) Create(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	createUserDto := dtos.CreateUserDTO{}

	json.NewDecoder(r.Body).Decode(&createUserDto)

	// Call service to create new user
	res, err := u.Service.Create(ctx, createUserDto)

	if err != nil {
		http.Error(rw, http.StatusText(http.StatusForbidden), http.StatusForbidden)
	}

	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(res)
}

func (u UserController) Get(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Call service to get users
	res, err := u.Service.Get(ctx)

	if err != nil {
		http.Error(rw, http.StatusText(http.StatusForbidden), http.StatusForbidden)
	}

	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(res)
}
