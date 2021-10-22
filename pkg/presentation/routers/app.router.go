package routers

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
)

type AppRouter struct {
	HealthRouter HealthRouter
	Router       *chi.Mux
}

func NewAppRouter(port string) (r AppRouter) {
	r.HealthRouter = NewHealthRouter()
	r.Router = chi.NewRouter()
	r.Router.Use(middleware.Logger)
	r.Router.Use(middleware.StripSlashes)
	r.Router.Mount("/", r.HealthRouter.Router)
	return r
}
