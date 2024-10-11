// Code generated by mockery v2.46.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// BatchConsumerCallback is an autogenerated mock type for the BatchConsumerCallback type
type BatchConsumerCallback struct {
	mock.Mock
}

type BatchConsumerCallback_Expecter struct {
	mock *mock.Mock
}

func (_m *BatchConsumerCallback) EXPECT() *BatchConsumerCallback_Expecter {
	return &BatchConsumerCallback_Expecter{mock: &_m.Mock}
}

// Consume provides a mock function with given fields: ctx, models, attributes
func (_m *BatchConsumerCallback) Consume(ctx context.Context, models []interface{}, attributes []map[string]string) ([]bool, error) {
	ret := _m.Called(ctx, models, attributes)

	if len(ret) == 0 {
		panic("no return value specified for Consume")
	}

	var r0 []bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []interface{}, []map[string]string) ([]bool, error)); ok {
		return rf(ctx, models, attributes)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []interface{}, []map[string]string) []bool); ok {
		r0 = rf(ctx, models, attributes)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]bool)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []interface{}, []map[string]string) error); ok {
		r1 = rf(ctx, models, attributes)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BatchConsumerCallback_Consume_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Consume'
type BatchConsumerCallback_Consume_Call struct {
	*mock.Call
}

// Consume is a helper method to define mock.On call
//   - ctx context.Context
//   - models []interface{}
//   - attributes []map[string]string
func (_e *BatchConsumerCallback_Expecter) Consume(ctx interface{}, models interface{}, attributes interface{}) *BatchConsumerCallback_Consume_Call {
	return &BatchConsumerCallback_Consume_Call{Call: _e.mock.On("Consume", ctx, models, attributes)}
}

func (_c *BatchConsumerCallback_Consume_Call) Run(run func(ctx context.Context, models []interface{}, attributes []map[string]string)) *BatchConsumerCallback_Consume_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]interface{}), args[2].([]map[string]string))
	})
	return _c
}

func (_c *BatchConsumerCallback_Consume_Call) Return(_a0 []bool, _a1 error) *BatchConsumerCallback_Consume_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *BatchConsumerCallback_Consume_Call) RunAndReturn(run func(context.Context, []interface{}, []map[string]string) ([]bool, error)) *BatchConsumerCallback_Consume_Call {
	_c.Call.Return(run)
	return _c
}

// GetModel provides a mock function with given fields: attributes
func (_m *BatchConsumerCallback) GetModel(attributes map[string]string) interface{} {
	ret := _m.Called(attributes)

	if len(ret) == 0 {
		panic("no return value specified for GetModel")
	}

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(map[string]string) interface{}); ok {
		r0 = rf(attributes)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// BatchConsumerCallback_GetModel_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetModel'
type BatchConsumerCallback_GetModel_Call struct {
	*mock.Call
}

// GetModel is a helper method to define mock.On call
//   - attributes map[string]string
func (_e *BatchConsumerCallback_Expecter) GetModel(attributes interface{}) *BatchConsumerCallback_GetModel_Call {
	return &BatchConsumerCallback_GetModel_Call{Call: _e.mock.On("GetModel", attributes)}
}

func (_c *BatchConsumerCallback_GetModel_Call) Run(run func(attributes map[string]string)) *BatchConsumerCallback_GetModel_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(map[string]string))
	})
	return _c
}

func (_c *BatchConsumerCallback_GetModel_Call) Return(_a0 interface{}) *BatchConsumerCallback_GetModel_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BatchConsumerCallback_GetModel_Call) RunAndReturn(run func(map[string]string) interface{}) *BatchConsumerCallback_GetModel_Call {
	_c.Call.Return(run)
	return _c
}

// NewBatchConsumerCallback creates a new instance of BatchConsumerCallback. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBatchConsumerCallback(t interface {
	mock.TestingT
	Cleanup(func())
}) *BatchConsumerCallback {
	mock := &BatchConsumerCallback{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
