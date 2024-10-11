// Code generated by mockery v2.46.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// FixtureSet is an autogenerated mock type for the FixtureSet type
type FixtureSet struct {
	mock.Mock
}

type FixtureSet_Expecter struct {
	mock *mock.Mock
}

func (_m *FixtureSet) EXPECT() *FixtureSet_Expecter {
	return &FixtureSet_Expecter{mock: &_m.Mock}
}

// Write provides a mock function with given fields: ctx
func (_m *FixtureSet) Write(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Write")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FixtureSet_Write_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Write'
type FixtureSet_Write_Call struct {
	*mock.Call
}

// Write is a helper method to define mock.On call
//   - ctx context.Context
func (_e *FixtureSet_Expecter) Write(ctx interface{}) *FixtureSet_Write_Call {
	return &FixtureSet_Write_Call{Call: _e.mock.On("Write", ctx)}
}

func (_c *FixtureSet_Write_Call) Run(run func(ctx context.Context)) *FixtureSet_Write_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *FixtureSet_Write_Call) Return(_a0 error) *FixtureSet_Write_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FixtureSet_Write_Call) RunAndReturn(run func(context.Context) error) *FixtureSet_Write_Call {
	_c.Call.Return(run)
	return _c
}

// NewFixtureSet creates a new instance of FixtureSet. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewFixtureSet(t interface {
	mock.TestingT
	Cleanup(func())
}) *FixtureSet {
	mock := &FixtureSet{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
