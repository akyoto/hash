# {name}

{go:header}

Ultra-fast hashing of bytes, strings and reader contents. The algorithm is meant to be used for cache (in)validation.

## Benchmarks

[![Hash performance benchmarks](docs/hash-performance.png)](https://github.com/akyoto/hash-benchmarks)

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
