package main

import (
	"fmt"
	"reflect"
)

func main() {
	var v interface{}

	v = 123
	printType(v)

	v = "hello"
	printType(v)

	v = true
	printType(v)

	c := make(chan int)
	v = c
	printType(v)
}

func printType(v interface{}) {
	fmt.Printf("Type of %v is %v\n", v, reflect.TypeOf(v))
}
