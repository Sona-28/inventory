// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package netxd_project

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

// InventoryServiceClient is the client API for InventoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InventoryServiceClient interface {
	GetAllItems(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*AllInventoryItems, error)
	GetInventoryItemByItemName(ctx context.Context, in *ItemName, opts ...grpc.CallOption) (*InventoryItem, error)
}

type inventoryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInventoryServiceClient(cc grpc.ClientConnInterface) InventoryServiceClient {
	return &inventoryServiceClient{cc}
}

func (c *inventoryServiceClient) GetAllItems(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*AllInventoryItems, error) {
	out := new(AllInventoryItems)
	err := c.cc.Invoke(ctx, "/inventory.InventoryService/GetAllItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryServiceClient) GetInventoryItemByItemName(ctx context.Context, in *ItemName, opts ...grpc.CallOption) (*InventoryItem, error) {
	out := new(InventoryItem)
	err := c.cc.Invoke(ctx, "/inventory.InventoryService/GetInventoryItemByItemName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InventoryServiceServer is the server API for InventoryService service.
// All implementations must embed UnimplementedInventoryServiceServer
// for forward compatibility
type InventoryServiceServer interface {
	GetAllItems(context.Context, *Empty) (*AllInventoryItems, error)
	GetInventoryItemByItemName(context.Context, *ItemName) (*InventoryItem, error)
	mustEmbedUnimplementedInventoryServiceServer()
}

// UnimplementedInventoryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedInventoryServiceServer struct {
}

func (UnimplementedInventoryServiceServer) GetAllItems(context.Context, *Empty) (*AllInventoryItems, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllItems not implemented")
}
func (UnimplementedInventoryServiceServer) GetInventoryItemByItemName(context.Context, *ItemName) (*InventoryItem, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInventoryItemByItemName not implemented")
}
func (UnimplementedInventoryServiceServer) mustEmbedUnimplementedInventoryServiceServer() {}

// UnsafeInventoryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InventoryServiceServer will
// result in compilation errors.
type UnsafeInventoryServiceServer interface {
	mustEmbedUnimplementedInventoryServiceServer()
}

func RegisterInventoryServiceServer(s grpc.ServiceRegistrar, srv InventoryServiceServer) {
	s.RegisterService(&InventoryService_ServiceDesc, srv)
}

func _InventoryService_GetAllItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServiceServer).GetAllItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inventory.InventoryService/GetAllItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServiceServer).GetAllItems(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _InventoryService_GetInventoryItemByItemName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ItemName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServiceServer).GetInventoryItemByItemName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inventory.InventoryService/GetInventoryItemByItemName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServiceServer).GetInventoryItemByItemName(ctx, req.(*ItemName))
	}
	return interceptor(ctx, in, info, handler)
}

// InventoryService_ServiceDesc is the grpc.ServiceDesc for InventoryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InventoryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "inventory.InventoryService",
	HandlerType: (*InventoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllItems",
			Handler:    _InventoryService_GetAllItems_Handler,
		},
		{
			MethodName: "GetInventoryItemByItemName",
			Handler:    _InventoryService_GetInventoryItemByItemName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/inventory.proto",
}
