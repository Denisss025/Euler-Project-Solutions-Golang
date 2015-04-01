package math

import "fmt"

func Exp2(pow uint8) (uint64, error) {
	if pow > 63 {
		return 0, fmt.Errorf("Integer overflow")
	}
	return uint64(1) << pow, nil
}

func Pow(x uint64, pow uint8) (uint64, error) {
	var result uint64 = 1

	for pow != 0 {
		if pow&0x1 != 0 {
			if MaxU64/x < result {
				return 0, fmt.Errorf(
					"Integer overflow: result = %v; x = %v",
					result, x,
				)
			}
			result *= x
			pow--
		}
		if pow != 0 && MaxU64/x < x {
			return 0, fmt.Errorf("Integer overflow: x(%v vs %v)",
				x, MaxU64/x)
		}
		x *= x
		pow >>= 1
	}

	return result, nil
}
