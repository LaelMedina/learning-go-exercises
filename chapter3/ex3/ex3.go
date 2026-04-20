package main

import "fmt"

/*
“Write a program that defines a struct called Employee with three fields:
firstName, lastName, and id. The first two fields are of type string, and the last field (id) is of type int.
Create three instances of this struct using whatever values you’d like.
Initialize the first one using the struct literal style without names,
the second using the struct literal style with names,
and the third with a var declaration.
Use dot notation to populate the fields in the third struct. Print out all three structs.”

Excerpt From
Learning Go (for Mark nefedov)
Jon Bodner
This material may be protected by copyright.
*/

type Employee struct {
	id        int
	firstName string
	lastName  string
}

func main() {

	empRoy := Employee{
		1,
		"Roy",
		"Mustang",
	}

	empErwin := Employee{
		id:        2,
		firstName: "Erwin",
		lastName:  "Smith",
	}

	var empThorfin Employee

	empThorfin.id = 3
	empThorfin.firstName = "Thorfin"
	empThorfin.lastName = "Thorsson"

	fmt.Println("Employee 1: ", empRoy)
	fmt.Println("Employee 2: ", empErwin)
	fmt.Println("Employee 3: ", empThorfin)

}
