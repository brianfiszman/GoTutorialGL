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
	var http *routers.AppRouter = &routers.AppRouter{
		UserRoutes: userRoutes,
	}
	http.Router = http.NewAppRouter()

	var serverContainer ServerContainer = ServerContainer{&server.Server{
		AppRouter: http,
		HTTP_PORT: os.Getenv("HTTP_PORT"),
	}}
	return serverContainer
}
