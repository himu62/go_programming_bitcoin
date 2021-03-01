package ecc

import (
	"errors"
	"fmt"
)

type FieldElement struct {
	num int
	prime int
}

func NewFieldElement(num int, prime int) (*FieldElement, error) {
	if num >= prime || num < 0 {
		return nil, errors.New(fmt.Sprintf("Num %d not in field range 0 to %d", num, prime - 1))
	}
	return &FieldElement{
		num: num,
		prime: prime,
	}, nil
}

func (fe *FieldElement) String() string {
	return fmt.Sprintf("FieldElement_%d(%d)", fe.prime, fe.num)
}

func (fe *FieldElement) Equals(other *FieldElement) bool {
	return fe.num == other.num && fe.prime == other.prime
}

// practice 1
func (fe *FieldElement) Unequals(other *FieldElement) bool {
	return fe.num != other.num || fe.prime != other.prime
}
