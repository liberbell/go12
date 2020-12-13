package main

import "fmt"

func main() {
	slice := []int{7, 8, 9, 10, 15}
	fmt.Println("Slice == ", slice)

	fmt.Println("Slice[1:4]", slice[1:4])

	fmt.Println("Slice[:3]", slice[:3])
}
