package main

import (
	"fmt"
)

func main() {
	var emptyI interface{}
	describeInterfaceEmpty(emptyI)
	emptyI = 10
	describeInterfaceEmpty(emptyI)
	emptyI = "Hello"
	describeInterfaceEmpty(emptyI)

	//Type Assertions
	a, ok := emptyI.(string)
	fmt.Println(a, ok)

	b, ok := emptyI.(int)
	fmt.Println(b, ok)

	c := emptyI.(string)
	fmt.Println(c)

	//d := emptyI.(int)
	//fmt.Println(d)

	//Switch Types
	switch v := emptyI.(type) {
	case int:
		fmt.Println("int ", v)
	case string:
		fmt.Print("string ", v)
	default:
		fmt.Print("I have no idea ", v)
	}

}

//Empty interfaces are used by code that handles values of unknown type
func describeInterfaceEmpty(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}