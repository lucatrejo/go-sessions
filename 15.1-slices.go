package main

import "fmt"

func main() {
	names := [4]string {
		"John", "Paul", "George", "Ringo",
	}

	fmt.Println(names)


	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)


	b[0] = "XXX" //slices are references to arrays
	fmt.Println(a, b)
	fmt.Println(names)


	//literals
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

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
	fmt.Println(n[:])

	//length and capacity
	printSlice(n)

	// Slice the slice to give it zero length.
	n = n[:0]
	printSlice(n)

	// Extend its length.
	n = n[:4]
	printSlice(n)

	// Null slice
	var nullSlice []int
	fmt.Print("null slice: ")
	printSlice(nullSlice)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
