// Copyright (c) Abstract Machines
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: auth/v1/auth.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AuthNReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *AuthNReq) Reset() {
	*x = AuthNReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_v1_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthNReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthNReq) ProtoMessage() {}

func (x *AuthNReq) ProtoReflect() protoreflect.Message {
	mi := &file_auth_v1_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthNReq.ProtoReflect.Descriptor instead.
func (*AuthNReq) Descriptor() ([]byte, []int) {
	return file_auth_v1_auth_proto_rawDescGZIP(), []int{0}
}

func (x *AuthNReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type AuthNRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                             // IMPROVEMENT NOTE: change name from "id" to "subject" , sub in jwt = user id  + domain id //
	UserId   string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`       // user id
	DomainId string `protobuf:"bytes,3,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"` // domain id
}

func (x *AuthNRes) Reset() {
	*x = AuthNRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_v1_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthNRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthNRes) ProtoMessage() {}

func (x *AuthNRes) ProtoReflect() protoreflect.Message {
	mi := &file_auth_v1_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthNRes.ProtoReflect.Descriptor instead.
func (*AuthNRes) Descriptor() ([]byte, []int) {
	return file_auth_v1_auth_proto_rawDescGZIP(), []int{1}
}

func (x *AuthNRes) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AuthNRes) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AuthNRes) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

type AuthZReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Domain          string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`                                          // Domain
	SubjectType     string `protobuf:"bytes,2,opt,name=subject_type,json=subjectType,proto3" json:"subject_type,omitempty"`             // Client or User
	SubjectKind     string `protobuf:"bytes,3,opt,name=subject_kind,json=subjectKind,proto3" json:"subject_kind,omitempty"`             // ID or Token
	SubjectRelation string `protobuf:"bytes,4,opt,name=subject_relation,json=subjectRelation,proto3" json:"subject_relation,omitempty"` // Subject relation
	Subject         string `protobuf:"bytes,5,opt,name=subject,proto3" json:"subject,omitempty"`                                        // Subject value (id or token, depending on kind)
	Relation        string `protobuf:"bytes,6,opt,name=relation,proto3" json:"relation,omitempty"`                                      // Relation to filter
	Permission      string `protobuf:"bytes,7,opt,name=permission,proto3" json:"permission,omitempty"`                                  // Action
	Object          string `protobuf:"bytes,8,opt,name=object,proto3" json:"object,omitempty"`                                          // Object ID
	ObjectType      string `protobuf:"bytes,9,opt,name=object_type,json=objectType,proto3" json:"object_type,omitempty"`                // Client, User, Group
}

func (x *AuthZReq) Reset() {
	*x = AuthZReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_v1_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthZReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthZReq) ProtoMessage() {}

func (x *AuthZReq) ProtoReflect() protoreflect.Message {
	mi := &file_auth_v1_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthZReq.ProtoReflect.Descriptor instead.
func (*AuthZReq) Descriptor() ([]byte, []int) {
	return file_auth_v1_auth_proto_rawDescGZIP(), []int{2}
}

func (x *AuthZReq) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *AuthZReq) GetSubjectType() string {
	if x != nil {
		return x.SubjectType
	}
	return ""
}

func (x *AuthZReq) GetSubjectKind() string {
	if x != nil {
		return x.SubjectKind
	}
	return ""
}

func (x *AuthZReq) GetSubjectRelation() string {
	if x != nil {
		return x.SubjectRelation
	}
	return ""
}

func (x *AuthZReq) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *AuthZReq) GetRelation() string {
	if x != nil {
		return x.Relation
	}
	return ""
}

func (x *AuthZReq) GetPermission() string {
	if x != nil {
		return x.Permission
	}
	return ""
}

func (x *AuthZReq) GetObject() string {
	if x != nil {
		return x.Object
	}
	return ""
}

func (x *AuthZReq) GetObjectType() string {
	if x != nil {
		return x.ObjectType
	}
	return ""
}

type AuthZRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Authorized bool   `protobuf:"varint,1,opt,name=authorized,proto3" json:"authorized,omitempty"`
	Id         string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AuthZRes) Reset() {
	*x = AuthZRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_v1_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthZRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthZRes) ProtoMessage() {}

func (x *AuthZRes) ProtoReflect() protoreflect.Message {
	mi := &file_auth_v1_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthZRes.ProtoReflect.Descriptor instead.
func (*AuthZRes) Descriptor() ([]byte, []int) {
	return file_auth_v1_auth_proto_rawDescGZIP(), []int{3}
}

func (x *AuthZRes) GetAuthorized() bool {
	if x != nil {
		return x.Authorized
	}
	return false
}

func (x *AuthZRes) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_auth_v1_auth_proto protoreflect.FileDescriptor

var file_auth_v1_auth_proto_rawDesc = []byte{
	0x0a, 0x12, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x22, 0x20, 0x0a,
	0x08, 0x41, 0x75, 0x74, 0x68, 0x4e, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22,
	0x50, 0x0a, 0x08, 0x41, 0x75, 0x74, 0x68, 0x4e, 0x52, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49,
	0x64, 0x22, 0xa2, 0x02, 0x0a, 0x08, 0x41, 0x75, 0x74, 0x68, 0x5a, 0x52, 0x65, 0x71, 0x12, 0x16,
	0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x75,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x75, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x29, 0x0a, 0x10,
	0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a,
	0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a,
	0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x0a, 0x08, 0x41, 0x75, 0x74, 0x68, 0x5a, 0x52,
	0x65, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a,
	0x65, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x32, 0x7a, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x33, 0x0a, 0x09, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x12, 0x11,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x5a, 0x52, 0x65,
	0x71, 0x1a, 0x11, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68,
	0x5a, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x0c, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e,
	0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x75, 0x74, 0x68, 0x4e, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x4e, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42, 0x35,
	0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x62, 0x73,
	0x6d, 0x61, 0x63, 0x68, 0x2f, 0x6d, 0x61, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x61, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x61, 0x75,
	0x74, 0x68, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_auth_v1_auth_proto_rawDescOnce sync.Once
	file_auth_v1_auth_proto_rawDescData = file_auth_v1_auth_proto_rawDesc
)

func file_auth_v1_auth_proto_rawDescGZIP() []byte {
	file_auth_v1_auth_proto_rawDescOnce.Do(func() {
		file_auth_v1_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_v1_auth_proto_rawDescData)
	})
	return file_auth_v1_auth_proto_rawDescData
}

var file_auth_v1_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_auth_v1_auth_proto_goTypes = []any{
	(*AuthNReq)(nil), // 0: auth.v1.AuthNReq
	(*AuthNRes)(nil), // 1: auth.v1.AuthNRes
	(*AuthZReq)(nil), // 2: auth.v1.AuthZReq
	(*AuthZRes)(nil), // 3: auth.v1.AuthZRes
}
var file_auth_v1_auth_proto_depIdxs = []int32{
	2, // 0: auth.v1.AuthService.Authorize:input_type -> auth.v1.AuthZReq
	0, // 1: auth.v1.AuthService.Authenticate:input_type -> auth.v1.AuthNReq
	3, // 2: auth.v1.AuthService.Authorize:output_type -> auth.v1.AuthZRes
	1, // 3: auth.v1.AuthService.Authenticate:output_type -> auth.v1.AuthNRes
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_auth_v1_auth_proto_init() }
func file_auth_v1_auth_proto_init() {
	if File_auth_v1_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_auth_v1_auth_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*AuthNReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_auth_v1_auth_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*AuthNRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_auth_v1_auth_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*AuthZReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_auth_v1_auth_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*AuthZRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_auth_v1_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auth_v1_auth_proto_goTypes,
		DependencyIndexes: file_auth_v1_auth_proto_depIdxs,
		MessageInfos:      file_auth_v1_auth_proto_msgTypes,
	}.Build()
	File_auth_v1_auth_proto = out.File
	file_auth_v1_auth_proto_rawDesc = nil
	file_auth_v1_auth_proto_goTypes = nil
	file_auth_v1_auth_proto_depIdxs = nil
}
