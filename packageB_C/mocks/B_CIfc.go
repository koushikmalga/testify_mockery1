// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// B_CIfc is an autogenerated mock type for the B_CIfc type
type B_CIfc struct {
	mock.Mock
}

// Sub_modified provides a mock function with given fields: _a0, _a1
func (_m *B_CIfc) Sub_modified(_a0 int, _a1 int) int {
	ret := _m.Called(_a0, _a1)

	var r0 int
	if rf, ok := ret.Get(0).(func(int, int) int); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// Sub_modified1 provides a mock function with given fields: _a0, _a1
func (_m *B_CIfc) Sub_modified1(_a0 int, _a1 int) int {
	ret := _m.Called(_a0, _a1)

	var r0 int
	if rf, ok := ret.Get(0).(func(int, int) int); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// NewB_CIfc creates a new instance of B_CIfc. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewB_CIfc(t interface {
	mock.TestingT
	Cleanup(func())
}) *B_CIfc {
	mock := &B_CIfc{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
