package euler003

import (
	"euler"
)

func maxPrimeFactor(val uint64) uint64 {
	var divisor uint64 = 3

	for val&0x1 == 0 {
		val >>= 1
	}

	for divisor < val {
		for val%divisor == 0 {
			val /= divisor
		}
		divisor += 2
	}

	return val
}

func Solution() (uint64, error) {
	return maxPrimeFactor(600851475143), nil
}

func init() {
	euler.RegisterSolution("003", Solution)
}
