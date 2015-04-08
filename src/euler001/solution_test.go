package euler001

import "testing"

func TestSum3N5(t *testing.T) {
	sum := uint64(0)

	for i := uint64(0); i < 10000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}

		if r, e := sum3n5(i + 1); e != nil {
			t.Errorf("Unexpected error: %v", e)
		} else if r != sum {
			t.Errorf("Wrong result for %v. Expected %v. Acquired %v",
				i, sum, r)
		}
	}
}
