package wyhash

import (
	"runtime"
	"testing"
)

func TestRng(t *testing.T) {
	var rng RNG
	for i, want := range rngvecs {
		got := rng.Uint64()
		if got != want {
			t.Errorf("rng.Next()[%d]=%x, want %x", i, got, want)
		}
	}
}

func BenchmarkRNG(b *testing.B) {
	var blackholeUint64 uint64
	var blackholeFloat64 float64

	b.Run("Uint64", func(b *testing.B) {
		rng := RNG(2345)
		for i := 0; i < b.N; i++ {
			blackholeUint64 += rng.Uint64()
		}
	})

	b.Run("Uint64n", func(b *testing.B) {
		b.Run("Large", func(b *testing.B) {
			rng := RNG(2345)
			for i := 0; i < b.N; i++ {
				blackholeUint64 += rng.Uint64n(1<<63 + 1)
			}
		})

		b.Run("Med", func(b *testing.B) {
			rng := RNG(2345)
			for i := 0; i < b.N; i++ {
				blackholeUint64 += rng.Uint64n(1<<31 + 1)
			}
		})

		b.Run("Small", func(b *testing.B) {
			rng := RNG(2345)
			for i := 0; i < b.N; i++ {
				blackholeUint64 += rng.Uint64n(1000)
			}
		})
	})

	b.Run("Float64", func(b *testing.B) {
		rng := RNG(2345)
		for i := 0; i < b.N; i++ {
			blackholeFloat64 += rng.Float64()
		}
	})

	runtime.KeepAlive(blackholeUint64)
	runtime.KeepAlive(blackholeFloat64)
}
