package services

import (
	"context"

	"github.com/SaiHtetMyatHtut/potatoverse/src/core/data/models"
	"github.com/SaiHtetMyatHtut/potatoverse/src/core/data/repositories"
	"go.uber.org/dig"
)

type UserService struct {
	userRepository repositories.UserRepository
}

type UserServiceDependencies struct {
	dig.In

	UserRepository repositories.UserRepository `name:"UserRepository"`
}

func NewUserService(deps UserServiceDependencies) UserService {
	return UserService{
		userRepository: deps.UserRepository,
	}
}

func (us *UserService) GetAllUsers(ctx context.Context) ([]models.User, error) {
	return us.userRepository.ReadAll(ctx)
}

func (us *UserService) GetUserById(ctx context.Context, id int64) (models.User, error) {
	return us.userRepository.ReadByID(ctx, id)
}
