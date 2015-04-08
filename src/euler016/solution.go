package euler016

import (
	"euler"
)

func calcDigitsForPow2(pow uint64) uint64 {
	l := 1 + (pow >> 5)
	v := make([]uint32, l, l)
	v[0] = uint32(1) << (pow & 0x1F)

	q := uint64(0)
	sum := uint64(0)
	const d = uint64(1000000000)

	for len(v) != 0 {
		for i := range v {
			q <<= 32
			q |= uint64(v[i])
			v[i] = uint32(q / d)
			q -= uint64(v[i]) * d
		}
		for q != 0 {
			sum += q % 10
			q /= 10
		}
		if v[0] == 0 {
			v = v[1:]
		}
	}

	return sum
}

func Solution() (uint64, error) {
	return calcDigitsForPow2(1000), nil
}

func init() {
	euler.RegisterSolution("016", Solution)
}
