// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"
	entity "faceit/domain/user/entity"

	mock "github.com/stretchr/testify/mock"
)

// IUsersRepository is an autogenerated mock type for the IUsersRepository type
type IUsersRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, user
func (_m *IUsersRepository) Create(ctx context.Context, user *entity.User) (*entity.User, error) {
	ret := _m.Called(ctx, user)

	var r0 *entity.User
	if rf, ok := ret.Get(0).(func(context.Context, *entity.User) *entity.User); ok {
		r0 = rf(ctx, user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *entity.User) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: ctx, filter, page, pageSize
func (_m *IUsersRepository) Get(ctx context.Context, filter *entity.Filter, page int64, pageSize int64) ([]*entity.User, error) {
	ret := _m.Called(ctx, filter, page, pageSize)

	var r0 []*entity.User
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Filter, int64, int64) []*entity.User); ok {
		r0 = rf(ctx, filter, page, pageSize)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *entity.Filter, int64, int64) error); ok {
		r1 = rf(ctx, filter, page, pageSize)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByEmail provides a mock function with given fields: ctx, email
func (_m *IUsersRepository) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	ret := _m.Called(ctx, email)

	var r0 *entity.User
	if rf, ok := ret.Get(0).(func(context.Context, string) *entity.User); ok {
		r0 = rf(ctx, email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, ID
func (_m *IUsersRepository) GetByID(ctx context.Context, ID int64) (*entity.User, error) {
	ret := _m.Called(ctx, ID)

	var r0 *entity.User
	if rf, ok := ret.Get(0).(func(context.Context, int64) *entity.User); ok {
		r0 = rf(ctx, ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByNickName provides a mock function with given fields: ctx, nickName
func (_m *IUsersRepository) GetByNickName(ctx context.Context, nickName string) (*entity.User, error) {
	ret := _m.Called(ctx, nickName)

	var r0 *entity.User
	if rf, ok := ret.Get(0).(func(context.Context, string) *entity.User); ok {
		r0 = rf(ctx, nickName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, nickName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCount provides a mock function with given fields: ctx, filter
func (_m *IUsersRepository) GetCount(ctx context.Context, filter *entity.Filter) (uint64, error) {
	ret := _m.Called(ctx, filter)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Filter) uint64); ok {
		r0 = rf(ctx, filter)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *entity.Filter) error); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Remove provides a mock function with given fields: ctx, ID
func (_m *IUsersRepository) Remove(ctx context.Context, ID int64) error {
	ret := _m.Called(ctx, ID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, ID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: ctx, user
func (_m *IUsersRepository) Update(ctx context.Context, user *entity.User) error {
	ret := _m.Called(ctx, user)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.User) error); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewIUsersRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewIUsersRepository creates a new instance of IUsersRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIUsersRepository(t mockConstructorTestingTNewIUsersRepository) *IUsersRepository {
	mock := &IUsersRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
