package math

const (
	MaxU8  uint8  = ^uint8(0)
	MaxU16 uint16 = ^uint16(0)
	MaxU32 uint32 = ^uint32(0)
	MaxU64 uint64 = ^uint64(0)
)

func Log2U(val uint64) uint8 {
	if val == 0 {
		return 0
	}

	n := uint8(0)
	if val&0xffffffff00000000 != 0 {
		n += 32
		val >>= 32
	}

	if val&0x00000000ffff0000 != 0 {
		n += 16
		val >>= 16
	}

	if val&0x000000000000ff00 != 0 {
		n += 8
		val >>= 8
	}

	if val&0x00000000000000f0 != 0 {
		n += 4
		val >>= 4
	}

	if val&0xc != 0 {
		n += 2
		val >>= 2
	}

	return n + uint8(val>>1)
}

func Sqrt(val uint64) uint64 {
	var bits = Log2U(val)
	var x uint64 = 1 << ((bits >> 1) + 1)
	var calc_next_x func(uint64) uint64
	if bits < 32 {
		calc_next_x = func(prev_x uint64) uint64 {
			return ((prev_x*prev_x + val) / prev_x) >> 1
		}
	} else {
		calc_next_x = func(prev_x uint64) uint64 {
			return (val/prev_x + prev_x) >> 1
		}
	}
	next_x := calc_next_x(x)

	for next_x < x {
		x, next_x = next_x, calc_next_x(next_x)
	}

	return x
}
