package main

import "fmt"

func main() {
	a := 100

	if a < 101 {
		fmt.Println("if condition true")
	}

	//init statement
	if b := 101; b < 102 {
		fmt.Println("if condition true")
	}

	//else with vars on init statement
	if c := 102; c < 102 {
		fmt.Println("if condition true")
	} else {
		fmt.Println("if condition false, value: ", c)
	}
}
