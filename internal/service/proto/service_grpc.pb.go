// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: service.proto

package __

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

// URLServiceClient is the client API for URLService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type URLServiceClient interface {
	Create(ctx context.Context, in *URL, opts ...grpc.CallOption) (*URL, error)
	Get(ctx context.Context, in *URL, opts ...grpc.CallOption) (*URL, error)
}

type uRLServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewURLServiceClient(cc grpc.ClientConnInterface) URLServiceClient {
	return &uRLServiceClient{cc}
}

func (c *uRLServiceClient) Create(ctx context.Context, in *URL, opts ...grpc.CallOption) (*URL, error) {
	out := new(URL)
	err := c.cc.Invoke(ctx, "/service.URLService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uRLServiceClient) Get(ctx context.Context, in *URL, opts ...grpc.CallOption) (*URL, error) {
	out := new(URL)
	err := c.cc.Invoke(ctx, "/service.URLService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// URLServiceServer is the server API for URLService service.
// All implementations must embed UnimplementedURLServiceServer
// for forward compatibility
type URLServiceServer interface {
	Create(context.Context, *URL) (*URL, error)
	Get(context.Context, *URL) (*URL, error)
	mustEmbedUnimplementedURLServiceServer()
}

// UnimplementedURLServiceServer must be embedded to have forward compatible implementations.
type UnimplementedURLServiceServer struct {
}

func (UnimplementedURLServiceServer) Create(context.Context, *URL) (*URL, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedURLServiceServer) Get(context.Context, *URL) (*URL, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedURLServiceServer) mustEmbedUnimplementedURLServiceServer() {}

// UnsafeURLServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to URLServiceServer will
// result in compilation errors.
type UnsafeURLServiceServer interface {
	mustEmbedUnimplementedURLServiceServer()
}

func RegisterURLServiceServer(s grpc.ServiceRegistrar, srv URLServiceServer) {
	s.RegisterService(&URLService_ServiceDesc, srv)
}

func _URLService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(URL)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(URLServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.URLService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(URLServiceServer).Create(ctx, req.(*URL))
	}
	return interceptor(ctx, in, info, handler)
}

func _URLService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(URL)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(URLServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.URLService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(URLServiceServer).Get(ctx, req.(*URL))
	}
	return interceptor(ctx, in, info, handler)
}

// URLService_ServiceDesc is the grpc.ServiceDesc for URLService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var URLService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.URLService",
	HandlerType: (*URLServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _URLService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _URLService_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
