package main

import (
	"fmt"
)

type MyStruct struct {
	X int
	Y int
	Z int
}

func MultiplyFunc(myStruct MyStruct) int {
	result := myStruct.X * myStruct.Y
	myStruct.Z = result
	return result
}

//Method is just a func with a receiver argument (v MyStruct)
func (myStruct MyStruct) Multiply() int {
	result := myStruct.X * myStruct.Y
	myStruct.Z = result
	return result
}

func main() {
	//try with pointers!
	myStruct := MyStruct{X: 5, Y: 5}

	fmt.Println(MultiplyFunc(myStruct))
	fmt.Println(myStruct)

	fmt.Println(myStruct.Multiply())
	fmt.Println(myStruct)
}