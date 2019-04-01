package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	matrix := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		line := make([]uint8, dx)
		for x := 0; x < dx; x++ {
			line[x] = uint8(x ^ y)
		}
		matrix[y] = line
	}

	return matrix
}

func main() {
	pic.Show(Pic)
}
