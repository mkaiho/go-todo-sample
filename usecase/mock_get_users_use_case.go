// Code generated by mockery v2.9.4. DO NOT EDIT.

package usecase

import (
	entity "github.com/mkaiho/go-todo-sample/entity"
	mock "github.com/stretchr/testify/mock"
)

// MockGetUsersUseCase is an autogenerated mock type for the GetUsersUseCase type
type MockGetUsersUseCase struct {
	mock.Mock
}

// Get provides a mock function with given fields: input
func (_m *MockGetUsersUseCase) Get(input GetUsersUseCaseInput) (entity.User, UseCaseError) {
	ret := _m.Called(input)

	var r0 entity.User
	if rf, ok := ret.Get(0).(func(GetUsersUseCaseInput) entity.User); ok {
		r0 = rf(input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(entity.User)
		}
	}

	var r1 UseCaseError
	if rf, ok := ret.Get(1).(func(GetUsersUseCaseInput) UseCaseError); ok {
		r1 = rf(input)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(UseCaseError)
		}
	}

	return r0, r1
}