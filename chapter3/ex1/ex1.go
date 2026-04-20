package main

import "fmt"

/*
“Write a program that defines a variable named greetings of type slice of strings with the following values:
"Hello", "Hola", "नमस्कार", "こんにちは", and "Привіт". Create a subslice containing the first two values;
a second subslice with the second, third, and fourth values; and a third subslice with the fourth and fifth values.
Print out all four slices.”

Excerpt From
Learning Go (for Mark nefedov)
Jon Bodner
This material may be protected by copyright.
*/
func main() {
	greetings := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}

	firstTwo := make([]string, 2)
	copy(firstTwo, greetings[:2])

	secondThirtFourth := make([]string, 3)
	copy(secondThirtFourth, greetings[1:4])

	fourthFifth := make([]string, 2)
	copy(fourthFifth, greetings[3:])

	fmt.Println("First two: \t\t", firstTwo)
	fmt.Println("Second, Third, Fourth: \t", secondThirtFourth)
	fmt.Println("Fourth, Fifth: \t\t", fourthFifth)
}
