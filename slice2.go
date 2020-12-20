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
}
