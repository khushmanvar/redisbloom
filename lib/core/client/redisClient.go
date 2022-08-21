package client

import "github.com/go-redis/redis"

func NewRedisClient(redisAddr string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})

	return client
}