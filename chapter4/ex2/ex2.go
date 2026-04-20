package main

import (
	"fmt"
	"math/rand/v2"
)

/*

“Loop over the slice you created in exercise 1. For each value in the slice, apply the following rules:
If the value is divisible by 2, print “Two!”
If the value is divisible by 3, print “Three!”
IIf the value is divisible by 2 and 3, print “Six!”. Don’t print anything else.
Otherwise, print “Never mind”.

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

	for i := 0; i < 10; i++ {
		if intSlice[i]%2 == 0 && intSlice[i]%3 == 0 {
			fmt.Printf("Six!  (%d)\n", intSlice[i])
			continue
		}

		if intSlice[i]%2 == 0 {
			fmt.Printf("Two!  (%d)\n", intSlice[i])
			continue
		}

		if intSlice[i]%3 == 0 {
			fmt.Printf("Three!  (%d)\n", intSlice[i])
			continue
		}

		fmt.Printf("Never mind\n")
	}
}
