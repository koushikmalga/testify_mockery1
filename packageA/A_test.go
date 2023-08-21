package packageA

import (
	"testify_mockery1/packageB"
	"testify_mockery1/packageB_C/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestA(t *testing.T) {
	// Creating a mock from the one  genrated by mockery
	mockc1 := &mocks.B_CIfc{}

	// On calling Sub method returning
	mockc1.On("Sub_modified1", 5, 2).Return(3)

	// Building NewA and NewB by using mock genrated by mockery
	a := NewA(packageB.NewB(mockc1))

	var ele1, ele2 int
	ele1 = 5
	ele2 = 2
	result := a.func1(ele1, ele2)

	exp_result := 10

	// testing if two values are equal
	assert.Equal(t, exp_result, result)

	//checks whether all conditions for calling function met
	mockc1.AssertExpectations(t)

}
