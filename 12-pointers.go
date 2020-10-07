package main

import "fmt"

func main() {
	i := 50
	fmt.Println("i:", i, "&i:", &i)


	p := &i
	fmt.Println("p:", p, "&p", &p)

	fmt.Println("*p: ", *p)


	*p = 52
	fmt.Println("*p: ", *p)
	fmt.Println("i: ", i)

}
