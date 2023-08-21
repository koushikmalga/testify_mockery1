package packageB

import (
	"testify_mockery1/packageB_C"
)

// Interface with 3 moethods where one has external dependency
type BIfc interface {
	Sum(int, int) int
	UseC(int, int) int
	Prod(int, int) int
}

// A struct with packageB_C interface
type B struct {
	c packageB_C.B_CIfc
}

func NewB(c packageB_C.B_CIfc) *B {
	return &B{c: c}
}

func (b *B) Sum(ele1, ele2 int) int {

	return ele1 + ele2
}

// function calling packageB_C
func (b *B) UseC(ele1, ele2 int) int {

	return b.c.Sub_modified1(ele1, ele2)
}

func (b *B) Prod(ele1, ele2 int) int {
	return ele1 * ele2
}
