package main

import "fmt"

func main() {
	letters := []string{"a", "b", "c", "d"}

	slice := letters[:]

	fmt.Printf("cap(letters): %d    cap(slice): %d\n", cap(letters), cap(slice))

	slice = append(slice, "e")

	fmt.Printf("cap(letters): %d    cap(slice): %d\n", cap(letters), cap(slice))
	fmt.Println("letters:", letters)
	fmt.Println("slice:", slice)
}
