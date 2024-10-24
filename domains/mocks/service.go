// Code generated by mockery v2.43.2. DO NOT EDIT.

// Copyright (c) Abstract Machines

package mocks

import (
	context "context"

	authn "github.com/absmach/magistrala/pkg/authn"

	domains "github.com/absmach/magistrala/domains"

	mock "github.com/stretchr/testify/mock"

	roles "github.com/absmach/magistrala/pkg/roles"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// AddRole provides a mock function with given fields: ctx, session, entityID, roleName, optionalActions, optionalMembers
func (_m *Service) AddRole(ctx context.Context, session authn.Session, entityID string, roleName string, optionalActions []string, optionalMembers []string) (roles.Role, error) {
	ret := _m.Called(ctx, session, entityID, roleName, optionalActions, optionalMembers)

	if len(ret) == 0 {
		panic("no return value specified for AddRole")
	}

	var r0 roles.Role
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, authn.Session, string, string, []string, []string) (roles.Role, error)); ok {
		return rf(ctx, session, entityID, roleName, optionalActions, optionalMembers)
	}
	if rf, ok := ret.Get(0).(func(context.Context, authn.Session, string, string, []string, []string) roles.Role); ok {
		r0 = rf(ctx, session, entityID, roleName, optionalActions, optionalMembers)
	} else {
		r0 = ret.Get(0).(roles.Role)
	}

	if rf, ok := ret.Get(1).(func(context.Context, authn.Session, string, string, []string, []string) error); ok {
		r1 = rf(ctx, session, entityID, roleName, optionalActions, optionalMembers)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateDomain provides a mock function with given fields: ctx, sesssion, d
func (_m *Service) CreateDomain(ctx context.Context, sesssion authn.Session, d domains.Domain) (domains.Domain, error) {
	ret := _m.Called(ctx, sesssion, d)

	if len(ret) == 0 {
		panic("no return value specified for CreateDomain")
	}

	var r0 domains.Domain
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, authn.Session, domains.Domain) (domains.Domain, error)); ok {
		return rf(ctx, sesssion, d)
	}
	if rf, ok := ret.Get(0).(func(context.Context, authn.Session, domains.Domain) domains.Domain); ok {
		r0 = rf(ctx, sesssion, d)
	} else {
		r0 = ret.Get(0).(domains.Domain)
	}

	if rf, ok := ret.Get(1).(func(context.Context, authn.Session, domains.Domain) error); ok {
		r1 = rf(ctx, sesssion, d)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteUserFromDomains provides a mock function with given fields: ctx, id
func (_m *Service) DeleteUserFromDomains(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteUserFromDomains")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DisableDomain provides a mock function with given fields: ctx, sesssion, id
func (_m *Service) DisableDomain(ctx context.Context, sesssion authn.Session, id string) (domains.Domain, error) {
	ret := _m.Called(ctx, sesssion, id)

	if len(ret) == 0 {
		panic("no return value specified for DisableDomain")
	}

	var r0 domains.Domain
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, authn.Session, string) (domains.Domain, error)); ok {
		return rf(ctx, sesssion, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, authn.Session, string) domains.Domain); ok {
		r0 = rf(ctx, sesssion, id)
	} else {
		r0 = ret.Get(0).(domains.Domain)
	}

	if rf, ok := ret.Get(1).(func(context.Context, authn.Session, string) error); ok {
		r1 = rf(ctx, sesssion, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EnableDomain provides a mock function with given fields: ctx, sesssion, id
func (_m *Service) EnableDomain(ctx context.Context, sesssion authn.Session, id string) (domains.Domain, error) {
	ret := _m.Called(ctx, sesssion, id)

	if len(ret) == 0 {
		panic("no return value specified for EnableDomain")
	}

	var r0 domains.Domain
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, authn.Session, string) (domains.Domain, error)); ok {
		return rf(ctx, sesssion, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, authn.Session, string) domains.Domain); ok {
		r0 = rf(ctx, sesssion, id)
	} else {
		r0 = ret.Get(0).(domains.Domain)
	}

	if rf, ok := ret.Get(1).(func(context.Context, authn.Session, string) error); ok {
		r1 = rf(ctx, sesssion, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FreezeDomain provides a mock function with given fields: ctx, sesssion, id
func (_m *Service) FreezeDomain(ctx context.Context, sesssion authn.Session, id string) (domains.Domain, error) {
	ret := _m.Called(ctx, sesssion, id)

	if len(ret) == 0 {
		panic("no return value specified for FreezeDomain")
	}

	var r0 domains.Domain
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, authn.Session, string) (domains.Domain, error)); ok {
		return rf(ctx, sesssion, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, authn.Session, string) domains.Domain); ok {
		r0 = rf(ctx, sesssion, id)
	} else {
		r0 = ret.Get(0).(domains.Domain)
	}

	if rf, ok := ret.Get(1).(func(context.Context, authn.Session, string) error); ok {
		r1 = rf(ctx, sesssion, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAvailableActions provides a mock function with given fields: ctx, session
func (_m *Service) ListAvailableActions(ctx context.Context, session authn.Session) ([]string, error) {
	ret := _m.Called(ctx, session)

	if len(ret) == 0 {
		panic("no return value specified for ListAvailableActions")
	}

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, authn.Session) ([]string, error)); ok {
		return rf(ctx, session)
	}
	if rf, ok := ret.Get(0).(func(context.Context, authn.Session) []string); ok {
		r0 = rf(ctx, session)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, authn.Session) error); ok {
		r1 = rf(ctx, session)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListDomains provides a mock function with given fields: ctx, sesssion, page
func (_m *Service) ListDomains(ctx context.Context, sesssion authn.Session, page domains.Page) (domains.DomainsPage, error) {
	ret := _m.Called(ctx, sesssion, page)

	if len(ret) == 0 {
		panic("no return value specified for ListDomains")
	}

	var r0 domains.DomainsPage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, authn.Session, domains.Page) (domains.DomainsPage, error)); ok {
		return rf(ctx, sesssion, page)
	}
	if rf, ok := ret.Get(0).(func(context.Context, authn.Session, domains.Page) domains.DomainsPage); ok {
		r0 = rf(ctx, sesssion, page)
	} else {
		r0 = ret.Get(0).(domains.DomainsPage)
	}

	if rf, ok := ret.Get(1).(func(context.Context, authn.Session, domains.Page) error); ok {
		r1 = rf(ctx, sesssion, page)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveMemberFromAllRoles provides a mock function with given fields: ctx, session, memberID
func (_m *Service) RemoveMemberFromAllRoles(ctx context.Context, session authn.Session, memberID string) error {
	ret := _m.Called(ctx, session, memberID)

	if len(ret) == 0 {
		panic("no return value specified for RemoveMemberFromAllRoles")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, authn.Session, string) error); ok {
		r0 = rf(ctx, session, memberID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveRole provides a mock function with given fields: ctx, session, entityID, roleName
func (_m *Service) RemoveRole(ctx context.Context, session authn.Session, entityID string, roleName string) error {
	ret := _m.Called(ctx, session, entityID, roleName)

	if len(ret) == 0 {
		panic("no return value specified for RemoveRole")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, authn.Session, string, string) error); ok {
		r0 = rf(ctx, session, entityID, roleName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RetrieveAllRoles provides a mock function with given fields: ctx, session, entityID, limit, offset
func (_m *Service) RetrieveAllRoles(ctx context.Context, session authn.Session, entityID string, limit uint64, offset uint64) (roles.RolePage, error) {
	ret := _m.Called(ctx, session, entityID, limit, offset)

	if len(ret) == 0 {
		panic("no return value specified for RetrieveAllRoles")
	}

	var r0 roles.RolePage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, authn.Session, string, uint64, uint64) (roles.RolePage, error)); ok {
		return rf(ctx, session, entityID, limit, offset)
	}
	if rf, ok := ret.Get(0).(func(context.Context, authn.Session, string, uint64, uint64) roles.RolePage); ok {
		r0 = rf(ctx, session, entityID, limit, offset)
	} else {
		r0 = ret.Get(0).(roles.RolePage)
	}

	if rf, ok := ret.Get(1).(func(context.Context, authn.Session, string, uint64, uint64) error); ok {
		r1 = rf(ctx, session, entityID, limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RetrieveDomain provides a mock function with given fields: ctx, sesssion, id
func (_m *Service) RetrieveDomain(ctx context.Context, sesssion authn.Session, id string) (domains.Domain, error) {
	ret := _m.Called(ctx, sesssion, id)

	if len(ret) == 0 {
		panic("no return value specified for RetrieveDomain")
	}

	var r0 domains.Domain
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, authn.Session, string) (domains.Domain, error)); ok {
		return rf(ctx, sesssion, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, authn.Session, string) domains.Domain); ok {
		r0 = rf(ctx, sesssion, id)
	} else {
		r0 = ret.Get(0).(domains.Domain)
	}

	if rf, ok := ret.Get(1).(func(context.Context, authn.Session, string) error); ok {
		r1 = rf(ctx, sesssion, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RetrieveRole provides a mock function with given fields: ctx, session, entityID, roleName
func (_m *Service) RetrieveRole(ctx context.Context, session authn.Session, entityID string, roleName string) (roles.Role, error) {
	ret := _m.Called(ctx, session, entityID, roleName)

	if len(ret) == 0 {
		panic("no return value specified for RetrieveRole")
	}

	var r0 roles.Role
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, authn.Session, string, string) (roles.Role, error)); ok {
		return rf(ctx, session, entityID, roleName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, authn.Session, string, string) roles.Role); ok {
		r0 = rf(ctx, session, entityID, roleName)
	} else {
		r0 = ret.Get(0).(roles.Role)
	}

	if rf, ok := ret.Get(1).(func(context.Context, authn.Session, string, string) error); ok {
		r1 = rf(ctx, session, entityID, roleName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RoleAddActions provides a mock function with given fields: ctx, session, entityID, roleName, actions
func (_m *Service) RoleAddActions(ctx context.Context, session authn.Session, entityID string, roleName string, actions []string) ([]string, error) {
	ret := _m.Called(ctx, session, entityID, roleName, actions)

	if len(ret) == 0 {
		panic("no return value specified for RoleAddActions")
	}

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, authn.Session, string, string, []string) ([]string, error)); ok {
		return rf(ctx, session, entityID, roleName, actions)
	}
	if rf, ok := ret.Get(0).(func(context.Context, authn.Session, string, string, []string) []string); ok {
		r0 = rf(ctx, session, entityID, roleName, actions)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, authn.Session, string, string, []string) error); ok {
		r1 = rf(ctx, session, entityID, roleName, actions)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RoleAddMembers provides a mock function with given fields: ctx, session, entityID, roleName, members
func (_m *Service) RoleAddMembers(ctx context.Context, session authn.Session, entityID string, roleName string, members []string) ([]string, error) {
	ret := _m.Called(ctx, session, entityID, roleName, members)

	if len(ret) == 0 {
		panic("no return value specified for RoleAddMembers")
	}

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, authn.Session, string, string, []string) ([]string, error)); ok {
		return rf(ctx, session, entityID, roleName, members)
	}
	if rf, ok := ret.Get(0).(func(context.Context, authn.Session, string, string, []string) []string); ok {
		r0 = rf(ctx, session, entityID, roleName, members)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, authn.Session, string, string, []string) error); ok {
		r1 = rf(ctx, session, entityID, roleName, members)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RoleCheckActionsExists provides a mock function with given fields: ctx, session, entityID, roleName, actions
func (_m *Service) RoleCheckActionsExists(ctx context.Context, session authn.Session, entityID string, roleName string, actions []string) (bool, error) {
	ret := _m.Called(ctx, session, entityID, roleName, actions)

	if len(ret) == 0 {
		panic("no return value specified for RoleCheckActionsExists")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, authn.Session, string, string, []string) (bool, error)); ok {
		return rf(ctx, session, entityID, roleName, actions)
	}
	if rf, ok := ret.Get(0).(func(context.Context, authn.Session, string, string, []string) bool); ok {
		r0 = rf(ctx, session, entityID, roleName, actions)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, authn.Session, string, string, []string) error); ok {
		r1 = rf(ctx, session, entityID, roleName, actions)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RoleCheckMembersExists provides a mock function with given fields: ctx, session, entityID, roleName, members
func (_m *Service) RoleCheckMembersExists(ctx context.Context, session authn.Session, entityID string, roleName string, members []string) (bool, error) {
	ret := _m.Called(ctx, session, entityID, roleName, members)

	if len(ret) == 0 {
		panic("no return value specified for RoleCheckMembersExists")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, authn.Session, string, string, []string) (bool, error)); ok {
		return rf(ctx, session, entityID, roleName, members)
	}
	if rf, ok := ret.Get(0).(func(context.Context, authn.Session, string, string, []string) bool); ok {
		r0 = rf(ctx, session, entityID, roleName, members)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, authn.Session, string, string, []string) error); ok {
		r1 = rf(ctx, session, entityID, roleName, members)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RoleListActions provides a mock function with given fields: ctx, session, entityID, roleName
func (_m *Service) RoleListActions(ctx context.Context, session authn.Session, entityID string, roleName string) ([]string, error) {
	ret := _m.Called(ctx, session, entityID, roleName)

	if len(ret) == 0 {
		panic("no return value specified for RoleListActions")
	}

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, authn.Session, string, string) ([]string, error)); ok {
		return rf(ctx, session, entityID, roleName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, authn.Session, string, string) []string); ok {
		r0 = rf(ctx, session, entityID, roleName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, authn.Session, string, string) error); ok {
		r1 = rf(ctx, session, entityID, roleName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RoleListMembers provides a mock function with given fields: ctx, session, entityID, roleName, limit, offset
func (_m *Service) RoleListMembers(ctx context.Context, session authn.Session, entityID string, roleName string, limit uint64, offset uint64) (roles.MembersPage, error) {
	ret := _m.Called(ctx, session, entityID, roleName, limit, offset)

	if len(ret) == 0 {
		panic("no return value specified for RoleListMembers")
	}

	var r0 roles.MembersPage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, authn.Session, string, string, uint64, uint64) (roles.MembersPage, error)); ok {
		return rf(ctx, session, entityID, roleName, limit, offset)
	}
	if rf, ok := ret.Get(0).(func(context.Context, authn.Session, string, string, uint64, uint64) roles.MembersPage); ok {
		r0 = rf(ctx, session, entityID, roleName, limit, offset)
	} else {
		r0 = ret.Get(0).(roles.MembersPage)
	}

	if rf, ok := ret.Get(1).(func(context.Context, authn.Session, string, string, uint64, uint64) error); ok {
		r1 = rf(ctx, session, entityID, roleName, limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RoleRemoveActions provides a mock function with given fields: ctx, session, entityID, roleName, actions
func (_m *Service) RoleRemoveActions(ctx context.Context, session authn.Session, entityID string, roleName string, actions []string) error {
	ret := _m.Called(ctx, session, entityID, roleName, actions)

	if len(ret) == 0 {
		panic("no return value specified for RoleRemoveActions")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, authn.Session, string, string, []string) error); ok {
		r0 = rf(ctx, session, entityID, roleName, actions)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RoleRemoveAllActions provides a mock function with given fields: ctx, session, entityID, roleName
func (_m *Service) RoleRemoveAllActions(ctx context.Context, session authn.Session, entityID string, roleName string) error {
	ret := _m.Called(ctx, session, entityID, roleName)

	if len(ret) == 0 {
		panic("no return value specified for RoleRemoveAllActions")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, authn.Session, string, string) error); ok {
		r0 = rf(ctx, session, entityID, roleName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RoleRemoveAllMembers provides a mock function with given fields: ctx, session, entityID, roleName
func (_m *Service) RoleRemoveAllMembers(ctx context.Context, session authn.Session, entityID string, roleName string) error {
	ret := _m.Called(ctx, session, entityID, roleName)

	if len(ret) == 0 {
		panic("no return value specified for RoleRemoveAllMembers")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, authn.Session, string, string) error); ok {
		r0 = rf(ctx, session, entityID, roleName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RoleRemoveMembers provides a mock function with given fields: ctx, session, entityID, roleName, members
func (_m *Service) RoleRemoveMembers(ctx context.Context, session authn.Session, entityID string, roleName string, members []string) error {
	ret := _m.Called(ctx, session, entityID, roleName, members)

	if len(ret) == 0 {
		panic("no return value specified for RoleRemoveMembers")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, authn.Session, string, string, []string) error); ok {
		r0 = rf(ctx, session, entityID, roleName, members)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateDomain provides a mock function with given fields: ctx, sesssion, id, d
func (_m *Service) UpdateDomain(ctx context.Context, sesssion authn.Session, id string, d domains.DomainReq) (domains.Domain, error) {
	ret := _m.Called(ctx, sesssion, id, d)

	if len(ret) == 0 {
		panic("no return value specified for UpdateDomain")
	}

	var r0 domains.Domain
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, authn.Session, string, domains.DomainReq) (domains.Domain, error)); ok {
		return rf(ctx, sesssion, id, d)
	}
	if rf, ok := ret.Get(0).(func(context.Context, authn.Session, string, domains.DomainReq) domains.Domain); ok {
		r0 = rf(ctx, sesssion, id, d)
	} else {
		r0 = ret.Get(0).(domains.Domain)
	}

	if rf, ok := ret.Get(1).(func(context.Context, authn.Session, string, domains.DomainReq) error); ok {
		r1 = rf(ctx, sesssion, id, d)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateRoleName provides a mock function with given fields: ctx, session, entityID, oldRoleName, newRoleName
func (_m *Service) UpdateRoleName(ctx context.Context, session authn.Session, entityID string, oldRoleName string, newRoleName string) (roles.Role, error) {
	ret := _m.Called(ctx, session, entityID, oldRoleName, newRoleName)

	if len(ret) == 0 {
		panic("no return value specified for UpdateRoleName")
	}

	var r0 roles.Role
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, authn.Session, string, string, string) (roles.Role, error)); ok {
		return rf(ctx, session, entityID, oldRoleName, newRoleName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, authn.Session, string, string, string) roles.Role); ok {
		r0 = rf(ctx, session, entityID, oldRoleName, newRoleName)
	} else {
		r0 = ret.Get(0).(roles.Role)
	}

	if rf, ok := ret.Get(1).(func(context.Context, authn.Session, string, string, string) error); ok {
		r1 = rf(ctx, session, entityID, oldRoleName, newRoleName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewService creates a new instance of Service. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewService(t interface {
	mock.TestingT
	Cleanup(func())
}) *Service {
	mock := &Service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
