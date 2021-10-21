package containers

type AppContainer struct {
	DatabaseContainer *DatabaseContainer
	HealthContainer   HealthContainer
	HealthRouter      *HealthRouter
}

func NewAppContainer() AppContainer {
	var databaseContainer *DatabaseContainer = NewDatabaseContainer()
	var healthContainer HealthContainer = NewHealthContainer()
	var healthRouter *HealthRouter = NewHealthRouter()
	return AppContainer{
		DatabaseContainer: databaseContainer,
		HealthContainer:   healthContainer,
	}
}
