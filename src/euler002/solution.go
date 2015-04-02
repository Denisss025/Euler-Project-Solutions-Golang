package euler002

import "euler"

func fibGen(fib chan uint64) {
	fib <- 0

	prev, curr := uint64(0), uint64(1)

	for {
		fib <- curr
		prev, curr = curr, curr+prev
	}
}

func sumEvenFib(limit uint64) (sum uint64, err error) {
	fib := make(chan uint64, 10)
	go fibGen(fib)

	for fibNum := <-fib; fibNum < limit; fibNum = <-fib {
		if fibNum&0x1 == 0x0 {
			sum += fibNum
		}
	}

	return
}

func Solution() (uint64, error) {
	return sumEvenFib(4000000)
}

func init() {
	euler.RegisterSolution("002", Solution)
}
