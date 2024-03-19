package services

import (
	"context"
	"time"

	"github.com/SaiHtetMyatHtut/potatoverse/src/core/data/models"
	"github.com/SaiHtetMyatHtut/potatoverse/src/core/data/repositories"
	authschemas "github.com/SaiHtetMyatHtut/potatoverse/src/schemas/auth_schemas"
	userschemas "github.com/SaiHtetMyatHtut/potatoverse/src/schemas/user_schemas"
	"github.com/SaiHtetMyatHtut/potatoverse/src/utils"
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

func (us *UserService) CreateUser(ctx context.Context, userDTO authschemas.UserSignUpSchema) (models.User, error) {
	// TODO: add logic to check the username duplicates.
	newHashPassword, err := utils.HashPassword(userDTO.Password)

	if err != nil {
		return models.User{
			Username:       userDTO.Username,
			HashedPassword: userDTO.Password,
			CreatedAt:      time.Now().UTC(),
			LastLogin:      time.Now().UTC(),
		}, err
	}

	user, err := us.userRepository.Insert(ctx, models.User{
		Username:       userDTO.Username,
		HashedPassword: newHashPassword,
		CreatedAt:      time.Now().UTC(),
		LastLogin:      time.Now().UTC(),
	})

	if err != nil {
		return user, err
	}

	return user, nil
}

func (us *UserService) UpdateUser(ctx context.Context, userUpdateDTO userschemas.UpdateUserSchema) error {
	// TODO
	return nil
}
