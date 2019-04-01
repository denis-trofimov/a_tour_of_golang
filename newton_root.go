package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (float64, float64, int) {
	const stop_value = 0.001
	var z_n float64
	z := float64(x)
	i := 0
	for {
		i++
		z_n = z
		z = z/2 + x/2/z
		if math.Abs(z_n-z) < stop_value {
			break
		}
	}
	return z, z_n, i
}

func main() {
	fmt.Println(Sqrt(16))
}
