package repository

import (
	"context"
	"errors"

	"github.com/ThePotatoVerse/internal/app/model"
)

// Common errors
var (
	ErrNotFound = errors.New("not found")
)

// UserRepository defines the interface for user data access
type UserRepository interface {
	FindAll(ctx context.Context) ([]model.User, error)
	FindByID(ctx context.Context, id string) (model.User, error)
	Create(ctx context.Context, user model.User) (model.User, error)
	Update(ctx context.Context, user model.User) error
	Delete(ctx context.Context, id string) error
}
