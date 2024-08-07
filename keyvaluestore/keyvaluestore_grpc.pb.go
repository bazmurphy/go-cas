// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.21.12
// source: keyvaluestore.proto

package keyvaluestore

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
	KeyValueStoreService_Get_FullMethodName = "/keyvaluestore.KeyValueStoreService/Get"
	KeyValueStoreService_Set_FullMethodName = "/keyvaluestore.KeyValueStoreService/Set"
	KeyValueStoreService_Cas_FullMethodName = "/keyvaluestore.KeyValueStoreService/Cas"
)

// KeyValueStoreServiceClient is the client API for KeyValueStoreService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KeyValueStoreServiceClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	Set(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*SetResponse, error)
	Cas(ctx context.Context, in *CasRequest, opts ...grpc.CallOption) (*CasResponse, error)
}

type keyValueStoreServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKeyValueStoreServiceClient(cc grpc.ClientConnInterface) KeyValueStoreServiceClient {
	return &keyValueStoreServiceClient{cc}
}

func (c *keyValueStoreServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, KeyValueStoreService_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyValueStoreServiceClient) Set(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*SetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetResponse)
	err := c.cc.Invoke(ctx, KeyValueStoreService_Set_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyValueStoreServiceClient) Cas(ctx context.Context, in *CasRequest, opts ...grpc.CallOption) (*CasResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CasResponse)
	err := c.cc.Invoke(ctx, KeyValueStoreService_Cas_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeyValueStoreServiceServer is the server API for KeyValueStoreService service.
// All implementations must embed UnimplementedKeyValueStoreServiceServer
// for forward compatibility
type KeyValueStoreServiceServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
	Set(context.Context, *SetRequest) (*SetResponse, error)
	Cas(context.Context, *CasRequest) (*CasResponse, error)
	mustEmbedUnimplementedKeyValueStoreServiceServer()
}

// UnimplementedKeyValueStoreServiceServer must be embedded to have forward compatible implementations.
type UnimplementedKeyValueStoreServiceServer struct {
}

func (UnimplementedKeyValueStoreServiceServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedKeyValueStoreServiceServer) Set(context.Context, *SetRequest) (*SetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (UnimplementedKeyValueStoreServiceServer) Cas(context.Context, *CasRequest) (*CasResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Cas not implemented")
}
func (UnimplementedKeyValueStoreServiceServer) mustEmbedUnimplementedKeyValueStoreServiceServer() {}

// UnsafeKeyValueStoreServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KeyValueStoreServiceServer will
// result in compilation errors.
type UnsafeKeyValueStoreServiceServer interface {
	mustEmbedUnimplementedKeyValueStoreServiceServer()
}

func RegisterKeyValueStoreServiceServer(s grpc.ServiceRegistrar, srv KeyValueStoreServiceServer) {
	s.RegisterService(&KeyValueStoreService_ServiceDesc, srv)
}

func _KeyValueStoreService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyValueStoreServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeyValueStoreService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyValueStoreServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyValueStoreService_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyValueStoreServiceServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeyValueStoreService_Set_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyValueStoreServiceServer).Set(ctx, req.(*SetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyValueStoreService_Cas_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CasRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyValueStoreServiceServer).Cas(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeyValueStoreService_Cas_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyValueStoreServiceServer).Cas(ctx, req.(*CasRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// KeyValueStoreService_ServiceDesc is the grpc.ServiceDesc for KeyValueStoreService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KeyValueStoreService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "keyvaluestore.KeyValueStoreService",
	HandlerType: (*KeyValueStoreServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _KeyValueStoreService_Get_Handler,
		},
		{
			MethodName: "Set",
			Handler:    _KeyValueStoreService_Set_Handler,
		},
		{
			MethodName: "Cas",
			Handler:    _KeyValueStoreService_Cas_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "keyvaluestore.proto",
}
