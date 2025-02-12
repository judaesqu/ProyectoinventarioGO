// Code generated by mockery v2.43.2. DO NOT EDIT.

package service

import (
	context "context"

	models "github.com/judaesqu/ProyectoinventarioGO/internal/models"
	mock "github.com/stretchr/testify/mock"
)

// MockService is an autogenerated mock type for the Service type
type MockService struct {
	mock.Mock
}

// AddProduct provides a mock function with given fields: ctx, product, userEmail
func (_m *MockService) AddProduct(ctx context.Context, product models.Product, userEmail string) error {
	ret := _m.Called(ctx, product, userEmail)

	if len(ret) == 0 {
		panic("no return value specified for AddProduct")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.Product, string) error); ok {
		r0 = rf(ctx, product, userEmail)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddUserRole provides a mock function with given fields: ctx, userID, roleID
func (_m *MockService) AddUserRole(ctx context.Context, userID int64, roleID int64) error {
	ret := _m.Called(ctx, userID, roleID)

	if len(ret) == 0 {
		panic("no return value specified for AddUserRole")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64) error); ok {
		r0 = rf(ctx, userID, roleID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetProduct provides a mock function with given fields: ctx, id
func (_m *MockService) GetProduct(ctx context.Context, id int64) (*models.Product, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetProduct")
	}

	var r0 *models.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (*models.Product, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) *models.Product); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Product)
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
func (_m *MockService) GetProducts(ctx context.Context) ([]models.Product, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetProducts")
	}

	var r0 []models.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]models.Product, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []models.Product); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Product)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LoginUser provides a mock function with given fields: ctx, email, password
func (_m *MockService) LoginUser(ctx context.Context, email string, password string) (*models.User, error) {
	ret := _m.Called(ctx, email, password)

	if len(ret) == 0 {
		panic("no return value specified for LoginUser")
	}

	var r0 *models.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*models.User, error)); ok {
		return rf(ctx, email, password)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *models.User); ok {
		r0 = rf(ctx, email, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, email, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterUser provides a mock function with given fields: ctx, email, name, password
func (_m *MockService) RegisterUser(ctx context.Context, email string, name string, password string) error {
	ret := _m.Called(ctx, email, name, password)

	if len(ret) == 0 {
		panic("no return value specified for RegisterUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) error); ok {
		r0 = rf(ctx, email, name, password)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveUserRole provides a mock function with given fields: ctx, userID, roleID
func (_m *MockService) RemoveUserRole(ctx context.Context, userID int64, roleID int64) error {
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

// NewMockService creates a new instance of MockService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockService {
	mock := &MockService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
