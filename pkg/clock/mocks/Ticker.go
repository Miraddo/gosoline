// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// Ticker is an autogenerated mock type for the Ticker type
type Ticker struct {
	mock.Mock
}

type Ticker_Expecter struct {
	mock *mock.Mock
}

func (_m *Ticker) EXPECT() *Ticker_Expecter {
	return &Ticker_Expecter{mock: &_m.Mock}
}

// Chan provides a mock function with given fields:
func (_m *Ticker) Chan() <-chan time.Time {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Chan")
	}

	var r0 <-chan time.Time
	if rf, ok := ret.Get(0).(func() <-chan time.Time); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan time.Time)
		}
	}

	return r0
}

// Ticker_Chan_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Chan'
type Ticker_Chan_Call struct {
	*mock.Call
}

// Chan is a helper method to define mock.On call
func (_e *Ticker_Expecter) Chan() *Ticker_Chan_Call {
	return &Ticker_Chan_Call{Call: _e.mock.On("Chan")}
}

func (_c *Ticker_Chan_Call) Run(run func()) *Ticker_Chan_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Ticker_Chan_Call) Return(_a0 <-chan time.Time) *Ticker_Chan_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Ticker_Chan_Call) RunAndReturn(run func() <-chan time.Time) *Ticker_Chan_Call {
	_c.Call.Return(run)
	return _c
}

// Reset provides a mock function with given fields: d
func (_m *Ticker) Reset(d time.Duration) {
	_m.Called(d)
}

// Ticker_Reset_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Reset'
type Ticker_Reset_Call struct {
	*mock.Call
}

// Reset is a helper method to define mock.On call
//   - d time.Duration
func (_e *Ticker_Expecter) Reset(d interface{}) *Ticker_Reset_Call {
	return &Ticker_Reset_Call{Call: _e.mock.On("Reset", d)}
}

func (_c *Ticker_Reset_Call) Run(run func(d time.Duration)) *Ticker_Reset_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(time.Duration))
	})
	return _c
}

func (_c *Ticker_Reset_Call) Return() *Ticker_Reset_Call {
	_c.Call.Return()
	return _c
}

func (_c *Ticker_Reset_Call) RunAndReturn(run func(time.Duration)) *Ticker_Reset_Call {
	_c.Call.Return(run)
	return _c
}

// Stop provides a mock function with given fields:
func (_m *Ticker) Stop() {
	_m.Called()
}

// Ticker_Stop_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Stop'
type Ticker_Stop_Call struct {
	*mock.Call
}

// Stop is a helper method to define mock.On call
func (_e *Ticker_Expecter) Stop() *Ticker_Stop_Call {
	return &Ticker_Stop_Call{Call: _e.mock.On("Stop")}
}

func (_c *Ticker_Stop_Call) Run(run func()) *Ticker_Stop_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Ticker_Stop_Call) Return() *Ticker_Stop_Call {
	_c.Call.Return()
	return _c
}

func (_c *Ticker_Stop_Call) RunAndReturn(run func()) *Ticker_Stop_Call {
	_c.Call.Return(run)
	return _c
}

// NewTicker creates a new instance of Ticker. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTicker(t interface {
	mock.TestingT
	Cleanup(func())
}) *Ticker {
	mock := &Ticker{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
