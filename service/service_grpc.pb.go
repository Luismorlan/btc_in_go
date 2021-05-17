// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package service

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

// FullNodeServiceClient is the client API for FullNodeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FullNodeServiceClient interface {
	// SetTransaction takes in a transaction posted by either self or client,
	// process and broadcast to other full node.
	SetTransaction(ctx context.Context, in *SetTransactionRequest, opts ...grpc.CallOption) (*SetTransactionResponse, error)
	// When a fullnode discovered a new block, it announce the block to other fullnode.
	SetBlock(ctx context.Context, in *SetBlockRequest, opts ...grpc.CallOption) (*SetBlockResponse, error)
	// Return balance for any client that want to query with their public key in the
	// form of UTXO to Output pairs.
	GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*GetBalanceResponse, error)
	// Add a new peer to this full node and create a bidirection connection.
	AddPeer(ctx context.Context, in *AddPeerRequest, opts ...grpc.CallOption) (*AddPeerResponse, error)
	// Return blocks in blockchain to help peers catching up with the system.
	Sync(ctx context.Context, in *SyncRequest, opts ...grpc.CallOption) (*SyncResponse, error)
}

type fullNodeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFullNodeServiceClient(cc grpc.ClientConnInterface) FullNodeServiceClient {
	return &fullNodeServiceClient{cc}
}

func (c *fullNodeServiceClient) SetTransaction(ctx context.Context, in *SetTransactionRequest, opts ...grpc.CallOption) (*SetTransactionResponse, error) {
	out := new(SetTransactionResponse)
	err := c.cc.Invoke(ctx, "/FullNodeService/SetTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fullNodeServiceClient) SetBlock(ctx context.Context, in *SetBlockRequest, opts ...grpc.CallOption) (*SetBlockResponse, error) {
	out := new(SetBlockResponse)
	err := c.cc.Invoke(ctx, "/FullNodeService/SetBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fullNodeServiceClient) GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*GetBalanceResponse, error) {
	out := new(GetBalanceResponse)
	err := c.cc.Invoke(ctx, "/FullNodeService/GetBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fullNodeServiceClient) AddPeer(ctx context.Context, in *AddPeerRequest, opts ...grpc.CallOption) (*AddPeerResponse, error) {
	out := new(AddPeerResponse)
	err := c.cc.Invoke(ctx, "/FullNodeService/AddPeer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fullNodeServiceClient) Sync(ctx context.Context, in *SyncRequest, opts ...grpc.CallOption) (*SyncResponse, error) {
	out := new(SyncResponse)
	err := c.cc.Invoke(ctx, "/FullNodeService/Sync", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FullNodeServiceServer is the server API for FullNodeService service.
// All implementations must embed UnimplementedFullNodeServiceServer
// for forward compatibility
type FullNodeServiceServer interface {
	// SetTransaction takes in a transaction posted by either self or client,
	// process and broadcast to other full node.
	SetTransaction(context.Context, *SetTransactionRequest) (*SetTransactionResponse, error)
	// When a fullnode discovered a new block, it announce the block to other fullnode.
	SetBlock(context.Context, *SetBlockRequest) (*SetBlockResponse, error)
	// Return balance for any client that want to query with their public key in the
	// form of UTXO to Output pairs.
	GetBalance(context.Context, *GetBalanceRequest) (*GetBalanceResponse, error)
	// Add a new peer to this full node and create a bidirection connection.
	AddPeer(context.Context, *AddPeerRequest) (*AddPeerResponse, error)
	// Return blocks in blockchain to help peers catching up with the system.
	Sync(context.Context, *SyncRequest) (*SyncResponse, error)
	mustEmbedUnimplementedFullNodeServiceServer()
}

// UnimplementedFullNodeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFullNodeServiceServer struct {
}

func (UnimplementedFullNodeServiceServer) SetTransaction(context.Context, *SetTransactionRequest) (*SetTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetTransaction not implemented")
}
func (UnimplementedFullNodeServiceServer) SetBlock(context.Context, *SetBlockRequest) (*SetBlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetBlock not implemented")
}
func (UnimplementedFullNodeServiceServer) GetBalance(context.Context, *GetBalanceRequest) (*GetBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBalance not implemented")
}
func (UnimplementedFullNodeServiceServer) AddPeer(context.Context, *AddPeerRequest) (*AddPeerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPeer not implemented")
}
func (UnimplementedFullNodeServiceServer) Sync(context.Context, *SyncRequest) (*SyncResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sync not implemented")
}
func (UnimplementedFullNodeServiceServer) mustEmbedUnimplementedFullNodeServiceServer() {}

// UnsafeFullNodeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FullNodeServiceServer will
// result in compilation errors.
type UnsafeFullNodeServiceServer interface {
	mustEmbedUnimplementedFullNodeServiceServer()
}

func RegisterFullNodeServiceServer(s grpc.ServiceRegistrar, srv FullNodeServiceServer) {
	s.RegisterService(&FullNodeService_ServiceDesc, srv)
}

func _FullNodeService_SetTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FullNodeServiceServer).SetTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FullNodeService/SetTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FullNodeServiceServer).SetTransaction(ctx, req.(*SetTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FullNodeService_SetBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FullNodeServiceServer).SetBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FullNodeService/SetBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FullNodeServiceServer).SetBlock(ctx, req.(*SetBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FullNodeService_GetBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FullNodeServiceServer).GetBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FullNodeService/GetBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FullNodeServiceServer).GetBalance(ctx, req.(*GetBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FullNodeService_AddPeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPeerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FullNodeServiceServer).AddPeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FullNodeService/AddPeer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FullNodeServiceServer).AddPeer(ctx, req.(*AddPeerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FullNodeService_Sync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FullNodeServiceServer).Sync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FullNodeService/Sync",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FullNodeServiceServer).Sync(ctx, req.(*SyncRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FullNodeService_ServiceDesc is the grpc.ServiceDesc for FullNodeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FullNodeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "FullNodeService",
	HandlerType: (*FullNodeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetTransaction",
			Handler:    _FullNodeService_SetTransaction_Handler,
		},
		{
			MethodName: "SetBlock",
			Handler:    _FullNodeService_SetBlock_Handler,
		},
		{
			MethodName: "GetBalance",
			Handler:    _FullNodeService_GetBalance_Handler,
		},
		{
			MethodName: "AddPeer",
			Handler:    _FullNodeService_AddPeer_Handler,
		},
		{
			MethodName: "Sync",
			Handler:    _FullNodeService_Sync_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service/service.proto",
}
