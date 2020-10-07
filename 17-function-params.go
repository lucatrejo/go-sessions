package main

import (
	"fmt"
)

func firstFunc(secondFunc func(string, string) string) string {
	return "First Function: " + secondFunc("World", "Hello")
}

func main() {
	secondFunc := func(first, second string) string {
		return second + " " + first + ";"
	}

	fmt.Println(secondFunc("World", "Hello"))
	fmt.Println(firstFunc(secondFunc))
}