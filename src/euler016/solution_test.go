package euler016

import "testing"

func TestCalc15(t *testing.T) {
	if r := calcDigitsForPow2(15); r != 26 {
		t.Errorf("Wrong result for value 15. Expected 26. Acquired %v",
			r)
	}
}

func benchImpl(b *testing.B, val uint64) {
	for n := 0; n < b.N; n++ {
		calcDigitsForPow2(val)
	}
}

func BenchmarkCalc0(b *testing.B) {
	benchImpl(b, 0)
}

func BenchmarkCalc50(b *testing.B) {
	benchImpl(b, 50)
}

func BenchmarkCalc100(b *testing.B) {
	benchImpl(b, 100)
}

func BenchmarkCalc1000(b *testing.B) {
	benchImpl(b, 1000)
}
