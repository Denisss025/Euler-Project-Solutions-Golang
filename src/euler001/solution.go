package euler001

import (
	"euler"
	"euler/math"
	"fmt"
)

func sum3n5(limit uint64) (uint64, error) {
	var limDiv03 uint64 = (limit - 1) / 3
	var limDiv05 uint64 = (limit - 1) / 5
	var limDiv15 uint64 = (limit - 1) / 15

	if math.Sqrt(math.MaxU64) < 1+limDiv03 {
		return 0, fmt.Errorf("%s", "Integer overflow for limDiv03")
	}

	var sum03 uint64 = (limDiv03*(limDiv03+1) - 5*limDiv15*(limDiv15+1)) / 2
	if math.MaxU64/3 < sum03 {
		return 0, fmt.Errorf("%s", "Integer overflow for sum03 * 3")
	}
	sum03 *= 3

	if (math.Sqrt(math.MaxU64-sum03)/5)*2 < 1+limDiv05 {
		return 0, fmt.Errorf("%s", "Integer overflow for final result")
	}

	return sum03 + 5*limDiv05*(limDiv05+1)/2, nil
}

func Solution() (uint64, error) {
	return sum3n5(1000)
}

func init() {
	euler.RegisterSolution("001", Solution)
}
