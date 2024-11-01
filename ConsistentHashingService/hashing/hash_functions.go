package hashing

import "crypto/sha1"

func HashFunction(key string) []byte {
	hasher := sha1.New()
	hasher.Write([]byte(key))
	hash_value := hasher.Sum(nil)
	return hash_value
}
