package main

import "fmt"

func main() {
	primes := [6]int {2, 3, 5, 7, 11, 13}

	//include first element and excludes last one
	var s []int = primes[0:3]
	fmt.Println(s)

	names := [4]string {
		"John", "Paul", "George", "Ringo",
	}

	fmt.Println(names)

	//slices are references to arrays
	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)
	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	//literals
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)
	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	m := []struct{
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(m)

	//default limits on slices low = 0 & high = array length
	n := []int{2, 3, 5, 7, 11, 13}
	n1 := n[:]
	fmt.Println(n1)
	n2 := n[1:]
	fmt.Println(n2)
	n3 := n[:3]
	fmt.Println(n3)

	//length and capacity
	s3 := []int{2, 3, 5, 7, 11, 13}
	printSlice(s3)

	// Slice the slice to give it zero length.
	s3 = s3[:0]
	printSlice(s3)

	// Extend its length.
	s3 = s3[:4]
	printSlice(s3)

	// Drop its first two values.
	s3 = s3[2:]
	printSlice(s3)

	// Null slice
	var s4 []int
	fmt.Print("null slice: ")
	printSlice(s4)


}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
