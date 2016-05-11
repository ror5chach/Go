package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	picture := make([][]uint8, dy)
	for y,_ := range picture {
		picture[y] = make([]uint8, dx)
		
		for x,_ := range picture[y] {	
			picture[y][x] = uint8(x^y)//(x*y)//((x+y/2))
		}
	}
		
	return picture
}

func main() {
	pic.Show(Pic)
}
