package main

import "fmt"

func main() {
	w := make(map[string]int)
	w["Answer"] = 10

	fmt.Println("The value: ", w["Answer"])
}
