package di

import (
	"log"

	"github.com/SaiHtetMyatHtut/potatoverse/src/controllers"
	"github.com/SaiHtetMyatHtut/potatoverse/src/core/data/repositories"
	"github.com/SaiHtetMyatHtut/potatoverse/src/core/domain/services"
	"github.com/SaiHtetMyatHtut/potatoverse/src/db"
	"go.uber.org/dig"
)

type dependency struct {
	Construstor interface{}
	Token       string
}

func Initialize() *dig.Container {
	deps := []dependency{
		{
			Construstor: db.NewRedisClient,
			Token:       "Redis",
		},
		{
			Construstor: repositories.NewNewUserRepository,
			Token:       "UserRepository",
		},
		{
			Construstor: services.NewUserService,
			Token:       "UserService",
		},
		{
			Construstor: controllers.NewUserController,
			Token:       "UserController",
		},
	}
	container := dig.New()

	for _, dep := range deps {
		err := container.Provide(
			dep.Construstor,
			dig.Name(dep.Token),
		)
		if err != nil {
			log.Fatal(err)
		}
	}

	return container
}
