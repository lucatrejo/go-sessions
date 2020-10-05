package main

import "fmt"

func main() {
	fmt.Println("First")

	defer fmt.Println("First IN - Last OUT")
	defer fmt.Println("Last IN - First OUT")

	fmt.Println("Second")
}
