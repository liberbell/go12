package main

import "fmt"

type Rect struct {
	Width, Height int
}

var m = map[string]Rect{
	"Rect1": Rect{1, 2},
	"Rect2": Rect{4, 6},
}

func main() {
	fmt.Println(m)
}
