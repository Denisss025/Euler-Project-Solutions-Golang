package euler002

import (
    "euler"
)

func Solution() (uint64, error) {
    var prev uint64
    var curr uint64 = 1
    var sum uint64
    
    for curr < 4000000 {
        if (curr & 0x1 == 0x0) {
            sum += curr
        }
        
        prev, curr = curr, curr + prev
    }
    
    return sum, nil
}

func init() {
    euler.RegisterSolution("002", Solution)
}
