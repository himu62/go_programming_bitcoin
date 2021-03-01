package main

import (
	"github.com/himu62/go_programming_bitcoin/01_finite_fields/ecc"
)

func N(num int) *ecc.FieldElement {
	el, _ := ecc.NewFieldElement(num, 57)
	return el
}

func main() {
	// practice 2
	p1, _ := N(44).Add(N(33))
	println(p1.Num())  // => 20

	p2, _ := N(9).Sub(N(29))
	println(p2.Num())  // => 37

	p3, _ := N(17).Add(N(42))
	p3, _ = p3.Add(N(49))
	println(p3.Num())  // => 51

	p4, _ := N(52).Sub(N(30))
	p4, _ = p4.Sub(N(38))
	println(p4.Num())  // => 41
}
