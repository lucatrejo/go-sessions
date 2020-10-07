package main

import(
	"fmt"
)

func add(x int, y int) int {
	return x + y
}

func addResumed(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println("function: ", add(10, 10))
	fmt.Println("function2: ", addResumed(10, 100))

	a, b := swap("world", "hello")
	fmt.Println("function multiple result: ", a, b)

	x, y := split(100)
	fmt.Println("function named result: ", x, y)
}