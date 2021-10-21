package containers

import (
	"training/tutorialGL/pkg/presentation/routers"
)

type HealthContainer struct {
	Routers routers.HealthRouter
}

func NewHealthContainer() HealthContainer {

	var h HealthContainer = HealthContainer{}
	h.Routers.Router = h.Routers.NewHealthRouter()

	return h
}
