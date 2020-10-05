package main

import "fmt"

var (
	Boolean bool
	String string
	Int int
	Float float32
	Complex complex64
)

func main() {
	fmt.Printf("Type: %T - DefaultValue: %v\n", Boolean, Boolean)
	fmt.Printf("Type: %T - DefaultValue: %v\n", String, String)
	fmt.Printf("Type: %T - DefaultValue: %v\n", Int, Int)
	fmt.Printf("Type: %T - DefaultValue: %v\n", Float, Float)
	fmt.Printf("Type: %T - DefaultValue: %v\n", Complex, Complex)

	//Conversions
	Int := 1
	Float = float32(Int)
	fmt.Println("Float from Int: ", Float)

	//Inference
	var i int
	j := i
	fmt.Printf("J is type: %T\n", j)

	h := 42	//int
	f := 42.123 //float
	g := 42.123 + 0.5i //complex

	fmt.Printf("Type: %T - DefaultValue: %v\n", h, h)
	fmt.Printf("Type: %T - DefaultValue: %v\n", f, f)
	fmt.Printf("Type: %T - DefaultValue: %v\n", g, g)

}
