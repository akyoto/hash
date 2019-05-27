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
	iterations = 100000
	sizeMin    = 32
	sizeMax    = 8192
)

func TestCollisions(t *testing.T) {
	hashes := map[uint64]bool{}

	// Generate payloads
	for size := sizeMin; size < sizeMax; size *= 2 {
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
