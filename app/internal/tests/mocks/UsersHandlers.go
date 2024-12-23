// Code generated by mockery v2.47.0. DO NOT EDIT.

package mocks

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	mock "github.com/stretchr/testify/mock"
)

// UsersHandlers is an autogenerated mock type for the UsersHandlers type
type UsersHandlers struct {
	mock.Mock
}

type UsersHandlers_Expecter struct {
	mock *mock.Mock
}

func (_m *UsersHandlers) EXPECT() *UsersHandlers_Expecter {
	return &UsersHandlers_Expecter{mock: &_m.Mock}
}

// PostUser provides a mock function with given fields: update
func (_m *UsersHandlers) PostUser(update tgbotapi.Update) error {
	ret := _m.Called(update)

	if len(ret) == 0 {
		panic("no return value specified for PostUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(tgbotapi.Update) error); ok {
		r0 = rf(update)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UsersHandlers_PostUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PostUser'
type UsersHandlers_PostUser_Call struct {
	*mock.Call
}

// PostUser is a helper method to define mock.On call
//   - update tgbotapi.Update
func (_e *UsersHandlers_Expecter) PostUser(update interface{}) *UsersHandlers_PostUser_Call {
	return &UsersHandlers_PostUser_Call{Call: _e.mock.On("PostUser", update)}
}

func (_c *UsersHandlers_PostUser_Call) Run(run func(update tgbotapi.Update)) *UsersHandlers_PostUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(tgbotapi.Update))
	})
	return _c
}

func (_c *UsersHandlers_PostUser_Call) Return(_a0 error) *UsersHandlers_PostUser_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *UsersHandlers_PostUser_Call) RunAndReturn(run func(tgbotapi.Update) error) *UsersHandlers_PostUser_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateUserCurrency provides a mock function with given fields: update
func (_m *UsersHandlers) UpdateUserCurrency(update tgbotapi.Update) error {
	ret := _m.Called(update)

	if len(ret) == 0 {
		panic("no return value specified for UpdateUserCurrency")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(tgbotapi.Update) error); ok {
		r0 = rf(update)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UsersHandlers_UpdateUserCurrency_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateUserCurrency'
type UsersHandlers_UpdateUserCurrency_Call struct {
	*mock.Call
}

// UpdateUserCurrency is a helper method to define mock.On call
//   - update tgbotapi.Update
func (_e *UsersHandlers_Expecter) UpdateUserCurrency(update interface{}) *UsersHandlers_UpdateUserCurrency_Call {
	return &UsersHandlers_UpdateUserCurrency_Call{Call: _e.mock.On("UpdateUserCurrency", update)}
}

func (_c *UsersHandlers_UpdateUserCurrency_Call) Run(run func(update tgbotapi.Update)) *UsersHandlers_UpdateUserCurrency_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(tgbotapi.Update))
	})
	return _c
}

func (_c *UsersHandlers_UpdateUserCurrency_Call) Return(_a0 error) *UsersHandlers_UpdateUserCurrency_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *UsersHandlers_UpdateUserCurrency_Call) RunAndReturn(run func(tgbotapi.Update) error) *UsersHandlers_UpdateUserCurrency_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateUserName provides a mock function with given fields: update
func (_m *UsersHandlers) UpdateUserName(update tgbotapi.Update) error {
	ret := _m.Called(update)

	if len(ret) == 0 {
		panic("no return value specified for UpdateUserName")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(tgbotapi.Update) error); ok {
		r0 = rf(update)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UsersHandlers_UpdateUserName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateUserName'
type UsersHandlers_UpdateUserName_Call struct {
	*mock.Call
}

// UpdateUserName is a helper method to define mock.On call
//   - update tgbotapi.Update
func (_e *UsersHandlers_Expecter) UpdateUserName(update interface{}) *UsersHandlers_UpdateUserName_Call {
	return &UsersHandlers_UpdateUserName_Call{Call: _e.mock.On("UpdateUserName", update)}
}

func (_c *UsersHandlers_UpdateUserName_Call) Run(run func(update tgbotapi.Update)) *UsersHandlers_UpdateUserName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(tgbotapi.Update))
	})
	return _c
}

func (_c *UsersHandlers_UpdateUserName_Call) Return(_a0 error) *UsersHandlers_UpdateUserName_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *UsersHandlers_UpdateUserName_Call) RunAndReturn(run func(tgbotapi.Update) error) *UsersHandlers_UpdateUserName_Call {
	_c.Call.Return(run)
	return _c
}

// NewUsersHandlers creates a new instance of UsersHandlers. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUsersHandlers(t interface {
	mock.TestingT
	Cleanup(func())
}) *UsersHandlers {
	mock := &UsersHandlers{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
