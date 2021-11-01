package routers

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type AppRouter struct {
	Router     *chi.Mux
	UserRoutes UserRoutes
	// HealthRoutes HealthRoutes
}

func (r *AppRouter) NewAppRouter() *chi.Mux {
	// r.HealthRoutes = NewHealthRoutes()
	r.Router = chi.NewRouter()

	r.Router.Use(middleware.Logger)
	r.Router.Use(middleware.StripSlashes)

	r.Router.Mount("/users", r.UserRoutes.Router)
	// r.Router.Mount("/", r.HealthRoutes.Router)
	return r.Router
}
