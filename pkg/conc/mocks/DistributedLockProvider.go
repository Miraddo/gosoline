// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	conc "github.com/justtrackio/gosoline/pkg/conc"

	mock "github.com/stretchr/testify/mock"
)

// DistributedLockProvider is an autogenerated mock type for the DistributedLockProvider type
type DistributedLockProvider struct {
	mock.Mock
}

type DistributedLockProvider_Expecter struct {
	mock *mock.Mock
}

func (_m *DistributedLockProvider) EXPECT() *DistributedLockProvider_Expecter {
	return &DistributedLockProvider_Expecter{mock: &_m.Mock}
}

// Acquire provides a mock function with given fields: ctx, resource
func (_m *DistributedLockProvider) Acquire(ctx context.Context, resource string) (conc.DistributedLock, error) {
	ret := _m.Called(ctx, resource)

	if len(ret) == 0 {
		panic("no return value specified for Acquire")
	}

	var r0 conc.DistributedLock
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (conc.DistributedLock, error)); ok {
		return rf(ctx, resource)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) conc.DistributedLock); ok {
		r0 = rf(ctx, resource)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(conc.DistributedLock)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, resource)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DistributedLockProvider_Acquire_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Acquire'
type DistributedLockProvider_Acquire_Call struct {
	*mock.Call
}

// Acquire is a helper method to define mock.On call
//   - ctx context.Context
//   - resource string
func (_e *DistributedLockProvider_Expecter) Acquire(ctx interface{}, resource interface{}) *DistributedLockProvider_Acquire_Call {
	return &DistributedLockProvider_Acquire_Call{Call: _e.mock.On("Acquire", ctx, resource)}
}

func (_c *DistributedLockProvider_Acquire_Call) Run(run func(ctx context.Context, resource string)) *DistributedLockProvider_Acquire_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *DistributedLockProvider_Acquire_Call) Return(_a0 conc.DistributedLock, _a1 error) *DistributedLockProvider_Acquire_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DistributedLockProvider_Acquire_Call) RunAndReturn(run func(context.Context, string) (conc.DistributedLock, error)) *DistributedLockProvider_Acquire_Call {
	_c.Call.Return(run)
	return _c
}

// NewDistributedLockProvider creates a new instance of DistributedLockProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDistributedLockProvider(t interface {
	mock.TestingT
	Cleanup(func())
}) *DistributedLockProvider {
	mock := &DistributedLockProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
