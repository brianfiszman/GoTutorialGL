package routers

import (
	"training/tutorialGL/pkg/application/controllers"

	"github.com/go-chi/chi/v5"
)

type HealthRoutes struct {
	Router     *chi.Mux
	Controller controllers.HealthController
}

func (h HealthRoutes) NewHealthRoutes() *chi.Mux {
	h.Router = chi.NewRouter()

	h.Router.Get("/", h.Controller.HealthChecks)
	return h.Router
}
