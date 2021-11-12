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
		Router: routers.HealthRoutes{
			Controller: controllers.HealthController{},
		},
	}

	h.Router.Controller = h.Controller

	return h
}
