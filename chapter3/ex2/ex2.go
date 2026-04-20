package main

import (
	"fmt"
)

/*
“Write a program that defines a string variable called message with the value "Hi 🧔🏻‍♂️ and 👩🏻‍🦰" and prints the fourth rune in it as a character, not a number.”

Excerpt From
Learning Go (for Mark nefedov)
Jon Bodner
This material may be protected by copyright.
*/

func main() {

	message := "Hi 🧔🏻‍♂️ and 👩🏻‍🦰"

	runeMessage := []rune(message)

	emoji := rune(runeMessage[3])

	fmt.Println("Emoji: ", string(emoji))
}
