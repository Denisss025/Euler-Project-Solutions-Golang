package math

func GenPrimes(max uint64, prime chan uint64) {
	if max >= 2 {
		prime <- 2
	}

	if max >= 3 {
		prime <- 3
	}

	prVec := [2][]bool{
		make([]bool, 6+max/6),
		make([]bool, 6+max/6),
	}

	for i := range prVec {
		for j := range prVec[i] {
			prVec[i][j] = true
		}
	}

	n := len(prVec[0])

	num := uint64(0)

	for i := range prVec[0] {
		num += 5
		if num > max {
			prime <- 0
			return
		}

		if prVec[0][i] {
			prime <- num
			for j := i + int(num); j < n; j += int(num) {
				prVec[0][j] = false
			}
			for j := int(num) - i - 2; j < n; j += int(num) {
				prVec[1][j] = false
			}
		}

		num += 2
		if num > max {
			prime <- 0
			return
		}
		if prVec[1][i] {
			prime <- num
			for j := i + int(num); j < n; j += int(num) {
				prVec[1][j] = false
			}
			for j := int(num) - i - 2; j < n; j += int(num) {
				prVec[0][j] = false
			}
		}
		num--
	}
	prime <- 0
}

type PrimeGenerator struct {
	prime chan uint64
}

func NewPrimeGenerator(max uint64) *PrimeGenerator {
	pg := &PrimeGenerator{prime: make(chan uint64, 10)}
	go GenPrimes(max, pg.prime)
	return pg
}

func (pg *PrimeGenerator) Next() uint64 {
	return <-pg.prime
}
