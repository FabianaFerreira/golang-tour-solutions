package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func NewSqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	z := INITIAL_VALUE
	for zn := nextZ(z, x); math.Abs(zn-z) > DELTA; {
		z = zn
		zn = nextZ(z, x)
	}

	return z, nil
}

func newSqrt() {
	fmt.Println(NewSqrt(2))
	fmt.Println(NewSqrt(-2))
}
