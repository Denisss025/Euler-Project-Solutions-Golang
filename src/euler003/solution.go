package euler003

import (
	"euler"
	"euler/math"
)

func maxPrimeFactor(val uint64) uint64 {
	for val > 2 && val&0x1 == 0 {
		val >>= 1
	}

	for val > 3 && val%3 == 0 {
		val /= 3
	}

	for divisor := uint64(5); divisor < math.Sqrt(val); divisor += 4 {
		for val > divisor && val%divisor == 0 {
			val /= divisor
		}

		divisor += 2
		for val > divisor && val%divisor == 0 {
			val /= divisor
		}
	}
	return val
}

func Solution() (uint64, error) {
	return maxPrimeFactor(600851475143), nil
}

func init() {
	euler.RegisterSolution("003", Solution)
}
