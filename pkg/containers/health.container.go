package containers

import (
	"training/tutorialGL/pkg/application/controllers"
	"training/tutorialGL/pkg/presentation/routers"
)

type HealthContainer struct {
	Router     routers.HealthRoutes
	Controller controllers.HealthController
}

func NewHealthContainer() HealthContainer {
	var h HealthContainer = HealthContainer{
		Controller: controllers.HealthController{},
		Router:     routers.NewHealthRoutes(),
	}

	h.Router.Controller = h.Controller

	return h
}
