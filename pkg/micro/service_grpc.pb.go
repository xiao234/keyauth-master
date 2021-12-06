// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package micro

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MicroServiceClient is the client API for MicroService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MicroServiceClient interface {
	ValidateClientCredential(ctx context.Context, in *ValidateClientCredentialRequest, opts ...grpc.CallOption) (*Micro, error)
	CreateService(ctx context.Context, in *CreateMicroRequest, opts ...grpc.CallOption) (*Micro, error)
	QueryService(ctx context.Context, in *QueryMicroRequest, opts ...grpc.CallOption) (*Set, error)
	DescribeService(ctx context.Context, in *DescribeMicroRequest, opts ...grpc.CallOption) (*Micro, error)
	DeleteService(ctx context.Context, in *DeleteMicroRequest, opts ...grpc.CallOption) (*Micro, error)
	RefreshServiceClientSecret(ctx context.Context, in *DescribeMicroRequest, opts ...grpc.CallOption) (*Micro, error)
}

type microServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMicroServiceClient(cc grpc.ClientConnInterface) MicroServiceClient {
	return &microServiceClient{cc}
}

func (c *microServiceClient) ValidateClientCredential(ctx context.Context, in *ValidateClientCredentialRequest, opts ...grpc.CallOption) (*Micro, error) {
	out := new(Micro)
	err := c.cc.Invoke(ctx, "/keyauth.micro.MicroService/ValidateClientCredential", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *microServiceClient) CreateService(ctx context.Context, in *CreateMicroRequest, opts ...grpc.CallOption) (*Micro, error) {
	out := new(Micro)
	err := c.cc.Invoke(ctx, "/keyauth.micro.MicroService/CreateService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *microServiceClient) QueryService(ctx context.Context, in *QueryMicroRequest, opts ...grpc.CallOption) (*Set, error) {
	out := new(Set)
	err := c.cc.Invoke(ctx, "/keyauth.micro.MicroService/QueryService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *microServiceClient) DescribeService(ctx context.Context, in *DescribeMicroRequest, opts ...grpc.CallOption) (*Micro, error) {
	out := new(Micro)
	err := c.cc.Invoke(ctx, "/keyauth.micro.MicroService/DescribeService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *microServiceClient) DeleteService(ctx context.Context, in *DeleteMicroRequest, opts ...grpc.CallOption) (*Micro, error) {
	out := new(Micro)
	err := c.cc.Invoke(ctx, "/keyauth.micro.MicroService/DeleteService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *microServiceClient) RefreshServiceClientSecret(ctx context.Context, in *DescribeMicroRequest, opts ...grpc.CallOption) (*Micro, error) {
	out := new(Micro)
	err := c.cc.Invoke(ctx, "/keyauth.micro.MicroService/RefreshServiceClientSecret", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MicroServiceServer is the server API for MicroService service.
// All implementations must embed UnimplementedMicroServiceServer
// for forward compatibility
type MicroServiceServer interface {
	ValidateClientCredential(context.Context, *ValidateClientCredentialRequest) (*Micro, error)
	CreateService(context.Context, *CreateMicroRequest) (*Micro, error)
	QueryService(context.Context, *QueryMicroRequest) (*Set, error)
	DescribeService(context.Context, *DescribeMicroRequest) (*Micro, error)
	DeleteService(context.Context, *DeleteMicroRequest) (*Micro, error)
	RefreshServiceClientSecret(context.Context, *DescribeMicroRequest) (*Micro, error)
	mustEmbedUnimplementedMicroServiceServer()
}

// UnimplementedMicroServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMicroServiceServer struct {
}

func (UnimplementedMicroServiceServer) ValidateClientCredential(context.Context, *ValidateClientCredentialRequest) (*Micro, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateClientCredential not implemented")
}
func (UnimplementedMicroServiceServer) CreateService(context.Context, *CreateMicroRequest) (*Micro, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateService not implemented")
}
func (UnimplementedMicroServiceServer) QueryService(context.Context, *QueryMicroRequest) (*Set, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryService not implemented")
}
func (UnimplementedMicroServiceServer) DescribeService(context.Context, *DescribeMicroRequest) (*Micro, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeService not implemented")
}
func (UnimplementedMicroServiceServer) DeleteService(context.Context, *DeleteMicroRequest) (*Micro, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteService not implemented")
}
func (UnimplementedMicroServiceServer) RefreshServiceClientSecret(context.Context, *DescribeMicroRequest) (*Micro, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshServiceClientSecret not implemented")
}
func (UnimplementedMicroServiceServer) mustEmbedUnimplementedMicroServiceServer() {}

// UnsafeMicroServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MicroServiceServer will
// result in compilation errors.
type UnsafeMicroServiceServer interface {
	mustEmbedUnimplementedMicroServiceServer()
}

func RegisterMicroServiceServer(s grpc.ServiceRegistrar, srv MicroServiceServer) {
	s.RegisterService(&MicroService_ServiceDesc, srv)
}

func _MicroService_ValidateClientCredential_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateClientCredentialRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MicroServiceServer).ValidateClientCredential(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/keyauth.micro.MicroService/ValidateClientCredential",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MicroServiceServer).ValidateClientCredential(ctx, req.(*ValidateClientCredentialRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MicroService_CreateService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMicroRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MicroServiceServer).CreateService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/keyauth.micro.MicroService/CreateService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MicroServiceServer).CreateService(ctx, req.(*CreateMicroRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MicroService_QueryService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryMicroRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MicroServiceServer).QueryService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/keyauth.micro.MicroService/QueryService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MicroServiceServer).QueryService(ctx, req.(*QueryMicroRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MicroService_DescribeService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeMicroRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MicroServiceServer).DescribeService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/keyauth.micro.MicroService/DescribeService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MicroServiceServer).DescribeService(ctx, req.(*DescribeMicroRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MicroService_DeleteService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMicroRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MicroServiceServer).DeleteService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/keyauth.micro.MicroService/DeleteService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MicroServiceServer).DeleteService(ctx, req.(*DeleteMicroRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MicroService_RefreshServiceClientSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeMicroRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MicroServiceServer).RefreshServiceClientSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/keyauth.micro.MicroService/RefreshServiceClientSecret",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MicroServiceServer).RefreshServiceClientSecret(ctx, req.(*DescribeMicroRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MicroService_ServiceDesc is the grpc.ServiceDesc for MicroService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MicroService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "keyauth.micro.MicroService",
	HandlerType: (*MicroServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ValidateClientCredential",
			Handler:    _MicroService_ValidateClientCredential_Handler,
		},
		{
			MethodName: "CreateService",
			Handler:    _MicroService_CreateService_Handler,
		},
		{
			MethodName: "QueryService",
			Handler:    _MicroService_QueryService_Handler,
		},
		{
			MethodName: "DescribeService",
			Handler:    _MicroService_DescribeService_Handler,
		},
		{
			MethodName: "DeleteService",
			Handler:    _MicroService_DeleteService_Handler,
		},
		{
			MethodName: "RefreshServiceClientSecret",
			Handler:    _MicroService_RefreshServiceClientSecret_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/micro/pb/service.proto",
}
