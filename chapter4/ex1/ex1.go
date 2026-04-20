package main

import (
	"fmt"
	"math/rand/v2"
)

/*
“Write a for loop that puts 100 random numbers between 0 and 100 into an int slice.”

Excerpt From
Learning Go (for Mark nefedov)
Jon Bodner
This material may be protected by copyright.
*/

func main() {
	intSlice := make([]int, 10)

	for i := 0; i < 10; i++ {
		intSlice[i] = rand.IntN(10)
	}

	fmt.Println("slice: ", intSlice)
}
