package redisbloom

import (
	"redisbloom/lib/core/beans"

 	"github.com/go-redis/redis"
)

func NewBloomFilter(redisAddr, key string, size uint64, hashes int) *beans.BloomFilter {
	client := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})
	return &beans.BloomFilter{
		RedisClient: client,
		Key:         key,
		Size:        size,
		Hashes:      hashes,
	}
}

func NewThreadSafeBloomFilter(redisAddr, key string, size uint64, hashes int) *beans.ThreadSafeBloomFilter {
    return &beans.ThreadSafeBloomFilter{
        Filter: NewBloomFilter(redisAddr, key, size, hashes),
    }
}