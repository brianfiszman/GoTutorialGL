package containers

import (
	"os"
	"training/tutorialGL/pkg/application/server"
	"training/tutorialGL/pkg/presentation/routers"
)

type ServerContainer struct {
	Server *server.Server
}

func NewServerContainer(userRoutes routers.UserRoutes, healthRoutes routers.HealthRoutes) (s ServerContainer) {
	var httpConnector *routers.AppRouter = &routers.AppRouter{
		UserRoutes:   userRoutes,
		HealthRoutes: healthRoutes,
	}
	httpConnector.Router = httpConnector.NewAppRouter()

	var serverContainer ServerContainer = ServerContainer{&server.Server{
		AppRouter: httpConnector,
		HTTP_PORT: os.Getenv("HTTP_PORT"),
	}}
	return serverContainer
}
