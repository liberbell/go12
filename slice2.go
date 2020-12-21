package main

import "fmt"

func main() {
	slice := []int{10, 15, 20, 25}
	fmt.Println("\nHere is our slice: ")
	fmt.Println("slice == ", slice)
	fmt.Println("slice[1:4]", slice[1:4])
	fmt.Println("slice[:3]", slice[:3])
	fmt.Println("slice[2:]", slice[2:])
	fmt.Println("len slice ==", len(slice))
	fmt.Println("cap slice ==", cap(slice))

	for i, v := range slice {
		slice[i] = v - 5
	}
	fmt.Println("\nThe new values in our slice:")
	fmt.Println("slice:", slice)

	fmt.Println("\nNow we'll append 2 values to slice (what happen?)")
	slice = append(slice, 10, 20)
	fmt.Println("slice:", slice)

	fmt.Println("\nNow we'll append 8 more values to slice (now what happen)")
	slice = resize(slice)
	fmt.Println("slice:", slice)
}

func resize(slice []int) []int {
	for i := 0; i < 8; i++ {
		slice = append(slice, 1)
	}
	return slice
}
