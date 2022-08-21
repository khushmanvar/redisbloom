package beans

import "github.com/go-redis/redis"

type BloomFilter struct {
	RedisClient *redis.Client
	Key         string
	Size        uint64
	Hashes      int
}