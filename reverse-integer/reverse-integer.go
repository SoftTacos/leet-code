package reverseinteger

import (
	"math"
)

// this shouldn't work because the values here are going to be 64 bit ints, you can't
// compare the value of an int32 to see if it is larger than the max val of an int32,
// it'll just overflow before you get there, you could check for wrap around
// pretty sure the """ideal""" solution is to use a linked list, this would work
// for any kind of int, but that's also super slow and apparently this is fine
func reverse(x int) (r int) {
	if x == 0 {
		return
	}
	var maxSigFig int
	var digit int
	var v int
	var pow = 1
	var iterations = int(math.Floor(math.Log10(math.Abs(float64(x))))) + 1
	// 10^11 is larger than abs(MaxInt32)
	if iterations > 10 {
		return 0
	}
	for digit = 0; digit < iterations; digit++ {
		v = x % (pow * 10) / pow
		r = r*10 + v
		if v != 0 {
			maxSigFig = digit + 1
		}
		pow *= 10
	}
	r = r / int(math.Pow10(digit-maxSigFig))
	if r > math.MaxInt32 || r < math.MinInt32 {
		r = 0
	}
	return
}
