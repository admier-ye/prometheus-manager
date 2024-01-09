// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: agent/load.proto

package agent

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

// LoadClient is the client API for Load service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LoadClient interface {
	// 第一次全部加载
	StrategyGroupAll(ctx context.Context, in *StrategyGroupAllRequest, opts ...grpc.CallOption) (*StrategyGroupAllReply, error)
	// 加载变化的规则
	StrategyGroupDiff(ctx context.Context, in *StrategyGroupDiffRequest, opts ...grpc.CallOption) (*StrategyGroupDiffReply, error)
}

type loadClient struct {
	cc grpc.ClientConnInterface
}

func NewLoadClient(cc grpc.ClientConnInterface) LoadClient {
	return &loadClient{cc}
}

func (c *loadClient) StrategyGroupAll(ctx context.Context, in *StrategyGroupAllRequest, opts ...grpc.CallOption) (*StrategyGroupAllReply, error) {
	out := new(StrategyGroupAllReply)
	err := c.cc.Invoke(ctx, "/api.agent.Load/StrategyGroupAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loadClient) StrategyGroupDiff(ctx context.Context, in *StrategyGroupDiffRequest, opts ...grpc.CallOption) (*StrategyGroupDiffReply, error) {
	out := new(StrategyGroupDiffReply)
	err := c.cc.Invoke(ctx, "/api.agent.Load/StrategyGroupDiff", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoadServer is the server API for Load service.
// All implementations must embed UnimplementedLoadServer
// for forward compatibility
type LoadServer interface {
	// 第一次全部加载
	StrategyGroupAll(context.Context, *StrategyGroupAllRequest) (*StrategyGroupAllReply, error)
	// 加载变化的规则
	StrategyGroupDiff(context.Context, *StrategyGroupDiffRequest) (*StrategyGroupDiffReply, error)
	mustEmbedUnimplementedLoadServer()
}

// UnimplementedLoadServer must be embedded to have forward compatible implementations.
type UnimplementedLoadServer struct {
}

func (UnimplementedLoadServer) StrategyGroupAll(context.Context, *StrategyGroupAllRequest) (*StrategyGroupAllReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StrategyGroupAll not implemented")
}
func (UnimplementedLoadServer) StrategyGroupDiff(context.Context, *StrategyGroupDiffRequest) (*StrategyGroupDiffReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StrategyGroupDiff not implemented")
}
func (UnimplementedLoadServer) mustEmbedUnimplementedLoadServer() {}

// UnsafeLoadServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LoadServer will
// result in compilation errors.
type UnsafeLoadServer interface {
	mustEmbedUnimplementedLoadServer()
}

func RegisterLoadServer(s grpc.ServiceRegistrar, srv LoadServer) {
	s.RegisterService(&Load_ServiceDesc, srv)
}

func _Load_StrategyGroupAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StrategyGroupAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoadServer).StrategyGroupAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.agent.Load/StrategyGroupAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoadServer).StrategyGroupAll(ctx, req.(*StrategyGroupAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Load_StrategyGroupDiff_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StrategyGroupDiffRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoadServer).StrategyGroupDiff(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.agent.Load/StrategyGroupDiff",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoadServer).StrategyGroupDiff(ctx, req.(*StrategyGroupDiffRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Load_ServiceDesc is the grpc.ServiceDesc for Load service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Load_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.agent.Load",
	HandlerType: (*LoadServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StrategyGroupAll",
			Handler:    _Load_StrategyGroupAll_Handler,
		},
		{
			MethodName: "StrategyGroupDiff",
			Handler:    _Load_StrategyGroupDiff_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "agent/load.proto",
}
