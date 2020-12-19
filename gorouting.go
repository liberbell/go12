package main

import "fmt"

func f(msg int) {
	for i := 0; i < 10; i++ {
		fmt.Println(msg, ":", i)
	}
}
