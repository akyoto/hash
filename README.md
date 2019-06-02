# hash

[![Godoc][godoc-image]][godoc-url]
[![Report][report-image]][report-url]
[![Tests][tests-image]][tests-url]
[![Coverage][coverage-image]][coverage-url]
[![Sponsor][sponsor-image]][sponsor-url]

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

## Collisions

The algorithm in its current state has a possible collision if equally sized payloads bigger than 64 bytes contain a +1/-1 diff on an 8-byte index difference.
If this turns out to be a problem for your use case, please consider other hash libraries or feel free to contribute and [send a PR](https://github.com/akyoto/hash/pulls).

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
