// Code generated by mockery v2.45.1. DO NOT EDIT.

//go:build !compile

package repo

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	"go.etcd.io/etcd/clientv3"
)

// MockLease is an autogenerated mock type for the Lease type
type MockLease struct {
	mock.Mock
}

type MockLease_Expecter struct {
	mock *mock.Mock
}

func (_m *MockLease) EXPECT() *MockLease_Expecter {
	return &MockLease_Expecter{mock: &_m.Mock}
}

// Close provides a mock function with given fields:
func (_m *MockLease) Close() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Close")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockLease_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type MockLease_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *MockLease_Expecter) Close() *MockLease_Close_Call {
	return &MockLease_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *MockLease_Close_Call) Run(run func()) *MockLease_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockLease_Close_Call) Return(_a0 error) *MockLease_Close_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockLease_Close_Call) RunAndReturn(run func() error) *MockLease_Close_Call {
	_c.Call.Return(run)
	return _c
}

// Grant provides a mock function with given fields: ctx, ttl
func (_m *MockLease) Grant(ctx context.Context, ttl int64) (*clientv3.LeaseGrantResponse, error) {
	ret := _m.Called(ctx, ttl)

	if len(ret) == 0 {
		panic("no return value specified for Grant")
	}

	var r0 *clientv3.LeaseGrantResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (*clientv3.LeaseGrantResponse, error)); ok {
		return rf(ctx, ttl)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) *clientv3.LeaseGrantResponse); ok {
		r0 = rf(ctx, ttl)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientv3.LeaseGrantResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, ttl)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockLease_Grant_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Grant'
type MockLease_Grant_Call struct {
	*mock.Call
}

// Grant is a helper method to define mock.On call
//   - ctx context.Context
//   - ttl int64
func (_e *MockLease_Expecter) Grant(ctx interface{}, ttl interface{}) *MockLease_Grant_Call {
	return &MockLease_Grant_Call{Call: _e.mock.On("Grant", ctx, ttl)}
}

func (_c *MockLease_Grant_Call) Run(run func(ctx context.Context, ttl int64)) *MockLease_Grant_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *MockLease_Grant_Call) Return(_a0 *clientv3.LeaseGrantResponse, _a1 error) *MockLease_Grant_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockLease_Grant_Call) RunAndReturn(run func(context.Context, int64) (*clientv3.LeaseGrantResponse, error)) *MockLease_Grant_Call {
	_c.Call.Return(run)
	return _c
}

// KeepAlive provides a mock function with given fields: ctx, id
func (_m *MockLease) KeepAlive(ctx context.Context, id clientv3.LeaseID) (<-chan *clientv3.LeaseKeepAliveResponse, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for KeepAlive")
	}

	var r0 <-chan *clientv3.LeaseKeepAliveResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, clientv3.LeaseID) (<-chan *clientv3.LeaseKeepAliveResponse, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, clientv3.LeaseID) <-chan *clientv3.LeaseKeepAliveResponse); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan *clientv3.LeaseKeepAliveResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, clientv3.LeaseID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockLease_KeepAlive_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'KeepAlive'
type MockLease_KeepAlive_Call struct {
	*mock.Call
}

// KeepAlive is a helper method to define mock.On call
//   - ctx context.Context
//   - id clientv3.LeaseID
func (_e *MockLease_Expecter) KeepAlive(ctx interface{}, id interface{}) *MockLease_KeepAlive_Call {
	return &MockLease_KeepAlive_Call{Call: _e.mock.On("KeepAlive", ctx, id)}
}

func (_c *MockLease_KeepAlive_Call) Run(run func(ctx context.Context, id clientv3.LeaseID)) *MockLease_KeepAlive_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(clientv3.LeaseID))
	})
	return _c
}

func (_c *MockLease_KeepAlive_Call) Return(_a0 <-chan *clientv3.LeaseKeepAliveResponse, _a1 error) *MockLease_KeepAlive_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockLease_KeepAlive_Call) RunAndReturn(run func(context.Context, clientv3.LeaseID) (<-chan *clientv3.LeaseKeepAliveResponse, error)) *MockLease_KeepAlive_Call {
	_c.Call.Return(run)
	return _c
}

// KeepAliveOnce provides a mock function with given fields: ctx, id
func (_m *MockLease) KeepAliveOnce(ctx context.Context, id clientv3.LeaseID) (*clientv3.LeaseKeepAliveResponse, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for KeepAliveOnce")
	}

	var r0 *clientv3.LeaseKeepAliveResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, clientv3.LeaseID) (*clientv3.LeaseKeepAliveResponse, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, clientv3.LeaseID) *clientv3.LeaseKeepAliveResponse); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientv3.LeaseKeepAliveResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, clientv3.LeaseID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockLease_KeepAliveOnce_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'KeepAliveOnce'
type MockLease_KeepAliveOnce_Call struct {
	*mock.Call
}

// KeepAliveOnce is a helper method to define mock.On call
//   - ctx context.Context
//   - id clientv3.LeaseID
func (_e *MockLease_Expecter) KeepAliveOnce(ctx interface{}, id interface{}) *MockLease_KeepAliveOnce_Call {
	return &MockLease_KeepAliveOnce_Call{Call: _e.mock.On("KeepAliveOnce", ctx, id)}
}

func (_c *MockLease_KeepAliveOnce_Call) Run(run func(ctx context.Context, id clientv3.LeaseID)) *MockLease_KeepAliveOnce_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(clientv3.LeaseID))
	})
	return _c
}

func (_c *MockLease_KeepAliveOnce_Call) Return(_a0 *clientv3.LeaseKeepAliveResponse, _a1 error) *MockLease_KeepAliveOnce_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockLease_KeepAliveOnce_Call) RunAndReturn(run func(context.Context, clientv3.LeaseID) (*clientv3.LeaseKeepAliveResponse, error)) *MockLease_KeepAliveOnce_Call {
	_c.Call.Return(run)
	return _c
}

// Leases provides a mock function with given fields: ctx
func (_m *MockLease) Leases(ctx context.Context) (*clientv3.LeaseLeasesResponse, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Leases")
	}

	var r0 *clientv3.LeaseLeasesResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*clientv3.LeaseLeasesResponse, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *clientv3.LeaseLeasesResponse); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientv3.LeaseLeasesResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockLease_Leases_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Leases'
type MockLease_Leases_Call struct {
	*mock.Call
}

// Leases is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockLease_Expecter) Leases(ctx interface{}) *MockLease_Leases_Call {
	return &MockLease_Leases_Call{Call: _e.mock.On("Leases", ctx)}
}

func (_c *MockLease_Leases_Call) Run(run func(ctx context.Context)) *MockLease_Leases_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockLease_Leases_Call) Return(_a0 *clientv3.LeaseLeasesResponse, _a1 error) *MockLease_Leases_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockLease_Leases_Call) RunAndReturn(run func(context.Context) (*clientv3.LeaseLeasesResponse, error)) *MockLease_Leases_Call {
	_c.Call.Return(run)
	return _c
}

// Revoke provides a mock function with given fields: ctx, id
func (_m *MockLease) Revoke(ctx context.Context, id clientv3.LeaseID) (*clientv3.LeaseRevokeResponse, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Revoke")
	}

	var r0 *clientv3.LeaseRevokeResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, clientv3.LeaseID) (*clientv3.LeaseRevokeResponse, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, clientv3.LeaseID) *clientv3.LeaseRevokeResponse); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientv3.LeaseRevokeResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, clientv3.LeaseID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockLease_Revoke_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Revoke'
type MockLease_Revoke_Call struct {
	*mock.Call
}

// Revoke is a helper method to define mock.On call
//   - ctx context.Context
//   - id clientv3.LeaseID
func (_e *MockLease_Expecter) Revoke(ctx interface{}, id interface{}) *MockLease_Revoke_Call {
	return &MockLease_Revoke_Call{Call: _e.mock.On("Revoke", ctx, id)}
}

func (_c *MockLease_Revoke_Call) Run(run func(ctx context.Context, id clientv3.LeaseID)) *MockLease_Revoke_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(clientv3.LeaseID))
	})
	return _c
}

func (_c *MockLease_Revoke_Call) Return(_a0 *clientv3.LeaseRevokeResponse, _a1 error) *MockLease_Revoke_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockLease_Revoke_Call) RunAndReturn(run func(context.Context, clientv3.LeaseID) (*clientv3.LeaseRevokeResponse, error)) *MockLease_Revoke_Call {
	_c.Call.Return(run)
	return _c
}

// TimeToLive provides a mock function with given fields: ctx, id, opts
func (_m *MockLease) TimeToLive(ctx context.Context, id clientv3.LeaseID, opts ...clientv3.LeaseOption) (*clientv3.LeaseTimeToLiveResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, id)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for TimeToLive")
	}

	var r0 *clientv3.LeaseTimeToLiveResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, clientv3.LeaseID, ...clientv3.LeaseOption) (*clientv3.LeaseTimeToLiveResponse, error)); ok {
		return rf(ctx, id, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, clientv3.LeaseID, ...clientv3.LeaseOption) *clientv3.LeaseTimeToLiveResponse); ok {
		r0 = rf(ctx, id, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientv3.LeaseTimeToLiveResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, clientv3.LeaseID, ...clientv3.LeaseOption) error); ok {
		r1 = rf(ctx, id, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockLease_TimeToLive_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TimeToLive'
type MockLease_TimeToLive_Call struct {
	*mock.Call
}

// TimeToLive is a helper method to define mock.On call
//   - ctx context.Context
//   - id clientv3.LeaseID
//   - opts ...clientv3.LeaseOption
func (_e *MockLease_Expecter) TimeToLive(ctx interface{}, id interface{}, opts ...interface{}) *MockLease_TimeToLive_Call {
	return &MockLease_TimeToLive_Call{Call: _e.mock.On("TimeToLive",
		append([]interface{}{ctx, id}, opts...)...)}
}

func (_c *MockLease_TimeToLive_Call) Run(run func(ctx context.Context, id clientv3.LeaseID, opts ...clientv3.LeaseOption)) *MockLease_TimeToLive_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]clientv3.LeaseOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(clientv3.LeaseOption)
			}
		}
		run(args[0].(context.Context), args[1].(clientv3.LeaseID), variadicArgs...)
	})
	return _c
}

func (_c *MockLease_TimeToLive_Call) Return(_a0 *clientv3.LeaseTimeToLiveResponse, _a1 error) *MockLease_TimeToLive_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockLease_TimeToLive_Call) RunAndReturn(run func(context.Context, clientv3.LeaseID, ...clientv3.LeaseOption) (*clientv3.LeaseTimeToLiveResponse, error)) *MockLease_TimeToLive_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockLease creates a new instance of MockLease. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockLease(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockLease {
	mock := &MockLease{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}