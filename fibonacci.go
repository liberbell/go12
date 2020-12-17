package main

import "fmt"

func fibonacci() func() {
	x := 0
	y := 1
	fmt.Println(x)

	return func(x) int {
		x, y = y, x+y
		return x
	}
}
