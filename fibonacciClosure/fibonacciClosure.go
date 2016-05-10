package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	counter := 0
	x, y := 0, 1
	return func() int {
		switch {
			case counter == 0: 
							counter ++
							return 0
			default:
							counter ++
		}
		y, x = x, x + y
		return x
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
