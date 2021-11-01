package containers

import (
	"training/tutorialGL/pkg/application/controllers"
	"training/tutorialGL/pkg/domain/services"
	"training/tutorialGL/pkg/interfaces"
	"training/tutorialGL/pkg/persistence/repositories"
	"training/tutorialGL/pkg/presentation/routers"
)

type UserContainer struct {
	Router     routers.UserRoutes
	Controller controllers.UserController
	Service    services.UserService
	Repository repositories.UserRepository
}

func NewUserContainer(d interfaces.DatabaseInterface) UserContainer {
	var userRepository = &repositories.UserRepository{Database: d}
	var userService = services.UserService{Repository: *userRepository}
	var u UserContainer = UserContainer{
		Repository: *userRepository,
		Service:    userService,
		Router: routers.UserRoutes{
			Controller: controllers.UserController{
				Service: userService,
			},
		},
	}
	u.Router.Router = u.Router.NewUserRoutes()
	return u
}
