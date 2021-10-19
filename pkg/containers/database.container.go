package containers

import (
	"training/tutorialGL/pkg/persistence/adapters"
	"training/tutorialGL/pkg/persistence/config"
)

type DatabaseContainer struct {
	Database Database
}

type Database interface {
	Connect()
	// Insert()
	// Select()
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
