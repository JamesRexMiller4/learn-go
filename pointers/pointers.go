package main

import "fmt"

func main() {
	a := 4
	// b := &a
	// c := a

	a = 20

	// fmt.Println(a, b, *b, c)

	// *b = 30

	// fmt.Println(a, b, *b, c)

	// c = 40

	// fmt.Println(a, b, *b, c)

	// var d *int
	// e := new(int)

	// fmt.Println(d)
	// fmt.Println(*d)
	// fmt.Println(e, *e)
	fmt.Println(a)
	setTo10(&a)
	fmt.Println(a)

}

func setTo10(a *int) {
	*a = 10
}
