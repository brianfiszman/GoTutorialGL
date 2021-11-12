package routers

import (
	"training/tutorialGL/pkg/application/controllers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

type AuthRoutes struct {
	Router     *chi.Mux
	Controller controllers.AuthController
	JWTAuth    *jwtauth.JWTAuth
}

func (a AuthRoutes) NewAuthRoutes() *chi.Mux {

	a.Router = chi.NewRouter()

	a.Router.Post("/", a.Controller.Authenticate)

	return a.Router
}
