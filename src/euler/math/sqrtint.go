package math

const maxU64 uint64 = 0xffffffffffffffff

func sbc(val uint64) uint64 {
	var c uint64
	for val != 0 {
		c++
		val >>= 1
	}

	return c
}

func Sqrt(val uint64) uint64 {
	var bits = sbc(val)
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
