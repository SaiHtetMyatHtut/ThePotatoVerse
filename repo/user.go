package repo

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/SaiHtetMyatHtut/potatoverse/db"
	"github.com/SaiHtetMyatHtut/potatoverse/model"
)

// userKey generates a key for a user with the given ID.
func userKey(id int) string {
	return fmt.Sprintf("User with ID %d", id)
}

// Insert inserts a user into the Redis repository.
func Insert(ctx context.Context, user model.User) error {
	db := db.NewRedisRepo()
	// Marshal user data into JSON format.
	userData, err := json.Marshal(user)
	if err != nil {
		return fmt.Errorf("> Error Marshalling User: %w", err)
	}

	// userKey generates a unique key for the given user ID.
	key := userKey(user.ID)

	// txn is a pipeline for executing multiple Redis commands in a transaction.
	txn := db.Client.TxPipeline()

	// SetNX sets the value of the key in the transaction if it does not already exist.
	// It returns true if the key was set successfully, false otherwise.
	res := txn.SetNX(ctx, key, string(userData), 0)
	if res.Err() != nil {
		txn.Discard()
		db.Close()
		return fmt.Errorf("> Error Setting User: %w", res.Err())
	}

	// AddUserToSet adds a user to a set in Redis.
	// It takes a Redis transaction (txn), a context (ctx), and a key as parameters.
	// If the operation fails, it returns an error.
	if err := txn.SAdd(ctx, "users", key).Err(); err != nil {
		txn.Discard()
		db.Close()
		return fmt.Errorf("> Error Adding User to Set: %w", err)
	}

	// Exec executes the transaction and returns any error encountered.
	if _, err := txn.Exec(ctx); err != nil {
		db.Close()
		return fmt.Errorf("> Error Executing Transaction: %w", err)
	}
	db.Close()
	return nil
}

func ReadAll(ctx context.Context) ([]model.User, error) {
	db := db.NewRedisRepo()
	var users []model.User

	// SMembers returns all the members of the set in the Redis repository.
	// If the set does not exist, it returns an empty slice.
	keys, err := db.Client.SMembers(ctx, "users").Result()
	if err != nil {
		db.Close()
		return users, fmt.Errorf("> Error Getting Users: %w", err)
	}

	// Iterate over the keys and retrieve the user data from the Redis repository.
	for _, key := range keys {
		val, err := db.Client.Get(ctx, key).Result()
		if err != nil {
			db.Close()
			return users, fmt.Errorf("> Error Getting User: %w", err)
		}

		var user model.User
		if err := json.Unmarshal([]byte(val), &user); err != nil {
			db.Close()
			return users, fmt.Errorf("> Error Unmarshalling User: %w", err)
		}

		users = append(users, user)
	}
	db.Close()
	return users, nil
}

func ReadByID(ctx context.Context, id int) (model.User, error) {
	db := db.NewRedisRepo()
	var user model.User

	// userKey generates a unique key for the given user ID.
	key := userKey(id)

	// Get retrieves the value of the key from the Redis repository.
	// If the key does not exist, it returns a nil value.
	val, err := db.Client.Get(ctx, key).Result()
	if err != nil {
		db.Close()
		return user, fmt.Errorf("> Error Getting User: %w", err)
	}

	// Unmarshal user data from JSON format.
	if err := json.Unmarshal([]byte(val), &user); err != nil {
		db.Close()
		return user, fmt.Errorf("> Error Unmarshalling User: %w", err)
	}
	db.Close()
	return user, nil
}

func Update(ctx context.Context, user model.User) error {
	db := db.NewRedisRepo()
	// Marshal user data into JSON format.
	userData, err := json.Marshal(user)
	if err != nil {
		return fmt.Errorf("> Error Marshalling User: %w", err)
	}

	// userKey generates a unique key for the given user ID.
	key := userKey(user.ID)

	// Set updates the value of the key in the Redis repository.
	// It returns an error if the operation fails.
	if err := db.Client.Set(ctx, key, string(userData), 0).Err(); err != nil {
		return fmt.Errorf("> Error Setting User: %w", err)
	}
	db.Close()
	return nil
}

func Delete(ctx context.Context, id int) error {
	db := db.NewRedisRepo()
	// userKey generates a unique key for the given user ID.
	key := userKey(id)

	// Del removes the key from the Redis repository.
	// It returns the number of keys that were removed.
	if err := db.Client.Del(ctx, key).Err(); err != nil {
		db.Close()
		return fmt.Errorf("> Error Deleting User: %w", err)
	}

	// SRem removes the specified members from the set stored at the key.
	// It returns the number of members that were removed from the set.
	if err := db.Client.SRem(ctx, "users", key).Err(); err != nil {
		db.Close()
		return fmt.Errorf("> Error Removing User from Set: %w", err)
	}
	db.Close()
	return nil
}
