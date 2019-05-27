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
BenchmarkBytes-12       31116072                36.9 ns/op             0 B/op          0 allocs/op
BenchmarkString-12      30273320                37.2 ns/op             0 B/op          0 allocs/op
BenchmarkReader-12      18604939                61.9 ns/op            32 B/op          1 allocs/op
```

{go:footer}
