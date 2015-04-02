package euler006

import (
	"euler"
	"euler/math"
	"fmt"
)

func diff(max uint64) (result uint64, err error) {
	if max < 3 {
		return (0x2 & max) * max, nil
	}

	if math.MaxU64/max < max {
		return 0, fmt.Errorf("Integer overflow")
	}

	if math.MaxU64/(max*max) < (max * max) {
		return 0, fmt.Errorf("Integer overflow")
	}

	a, b := (max+1)>>2, (max-1)>>2
	a *= (b << 1) + 1

	result, b = a*(max-1), (max-3)>>2
	for a = 0; a < b; a++ {
		result -= ((a*a + a) << 3) + 2
	}
	a <<= 2
	a += 2
	result -= b * a
	if (max-1)&0x2 == 0 {
		result -= (b + 1) * a
	}

	a = 2 + 3*max
	result *= a
	a *= (1 - (a & 0x1)) * max * max
	result += a >> 3
	return
}

func Solution() (uint64, error) {
	return diff(100)
}

func init() {
	euler.RegisterSolution("006", Solution)
}
