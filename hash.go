package hash

import (
	"unsafe"

	utils "github.com/akyoto/stringutils/unsafe"
)

// Bytes hashes the given byte slice.
func Bytes(in []byte) uint64 {
	x := uint64(len(in))
	i := 0

	// Cache lines on modern processors are 64 bytes long.
	// A single uint64 consumes 64 bits (8 bytes).
	// That means we should read 8 integers at a time.
	for ; i < len(in)-63; i += 64 {
		words := (*[8]uint64)(unsafe.Pointer(&in[i]))
		x += words[0]
		x += words[1]
		x += words[2]
		x += words[3]
		x += words[4]
		x += words[5]
		x += words[6]
		x += words[7]
		x = (x << 1) | (x >> (64 - 1))
	}

	// Remaining bytes
	for ; i < len(in); i++ {
		x += uint64(in[i])
		x = (x << 1) | (x >> (64 - 1))
	}

	return x
}

// String hashes the given string.
func String(in string) uint64 {
	return Bytes(utils.StringToBytes(in))
}
