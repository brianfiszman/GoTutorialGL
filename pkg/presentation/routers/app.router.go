package routers

import (
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

type AppRouter struct {
	Router       *chi.Mux
	UserRoutes   UserRoutes
	HealthRoutes HealthRoutes
	AuthRoutes   AuthRoutes
}

func (r *AppRouter) NewAppRouter() *chi.Mux {
	r.Router = chi.NewRouter()

	r.Router.Use(middleware.Logger)
	r.Router.Use(middleware.StripSlashes)

	r.Router.Mount("/health", r.HealthRoutes.NewHealthRoutes())

	r.Router.Mount("/auth", r.AuthRoutes.NewAuthRoutes())

	r.Router.Mount("/users", r.UserRoutes.NewUserRoutes())

	return r.Router
}
