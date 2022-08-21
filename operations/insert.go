package operations

import (
	"redisbloom/lib/core/beans"
	"redisbloom/lib/core/crypto"
)

func Insert(bf *beans.BloomFilter, item []byte) error {
	for i := 0; i < bf.Hashes; i++ {
		hash := crypto.ComputeHash(item, i)
		bit := hash % bf.Size
		if err := bf.RedisClient.SetBit(bf.Key, int64(bit), 1).Err(); err != nil {
			return err
		}
	}
	return nil
}