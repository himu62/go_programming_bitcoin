package main

import "github.com/himu62/go_programming_bitcoin/01_finite_fields/ecc"

func main() {
	a, _ := ecc.NewFieldElement(7, 13)
	b, _ := ecc.NewFieldElement(6, 13)
	println(a.Equals(b))  // => false
	println(a.Equals(a))  // => true

	// practice 1
	println(a.Unequals(b))  // => true
	println(a.Unequals(a))  // => false
}
