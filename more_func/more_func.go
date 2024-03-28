package main

import "fmt"

// func addOne(a int) int {
// 	return a + 1
// }

// func addTwo(a int) int {
// 	return a + 2
// }

//	func printOperation(a int, f func(int) int) {
//		fmt.Println(f(a))
//	}

func makeAdder(b int) func(int) int {
	return func(a int) int {
		return a + b
	}
}

func makeDoubler(f func(int) int) func(int) int {
	return func(a int) int {
		b := f(a)
		return b * 2
	}
}

func main() {
	// myAddOne := func(a int) int {
	// 	return a + 1
	// }
	// fmt.Println(myAddOne(1))

	// printOperation(1, addOne)
	// printOperation(1, addTwo)

	// b := 2
	// myAddOne := func(a int) {
	// 	b = a + b
	// }

	// myAddOne(1)
	// fmt.Println(b)

	addOne := makeAdder(1)
	doubleAddOne := makeDoubler(addOne)

	fmt.Println(addOne(1))
	fmt.Println(doubleAddOne(1))
}
