package main

import (
	"errors"
	"fmt"
	"strconv"
)

/*
“The simple calculator program doesn’t handle one error case: division by zero.
Change the function signature for the math operations to return both an int and an error.
In the div function, if the divisor is 0, return errors.New("division by zero") for the error.
In all other cases, return nil. Adjust the main function to check for this error.”

Excerpt From
Learning Go (for Mark nefedov)
Jon Bodner
This material may be protected by copyright.
*/

var opMap = map[string]func(int, int) (int, error){
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
}

func main() {
	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "0"},
		{"2", "%", "3"},
		{"two", "+", "three"},
		{"5"},
	}

	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Println("invalid expression:", expression)
			continue
		}

		p1, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Println(err)
			continue
		}

		op := expression[1]

		opFunc, ok := opMap[op]
		if !ok {
			fmt.Println("unsupported operator:", op)
			continue
		}
		p2, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Println(err)
			continue
		}
		result, err := opFunc(p1, p2)
		if err != nil {
			fmt.Printf("An error has occured: %v\n", err)
			return
		}
		fmt.Println(result)
	}

}

func add(a, b int) (int, error) { return a + b, nil }
func sub(a, b int) (int, error) { return a - b, nil }
func mul(a, b int) (int, error) { return a * b, nil }
func div(a, b int) (int, error) {

	if b == 0 {
		return 0, errors.New("division by zero")
	}

	return a / b, nil
}
