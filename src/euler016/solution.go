package euler016

import (
	"euler"
)

var dsum = [100]byte{
	0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
	1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	3, 4, 5, 6, 7, 8, 9, 10, 11, 12,
	4, 5, 6, 7, 8, 9, 10, 11, 12, 13,
	5, 6, 7, 8, 9, 10, 11, 12, 13, 14,
	6, 7, 8, 9, 10, 11, 12, 13, 14, 15,
	7, 8, 9, 10, 11, 12, 13, 14, 15, 16,
	8, 9, 10, 11, 12, 13, 14, 15, 16, 17,
	9, 10, 11, 12, 13, 14, 15, 16, 17, 18,
}

func calcDigitsForPow2(pow uint64) (sum uint64) {
	l := 1 + (pow >> 5)
	v := make([]uint32, l, l)
	v[0] = uint32(1) << (pow & 0x1F)

	q := uint64(0)
	const d = uint64(1000000000)

	for len(v) != 0 {
		for i := range v {
			q <<= 32
			q |= uint64(v[i])
			v[i] = uint32(q / d)
			q -= uint64(v[i]) * d
		}
		for q != 0 {
			sum += uint64(dsum[q%100])
			q /= 100
		}
		if v[0] == 0 {
			v = v[1:]
		}
	}

	return
}

func Solution() (uint64, error) {
	return calcDigitsForPow2(1000), nil
}

func init() {
	euler.RegisterSolution("016", Solution)
}
