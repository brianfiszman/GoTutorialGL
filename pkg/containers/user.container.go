package containers

import (
	"training/tutorialGL/pkg/application/controllers"
	"training/tutorialGL/pkg/domain/services"
	"training/tutorialGL/pkg/interfaces"
	"training/tutorialGL/pkg/persistence/repositories"
	"training/tutorialGL/pkg/presentation/routers"

	"github.com/go-chi/jwtauth/v5"
)

type UserContainer struct {
	Router     routers.UserRoutes
	Controller controllers.UserController
	Service    services.UserService
	Repository repositories.UserRepository
}

func NewUserContainer(d interfaces.DatabaseInterface, t *jwtauth.JWTAuth) UserContainer {
	var userRepository = &repositories.UserRepository{Database: d}
	var userService = services.UserService{Repository: *userRepository}
	var u UserContainer = UserContainer{
		Repository: *userRepository,
		Service:    userService,
		Router: routers.UserRoutes{
			Controller: controllers.UserController{
				Service: userService,
			},
			TokenAuth: t,
		},
	}
	u.Router.TokenAuth = t
	u.Router.Router = u.Router.NewUserRoutes()
	return u
}
