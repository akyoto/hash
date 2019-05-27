package hash

import (
	"unsafe"

	utils "github.com/akyoto/stringutils/unsafe"
)

// Bytes hashes the given byte slice.
func Bytes(in []byte) uint64 {
	return add(0, in)
}

// String hashes the given string.
func String(in string) uint64 {
	return Bytes(utils.StringToBytes(in))
}

// add implements the actual hashing.
func add(x uint64, in []byte) uint64 {
	var i int

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
		x += 0xAA
		x = (x << 1) | (x >> (64 - 1))
	}

	// If we have at least 8 bytes left, convert them to uint64.
	for ; i < len(in)-7; i += 8 {
		x += *(*uint64)(unsafe.Pointer(&in[i]))
		x += 0x2A
		x = (x << 1) | (x >> (64 - 1))
	}

	// Hash the remaining bytes.
	// At this point we know that there are less than 8 bytes left,
	// so we can shift each byte by 8 bits to assure that small
	// data hashes are always unique.
	// The "0xA" helps to avoid clashes between different lengths
	// of all-zero bytes by making the data length significant.
	for ; i < len(in); i++ {
		x += uint64(in[i])
		x += 0xA
		x = (x << 8) | (x >> (64 - 8))
	}

	return x
}
