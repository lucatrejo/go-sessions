package main

import "fmt"

func main() {
	var a[2] string
	a[0] = "hello"
	a[1] = "world"

	fmt.Println(a)


	fmt.Println(a[0], a[1])

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}
