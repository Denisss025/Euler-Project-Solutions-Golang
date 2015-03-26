package euler001

import (
    "euler"
	"euler/math"
	"fmt"
)

const (
	maxUInt64 uint64 = 0xffffffffffffffff
	limit     uint64 = 1000
)

func Solution() (uint64, error) {
	const limDiv03 uint64 = (limit - 1) / 3
	const limDiv05 uint64 = (limit - 1) / 5
	const limDiv15 uint64 = (limit - 1) / 15

	if math.Sqrt(maxUInt64) < 1+limDiv03 {
		return 0, fmt.Errorf("%s", "Integer overflow for limDiv03")
	}

	var sum03 uint64 = (limDiv03*(limDiv03+1) - 5*limDiv15*(limDiv15+1)) / 2
	if maxUInt64/3 < sum03 {
		return 0, fmt.Errorf("%s", "Integer overflow for sum03 * 3")
	}
	sum03 *= 3

	if (math.Sqrt(maxUInt64-sum03)/5)*2 < 1+limDiv05 {
		return 0, fmt.Errorf("%s", "Integer overflow for final result")
	}

	return sum03 + 5*limDiv05*(limDiv05+1)/2, nil
}

func init() {
    euler.RegisterSolution("001", Solution)
}
