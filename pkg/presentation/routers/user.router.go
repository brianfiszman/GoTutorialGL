package routers

import (
	"training/tutorialGL/pkg/application/controllers"

	"github.com/go-chi/chi/v5"
)

type UserRoutes struct {
	Router     *chi.Mux
	Controller controllers.UserController
}

func (u UserRoutes) NewUserRoutes() *chi.Mux {
	u.Router = chi.NewRouter()

	u.Router.Post("/users", u.Controller.Create)
	u.Router.Get("/users", u.Controller.Get)
	return u.Router
}
