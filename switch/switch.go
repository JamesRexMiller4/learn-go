package main

import (
	"fmt"
	"os"
)

func main() {
	word := os.Args[1]

	switch l := len(word); {
	case word == "hello":
		fmt.Println("Hello yourself")
		fallthrough
	case word == "hi":
		fmt.Println("How dare you")
		fallthrough
	case l < 4 && l > 6:
		fmt.Println("Oh...wasn't expecting that")
	case word == "goodbye":
		fmt.Println("See ya!")
	default:
		fmt.Println("I don't know what you said")
	}
}
