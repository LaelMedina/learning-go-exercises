package main

import "fmt"

/*
“Start a new program. In main, declare an int variable called total.
Write a for loop that uses a variable named i to iterate from 0 (inclusive) to 10 (exclusive).
The body of the for loop should be as follows:

total := total + i
fmt.Println(total)

After the for loop, print out the value of total.  What is printed out? What is the likely bug in this code?”

Excerpt From
Learning Go (for Mark nefedov)
Jon Bodner
This material may be protected by copyright.
*/

func main() {
	var total int

	for i := 0; i < 10; i++ {
		total := total + i
		// Bug explanation: a new variable is being created inside the for loop scope,
		// the total varible outside of it it's never being incremented.
		fmt.Println("in total: ", total)
	}

	fmt.Println("total: ", total)
}
