// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: bpfd.proto

package gobpfd

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

// LoaderClient is the client API for Loader service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LoaderClient interface {
	Load(ctx context.Context, in *LoadRequest, opts ...grpc.CallOption) (*LoadResponse, error)
	Unload(ctx context.Context, in *UnloadRequest, opts ...grpc.CallOption) (*UnloadResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	GetMap(ctx context.Context, in *GetMapRequest, opts ...grpc.CallOption) (*GetMapResponse, error)
}

type loaderClient struct {
	cc grpc.ClientConnInterface
}

func NewLoaderClient(cc grpc.ClientConnInterface) LoaderClient {
	return &loaderClient{cc}
}

func (c *loaderClient) Load(ctx context.Context, in *LoadRequest, opts ...grpc.CallOption) (*LoadResponse, error) {
	out := new(LoadResponse)
	err := c.cc.Invoke(ctx, "/bpfd.Loader/Load", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loaderClient) Unload(ctx context.Context, in *UnloadRequest, opts ...grpc.CallOption) (*UnloadResponse, error) {
	out := new(UnloadResponse)
	err := c.cc.Invoke(ctx, "/bpfd.Loader/Unload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loaderClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/bpfd.Loader/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loaderClient) GetMap(ctx context.Context, in *GetMapRequest, opts ...grpc.CallOption) (*GetMapResponse, error) {
	out := new(GetMapResponse)
	err := c.cc.Invoke(ctx, "/bpfd.Loader/GetMap", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoaderServer is the server API for Loader service.
// All implementations must embed UnimplementedLoaderServer
// for forward compatibility
type LoaderServer interface {
	Load(context.Context, *LoadRequest) (*LoadResponse, error)
	Unload(context.Context, *UnloadRequest) (*UnloadResponse, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
	GetMap(context.Context, *GetMapRequest) (*GetMapResponse, error)
	mustEmbedUnimplementedLoaderServer()
}

// UnimplementedLoaderServer must be embedded to have forward compatible implementations.
type UnimplementedLoaderServer struct {
}

func (UnimplementedLoaderServer) Load(context.Context, *LoadRequest) (*LoadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Load not implemented")
}
func (UnimplementedLoaderServer) Unload(context.Context, *UnloadRequest) (*UnloadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unload not implemented")
}
func (UnimplementedLoaderServer) List(context.Context, *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedLoaderServer) GetMap(context.Context, *GetMapRequest) (*GetMapResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMap not implemented")
}
func (UnimplementedLoaderServer) mustEmbedUnimplementedLoaderServer() {}

// UnsafeLoaderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LoaderServer will
// result in compilation errors.
type UnsafeLoaderServer interface {
	mustEmbedUnimplementedLoaderServer()
}

func RegisterLoaderServer(s grpc.ServiceRegistrar, srv LoaderServer) {
	s.RegisterService(&Loader_ServiceDesc, srv)
}

func _Loader_Load_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoaderServer).Load(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bpfd.Loader/Load",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoaderServer).Load(ctx, req.(*LoadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Loader_Unload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnloadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoaderServer).Unload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bpfd.Loader/Unload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoaderServer).Unload(ctx, req.(*UnloadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Loader_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoaderServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bpfd.Loader/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoaderServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Loader_GetMap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMapRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoaderServer).GetMap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bpfd.Loader/GetMap",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoaderServer).GetMap(ctx, req.(*GetMapRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Loader_ServiceDesc is the grpc.ServiceDesc for Loader service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Loader_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bpfd.Loader",
	HandlerType: (*LoaderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Load",
			Handler:    _Loader_Load_Handler,
		},
		{
			MethodName: "Unload",
			Handler:    _Loader_Unload_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Loader_List_Handler,
		},
		{
			MethodName: "GetMap",
			Handler:    _Loader_GetMap_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bpfd.proto",
}
