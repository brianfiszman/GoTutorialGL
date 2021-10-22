package containers

import (
	"training/tutorialGL/pkg/application/server"
)

type ServerContainer struct {
	Server server.Server
}

func NewServerContainer() (s ServerContainer) {
	s.Server = server.NewServer()
	return s
}
