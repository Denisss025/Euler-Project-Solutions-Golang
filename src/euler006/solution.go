package euler006

import "euler"

func diff(n uint64) (uint64, error) {
	s := n
	n++
	s *= n
	// sum = n*(n+1)/2
	s >>= 1
	n <<= 1
	n--
	n *= s
	// n = sum * (2*(n + 1) - 1) / 3 =
	//   = n(n+1)(2n + 1)/6 =
	//   = (2*n^3 + 3*n^2 + n)/6 =
	//   = (n^3)/3 + (n^2)/2 + n/6
	n /= 3
	// s = square(sum)
	s *= s
	return s - n, nil

}

func Solution() (uint64, error) {
	return diff(100)
}

func init() {
	euler.RegisterSolution("006", Solution)
}
