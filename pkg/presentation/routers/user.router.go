package routers

import (
	"training/tutorialGL/pkg/application/controllers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

type UserRoutes struct {
	Router     *chi.Mux
	Controller controllers.UserController
	TokenAuth  *jwtauth.JWTAuth
}

func (u UserRoutes) NewUserRoutes() *chi.Mux {
	u.Router = chi.NewRouter()

	u.Router.Post("/", u.Controller.Create)

	u.Router.Group(func(router chi.Router) {

		router.Use(jwtauth.Verifier(u.TokenAuth))

		router.Use(jwtauth.Authenticator)

		router.Get("/", u.Controller.Get)
	})

	return u.Router
}
