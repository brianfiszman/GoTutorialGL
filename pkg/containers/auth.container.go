package containers

import (
	"os"
	"training/tutorialGL/pkg/application/controllers"
	"training/tutorialGL/pkg/presentation/routers"

	"github.com/go-chi/jwtauth/v5"
)

type AuthContainer struct {
	Router     routers.AuthRoutes
	Controller controllers.AuthController
	JWTAuth    *jwtauth.JWTAuth
}

func NewAuthContainer() AuthContainer {
	var tokenAuth = jwtauth.New("HS256", []byte(os.Getenv("JWT_SECRET")), nil)

	var a AuthContainer = AuthContainer{
		Router: routers.AuthRoutes{
			Controller: controllers.AuthController{
				TokenAuth: tokenAuth,
			},
			JWTAuth: tokenAuth,
		},
	}

	return a
}
