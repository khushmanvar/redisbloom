package crypto

import (
	"encoding/binary"
)

const prime64 uint64 = 0x9e3779b97f4a7c15

func ComputeHash(item []byte, seed int) uint64 {
	hash := uint64(seed) + prime64
	hash = mixHash(item, hash)
	hash = finalizeHash(hash)
	return hash
}

func mixHash(item []byte, hash uint64) uint64 {
	for len(item) >= 8 {
		k := binary.LittleEndian.Uint64(item)
		k *= prime64
		k ^= k >> 33
		k *= prime64
		hash ^= k
		hash *= prime64
		item = item[8:]
	}
	if len(item) > 0 {
		var k uint64
		for i := len(item) - 1; i >= 0; i-- {
			k <<= 8
			k |= uint64(item[i])
		}
		hash ^= k
		hash *= prime64
	}
	return hash
}

func finalizeHash(hash uint64) uint64 {
	hash ^= hash >> 33
	hash *= prime64
	hash ^= hash >> 29
	hash *= prime64
	hash ^= hash >> 32
	return hash
}
