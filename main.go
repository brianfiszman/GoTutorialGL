package main

import "training/tutorialGL/pkg/containers"

func main() {
	var app containers.AppContainer = containers.NewAppContainer()

	// Server initializes listening
	app.HealthContainer.Routers.NewHealthRouter()
}
