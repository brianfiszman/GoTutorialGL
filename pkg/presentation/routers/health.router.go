package routers

import (
	"training/tutorialGL/pkg/application/controllers"

	"github.com/go-chi/chi"
)

type HealthRoutes struct {
	Router     *chi.Mux
	Controller controllers.HealthController
}

func NewHealthRoutes() (h HealthRoutes) {
	h.Router = chi.NewRouter()

	h.Router.Get("/", h.Controller.HealthChecks)
	return h
}
