package db

import (
	"github.com/SaiHtetMyatHtut/potatoverse/configs"
	"github.com/redis/go-redis/v9"
)

type RedisRepo struct {
	Client *redis.Client
}

func NewRedisRepo() *RedisRepo {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     configs.Env.Redis.Host + configs.Env.Redis.Port,
		Password: "",
		DB:       0,
	})
	return &RedisRepo{
		Client: redisClient,
	}
}

func NewRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     configs.Env.Redis.Host + configs.Env.Redis.Port,
		Password: "",
		DB:       0,
	})
}

func (r *RedisRepo) Close() error {
	return r.Client.Close()
}
