package containers

import (
	"training/tutorialGL/pkg/application/controllers"
	"training/tutorialGL/pkg/presentation/routers"
)

type HealthContainer struct {
	Router     routers.HealthRouter
	Controller controllers.HealthController
}

func NewHealthContainer() HealthContainer {
	var h HealthContainer = HealthContainer{
		Controller: controllers.HealthController{},
		Router:     routers.NewHealthRouter(),
	}

	h.Router.HealthController = h.Controller

	return h
}
