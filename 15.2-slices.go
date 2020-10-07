package main

import (
	"fmt"
)

func main() {
	//slice of slice
	sliceParent := [][]int{
		[]int{1, 1, 1, 1},
		[]int{2, 2, 2, 2},
		[]int{3, 3, 3, 3},
	}
	printSlices("Slice of Slice: ", sliceParent)

	a := make([]int, 3, 6) //type, length, capacity
	printSliceInt("slice with make: ", a)
	a = append(a, 1)
	a = append(a, 2)
	a = append(a, 3)	//if slice is initialized with capacity defined
	a = append(a, 4)	//when the cap is exceeded the cap is duplicated
	printSliceInt("slice after append: ", a)

	var b []int
	b = append(b, 1,2,3,4,5,6,7)
	printSliceInt("slice", b) //the cap grows as needed

	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	//omit the index or value with _ or remove the second parameter
	for i, v := range pow { //i: index - v: copy of element in that index
		if i == 7 {
			v = 999
			//pow[7] = 999
		}
		fmt.Printf("2**%d = %d\n", i, v)
	}

	printSliceInt("Slice after iteration: ", pow)
}

func printSliceInt(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func printSlices(s string, x [][]int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}