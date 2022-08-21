package redisbloom

import (
	"github.com/go-redis/redis/v8"
)

type BloomFilter struct {
	redisClient *redis.Client
	key         string
	size        uint64
	hashes      int
}

func NewBloomFilter(redisAddr, key string, size uint64, hashes int) *BloomFilter {
	client := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})
	return &BloomFilter{
		redisClient: client,
		key:         key,
		size:        size,
		hashes:      hashes,
	}
}