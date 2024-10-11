// Code generated by mockery v2.46.0. DO NOT EDIT.

package mocks

import (
	tracing "github.com/justtrackio/gosoline/pkg/tracing"
	mock "github.com/stretchr/testify/mock"
)

// TraceAble is an autogenerated mock type for the TraceAble type
type TraceAble struct {
	mock.Mock
}

type TraceAble_Expecter struct {
	mock *mock.Mock
}

func (_m *TraceAble) EXPECT() *TraceAble_Expecter {
	return &TraceAble_Expecter{mock: &_m.Mock}
}

// GetTrace provides a mock function with given fields:
func (_m *TraceAble) GetTrace() *tracing.Trace {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetTrace")
	}

	var r0 *tracing.Trace
	if rf, ok := ret.Get(0).(func() *tracing.Trace); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*tracing.Trace)
		}
	}

	return r0
}

// TraceAble_GetTrace_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTrace'
type TraceAble_GetTrace_Call struct {
	*mock.Call
}

// GetTrace is a helper method to define mock.On call
func (_e *TraceAble_Expecter) GetTrace() *TraceAble_GetTrace_Call {
	return &TraceAble_GetTrace_Call{Call: _e.mock.On("GetTrace")}
}

func (_c *TraceAble_GetTrace_Call) Run(run func()) *TraceAble_GetTrace_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *TraceAble_GetTrace_Call) Return(_a0 *tracing.Trace) *TraceAble_GetTrace_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *TraceAble_GetTrace_Call) RunAndReturn(run func() *tracing.Trace) *TraceAble_GetTrace_Call {
	_c.Call.Return(run)
	return _c
}

// NewTraceAble creates a new instance of TraceAble. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTraceAble(t interface {
	mock.TestingT
	Cleanup(func())
}) *TraceAble {
	mock := &TraceAble{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
