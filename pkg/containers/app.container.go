package containers

type AppContainer struct {
	DatabaseContainer *DatabaseContainer
	HealthContainer   HealthContainer
}

func NewAppContainer() AppContainer {
	var databaseContainer *DatabaseContainer = NewDatabaseContainer()
	var healthContainer HealthContainer = NewHealthContainer()
	return AppContainer{
		DatabaseContainer: databaseContainer,
		HealthContainer:   healthContainer,
	}
}
