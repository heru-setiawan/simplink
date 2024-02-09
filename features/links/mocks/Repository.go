// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	context "context"
	core "simplink/features/links/core"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, data
func (_m *Repository) Create(ctx context.Context, data core.Link) error {
	ret := _m.Called(ctx, data)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, core.Link) error); ok {
		r0 = rf(ctx, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetByShort provides a mock function with given fields: ctx, short
func (_m *Repository) GetByShort(ctx context.Context, short string) (*core.Link, error) {
	ret := _m.Called(ctx, short)

	if len(ret) == 0 {
		panic("no return value specified for GetByShort")
	}

	var r0 *core.Link
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*core.Link, error)); ok {
		return rf(ctx, short)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *core.Link); ok {
		r0 = rf(ctx, short)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.Link)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, short)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
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
