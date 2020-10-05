package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("random number: ", rand.Intn(10))
	fmt.Println("pi: ", math.Pi)
}