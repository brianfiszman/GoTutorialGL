package containers

import (
	"training/tutorialGL/pkg/interfaces"
	"training/tutorialGL/pkg/persistence/adapters"
	"training/tutorialGL/pkg/persistence/config"
)

type DatabaseContainer struct {
	Database interfaces.DatabaseInterface
}

func NewDatabaseContainer() *DatabaseContainer {
	var db *DatabaseContainer = &DatabaseContainer{
		Database: &adapters.PostgreSQL{
			Config: config.GetDatabaseConfig(),
		},
	}

	db.Database.Connect()

	return db
}
