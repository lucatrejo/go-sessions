package main

import "fmt"

func main() {
	a := 0
	for i := 0; i < 10; i++ {
		a++
	}
	fmt.Println(a)

	//optional init and post statements "WHILE"
	b := 0
	for b < 100 {
		b++
	}
	fmt.Println(b)
}
