# {name}

{go:header}

Ultra-fast hashing of arbitrary bytes.

## Usage

```go
x := hash.Bytes([]byte("Hello World"))
```

## Benchmarks

```text
BenchmarkBytes-12       26772571                42.2 ns/op             0 B/op          0 allocs/op
BenchmarkString-12      27747283                42.3 ns/op             0 B/op          0 allocs/op
```

{go:footer}
