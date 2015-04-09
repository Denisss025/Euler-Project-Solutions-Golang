package euler018

import "euler"

var tr = [][]byte{
	[]byte{75},
	[]byte{95, 64},
	[]byte{17, 47, 82},
	[]byte{18, 35, 87, 10},
	[]byte{20, 4, 82, 47, 65},
	[]byte{19, 1, 23, 75, 3, 34},
	[]byte{88, 2, 77, 73, 7, 63, 67},
	[]byte{99, 65, 4, 28, 6, 16, 70, 92},
	[]byte{41, 41, 26, 56, 83, 40, 80, 70, 33},
	[]byte{41, 48, 72, 33, 47, 32, 37, 16, 94, 29},
	[]byte{53, 71, 44, 65, 25, 43, 91, 52, 97, 51, 14},
	[]byte{70, 11, 33, 28, 77, 73, 17, 78, 39, 68, 17, 57},
	[]byte{91, 71, 52, 38, 17, 14, 91, 43, 58, 50, 27, 29, 48},
	[]byte{63, 66, 4, 68, 89, 53, 67, 30, 73, 16, 69, 87, 40, 31},
	[]byte{4, 62, 98, 27, 23, 9, 70, 98, 73, 93, 38, 53, 60, 4, 23},
}

func maxU64(a, b uint64) uint64 {
	if a > b {
		return a
	}
	return b
}

func maxTotal(triangle [][]byte) (result uint64) {
	tLen := len(triangle[len(triangle)-1]) + 1
	totals := make([]uint64, tLen, tLen)

	swp := uint64(0)
	for _, line := range triangle {
		swp = totals[0]
		for j, val := range line {
			if swp > totals[j+1] {
				swp, totals[j+1] = totals[j+1], swp
			} else {
				swp = totals[j+1]
			}
			totals[j+1] += uint64(val)
		}
	}

	result = 0
	for _, val := range totals {
		if val > result {
			result = val
		}
	}
	return
}

func Solution() (uint64, error) {
	return maxTotal(tr), nil
}

func init() {
	euler.RegisterSolution("018", Solution)
}
