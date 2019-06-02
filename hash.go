package hash

import (
	"github.com/zeebo/xxh3"
)

// Bytes hashes the given byte slice.
func Bytes(in []byte) uint64 {
	return xxh3.Hash(in)
}

// String hashes the given string.
func String(in string) uint64 {
	return xxh3.HashString(in)
}
