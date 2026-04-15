package main

import "fmt"

/*
“Write a program that declares an integer variable called i with the value 20. Assign i to a floating-point variable named f. Print out i and f.”

Excerpt From
Learning Go (for Mark nefedov)
Jon Bodner
This material may be protected by copyright.
*/

func main() {

	var i int = 20
	var f float32

	f = float32(i)

	fmt.Printf("i: %d  f: %.2f\n", i, f)
}
