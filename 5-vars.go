package main

import "fmt"

var c, python, java bool
var j = 2

func main() {
	var i int
	fmt.Println(i, c , python, java)

	var a, b, d = true, false, "vars with omitted type"
	fmt.Println(i, j, a, b, d)

	x := "short var declaraion"
	fmt.Println(x)
}
