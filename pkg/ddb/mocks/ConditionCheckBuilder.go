// Code generated by mockery v2.46.0. DO NOT EDIT.

package mocks

import (
	expression "github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	ddb "github.com/justtrackio/gosoline/pkg/ddb"
	mock "github.com/stretchr/testify/mock"

	types "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// ConditionCheckBuilder is an autogenerated mock type for the ConditionCheckBuilder type
type ConditionCheckBuilder struct {
	mock.Mock
}

type ConditionCheckBuilder_Expecter struct {
	mock *mock.Mock
}

func (_m *ConditionCheckBuilder) EXPECT() *ConditionCheckBuilder_Expecter {
	return &ConditionCheckBuilder_Expecter{mock: &_m.Mock}
}

// Build provides a mock function with given fields: result
func (_m *ConditionCheckBuilder) Build(result interface{}) (*types.ConditionCheck, error) {
	ret := _m.Called(result)

	if len(ret) == 0 {
		panic("no return value specified for Build")
	}

	var r0 *types.ConditionCheck
	var r1 error
	if rf, ok := ret.Get(0).(func(interface{}) (*types.ConditionCheck, error)); ok {
		return rf(result)
	}
	if rf, ok := ret.Get(0).(func(interface{}) *types.ConditionCheck); ok {
		r0 = rf(result)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.ConditionCheck)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}) error); ok {
		r1 = rf(result)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ConditionCheckBuilder_Build_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Build'
type ConditionCheckBuilder_Build_Call struct {
	*mock.Call
}

// Build is a helper method to define mock.On call
//   - result interface{}
func (_e *ConditionCheckBuilder_Expecter) Build(result interface{}) *ConditionCheckBuilder_Build_Call {
	return &ConditionCheckBuilder_Build_Call{Call: _e.mock.On("Build", result)}
}

func (_c *ConditionCheckBuilder_Build_Call) Run(run func(result interface{})) *ConditionCheckBuilder_Build_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *ConditionCheckBuilder_Build_Call) Return(_a0 *types.ConditionCheck, _a1 error) *ConditionCheckBuilder_Build_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ConditionCheckBuilder_Build_Call) RunAndReturn(run func(interface{}) (*types.ConditionCheck, error)) *ConditionCheckBuilder_Build_Call {
	_c.Call.Return(run)
	return _c
}

// ReturnAllOld provides a mock function with given fields:
func (_m *ConditionCheckBuilder) ReturnAllOld() ddb.ConditionCheckBuilder {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ReturnAllOld")
	}

	var r0 ddb.ConditionCheckBuilder
	if rf, ok := ret.Get(0).(func() ddb.ConditionCheckBuilder); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ddb.ConditionCheckBuilder)
		}
	}

	return r0
}

// ConditionCheckBuilder_ReturnAllOld_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReturnAllOld'
type ConditionCheckBuilder_ReturnAllOld_Call struct {
	*mock.Call
}

// ReturnAllOld is a helper method to define mock.On call
func (_e *ConditionCheckBuilder_Expecter) ReturnAllOld() *ConditionCheckBuilder_ReturnAllOld_Call {
	return &ConditionCheckBuilder_ReturnAllOld_Call{Call: _e.mock.On("ReturnAllOld")}
}

func (_c *ConditionCheckBuilder_ReturnAllOld_Call) Run(run func()) *ConditionCheckBuilder_ReturnAllOld_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ConditionCheckBuilder_ReturnAllOld_Call) Return(_a0 ddb.ConditionCheckBuilder) *ConditionCheckBuilder_ReturnAllOld_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ConditionCheckBuilder_ReturnAllOld_Call) RunAndReturn(run func() ddb.ConditionCheckBuilder) *ConditionCheckBuilder_ReturnAllOld_Call {
	_c.Call.Return(run)
	return _c
}

// ReturnNone provides a mock function with given fields:
func (_m *ConditionCheckBuilder) ReturnNone() ddb.ConditionCheckBuilder {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ReturnNone")
	}

	var r0 ddb.ConditionCheckBuilder
	if rf, ok := ret.Get(0).(func() ddb.ConditionCheckBuilder); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ddb.ConditionCheckBuilder)
		}
	}

	return r0
}

// ConditionCheckBuilder_ReturnNone_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReturnNone'
type ConditionCheckBuilder_ReturnNone_Call struct {
	*mock.Call
}

// ReturnNone is a helper method to define mock.On call
func (_e *ConditionCheckBuilder_Expecter) ReturnNone() *ConditionCheckBuilder_ReturnNone_Call {
	return &ConditionCheckBuilder_ReturnNone_Call{Call: _e.mock.On("ReturnNone")}
}

func (_c *ConditionCheckBuilder_ReturnNone_Call) Run(run func()) *ConditionCheckBuilder_ReturnNone_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ConditionCheckBuilder_ReturnNone_Call) Return(_a0 ddb.ConditionCheckBuilder) *ConditionCheckBuilder_ReturnNone_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ConditionCheckBuilder_ReturnNone_Call) RunAndReturn(run func() ddb.ConditionCheckBuilder) *ConditionCheckBuilder_ReturnNone_Call {
	_c.Call.Return(run)
	return _c
}

// WithCondition provides a mock function with given fields: cond
func (_m *ConditionCheckBuilder) WithCondition(cond expression.ConditionBuilder) ddb.ConditionCheckBuilder {
	ret := _m.Called(cond)

	if len(ret) == 0 {
		panic("no return value specified for WithCondition")
	}

	var r0 ddb.ConditionCheckBuilder
	if rf, ok := ret.Get(0).(func(expression.ConditionBuilder) ddb.ConditionCheckBuilder); ok {
		r0 = rf(cond)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ddb.ConditionCheckBuilder)
		}
	}

	return r0
}

// ConditionCheckBuilder_WithCondition_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithCondition'
type ConditionCheckBuilder_WithCondition_Call struct {
	*mock.Call
}

// WithCondition is a helper method to define mock.On call
//   - cond expression.ConditionBuilder
func (_e *ConditionCheckBuilder_Expecter) WithCondition(cond interface{}) *ConditionCheckBuilder_WithCondition_Call {
	return &ConditionCheckBuilder_WithCondition_Call{Call: _e.mock.On("WithCondition", cond)}
}

func (_c *ConditionCheckBuilder_WithCondition_Call) Run(run func(cond expression.ConditionBuilder)) *ConditionCheckBuilder_WithCondition_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(expression.ConditionBuilder))
	})
	return _c
}

func (_c *ConditionCheckBuilder_WithCondition_Call) Return(_a0 ddb.ConditionCheckBuilder) *ConditionCheckBuilder_WithCondition_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ConditionCheckBuilder_WithCondition_Call) RunAndReturn(run func(expression.ConditionBuilder) ddb.ConditionCheckBuilder) *ConditionCheckBuilder_WithCondition_Call {
	_c.Call.Return(run)
	return _c
}

// WithHash provides a mock function with given fields: hashValue
func (_m *ConditionCheckBuilder) WithHash(hashValue interface{}) ddb.ConditionCheckBuilder {
	ret := _m.Called(hashValue)

	if len(ret) == 0 {
		panic("no return value specified for WithHash")
	}

	var r0 ddb.ConditionCheckBuilder
	if rf, ok := ret.Get(0).(func(interface{}) ddb.ConditionCheckBuilder); ok {
		r0 = rf(hashValue)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ddb.ConditionCheckBuilder)
		}
	}

	return r0
}

// ConditionCheckBuilder_WithHash_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithHash'
type ConditionCheckBuilder_WithHash_Call struct {
	*mock.Call
}

// WithHash is a helper method to define mock.On call
//   - hashValue interface{}
func (_e *ConditionCheckBuilder_Expecter) WithHash(hashValue interface{}) *ConditionCheckBuilder_WithHash_Call {
	return &ConditionCheckBuilder_WithHash_Call{Call: _e.mock.On("WithHash", hashValue)}
}

func (_c *ConditionCheckBuilder_WithHash_Call) Run(run func(hashValue interface{})) *ConditionCheckBuilder_WithHash_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *ConditionCheckBuilder_WithHash_Call) Return(_a0 ddb.ConditionCheckBuilder) *ConditionCheckBuilder_WithHash_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ConditionCheckBuilder_WithHash_Call) RunAndReturn(run func(interface{}) ddb.ConditionCheckBuilder) *ConditionCheckBuilder_WithHash_Call {
	_c.Call.Return(run)
	return _c
}

// WithKeys provides a mock function with given fields: keys
func (_m *ConditionCheckBuilder) WithKeys(keys ...interface{}) ddb.ConditionCheckBuilder {
	var _ca []interface{}
	_ca = append(_ca, keys...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for WithKeys")
	}

	var r0 ddb.ConditionCheckBuilder
	if rf, ok := ret.Get(0).(func(...interface{}) ddb.ConditionCheckBuilder); ok {
		r0 = rf(keys...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ddb.ConditionCheckBuilder)
		}
	}

	return r0
}

// ConditionCheckBuilder_WithKeys_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithKeys'
type ConditionCheckBuilder_WithKeys_Call struct {
	*mock.Call
}

// WithKeys is a helper method to define mock.On call
//   - keys ...interface{}
func (_e *ConditionCheckBuilder_Expecter) WithKeys(keys ...interface{}) *ConditionCheckBuilder_WithKeys_Call {
	return &ConditionCheckBuilder_WithKeys_Call{Call: _e.mock.On("WithKeys",
		append([]interface{}{}, keys...)...)}
}

func (_c *ConditionCheckBuilder_WithKeys_Call) Run(run func(keys ...interface{})) *ConditionCheckBuilder_WithKeys_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *ConditionCheckBuilder_WithKeys_Call) Return(_a0 ddb.ConditionCheckBuilder) *ConditionCheckBuilder_WithKeys_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ConditionCheckBuilder_WithKeys_Call) RunAndReturn(run func(...interface{}) ddb.ConditionCheckBuilder) *ConditionCheckBuilder_WithKeys_Call {
	_c.Call.Return(run)
	return _c
}

// WithRange provides a mock function with given fields: rangeValue
func (_m *ConditionCheckBuilder) WithRange(rangeValue interface{}) ddb.ConditionCheckBuilder {
	ret := _m.Called(rangeValue)

	if len(ret) == 0 {
		panic("no return value specified for WithRange")
	}

	var r0 ddb.ConditionCheckBuilder
	if rf, ok := ret.Get(0).(func(interface{}) ddb.ConditionCheckBuilder); ok {
		r0 = rf(rangeValue)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ddb.ConditionCheckBuilder)
		}
	}

	return r0
}

// ConditionCheckBuilder_WithRange_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithRange'
type ConditionCheckBuilder_WithRange_Call struct {
	*mock.Call
}

// WithRange is a helper method to define mock.On call
//   - rangeValue interface{}
func (_e *ConditionCheckBuilder_Expecter) WithRange(rangeValue interface{}) *ConditionCheckBuilder_WithRange_Call {
	return &ConditionCheckBuilder_WithRange_Call{Call: _e.mock.On("WithRange", rangeValue)}
}

func (_c *ConditionCheckBuilder_WithRange_Call) Run(run func(rangeValue interface{})) *ConditionCheckBuilder_WithRange_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *ConditionCheckBuilder_WithRange_Call) Return(_a0 ddb.ConditionCheckBuilder) *ConditionCheckBuilder_WithRange_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ConditionCheckBuilder_WithRange_Call) RunAndReturn(run func(interface{}) ddb.ConditionCheckBuilder) *ConditionCheckBuilder_WithRange_Call {
	_c.Call.Return(run)
	return _c
}

// NewConditionCheckBuilder creates a new instance of ConditionCheckBuilder. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewConditionCheckBuilder(t interface {
	mock.TestingT
	Cleanup(func())
}) *ConditionCheckBuilder {
	mock := &ConditionCheckBuilder{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
