package main

import "fmt"

func main() {
	fmt.Println("\nLast in ---> First Out\n")
	for x := 0; x < 5; x++ {
		defer fmt.Println(x, "popped.")
	}
}
