package euler003

import "euler"

func maxPrimeFactor(val uint64) uint64 {
	for val != 0 && val&0x1 == 0 {
		val >>= 1
	}

	for divisor := uint64(3); divisor < val; divisor += 2 {
		for val%divisor == 0 {
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
