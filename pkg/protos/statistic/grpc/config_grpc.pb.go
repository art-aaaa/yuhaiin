// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.2
// source: statistic/grpc/config.proto

package service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Connections_Conns_FullMethodName     = "/yuhaiin.protos.statistic.service.connections/conns"
	Connections_CloseConn_FullMethodName = "/yuhaiin.protos.statistic.service.connections/close_conn"
	Connections_Total_FullMethodName     = "/yuhaiin.protos.statistic.service.connections/total"
	Connections_Notify_FullMethodName    = "/yuhaiin.protos.statistic.service.connections/notify"
)

// ConnectionsClient is the client API for Connections service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConnectionsClient interface {
	Conns(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*NotifyNewConnections, error)
	CloseConn(ctx context.Context, in *NotifyRemoveConnections, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Total(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*TotalFlow, error)
	Notify(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Connections_NotifyClient, error)
}

type connectionsClient struct {
	cc grpc.ClientConnInterface
}

func NewConnectionsClient(cc grpc.ClientConnInterface) ConnectionsClient {
	return &connectionsClient{cc}
}

func (c *connectionsClient) Conns(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*NotifyNewConnections, error) {
	out := new(NotifyNewConnections)
	err := c.cc.Invoke(ctx, Connections_Conns_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionsClient) CloseConn(ctx context.Context, in *NotifyRemoveConnections, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Connections_CloseConn_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionsClient) Total(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*TotalFlow, error) {
	out := new(TotalFlow)
	err := c.cc.Invoke(ctx, Connections_Total_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionsClient) Notify(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Connections_NotifyClient, error) {
	stream, err := c.cc.NewStream(ctx, &Connections_ServiceDesc.Streams[0], Connections_Notify_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &connectionsNotifyClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Connections_NotifyClient interface {
	Recv() (*NotifyData, error)
	grpc.ClientStream
}

type connectionsNotifyClient struct {
	grpc.ClientStream
}

func (x *connectionsNotifyClient) Recv() (*NotifyData, error) {
	m := new(NotifyData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ConnectionsServer is the server API for Connections service.
// All implementations must embed UnimplementedConnectionsServer
// for forward compatibility
type ConnectionsServer interface {
	Conns(context.Context, *emptypb.Empty) (*NotifyNewConnections, error)
	CloseConn(context.Context, *NotifyRemoveConnections) (*emptypb.Empty, error)
	Total(context.Context, *emptypb.Empty) (*TotalFlow, error)
	Notify(*emptypb.Empty, Connections_NotifyServer) error
	mustEmbedUnimplementedConnectionsServer()
}

// UnimplementedConnectionsServer must be embedded to have forward compatible implementations.
type UnimplementedConnectionsServer struct {
}

func (UnimplementedConnectionsServer) Conns(context.Context, *emptypb.Empty) (*NotifyNewConnections, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Conns not implemented")
}
func (UnimplementedConnectionsServer) CloseConn(context.Context, *NotifyRemoveConnections) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseConn not implemented")
}
func (UnimplementedConnectionsServer) Total(context.Context, *emptypb.Empty) (*TotalFlow, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Total not implemented")
}
func (UnimplementedConnectionsServer) Notify(*emptypb.Empty, Connections_NotifyServer) error {
	return status.Errorf(codes.Unimplemented, "method Notify not implemented")
}
func (UnimplementedConnectionsServer) mustEmbedUnimplementedConnectionsServer() {}

// UnsafeConnectionsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConnectionsServer will
// result in compilation errors.
type UnsafeConnectionsServer interface {
	mustEmbedUnimplementedConnectionsServer()
}

func RegisterConnectionsServer(s grpc.ServiceRegistrar, srv ConnectionsServer) {
	s.RegisterService(&Connections_ServiceDesc, srv)
}

func _Connections_Conns_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionsServer).Conns(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Connections_Conns_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionsServer).Conns(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Connections_CloseConn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotifyRemoveConnections)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionsServer).CloseConn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Connections_CloseConn_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionsServer).CloseConn(ctx, req.(*NotifyRemoveConnections))
	}
	return interceptor(ctx, in, info, handler)
}

func _Connections_Total_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionsServer).Total(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Connections_Total_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionsServer).Total(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Connections_Notify_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ConnectionsServer).Notify(m, &connectionsNotifyServer{stream})
}

type Connections_NotifyServer interface {
	Send(*NotifyData) error
	grpc.ServerStream
}

type connectionsNotifyServer struct {
	grpc.ServerStream
}

func (x *connectionsNotifyServer) Send(m *NotifyData) error {
	return x.ServerStream.SendMsg(m)
}

// Connections_ServiceDesc is the grpc.ServiceDesc for Connections service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Connections_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yuhaiin.protos.statistic.service.connections",
	HandlerType: (*ConnectionsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "conns",
			Handler:    _Connections_Conns_Handler,
		},
		{
			MethodName: "close_conn",
			Handler:    _Connections_CloseConn_Handler,
		},
		{
			MethodName: "total",
			Handler:    _Connections_Total_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "notify",
			Handler:       _Connections_Notify_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "statistic/grpc/config.proto",
}
