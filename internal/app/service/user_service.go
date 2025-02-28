package service

import (
	"context"
	"errors"

	"github.com/ThePotatoVerse/internal/app/model"
	"github.com/ThePotatoVerse/internal/app/repository"
	"github.com/ThePotatoVerse/pkg/logger"
)

// Common errors
var (
	ErrUserNotFound = errors.New("user not found")
	ErrInvalidInput = errors.New("invalid input")
)

// UserService defines the interface for user business logic
type UserService interface {
	List(ctx context.Context) ([]model.User, error)
	Create(ctx context.Context, user model.User) (model.User, error)
	Get(ctx context.Context, id string) (model.User, error)
	Update(ctx context.Context, user model.User) error
	Delete(ctx context.Context, id string) error
}

// userService implements UserService
type userService struct {
	log      logger.Logger
	userRepo repository.UserRepository
}

// NewUserService creates a new user service
func NewUserService(log logger.Logger, userRepo repository.UserRepository) UserService {
	return &userService{
		log:      log,
		userRepo: userRepo,
	}
}

// List returns all users
func (s *userService) List(ctx context.Context) ([]model.User, error) {
	s.log.Info("Listing users")
	return s.userRepo.FindAll(ctx)
}

// Create creates a new user
func (s *userService) Create(ctx context.Context, user model.User) (model.User, error) {
	s.log.Info("Creating user")

	// Validate user
	if user.Name == "" {
		return model.User{}, ErrInvalidInput
	}

	return s.userRepo.Create(ctx, user)
}

// Get returns a user by ID
func (s *userService) Get(ctx context.Context, id string) (model.User, error) {
	s.log.Info("Getting user", "id", id)

	user, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return model.User{}, ErrUserNotFound
		}
		return model.User{}, err
	}

	return user, nil
}

// Update updates a user
func (s *userService) Update(ctx context.Context, user model.User) error {
	s.log.Info("Updating user", "id", user.ID)

	// Validate user
	if user.ID == "" || user.Name == "" {
		return ErrInvalidInput
	}

	// Check if user exists
	_, err := s.userRepo.FindByID(ctx, user.ID)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return ErrUserNotFound
		}
		return err
	}

	return s.userRepo.Update(ctx, user)
}

// Delete deletes a user
func (s *userService) Delete(ctx context.Context, id string) error {
	s.log.Info("Deleting user", "id", id)

	// Check if user exists
	_, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return ErrUserNotFound
		}
		return err
	}

	return s.userRepo.Delete(ctx, id)
}
