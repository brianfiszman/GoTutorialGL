package containers

type AppContainer struct {
	DatabaseContainer *DatabaseContainer
	HealthContainer   HealthContainer
	ServerContainer   ServerContainer
}

func NewAppContainer() AppContainer {
	var databaseContainer *DatabaseContainer = NewDatabaseContainer()
	var healthContainer HealthContainer = NewHealthContainer()
	var serverContainer ServerContainer = NewServerContainer()

	return AppContainer{
		DatabaseContainer: databaseContainer,
		HealthContainer:   healthContainer,
		ServerContainer:   serverContainer,
	}
}
