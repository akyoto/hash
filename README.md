# hash

[![Godoc][godoc-image]][godoc-url]
[![Report][report-image]][report-url]
[![Tests][tests-image]][tests-url]
[![Coverage][coverage-image]][coverage-url]
[![Sponsor][sponsor-image]][sponsor-url]

Ultra-fast hashing of bytes, strings and reader contents for non-cryptographic purposes such as cache validation.

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

## Style

Please take a look at the [style guidelines](https://github.com/akyoto/quality/blob/master/STYLE.md) if you'd like to make a pull request.

## Sponsors

| [![Cedric Fung](https://avatars3.githubusercontent.com/u/2269238?s=70&v=4)](https://github.com/cedricfung) | [![Scott Rayapoullé](https://avatars3.githubusercontent.com/u/11772084?s=70&v=4)](https://github.com/soulcramer) | [![Eduard Urbach](https://avatars3.githubusercontent.com/u/438936?s=70&v=4)](https://twitter.com/eduardurbach) |
| --- | --- | --- |
| [Cedric Fung](https://github.com/cedricfung) | [Scott Rayapoullé](https://github.com/soulcramer) | [Eduard Urbach](https://eduardurbach.com) |

Want to see [your own name here?](https://github.com/users/akyoto/sponsorship)

[godoc-image]: https://godoc.org/github.com/akyoto/hash?status.svg
[godoc-url]: https://godoc.org/github.com/akyoto/hash
[report-image]: https://goreportcard.com/badge/github.com/akyoto/hash
[report-url]: https://goreportcard.com/report/github.com/akyoto/hash
[tests-image]: https://cloud.drone.io/api/badges/akyoto/hash/status.svg
[tests-url]: https://cloud.drone.io/akyoto/hash
[coverage-image]: https://codecov.io/gh/akyoto/hash/graph/badge.svg
[coverage-url]: https://codecov.io/gh/akyoto/hash
[sponsor-image]: https://img.shields.io/badge/github-donate-green.svg
[sponsor-url]: https://github.com/users/akyoto/sponsorship
