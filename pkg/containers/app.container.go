package containers

type AppContainer struct {
	DatabaseContainer *DatabaseContainer
	HealthContainer   HealthContainer
	ServerContainer   ServerContainer
	UserContainer     UserContainer
}

func NewAppContainer() AppContainer {
	var databaseContainer *DatabaseContainer = NewDatabaseContainer()
	var healthContainer HealthContainer = NewHealthContainer()
	var userContainer UserContainer = NewUserContainer(databaseContainer.Database)
	var serverContainer ServerContainer = NewServerContainer(userContainer.Router, healthContainer.Router)

	return AppContainer{
		DatabaseContainer: databaseContainer,
		HealthContainer:   healthContainer,
		UserContainer:     userContainer,
		ServerContainer:   serverContainer,
	}
}
