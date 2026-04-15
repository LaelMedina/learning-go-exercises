package main

import "fmt"

/*
“Write a program with three variables, one named b of type byte, one named smallI of type int32, and one named bigI of type uint64.
Assign each variable the maximum legal value for its type; then add 1 to each variable. Print out their values.”

Excerpt From
Learning Go (for Mark nefedov)
Jon Bodner
This material may be protected by copyright.
*/

func main() {
	var b byte = 255
	var smallI int32 = 2147483647
	var bigI int64 = 9223372036854775807

	b += 1
	smallI += 1
	bigI += 1

	// Adding 1 to the maximun value creates an overflow, in this case, the min legal value will be assined
	//b: 0
	// smallI: -2147483648
	// bigI: -9223372036854775808

	fmt.Printf("b: %v\n", b)
	fmt.Printf("smallI: %v\n", smallI)
	fmt.Printf("bigI: %v\n", bigI)
}
