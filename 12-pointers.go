package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i
	fmt.Println("pointer to i: ", p)
	fmt.Println("read through the pointer: ", *p)
	*p = 52

	fmt.Println("pointer value updated: ", *p)
	fmt.Println("original var value: ", i)
	fmt.Println("memory dir on p", &p)
	fmt.Println("memory dir on i", &i)
	fmt.Println("value of p", p)

	p = &j
	fmt.Println("assign new value")
	fmt.Println("pointer value updated:", *p)
	*p = *p / 37
	fmt.Println("value of j:", j)

}
