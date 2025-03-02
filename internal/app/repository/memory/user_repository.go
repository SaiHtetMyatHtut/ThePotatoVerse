package memory

import (
	"context"
	"sync"
	"time"

	"github.com/ThePotatoVerse/internal/app/model"
	"github.com/ThePotatoVerse/internal/app/repository"
	"github.com/google/uuid"
)

// userRepository implements repository.UserRepository with an in-memory store
type userRepository struct {
	mu    sync.RWMutex
	users map[string]model.User
}

// NewUserRepository creates a new in-memory user repository
func NewUserRepository() repository.UserRepository {
	return &userRepository{
		users: make(map[string]model.User),
	}
}

// FindAll returns all users
func (r *userRepository) FindAll(ctx context.Context) ([]model.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	users := make([]model.User, 0, len(r.users))
	for _, user := range r.users {
		users = append(users, user)
	}

	return users, nil
}

// FindByID returns a user by ID
func (r *userRepository) FindByID(ctx context.Context, id string) (model.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, ok := r.users[id]
	if !ok {
		return model.User{}, repository.ErrNotFound
	}

	return user, nil
}

// Create creates a new user
func (r *userRepository) Create(ctx context.Context, user model.User) (model.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Generate ID if not provided
	if user.ID == "" {
		user.ID = uuid.New().String()
	}

	// Set timestamps
	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now

	// Store user
	r.users[user.ID] = user

	return user, nil
}

// Update updates a user
func (r *userRepository) Update(ctx context.Context, user model.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Check if user exists
	_, ok := r.users[user.ID]
	if !ok {
		return repository.ErrNotFound
	}

	// Update timestamp
	user.UpdatedAt = time.Now()

	// Store updated user
	r.users[user.ID] = user

	return nil
}

// Delete deletes a user
func (r *userRepository) Delete(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Check if user exists
	_, ok := r.users[id]
	if !ok {
		return repository.ErrNotFound
	}

	// Delete user
	delete(r.users, id)

	return nil
}

// FindByEmail returns a user by email
func (r *userRepository) FindByEmail(ctx context.Context, email string) (model.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, user := range r.users {
		if user.Email == email {
			return user, nil
		}
	}

	return model.User{}, repository.ErrNotFound
}

// ExistsByEmail checks if a user with the given email exists
func (r *userRepository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, user := range r.users {
		if user.Email == email {
			return true, nil
		}
	}

	return false, nil
}
