// Code generated by mockery v2.50.0. DO NOT EDIT.

package mocks

import (
	models "cachManagerApp/app/db/models"

	mock "github.com/stretchr/testify/mock"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// ExpenseAnalyticHandler is an autogenerated mock type for the ExpenseAnalyticHandler type
type ExpenseAnalyticHandler struct {
	mock.Mock
}

type ExpenseAnalyticHandler_Expecter struct {
	mock *mock.Mock
}

func (_m *ExpenseAnalyticHandler) EXPECT() *ExpenseAnalyticHandler_Expecter {
	return &ExpenseAnalyticHandler_Expecter{mock: &_m.Mock}
}

// ExpenseDayAnalytic provides a mock function with given fields: update
func (_m *ExpenseAnalyticHandler) ExpenseDayAnalytic(update tgbotapi.Update) ([]models.Transactions, error) {
	ret := _m.Called(update)

	if len(ret) == 0 {
		panic("no return value specified for ExpenseDayAnalytic")
	}

	var r0 []models.Transactions
	var r1 error
	if rf, ok := ret.Get(0).(func(tgbotapi.Update) ([]models.Transactions, error)); ok {
		return rf(update)
	}
	if rf, ok := ret.Get(0).(func(tgbotapi.Update) []models.Transactions); ok {
		r0 = rf(update)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Transactions)
		}
	}

	if rf, ok := ret.Get(1).(func(tgbotapi.Update) error); ok {
		r1 = rf(update)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExpenseAnalyticHandler_ExpenseDayAnalytic_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ExpenseDayAnalytic'
type ExpenseAnalyticHandler_ExpenseDayAnalytic_Call struct {
	*mock.Call
}

// ExpenseDayAnalytic is a helper method to define mock.On call
//   - update tgbotapi.Update
func (_e *ExpenseAnalyticHandler_Expecter) ExpenseDayAnalytic(update interface{}) *ExpenseAnalyticHandler_ExpenseDayAnalytic_Call {
	return &ExpenseAnalyticHandler_ExpenseDayAnalytic_Call{Call: _e.mock.On("ExpenseDayAnalytic", update)}
}

func (_c *ExpenseAnalyticHandler_ExpenseDayAnalytic_Call) Run(run func(update tgbotapi.Update)) *ExpenseAnalyticHandler_ExpenseDayAnalytic_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(tgbotapi.Update))
	})
	return _c
}

func (_c *ExpenseAnalyticHandler_ExpenseDayAnalytic_Call) Return(_a0 []models.Transactions, _a1 error) *ExpenseAnalyticHandler_ExpenseDayAnalytic_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ExpenseAnalyticHandler_ExpenseDayAnalytic_Call) RunAndReturn(run func(tgbotapi.Update) ([]models.Transactions, error)) *ExpenseAnalyticHandler_ExpenseDayAnalytic_Call {
	_c.Call.Return(run)
	return _c
}

// ExpenseMonthAnalytic provides a mock function with given fields: update
func (_m *ExpenseAnalyticHandler) ExpenseMonthAnalytic(update tgbotapi.Update) (map[string]uint64, error) {
	ret := _m.Called(update)

	if len(ret) == 0 {
		panic("no return value specified for ExpenseMonthAnalytic")
	}

	var r0 map[string]uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(tgbotapi.Update) (map[string]uint64, error)); ok {
		return rf(update)
	}
	if rf, ok := ret.Get(0).(func(tgbotapi.Update) map[string]uint64); ok {
		r0 = rf(update)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]uint64)
		}
	}

	if rf, ok := ret.Get(1).(func(tgbotapi.Update) error); ok {
		r1 = rf(update)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExpenseAnalyticHandler_ExpenseMonthAnalytic_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ExpenseMonthAnalytic'
type ExpenseAnalyticHandler_ExpenseMonthAnalytic_Call struct {
	*mock.Call
}

// ExpenseMonthAnalytic is a helper method to define mock.On call
//   - update tgbotapi.Update
func (_e *ExpenseAnalyticHandler_Expecter) ExpenseMonthAnalytic(update interface{}) *ExpenseAnalyticHandler_ExpenseMonthAnalytic_Call {
	return &ExpenseAnalyticHandler_ExpenseMonthAnalytic_Call{Call: _e.mock.On("ExpenseMonthAnalytic", update)}
}

func (_c *ExpenseAnalyticHandler_ExpenseMonthAnalytic_Call) Run(run func(update tgbotapi.Update)) *ExpenseAnalyticHandler_ExpenseMonthAnalytic_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(tgbotapi.Update))
	})
	return _c
}

func (_c *ExpenseAnalyticHandler_ExpenseMonthAnalytic_Call) Return(_a0 map[string]uint64, _a1 error) *ExpenseAnalyticHandler_ExpenseMonthAnalytic_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ExpenseAnalyticHandler_ExpenseMonthAnalytic_Call) RunAndReturn(run func(tgbotapi.Update) (map[string]uint64, error)) *ExpenseAnalyticHandler_ExpenseMonthAnalytic_Call {
	_c.Call.Return(run)
	return _c
}

// ExpenseWeekAnalytic provides a mock function with given fields: update
func (_m *ExpenseAnalyticHandler) ExpenseWeekAnalytic(update tgbotapi.Update) (map[string]uint64, error) {
	ret := _m.Called(update)

	if len(ret) == 0 {
		panic("no return value specified for ExpenseWeekAnalytic")
	}

	var r0 map[string]uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(tgbotapi.Update) (map[string]uint64, error)); ok {
		return rf(update)
	}
	if rf, ok := ret.Get(0).(func(tgbotapi.Update) map[string]uint64); ok {
		r0 = rf(update)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]uint64)
		}
	}

	if rf, ok := ret.Get(1).(func(tgbotapi.Update) error); ok {
		r1 = rf(update)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExpenseAnalyticHandler_ExpenseWeekAnalytic_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ExpenseWeekAnalytic'
type ExpenseAnalyticHandler_ExpenseWeekAnalytic_Call struct {
	*mock.Call
}

// ExpenseWeekAnalytic is a helper method to define mock.On call
//   - update tgbotapi.Update
func (_e *ExpenseAnalyticHandler_Expecter) ExpenseWeekAnalytic(update interface{}) *ExpenseAnalyticHandler_ExpenseWeekAnalytic_Call {
	return &ExpenseAnalyticHandler_ExpenseWeekAnalytic_Call{Call: _e.mock.On("ExpenseWeekAnalytic", update)}
}

func (_c *ExpenseAnalyticHandler_ExpenseWeekAnalytic_Call) Run(run func(update tgbotapi.Update)) *ExpenseAnalyticHandler_ExpenseWeekAnalytic_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(tgbotapi.Update))
	})
	return _c
}

func (_c *ExpenseAnalyticHandler_ExpenseWeekAnalytic_Call) Return(_a0 map[string]uint64, _a1 error) *ExpenseAnalyticHandler_ExpenseWeekAnalytic_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ExpenseAnalyticHandler_ExpenseWeekAnalytic_Call) RunAndReturn(run func(tgbotapi.Update) (map[string]uint64, error)) *ExpenseAnalyticHandler_ExpenseWeekAnalytic_Call {
	_c.Call.Return(run)
	return _c
}

// NewExpenseAnalyticHandler creates a new instance of ExpenseAnalyticHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewExpenseAnalyticHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *ExpenseAnalyticHandler {
	mock := &ExpenseAnalyticHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
