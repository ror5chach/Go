package main

import "fmt"

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %.f", float64(e))
	//return fmt.Sprint(float64(e))
}

func Sqrt(x float64) (float64, error) {
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
		return z, nil
	} else {
		if (x == 0) {
			return 0, nil
		} else {
			return -1, ErrNegativeSqrt(x)
		}
	}
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}