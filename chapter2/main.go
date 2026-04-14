package main

import "fmt"

func main() {

	var age uint8

	fmt.Print("Enter your age: ")

	_, err := fmt.Scanf("%d", &age)

	if err != nil {
		fmt.Println("Error reading input: ", err)
		return
	}

	fmt.Printf("\nThe age is: %d\n", age)
}
