// Code generated by mockery v2.46.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// TypedModule is an autogenerated mock type for the TypedModule type
type TypedModule struct {
	mock.Mock
}

type TypedModule_Expecter struct {
	mock *mock.Mock
}

func (_m *TypedModule) EXPECT() *TypedModule_Expecter {
	return &TypedModule_Expecter{mock: &_m.Mock}
}

// IsBackground provides a mock function with given fields:
func (_m *TypedModule) IsBackground() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IsBackground")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// TypedModule_IsBackground_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsBackground'
type TypedModule_IsBackground_Call struct {
	*mock.Call
}

// IsBackground is a helper method to define mock.On call
func (_e *TypedModule_Expecter) IsBackground() *TypedModule_IsBackground_Call {
	return &TypedModule_IsBackground_Call{Call: _e.mock.On("IsBackground")}
}

func (_c *TypedModule_IsBackground_Call) Run(run func()) *TypedModule_IsBackground_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *TypedModule_IsBackground_Call) Return(_a0 bool) *TypedModule_IsBackground_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *TypedModule_IsBackground_Call) RunAndReturn(run func() bool) *TypedModule_IsBackground_Call {
	_c.Call.Return(run)
	return _c
}

// IsEssential provides a mock function with given fields:
func (_m *TypedModule) IsEssential() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IsEssential")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// TypedModule_IsEssential_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsEssential'
type TypedModule_IsEssential_Call struct {
	*mock.Call
}

// IsEssential is a helper method to define mock.On call
func (_e *TypedModule_Expecter) IsEssential() *TypedModule_IsEssential_Call {
	return &TypedModule_IsEssential_Call{Call: _e.mock.On("IsEssential")}
}

func (_c *TypedModule_IsEssential_Call) Run(run func()) *TypedModule_IsEssential_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *TypedModule_IsEssential_Call) Return(_a0 bool) *TypedModule_IsEssential_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *TypedModule_IsEssential_Call) RunAndReturn(run func() bool) *TypedModule_IsEssential_Call {
	_c.Call.Return(run)
	return _c
}

// NewTypedModule creates a new instance of TypedModule. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTypedModule(t interface {
	mock.TestingT
	Cleanup(func())
}) *TypedModule {
	mock := &TypedModule{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
