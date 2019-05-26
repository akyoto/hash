package hash_test

import (
	"bytes"
	"crypto/rand"
	"strings"
	"testing"

	"github.com/akyoto/hash"
)

const (
	iterations = 100
	sizeMin    = 8
	sizeMax    = 8 * 2048
	sizeStep   = 8
)

func TestCollisions(t *testing.T) {
	hashes := map[uint64]bool{}

	// Generate payloads
	for size := sizeMin; size < sizeMax; size += sizeStep {
		for i := 0; i < iterations; i++ {
			data := make([]byte, size)
			_, _ = rand.Read(data)
			h := hash.Bytes(data)
			_, exists := hashes[h]

			if exists {
				t.Fatalf("Hash '%d' for '%s' already exists", h, data)
			}

			hashes[h] = true
		}
	}
}

func TestString(t *testing.T) {
	x := hash.String("Hello World")

	if x == 0 {
		t.Fatal("String hashing is most likely broken")
	}
}

func BenchmarkBytes(b *testing.B) {
	data := bytes.Repeat([]byte("Hello World"), 1000)

	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			hash.Bytes(data)
		}
	})
}

func BenchmarkString(b *testing.B) {
	data := strings.Repeat("Hello World", 1000)

	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			hash.String(data)
		}
	})
}
