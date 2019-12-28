package wyhash

import "sync"

// global is a parallel rng for the package functions.
var global struct {
	sync.Mutex
	RNG
}

// Int returns a random int.
// Safe for concurrent callers.
func Int() int {
	global.Lock()
	out := global.Int()
	global.Unlock()
	return out
}

// Intn returns a int uniformly in [0, n).
// Safe for concurrent callers.
func Intn(n int) int {
	global.Lock()
	out := global.Intn(n)
	global.Unlock()
	return out
}

// Uint64 returns a random uint64.
// Safe for concurrent callers.
func Uint64() uint64 {
	global.Lock()
	out := global.Uint64()
	global.Unlock()
	return out
}

// Uint64n returns a uint64 uniformly in [0, n).
// Safe for concurrent callers.
func Uint64n(n uint64) uint64 {
	global.Lock()
	out := global.Uint64n(n)
	global.Unlock()
	return out
}

// Float64 returns a float64 uniformly in [0, 1).
// Safe for concurrent callers.
func Float64() float64 {
	global.Lock()
	out := global.Float64()
	global.Unlock()
	return out
}
