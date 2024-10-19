// Code generated by mockery v2.45.1. DO NOT EDIT.

//go:build !compile

package handler

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockWaitComposer is an autogenerated mock type for the WaitComposer type
type MockWaitComposer struct {
	mock.Mock
}

type MockWaitComposer_Expecter struct {
	mock *mock.Mock
}

func (_m *MockWaitComposer) EXPECT() *MockWaitComposer_Expecter {
	return &MockWaitComposer_Expecter{mock: &_m.Mock}
}

// Compose provides a mock function with given fields:
func (_m *MockWaitComposer) Compose() (map[string]interface{}, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Compose")
	}

	var r0 map[string]interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func() (map[string]interface{}, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() map[string]interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockWaitComposer_Compose_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Compose'
type MockWaitComposer_Compose_Call struct {
	*mock.Call
}

// Compose is a helper method to define mock.On call
func (_e *MockWaitComposer_Expecter) Compose() *MockWaitComposer_Compose_Call {
	return &MockWaitComposer_Compose_Call{Call: _e.mock.On("Compose")}
}

func (_c *MockWaitComposer_Compose_Call) Run(run func()) *MockWaitComposer_Compose_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockWaitComposer_Compose_Call) Return(_a0 map[string]interface{}, _a1 error) *MockWaitComposer_Compose_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockWaitComposer_Compose_Call) RunAndReturn(run func() (map[string]interface{}, error)) *MockWaitComposer_Compose_Call {
	_c.Call.Return(run)
	return _c
}

// Wait provides a mock function with given fields: ctx, name, parser, respChan
func (_m *MockWaitComposer) Wait(ctx context.Context, name string, parser Parser, respChan <-chan []byte) {
	_m.Called(ctx, name, parser, respChan)
}

// MockWaitComposer_Wait_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Wait'
type MockWaitComposer_Wait_Call struct {
	*mock.Call
}

// Wait is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - parser Parser
//   - respChan <-chan []byte
func (_e *MockWaitComposer_Expecter) Wait(ctx interface{}, name interface{}, parser interface{}, respChan interface{}) *MockWaitComposer_Wait_Call {
	return &MockWaitComposer_Wait_Call{Call: _e.mock.On("Wait", ctx, name, parser, respChan)}
}

func (_c *MockWaitComposer_Wait_Call) Run(run func(ctx context.Context, name string, parser Parser, respChan <-chan []byte)) *MockWaitComposer_Wait_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(Parser), args[3].(<-chan []byte))
	})
	return _c
}

func (_c *MockWaitComposer_Wait_Call) Return() *MockWaitComposer_Wait_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockWaitComposer_Wait_Call) RunAndReturn(run func(context.Context, string, Parser, <-chan []byte)) *MockWaitComposer_Wait_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockWaitComposer creates a new instance of MockWaitComposer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockWaitComposer(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockWaitComposer {
	mock := &MockWaitComposer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
