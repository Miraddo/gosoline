// Code generated by mockery v2.46.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// ShardReader is an autogenerated mock type for the ShardReader type
type ShardReader struct {
	mock.Mock
}

type ShardReader_Expecter struct {
	mock *mock.Mock
}

func (_m *ShardReader) EXPECT() *ShardReader_Expecter {
	return &ShardReader_Expecter{mock: &_m.Mock}
}

// Run provides a mock function with given fields: ctx, handler
func (_m *ShardReader) Run(ctx context.Context, handler func([]byte) error) error {
	ret := _m.Called(ctx, handler)

	if len(ret) == 0 {
		panic("no return value specified for Run")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, func([]byte) error) error); ok {
		r0 = rf(ctx, handler)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ShardReader_Run_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Run'
type ShardReader_Run_Call struct {
	*mock.Call
}

// Run is a helper method to define mock.On call
//   - ctx context.Context
//   - handler func([]byte) error
func (_e *ShardReader_Expecter) Run(ctx interface{}, handler interface{}) *ShardReader_Run_Call {
	return &ShardReader_Run_Call{Call: _e.mock.On("Run", ctx, handler)}
}

func (_c *ShardReader_Run_Call) Run(run func(ctx context.Context, handler func([]byte) error)) *ShardReader_Run_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(func([]byte) error))
	})
	return _c
}

func (_c *ShardReader_Run_Call) Return(_a0 error) *ShardReader_Run_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ShardReader_Run_Call) RunAndReturn(run func(context.Context, func([]byte) error) error) *ShardReader_Run_Call {
	_c.Call.Return(run)
	return _c
}

// NewShardReader creates a new instance of ShardReader. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewShardReader(t interface {
	mock.TestingT
	Cleanup(func())
}) *ShardReader {
	mock := &ShardReader{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
