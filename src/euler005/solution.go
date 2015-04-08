package euler005

import (
	"euler"
	emath "euler/math"
	"math"
)

func getSmallestDivisible(max uint8) (uint64, error) {
	if max < 3 {
		return uint64(max), nil
	}

	v := make([]bool, (max-1)/2)
	for i := range v {
		v[i] = true
	}

	log := math.Log2(float64(max))
	result := uint64(1) << emath.FloorLog2(uint64(max))

	pg := emath.NewPrimeGenerator(uint64(max))
	pg.Next() // skip 2
	for num := pg.Next(); num < uint64(max) && num != 0; num = pg.Next() {
		pow, err := emath.Pow(uint64(num),
			uint8(float64(log)/math.Log2(float64(num))))
		if err != nil {
			return result, err
		} else {
			result *= pow
		}
	}

	return result, nil
}

func Solution() (uint64, error) {
	return getSmallestDivisible(20)
}

func init() {
	euler.RegisterSolution("005", Solution)
}
