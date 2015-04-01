package math

import "testing"

func TestLog2U(t *testing.T) {
	tests := map[uint64]uint8{
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
		sx := Log2U(x2)
		if x != sx {
			t.Errorf("Log2U(%v) == %v and != %v",
				x2, sx, x)
		}
	}
}

func TestLog2UvsExp2(t *testing.T) {
	var i uint8
	for i = 0; i < 64; i++ {
		if x2, _ := Exp2(i); i != Log2U(x2) {
			t.Errorf("Log2U(Exp2(%v)) != %v", x2, i)
		}
	}
}
