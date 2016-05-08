// package with some functions for math calculations
package mathutil

import "math"

// returns squareroot of a floating point value
func Sqrt(x float64) float64 {
	if (x > 0) {
		// pick a starting point
		z := x/2		
		delta := 1.0
		for ; delta > 0.000001; {
			// Newtons formula
			nextZ := z - ((z * z) - x)/(2 * z)		
			delta = z - nextZ
			if (delta < 0) {
				delta = 0.0 - delta
			}
			z = nextZ
		}
		return z
	} else {
		if (x == 0) {
			return 0
		} else {
			return math.NaN()
		}
	}
}