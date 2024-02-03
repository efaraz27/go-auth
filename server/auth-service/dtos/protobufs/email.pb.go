// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: dtos/protobufs/email.proto

package protobufs

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

type SendVerificationEmailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	To    string `protobuf:"bytes,1,opt,name=To,proto3" json:"To,omitempty"`
	Token string `protobuf:"bytes,2,opt,name=Token,proto3" json:"Token,omitempty"`
}

func (x *SendVerificationEmailRequest) Reset() {
	*x = SendVerificationEmailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dtos_protobufs_email_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendVerificationEmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendVerificationEmailRequest) ProtoMessage() {}

func (x *SendVerificationEmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dtos_protobufs_email_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendVerificationEmailRequest.ProtoReflect.Descriptor instead.
func (*SendVerificationEmailRequest) Descriptor() ([]byte, []int) {
	return file_dtos_protobufs_email_proto_rawDescGZIP(), []int{0}
}

func (x *SendVerificationEmailRequest) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *SendVerificationEmailRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

var File_dtos_protobufs_email_proto protoreflect.FileDescriptor

var file_dtos_protobufs_email_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x64, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x73,
	0x2f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x73, 0x22, 0x44, 0x0a, 0x1c, 0x53, 0x65, 0x6e, 0x64, 0x56,
	0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x54, 0x6f, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x54, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x40, 0x5a,
	0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x66, 0x61, 0x72,
	0x61, 0x7a, 0x32, 0x37, 0x2f, 0x67, 0x6f, 0x2d, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x64, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dtos_protobufs_email_proto_rawDescOnce sync.Once
	file_dtos_protobufs_email_proto_rawDescData = file_dtos_protobufs_email_proto_rawDesc
)

func file_dtos_protobufs_email_proto_rawDescGZIP() []byte {
	file_dtos_protobufs_email_proto_rawDescOnce.Do(func() {
		file_dtos_protobufs_email_proto_rawDescData = protoimpl.X.CompressGZIP(file_dtos_protobufs_email_proto_rawDescData)
	})
	return file_dtos_protobufs_email_proto_rawDescData
}

var file_dtos_protobufs_email_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_dtos_protobufs_email_proto_goTypes = []interface{}{
	(*SendVerificationEmailRequest)(nil), // 0: protobufs.SendVerificationEmailRequest
}
var file_dtos_protobufs_email_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_dtos_protobufs_email_proto_init() }
func file_dtos_protobufs_email_proto_init() {
	if File_dtos_protobufs_email_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dtos_protobufs_email_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendVerificationEmailRequest); i {
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
			RawDescriptor: file_dtos_protobufs_email_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_dtos_protobufs_email_proto_goTypes,
		DependencyIndexes: file_dtos_protobufs_email_proto_depIdxs,
		MessageInfos:      file_dtos_protobufs_email_proto_msgTypes,
	}.Build()
	File_dtos_protobufs_email_proto = out.File
	file_dtos_protobufs_email_proto_rawDesc = nil
	file_dtos_protobufs_email_proto_goTypes = nil
	file_dtos_protobufs_email_proto_depIdxs = nil
}
