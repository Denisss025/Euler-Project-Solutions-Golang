package euler007

import (
	"euler"
	"euler/math"
)

func nthPrime(n uint16) uint64 {
	if n < 3 {
		return uint64(n) + 1
	}

	primeGen := math.NewPrimeGenerator(uint64(n) * 13)
	primeGen.Next() // skip 2
	primeGen.Next() // skip 3

	n -= 3
	for n != 0 {
		primeGen.Next()
		n--
	}
	return primeGen.Next()
}

func Solution() (uint64, error) {
	return nthPrime(10001), nil
}

func init() {
	euler.RegisterSolution("007", Solution)
}
