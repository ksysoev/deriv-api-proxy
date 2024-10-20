// Code generated by mockery v2.45.1. DO NOT EDIT.

//go:build !compile

package repo

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	"go.etcd.io/etcd/clientv3"
)

// MockKV is an autogenerated mock type for the KV type
type MockKV struct {
	mock.Mock
}

type MockKV_Expecter struct {
	mock *mock.Mock
}

func (_m *MockKV) EXPECT() *MockKV_Expecter {
	return &MockKV_Expecter{mock: &_m.Mock}
}

// Compact provides a mock function with given fields: ctx, rev, opts
func (_m *MockKV) Compact(ctx context.Context, rev int64, opts ...clientv3.CompactOption) (*clientv3.CompactResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, rev)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Compact")
	}

	var r0 *clientv3.CompactResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, ...clientv3.CompactOption) (*clientv3.CompactResponse, error)); ok {
		return rf(ctx, rev, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, ...clientv3.CompactOption) *clientv3.CompactResponse); ok {
		r0 = rf(ctx, rev, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientv3.CompactResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, ...clientv3.CompactOption) error); ok {
		r1 = rf(ctx, rev, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockKV_Compact_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Compact'
type MockKV_Compact_Call struct {
	*mock.Call
}

// Compact is a helper method to define mock.On call
//   - ctx context.Context
//   - rev int64
//   - opts ...clientv3.CompactOption
func (_e *MockKV_Expecter) Compact(ctx interface{}, rev interface{}, opts ...interface{}) *MockKV_Compact_Call {
	return &MockKV_Compact_Call{Call: _e.mock.On("Compact",
		append([]interface{}{ctx, rev}, opts...)...)}
}

func (_c *MockKV_Compact_Call) Run(run func(ctx context.Context, rev int64, opts ...clientv3.CompactOption)) *MockKV_Compact_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]clientv3.CompactOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(clientv3.CompactOption)
			}
		}
		run(args[0].(context.Context), args[1].(int64), variadicArgs...)
	})
	return _c
}

func (_c *MockKV_Compact_Call) Return(_a0 *clientv3.CompactResponse, _a1 error) *MockKV_Compact_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockKV_Compact_Call) RunAndReturn(run func(context.Context, int64, ...clientv3.CompactOption) (*clientv3.CompactResponse, error)) *MockKV_Compact_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, key, opts
func (_m *MockKV) Delete(ctx context.Context, key string, opts ...clientv3.OpOption) (*clientv3.DeleteResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, key)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 *clientv3.DeleteResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ...clientv3.OpOption) (*clientv3.DeleteResponse, error)); ok {
		return rf(ctx, key, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, ...clientv3.OpOption) *clientv3.DeleteResponse); ok {
		r0 = rf(ctx, key, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientv3.DeleteResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, ...clientv3.OpOption) error); ok {
		r1 = rf(ctx, key, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockKV_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockKV_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - key string
//   - opts ...clientv3.OpOption
func (_e *MockKV_Expecter) Delete(ctx interface{}, key interface{}, opts ...interface{}) *MockKV_Delete_Call {
	return &MockKV_Delete_Call{Call: _e.mock.On("Delete",
		append([]interface{}{ctx, key}, opts...)...)}
}

func (_c *MockKV_Delete_Call) Run(run func(ctx context.Context, key string, opts ...clientv3.OpOption)) *MockKV_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]clientv3.OpOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(clientv3.OpOption)
			}
		}
		run(args[0].(context.Context), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockKV_Delete_Call) Return(_a0 *clientv3.DeleteResponse, _a1 error) *MockKV_Delete_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockKV_Delete_Call) RunAndReturn(run func(context.Context, string, ...clientv3.OpOption) (*clientv3.DeleteResponse, error)) *MockKV_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Do provides a mock function with given fields: ctx, op
func (_m *MockKV) Do(ctx context.Context, op clientv3.Op) (clientv3.OpResponse, error) {
	ret := _m.Called(ctx, op)

	if len(ret) == 0 {
		panic("no return value specified for Do")
	}

	var r0 clientv3.OpResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, clientv3.Op) (clientv3.OpResponse, error)); ok {
		return rf(ctx, op)
	}
	if rf, ok := ret.Get(0).(func(context.Context, clientv3.Op) clientv3.OpResponse); ok {
		r0 = rf(ctx, op)
	} else {
		r0 = ret.Get(0).(clientv3.OpResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, clientv3.Op) error); ok {
		r1 = rf(ctx, op)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockKV_Do_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Do'
type MockKV_Do_Call struct {
	*mock.Call
}

// Do is a helper method to define mock.On call
//   - ctx context.Context
//   - op clientv3.Op
func (_e *MockKV_Expecter) Do(ctx interface{}, op interface{}) *MockKV_Do_Call {
	return &MockKV_Do_Call{Call: _e.mock.On("Do", ctx, op)}
}

func (_c *MockKV_Do_Call) Run(run func(ctx context.Context, op clientv3.Op)) *MockKV_Do_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(clientv3.Op))
	})
	return _c
}

func (_c *MockKV_Do_Call) Return(_a0 clientv3.OpResponse, _a1 error) *MockKV_Do_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockKV_Do_Call) RunAndReturn(run func(context.Context, clientv3.Op) (clientv3.OpResponse, error)) *MockKV_Do_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, key, opts
func (_m *MockKV) Get(ctx context.Context, key string, opts ...clientv3.OpOption) (*clientv3.GetResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, key)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *clientv3.GetResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ...clientv3.OpOption) (*clientv3.GetResponse, error)); ok {
		return rf(ctx, key, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, ...clientv3.OpOption) *clientv3.GetResponse); ok {
		r0 = rf(ctx, key, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientv3.GetResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, ...clientv3.OpOption) error); ok {
		r1 = rf(ctx, key, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockKV_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockKV_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - key string
//   - opts ...clientv3.OpOption
func (_e *MockKV_Expecter) Get(ctx interface{}, key interface{}, opts ...interface{}) *MockKV_Get_Call {
	return &MockKV_Get_Call{Call: _e.mock.On("Get",
		append([]interface{}{ctx, key}, opts...)...)}
}

func (_c *MockKV_Get_Call) Run(run func(ctx context.Context, key string, opts ...clientv3.OpOption)) *MockKV_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]clientv3.OpOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(clientv3.OpOption)
			}
		}
		run(args[0].(context.Context), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockKV_Get_Call) Return(_a0 *clientv3.GetResponse, _a1 error) *MockKV_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockKV_Get_Call) RunAndReturn(run func(context.Context, string, ...clientv3.OpOption) (*clientv3.GetResponse, error)) *MockKV_Get_Call {
	_c.Call.Return(run)
	return _c
}

// Put provides a mock function with given fields: ctx, key, val, opts
func (_m *MockKV) Put(ctx context.Context, key string, val string, opts ...clientv3.OpOption) (*clientv3.PutResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, key, val)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Put")
	}

	var r0 *clientv3.PutResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, ...clientv3.OpOption) (*clientv3.PutResponse, error)); ok {
		return rf(ctx, key, val, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, ...clientv3.OpOption) *clientv3.PutResponse); ok {
		r0 = rf(ctx, key, val, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientv3.PutResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, ...clientv3.OpOption) error); ok {
		r1 = rf(ctx, key, val, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockKV_Put_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Put'
type MockKV_Put_Call struct {
	*mock.Call
}

// Put is a helper method to define mock.On call
//   - ctx context.Context
//   - key string
//   - val string
//   - opts ...clientv3.OpOption
func (_e *MockKV_Expecter) Put(ctx interface{}, key interface{}, val interface{}, opts ...interface{}) *MockKV_Put_Call {
	return &MockKV_Put_Call{Call: _e.mock.On("Put",
		append([]interface{}{ctx, key, val}, opts...)...)}
}

func (_c *MockKV_Put_Call) Run(run func(ctx context.Context, key string, val string, opts ...clientv3.OpOption)) *MockKV_Put_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]clientv3.OpOption, len(args)-3)
		for i, a := range args[3:] {
			if a != nil {
				variadicArgs[i] = a.(clientv3.OpOption)
			}
		}
		run(args[0].(context.Context), args[1].(string), args[2].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockKV_Put_Call) Return(_a0 *clientv3.PutResponse, _a1 error) *MockKV_Put_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockKV_Put_Call) RunAndReturn(run func(context.Context, string, string, ...clientv3.OpOption) (*clientv3.PutResponse, error)) *MockKV_Put_Call {
	_c.Call.Return(run)
	return _c
}

// Txn provides a mock function with given fields: ctx
func (_m *MockKV) Txn(ctx context.Context) clientv3.Txn {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Txn")
	}

	var r0 clientv3.Txn
	if rf, ok := ret.Get(0).(func(context.Context) clientv3.Txn); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(clientv3.Txn)
		}
	}

	return r0
}

// MockKV_Txn_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Txn'
type MockKV_Txn_Call struct {
	*mock.Call
}

// Txn is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockKV_Expecter) Txn(ctx interface{}) *MockKV_Txn_Call {
	return &MockKV_Txn_Call{Call: _e.mock.On("Txn", ctx)}
}

func (_c *MockKV_Txn_Call) Run(run func(ctx context.Context)) *MockKV_Txn_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockKV_Txn_Call) Return(_a0 clientv3.Txn) *MockKV_Txn_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockKV_Txn_Call) RunAndReturn(run func(context.Context) clientv3.Txn) *MockKV_Txn_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockKV creates a new instance of MockKV. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockKV(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockKV {
	mock := &MockKV{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
