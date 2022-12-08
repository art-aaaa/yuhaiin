// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.10
// source: node/grpc/node.proto

package service

import (
	context "context"
	node "github.com/Asutorufa/yuhaiin/pkg/protos/node"
	latency "github.com/Asutorufa/yuhaiin/pkg/protos/node/latency"
	point "github.com/Asutorufa/yuhaiin/pkg/protos/node/point"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// NodeClient is the client API for Node service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NodeClient interface {
	Now(ctx context.Context, in *NowReq, opts ...grpc.CallOption) (*point.Point, error)
	// use req is hash string of point
	Use(ctx context.Context, in *UseReq, opts ...grpc.CallOption) (*point.Point, error)
	Get(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*point.Point, error)
	Save(ctx context.Context, in *point.Point, opts ...grpc.CallOption) (*point.Point, error)
	Remove(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Manager(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*node.Manager, error)
	Latency(ctx context.Context, in *latency.Requests, opts ...grpc.CallOption) (*latency.Response, error)
}

type nodeClient struct {
	cc grpc.ClientConnInterface
}

func NewNodeClient(cc grpc.ClientConnInterface) NodeClient {
	return &nodeClient{cc}
}

func (c *nodeClient) Now(ctx context.Context, in *NowReq, opts ...grpc.CallOption) (*point.Point, error) {
	out := new(point.Point)
	err := c.cc.Invoke(ctx, "/yuhaiin.protos.node.service.node/now", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) Use(ctx context.Context, in *UseReq, opts ...grpc.CallOption) (*point.Point, error) {
	out := new(point.Point)
	err := c.cc.Invoke(ctx, "/yuhaiin.protos.node.service.node/use", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) Get(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*point.Point, error) {
	out := new(point.Point)
	err := c.cc.Invoke(ctx, "/yuhaiin.protos.node.service.node/get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) Save(ctx context.Context, in *point.Point, opts ...grpc.CallOption) (*point.Point, error) {
	out := new(point.Point)
	err := c.cc.Invoke(ctx, "/yuhaiin.protos.node.service.node/save", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) Remove(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/yuhaiin.protos.node.service.node/remove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) Manager(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*node.Manager, error) {
	out := new(node.Manager)
	err := c.cc.Invoke(ctx, "/yuhaiin.protos.node.service.node/manager", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) Latency(ctx context.Context, in *latency.Requests, opts ...grpc.CallOption) (*latency.Response, error) {
	out := new(latency.Response)
	err := c.cc.Invoke(ctx, "/yuhaiin.protos.node.service.node/latency", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeServer is the server API for Node service.
// All implementations must embed UnimplementedNodeServer
// for forward compatibility
type NodeServer interface {
	Now(context.Context, *NowReq) (*point.Point, error)
	// use req is hash string of point
	Use(context.Context, *UseReq) (*point.Point, error)
	Get(context.Context, *wrapperspb.StringValue) (*point.Point, error)
	Save(context.Context, *point.Point) (*point.Point, error)
	Remove(context.Context, *wrapperspb.StringValue) (*emptypb.Empty, error)
	Manager(context.Context, *wrapperspb.StringValue) (*node.Manager, error)
	Latency(context.Context, *latency.Requests) (*latency.Response, error)
	mustEmbedUnimplementedNodeServer()
}

// UnimplementedNodeServer must be embedded to have forward compatible implementations.
type UnimplementedNodeServer struct {
}

func (UnimplementedNodeServer) Now(context.Context, *NowReq) (*point.Point, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Now not implemented")
}
func (UnimplementedNodeServer) Use(context.Context, *UseReq) (*point.Point, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Use not implemented")
}
func (UnimplementedNodeServer) Get(context.Context, *wrapperspb.StringValue) (*point.Point, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedNodeServer) Save(context.Context, *point.Point) (*point.Point, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Save not implemented")
}
func (UnimplementedNodeServer) Remove(context.Context, *wrapperspb.StringValue) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Remove not implemented")
}
func (UnimplementedNodeServer) Manager(context.Context, *wrapperspb.StringValue) (*node.Manager, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Manager not implemented")
}
func (UnimplementedNodeServer) Latency(context.Context, *latency.Requests) (*latency.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Latency not implemented")
}
func (UnimplementedNodeServer) mustEmbedUnimplementedNodeServer() {}

// UnsafeNodeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NodeServer will
// result in compilation errors.
type UnsafeNodeServer interface {
	mustEmbedUnimplementedNodeServer()
}

func RegisterNodeServer(s grpc.ServiceRegistrar, srv NodeServer) {
	s.RegisterService(&Node_ServiceDesc, srv)
}

func _Node_Now_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NowReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).Now(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yuhaiin.protos.node.service.node/now",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).Now(ctx, req.(*NowReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_Use_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UseReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).Use(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yuhaiin.protos.node.service.node/use",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).Use(ctx, req.(*UseReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrapperspb.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yuhaiin.protos.node.service.node/get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).Get(ctx, req.(*wrapperspb.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_Save_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(point.Point)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).Save(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yuhaiin.protos.node.service.node/save",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).Save(ctx, req.(*point.Point))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrapperspb.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yuhaiin.protos.node.service.node/remove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).Remove(ctx, req.(*wrapperspb.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_Manager_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrapperspb.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).Manager(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yuhaiin.protos.node.service.node/manager",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).Manager(ctx, req.(*wrapperspb.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_Latency_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(latency.Requests)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).Latency(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yuhaiin.protos.node.service.node/latency",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).Latency(ctx, req.(*latency.Requests))
	}
	return interceptor(ctx, in, info, handler)
}

// Node_ServiceDesc is the grpc.ServiceDesc for Node service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Node_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yuhaiin.protos.node.service.node",
	HandlerType: (*NodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "now",
			Handler:    _Node_Now_Handler,
		},
		{
			MethodName: "use",
			Handler:    _Node_Use_Handler,
		},
		{
			MethodName: "get",
			Handler:    _Node_Get_Handler,
		},
		{
			MethodName: "save",
			Handler:    _Node_Save_Handler,
		},
		{
			MethodName: "remove",
			Handler:    _Node_Remove_Handler,
		},
		{
			MethodName: "manager",
			Handler:    _Node_Manager_Handler,
		},
		{
			MethodName: "latency",
			Handler:    _Node_Latency_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "node/grpc/node.proto",
}

// SubscribeClient is the client API for Subscribe service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SubscribeClient interface {
	Save(ctx context.Context, in *SaveLinkReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Remove(ctx context.Context, in *LinkReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Update(ctx context.Context, in *LinkReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Get(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetLinksResp, error)
}

type subscribeClient struct {
	cc grpc.ClientConnInterface
}

func NewSubscribeClient(cc grpc.ClientConnInterface) SubscribeClient {
	return &subscribeClient{cc}
}

func (c *subscribeClient) Save(ctx context.Context, in *SaveLinkReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/yuhaiin.protos.node.service.subscribe/save", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscribeClient) Remove(ctx context.Context, in *LinkReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/yuhaiin.protos.node.service.subscribe/remove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscribeClient) Update(ctx context.Context, in *LinkReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/yuhaiin.protos.node.service.subscribe/update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscribeClient) Get(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetLinksResp, error) {
	out := new(GetLinksResp)
	err := c.cc.Invoke(ctx, "/yuhaiin.protos.node.service.subscribe/get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SubscribeServer is the server API for Subscribe service.
// All implementations must embed UnimplementedSubscribeServer
// for forward compatibility
type SubscribeServer interface {
	Save(context.Context, *SaveLinkReq) (*emptypb.Empty, error)
	Remove(context.Context, *LinkReq) (*emptypb.Empty, error)
	Update(context.Context, *LinkReq) (*emptypb.Empty, error)
	Get(context.Context, *emptypb.Empty) (*GetLinksResp, error)
	mustEmbedUnimplementedSubscribeServer()
}

// UnimplementedSubscribeServer must be embedded to have forward compatible implementations.
type UnimplementedSubscribeServer struct {
}

func (UnimplementedSubscribeServer) Save(context.Context, *SaveLinkReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Save not implemented")
}
func (UnimplementedSubscribeServer) Remove(context.Context, *LinkReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Remove not implemented")
}
func (UnimplementedSubscribeServer) Update(context.Context, *LinkReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedSubscribeServer) Get(context.Context, *emptypb.Empty) (*GetLinksResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedSubscribeServer) mustEmbedUnimplementedSubscribeServer() {}

// UnsafeSubscribeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SubscribeServer will
// result in compilation errors.
type UnsafeSubscribeServer interface {
	mustEmbedUnimplementedSubscribeServer()
}

func RegisterSubscribeServer(s grpc.ServiceRegistrar, srv SubscribeServer) {
	s.RegisterService(&Subscribe_ServiceDesc, srv)
}

func _Subscribe_Save_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveLinkReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscribeServer).Save(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yuhaiin.protos.node.service.subscribe/save",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscribeServer).Save(ctx, req.(*SaveLinkReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Subscribe_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LinkReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscribeServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yuhaiin.protos.node.service.subscribe/remove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscribeServer).Remove(ctx, req.(*LinkReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Subscribe_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LinkReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscribeServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yuhaiin.protos.node.service.subscribe/update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscribeServer).Update(ctx, req.(*LinkReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Subscribe_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscribeServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yuhaiin.protos.node.service.subscribe/get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscribeServer).Get(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Subscribe_ServiceDesc is the grpc.ServiceDesc for Subscribe service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Subscribe_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yuhaiin.protos.node.service.subscribe",
	HandlerType: (*SubscribeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "save",
			Handler:    _Subscribe_Save_Handler,
		},
		{
			MethodName: "remove",
			Handler:    _Subscribe_Remove_Handler,
		},
		{
			MethodName: "update",
			Handler:    _Subscribe_Update_Handler,
		},
		{
			MethodName: "get",
			Handler:    _Subscribe_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "node/grpc/node.proto",
}
