package main

import (
	"fmt"
)

type Operational interface {
	Multiply() int
}


type MyOpStruct struct {
	A int
	B int
	Result int
}

//This method means type MyOpStruct implements the interface Operational, don't need to explicitly declare it
func (x MyOpStruct) Multiply() int {
	result := x.A * x.B
	x.Result = result
	return result
}


type MyOpInt int
func (x MyOpInt) Multiply() int {
	result := x * x
	x = result
	return int(result)
}

func main() {
	var op Operational

	myStruct := MyOpStruct{A: 5, B: 5}
	myInt := MyOpInt(10)

	op = myStruct
	fmt.Println(op.Multiply())
	describe(op)

	op = myInt
	fmt.Println(op.Multiply())
	describe(op)
}

//An interface value holds a value of a specific underlying concrete type.
func describe(i Operational) {
	fmt.Printf("(%v, %T)\n", i, i)
}