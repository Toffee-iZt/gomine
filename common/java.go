package common

import (
	"crypto/sha256"
	"encoding/binary"
)

// JavaStringHashCode ...
func JavaStringHashCode(value string) int64 {
	var h int32

	if len(value) > 0 {
		for _, r := range value {
			h = 31*h + r
		}
	}

	return int64(h)
}

// JavaSHA256HashLong ...
func JavaSHA256HashLong(value int64) []byte {
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, uint64(value))

	hash := sha256.Sum256(bytes)

	return hash[:]
}
