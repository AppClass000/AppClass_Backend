package containers

import (
	"backend/internal/app/repositories"
	"backend/internal/app/services"
	"backend/internal/database"
	"backend/internal/handlers"
)

type AppContainer struct {
	UserRepository    repositories.UserRepository
	UserServise       services.UserServise
	UserHandler       handlers.UserHandler
	ClassesRepository repositories.ClassesRepository
	ClassesServise    services.ClassesServise
	ClassesHandler    handlers.ClassesHandler
}

func NewAppContainer() *AppContainer {
	database.InitDB()
	db := database.GetDB()

	userrepository := repositories.NewUserRepository(db)
	userservise := services.NewUserServise(&userrepository)
	userhandler := handlers.NewUserHandler(userservise)

	classesrepository := repositories.NewClassRepository(db)
	classesservice := services.NewClassesServise(classesrepository)
	classeshandler := handlers.NewClassesHandler(classesservice)

	return &AppContainer{
		UserRepository:    userrepository,
		UserServise:       userservise,
		UserHandler:       userhandler,
		ClassesRepository: classesrepository,
		ClassesServise:    classesservice,
		ClassesHandler:    classeshandler,
	}
}
