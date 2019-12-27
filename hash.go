package wyhash

import (
	"encoding/binary"
	"math/bits"
	"unsafe"
)

type (
	ptr  = unsafe.Pointer
	uptr = uintptr

	u32 = uint32
	u64 = uint64
)

const (
	_wyp0 = 0xa0761d6478bd642f
	_wyp1 = 0xe7037ed1a0b428db
	_wyp2 = 0x8ebc6af09c88c6e3
	_wyp3 = 0x589965cc75374cc3
	_wyp4 = 0x1d8e4e27c47d124f
)

func i(p ptr, n uptr) ptr { return ptr(uptr(p) + n) }

func _wymum(A, B u64) u64 {
	hi, lo := bits.Mul64(A, B)
	return hi ^ lo
}

func _wyr8(p ptr) u64 {
	return binary.LittleEndian.Uint64((*[8]byte)(p)[:])
}

func _wyr4(p ptr) u64 {
	return u64(binary.LittleEndian.Uint32((*[4]byte)(p)[:]))
}

func _wyr3(p ptr, k uptr) u64 {
	b0 := u64(*(*byte)(p))
	b1 := u64(*(*byte)(i(p, k>>1)))
	b2 := u64(*(*byte)(i(p, k-1)))
	return b0<<16 | b1<<8 | b2
}

func _wyr9(p ptr) u64 {
	b := (*[8]byte)(p)
	return u64(u32(b[0])|u32(b[1])<<8|u32(b[2])<<16|u32(b[3])<<24)<<32 |
		u64(u32(b[4])|u32(b[5])<<8|u32(b[6])<<16|u32(b[7])<<24)
}

func hash(data string, seed u64) u64 {
	p, len := *(*ptr)(ptr(&data)), uptr(len(data))
	see1, off := seed, len

	switch {
	case len <= 0x03:
		return _wymum(_wymum(_wyr3(p, len)^seed^_wyp0, seed^_wyp1)^seed, u64(len)^_wyp4)

	case len <= 0x08:
		return _wymum(_wymum(_wyr4(i(p, 0x00))^seed^_wyp0, _wyr4(i(p, len-0x04))^seed^_wyp1)^seed, u64(len)^_wyp4)

	case len <= 0x10:
		return _wymum(_wymum(_wyr9(i(p, 0x00))^seed^_wyp0, _wyr9(i(p, len-0x08))^seed^_wyp1)^seed, u64(len)^_wyp4)

	case len <= 0x18:
		return _wymum(_wymum(_wyr9(i(p, 0x00))^seed^_wyp0, _wyr9(i(p, 0x08))^seed^_wyp1)^_wymum(_wyr9(i(p, len-0x08))^seed^_wyp2, seed^_wyp3), u64(len)^_wyp4)

	case len <= 0x20:
		return _wymum(_wymum(_wyr9(i(p, 0x00))^seed^_wyp0, _wyr9(i(p, 0x08))^seed^_wyp1)^_wymum(_wyr9(i(p, 0x10))^seed^_wyp2, _wyr9(i(p, len-0x08))^seed^_wyp3), u64(len)^_wyp4)

	case len <= 0x100:
		seed = _wymum(_wyr8(i(p, 0x00))^seed^_wyp0, _wyr8(i(p, 0x08))^seed^_wyp1)
		see1 = _wymum(_wyr8(i(p, 0x10))^see1^_wyp2, _wyr8(i(p, 0x18))^see1^_wyp3)
		if len > 0x40 {
			seed = _wymum(_wyr8(i(p, 0x20))^seed^_wyp0, _wyr8(i(p, 0x28))^seed^_wyp1)
			see1 = _wymum(_wyr8(i(p, 0x30))^see1^_wyp2, _wyr8(i(p, 0x38))^see1^_wyp3)
			if len > 0x60 {
				seed = _wymum(_wyr8(i(p, 0x40))^seed^_wyp0, _wyr8(i(p, 0x48))^seed^_wyp1)
				see1 = _wymum(_wyr8(i(p, 0x50))^see1^_wyp2, _wyr8(i(p, 0x58))^see1^_wyp3)
				if len > 0x80 {
					seed = _wymum(_wyr8(i(p, 0x60))^seed^_wyp0, _wyr8(i(p, 0x68))^seed^_wyp1)
					see1 = _wymum(_wyr8(i(p, 0x70))^see1^_wyp2, _wyr8(i(p, 0x78))^see1^_wyp3)
					if len > 0xa0 {
						seed = _wymum(_wyr8(i(p, 0x80))^seed^_wyp0, _wyr8(i(p, 0x88))^seed^_wyp1)
						see1 = _wymum(_wyr8(i(p, 0x90))^see1^_wyp2, _wyr8(i(p, 0x98))^see1^_wyp3)
						if len > 0xc0 {
							seed = _wymum(_wyr8(i(p, 0xa0))^seed^_wyp0, _wyr8(i(p, 0xa8))^seed^_wyp1)
							see1 = _wymum(_wyr8(i(p, 0xb0))^see1^_wyp2, _wyr8(i(p, 0xb8))^see1^_wyp3)
							if len > 0xe0 {
								seed = _wymum(_wyr8(i(p, 0xc0))^seed^_wyp0, _wyr8(i(p, 0xc8))^seed^_wyp1)
								see1 = _wymum(_wyr8(i(p, 0xd0))^see1^_wyp2, _wyr8(i(p, 0xd8))^see1^_wyp3)
							}
						}
					}
				}
			}
		}

		off = (off-1)%0x20 + 1
		p = i(p, len-off)

	default:
		for ; off > 0x100; off, p = off-0x100, i(p, 0x100) {
			seed = _wymum(_wyr8(i(p, 0x00))^seed^_wyp0, _wyr8(i(p, 0x08))^seed^_wyp1) ^ _wymum(_wyr8(i(p, 0x10))^seed^_wyp2, _wyr8(i(p, 0x18))^seed^_wyp3)
			see1 = _wymum(_wyr8(i(p, 0x20))^see1^_wyp1, _wyr8(i(p, 0x28))^see1^_wyp2) ^ _wymum(_wyr8(i(p, 0x30))^see1^_wyp3, _wyr8(i(p, 0x38))^see1^_wyp0)
			seed = _wymum(_wyr8(i(p, 0x40))^seed^_wyp0, _wyr8(i(p, 0x48))^seed^_wyp1) ^ _wymum(_wyr8(i(p, 0x50))^seed^_wyp2, _wyr8(i(p, 0x58))^seed^_wyp3)
			see1 = _wymum(_wyr8(i(p, 0x60))^see1^_wyp1, _wyr8(i(p, 0x68))^see1^_wyp2) ^ _wymum(_wyr8(i(p, 0x70))^see1^_wyp3, _wyr8(i(p, 0x78))^see1^_wyp0)
			seed = _wymum(_wyr8(i(p, 0x80))^seed^_wyp0, _wyr8(i(p, 0x88))^seed^_wyp1) ^ _wymum(_wyr8(i(p, 0x90))^seed^_wyp2, _wyr8(i(p, 0x98))^seed^_wyp3)
			see1 = _wymum(_wyr8(i(p, 0xa0))^see1^_wyp1, _wyr8(i(p, 0xa8))^see1^_wyp2) ^ _wymum(_wyr8(i(p, 0xb0))^see1^_wyp3, _wyr8(i(p, 0xb8))^see1^_wyp0)
			seed = _wymum(_wyr8(i(p, 0xc0))^seed^_wyp0, _wyr8(i(p, 0xc8))^seed^_wyp1) ^ _wymum(_wyr8(i(p, 0xd0))^seed^_wyp2, _wyr8(i(p, 0xd8))^seed^_wyp3)
			see1 = _wymum(_wyr8(i(p, 0xe0))^see1^_wyp1, _wyr8(i(p, 0xe8))^see1^_wyp2) ^ _wymum(_wyr8(i(p, 0xf0))^see1^_wyp3, _wyr8(i(p, 0xf8))^see1^_wyp0)
		}
		for ; off > 0x20; off, p = off-0x20, i(p, 0x20) {
			seed = _wymum(_wyr8(i(p, 0x00))^seed^_wyp0, _wyr8(i(p, 0x08))^seed^_wyp1)
			see1 = _wymum(_wyr8(i(p, 0x10))^see1^_wyp2, _wyr8(i(p, 0x18))^see1^_wyp3)
		}
	}

	switch {
	case off > 0x18:
		seed = _wymum(_wyr9(i(p, 0x00))^seed^_wyp0, _wyr9(i(p, 0x08))^seed^_wyp1)
		see1 = _wymum(_wyr9(i(p, 0x10))^see1^_wyp2, _wyr9(i(p, off-0x08))^see1^_wyp3)

	case off > 0x10:
		seed = _wymum(_wyr9(i(p, 0x00))^seed^_wyp0, _wyr9(i(p, 0x08))^seed^_wyp1)
		see1 = _wymum(_wyr9(i(p, off-0x08))^see1^_wyp2, see1^_wyp3)

	case off > 0x08:
		seed = _wymum(_wyr9(i(p, 0x00))^seed^_wyp0, _wyr9(i(p, off-0x08))^seed^_wyp1)

	case off > 0x03:
		seed = _wymum(_wyr4(i(p, 0x00))^seed^_wyp0, _wyr4(i(p, off-0x04))^seed^_wyp1)

	default:
		seed = _wymum(_wyr3(p, off)^seed^_wyp0, seed^_wyp1)
	}

	return _wymum(seed^see1, u64(len)^_wyp4)
}

// Hash returns a 64bit digest of the data with different ones for every seed.
func Hash(data []byte, seed uint64) uint64 {
	if len(data) == 0 {
		return seed
	}
	return hash(*(*string)(ptr(&data)), seed)
}

// HashString returns a 64bit digest of the data with different ones for every seed.
func HashString(data string, seed uint64) uint64 {
	if len(data) == 0 {
		return seed
	}
	return hash(*(*string)(ptr(&data)), seed)
}
