# {name}

{go:header}

Ultra-fast hashing of arbitrary bytes.

## Usage

```go
x := hash.Bytes([]byte("Hello World"))
```

## Benchmarks

Hashing 10 KB:

```text
BenchmarkBytes-12       30485769                37.9 ns/op             0 B/op          0 allocs/op
BenchmarkString-12      29912976                38.3 ns/op             0 B/op          0 allocs/op
BenchmarkReader-12      19175532                61.3 ns/op            32 B/op          1 allocs/op
```

{go:footer}
