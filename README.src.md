# {name}

{go:header}

Ultra-fast hashing of bytes, strings and reader contents. The algorithm is meant to be used for cache (in)validation. It is about 3 times faster than the popular `xxhash`.

## Usage

```go
x := hash.Bytes([]byte("Hello World"))
```

## Benchmarks

Hashing 10 KB:

```text
BenchmarkBytes-12       31416195                35.9 ns/op             0 B/op          0 allocs/op
BenchmarkString-12      31345899                36.3 ns/op             0 B/op          0 allocs/op
BenchmarkReader-12      19575999                59.9 ns/op            32 B/op          1 allocs/op
```

{go:footer}
