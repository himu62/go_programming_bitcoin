package ecc

import "fmt"

type FieldElement struct {
	num int
	prime int
}

func NewFieldElement(num int, prime int) (*FieldElement, error) {
	if num >= prime || num < 0 {
		return nil, fmt.Errorf("num %d not in field range 0 to %d", num, prime - 1)
	}
	return &FieldElement{
		num: num,
		prime: prime,
	}, nil
}

func (fe *FieldElement) String() string {
	return fmt.Sprintf("FieldElement_%d(%d)", fe.prime, fe.num)
}

// display
func (fe *FieldElement) Num() int {
	return fe.num
}

func (fe *FieldElement) Equals(other *FieldElement) bool {
	return fe.num == other.num && fe.prime == other.prime
}

// practice 1
func (fe *FieldElement) Unequals(other *FieldElement) bool {
	return fe.num != other.num || fe.prime != other.prime
}

// practice 2
func (fe *FieldElement) Add(other *FieldElement) (*FieldElement, error) {
	if fe.prime != other.prime {
		return &FieldElement{}, fmt.Errorf("unequal primes %d != %d", fe.prime, other.prime)
	}
	mod := (fe.num + other.num) % fe.prime
	if mod < 0 {
		mod = fe.prime + mod
	}
	return &FieldElement{
		num: mod,
		prime: fe.prime,
	}, nil
}

// practice 3
func (fe *FieldElement) Sub(other *FieldElement) (*FieldElement, error) {
	if fe.prime != other.prime {
		return &FieldElement{}, fmt.Errorf("unequal primes %d != %d", fe.prime, other.prime)
	}
	mod := (fe.num - other.num) % fe.prime
	if mod < 0 {
		mod = fe.prime + mod
	}
	return &FieldElement{
		num: mod,
		prime: fe.prime,
	}, nil
}
