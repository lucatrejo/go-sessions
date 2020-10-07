package main

import (
	"fmt"
)

func adder() func(int) int {
	sum := 0

	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	add := adder()

	fmt.Println(add(1))
	fmt.Println(add(2))
	fmt.Println(add(10))
}
