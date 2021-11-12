package containers

type AppContainer struct {
	DatabaseContainer *DatabaseContainer
	ServerContainer   ServerContainer
	AuthContainer     AuthContainer
	HealthContainer   HealthContainer
	UserContainer     UserContainer
}

func NewAppContainer() AppContainer {
	var databaseContainer *DatabaseContainer = NewDatabaseContainer()
	var authContainer AuthContainer = NewAuthContainer()
	var healthContainer HealthContainer = NewHealthContainer()
	var userContainer UserContainer = NewUserContainer(databaseContainer.Database, authContainer.Router.JWTAuth)
	var serverContainer ServerContainer = NewServerContainer(userContainer.Router, healthContainer.Router, authContainer.Router)

	return AppContainer{
		DatabaseContainer: databaseContainer,
		AuthContainer:     authContainer,
		HealthContainer:   healthContainer,
		UserContainer:     userContainer,
		ServerContainer:   serverContainer,
	}
}
