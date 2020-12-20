package main

import "fmt"

func main() {
	slice := []int{10, 15, 20, 25}
	fmt.Println("\nHere is our slice: ")
	fmt.Println("slice == ", slice)
	fmt.Println("slice[1:4]", slice[1:4])
}
