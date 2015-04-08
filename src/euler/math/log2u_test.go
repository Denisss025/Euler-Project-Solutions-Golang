package math

import (
	"math"
	"testing"
)

func TestFloorLog2(t *testing.T) {
	tests := map[uint64]uint64{
		1:     0,
		2:     1,
		3:     1,
		4:     2,
		8:     3,
		16:    4,
		32:    5,
		64:    6,
		128:   7,
		256:   8,
		512:   9,
		1024:  10,
		2048:  11,
		4096:  12,
		8192:  13,
		16384: 14,
		32768: 15,
		65536: 16,
	}

	for x2, x := range tests {
		sx := FloorLog2(x2)
		if x != sx {
			t.Errorf("Log2U(%v) == %v and != %v",
				x2, sx, x)
		}
	}

	tests2 := []uint64{10, 100, 1000, 1000000, 1000000000}
	for _, x := range tests2 {
		log2x1 := mathFloorLog2(x)
		log2x2 := FloorLog2(x)
		if log2x1 != log2x2 {
			t.Errorf("Wrong answer for value %v. "+
				"Expected %v. Acquired %v.", x, log2x1, log2x2)
		}
	}
}

func TestFloorLog2vsExp2(t *testing.T) {
	for i := uint64(0); i < 64; i++ {
		if x2, _ := Exp2(uint8(i)); i != FloorLog2(x2) {
			t.Errorf("Log2U(Exp2(%v)) != %v", x2, i)
		}
	}
}

func mathFloorLog2(val uint64) uint64 {
	return uint64(math.Log2(float64(val)))
}

func benchFloorLog2Impl(b *testing.B, val uint64) {
	for n := 0; n < b.N; n++ {
		FloorLog2(val)
	}
}

func benchMathFloorLog2Impl(b *testing.B, val uint64) {
	for n := 0; n < b.N; n++ {
		mathFloorLog2(val)
	}
}

func BenchmarkFloorLog2K(b *testing.B) {
	benchFloorLog2Impl(b, 1000)
}

func BenchmarkMathFloorLog2K(b *testing.B) {
	benchMathFloorLog2Impl(b, 1000)
}

func BenchmarkFloorLog2M(b *testing.B) {
	benchFloorLog2Impl(b, 1000000)
}

func BenchmarkMathFloorLog2M(b *testing.B) {
	benchMathFloorLog2Impl(b, 1000000)
}

func BenchmarkFloorLog2G(b *testing.B) {
	benchFloorLog2Impl(b, 1000000000)
}

func BenchmarkMathFloorLog2G(b *testing.B) {
	benchMathFloorLog2Impl(b, 1000000000)
}
