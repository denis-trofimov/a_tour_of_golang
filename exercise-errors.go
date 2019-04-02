package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	const stop_value = 0.000001
	var z_n float64
	z := float64(x)
	for {
		z_n = z
		z = z/2 + x/2/z
		if math.Abs(z_n-z) < stop_value {
			break
		}
	}
	return z, nil
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))

	s, err := Sqrt(2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(s)
	}
	s, err = Sqrt(-2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(s)
	}
}
