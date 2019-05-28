# {name}

{go:header}

Ultra-fast hashing of bytes, strings and reader contents. The algorithm is meant to be used for cache (in)validation. It is about 3 times faster than the popular `xxhash`.

## Comparison

```text
BenchmarkAkyotoHash_4B-12       1000000000               0.702 ns/op           0 B/op          0 allocs/op
BenchmarkCRC32_4B-12            748350471                1.59 ns/op            0 B/op          0 allocs/op
BenchmarkXXHash_4B-12           533528487                2.22 ns/op            0 B/op          0 allocs/op
BenchmarkSipHash_4B-12          833339200                1.43 ns/op            0 B/op          0 allocs/op
BenchmarkFNV_4B-12              1000000000               0.703 ns/op           0 B/op          0 allocs/op

BenchmarkAkyotoHash_8B-12       1000000000               0.469 ns/op           0 B/op          0 allocs/op
BenchmarkCRC32_8B-12            547803841                2.17 ns/op            0 B/op          0 allocs/op
BenchmarkXXHash_8B-12           516036067                2.24 ns/op            0 B/op          0 allocs/op
BenchmarkSipHash_8B-12          682836061                1.74 ns/op            0 B/op          0 allocs/op
BenchmarkFNV_8B-12              1000000000               1.01 ns/op            0 B/op          0 allocs/op

BenchmarkAkyotoHash_10KB-12     30067696                37.0 ns/op             0 B/op          0 allocs/op
BenchmarkCRC32_10KB-12          22916304                50.3 ns/op             0 B/op          0 allocs/op
BenchmarkXXHash_10KB-12         11700764               101 ns/op               0 B/op          0 allocs/op
BenchmarkSipHash_10KB-12         2815754               423 ns/op               0 B/op          0 allocs/op
BenchmarkFNV_10KB-12             1534830               781 ns/op               0 B/op          0 allocs/op
```

Run these for yourself using [hash-benchmarks](https://github.com/akyoto/hash-benchmarks).

## Usage

### []byte

```go
x := hash.Bytes([]byte("Hello World"))
```

### string

```go
x := hash.String("Hello World")
```

### io.Reader

```go
x := hash.Reader(file)
```

{go:footer}
