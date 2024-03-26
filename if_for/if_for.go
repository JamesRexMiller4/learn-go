package main

import "fmt"

func main() {
	a := 5
	if b := 7; a < b {
		fmt.Println(a, "is less than", b)
	} else {
		fmt.Println(a, "is greater than", b)
	}

	// Tradish
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// While
	i := 0
	for i < 6 {
		println(i)
		i++
	}

	// Eternity
	for {
		fmt.Println(i)
		if i == 10 {
			break
		}
		i++
	}
}
