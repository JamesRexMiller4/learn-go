package main

import (
	"Users/jamesmiller/Desktop/learn-go/leftpad"
	"fmt"
)

func main() {
	fmt.Println(leftpad.Format("hello", 15))
	fmt.Println(leftpad.Format("goodbye", 15))
	fmt.Println(leftpad.Format("Oh my", 15))
	fmt.Println(leftpad.Format("Internationalization", 15))
	fmt.Println(leftpad.FormatRune("hello", 15, '🥴'))
	fmt.Println(leftpad.FormatRune("goodbye", 15, '🥴'))
	fmt.Println(leftpad.FormatRune("Oh my", 15, '🥴'))
}
