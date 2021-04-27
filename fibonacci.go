package main

import "fmt"

func fibonacciSequence() func() int {
	an2, an1 := 0, 1

	return func() int {
		an := an2
		an2, an1 = an1, an+an1

		return an
	}
}

func fibonacci() {
	f := fibonacciSequence()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
