//go:generate  mcube enum -m -p
//
// Code generated by protoc-gen-go-ext. DO NOT EDIT.
// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.1
// source: pkg/verifycode/pb/request.proto

package verifycode

import (
	_ "github.com/infraboard/mcube/cmd/protoc-gen-go-ext/extension/tag"
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

// IssueCodeRequest 验证码申请请求
type IssueCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IssueType    IssueType `protobuf:"varint,1,opt,name=issue_type,json=issueType,proto3,enum=keyauth.verifycode.IssueType" json:"issue_type"`
	Username     string    `protobuf:"bytes,2,opt,name=username,proto3" json:"username" validate:"required"`
	Password     string    `protobuf:"bytes,3,opt,name=password,proto3" json:"password" validate:"required"`
	ClientId     string    `protobuf:"bytes,4,opt,name=client_id,json=clientId,proto3" json:"client_id" validate:"required"`
	ClientSecret string    `protobuf:"bytes,5,opt,name=client_secret,json=clientSecret,proto3" json:"client_secret" validate:"required"`
}

func (x *IssueCodeRequest) Reset() {
	*x = IssueCodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_verifycode_pb_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IssueCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IssueCodeRequest) ProtoMessage() {}

func (x *IssueCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_verifycode_pb_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IssueCodeRequest.ProtoReflect.Descriptor instead.
func (*IssueCodeRequest) Descriptor() ([]byte, []int) {
	return file_pkg_verifycode_pb_request_proto_rawDescGZIP(), []int{0}
}

func (x *IssueCodeRequest) GetIssueType() IssueType {
	if x != nil {
		return x.IssueType
	}
	return IssueType_PASS
}

func (x *IssueCodeRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *IssueCodeRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *IssueCodeRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *IssueCodeRequest) GetClientSecret() string {
	if x != nil {
		return x.ClientSecret
	}
	return ""
}

// IssueCodeResponse todo
type IssueCodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message"`
}

func (x *IssueCodeResponse) Reset() {
	*x = IssueCodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_verifycode_pb_request_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IssueCodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IssueCodeResponse) ProtoMessage() {}

func (x *IssueCodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_verifycode_pb_request_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IssueCodeResponse.ProtoReflect.Descriptor instead.
func (*IssueCodeResponse) Descriptor() ([]byte, []int) {
	return file_pkg_verifycode_pb_request_proto_rawDescGZIP(), []int{1}
}

func (x *IssueCodeResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// CheckCodeRequest 验证码校验请求
type CheckCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username" validate:"required"`
	Number   string `protobuf:"bytes,2,opt,name=number,proto3" json:"number" validate:"required"`
}

func (x *CheckCodeRequest) Reset() {
	*x = CheckCodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_verifycode_pb_request_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckCodeRequest) ProtoMessage() {}

func (x *CheckCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_verifycode_pb_request_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckCodeRequest.ProtoReflect.Descriptor instead.
func (*CheckCodeRequest) Descriptor() ([]byte, []int) {
	return file_pkg_verifycode_pb_request_proto_rawDescGZIP(), []int{2}
}

func (x *CheckCodeRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *CheckCodeRequest) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

// CheckCodeResponse todo
type CheckCodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message"`
}

func (x *CheckCodeResponse) Reset() {
	*x = CheckCodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_verifycode_pb_request_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckCodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckCodeResponse) ProtoMessage() {}

func (x *CheckCodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_verifycode_pb_request_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckCodeResponse.ProtoReflect.Descriptor instead.
func (*CheckCodeResponse) Descriptor() ([]byte, []int) {
	return file_pkg_verifycode_pb_request_proto_rawDescGZIP(), []int{3}
}

func (x *CheckCodeResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_pkg_verifycode_pb_request_proto protoreflect.FileDescriptor

var file_pkg_verifycode_pb_request_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x6b, 0x67, 0x2f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x63, 0x6f, 0x64, 0x65,
	0x2f, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x12, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x63, 0x6f, 0x64, 0x65, 0x1a, 0x49, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x63, 0x75,
	0x62, 0x65, 0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65,
	0x6e, 0x2d, 0x67, 0x6f, 0x2d, 0x65, 0x78, 0x74, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x2f, 0x74, 0x61, 0x67, 0x2f, 0x74, 0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x70, 0x6b, 0x67, 0x2f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x63, 0x6f, 0x64, 0x65,
	0x2f, 0x70, 0x62, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x95,
	0x03, 0x0a, 0x10, 0x49, 0x73, 0x73, 0x75, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x55, 0x0a, 0x0a, 0x69, 0x73, 0x73, 0x75, 0x65, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x49, 0x73, 0x73,
	0x75, 0x65, 0x54, 0x79, 0x70, 0x65, 0x42, 0x17, 0xc2, 0xde, 0x1f, 0x13, 0x0a, 0x11, 0x6a, 0x73,
	0x6f, 0x6e, 0x3a, 0x22, 0x69, 0x73, 0x73, 0x75, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0x52,
	0x09, 0x69, 0x73, 0x73, 0x75, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x45, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x29, 0xc2, 0xde,
	0x1f, 0x25, 0x0a, 0x23, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x20, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x3a, 0x22, 0x72, 0x65,
	0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x45, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x29, 0xc2, 0xde, 0x1f, 0x25, 0x0a, 0x23, 0x6a, 0x73, 0x6f, 0x6e, 0x3a,
	0x22, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x20, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x3a, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x47, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2a, 0xc2, 0xde, 0x1f,
	0x26, 0x0a, 0x24, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x22, 0x20, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x3a, 0x22, 0x72, 0x65,
	0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x53, 0x0a, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2e, 0xc2, 0xde, 0x1f, 0x2a, 0x0a, 0x28,
	0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x22, 0x20, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x3a, 0x22, 0x72,
	0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x52, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x22, 0x43, 0x0a, 0x11, 0x49, 0x73, 0x73, 0x75, 0x65, 0x43,
	0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x14, 0xc2, 0xde,
	0x1f, 0x10, 0x0a, 0x0e, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x9a, 0x01, 0x0a, 0x10,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x45, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x29, 0xc2, 0xde, 0x1f, 0x25, 0x0a, 0x23, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x20, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x3a, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3f, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x27, 0xc2, 0xde, 0x1f, 0x23, 0x0a, 0x21, 0x6a,
	0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x20, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x3a, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22,
	0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x43, 0x0a, 0x11, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x14,
	0xc2, 0xde, 0x1f, 0x10, 0x0a, 0x0e, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x2e, 0x5a,
	0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x63, 0x6f, 0x64, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_verifycode_pb_request_proto_rawDescOnce sync.Once
	file_pkg_verifycode_pb_request_proto_rawDescData = file_pkg_verifycode_pb_request_proto_rawDesc
)

func file_pkg_verifycode_pb_request_proto_rawDescGZIP() []byte {
	file_pkg_verifycode_pb_request_proto_rawDescOnce.Do(func() {
		file_pkg_verifycode_pb_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_verifycode_pb_request_proto_rawDescData)
	})
	return file_pkg_verifycode_pb_request_proto_rawDescData
}

var file_pkg_verifycode_pb_request_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pkg_verifycode_pb_request_proto_goTypes = []interface{}{
	(*IssueCodeRequest)(nil),  // 0: keyauth.verifycode.IssueCodeRequest
	(*IssueCodeResponse)(nil), // 1: keyauth.verifycode.IssueCodeResponse
	(*CheckCodeRequest)(nil),  // 2: keyauth.verifycode.CheckCodeRequest
	(*CheckCodeResponse)(nil), // 3: keyauth.verifycode.CheckCodeResponse
	(IssueType)(0),            // 4: keyauth.verifycode.IssueType
}
var file_pkg_verifycode_pb_request_proto_depIdxs = []int32{
	4, // 0: keyauth.verifycode.IssueCodeRequest.issue_type:type_name -> keyauth.verifycode.IssueType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_verifycode_pb_request_proto_init() }
func file_pkg_verifycode_pb_request_proto_init() {
	if File_pkg_verifycode_pb_request_proto != nil {
		return
	}
	file_pkg_verifycode_pb_enum_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pkg_verifycode_pb_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IssueCodeRequest); i {
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
		file_pkg_verifycode_pb_request_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IssueCodeResponse); i {
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
		file_pkg_verifycode_pb_request_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckCodeRequest); i {
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
		file_pkg_verifycode_pb_request_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckCodeResponse); i {
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
			RawDescriptor: file_pkg_verifycode_pb_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_verifycode_pb_request_proto_goTypes,
		DependencyIndexes: file_pkg_verifycode_pb_request_proto_depIdxs,
		MessageInfos:      file_pkg_verifycode_pb_request_proto_msgTypes,
	}.Build()
	File_pkg_verifycode_pb_request_proto = out.File
	file_pkg_verifycode_pb_request_proto_rawDesc = nil
	file_pkg_verifycode_pb_request_proto_goTypes = nil
	file_pkg_verifycode_pb_request_proto_depIdxs = nil
}
