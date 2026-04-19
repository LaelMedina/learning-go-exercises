package main

import "fmt"

func main() {
	odds := []int{1, 3, 5, 7}

	oddsCopy := make([]int, 4)

	result := copy(oddsCopy, odds)

	fmt.Println(oddsCopy, result)
}
