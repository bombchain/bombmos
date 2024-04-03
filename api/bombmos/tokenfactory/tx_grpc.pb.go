// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: bombmos/tokenfactory/tx.proto

package tokenfactory

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

const (
	Msg_UpdateParams_FullMethodName = "/bombmos.tokenfactory.Msg/UpdateParams"
	Msg_CreateDenom_FullMethodName  = "/bombmos.tokenfactory.Msg/CreateDenom"
	Msg_UpdateDenom_FullMethodName  = "/bombmos.tokenfactory.Msg/UpdateDenom"
	Msg_DeleteDenom_FullMethodName  = "/bombmos.tokenfactory.Msg/DeleteDenom"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
	CreateDenom(ctx context.Context, in *MsgCreateDenom, opts ...grpc.CallOption) (*MsgCreateDenomResponse, error)
	UpdateDenom(ctx context.Context, in *MsgUpdateDenom, opts ...grpc.CallOption) (*MsgUpdateDenomResponse, error)
	DeleteDenom(ctx context.Context, in *MsgDeleteDenom, opts ...grpc.CallOption) (*MsgDeleteDenomResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateParams_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CreateDenom(ctx context.Context, in *MsgCreateDenom, opts ...grpc.CallOption) (*MsgCreateDenomResponse, error) {
	out := new(MsgCreateDenomResponse)
	err := c.cc.Invoke(ctx, Msg_CreateDenom_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateDenom(ctx context.Context, in *MsgUpdateDenom, opts ...grpc.CallOption) (*MsgUpdateDenomResponse, error) {
	out := new(MsgUpdateDenomResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateDenom_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DeleteDenom(ctx context.Context, in *MsgDeleteDenom, opts ...grpc.CallOption) (*MsgDeleteDenomResponse, error) {
	out := new(MsgDeleteDenomResponse)
	err := c.cc.Invoke(ctx, Msg_DeleteDenom_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
	CreateDenom(context.Context, *MsgCreateDenom) (*MsgCreateDenomResponse, error)
	UpdateDenom(context.Context, *MsgUpdateDenom) (*MsgUpdateDenomResponse, error)
	DeleteDenom(context.Context, *MsgDeleteDenom) (*MsgDeleteDenomResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}
func (UnimplementedMsgServer) CreateDenom(context.Context, *MsgCreateDenom) (*MsgCreateDenomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDenom not implemented")
}
func (UnimplementedMsgServer) UpdateDenom(context.Context, *MsgUpdateDenom) (*MsgUpdateDenomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDenom not implemented")
}
func (UnimplementedMsgServer) DeleteDenom(context.Context, *MsgDeleteDenom) (*MsgDeleteDenomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDenom not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_UpdateParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateParams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, req.(*MsgUpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CreateDenom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateDenom)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateDenom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CreateDenom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateDenom(ctx, req.(*MsgCreateDenom))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateDenom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateDenom)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateDenom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateDenom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateDenom(ctx, req.(*MsgUpdateDenom))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DeleteDenom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeleteDenom)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DeleteDenom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_DeleteDenom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DeleteDenom(ctx, req.(*MsgDeleteDenom))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bombmos.tokenfactory.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
		{
			MethodName: "CreateDenom",
			Handler:    _Msg_CreateDenom_Handler,
		},
		{
			MethodName: "UpdateDenom",
			Handler:    _Msg_UpdateDenom_Handler,
		},
		{
			MethodName: "DeleteDenom",
			Handler:    _Msg_DeleteDenom_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bombmos/tokenfactory/tx.proto",
}
