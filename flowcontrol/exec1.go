package main

import (
	"fmt"
	"math"
)

const DELTA = 0.0000001
const INITIAL_VALUE = 100.0

func nextZ(z, x float64) float64 { return z - (z*z - x) / (2 * z) }

func Sqrt(x float64) (z float64) {
	z = INITIAL_VALUE

	for zn := nextZ(z, x); math.Abs(zn - z) > DELTA; {
		z = zn
		zn = nextZ(z, x)
	}

	return
}

func main() {
	fmt.Println(Sqrt(2))
}