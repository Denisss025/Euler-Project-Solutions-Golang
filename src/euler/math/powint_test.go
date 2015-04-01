package math

import "testing"

func TestExp2(t *testing.T) {
	tests := map[uint8]uint64{
		0:  1,
		1:  2,
		2:  4,
		3:  8,
		4:  16,
		5:  32,
		6:  64,
		7:  128,
		8:  256,
		9:  512,
		10: 1024,
		11: 2048,
		12: 4096,
		13: 8192,
		14: 16384,
		15: 32768,
		16: 65536,
	}

	for x, x2 := range tests {
		xx, _ := Exp2(x)
		if x2 != xx {
			t.Errorf("Exp2(%d) == %d and != %d", x, xx, x2)
		}
	}

	if _, err := Exp2(64); err == nil {
		t.Errorf("Exp2(64) does not returned error")
	}
}

func TestExp2vsPow(t *testing.T) {
	for i := uint8(0); i < 64; i++ {
		e2, err1 := Exp2(i)
		p2, err2 := Pow(2, i)

		if err1 != nil || err2 != nil {
			t.Errorf("Should not get error %d: (%v), (%v)",
				i, err1, err2)
		} else if e2 != p2 {
			t.Errorf("Got different results for %d: %v != %v",
				i, e2, p2)
		}
	}

	if _, err := Pow(2, 64); err == nil {
		t.Errorf("Should get error")
	}
}

func TestPow0(t *testing.T) {
	for i := uint64(1); i < 10000; i += 200 {
		if p0, err := Pow(i, 0); err != nil {
			t.Errorf("Got error on %v", i)
		} else if p0 != 1 {
			t.Errorf("Got wrong result for %v: %v", i, p0)
		}
	}
}

func TestPow(t *testing.T) {
	for i := uint8(2); i < 10; i++ {
		for j := uint64(1); j < 100; j++ {
			p := uint64(1)
			for pow := i; pow != 0; pow-- {
				p *= j
			}
			if p2, err := Pow(j, i); err != nil {
				t.Errorf("Got error on Pow(%v, %v): %v",
					j, i, err)
			} else if p != p2 {
				t.Errorf("Pow(%v, %v) == %v and != %v",
					j, i, p2, p)
			}
		}
	}
}
