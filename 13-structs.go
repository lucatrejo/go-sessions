package main

import "fmt"

type Vertex struct {
	X int
	Y int
	z int
}

func main() {
	v := Vertex{1,2, 3}

	fmt.Println(v)

	fmt.Println(v.X)
	fmt.Println(v.Y)
	fmt.Println(v.z)


	v.X = 10
	fmt.Println(v)
	fmt.Println(v.X)

	v1 := Vertex{Y: 10}
	v2 := Vertex{}
	fmt.Println(v1, v2)
}
