package routers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"training/tutorialGL/pkg/application/controllers"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

type HealthRouter struct {
	Router *chi.Mux
}

func (r *HealthRouter) NewHealthRouter() *chi.Mux {
	var port string = os.Getenv("HTTP_PORT")
	r.Router = chi.NewRouter()
	r.Router.Use(middleware.Logger)
	r.Router.Get("/", controllers.HealthChecks)
	log.Println("Servirn the app in the port:", port)
	err := http.ListenAndServe(":"+port, r.Router)
	if err != nil {
		fmt.Println(err)
	}
	return r.Router
}
