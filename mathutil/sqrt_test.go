package mathutil

import "testing"
//import "math"

func TestSqrt(t *testing.T) {
	cases := []struct {
		in, out float64
	} {
		{4, 2},
		{25, 5},
		{2, 1.414},
		{10, 3.162},
		{0, 0},
		//{-1, math.NaN()},
	}	
	for _, c := range cases {
		got := Sqrt(c.in)
		if (!((got - c.out) < 0.001)) {
			t.Errorf("Sqrt(%f) == %f, out %.3f", c.in, got, c.out)
		}
	}	
}