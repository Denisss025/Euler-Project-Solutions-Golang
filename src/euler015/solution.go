package euler015

import (
	"errors"
	"euler"
	"euler/math"
)

func countRotes(n uint8) (result uint64, err error) {
	if n == 0 {
		return
	}

	result = 1
	n2 := uint64(n)*2 + 1
	for i := uint64(2); i < n2; i++ {
		if math.MaxU64/i < result {
			err = errors.New("Integer overflow")
			result = 0
			return
		}
		result *= i
		result /= (i + 1) >> 1
	}

	return
}

func Solution() (uint64, error) {
	return countRotes(20)
}

func init() {
	euler.RegisterSolution("015", Solution)
}
