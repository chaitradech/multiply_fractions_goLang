package main

import (
	"fmt"
	. "arithmetic/multiply_fractions"
)

func main() {
	a := Fraction{Numerator: 6, Denominator: 10}
	b := Fraction{Numerator: 4, Denominator: 10}
	fmt.Println(MultiplyFractions(a, b))
}
