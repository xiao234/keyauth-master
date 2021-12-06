// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package namespace

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

// NamespaceServiceClient is the client API for NamespaceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NamespaceServiceClient interface {
	CreateNamespace(ctx context.Context, in *CreateNamespaceRequest, opts ...grpc.CallOption) (*Namespace, error)
	QueryNamespace(ctx context.Context, in *QueryNamespaceRequest, opts ...grpc.CallOption) (*Set, error)
	DescribeNamespace(ctx context.Context, in *DescriptNamespaceRequest, opts ...grpc.CallOption) (*Namespace, error)
	DeleteNamespace(ctx context.Context, in *DeleteNamespaceRequest, opts ...grpc.CallOption) (*Namespace, error)
}

type namespaceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNamespaceServiceClient(cc grpc.ClientConnInterface) NamespaceServiceClient {
	return &namespaceServiceClient{cc}
}

func (c *namespaceServiceClient) CreateNamespace(ctx context.Context, in *CreateNamespaceRequest, opts ...grpc.CallOption) (*Namespace, error) {
	out := new(Namespace)
	err := c.cc.Invoke(ctx, "/keyauth.namespace.NamespaceService/CreateNamespace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *namespaceServiceClient) QueryNamespace(ctx context.Context, in *QueryNamespaceRequest, opts ...grpc.CallOption) (*Set, error) {
	out := new(Set)
	err := c.cc.Invoke(ctx, "/keyauth.namespace.NamespaceService/QueryNamespace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *namespaceServiceClient) DescribeNamespace(ctx context.Context, in *DescriptNamespaceRequest, opts ...grpc.CallOption) (*Namespace, error) {
	out := new(Namespace)
	err := c.cc.Invoke(ctx, "/keyauth.namespace.NamespaceService/DescribeNamespace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *namespaceServiceClient) DeleteNamespace(ctx context.Context, in *DeleteNamespaceRequest, opts ...grpc.CallOption) (*Namespace, error) {
	out := new(Namespace)
	err := c.cc.Invoke(ctx, "/keyauth.namespace.NamespaceService/DeleteNamespace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NamespaceServiceServer is the server API for NamespaceService service.
// All implementations must embed UnimplementedNamespaceServiceServer
// for forward compatibility
type NamespaceServiceServer interface {
	CreateNamespace(context.Context, *CreateNamespaceRequest) (*Namespace, error)
	QueryNamespace(context.Context, *QueryNamespaceRequest) (*Set, error)
	DescribeNamespace(context.Context, *DescriptNamespaceRequest) (*Namespace, error)
	DeleteNamespace(context.Context, *DeleteNamespaceRequest) (*Namespace, error)
	mustEmbedUnimplementedNamespaceServiceServer()
}

// UnimplementedNamespaceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNamespaceServiceServer struct {
}

func (UnimplementedNamespaceServiceServer) CreateNamespace(context.Context, *CreateNamespaceRequest) (*Namespace, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNamespace not implemented")
}
func (UnimplementedNamespaceServiceServer) QueryNamespace(context.Context, *QueryNamespaceRequest) (*Set, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryNamespace not implemented")
}
func (UnimplementedNamespaceServiceServer) DescribeNamespace(context.Context, *DescriptNamespaceRequest) (*Namespace, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeNamespace not implemented")
}
func (UnimplementedNamespaceServiceServer) DeleteNamespace(context.Context, *DeleteNamespaceRequest) (*Namespace, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNamespace not implemented")
}
func (UnimplementedNamespaceServiceServer) mustEmbedUnimplementedNamespaceServiceServer() {}

// UnsafeNamespaceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NamespaceServiceServer will
// result in compilation errors.
type UnsafeNamespaceServiceServer interface {
	mustEmbedUnimplementedNamespaceServiceServer()
}

func RegisterNamespaceServiceServer(s grpc.ServiceRegistrar, srv NamespaceServiceServer) {
	s.RegisterService(&NamespaceService_ServiceDesc, srv)
}

func _NamespaceService_CreateNamespace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNamespaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NamespaceServiceServer).CreateNamespace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/keyauth.namespace.NamespaceService/CreateNamespace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NamespaceServiceServer).CreateNamespace(ctx, req.(*CreateNamespaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NamespaceService_QueryNamespace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryNamespaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NamespaceServiceServer).QueryNamespace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/keyauth.namespace.NamespaceService/QueryNamespace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NamespaceServiceServer).QueryNamespace(ctx, req.(*QueryNamespaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NamespaceService_DescribeNamespace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescriptNamespaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NamespaceServiceServer).DescribeNamespace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/keyauth.namespace.NamespaceService/DescribeNamespace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NamespaceServiceServer).DescribeNamespace(ctx, req.(*DescriptNamespaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NamespaceService_DeleteNamespace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNamespaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NamespaceServiceServer).DeleteNamespace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/keyauth.namespace.NamespaceService/DeleteNamespace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NamespaceServiceServer).DeleteNamespace(ctx, req.(*DeleteNamespaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NamespaceService_ServiceDesc is the grpc.ServiceDesc for NamespaceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NamespaceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "keyauth.namespace.NamespaceService",
	HandlerType: (*NamespaceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateNamespace",
			Handler:    _NamespaceService_CreateNamespace_Handler,
		},
		{
			MethodName: "QueryNamespace",
			Handler:    _NamespaceService_QueryNamespace_Handler,
		},
		{
			MethodName: "DescribeNamespace",
			Handler:    _NamespaceService_DescribeNamespace_Handler,
		},
		{
			MethodName: "DeleteNamespace",
			Handler:    _NamespaceService_DeleteNamespace_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/namespace/pb/service.proto",
}
