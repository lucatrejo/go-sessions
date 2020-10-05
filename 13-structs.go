package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1,2}

	fmt.Println(v)
	fmt.Println(v.X)
	fmt.Println(v.Y)

	v.X = 10
	fmt.Println(v)
	fmt.Println(v.X)

	//pointers
	p := &v
	p.X = 100
	fmt.Println(v)

	v1 := Vertex{Y: 10}
	v2 := Vertex{}
	v3 := &Vertex{10,20}
	fmt.Println(v1, v2, v3)

}
