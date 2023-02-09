// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	test "github.com/vektra/mockery/v2/pkg/fixtures"
)

// UsesOtherPkgIface is an autogenerated mock type for the UsesOtherPkgIface type
type UsesOtherPkgIface struct {
	mock.Mock
}

// DoSomethingElse provides a mock function with given fields: obj
func (_m *UsesOtherPkgIface) DoSomethingElse(obj test.Sibling) {
	_m.Called(obj)
}

type mockConstructorTestingTNewUsesOtherPkgIface interface {
	mock.TestingT
	Cleanup(func())
}

// NewUsesOtherPkgIface creates a new instance of UsesOtherPkgIface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUsesOtherPkgIface(t mockConstructorTestingTNewUsesOtherPkgIface) *UsesOtherPkgIface {
	mock := &UsesOtherPkgIface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}