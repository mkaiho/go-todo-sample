// Code generated by mockery v2.9.4. DO NOT EDIT.

package usecase

import (
	entity "github.com/mkaiho/go-todo-sample/entity"
	mock "github.com/stretchr/testify/mock"
)

// MockUsersDataAccess is an autogenerated mock type for the UsersDataAccess type
type MockUsersDataAccess struct {
	mock.Mock
}

// FindByID provides a mock function with given fields: id
func (_m *MockUsersDataAccess) FindByID(id entity.UserID) (entity.User, UseCaseError) {
	ret := _m.Called(id)

	var r0 entity.User
	if rf, ok := ret.Get(0).(func(entity.UserID) entity.User); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(entity.User)
		}
	}

	var r1 UseCaseError
	if rf, ok := ret.Get(1).(func(entity.UserID) UseCaseError); ok {
		r1 = rf(id)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(UseCaseError)
		}
	}

	return r0, r1
}
