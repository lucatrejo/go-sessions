package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	//has init statement
	//no need brakes

	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}

	//switch without expresion
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Print("Good Morning")
	case t.Hour() < 17:
		fmt.Print("Good Afternoon")
	default:
		fmt.Print("Good Evening")
	}
}
