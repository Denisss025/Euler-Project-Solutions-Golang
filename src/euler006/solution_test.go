package euler006

import "testing"

func TestDiff(t *testing.T) {
	sum := uint64(0)
	sumq := uint64(0)

	for i := uint64(0); i < 101; i++ {
		sum += i
		sumq += i * i
		d1 := sum*sum - sumq
		if d2, err := diff(i); err != nil {
			t.Errorf("Unexpected error: %v", err)
		} else if d2 != d1 {
			t.Errorf("Wrong result for %v. "+
				"Expected %v. Acquired %v.", i, d1, d2)
		}
	}
}

func benchImpl(b *testing.B, val uint64) {
	for n := 0; n < b.N; n++ {
		diff(val)
	}
}

func BenchmarkDiff00(b *testing.B) {
	benchImpl(b, 0)
}

func BenchmarkDiff05(b *testing.B) {
	benchImpl(b, 5)
}

func BenchmarkDiff10(b *testing.B) {
	benchImpl(b, 10)
}

func BenchmarkDiff20(b *testing.B) {
	benchImpl(b, 20)
}

func BenchmarkDiff25(b *testing.B) {
	benchImpl(b, 25)
}

func BenchmarkDiff50(b *testing.B) {
	benchImpl(b, 50)
}

func BenchmarkDiff75(b *testing.B) {
	benchImpl(b, 75)
}

func BenchmarkDiff99(b *testing.B) {
	benchImpl(b, 99)
}
