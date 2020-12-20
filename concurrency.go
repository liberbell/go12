package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("\nMessage from func main, I'm finished.")
}

func msg() {
	for i := 0; i <= 10; i++ {
		time.Sleep(time.Microsecond * 1000)
	}
}
