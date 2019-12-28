
# wyhash
[![GoDoc](https://godoc.org/github.com/zeebo/wyhash?status.svg)](https://godoc.org/github.com/zeebo/wyhash)
[![Sourcegraph](https://sourcegraph.com/github.com/zeebo/wyhash/-/badge.svg)](https://sourcegraph.com/github.com/zeebo/wyhash?badge)
[![Go Report Card](https://goreportcard.com/badge/github.com/zeebo/wyhash)](https://goreportcard.com/report/github.com/zeebo/wyhash)

This package is a port of the [wyhash](https://github.com/wangyi-fudan/wyhash) library v3 to Go.

---

# Hash Benchmarks

Run on my `i7-6700K CPU @ 4.00GHz`

| Bytes     | Rate                                 |
|-----------|--------------------------------------|
|` 0 `      |` 0.49 ns/op `                        |
|` 1-3 `    |` 3.46 ns/op (0.29 GB/s - 0.88 GB/s) `|
|` 4-8 `    |` 3.40 ns/op (1.17 GB/s - 2.38 GB/s) `|
|` 9-16 `   |` 3.64 ns/op (2.47 GB/s - 4.39 GB/s) `|
|` 17-32 `  |` 4.63 ns/op (3.68 GB/s - 6.11 GB/s) `|
|` 33-64 `  |` 7.51 ns/op (4.40 GB/s - 7.98 GB/s) `|
|` 65-96 `  |` 9.22 ns/op (7.05 GB/s - 9.98 GB/s) `|
|` 97-128 ` |` 10.7 ns/op (9.03 GB/s - 11.3 GB/s) `|
|` 240 `    |` 17.5 ns/op (13.7 GB/s) `            |
|` 512 `    |` 37.3 ns/op (13.7 GB/s) `            |
|` 1024 `   |` 69.0 ns/op (14.8 GB/s) `            |
|` 100KB `  |` 6059 ns/op (16.9 GB/s) `            |

# RNG Benchmarks

| Method            | Speed         |
|-------------------|---------------|
|` Uint64 `         |` 1.31 ns/op ` |
|` Uint64n(large) ` |` 12.7 ns/op ` |
|` Uint64n(med) `   |` 2.40 ns/op ` |
|` Uint64n(small) ` |` 2.43 ns/op ` |
|` Float64 `        |` 2.17 ns/op ` |

# Usage

#### func  Float64

```go
func Float64() float64
```
Float64 returns a float64 uniformly in [0, 1). Safe for concurrent callers.

#### func  Hash

```go
func Hash(data []byte, seed uint64) uint64
```
Hash returns a 64bit digest of the data with different ones for every seed.

#### func  HashString

```go
func HashString(data string, seed uint64) uint64
```
HashString returns a 64bit digest of the data with different ones for every
seed.

#### func  Int

```go
func Int() int
```
Int returns a random int. Safe for concurrent callers.

#### func  Intn

```go
func Intn(n int) int
```
Intn returns a int uniformly in [0, n). Safe for concurrent callers.

#### func  Uint64

```go
func Uint64() uint64
```
Uint64 returns a random uint64. Safe for concurrent callers.

#### func  Uint64n

```go
func Uint64n(n uint64) uint64
```
Uint64n returns a uint64 uniformly in [0, n). Safe for concurrent callers.

#### type RNG

```go
type RNG uint64
```

RNG is a random number generator. The zero value is valid.

#### func (*RNG) Float64

```go
func (r *RNG) Float64() (v float64)
```
Float64 returns a float64 uniformly in [0, 1). Not safe for concurrent callers.

#### func (*RNG) Int

```go
func (r *RNG) Int() int
```
Int returns a random int. Not safe for concurrent callers.

#### func (*RNG) Intn

```go
func (r *RNG) Intn(n int) int
```
Intn returns an int uniformly in [0, n). Not safe for concurrent callers.

#### func (*RNG) Uint64

```go
func (r *RNG) Uint64() uint64
```
Uint64 returns a random uint64. Not safe for concurrent callers.

#### func (*RNG) Uint64n

```go
func (r *RNG) Uint64n(n uint64) uint64
```
Uint64n returns a uint64 uniformly in [0, n). Not safe for concurrent callers.
