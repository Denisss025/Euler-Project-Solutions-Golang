package euler004

import (
	"euler"
	"fmt"
)

func lrMaxMin(digits uint8) (min uint64, max uint64) {
	min = 1
	for i := digits; i > 1; i-- {
		min *= 10
	}
	max = min * 10
	return
}

func palindromeGenerator(result chan uint64, digits uint8) {
	min, max := lrMaxMin(digits)

	for p := max - 1; p >= min; p-- {
		var palindrome uint64
		n := p
		for i := digits; i > 0; i-- {
			palindrome *= 10
			palindrome += n % 10
			n /= 10
		}
		palindrome += max * p
		result <- palindrome
	}
	result <- 0
}

func maxProductPalindrome(digits uint8) (uint64, error) {
	palindrome := make(chan uint64, 10)

	go palindromeGenerator(palindrome, digits)

	var min uint64 = 1
	for i := digits; i > 1; i-- {
		min *= 10
	}
	max := min * 10

	for pl := <-palindrome; pl != 0; pl = <-palindrome {
		for n := max; n >= min && n >= pl/n; n-- {
			if pl%n == 0 {
				return pl, nil
			}
		}
	}

	return 0, fmt.Errorf("%s", "Cannot find solution")
}

func Solution() (uint64, error) {
	return maxProductPalindrome(3)
}

func init() {
	euler.RegisterSolution("004", Solution)
}
