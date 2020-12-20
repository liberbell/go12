package main

import "fmt"

func main() {
	c := make(chan int, 5)
	c <- 1
	c <- 3
	c <- 4
	c <- 8
	c <- 10
	c <- 13
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}
