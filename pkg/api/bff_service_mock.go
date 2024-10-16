// Code generated by mockery v2.45.1. DO NOT EDIT.

//go:build !compile

package api

import (
	core "github.com/ksysoev/deriv-api-bff/pkg/core"
	mock "github.com/stretchr/testify/mock"

	wasabi "github.com/ksysoev/wasabi"
)

// MockBFFService is an autogenerated mock type for the BFFService type
type MockBFFService struct {
	mock.Mock
}

type MockBFFService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockBFFService) EXPECT() *MockBFFService_Expecter {
	return &MockBFFService_Expecter{mock: &_m.Mock}
}

// PassThrough provides a mock function with given fields: clientConn, req
func (_m *MockBFFService) PassThrough(clientConn wasabi.Connection, req *core.Request) error {
	ret := _m.Called(clientConn, req)

	if len(ret) == 0 {
		panic("no return value specified for PassThrough")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(wasabi.Connection, *core.Request) error); ok {
		r0 = rf(clientConn, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockBFFService_PassThrough_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PassThrough'
type MockBFFService_PassThrough_Call struct {
	*mock.Call
}

// PassThrough is a helper method to define mock.On call
//   - clientConn wasabi.Connection
//   - req *core.Request
func (_e *MockBFFService_Expecter) PassThrough(clientConn interface{}, req interface{}) *MockBFFService_PassThrough_Call {
	return &MockBFFService_PassThrough_Call{Call: _e.mock.On("PassThrough", clientConn, req)}
}

func (_c *MockBFFService_PassThrough_Call) Run(run func(clientConn wasabi.Connection, req *core.Request)) *MockBFFService_PassThrough_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(wasabi.Connection), args[1].(*core.Request))
	})
	return _c
}

func (_c *MockBFFService_PassThrough_Call) Return(_a0 error) *MockBFFService_PassThrough_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBFFService_PassThrough_Call) RunAndReturn(run func(wasabi.Connection, *core.Request) error) *MockBFFService_PassThrough_Call {
	_c.Call.Return(run)
	return _c
}

// ProcessRequest provides a mock function with given fields: clientConn, req
func (_m *MockBFFService) ProcessRequest(clientConn wasabi.Connection, req *core.Request) error {
	ret := _m.Called(clientConn, req)

	if len(ret) == 0 {
		panic("no return value specified for ProcessRequest")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(wasabi.Connection, *core.Request) error); ok {
		r0 = rf(clientConn, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockBFFService_ProcessRequest_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ProcessRequest'
type MockBFFService_ProcessRequest_Call struct {
	*mock.Call
}

// ProcessRequest is a helper method to define mock.On call
//   - clientConn wasabi.Connection
//   - req *core.Request
func (_e *MockBFFService_Expecter) ProcessRequest(clientConn interface{}, req interface{}) *MockBFFService_ProcessRequest_Call {
	return &MockBFFService_ProcessRequest_Call{Call: _e.mock.On("ProcessRequest", clientConn, req)}
}

func (_c *MockBFFService_ProcessRequest_Call) Run(run func(clientConn wasabi.Connection, req *core.Request)) *MockBFFService_ProcessRequest_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(wasabi.Connection), args[1].(*core.Request))
	})
	return _c
}

func (_c *MockBFFService_ProcessRequest_Call) Return(_a0 error) *MockBFFService_ProcessRequest_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBFFService_ProcessRequest_Call) RunAndReturn(run func(wasabi.Connection, *core.Request) error) *MockBFFService_ProcessRequest_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockBFFService creates a new instance of MockBFFService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockBFFService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockBFFService {
	mock := &MockBFFService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
