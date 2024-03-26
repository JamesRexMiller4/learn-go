package main

import "fmt"

func main() {
	s := "Hello World!"
	s2 := s[4]
	s3 := s[7]
	s4 := s[0:4]
	s5 := s[5:9]
	fmt.Println(s, s2, string(s3), s4, s5)
}
