package operations

import (
	"redisbloom/lib/core/beans"
	"redisbloom/lib/core/crypto"
)

func Exists(bf *beans.BloomFilter, item []byte) (bool, error) {
	for i := 0; i < bf.Hashes; i++ {
		hash := crypto.ComputeHash(item, i)
		bit := hash % bf.Size
		exists, err := bf.RedisClient.GetBit(bf.Key, int64(bit)).Result()
		if err != nil {
			return false, err
		}
		if exists == 0 {
			return false, nil
		}
	}
	return true, nil
}

func ThreadSafeExists(tsbf *beans.ThreadSafeBloomFilter, item []byte) (bool, error) {
    tsbf.Mutex.RLock()
    defer tsbf.Mutex.RUnlock()
    return Exists(tsbf.Filter, item)
}
