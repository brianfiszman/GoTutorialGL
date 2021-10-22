package routers

import (
	"training/tutorialGL/pkg/application/controllers"

	"github.com/go-chi/chi"
)

type HealthRouter struct {
	Router           *chi.Mux
	HealthController controllers.HealthController
}

func NewHealthRouter() (h HealthRouter) {
	h.Router = chi.NewRouter()
	h.Router.Get("/", h.HealthController.HealthChecks)
	return h
}
