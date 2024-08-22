package beans

import (
	"sync"

	"github.com/go-redis/redis"
)

type BloomFilter struct {
	RedisClient *redis.Client
	Key         string
	Size        uint64
	Hashes      int
}

type ThreadSafeBloomFilter struct {
    Filter *BloomFilter
    Mutex  sync.RWMutex
}