// Code generated by mockery v2.43.1. DO NOT EDIT.

package mocks

import (
	command "testwork/internal/command"

	mock "github.com/stretchr/testify/mock"
)

// Storage is an autogenerated mock type for the Storage type
type Storage struct {
	mock.Mock
}

// Delete provides a mock function with given fields: key
func (_m *Storage) Delete(key string) {
	_m.Called(key)
}

// Get provides a mock function with given fields: key
func (_m *Storage) Get(key string) *command.StoredItem {
	ret := _m.Called(key)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *command.StoredItem
	if rf, ok := ret.Get(0).(func(string) *command.StoredItem); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*command.StoredItem)
		}
	}

	return r0
}

// GetAll provides a mock function with given fields:
func (_m *Storage) GetAll() []*command.StoredItem {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAll")
	}

	var r0 []*command.StoredItem
	if rf, ok := ret.Get(0).(func() []*command.StoredItem); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*command.StoredItem)
		}
	}

	return r0
}

// Set provides a mock function with given fields: key, value
func (_m *Storage) Set(key string, value string) {
	_m.Called(key, value)
}

// NewStorage creates a new instance of Storage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStorage(t interface {
	mock.TestingT
	Cleanup(func())
}) *Storage {
	mock := &Storage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}