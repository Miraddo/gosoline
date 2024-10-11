// Code generated by mockery v2.46.0. DO NOT EDIT.

package mocks

import (
	context "context"

	stream "github.com/justtrackio/gosoline/pkg/stream"
	mock "github.com/stretchr/testify/mock"
)

// RetryHandler is an autogenerated mock type for the RetryHandler type
type RetryHandler struct {
	mock.Mock
}

type RetryHandler_Expecter struct {
	mock *mock.Mock
}

func (_m *RetryHandler) EXPECT() *RetryHandler_Expecter {
	return &RetryHandler_Expecter{mock: &_m.Mock}
}

// Put provides a mock function with given fields: ctx, msg
func (_m *RetryHandler) Put(ctx context.Context, msg *stream.Message) error {
	ret := _m.Called(ctx, msg)

	if len(ret) == 0 {
		panic("no return value specified for Put")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *stream.Message) error); ok {
		r0 = rf(ctx, msg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RetryHandler_Put_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Put'
type RetryHandler_Put_Call struct {
	*mock.Call
}

// Put is a helper method to define mock.On call
//   - ctx context.Context
//   - msg *stream.Message
func (_e *RetryHandler_Expecter) Put(ctx interface{}, msg interface{}) *RetryHandler_Put_Call {
	return &RetryHandler_Put_Call{Call: _e.mock.On("Put", ctx, msg)}
}

func (_c *RetryHandler_Put_Call) Run(run func(ctx context.Context, msg *stream.Message)) *RetryHandler_Put_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*stream.Message))
	})
	return _c
}

func (_c *RetryHandler_Put_Call) Return(_a0 error) *RetryHandler_Put_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RetryHandler_Put_Call) RunAndReturn(run func(context.Context, *stream.Message) error) *RetryHandler_Put_Call {
	_c.Call.Return(run)
	return _c
}

// NewRetryHandler creates a new instance of RetryHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRetryHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *RetryHandler {
	mock := &RetryHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
