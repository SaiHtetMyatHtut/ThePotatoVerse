package repositories

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/SaiHtetMyatHtut/potatoverse/src/core/data/models"
	"github.com/redis/go-redis/v9"
	"go.uber.org/dig"
)

func userKey(id int64) string {
	return fmt.Sprintf("users:%d", id)
}

type UserRepository struct {
	client *redis.Client
}

type UserRepositoryDependencies struct {
	dig.In

	Client *redis.Client `name:"Redis"`
}

func NewUserRepository(client *redis.Client) *UserRepository {
	return &UserRepository{client: client}
}

func NewNewUserRepository(deps UserRepositoryDependencies) UserRepository {
	return UserRepository{client: deps.Client}
}

// Insert inserts a user into the Redis repository.
func (r UserRepository) Insert(ctx context.Context, user models.User) (models.User, error) {
	// Increment the user ID.
	id, err := r.client.Incr(ctx, "user:id").Result()
	if err != nil {
		// r.client.Close()
		return user, fmt.Errorf("> Error Incrementing User ID: %w", err)
	}
	user.ID = id
	key := userKey(id)

	// Marshal user data into JSON format.
	userData, err := json.Marshal(user)
	if err != nil {
		return user, fmt.Errorf("> Error Marshalling User: %w", err)
	}

	// txn is a pipeline for executing multiple Redis commands in a transaction.
	txn := r.client.TxPipeline()

	// SetNX sets the value of the key in the transaction if it does not already exist.
	// It returns true if the key was set successfully, false otherwise.
	res := txn.SetNX(ctx, key, string(userData), 0)
	if res.Err() != nil {
		txn.Discard()
		// r.client.Close()
		return user, fmt.Errorf("> Error Setting User: %w", res.Err())
	}

	// AddUserToSet adds a user to a set in Redis.
	// It takes a Redis transaction (txn), a context (ctx), and a key as parameters.
	// If the operation fails, it returns an error.
	if err := txn.SAdd(ctx, "users", key).Err(); err != nil {
		txn.Discard()
		// r.client.Close()
		return user, fmt.Errorf("> Error Adding User to Set: %w", err)
	}

	// Exec executes the transaction and returns any error encountered.
	if _, err := txn.Exec(ctx); err != nil {
		// r.client.Close()
		return user, fmt.Errorf("> Error Executing Transaction: %w", err)
	}
	// r.client.Close()
	return user, nil
}

func (r UserRepository) ReadAll(ctx context.Context) ([]models.User, error) {
	var users []models.User

	// SMembers returns all the members of the set in the Redis repository.
	// If the set does not exist, it returns an empty slice.
	keys, err := r.client.SMembers(ctx, "users").Result()
	if err != nil {
		// r.client.Close()
		return users, fmt.Errorf("> Error Getting Users: %w", err)
	}

	// Iterate over the keys and retrieve the user data from the Redis repository.
	for _, key := range keys {
		val, err := r.client.Get(ctx, key).Result()
		if err != nil {
			// r.client.Close()
			return users, fmt.Errorf("> Error Getting User: %w", err)
		}

		var user models.User
		if err := json.Unmarshal([]byte(val), &user); err != nil {
			// r.client.Close()
			return users, fmt.Errorf("> Error Unmarshalling User: %w", err)
		}

		users = append(users, user)
	}
	// r.client.Close()
	return users, nil
}

func (r UserRepository) ReadByID(ctx context.Context, id int64) (models.User, error) {
	var user models.User

	// userKey generates a unique key for the given user ID.
	key := userKey(id)

	// Get retrieves the value of the key from the Redis repository.
	// If the key does not exist, it returns a nil value.
	val, err := r.client.Get(ctx, key).Result()
	if err != nil {
		// r.client.Close()
		if err == redis.Nil {
			return user, fmt.Errorf("> User Not Found")
		} else {
			return user, fmt.Errorf("> Error Getting User: %w", err)
		}
	}

	// Unmarshal user data from JSON format.
	if err := json.Unmarshal([]byte(val), &user); err != nil {
		// r.client.Close()
		return user, fmt.Errorf("> Error Unmarshalling User: %w", err)
	}
	// r.client.Close()
	return user, nil
}

func (r UserRepository) Update(ctx context.Context, user models.User) error {
	// Marshal user data into JSON format.
	userData, err := json.Marshal(user)
	if err != nil {
		return fmt.Errorf("> Error Marshalling User: %w", err)
	}

	// userKey generates a unique key for the given user ID.
	key := userKey(user.ID)

	// Set updates the value of the key in the Redis repository.
	// It returns an error if the operation fails.
	if err := r.client.Set(ctx, key, string(userData), 0).Err(); err != nil {
		return fmt.Errorf("> Error Setting User: %w", err)
	}
	// r.client.Close()
	return nil
}

func (r UserRepository) Delete(ctx context.Context, id int64) error {
	// userKey generates a unique key for the given user ID.
	key := userKey(id)

	// Del removes the key from the Redis repository.
	// It returns the number of keys that were removed.
	if err := r.client.Del(ctx, key).Err(); err != nil {
		// r.client.Close()
		return fmt.Errorf("> Error Deleting User: %w", err)
	}

	// SRem removes the specified members from the set stored at the key.
	// It returns the number of members that were removed from the set.
	if err := r.client.SRem(ctx, "users", key).Err(); err != nil {
		// r.client.Close()
		return fmt.Errorf("> Error Removing User from Set: %w", err)
	}
	// r.client.Close()
	return nil
}
