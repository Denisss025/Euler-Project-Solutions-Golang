package euler010

import (
	"euler"
	"euler/math"
	"fmt"
)

func sumPrimesLessThan(val uint64) (sum uint64, err error) {
	primeGen := math.NewPrimeGenerator(val)
	for i := primeGen.Next(); i != 0; i = primeGen.Next() {
		if math.MaxU64-i < sum {
			err = fmt.Errorf("Integer overflow")
			return
		}
		sum += i
	}

	return
}

func Solution() (uint64, error) {
	return sumPrimesLessThan(2000000)
}

func init() {
	euler.RegisterSolution("010", Solution)
}
