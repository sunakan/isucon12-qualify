// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: isuxportal/services/admin/env_checks.proto

package admin

import (
	resources "github.com/isucon/isucon12-portal/proto.go/isuxportal/resources"
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

type ListEnvChecksQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeamId int64 `protobuf:"varint,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
}

func (x *ListEnvChecksQuery) Reset() {
	*x = ListEnvChecksQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_isuxportal_services_admin_env_checks_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEnvChecksQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEnvChecksQuery) ProtoMessage() {}

func (x *ListEnvChecksQuery) ProtoReflect() protoreflect.Message {
	mi := &file_isuxportal_services_admin_env_checks_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEnvChecksQuery.ProtoReflect.Descriptor instead.
func (*ListEnvChecksQuery) Descriptor() ([]byte, []int) {
	return file_isuxportal_services_admin_env_checks_proto_rawDescGZIP(), []int{0}
}

func (x *ListEnvChecksQuery) GetTeamId() int64 {
	if x != nil {
		return x.TeamId
	}
	return 0
}

type ListEnvChecksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EnvChecks []*resources.EnvCheck `protobuf:"bytes,1,rep,name=env_checks,json=envChecks,proto3" json:"env_checks,omitempty"`
}

func (x *ListEnvChecksResponse) Reset() {
	*x = ListEnvChecksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_isuxportal_services_admin_env_checks_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEnvChecksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEnvChecksResponse) ProtoMessage() {}

func (x *ListEnvChecksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_isuxportal_services_admin_env_checks_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEnvChecksResponse.ProtoReflect.Descriptor instead.
func (*ListEnvChecksResponse) Descriptor() ([]byte, []int) {
	return file_isuxportal_services_admin_env_checks_proto_rawDescGZIP(), []int{1}
}

func (x *ListEnvChecksResponse) GetEnvChecks() []*resources.EnvCheck {
	if x != nil {
		return x.EnvChecks
	}
	return nil
}

var File_isuxportal_services_admin_env_checks_proto protoreflect.FileDescriptor

var file_isuxportal_services_admin_env_checks_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x69, 0x73, 0x75, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x65, 0x6e, 0x76, 0x5f,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x69, 0x73,
	0x75, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x1a, 0x24, 0x69,
	0x73, 0x75, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2f, 0x65, 0x6e, 0x76, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x2d, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x76, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x65, 0x61,
	0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x74, 0x65, 0x61, 0x6d,
	0x49, 0x64, 0x22, 0x5c, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x76, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x0a, 0x65,
	0x6e, 0x76, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x24, 0x2e, 0x69, 0x73, 0x75, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x45, 0x6e, 0x76,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x09, 0x65, 0x6e, 0x76, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x73,
	0x42, 0x46, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69,
	0x73, 0x75, 0x63, 0x6f, 0x6e, 0x2f, 0x69, 0x73, 0x75, 0x63, 0x6f, 0x6e, 0x31, 0x32, 0x2d, 0x70,
	0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x6f, 0x2f, 0x69,
	0x73, 0x75, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_isuxportal_services_admin_env_checks_proto_rawDescOnce sync.Once
	file_isuxportal_services_admin_env_checks_proto_rawDescData = file_isuxportal_services_admin_env_checks_proto_rawDesc
)

func file_isuxportal_services_admin_env_checks_proto_rawDescGZIP() []byte {
	file_isuxportal_services_admin_env_checks_proto_rawDescOnce.Do(func() {
		file_isuxportal_services_admin_env_checks_proto_rawDescData = protoimpl.X.CompressGZIP(file_isuxportal_services_admin_env_checks_proto_rawDescData)
	})
	return file_isuxportal_services_admin_env_checks_proto_rawDescData
}

var file_isuxportal_services_admin_env_checks_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_isuxportal_services_admin_env_checks_proto_goTypes = []interface{}{
	(*ListEnvChecksQuery)(nil),    // 0: isuxportal.proto.services.admin.ListEnvChecksQuery
	(*ListEnvChecksResponse)(nil), // 1: isuxportal.proto.services.admin.ListEnvChecksResponse
	(*resources.EnvCheck)(nil),    // 2: isuxportal.proto.resources.EnvCheck
}
var file_isuxportal_services_admin_env_checks_proto_depIdxs = []int32{
	2, // 0: isuxportal.proto.services.admin.ListEnvChecksResponse.env_checks:type_name -> isuxportal.proto.resources.EnvCheck
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_isuxportal_services_admin_env_checks_proto_init() }
func file_isuxportal_services_admin_env_checks_proto_init() {
	if File_isuxportal_services_admin_env_checks_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_isuxportal_services_admin_env_checks_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEnvChecksQuery); i {
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
		file_isuxportal_services_admin_env_checks_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEnvChecksResponse); i {
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
			RawDescriptor: file_isuxportal_services_admin_env_checks_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_isuxportal_services_admin_env_checks_proto_goTypes,
		DependencyIndexes: file_isuxportal_services_admin_env_checks_proto_depIdxs,
		MessageInfos:      file_isuxportal_services_admin_env_checks_proto_msgTypes,
	}.Build()
	File_isuxportal_services_admin_env_checks_proto = out.File
	file_isuxportal_services_admin_env_checks_proto_rawDesc = nil
	file_isuxportal_services_admin_env_checks_proto_goTypes = nil
	file_isuxportal_services_admin_env_checks_proto_depIdxs = nil
}