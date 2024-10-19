// Code generated by mockery v2.45.1. DO NOT EDIT.

//go:build !compile

package handler

import (
	io "io"

	mock "github.com/stretchr/testify/mock"
)

// MockRenderParser is an autogenerated mock type for the RenderParser type
type MockRenderParser struct {
	mock.Mock
}

type MockRenderParser_Expecter struct {
	mock *mock.Mock
}

func (_m *MockRenderParser) EXPECT() *MockRenderParser_Expecter {
	return &MockRenderParser_Expecter{mock: &_m.Mock}
}

// DependsOn provides a mock function with given fields:
func (_m *MockRenderParser) DependsOn() []string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for DependsOn")
	}

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// MockRenderParser_DependsOn_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DependsOn'
type MockRenderParser_DependsOn_Call struct {
	*mock.Call
}

// DependsOn is a helper method to define mock.On call
func (_e *MockRenderParser_Expecter) DependsOn() *MockRenderParser_DependsOn_Call {
	return &MockRenderParser_DependsOn_Call{Call: _e.mock.On("DependsOn")}
}

func (_c *MockRenderParser_DependsOn_Call) Run(run func()) *MockRenderParser_DependsOn_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockRenderParser_DependsOn_Call) Return(_a0 []string) *MockRenderParser_DependsOn_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRenderParser_DependsOn_Call) RunAndReturn(run func() []string) *MockRenderParser_DependsOn_Call {
	_c.Call.Return(run)
	return _c
}

// Name provides a mock function with given fields:
func (_m *MockRenderParser) Name() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Name")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockRenderParser_Name_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Name'
type MockRenderParser_Name_Call struct {
	*mock.Call
}

// Name is a helper method to define mock.On call
func (_e *MockRenderParser_Expecter) Name() *MockRenderParser_Name_Call {
	return &MockRenderParser_Name_Call{Call: _e.mock.On("Name")}
}

func (_c *MockRenderParser_Name_Call) Run(run func()) *MockRenderParser_Name_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockRenderParser_Name_Call) Return(_a0 string) *MockRenderParser_Name_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRenderParser_Name_Call) RunAndReturn(run func() string) *MockRenderParser_Name_Call {
	_c.Call.Return(run)
	return _c
}

// Parse provides a mock function with given fields: data
func (_m *MockRenderParser) Parse(data []byte) (map[string]interface{}, map[string]interface{}, error) {
	ret := _m.Called(data)

	if len(ret) == 0 {
		panic("no return value specified for Parse")
	}

	var r0 map[string]interface{}
	var r1 map[string]interface{}
	var r2 error
	if rf, ok := ret.Get(0).(func([]byte) (map[string]interface{}, map[string]interface{}, error)); ok {
		return rf(data)
	}
	if rf, ok := ret.Get(0).(func([]byte) map[string]interface{}); ok {
		r0 = rf(data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	if rf, ok := ret.Get(1).(func([]byte) map[string]interface{}); ok {
		r1 = rf(data)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(map[string]interface{})
		}
	}

	if rf, ok := ret.Get(2).(func([]byte) error); ok {
		r2 = rf(data)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockRenderParser_Parse_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Parse'
type MockRenderParser_Parse_Call struct {
	*mock.Call
}

// Parse is a helper method to define mock.On call
//   - data []byte
func (_e *MockRenderParser_Expecter) Parse(data interface{}) *MockRenderParser_Parse_Call {
	return &MockRenderParser_Parse_Call{Call: _e.mock.On("Parse", data)}
}

func (_c *MockRenderParser_Parse_Call) Run(run func(data []byte)) *MockRenderParser_Parse_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte))
	})
	return _c
}

func (_c *MockRenderParser_Parse_Call) Return(_a0 map[string]interface{}, _a1 map[string]interface{}, _a2 error) *MockRenderParser_Parse_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockRenderParser_Parse_Call) RunAndReturn(run func([]byte) (map[string]interface{}, map[string]interface{}, error)) *MockRenderParser_Parse_Call {
	_c.Call.Return(run)
	return _c
}

// Render provides a mock function with given fields: w, reqID, params, deps
func (_m *MockRenderParser) Render(w io.Writer, reqID int64, params map[string]interface{}, deps map[string]interface{}) error {
	ret := _m.Called(w, reqID, params, deps)

	if len(ret) == 0 {
		panic("no return value specified for Render")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(io.Writer, int64, map[string]interface{}, map[string]interface{}) error); ok {
		r0 = rf(w, reqID, params, deps)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRenderParser_Render_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Render'
type MockRenderParser_Render_Call struct {
	*mock.Call
}

// Render is a helper method to define mock.On call
//   - w io.Writer
//   - reqID int64
//   - params map[string]interface{}
//   - deps map[string]interface{}
func (_e *MockRenderParser_Expecter) Render(w interface{}, reqID interface{}, params interface{}, deps interface{}) *MockRenderParser_Render_Call {
	return &MockRenderParser_Render_Call{Call: _e.mock.On("Render", w, reqID, params, deps)}
}

func (_c *MockRenderParser_Render_Call) Run(run func(w io.Writer, reqID int64, params map[string]interface{}, deps map[string]interface{})) *MockRenderParser_Render_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(io.Writer), args[1].(int64), args[2].(map[string]interface{}), args[3].(map[string]interface{}))
	})
	return _c
}

func (_c *MockRenderParser_Render_Call) Return(_a0 error) *MockRenderParser_Render_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRenderParser_Render_Call) RunAndReturn(run func(io.Writer, int64, map[string]interface{}, map[string]interface{}) error) *MockRenderParser_Render_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockRenderParser creates a new instance of MockRenderParser. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockRenderParser(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockRenderParser {
	mock := &MockRenderParser{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
