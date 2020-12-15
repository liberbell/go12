package main

import "fmt"

type Rect struct {
	Width, Height int
}

var (
	r1 = Rect{7, 8}     // type Rect
	r2 = Rect{Width: 4} // Height is implictly 0
)

func main() {
	fmt.Println()
}
