package db

import "github.com/redis/go-redis/v9"

type RedisRepo struct {
	Client *redis.Client
}

func NewRedisRepo() *RedisRepo {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return &RedisRepo{
		Client: redisClient,
	}
}

func (r *RedisRepo) Close() error {
	return r.Client.Close()
}
