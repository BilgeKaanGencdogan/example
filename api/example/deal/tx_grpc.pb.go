// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: example/deal/tx.proto

package deal

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
	Msg_UpdateParams_FullMethodName    = "/example.deal.Msg/UpdateParams"
	Msg_CreateDeal_FullMethodName      = "/example.deal.Msg/CreateDeal"
	Msg_CreateContract_FullMethodName  = "/example.deal.Msg/CreateContract"
	Msg_CommitContract_FullMethodName  = "/example.deal.Msg/CommitContract"
	Msg_ApproveContract_FullMethodName = "/example.deal.Msg/ApproveContract"
	Msg_ShipOrder_FullMethodName       = "/example.deal.Msg/ShipOrder"
	Msg_OrderDelivered_FullMethodName  = "/example.deal.Msg/OrderDelivered"
	Msg_CancelOrder_FullMethodName     = "/example.deal.Msg/CancelOrder"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
	CreateDeal(ctx context.Context, in *MsgCreateDeal, opts ...grpc.CallOption) (*MsgCreateDealResponse, error)
	CreateContract(ctx context.Context, in *MsgCreateContract, opts ...grpc.CallOption) (*MsgCreateContractResponse, error)
	CommitContract(ctx context.Context, in *MsgCommitContract, opts ...grpc.CallOption) (*MsgCommitContractResponse, error)
	ApproveContract(ctx context.Context, in *MsgApproveContract, opts ...grpc.CallOption) (*MsgApproveContractResponse, error)
	ShipOrder(ctx context.Context, in *MsgShipOrder, opts ...grpc.CallOption) (*MsgShipOrderResponse, error)
	OrderDelivered(ctx context.Context, in *MsgOrderDelivered, opts ...grpc.CallOption) (*MsgOrderDeliveredResponse, error)
	CancelOrder(ctx context.Context, in *MsgCancelOrder, opts ...grpc.CallOption) (*MsgCancelOrderResponse, error)
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

func (c *msgClient) CreateDeal(ctx context.Context, in *MsgCreateDeal, opts ...grpc.CallOption) (*MsgCreateDealResponse, error) {
	out := new(MsgCreateDealResponse)
	err := c.cc.Invoke(ctx, Msg_CreateDeal_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CreateContract(ctx context.Context, in *MsgCreateContract, opts ...grpc.CallOption) (*MsgCreateContractResponse, error) {
	out := new(MsgCreateContractResponse)
	err := c.cc.Invoke(ctx, Msg_CreateContract_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CommitContract(ctx context.Context, in *MsgCommitContract, opts ...grpc.CallOption) (*MsgCommitContractResponse, error) {
	out := new(MsgCommitContractResponse)
	err := c.cc.Invoke(ctx, Msg_CommitContract_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ApproveContract(ctx context.Context, in *MsgApproveContract, opts ...grpc.CallOption) (*MsgApproveContractResponse, error) {
	out := new(MsgApproveContractResponse)
	err := c.cc.Invoke(ctx, Msg_ApproveContract_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ShipOrder(ctx context.Context, in *MsgShipOrder, opts ...grpc.CallOption) (*MsgShipOrderResponse, error) {
	out := new(MsgShipOrderResponse)
	err := c.cc.Invoke(ctx, Msg_ShipOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) OrderDelivered(ctx context.Context, in *MsgOrderDelivered, opts ...grpc.CallOption) (*MsgOrderDeliveredResponse, error) {
	out := new(MsgOrderDeliveredResponse)
	err := c.cc.Invoke(ctx, Msg_OrderDelivered_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CancelOrder(ctx context.Context, in *MsgCancelOrder, opts ...grpc.CallOption) (*MsgCancelOrderResponse, error) {
	out := new(MsgCancelOrderResponse)
	err := c.cc.Invoke(ctx, Msg_CancelOrder_FullMethodName, in, out, opts...)
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
	CreateDeal(context.Context, *MsgCreateDeal) (*MsgCreateDealResponse, error)
	CreateContract(context.Context, *MsgCreateContract) (*MsgCreateContractResponse, error)
	CommitContract(context.Context, *MsgCommitContract) (*MsgCommitContractResponse, error)
	ApproveContract(context.Context, *MsgApproveContract) (*MsgApproveContractResponse, error)
	ShipOrder(context.Context, *MsgShipOrder) (*MsgShipOrderResponse, error)
	OrderDelivered(context.Context, *MsgOrderDelivered) (*MsgOrderDeliveredResponse, error)
	CancelOrder(context.Context, *MsgCancelOrder) (*MsgCancelOrderResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}
func (UnimplementedMsgServer) CreateDeal(context.Context, *MsgCreateDeal) (*MsgCreateDealResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDeal not implemented")
}
func (UnimplementedMsgServer) CreateContract(context.Context, *MsgCreateContract) (*MsgCreateContractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateContract not implemented")
}
func (UnimplementedMsgServer) CommitContract(context.Context, *MsgCommitContract) (*MsgCommitContractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommitContract not implemented")
}
func (UnimplementedMsgServer) ApproveContract(context.Context, *MsgApproveContract) (*MsgApproveContractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApproveContract not implemented")
}
func (UnimplementedMsgServer) ShipOrder(context.Context, *MsgShipOrder) (*MsgShipOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShipOrder not implemented")
}
func (UnimplementedMsgServer) OrderDelivered(context.Context, *MsgOrderDelivered) (*MsgOrderDeliveredResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderDelivered not implemented")
}
func (UnimplementedMsgServer) CancelOrder(context.Context, *MsgCancelOrder) (*MsgCancelOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelOrder not implemented")
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

func _Msg_CreateDeal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateDeal)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateDeal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CreateDeal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateDeal(ctx, req.(*MsgCreateDeal))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CreateContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateContract)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CreateContract_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateContract(ctx, req.(*MsgCreateContract))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CommitContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCommitContract)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CommitContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CommitContract_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CommitContract(ctx, req.(*MsgCommitContract))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ApproveContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgApproveContract)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ApproveContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_ApproveContract_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ApproveContract(ctx, req.(*MsgApproveContract))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ShipOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgShipOrder)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ShipOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_ShipOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ShipOrder(ctx, req.(*MsgShipOrder))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_OrderDelivered_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgOrderDelivered)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).OrderDelivered(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_OrderDelivered_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).OrderDelivered(ctx, req.(*MsgOrderDelivered))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CancelOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCancelOrder)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CancelOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CancelOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CancelOrder(ctx, req.(*MsgCancelOrder))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "example.deal.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
		{
			MethodName: "CreateDeal",
			Handler:    _Msg_CreateDeal_Handler,
		},
		{
			MethodName: "CreateContract",
			Handler:    _Msg_CreateContract_Handler,
		},
		{
			MethodName: "CommitContract",
			Handler:    _Msg_CommitContract_Handler,
		},
		{
			MethodName: "ApproveContract",
			Handler:    _Msg_ApproveContract_Handler,
		},
		{
			MethodName: "ShipOrder",
			Handler:    _Msg_ShipOrder_Handler,
		},
		{
			MethodName: "OrderDelivered",
			Handler:    _Msg_OrderDelivered_Handler,
		},
		{
			MethodName: "CancelOrder",
			Handler:    _Msg_CancelOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "example/deal/tx.proto",
}