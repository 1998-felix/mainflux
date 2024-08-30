// Code generated by mockery v2.43.2. DO NOT EDIT.

// Copyright (c) Abstract Machines

package mocks

import (
	context "context"

	policy "github.com/absmach/magistrala/pkg/policy"
	mock "github.com/stretchr/testify/mock"
)

// PolicyAgent is an autogenerated mock type for the PolicyAgent type
type PolicyAgent struct {
	mock.Mock
}

// AddPolicies provides a mock function with given fields: ctx, prs
func (_m *PolicyAgent) AddPolicies(ctx context.Context, prs []policy.PolicyReq) error {
	ret := _m.Called(ctx, prs)

	if len(ret) == 0 {
		panic("no return value specified for AddPolicies")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []policy.PolicyReq) error); ok {
		r0 = rf(ctx, prs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddPolicy provides a mock function with given fields: ctx, pr
func (_m *PolicyAgent) AddPolicy(ctx context.Context, pr policy.PolicyReq) error {
	ret := _m.Called(ctx, pr)

	if len(ret) == 0 {
		panic("no return value specified for AddPolicy")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, policy.PolicyReq) error); ok {
		r0 = rf(ctx, pr)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeletePolicies provides a mock function with given fields: ctx, prs
func (_m *PolicyAgent) DeletePolicies(ctx context.Context, prs []policy.PolicyReq) error {
	ret := _m.Called(ctx, prs)

	if len(ret) == 0 {
		panic("no return value specified for DeletePolicies")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []policy.PolicyReq) error); ok {
		r0 = rf(ctx, prs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeletePolicyFilter provides a mock function with given fields: ctx, pr
func (_m *PolicyAgent) DeletePolicyFilter(ctx context.Context, pr policy.PolicyReq) error {
	ret := _m.Called(ctx, pr)

	if len(ret) == 0 {
		panic("no return value specified for DeletePolicyFilter")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, policy.PolicyReq) error); ok {
		r0 = rf(ctx, pr)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RetrieveAllObjects provides a mock function with given fields: ctx, pr
func (_m *PolicyAgent) RetrieveAllObjects(ctx context.Context, pr policy.PolicyReq) ([]policy.PolicyRes, error) {
	ret := _m.Called(ctx, pr)

	if len(ret) == 0 {
		panic("no return value specified for RetrieveAllObjects")
	}

	var r0 []policy.PolicyRes
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, policy.PolicyReq) ([]policy.PolicyRes, error)); ok {
		return rf(ctx, pr)
	}
	if rf, ok := ret.Get(0).(func(context.Context, policy.PolicyReq) []policy.PolicyRes); ok {
		r0 = rf(ctx, pr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]policy.PolicyRes)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, policy.PolicyReq) error); ok {
		r1 = rf(ctx, pr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RetrieveAllObjectsCount provides a mock function with given fields: ctx, pr
func (_m *PolicyAgent) RetrieveAllObjectsCount(ctx context.Context, pr policy.PolicyReq) (uint64, error) {
	ret := _m.Called(ctx, pr)

	if len(ret) == 0 {
		panic("no return value specified for RetrieveAllObjectsCount")
	}

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, policy.PolicyReq) (uint64, error)); ok {
		return rf(ctx, pr)
	}
	if rf, ok := ret.Get(0).(func(context.Context, policy.PolicyReq) uint64); ok {
		r0 = rf(ctx, pr)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, policy.PolicyReq) error); ok {
		r1 = rf(ctx, pr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RetrieveAllSubjects provides a mock function with given fields: ctx, pr
func (_m *PolicyAgent) RetrieveAllSubjects(ctx context.Context, pr policy.PolicyReq) ([]policy.PolicyRes, error) {
	ret := _m.Called(ctx, pr)

	if len(ret) == 0 {
		panic("no return value specified for RetrieveAllSubjects")
	}

	var r0 []policy.PolicyRes
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, policy.PolicyReq) ([]policy.PolicyRes, error)); ok {
		return rf(ctx, pr)
	}
	if rf, ok := ret.Get(0).(func(context.Context, policy.PolicyReq) []policy.PolicyRes); ok {
		r0 = rf(ctx, pr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]policy.PolicyRes)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, policy.PolicyReq) error); ok {
		r1 = rf(ctx, pr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RetrieveAllSubjectsCount provides a mock function with given fields: ctx, pr
func (_m *PolicyAgent) RetrieveAllSubjectsCount(ctx context.Context, pr policy.PolicyReq) (uint64, error) {
	ret := _m.Called(ctx, pr)

	if len(ret) == 0 {
		panic("no return value specified for RetrieveAllSubjectsCount")
	}

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, policy.PolicyReq) (uint64, error)); ok {
		return rf(ctx, pr)
	}
	if rf, ok := ret.Get(0).(func(context.Context, policy.PolicyReq) uint64); ok {
		r0 = rf(ctx, pr)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, policy.PolicyReq) error); ok {
		r1 = rf(ctx, pr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RetrieveObjects provides a mock function with given fields: ctx, pr, nextPageToken, limit
func (_m *PolicyAgent) RetrieveObjects(ctx context.Context, pr policy.PolicyReq, nextPageToken string, limit uint64) ([]policy.PolicyRes, string, error) {
	ret := _m.Called(ctx, pr, nextPageToken, limit)

	if len(ret) == 0 {
		panic("no return value specified for RetrieveObjects")
	}

	var r0 []policy.PolicyRes
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, policy.PolicyReq, string, uint64) ([]policy.PolicyRes, string, error)); ok {
		return rf(ctx, pr, nextPageToken, limit)
	}
	if rf, ok := ret.Get(0).(func(context.Context, policy.PolicyReq, string, uint64) []policy.PolicyRes); ok {
		r0 = rf(ctx, pr, nextPageToken, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]policy.PolicyRes)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, policy.PolicyReq, string, uint64) string); ok {
		r1 = rf(ctx, pr, nextPageToken, limit)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(context.Context, policy.PolicyReq, string, uint64) error); ok {
		r2 = rf(ctx, pr, nextPageToken, limit)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// RetrievePermissions provides a mock function with given fields: ctx, pr, filterPermission
func (_m *PolicyAgent) RetrievePermissions(ctx context.Context, pr policy.PolicyReq, filterPermission []string) (policy.Permissions, error) {
	ret := _m.Called(ctx, pr, filterPermission)

	if len(ret) == 0 {
		panic("no return value specified for RetrievePermissions")
	}

	var r0 policy.Permissions
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, policy.PolicyReq, []string) (policy.Permissions, error)); ok {
		return rf(ctx, pr, filterPermission)
	}
	if rf, ok := ret.Get(0).(func(context.Context, policy.PolicyReq, []string) policy.Permissions); ok {
		r0 = rf(ctx, pr, filterPermission)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(policy.Permissions)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, policy.PolicyReq, []string) error); ok {
		r1 = rf(ctx, pr, filterPermission)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RetrieveSubjects provides a mock function with given fields: ctx, pr, nextPageToken, limit
func (_m *PolicyAgent) RetrieveSubjects(ctx context.Context, pr policy.PolicyReq, nextPageToken string, limit uint64) ([]policy.PolicyRes, string, error) {
	ret := _m.Called(ctx, pr, nextPageToken, limit)

	if len(ret) == 0 {
		panic("no return value specified for RetrieveSubjects")
	}

	var r0 []policy.PolicyRes
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, policy.PolicyReq, string, uint64) ([]policy.PolicyRes, string, error)); ok {
		return rf(ctx, pr, nextPageToken, limit)
	}
	if rf, ok := ret.Get(0).(func(context.Context, policy.PolicyReq, string, uint64) []policy.PolicyRes); ok {
		r0 = rf(ctx, pr, nextPageToken, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]policy.PolicyRes)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, policy.PolicyReq, string, uint64) string); ok {
		r1 = rf(ctx, pr, nextPageToken, limit)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(context.Context, policy.PolicyReq, string, uint64) error); ok {
		r2 = rf(ctx, pr, nextPageToken, limit)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// NewPolicyAgent creates a new instance of PolicyAgent. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPolicyAgent(t interface {
	mock.TestingT
	Cleanup(func())
}) *PolicyAgent {
	mock := &PolicyAgent{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
