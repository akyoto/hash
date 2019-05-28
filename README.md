# hash

[![Godoc][godoc-image]][godoc-url]
[![Report][report-image]][report-url]
[![Tests][tests-image]][tests-url]
[![Coverage][coverage-image]][coverage-url]
[![Sponsor][sponsor-image]][sponsor-url]

Ultra-fast hashing of bytes, strings and reader contents. The algorithm is meant to be used for cache (in)validation. It is about 3 times faster than the popular `xxhash`.

## Comparison

```text
BenchmarkAkyotoHash_4B-12       1000000000               0.705 ns/op           0 B/op          0 allocs/op
BenchmarkCRC32_4B-12            732247504                1.60 ns/op            0 B/op          0 allocs/op
BenchmarkXXHash_4B-12           697969192                1.70 ns/op            0 B/op          0 allocs/op
BenchmarkSipHash_4B-12          831329154                1.43 ns/op            0 B/op          0 allocs/op

BenchmarkAkyotoHash_8B-12       1000000000               0.520 ns/op           0 B/op          0 allocs/op
BenchmarkCRC32_8B-12            548619034                2.17 ns/op            0 B/op          0 allocs/op
BenchmarkXXHash_8B-12           670126293                1.77 ns/op            0 B/op          0 allocs/op
BenchmarkSipHash_8B-12          680157788                1.75 ns/op            0 B/op          0 allocs/op

BenchmarkAkyotoHash_10KB-12     30696001                37.0 ns/op             0 B/op          0 allocs/op
BenchmarkCRC32_10KB-12          22894020                50.3 ns/op             0 B/op          0 allocs/op
BenchmarkXXHash_10KB-12         11611434               101 ns/op               0 B/op          0 allocs/op
BenchmarkSipHash_10KB-12         2812972               424 ns/op               0 B/op          0 allocs/op
```

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
