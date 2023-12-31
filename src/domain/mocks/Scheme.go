// Code generated by mockery v2.28.1. DO NOT EDIT.

package mocks

import (
	domain "github.com/apm-dev/eth_getBalance-proxy/src/domain"
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// Scheme is an autogenerated mock type for the Scheme type
type Scheme struct {
	mock.Mock
}

// IsCacheSupported provides a mock function with given fields: req
func (_m *Scheme) IsCacheSupported(req *domain.JsonRpcRequest) (time.Duration, bool) {
	ret := _m.Called(req)

	var r0 time.Duration
	var r1 bool
	if rf, ok := ret.Get(0).(func(*domain.JsonRpcRequest) (time.Duration, bool)); ok {
		return rf(req)
	}
	if rf, ok := ret.Get(0).(func(*domain.JsonRpcRequest) time.Duration); ok {
		r0 = rf(req)
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	if rf, ok := ret.Get(1).(func(*domain.JsonRpcRequest) bool); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// IsJsonRpcResponseValid provides a mock function with given fields: resp
func (_m *Scheme) IsJsonRpcResponseValid(resp *domain.JsonRpcResponse) bool {
	ret := _m.Called(resp)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*domain.JsonRpcResponse) bool); ok {
		r0 = rf(resp)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IsSupportedRpcMethod provides a mock function with given fields: method
func (_m *Scheme) IsSupportedRpcMethod(method string) bool {
	ret := _m.Called(method)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(method)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ParseRequest provides a mock function with given fields: req
func (_m *Scheme) ParseRequest(req *domain.JsonRpcRequest) (*domain.JsonRpcResponse, error) {
	ret := _m.Called(req)

	var r0 *domain.JsonRpcResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(*domain.JsonRpcRequest) (*domain.JsonRpcResponse, error)); ok {
		return rf(req)
	}
	if rf, ok := ret.Get(0).(func(*domain.JsonRpcRequest) *domain.JsonRpcResponse); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.JsonRpcResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(*domain.JsonRpcRequest) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewScheme interface {
	mock.TestingT
	Cleanup(func())
}

// NewScheme creates a new instance of Scheme. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewScheme(t mockConstructorTestingTNewScheme) *Scheme {
	mock := &Scheme{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
