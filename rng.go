package wyhash

import "math/bits"

// RNG is a random number generator.
// The zero value is valid.
type RNG uint64

// Int returns a random positive int.
// Not safe for concurrent callers.
func (r *RNG) Int() int {
	return int(uint(r.Uint64()) >> 1)
}

// Intn returns an int uniformly in [0, n).
// Not safe for concurrent callers.
func (r *RNG) Intn(n int) int {
	if n <= 0 {
		return 0
	}
	return int(r.Uint64n(uint64(n)))
}

// Uint64 returns a random uint64.
// Not safe for concurrent callers.
func (r *RNG) Uint64() uint64 {
	*r += _wyp0
	return _wymum(u64(*r)^_wyp1, u64(*r))
}

// Uint64n returns a uint64 uniformly in [0, n).
// Not safe for concurrent callers.
func (r *RNG) Uint64n(n uint64) uint64 {
	if n == 0 {
		return 0
	}

	x := r.Uint64()
	h, l := bits.Mul64(x, n)

	if l < n {
		t := -n
		if t >= n {
			t -= n
			if t >= n {
				t = t % n
			}
		}

	again:
		if l < t {
			x = r.Uint64()
			h, l = bits.Mul64(x, n)
			goto again
		}
	}

	return h
}

// Float64 returns a float64 uniformly in [0, 1).
// Not safe for concurrent callers.
func (r *RNG) Float64() (v float64) {
again:
	v = float64(r.Uint64()>>11) / (1 << 53)
	if v == 1 {
		goto again
	}
	return
}
