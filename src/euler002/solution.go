package euler002

import "euler"

// FibGenerator is a Fibonacci sequence generator.
type FibGenerator struct{ n, p uint64 }

// NewFibGenerator creates a new Fibonacci sequence generator.
func NewFibGenerator(n, p uint64) *FibGenerator { return &FibGenerator{n, p} }

// Next calculates and returns the next Fibonacci number.
func (f *FibGenerator) Next() uint64 {
	f.p, f.n = f.n, f.n+f.p
	return f.p
}

// sumEvenFib calculates the sum of the even-valued terms of the Fibonacci
// sequence.
func sumEvenFib(limit uint64) (sum uint64, err error) {
	fib := NewFibGenerator(0, 1) // create new Fibonacci sequence generator
	for n := uint64(0); n < limit; n = fib.Next() {
		if n&1 == 0 {  // check if n is even
			sum += n
		}
	}
	return
}

// Solution solves the Project Euler Problem #2 and returns a solution.
func Solution() (uint64, error) {
	return sumEvenFib(4e6)
}

// init registers the Solution function.
func init() {
	euler.RegisterSolution("002", Solution)
}
