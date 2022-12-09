// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.11
// source: config/grpc/config.proto

package service

import (
	context "context"
	config "github.com/Asutorufa/yuhaiin/pkg/protos/config"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ConfigDaoClient is the client API for ConfigDao service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConfigDaoClient interface {
	Load(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*config.Setting, error)
	Save(ctx context.Context, in *config.Setting, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type configDaoClient struct {
	cc grpc.ClientConnInterface
}

func NewConfigDaoClient(cc grpc.ClientConnInterface) ConfigDaoClient {
	return &configDaoClient{cc}
}

func (c *configDaoClient) Load(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*config.Setting, error) {
	out := new(config.Setting)
	err := c.cc.Invoke(ctx, "/yuhaiin.protos.config.service.config_dao/load", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configDaoClient) Save(ctx context.Context, in *config.Setting, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/yuhaiin.protos.config.service.config_dao/save", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConfigDaoServer is the server API for ConfigDao service.
// All implementations must embed UnimplementedConfigDaoServer
// for forward compatibility
type ConfigDaoServer interface {
	Load(context.Context, *emptypb.Empty) (*config.Setting, error)
	Save(context.Context, *config.Setting) (*emptypb.Empty, error)
	mustEmbedUnimplementedConfigDaoServer()
}

// UnimplementedConfigDaoServer must be embedded to have forward compatible implementations.
type UnimplementedConfigDaoServer struct {
}

func (UnimplementedConfigDaoServer) Load(context.Context, *emptypb.Empty) (*config.Setting, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Load not implemented")
}
func (UnimplementedConfigDaoServer) Save(context.Context, *config.Setting) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Save not implemented")
}
func (UnimplementedConfigDaoServer) mustEmbedUnimplementedConfigDaoServer() {}

// UnsafeConfigDaoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConfigDaoServer will
// result in compilation errors.
type UnsafeConfigDaoServer interface {
	mustEmbedUnimplementedConfigDaoServer()
}

func RegisterConfigDaoServer(s grpc.ServiceRegistrar, srv ConfigDaoServer) {
	s.RegisterService(&ConfigDao_ServiceDesc, srv)
}

func _ConfigDao_Load_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigDaoServer).Load(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yuhaiin.protos.config.service.config_dao/load",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigDaoServer).Load(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigDao_Save_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(config.Setting)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigDaoServer).Save(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yuhaiin.protos.config.service.config_dao/save",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigDaoServer).Save(ctx, req.(*config.Setting))
	}
	return interceptor(ctx, in, info, handler)
}

// ConfigDao_ServiceDesc is the grpc.ServiceDesc for ConfigDao service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConfigDao_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yuhaiin.protos.config.service.config_dao",
	HandlerType: (*ConfigDaoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "load",
			Handler:    _ConfigDao_Load_Handler,
		},
		{
			MethodName: "save",
			Handler:    _ConfigDao_Save_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "config/grpc/config.proto",
}
