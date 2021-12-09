// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// Ticker is an autogenerated mock type for the Ticker type
type Ticker struct {
	mock.Mock
}

// Chan provides a mock function with given fields:
func (_m *Ticker) Chan() <-chan time.Time {
	ret := _m.Called()

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

// Reset provides a mock function with given fields: d
func (_m *Ticker) Reset(d time.Duration) {
	_m.Called(d)
}

// Stop provides a mock function with given fields:
func (_m *Ticker) Stop() {
	_m.Called()
}
