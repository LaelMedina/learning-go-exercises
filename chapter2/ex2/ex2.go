package main

import "fmt"

/*
“Write a program that declares a constant called value that can be assigned to both an integer and a floating-point variable.
Assign it to an integer called i and a floating-point variable called f. Print out i and f.”

Excerpt From
Learning Go (for Mark nefedov)
Jon Bodner
This material may be protected by copyright.
*/

func main() {

	const value = 10
	var i int = value
	var f float32 = value

	fmt.Printf("i: %d  f: %.2f\n", i, f)
}
