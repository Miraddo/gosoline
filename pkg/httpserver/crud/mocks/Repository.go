// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	db_repo "github.com/justtrackio/gosoline/pkg/db-repo"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

type Repository_Expecter struct {
	mock *mock.Mock
}

func (_m *Repository) EXPECT() *Repository_Expecter {
	return &Repository_Expecter{mock: &_m.Mock}
}

// Count provides a mock function with given fields: ctx, qb, model
func (_m *Repository) Count(ctx context.Context, qb *db_repo.QueryBuilder, model db_repo.ModelBased) (int, error) {
	ret := _m.Called(ctx, qb, model)

	if len(ret) == 0 {
		panic("no return value specified for Count")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *db_repo.QueryBuilder, db_repo.ModelBased) (int, error)); ok {
		return rf(ctx, qb, model)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *db_repo.QueryBuilder, db_repo.ModelBased) int); ok {
		r0 = rf(ctx, qb, model)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *db_repo.QueryBuilder, db_repo.ModelBased) error); ok {
		r1 = rf(ctx, qb, model)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Repository_Count_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Count'
type Repository_Count_Call struct {
	*mock.Call
}

// Count is a helper method to define mock.On call
//   - ctx context.Context
//   - qb *db_repo.QueryBuilder
//   - model db_repo.ModelBased
func (_e *Repository_Expecter) Count(ctx interface{}, qb interface{}, model interface{}) *Repository_Count_Call {
	return &Repository_Count_Call{Call: _e.mock.On("Count", ctx, qb, model)}
}

func (_c *Repository_Count_Call) Run(run func(ctx context.Context, qb *db_repo.QueryBuilder, model db_repo.ModelBased)) *Repository_Count_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*db_repo.QueryBuilder), args[2].(db_repo.ModelBased))
	})
	return _c
}

func (_c *Repository_Count_Call) Return(_a0 int, _a1 error) *Repository_Count_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Repository_Count_Call) RunAndReturn(run func(context.Context, *db_repo.QueryBuilder, db_repo.ModelBased) (int, error)) *Repository_Count_Call {
	_c.Call.Return(run)
	return _c
}

// Create provides a mock function with given fields: ctx, value
func (_m *Repository) Create(ctx context.Context, value db_repo.ModelBased) error {
	ret := _m.Called(ctx, value)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, db_repo.ModelBased) error); ok {
		r0 = rf(ctx, value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Repository_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type Repository_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - value db_repo.ModelBased
func (_e *Repository_Expecter) Create(ctx interface{}, value interface{}) *Repository_Create_Call {
	return &Repository_Create_Call{Call: _e.mock.On("Create", ctx, value)}
}

func (_c *Repository_Create_Call) Run(run func(ctx context.Context, value db_repo.ModelBased)) *Repository_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(db_repo.ModelBased))
	})
	return _c
}

func (_c *Repository_Create_Call) Return(_a0 error) *Repository_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Repository_Create_Call) RunAndReturn(run func(context.Context, db_repo.ModelBased) error) *Repository_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, value
func (_m *Repository) Delete(ctx context.Context, value db_repo.ModelBased) error {
	ret := _m.Called(ctx, value)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, db_repo.ModelBased) error); ok {
		r0 = rf(ctx, value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Repository_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type Repository_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - value db_repo.ModelBased
func (_e *Repository_Expecter) Delete(ctx interface{}, value interface{}) *Repository_Delete_Call {
	return &Repository_Delete_Call{Call: _e.mock.On("Delete", ctx, value)}
}

func (_c *Repository_Delete_Call) Run(run func(ctx context.Context, value db_repo.ModelBased)) *Repository_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(db_repo.ModelBased))
	})
	return _c
}

func (_c *Repository_Delete_Call) Return(_a0 error) *Repository_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Repository_Delete_Call) RunAndReturn(run func(context.Context, db_repo.ModelBased) error) *Repository_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// GetMetadata provides a mock function with given fields:
func (_m *Repository) GetMetadata() db_repo.Metadata {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetMetadata")
	}

	var r0 db_repo.Metadata
	if rf, ok := ret.Get(0).(func() db_repo.Metadata); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(db_repo.Metadata)
	}

	return r0
}

// Repository_GetMetadata_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetMetadata'
type Repository_GetMetadata_Call struct {
	*mock.Call
}

// GetMetadata is a helper method to define mock.On call
func (_e *Repository_Expecter) GetMetadata() *Repository_GetMetadata_Call {
	return &Repository_GetMetadata_Call{Call: _e.mock.On("GetMetadata")}
}

func (_c *Repository_GetMetadata_Call) Run(run func()) *Repository_GetMetadata_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Repository_GetMetadata_Call) Return(_a0 db_repo.Metadata) *Repository_GetMetadata_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Repository_GetMetadata_Call) RunAndReturn(run func() db_repo.Metadata) *Repository_GetMetadata_Call {
	_c.Call.Return(run)
	return _c
}

// Query provides a mock function with given fields: ctx, qb, result
func (_m *Repository) Query(ctx context.Context, qb *db_repo.QueryBuilder, result interface{}) error {
	ret := _m.Called(ctx, qb, result)

	if len(ret) == 0 {
		panic("no return value specified for Query")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *db_repo.QueryBuilder, interface{}) error); ok {
		r0 = rf(ctx, qb, result)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Repository_Query_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Query'
type Repository_Query_Call struct {
	*mock.Call
}

// Query is a helper method to define mock.On call
//   - ctx context.Context
//   - qb *db_repo.QueryBuilder
//   - result interface{}
func (_e *Repository_Expecter) Query(ctx interface{}, qb interface{}, result interface{}) *Repository_Query_Call {
	return &Repository_Query_Call{Call: _e.mock.On("Query", ctx, qb, result)}
}

func (_c *Repository_Query_Call) Run(run func(ctx context.Context, qb *db_repo.QueryBuilder, result interface{})) *Repository_Query_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*db_repo.QueryBuilder), args[2].(interface{}))
	})
	return _c
}

func (_c *Repository_Query_Call) Return(_a0 error) *Repository_Query_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Repository_Query_Call) RunAndReturn(run func(context.Context, *db_repo.QueryBuilder, interface{}) error) *Repository_Query_Call {
	_c.Call.Return(run)
	return _c
}

// Read provides a mock function with given fields: ctx, id, out
func (_m *Repository) Read(ctx context.Context, id *uint, out db_repo.ModelBased) error {
	ret := _m.Called(ctx, id, out)

	if len(ret) == 0 {
		panic("no return value specified for Read")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *uint, db_repo.ModelBased) error); ok {
		r0 = rf(ctx, id, out)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Repository_Read_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Read'
type Repository_Read_Call struct {
	*mock.Call
}

// Read is a helper method to define mock.On call
//   - ctx context.Context
//   - id *uint
//   - out db_repo.ModelBased
func (_e *Repository_Expecter) Read(ctx interface{}, id interface{}, out interface{}) *Repository_Read_Call {
	return &Repository_Read_Call{Call: _e.mock.On("Read", ctx, id, out)}
}

func (_c *Repository_Read_Call) Run(run func(ctx context.Context, id *uint, out db_repo.ModelBased)) *Repository_Read_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*uint), args[2].(db_repo.ModelBased))
	})
	return _c
}

func (_c *Repository_Read_Call) Return(_a0 error) *Repository_Read_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Repository_Read_Call) RunAndReturn(run func(context.Context, *uint, db_repo.ModelBased) error) *Repository_Read_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, value
func (_m *Repository) Update(ctx context.Context, value db_repo.ModelBased) error {
	ret := _m.Called(ctx, value)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, db_repo.ModelBased) error); ok {
		r0 = rf(ctx, value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Repository_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type Repository_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - value db_repo.ModelBased
func (_e *Repository_Expecter) Update(ctx interface{}, value interface{}) *Repository_Update_Call {
	return &Repository_Update_Call{Call: _e.mock.On("Update", ctx, value)}
}

func (_c *Repository_Update_Call) Run(run func(ctx context.Context, value db_repo.ModelBased)) *Repository_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(db_repo.ModelBased))
	})
	return _c
}

func (_c *Repository_Update_Call) Return(_a0 error) *Repository_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Repository_Update_Call) RunAndReturn(run func(context.Context, db_repo.ModelBased) error) *Repository_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
