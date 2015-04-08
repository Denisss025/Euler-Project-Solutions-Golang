package euler015

import "testing"

func benchImpl(b *testing.B, x uint8) {
	for n := 0; n < b.N; n++ {
		countRotes(x)
	}
}

func BenchmarkCountRoutes00(b *testing.B) {
	benchImpl(b, 0)
}

func BenchmarkCountRoutes05(b *testing.B) {
	benchImpl(b, 5)
}

func BenchmarkCountRoutes10(b *testing.B) {
	benchImpl(b, 10)
}

func BenchmarkCountRoutes20(b *testing.B) {
	benchImpl(b, 20)
}

func BenchmarkCountRoutesr30(b *testing.B) {
	benchImpl(b, 30)
}

func TestCountRoutes(t *testing.T) {
	tests := map[uint8]uint64{
		0:  0,
		1:  2,
		2:  6,
		3:  20,
		4:  70,
		5:  252,
		6:  924,
		20: 137846528820,
	}

	for v, k := range tests {
		if c, err := countRotes(v); err != nil {
			t.Errorf("Got error '%v' for %v", err, v)
		} else if c != k {
			t.Errorf("Routes for %v is %v, "+
				"but countRoutes returned %v value", v, k, c)
		}
	}
}
