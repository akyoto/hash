package hash_test

import (
	"bytes"
	"crypto/rand"
	"strings"
	"sync"
	"testing"

	"github.com/akyoto/hash"
)

const (
	iterations = 1000000
	sizeMin    = 1
	sizeMax    = 128
)

func TestRandomCollisions(t *testing.T) {
	hashes := map[uint64][]byte{}

	// Generate payloads
	for size := sizeMin; size <= sizeMax; size *= 2 {
		for i := 0; i < iterations; i++ {
			data := make([]byte, size)
			_, _ = rand.Read(data)
			h := hash.Bytes(data)
			existingData, exists := hashes[h]

			if exists && !bytes.Equal(data, existingData) {
				t.Fatalf("Hash '%d' already exists", h)
			}

			hashes[h] = data
		}
	}
}

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

		hashes[h] = data
	}
}

func TestSimilarCollisions(t *testing.T) {
	hashes := map[uint64][]byte{}

	// Generate payloads
	for size := sizeMin; size <= sizeMax; size *= 2 {
		data := make([]byte, size)
		index := 0

		for i := 0; i < iterations; i++ {
			data[index] += 1
			index = (index + 1) % len(data)
			h := hash.Bytes(data)
			existingData, exists := hashes[h]

			if exists && !bytes.Equal(data, existingData) {
				t.Fatalf("Hash '%d' already exists", h)
			}

			hashes[h] = data
		}
	}
}

func TestString(t *testing.T) {
	x := hash.String("Hello World")

	if x == 0 {
		t.Fatal("String hashing is most likely broken")
	}
}

func TestReader(t *testing.T) {
	a := hash.String("Hello World")
	b := hash.Reader(strings.NewReader("Hello World"))

	if a != b {
		t.Fatal("Reader hashing is most likely broken")
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

func BenchmarkReader(b *testing.B) {
	data := bytes.Repeat([]byte("HelloWorld"), 1000)

	pool := sync.Pool{
		New: func() interface{} {
			return bytes.NewReader(data)
		},
	}

	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		reader := pool.Get().(*bytes.Reader)

		for pb.Next() {
			_, _ = reader.Seek(0, 0)
			hash.Reader(reader)
		}

		pool.Put(reader)
	})
}
