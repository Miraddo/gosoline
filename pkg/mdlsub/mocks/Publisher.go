// Code generated by mockery v2.22.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// Publisher is an autogenerated mock type for the Publisher type
type Publisher struct {
	mock.Mock
}

type Publisher_Expecter struct {
	mock *mock.Mock
}

func (_m *Publisher) EXPECT() *Publisher_Expecter {
	return &Publisher_Expecter{mock: &_m.Mock}
}

// Publish provides a mock function with given fields: ctx, typ, version, value, customAttributes
func (_m *Publisher) Publish(ctx context.Context, typ string, version int, value interface{}, customAttributes ...map[string]string) error {
	_va := make([]interface{}, len(customAttributes))
	for _i := range customAttributes {
		_va[_i] = customAttributes[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, typ, version, value)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int, interface{}, ...map[string]string) error); ok {
		r0 = rf(ctx, typ, version, value, customAttributes...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Publisher_Publish_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Publish'
type Publisher_Publish_Call struct {
	*mock.Call
}

// Publish is a helper method to define mock.On call
//   - ctx context.Context
//   - typ string
//   - version int
//   - value interface{}
//   - customAttributes ...map[string]string
func (_e *Publisher_Expecter) Publish(ctx interface{}, typ interface{}, version interface{}, value interface{}, customAttributes ...interface{}) *Publisher_Publish_Call {
	return &Publisher_Publish_Call{Call: _e.mock.On("Publish",
		append([]interface{}{ctx, typ, version, value}, customAttributes...)...)}
}

func (_c *Publisher_Publish_Call) Run(run func(ctx context.Context, typ string, version int, value interface{}, customAttributes ...map[string]string)) *Publisher_Publish_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]map[string]string, len(args)-4)
		for i, a := range args[4:] {
			if a != nil {
				variadicArgs[i] = a.(map[string]string)
			}
		}
		run(args[0].(context.Context), args[1].(string), args[2].(int), args[3].(interface{}), variadicArgs...)
	})
	return _c
}

func (_c *Publisher_Publish_Call) Return(_a0 error) *Publisher_Publish_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Publisher_Publish_Call) RunAndReturn(run func(context.Context, string, int, interface{}, ...map[string]string) error) *Publisher_Publish_Call {
	_c.Call.Return(run)
	return _c
}

// PublishBatch provides a mock function with given fields: ctx, typ, version, values, customAttributes
func (_m *Publisher) PublishBatch(ctx context.Context, typ string, version int, values []interface{}, customAttributes ...map[string]string) error {
	_va := make([]interface{}, len(customAttributes))
	for _i := range customAttributes {
		_va[_i] = customAttributes[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, typ, version, values)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int, []interface{}, ...map[string]string) error); ok {
		r0 = rf(ctx, typ, version, values, customAttributes...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Publisher_PublishBatch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PublishBatch'
type Publisher_PublishBatch_Call struct {
	*mock.Call
}

// PublishBatch is a helper method to define mock.On call
//   - ctx context.Context
//   - typ string
//   - version int
//   - values []interface{}
//   - customAttributes ...map[string]string
func (_e *Publisher_Expecter) PublishBatch(ctx interface{}, typ interface{}, version interface{}, values interface{}, customAttributes ...interface{}) *Publisher_PublishBatch_Call {
	return &Publisher_PublishBatch_Call{Call: _e.mock.On("PublishBatch",
		append([]interface{}{ctx, typ, version, values}, customAttributes...)...)}
}

func (_c *Publisher_PublishBatch_Call) Run(run func(ctx context.Context, typ string, version int, values []interface{}, customAttributes ...map[string]string)) *Publisher_PublishBatch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]map[string]string, len(args)-4)
		for i, a := range args[4:] {
			if a != nil {
				variadicArgs[i] = a.(map[string]string)
			}
		}
		run(args[0].(context.Context), args[1].(string), args[2].(int), args[3].([]interface{}), variadicArgs...)
	})
	return _c
}

func (_c *Publisher_PublishBatch_Call) Return(_a0 error) *Publisher_PublishBatch_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Publisher_PublishBatch_Call) RunAndReturn(run func(context.Context, string, int, []interface{}, ...map[string]string) error) *Publisher_PublishBatch_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewPublisher interface {
	mock.TestingT
	Cleanup(func())
}

// NewPublisher creates a new instance of Publisher. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPublisher(t mockConstructorTestingTNewPublisher) *Publisher {
	mock := &Publisher{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
