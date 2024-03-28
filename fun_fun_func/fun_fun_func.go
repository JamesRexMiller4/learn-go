package main

import "fmt"

func main() {
	fmt.Println(addTwoNumbers(2, 4))
	fmt.Println(divRem(4, 2))

	div, _ := divRem(4, 2)
	_, rem := divRem(4, 2)
	fmt.Println(div)
	fmt.Println(rem)

	a := 1
	arr := [2]int{2, 4}
	s := "Hey"
	doubleFail(a, arr, s)
	fmt.Println(a, arr, s)
}

func addTwoNumbers(a int, b int) int {
	return a + b
}

func divRem(a int, b int) (int, int) {
	return a / b, a % b
}

func doubleFail(a int, arr [2]int, s string) {
	a = a * 2
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * 2
	}
	s = s + s
	fmt.Println(a, arr, s)
}
