//go:generate  mcube enum -m -p
//
// Code generated by protoc-gen-go-ext. DO NOT EDIT.
// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.1
// source: pkg/token/pb/service.proto

package token

import (
	_ "github.com/infraboard/mcube/pb/http"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_pkg_token_pb_service_proto protoreflect.FileDescriptor

var file_pkg_token_pb_service_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x6b, 0x67, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x70, 0x62, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6b, 0x65,
	0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x1a, 0x70, 0x6b, 0x67,
	0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x70, 0x6b, 0x67, 0x2f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x2f, 0x70, 0x62, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2f, 0x70,
	0x62, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x32, 0xc9, 0x07, 0x0a, 0x0c, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x68, 0x0a, 0x0a, 0x49, 0x73, 0x73, 0x75, 0x65, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x20, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x2e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x22, 0xa2, 0xa3, 0x8c, 0x4d, 0x1d,
	0x1a, 0x0e, 0x2f, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x32, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73,
	0x22, 0x04, 0x50, 0x4f, 0x53, 0x54, 0x2a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x6d, 0x0a,
	0x0d, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x23,
	0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x21, 0xa2, 0xa3, 0x8c, 0x4d, 0x1c,
	0x1a, 0x0e, 0x2f, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x32, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73,
	0x22, 0x03, 0x47, 0x45, 0x54, 0x2a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x79, 0x0a, 0x0d,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x23, 0x2e,
	0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x14, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x2d, 0xa2, 0xa3, 0x8c, 0x4d, 0x28, 0x1a,
	0x18, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x3a,
	0x69, 0x64, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x22, 0x03, 0x47, 0x45, 0x54, 0x2a, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x30, 0x01, 0x12, 0x78, 0x0a, 0x0b, 0x52, 0x65, 0x76, 0x6f, 0x6c,
	0x6b, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x21, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x52, 0x65, 0x76, 0x6f, 0x6c, 0x6b, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x6b, 0x65, 0x79, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22,
	0x30, 0xa2, 0xa3, 0x8c, 0x4d, 0x2b, 0x1a, 0x18, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x3a, 0x69, 0x64, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73,
	0x22, 0x06, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x2a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x30,
	0x01, 0x12, 0x73, 0x0a, 0x0a, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x20, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x14, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x2d, 0xa2, 0xa3, 0x8c, 0x4d, 0x28, 0x1a, 0x18,
	0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x3a, 0x69,
	0x64, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x22, 0x03, 0x50, 0x55, 0x54, 0x2a, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x30, 0x01, 0x12, 0x7d, 0x0a, 0x0f, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x25, 0x2e, 0x6b, 0x65, 0x79, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x14, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x2d, 0xa2, 0xa3, 0x8c, 0x4d, 0x28, 0x1a, 0x18, 0x2f,
	0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x3a, 0x69, 0x64,
	0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x22, 0x03, 0x50, 0x55, 0x54, 0x2a, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x30, 0x01, 0x12, 0x71, 0x0a, 0x0a, 0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x20, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x53, 0x65, 0x74, 0x22, 0x2d, 0xa2, 0xa3, 0x8c, 0x4d, 0x28,
	0x1a, 0x18, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x3a, 0x69, 0x64, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x22, 0x03, 0x47, 0x45, 0x54, 0x2a,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x30, 0x01, 0x12, 0x83, 0x01, 0x0a, 0x0b, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x21, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x6b, 0x65,
	0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x2d, 0xa2, 0xa3, 0x8c, 0x4d, 0x28, 0x1a, 0x18, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x3a, 0x69, 0x64, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73,
	0x22, 0x03, 0x47, 0x45, 0x54, 0x2a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x30, 0x01, 0x42, 0x29,
	0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var file_pkg_token_pb_service_proto_goTypes = []interface{}{
	(*IssueTokenRequest)(nil),      // 0: keyauth.token.IssueTokenRequest
	(*ValidateTokenRequest)(nil),   // 1: keyauth.token.ValidateTokenRequest
	(*DescribeTokenRequest)(nil),   // 2: keyauth.token.DescribeTokenRequest
	(*RevolkTokenRequest)(nil),     // 3: keyauth.token.RevolkTokenRequest
	(*BlockTokenRequest)(nil),      // 4: keyauth.token.BlockTokenRequest
	(*ChangeNamespaceRequest)(nil), // 5: keyauth.token.ChangeNamespaceRequest
	(*QueryTokenRequest)(nil),      // 6: keyauth.token.QueryTokenRequest
	(*DeleteTokenRequest)(nil),     // 7: keyauth.token.DeleteTokenRequest
	(*Token)(nil),                  // 8: keyauth.token.Token
	(*Set)(nil),                    // 9: keyauth.token.Set
	(*DeleteTokenResponse)(nil),    // 10: keyauth.token.DeleteTokenResponse
}
var file_pkg_token_pb_service_proto_depIdxs = []int32{
	0,  // 0: keyauth.token.TokenService.IssueToken:input_type -> keyauth.token.IssueTokenRequest
	1,  // 1: keyauth.token.TokenService.ValidateToken:input_type -> keyauth.token.ValidateTokenRequest
	2,  // 2: keyauth.token.TokenService.DescribeToken:input_type -> keyauth.token.DescribeTokenRequest
	3,  // 3: keyauth.token.TokenService.RevolkToken:input_type -> keyauth.token.RevolkTokenRequest
	4,  // 4: keyauth.token.TokenService.BlockToken:input_type -> keyauth.token.BlockTokenRequest
	5,  // 5: keyauth.token.TokenService.ChangeNamespace:input_type -> keyauth.token.ChangeNamespaceRequest
	6,  // 6: keyauth.token.TokenService.QueryToken:input_type -> keyauth.token.QueryTokenRequest
	7,  // 7: keyauth.token.TokenService.DeleteToken:input_type -> keyauth.token.DeleteTokenRequest
	8,  // 8: keyauth.token.TokenService.IssueToken:output_type -> keyauth.token.Token
	8,  // 9: keyauth.token.TokenService.ValidateToken:output_type -> keyauth.token.Token
	8,  // 10: keyauth.token.TokenService.DescribeToken:output_type -> keyauth.token.Token
	8,  // 11: keyauth.token.TokenService.RevolkToken:output_type -> keyauth.token.Token
	8,  // 12: keyauth.token.TokenService.BlockToken:output_type -> keyauth.token.Token
	8,  // 13: keyauth.token.TokenService.ChangeNamespace:output_type -> keyauth.token.Token
	9,  // 14: keyauth.token.TokenService.QueryToken:output_type -> keyauth.token.Set
	10, // 15: keyauth.token.TokenService.DeleteToken:output_type -> keyauth.token.DeleteTokenResponse
	8,  // [8:16] is the sub-list for method output_type
	0,  // [0:8] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_token_pb_service_proto_init() }
func file_pkg_token_pb_service_proto_init() {
	if File_pkg_token_pb_service_proto != nil {
		return
	}
	file_pkg_token_pb_request_proto_init()
	file_pkg_token_pb_token_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_token_pb_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_token_pb_service_proto_goTypes,
		DependencyIndexes: file_pkg_token_pb_service_proto_depIdxs,
	}.Build()
	File_pkg_token_pb_service_proto = out.File
	file_pkg_token_pb_service_proto_rawDesc = nil
	file_pkg_token_pb_service_proto_goTypes = nil
	file_pkg_token_pb_service_proto_depIdxs = nil
}