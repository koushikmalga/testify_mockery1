package packageB_C

// Considering C as a external dependency and need to mock it. So instead of directly calling C package from B.
// Creating a package B_C as the bridge between B and C and generating mocks for B_C

import (
	"testify_mockery1/packageC"
)

// By using mockery generating packageB_C Interfaces mock

//go:generate mockery --all --output=mocks
// The above line generates mocks of all the interfaces as we mentioned --all and saves it in the --output mentioned

// Creating an interface which inreturn calls packageC
type B_CIfc interface {
	Sub_modified(int, int) int
	Sub_modified1(int, int) int
}
type B_C struct {
	c packageC.CIfc
}

func (a *B_C) Sub_modified(ele1, ele2 int) int {
	res := a.c.Sub(ele1, ele2)

	return res
}

func (a *B_C) Sub_modified1(ele1, ele2 int) int {
	res := a.c.Sub(ele1, ele2)

	return res
}
