package main

import "fmt"

func main() {
	slice := []int{7, 8, 9, 10, 12}
	fmt.Println("Slice  ==", slice)

	for i := 0; i < len(slice); i++ {
		fmt.Printf("Slice[%d] == %d\n", i, slice[i])
	}
}
