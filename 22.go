package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(7000000000000000000)
	b := big.NewInt(5000000000000000000)

	// сложение
	sum := big.NewInt(0)
	sum.Add(a, b)
	fmt.Println(sum)

	// вычитание
	diff := big.NewInt(0)
	diff.Sub(b, a)
	fmt.Println(diff)

	// умножение
	prod := big.NewInt(0)
	prod.Mul(a, b)
	fmt.Println(prod)

	// деление
	quotient := big.NewInt(0)
	quotient.Div(b, a)
	fmt.Println(quotient)
}
