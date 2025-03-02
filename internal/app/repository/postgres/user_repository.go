package postgres

import (
	"context"
	"errors"
	"time"

	"github.com/ThePotatoVerse/internal/app/model"
	"github.com/ThePotatoVerse/internal/app/repository"
	"github.com/ThePotatoVerse/pkg/database"
	"github.com/ThePotatoVerse/pkg/logger"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
)

// userRepository implements repository.UserRepository with PostgreSQL
type userRepository struct {
	db  *database.Postgres
	log logger.Logger
}

// NewUserRepository creates a new PostgreSQL user repository
func NewUserRepository(db *database.Postgres, log logger.Logger) repository.UserRepository {
	return &userRepository{
		db:  db,
		log: log,
	}
}

// FindAll returns all users
func (r *userRepository) FindAll(ctx context.Context) ([]model.User, error) {
	query := `
		SELECT id, name, email, created_at, updated_at
		FROM users
		ORDER BY created_at DESC
	`

	rows, err := r.db.Pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// FindByID returns a user by ID
func (r *userRepository) FindByID(ctx context.Context, id string) (model.User, error) {
	query := `
		SELECT id, name, email, created_at, updated_at
		FROM users
		WHERE id = $1
	`

	var user model.User
	err := r.db.Pool.QueryRow(ctx, query, id).Scan(
		&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.User{}, repository.ErrNotFound
		}
		return model.User{}, err
	}

	return user, nil
}

// Create creates a new user
func (r *userRepository) Create(ctx context.Context, user model.User) (model.User, error) {
	query := `
		INSERT INTO users (id, name, email, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, name, email, created_at, updated_at
	`

	// Generate ID if not provided
	if user.ID == "" {
		user.ID = uuid.New().String()
	}

	// Set timestamps
	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now

	err := r.db.Pool.QueryRow(
		ctx, query, user.ID, user.Name, user.Email, user.CreatedAt, user.UpdatedAt,
	).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

// Update updates a user
func (r *userRepository) Update(ctx context.Context, user model.User) error {
	query := `
		UPDATE users
		SET name = $1, email = $2, updated_at = $3
		WHERE id = $4
	`

	// Update timestamp
	user.UpdatedAt = time.Now()

	result, err := r.db.Pool.Exec(ctx, query, user.Name, user.Email, user.UpdatedAt, user.ID)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return repository.ErrNotFound
	}

	return nil
}

// Delete deletes a user
func (r *userRepository) Delete(ctx context.Context, id string) error {
	query := `
		DELETE FROM users
		WHERE id = $1
	`

	result, err := r.db.Pool.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return repository.ErrNotFound
	}

	return nil
}

// FindByEmail returns a user by email
func (r *userRepository) FindByEmail(ctx context.Context, email string) (model.User, error) {
	query := `
		SELECT id, name, email, created_at, updated_at
		FROM users
		WHERE email = $1
	`

	var user model.User
	err := r.db.Pool.QueryRow(ctx, query, email).Scan(
		&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.User{}, repository.ErrNotFound
		}
		return model.User{}, err
	}

	return user, nil
}

// ExistsByEmail checks if a user with the given email exists
func (r *userRepository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	query := `
		SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)
	`

	var exists bool
	err := r.db.Pool.QueryRow(ctx, query, email).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
