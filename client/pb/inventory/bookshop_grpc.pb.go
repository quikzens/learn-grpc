// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.8
// source: bookshop.proto

package inventory

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

// InventoryClient is the client API for Inventory service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InventoryClient interface {
	GetBookList(ctx context.Context, in *GetBookListRequest, opts ...grpc.CallOption) (*GetBookListResponse, error)
}

type inventoryClient struct {
	cc grpc.ClientConnInterface
}

func NewInventoryClient(cc grpc.ClientConnInterface) InventoryClient {
	return &inventoryClient{cc}
}

func (c *inventoryClient) GetBookList(ctx context.Context, in *GetBookListRequest, opts ...grpc.CallOption) (*GetBookListResponse, error) {
	out := new(GetBookListResponse)
	err := c.cc.Invoke(ctx, "/Inventory/GetBookList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InventoryServer is the server API for Inventory service.
// All implementations must embed UnimplementedInventoryServer
// for forward compatibility
type InventoryServer interface {
	GetBookList(context.Context, *GetBookListRequest) (*GetBookListResponse, error)
	mustEmbedUnimplementedInventoryServer()
}

// UnimplementedInventoryServer must be embedded to have forward compatible implementations.
type UnimplementedInventoryServer struct {
}

func (UnimplementedInventoryServer) GetBookList(context.Context, *GetBookListRequest) (*GetBookListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBookList not implemented")
}
func (UnimplementedInventoryServer) mustEmbedUnimplementedInventoryServer() {}

// UnsafeInventoryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InventoryServer will
// result in compilation errors.
type UnsafeInventoryServer interface {
	mustEmbedUnimplementedInventoryServer()
}

func RegisterInventoryServer(s grpc.ServiceRegistrar, srv InventoryServer) {
	s.RegisterService(&Inventory_ServiceDesc, srv)
}

func _Inventory_GetBookList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServer).GetBookList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Inventory/GetBookList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServer).GetBookList(ctx, req.(*GetBookListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Inventory_ServiceDesc is the grpc.ServiceDesc for Inventory service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Inventory_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Inventory",
	HandlerType: (*InventoryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBookList",
			Handler:    _Inventory_GetBookList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bookshop.proto",
}
