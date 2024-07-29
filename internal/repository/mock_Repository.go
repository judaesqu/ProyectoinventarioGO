// Code generated by mockery v2.43.2. DO NOT EDIT.

package repository

import (
	context "context"

	entity "github.com/judaesqu/ProyectoinventarioGO/internal/entity"
	mock "github.com/stretchr/testify/mock"
)

// MockRepository is an autogenerated mock type for the Repository type
type MockRepository struct {
	mock.Mock
}

// GetProduct provides a mock function with given fields: ctx, id
func (_m *MockRepository) GetProduct(ctx context.Context, id int64) (*entity.Product, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetProduct")
	}

	var r0 *entity.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (*entity.Product, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) *entity.Product); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Product)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProducts provides a mock function with given fields: ctx
func (_m *MockRepository) GetProducts(ctx context.Context) ([]entity.Product, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetProducts")
	}

	var r0 []entity.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]entity.Product, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []entity.Product); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Product)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByEmail provides a mock function with given fields: ctx, email
func (_m *MockRepository) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	ret := _m.Called(ctx, email)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByEmail")
	}

	var r0 *entity.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*entity.User, error)); ok {
		return rf(ctx, email)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *entity.User); ok {
		r0 = rf(ctx, email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserRoles provides a mock function with given fields: ctx, userID
func (_m *MockRepository) GetUserRoles(ctx context.Context, userID int64) ([]entity.UserRole, error) {
	ret := _m.Called(ctx, userID)

	if len(ret) == 0 {
		panic("no return value specified for GetUserRoles")
	}

	var r0 []entity.UserRole
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) ([]entity.UserRole, error)); ok {
		return rf(ctx, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) []entity.UserRole); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.UserRole)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveUserRole provides a mock function with given fields: ctx, userID, roleID
func (_m *MockRepository) RemoveUserRole(ctx context.Context, userID int64, roleID int64) error {
	ret := _m.Called(ctx, userID, roleID)

	if len(ret) == 0 {
		panic("no return value specified for RemoveUserRole")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64) error); ok {
		r0 = rf(ctx, userID, roleID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveProduct provides a mock function with given fields: ctx, name, description, price, createdBy
func (_m *MockRepository) SaveProduct(ctx context.Context, name string, description string, price float32, createdBy int64) error {
	ret := _m.Called(ctx, name, description, price, createdBy)

	if len(ret) == 0 {
		panic("no return value specified for SaveProduct")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, float32, int64) error); ok {
		r0 = rf(ctx, name, description, price, createdBy)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveUser provides a mock function with given fields: ctx, email, name, password
func (_m *MockRepository) SaveUser(ctx context.Context, email string, name string, password string) error {
	ret := _m.Called(ctx, email, name, password)

	if len(ret) == 0 {
		panic("no return value specified for SaveUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) error); ok {
		r0 = rf(ctx, email, name, password)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveUserRole provides a mock function with given fields: ctx, userID, roleID
func (_m *MockRepository) SaveUserRole(ctx context.Context, userID int64, roleID int64) error {
	ret := _m.Called(ctx, userID, roleID)

	if len(ret) == 0 {
		panic("no return value specified for SaveUserRole")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64) error); ok {
		r0 = rf(ctx, userID, roleID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockRepository creates a new instance of MockRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockRepository {
	mock := &MockRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
