// Copyright (c) Abstract Machines
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: auth.proto

package magistrala

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	AuthzService_Authorize_FullMethodName = "/magistrala.AuthzService/Authorize"
)

// AuthzServiceClient is the client API for AuthzService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// AuthzService is a service that provides authorization functionalities
// for magistrala services.
type AuthzServiceClient interface {
	// Authorize checks if the subject is authorized to perform
	// the action on the object.
	Authorize(ctx context.Context, in *AuthorizeReq, opts ...grpc.CallOption) (*AuthorizeRes, error)
}

type authzServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthzServiceClient(cc grpc.ClientConnInterface) AuthzServiceClient {
	return &authzServiceClient{cc}
}

func (c *authzServiceClient) Authorize(ctx context.Context, in *AuthorizeReq, opts ...grpc.CallOption) (*AuthorizeRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AuthorizeRes)
	err := c.cc.Invoke(ctx, AuthzService_Authorize_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthzServiceServer is the server API for AuthzService service.
// All implementations must embed UnimplementedAuthzServiceServer
// for forward compatibility
//
// AuthzService is a service that provides authorization functionalities
// for magistrala services.
type AuthzServiceServer interface {
	// Authorize checks if the subject is authorized to perform
	// the action on the object.
	Authorize(context.Context, *AuthorizeReq) (*AuthorizeRes, error)
	mustEmbedUnimplementedAuthzServiceServer()
}

// UnimplementedAuthzServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthzServiceServer struct {
}

func (UnimplementedAuthzServiceServer) Authorize(context.Context, *AuthorizeReq) (*AuthorizeRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authorize not implemented")
}
func (UnimplementedAuthzServiceServer) mustEmbedUnimplementedAuthzServiceServer() {}

// UnsafeAuthzServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthzServiceServer will
// result in compilation errors.
type UnsafeAuthzServiceServer interface {
	mustEmbedUnimplementedAuthzServiceServer()
}

func RegisterAuthzServiceServer(s grpc.ServiceRegistrar, srv AuthzServiceServer) {
	s.RegisterService(&AuthzService_ServiceDesc, srv)
}

func _AuthzService_Authorize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorizeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthzServiceServer).Authorize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthzService_Authorize_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthzServiceServer).Authorize(ctx, req.(*AuthorizeReq))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthzService_ServiceDesc is the grpc.ServiceDesc for AuthzService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthzService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "magistrala.AuthzService",
	HandlerType: (*AuthzServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Authorize",
			Handler:    _AuthzService_Authorize_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}

const (
	AuthnService_Issue_FullMethodName    = "/magistrala.AuthnService/Issue"
	AuthnService_Refresh_FullMethodName  = "/magistrala.AuthnService/Refresh"
	AuthnService_Identify_FullMethodName = "/magistrala.AuthnService/Identify"
)

// AuthnServiceClient is the client API for AuthnService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// AuthnService is a service that provides authentication functionalities
// for magistrala services.
type AuthnServiceClient interface {
	Issue(ctx context.Context, in *IssueReq, opts ...grpc.CallOption) (*Token, error)
	Refresh(ctx context.Context, in *RefreshReq, opts ...grpc.CallOption) (*Token, error)
	Identify(ctx context.Context, in *IdentityReq, opts ...grpc.CallOption) (*IdentityRes, error)
}

type authnServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthnServiceClient(cc grpc.ClientConnInterface) AuthnServiceClient {
	return &authnServiceClient{cc}
}

func (c *authnServiceClient) Issue(ctx context.Context, in *IssueReq, opts ...grpc.CallOption) (*Token, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Token)
	err := c.cc.Invoke(ctx, AuthnService_Issue_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authnServiceClient) Refresh(ctx context.Context, in *RefreshReq, opts ...grpc.CallOption) (*Token, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Token)
	err := c.cc.Invoke(ctx, AuthnService_Refresh_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authnServiceClient) Identify(ctx context.Context, in *IdentityReq, opts ...grpc.CallOption) (*IdentityRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(IdentityRes)
	err := c.cc.Invoke(ctx, AuthnService_Identify_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthnServiceServer is the server API for AuthnService service.
// All implementations must embed UnimplementedAuthnServiceServer
// for forward compatibility
//
// AuthnService is a service that provides authentication functionalities
// for magistrala services.
type AuthnServiceServer interface {
	Issue(context.Context, *IssueReq) (*Token, error)
	Refresh(context.Context, *RefreshReq) (*Token, error)
	Identify(context.Context, *IdentityReq) (*IdentityRes, error)
	mustEmbedUnimplementedAuthnServiceServer()
}

// UnimplementedAuthnServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthnServiceServer struct {
}

func (UnimplementedAuthnServiceServer) Issue(context.Context, *IssueReq) (*Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Issue not implemented")
}
func (UnimplementedAuthnServiceServer) Refresh(context.Context, *RefreshReq) (*Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Refresh not implemented")
}
func (UnimplementedAuthnServiceServer) Identify(context.Context, *IdentityReq) (*IdentityRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Identify not implemented")
}
func (UnimplementedAuthnServiceServer) mustEmbedUnimplementedAuthnServiceServer() {}

// UnsafeAuthnServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthnServiceServer will
// result in compilation errors.
type UnsafeAuthnServiceServer interface {
	mustEmbedUnimplementedAuthnServiceServer()
}

func RegisterAuthnServiceServer(s grpc.ServiceRegistrar, srv AuthnServiceServer) {
	s.RegisterService(&AuthnService_ServiceDesc, srv)
}

func _AuthnService_Issue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IssueReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthnServiceServer).Issue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthnService_Issue_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthnServiceServer).Issue(ctx, req.(*IssueReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthnService_Refresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthnServiceServer).Refresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthnService_Refresh_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthnServiceServer).Refresh(ctx, req.(*RefreshReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthnService_Identify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdentityReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthnServiceServer).Identify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthnService_Identify_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthnServiceServer).Identify(ctx, req.(*IdentityReq))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthnService_ServiceDesc is the grpc.ServiceDesc for AuthnService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthnService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "magistrala.AuthnService",
	HandlerType: (*AuthnServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Issue",
			Handler:    _AuthnService_Issue_Handler,
		},
		{
			MethodName: "Refresh",
			Handler:    _AuthnService_Refresh_Handler,
		},
		{
			MethodName: "Identify",
			Handler:    _AuthnService_Identify_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}

const (
	PolicyService_AddPolicy_FullMethodName            = "/magistrala.PolicyService/AddPolicy"
	PolicyService_AddPolicies_FullMethodName          = "/magistrala.PolicyService/AddPolicies"
	PolicyService_DeletePolicyFilter_FullMethodName   = "/magistrala.PolicyService/DeletePolicyFilter"
	PolicyService_DeletePolicies_FullMethodName       = "/magistrala.PolicyService/DeletePolicies"
	PolicyService_ListObjects_FullMethodName          = "/magistrala.PolicyService/ListObjects"
	PolicyService_ListAllObjects_FullMethodName       = "/magistrala.PolicyService/ListAllObjects"
	PolicyService_CountObjects_FullMethodName         = "/magistrala.PolicyService/CountObjects"
	PolicyService_ListSubjects_FullMethodName         = "/magistrala.PolicyService/ListSubjects"
	PolicyService_ListAllSubjects_FullMethodName      = "/magistrala.PolicyService/ListAllSubjects"
	PolicyService_CountSubjects_FullMethodName        = "/magistrala.PolicyService/CountSubjects"
	PolicyService_ListPermissions_FullMethodName      = "/magistrala.PolicyService/ListPermissions"
	PolicyService_DeleteEntityPolicies_FullMethodName = "/magistrala.PolicyService/DeleteEntityPolicies"
)

// PolicyServiceClient is the client API for PolicyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// PolicyService is a service that provides policy CRUD
// functionalities for magistrala services.
type PolicyServiceClient interface {
	AddPolicy(ctx context.Context, in *AddPolicyReq, opts ...grpc.CallOption) (*AddPolicyRes, error)
	AddPolicies(ctx context.Context, in *AddPoliciesReq, opts ...grpc.CallOption) (*AddPoliciesRes, error)
	DeletePolicyFilter(ctx context.Context, in *DeletePolicyFilterReq, opts ...grpc.CallOption) (*DeletePolicyRes, error)
	DeletePolicies(ctx context.Context, in *DeletePoliciesReq, opts ...grpc.CallOption) (*DeletePolicyRes, error)
	ListObjects(ctx context.Context, in *ListObjectsReq, opts ...grpc.CallOption) (*ListObjectsRes, error)
	ListAllObjects(ctx context.Context, in *ListObjectsReq, opts ...grpc.CallOption) (*ListObjectsRes, error)
	CountObjects(ctx context.Context, in *CountObjectsReq, opts ...grpc.CallOption) (*CountObjectsRes, error)
	ListSubjects(ctx context.Context, in *ListSubjectsReq, opts ...grpc.CallOption) (*ListSubjectsRes, error)
	ListAllSubjects(ctx context.Context, in *ListSubjectsReq, opts ...grpc.CallOption) (*ListSubjectsRes, error)
	CountSubjects(ctx context.Context, in *CountSubjectsReq, opts ...grpc.CallOption) (*CountSubjectsRes, error)
	ListPermissions(ctx context.Context, in *ListPermissionsReq, opts ...grpc.CallOption) (*ListPermissionsRes, error)
	DeleteEntityPolicies(ctx context.Context, in *DeleteEntityPoliciesReq, opts ...grpc.CallOption) (*DeletePolicyRes, error)
}

type policyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPolicyServiceClient(cc grpc.ClientConnInterface) PolicyServiceClient {
	return &policyServiceClient{cc}
}

func (c *policyServiceClient) AddPolicy(ctx context.Context, in *AddPolicyReq, opts ...grpc.CallOption) (*AddPolicyRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddPolicyRes)
	err := c.cc.Invoke(ctx, PolicyService_AddPolicy_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *policyServiceClient) AddPolicies(ctx context.Context, in *AddPoliciesReq, opts ...grpc.CallOption) (*AddPoliciesRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddPoliciesRes)
	err := c.cc.Invoke(ctx, PolicyService_AddPolicies_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *policyServiceClient) DeletePolicyFilter(ctx context.Context, in *DeletePolicyFilterReq, opts ...grpc.CallOption) (*DeletePolicyRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeletePolicyRes)
	err := c.cc.Invoke(ctx, PolicyService_DeletePolicyFilter_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *policyServiceClient) DeletePolicies(ctx context.Context, in *DeletePoliciesReq, opts ...grpc.CallOption) (*DeletePolicyRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeletePolicyRes)
	err := c.cc.Invoke(ctx, PolicyService_DeletePolicies_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *policyServiceClient) ListObjects(ctx context.Context, in *ListObjectsReq, opts ...grpc.CallOption) (*ListObjectsRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListObjectsRes)
	err := c.cc.Invoke(ctx, PolicyService_ListObjects_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *policyServiceClient) ListAllObjects(ctx context.Context, in *ListObjectsReq, opts ...grpc.CallOption) (*ListObjectsRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListObjectsRes)
	err := c.cc.Invoke(ctx, PolicyService_ListAllObjects_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *policyServiceClient) CountObjects(ctx context.Context, in *CountObjectsReq, opts ...grpc.CallOption) (*CountObjectsRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CountObjectsRes)
	err := c.cc.Invoke(ctx, PolicyService_CountObjects_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *policyServiceClient) ListSubjects(ctx context.Context, in *ListSubjectsReq, opts ...grpc.CallOption) (*ListSubjectsRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListSubjectsRes)
	err := c.cc.Invoke(ctx, PolicyService_ListSubjects_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *policyServiceClient) ListAllSubjects(ctx context.Context, in *ListSubjectsReq, opts ...grpc.CallOption) (*ListSubjectsRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListSubjectsRes)
	err := c.cc.Invoke(ctx, PolicyService_ListAllSubjects_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *policyServiceClient) CountSubjects(ctx context.Context, in *CountSubjectsReq, opts ...grpc.CallOption) (*CountSubjectsRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CountSubjectsRes)
	err := c.cc.Invoke(ctx, PolicyService_CountSubjects_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *policyServiceClient) ListPermissions(ctx context.Context, in *ListPermissionsReq, opts ...grpc.CallOption) (*ListPermissionsRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListPermissionsRes)
	err := c.cc.Invoke(ctx, PolicyService_ListPermissions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *policyServiceClient) DeleteEntityPolicies(ctx context.Context, in *DeleteEntityPoliciesReq, opts ...grpc.CallOption) (*DeletePolicyRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeletePolicyRes)
	err := c.cc.Invoke(ctx, PolicyService_DeleteEntityPolicies_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PolicyServiceServer is the server API for PolicyService service.
// All implementations must embed UnimplementedPolicyServiceServer
// for forward compatibility
//
// PolicyService is a service that provides policy CRUD
// functionalities for magistrala services.
type PolicyServiceServer interface {
	AddPolicy(context.Context, *AddPolicyReq) (*AddPolicyRes, error)
	AddPolicies(context.Context, *AddPoliciesReq) (*AddPoliciesRes, error)
	DeletePolicyFilter(context.Context, *DeletePolicyFilterReq) (*DeletePolicyRes, error)
	DeletePolicies(context.Context, *DeletePoliciesReq) (*DeletePolicyRes, error)
	ListObjects(context.Context, *ListObjectsReq) (*ListObjectsRes, error)
	ListAllObjects(context.Context, *ListObjectsReq) (*ListObjectsRes, error)
	CountObjects(context.Context, *CountObjectsReq) (*CountObjectsRes, error)
	ListSubjects(context.Context, *ListSubjectsReq) (*ListSubjectsRes, error)
	ListAllSubjects(context.Context, *ListSubjectsReq) (*ListSubjectsRes, error)
	CountSubjects(context.Context, *CountSubjectsReq) (*CountSubjectsRes, error)
	ListPermissions(context.Context, *ListPermissionsReq) (*ListPermissionsRes, error)
	DeleteEntityPolicies(context.Context, *DeleteEntityPoliciesReq) (*DeletePolicyRes, error)
	mustEmbedUnimplementedPolicyServiceServer()
}

// UnimplementedPolicyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPolicyServiceServer struct {
}

func (UnimplementedPolicyServiceServer) AddPolicy(context.Context, *AddPolicyReq) (*AddPolicyRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPolicy not implemented")
}
func (UnimplementedPolicyServiceServer) AddPolicies(context.Context, *AddPoliciesReq) (*AddPoliciesRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPolicies not implemented")
}
func (UnimplementedPolicyServiceServer) DeletePolicyFilter(context.Context, *DeletePolicyFilterReq) (*DeletePolicyRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePolicyFilter not implemented")
}
func (UnimplementedPolicyServiceServer) DeletePolicies(context.Context, *DeletePoliciesReq) (*DeletePolicyRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePolicies not implemented")
}
func (UnimplementedPolicyServiceServer) ListObjects(context.Context, *ListObjectsReq) (*ListObjectsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListObjects not implemented")
}
func (UnimplementedPolicyServiceServer) ListAllObjects(context.Context, *ListObjectsReq) (*ListObjectsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAllObjects not implemented")
}
func (UnimplementedPolicyServiceServer) CountObjects(context.Context, *CountObjectsReq) (*CountObjectsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountObjects not implemented")
}
func (UnimplementedPolicyServiceServer) ListSubjects(context.Context, *ListSubjectsReq) (*ListSubjectsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSubjects not implemented")
}
func (UnimplementedPolicyServiceServer) ListAllSubjects(context.Context, *ListSubjectsReq) (*ListSubjectsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAllSubjects not implemented")
}
func (UnimplementedPolicyServiceServer) CountSubjects(context.Context, *CountSubjectsReq) (*CountSubjectsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountSubjects not implemented")
}
func (UnimplementedPolicyServiceServer) ListPermissions(context.Context, *ListPermissionsReq) (*ListPermissionsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPermissions not implemented")
}
func (UnimplementedPolicyServiceServer) DeleteEntityPolicies(context.Context, *DeleteEntityPoliciesReq) (*DeletePolicyRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEntityPolicies not implemented")
}
func (UnimplementedPolicyServiceServer) mustEmbedUnimplementedPolicyServiceServer() {}

// UnsafePolicyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PolicyServiceServer will
// result in compilation errors.
type UnsafePolicyServiceServer interface {
	mustEmbedUnimplementedPolicyServiceServer()
}

func RegisterPolicyServiceServer(s grpc.ServiceRegistrar, srv PolicyServiceServer) {
	s.RegisterService(&PolicyService_ServiceDesc, srv)
}

func _PolicyService_AddPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPolicyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyServiceServer).AddPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PolicyService_AddPolicy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyServiceServer).AddPolicy(ctx, req.(*AddPolicyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolicyService_AddPolicies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPoliciesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyServiceServer).AddPolicies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PolicyService_AddPolicies_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyServiceServer).AddPolicies(ctx, req.(*AddPoliciesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolicyService_DeletePolicyFilter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePolicyFilterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyServiceServer).DeletePolicyFilter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PolicyService_DeletePolicyFilter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyServiceServer).DeletePolicyFilter(ctx, req.(*DeletePolicyFilterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolicyService_DeletePolicies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePoliciesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyServiceServer).DeletePolicies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PolicyService_DeletePolicies_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyServiceServer).DeletePolicies(ctx, req.(*DeletePoliciesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolicyService_ListObjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListObjectsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyServiceServer).ListObjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PolicyService_ListObjects_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyServiceServer).ListObjects(ctx, req.(*ListObjectsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolicyService_ListAllObjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListObjectsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyServiceServer).ListAllObjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PolicyService_ListAllObjects_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyServiceServer).ListAllObjects(ctx, req.(*ListObjectsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolicyService_CountObjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountObjectsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyServiceServer).CountObjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PolicyService_CountObjects_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyServiceServer).CountObjects(ctx, req.(*CountObjectsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolicyService_ListSubjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSubjectsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyServiceServer).ListSubjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PolicyService_ListSubjects_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyServiceServer).ListSubjects(ctx, req.(*ListSubjectsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolicyService_ListAllSubjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSubjectsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyServiceServer).ListAllSubjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PolicyService_ListAllSubjects_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyServiceServer).ListAllSubjects(ctx, req.(*ListSubjectsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolicyService_CountSubjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountSubjectsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyServiceServer).CountSubjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PolicyService_CountSubjects_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyServiceServer).CountSubjects(ctx, req.(*CountSubjectsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolicyService_ListPermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPermissionsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyServiceServer).ListPermissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PolicyService_ListPermissions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyServiceServer).ListPermissions(ctx, req.(*ListPermissionsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolicyService_DeleteEntityPolicies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEntityPoliciesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyServiceServer).DeleteEntityPolicies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PolicyService_DeleteEntityPolicies_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyServiceServer).DeleteEntityPolicies(ctx, req.(*DeleteEntityPoliciesReq))
	}
	return interceptor(ctx, in, info, handler)
}

// PolicyService_ServiceDesc is the grpc.ServiceDesc for PolicyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PolicyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "magistrala.PolicyService",
	HandlerType: (*PolicyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddPolicy",
			Handler:    _PolicyService_AddPolicy_Handler,
		},
		{
			MethodName: "AddPolicies",
			Handler:    _PolicyService_AddPolicies_Handler,
		},
		{
			MethodName: "DeletePolicyFilter",
			Handler:    _PolicyService_DeletePolicyFilter_Handler,
		},
		{
			MethodName: "DeletePolicies",
			Handler:    _PolicyService_DeletePolicies_Handler,
		},
		{
			MethodName: "ListObjects",
			Handler:    _PolicyService_ListObjects_Handler,
		},
		{
			MethodName: "ListAllObjects",
			Handler:    _PolicyService_ListAllObjects_Handler,
		},
		{
			MethodName: "CountObjects",
			Handler:    _PolicyService_CountObjects_Handler,
		},
		{
			MethodName: "ListSubjects",
			Handler:    _PolicyService_ListSubjects_Handler,
		},
		{
			MethodName: "ListAllSubjects",
			Handler:    _PolicyService_ListAllSubjects_Handler,
		},
		{
			MethodName: "CountSubjects",
			Handler:    _PolicyService_CountSubjects_Handler,
		},
		{
			MethodName: "ListPermissions",
			Handler:    _PolicyService_ListPermissions_Handler,
		},
		{
			MethodName: "DeleteEntityPolicies",
			Handler:    _PolicyService_DeleteEntityPolicies_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}
