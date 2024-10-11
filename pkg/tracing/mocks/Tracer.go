// Code generated by mockery v2.46.0. DO NOT EDIT.

package mocks

import (
	context "context"
	http "net/http"

	mock "github.com/stretchr/testify/mock"

	tracing "github.com/justtrackio/gosoline/pkg/tracing"
)

// Tracer is an autogenerated mock type for the Tracer type
type Tracer struct {
	mock.Mock
}

type Tracer_Expecter struct {
	mock *mock.Mock
}

func (_m *Tracer) EXPECT() *Tracer_Expecter {
	return &Tracer_Expecter{mock: &_m.Mock}
}

// HttpHandler provides a mock function with given fields: h
func (_m *Tracer) HttpHandler(h http.Handler) http.Handler {
	ret := _m.Called(h)

	if len(ret) == 0 {
		panic("no return value specified for HttpHandler")
	}

	var r0 http.Handler
	if rf, ok := ret.Get(0).(func(http.Handler) http.Handler); ok {
		r0 = rf(h)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Handler)
		}
	}

	return r0
}

// Tracer_HttpHandler_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HttpHandler'
type Tracer_HttpHandler_Call struct {
	*mock.Call
}

// HttpHandler is a helper method to define mock.On call
//   - h http.Handler
func (_e *Tracer_Expecter) HttpHandler(h interface{}) *Tracer_HttpHandler_Call {
	return &Tracer_HttpHandler_Call{Call: _e.mock.On("HttpHandler", h)}
}

func (_c *Tracer_HttpHandler_Call) Run(run func(h http.Handler)) *Tracer_HttpHandler_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(http.Handler))
	})
	return _c
}

func (_c *Tracer_HttpHandler_Call) Return(_a0 http.Handler) *Tracer_HttpHandler_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Tracer_HttpHandler_Call) RunAndReturn(run func(http.Handler) http.Handler) *Tracer_HttpHandler_Call {
	_c.Call.Return(run)
	return _c
}

// StartSpan provides a mock function with given fields: name
func (_m *Tracer) StartSpan(name string) (context.Context, tracing.Span) {
	ret := _m.Called(name)

	if len(ret) == 0 {
		panic("no return value specified for StartSpan")
	}

	var r0 context.Context
	var r1 tracing.Span
	if rf, ok := ret.Get(0).(func(string) (context.Context, tracing.Span)); ok {
		return rf(name)
	}
	if rf, ok := ret.Get(0).(func(string) context.Context); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	if rf, ok := ret.Get(1).(func(string) tracing.Span); ok {
		r1 = rf(name)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(tracing.Span)
		}
	}

	return r0, r1
}

// Tracer_StartSpan_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StartSpan'
type Tracer_StartSpan_Call struct {
	*mock.Call
}

// StartSpan is a helper method to define mock.On call
//   - name string
func (_e *Tracer_Expecter) StartSpan(name interface{}) *Tracer_StartSpan_Call {
	return &Tracer_StartSpan_Call{Call: _e.mock.On("StartSpan", name)}
}

func (_c *Tracer_StartSpan_Call) Run(run func(name string)) *Tracer_StartSpan_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Tracer_StartSpan_Call) Return(_a0 context.Context, _a1 tracing.Span) *Tracer_StartSpan_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Tracer_StartSpan_Call) RunAndReturn(run func(string) (context.Context, tracing.Span)) *Tracer_StartSpan_Call {
	_c.Call.Return(run)
	return _c
}

// StartSpanFromContext provides a mock function with given fields: ctx, name
func (_m *Tracer) StartSpanFromContext(ctx context.Context, name string) (context.Context, tracing.Span) {
	ret := _m.Called(ctx, name)

	if len(ret) == 0 {
		panic("no return value specified for StartSpanFromContext")
	}

	var r0 context.Context
	var r1 tracing.Span
	if rf, ok := ret.Get(0).(func(context.Context, string) (context.Context, tracing.Span)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) context.Context); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) tracing.Span); ok {
		r1 = rf(ctx, name)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(tracing.Span)
		}
	}

	return r0, r1
}

// Tracer_StartSpanFromContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StartSpanFromContext'
type Tracer_StartSpanFromContext_Call struct {
	*mock.Call
}

// StartSpanFromContext is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
func (_e *Tracer_Expecter) StartSpanFromContext(ctx interface{}, name interface{}) *Tracer_StartSpanFromContext_Call {
	return &Tracer_StartSpanFromContext_Call{Call: _e.mock.On("StartSpanFromContext", ctx, name)}
}

func (_c *Tracer_StartSpanFromContext_Call) Run(run func(ctx context.Context, name string)) *Tracer_StartSpanFromContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *Tracer_StartSpanFromContext_Call) Return(_a0 context.Context, _a1 tracing.Span) *Tracer_StartSpanFromContext_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Tracer_StartSpanFromContext_Call) RunAndReturn(run func(context.Context, string) (context.Context, tracing.Span)) *Tracer_StartSpanFromContext_Call {
	_c.Call.Return(run)
	return _c
}

// StartSubSpan provides a mock function with given fields: ctx, name
func (_m *Tracer) StartSubSpan(ctx context.Context, name string) (context.Context, tracing.Span) {
	ret := _m.Called(ctx, name)

	if len(ret) == 0 {
		panic("no return value specified for StartSubSpan")
	}

	var r0 context.Context
	var r1 tracing.Span
	if rf, ok := ret.Get(0).(func(context.Context, string) (context.Context, tracing.Span)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) context.Context); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) tracing.Span); ok {
		r1 = rf(ctx, name)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(tracing.Span)
		}
	}

	return r0, r1
}

// Tracer_StartSubSpan_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StartSubSpan'
type Tracer_StartSubSpan_Call struct {
	*mock.Call
}

// StartSubSpan is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
func (_e *Tracer_Expecter) StartSubSpan(ctx interface{}, name interface{}) *Tracer_StartSubSpan_Call {
	return &Tracer_StartSubSpan_Call{Call: _e.mock.On("StartSubSpan", ctx, name)}
}

func (_c *Tracer_StartSubSpan_Call) Run(run func(ctx context.Context, name string)) *Tracer_StartSubSpan_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *Tracer_StartSubSpan_Call) Return(_a0 context.Context, _a1 tracing.Span) *Tracer_StartSubSpan_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Tracer_StartSubSpan_Call) RunAndReturn(run func(context.Context, string) (context.Context, tracing.Span)) *Tracer_StartSubSpan_Call {
	_c.Call.Return(run)
	return _c
}

// NewTracer creates a new instance of Tracer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTracer(t interface {
	mock.TestingT
	Cleanup(func())
}) *Tracer {
	mock := &Tracer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
