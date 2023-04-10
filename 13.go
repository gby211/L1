package main

import "fmt"

func main() {
	a := 5
	b := 8

	fmt.Printf(" a = %d, b = %d\n", a, b)

	a = a ^ b
	b = a ^ b
	a = a ^ b

	fmt.Printf("a = %d, b = %d\n", a, b)
}

//func main() {
//	a := 5
//	b := 8
//
//	a,b = b, a
//	fmt.Printf(" a = %d, b = %d\n", a, b)
//}
