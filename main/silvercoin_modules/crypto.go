package silvercoin_modules

import (
	// https://codesigningstore.com/hash-algorithm-comparison
	"crypto/sha256"
)

func Hash(data []byte) []byte {
	hash := sha256.Sum224(data)
	return hash[:]
}
