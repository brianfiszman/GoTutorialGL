package server

import (
	"log"
	"net/http"
	"training/tutorialGL/pkg/presentation/routers"
)

type Server struct {
	HTTP_PORT string
	AppRouter *routers.AppRouter
}

func (s *Server) Run() {
	log.Println("Listening on port:", s.HTTP_PORT)
	err := http.ListenAndServe(":"+s.HTTP_PORT, s.AppRouter.Router)
	if err != nil {
		log.Fatal(err)
	}
}
