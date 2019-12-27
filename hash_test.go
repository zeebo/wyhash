package wyhash

import (
	"fmt"
	"runtime"
	"testing"
)

func TestHash(t *testing.T) {
	var data []byte

	for i, want := range vecs {
		data = append(data, byte(i))
		got := Hash(data[:i], 0x0102030405060708)
		if i != 0 && got != want {
			t.Errorf("Hash(...%d)=%x, want %x", i, got, want)
		}
	}
}

func BenchmarkCompare(b *testing.B) {
	sizes := []int{
		0, 1, 3, 4, 8, 9, 16, 17, 32,
		33, 64, 65, 96, 97, 128, 129, 240, 241,
		512, 1024, 100 * 1024,
	}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
			b.SetBytes(int64(size))
			var acc uint64
			d := string(make([]byte, size))
			b.ReportAllocs()
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				acc = HashString(d, 0)
			}
			runtime.KeepAlive(acc)
		})
	}
}
