package euler015

import (
	"errors"
	"euler"
)

func countRotes(n uint8) (result uint64, err error) {
	if n == 0 {
		return
	}

	if n > 31 {
		err = errors.New("Integer overflow")
		return
	}

	result = 1
	n2 := uint64(n)*2 + 1
	for i := uint64(n + 1); i < n2; i++ {
		result *= i
		result /= i - uint64(n)
	}

	return
}

func Solution() (uint64, error) {
	return countRotes(20)
}

func init() {
	euler.RegisterSolution("015", Solution)
}
