// Code generated by mockery v2.22.1. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	oauth2 "google.golang.org/api/oauth2/v2"
)

// TokenInfoProvider is an autogenerated mock type for the TokenInfoProvider type
type TokenInfoProvider struct {
	mock.Mock
}

type TokenInfoProvider_Expecter struct {
	mock *mock.Mock
}

func (_m *TokenInfoProvider) EXPECT() *TokenInfoProvider_Expecter {
	return &TokenInfoProvider_Expecter{mock: &_m.Mock}
}

// GetTokenInfo provides a mock function with given fields: _a0
func (_m *TokenInfoProvider) GetTokenInfo(_a0 string) (*oauth2.Tokeninfo, error) {
	ret := _m.Called(_a0)

	var r0 *oauth2.Tokeninfo
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*oauth2.Tokeninfo, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(string) *oauth2.Tokeninfo); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*oauth2.Tokeninfo)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TokenInfoProvider_GetTokenInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTokenInfo'
type TokenInfoProvider_GetTokenInfo_Call struct {
	*mock.Call
}

// GetTokenInfo is a helper method to define mock.On call
//   - _a0 string
func (_e *TokenInfoProvider_Expecter) GetTokenInfo(_a0 interface{}) *TokenInfoProvider_GetTokenInfo_Call {
	return &TokenInfoProvider_GetTokenInfo_Call{Call: _e.mock.On("GetTokenInfo", _a0)}
}

func (_c *TokenInfoProvider_GetTokenInfo_Call) Run(run func(_a0 string)) *TokenInfoProvider_GetTokenInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *TokenInfoProvider_GetTokenInfo_Call) Return(_a0 *oauth2.Tokeninfo, _a1 error) *TokenInfoProvider_GetTokenInfo_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *TokenInfoProvider_GetTokenInfo_Call) RunAndReturn(run func(string) (*oauth2.Tokeninfo, error)) *TokenInfoProvider_GetTokenInfo_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewTokenInfoProvider interface {
	mock.TestingT
	Cleanup(func())
}

// NewTokenInfoProvider creates a new instance of TokenInfoProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTokenInfoProvider(t mockConstructorTestingTNewTokenInfoProvider) *TokenInfoProvider {
	mock := &TokenInfoProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}