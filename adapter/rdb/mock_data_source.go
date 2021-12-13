// Code generated by mockery v2.9.4. DO NOT EDIT.

package rdb

import (
	usecase "github.com/mkaiho/go-todo-sample/usecase"
	mock "github.com/stretchr/testify/mock"
)

// MockDataSource is an autogenerated mock type for the DataSource type
type MockDataSource struct {
	mock.Mock
}

// Execute provides a mock function with given fields: sql, args
func (_m *MockDataSource) Execute(sql string, args ...interface{}) usecase.UseCaseError {
	var _ca []interface{}
	_ca = append(_ca, sql)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 usecase.UseCaseError
	if rf, ok := ret.Get(0).(func(string, ...interface{}) usecase.UseCaseError); ok {
		r0 = rf(sql, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(usecase.UseCaseError)
		}
	}

	return r0
}

// Query provides a mock function with given fields: dest, query
func (_m *MockDataSource) Query(dest interface{}, query Query) usecase.UseCaseError {
	ret := _m.Called(dest, query)

	var r0 usecase.UseCaseError
	if rf, ok := ret.Get(0).(func(interface{}, Query) usecase.UseCaseError); ok {
		r0 = rf(dest, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(usecase.UseCaseError)
		}
	}

	return r0
}
