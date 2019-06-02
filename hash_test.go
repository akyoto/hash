package hash_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/akyoto/hash"
)

const (
	iterations = 1000000
	sizeMin    = 1
	sizeMax    = 128
)

func TestZeroedCollisions(t *testing.T) {
	hashes := map[uint64][]byte{}

	// Generate payloads
	for size := sizeMin; size <= sizeMax; size *= 2 {
		data := make([]byte, size)
		h := hash.Bytes(data)
		existingData, exists := hashes[h]

		if exists && !bytes.Equal(data, existingData) {
			t.Fatalf("Hash '%d' collision between zeroed sizes '%d' and '%d'", h, len(existingData), size)
		}

		tmp := make([]byte, len(data))
		copy(tmp, data)
		hashes[h] = tmp
	}
}

func TestSimilarCollisions(t *testing.T) {
	hashes := map[uint64][]byte{}
	offsetMax := 8

	// Generate payloads
	for size := sizeMin; size <= sizeMax; size *= 2 {
		for offset := 1; offset <= offsetMax; offset++ {
			data := make([]byte, size)
			index := 0

			for i := 0; i < iterations; i++ {
				data[index] += 1
				index = (index + offset) % len(data)
				h := hash.Bytes(data)
				existingData, exists := hashes[h]

				if exists && !bytes.Equal(data, existingData) {
					t.Logf("Len A: %d", len(data))

					for _, b := range data {
						t.Logf("A: %b (%d)", b, b)
					}

					t.Logf("Len B: %d", len(existingData))

					for _, b := range existingData {
						t.Logf("B: %b (%d)", b, b)
					}

					t.Fatalf("Hash %b (%d) already exists", h, h)
				}

				tmp := make([]byte, len(data))
				copy(tmp, data)
				hashes[h] = tmp
			}
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
	data := bytes.Repeat([]byte("HelloWorld"), 1000)

	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			hash.Bytes(data)
		}
	})
}

func BenchmarkString(b *testing.B) {
	data := strings.Repeat("HelloWorld", 1000)

	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			hash.String(data)
		}
	})
}
