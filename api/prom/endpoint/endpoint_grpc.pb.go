// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: prom/endpoint/endpoint.proto

package endpoint

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
	Endpoint_AppendEndpoint_FullMethodName          = "/api.prom.endpoint.Endpoint/AppendEndpoint"
	Endpoint_DeleteEndpoint_FullMethodName          = "/api.prom.endpoint.Endpoint/DeleteEndpoint"
	Endpoint_EditEndpoint_FullMethodName            = "/api.prom.endpoint.Endpoint/EditEndpoint"
	Endpoint_BatchEditEndpointStatus_FullMethodName = "/api.prom.endpoint.Endpoint/BatchEditEndpointStatus"
	Endpoint_GetEndpoint_FullMethodName             = "/api.prom.endpoint.Endpoint/GetEndpoint"
	Endpoint_ListEndpoint_FullMethodName            = "/api.prom.endpoint.Endpoint/ListEndpoint"
	Endpoint_SelectEndpoint_FullMethodName          = "/api.prom.endpoint.Endpoint/SelectEndpoint"
)

// EndpointClient is the client API for Endpoint service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EndpointClient interface {
	// 添加监控端点
	AppendEndpoint(ctx context.Context, in *AppendEndpointRequest, opts ...grpc.CallOption) (*AppendEndpointReply, error)
	// 删除监控端点
	DeleteEndpoint(ctx context.Context, in *DeleteEndpointRequest, opts ...grpc.CallOption) (*DeleteEndpointReply, error)
	// 编辑端点信息
	EditEndpoint(ctx context.Context, in *EditEndpointRequest, opts ...grpc.CallOption) (*EditEndpointReply, error)
	// 批量修改端点状态
	BatchEditEndpointStatus(ctx context.Context, in *BatchEditEndpointStatusRequest, opts ...grpc.CallOption) (*BatchEditEndpointStatusReply, error)
	// 获取监控端点详情
	GetEndpoint(ctx context.Context, in *GetEndpointRequest, opts ...grpc.CallOption) (*GetEndpointReply, error)
	// 获取监控端点列表
	ListEndpoint(ctx context.Context, in *ListEndpointRequest, opts ...grpc.CallOption) (*ListEndpointReply, error)
	// 获取监控端点下拉列表
	SelectEndpoint(ctx context.Context, in *SelectEndpointRequest, opts ...grpc.CallOption) (*SelectEndpointReply, error)
}

type endpointClient struct {
	cc grpc.ClientConnInterface
}

func NewEndpointClient(cc grpc.ClientConnInterface) EndpointClient {
	return &endpointClient{cc}
}

func (c *endpointClient) AppendEndpoint(ctx context.Context, in *AppendEndpointRequest, opts ...grpc.CallOption) (*AppendEndpointReply, error) {
	out := new(AppendEndpointReply)
	err := c.cc.Invoke(ctx, Endpoint_AppendEndpoint_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *endpointClient) DeleteEndpoint(ctx context.Context, in *DeleteEndpointRequest, opts ...grpc.CallOption) (*DeleteEndpointReply, error) {
	out := new(DeleteEndpointReply)
	err := c.cc.Invoke(ctx, Endpoint_DeleteEndpoint_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *endpointClient) EditEndpoint(ctx context.Context, in *EditEndpointRequest, opts ...grpc.CallOption) (*EditEndpointReply, error) {
	out := new(EditEndpointReply)
	err := c.cc.Invoke(ctx, Endpoint_EditEndpoint_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *endpointClient) BatchEditEndpointStatus(ctx context.Context, in *BatchEditEndpointStatusRequest, opts ...grpc.CallOption) (*BatchEditEndpointStatusReply, error) {
	out := new(BatchEditEndpointStatusReply)
	err := c.cc.Invoke(ctx, Endpoint_BatchEditEndpointStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *endpointClient) GetEndpoint(ctx context.Context, in *GetEndpointRequest, opts ...grpc.CallOption) (*GetEndpointReply, error) {
	out := new(GetEndpointReply)
	err := c.cc.Invoke(ctx, Endpoint_GetEndpoint_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *endpointClient) ListEndpoint(ctx context.Context, in *ListEndpointRequest, opts ...grpc.CallOption) (*ListEndpointReply, error) {
	out := new(ListEndpointReply)
	err := c.cc.Invoke(ctx, Endpoint_ListEndpoint_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *endpointClient) SelectEndpoint(ctx context.Context, in *SelectEndpointRequest, opts ...grpc.CallOption) (*SelectEndpointReply, error) {
	out := new(SelectEndpointReply)
	err := c.cc.Invoke(ctx, Endpoint_SelectEndpoint_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EndpointServer is the server API for Endpoint service.
// All implementations must embed UnimplementedEndpointServer
// for forward compatibility
type EndpointServer interface {
	// 添加监控端点
	AppendEndpoint(context.Context, *AppendEndpointRequest) (*AppendEndpointReply, error)
	// 删除监控端点
	DeleteEndpoint(context.Context, *DeleteEndpointRequest) (*DeleteEndpointReply, error)
	// 编辑端点信息
	EditEndpoint(context.Context, *EditEndpointRequest) (*EditEndpointReply, error)
	// 批量修改端点状态
	BatchEditEndpointStatus(context.Context, *BatchEditEndpointStatusRequest) (*BatchEditEndpointStatusReply, error)
	// 获取监控端点详情
	GetEndpoint(context.Context, *GetEndpointRequest) (*GetEndpointReply, error)
	// 获取监控端点列表
	ListEndpoint(context.Context, *ListEndpointRequest) (*ListEndpointReply, error)
	// 获取监控端点下拉列表
	SelectEndpoint(context.Context, *SelectEndpointRequest) (*SelectEndpointReply, error)
	mustEmbedUnimplementedEndpointServer()
}

// UnimplementedEndpointServer must be embedded to have forward compatible implementations.
type UnimplementedEndpointServer struct {
}

func (UnimplementedEndpointServer) AppendEndpoint(context.Context, *AppendEndpointRequest) (*AppendEndpointReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppendEndpoint not implemented")
}
func (UnimplementedEndpointServer) DeleteEndpoint(context.Context, *DeleteEndpointRequest) (*DeleteEndpointReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEndpoint not implemented")
}
func (UnimplementedEndpointServer) EditEndpoint(context.Context, *EditEndpointRequest) (*EditEndpointReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditEndpoint not implemented")
}
func (UnimplementedEndpointServer) BatchEditEndpointStatus(context.Context, *BatchEditEndpointStatusRequest) (*BatchEditEndpointStatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchEditEndpointStatus not implemented")
}
func (UnimplementedEndpointServer) GetEndpoint(context.Context, *GetEndpointRequest) (*GetEndpointReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEndpoint not implemented")
}
func (UnimplementedEndpointServer) ListEndpoint(context.Context, *ListEndpointRequest) (*ListEndpointReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEndpoint not implemented")
}
func (UnimplementedEndpointServer) SelectEndpoint(context.Context, *SelectEndpointRequest) (*SelectEndpointReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SelectEndpoint not implemented")
}
func (UnimplementedEndpointServer) mustEmbedUnimplementedEndpointServer() {}

// UnsafeEndpointServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EndpointServer will
// result in compilation errors.
type UnsafeEndpointServer interface {
	mustEmbedUnimplementedEndpointServer()
}

func RegisterEndpointServer(s grpc.ServiceRegistrar, srv EndpointServer) {
	s.RegisterService(&Endpoint_ServiceDesc, srv)
}

func _Endpoint_AppendEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppendEndpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndpointServer).AppendEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Endpoint_AppendEndpoint_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndpointServer).AppendEndpoint(ctx, req.(*AppendEndpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Endpoint_DeleteEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEndpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndpointServer).DeleteEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Endpoint_DeleteEndpoint_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndpointServer).DeleteEndpoint(ctx, req.(*DeleteEndpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Endpoint_EditEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditEndpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndpointServer).EditEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Endpoint_EditEndpoint_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndpointServer).EditEndpoint(ctx, req.(*EditEndpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Endpoint_BatchEditEndpointStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchEditEndpointStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndpointServer).BatchEditEndpointStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Endpoint_BatchEditEndpointStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndpointServer).BatchEditEndpointStatus(ctx, req.(*BatchEditEndpointStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Endpoint_GetEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEndpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndpointServer).GetEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Endpoint_GetEndpoint_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndpointServer).GetEndpoint(ctx, req.(*GetEndpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Endpoint_ListEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEndpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndpointServer).ListEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Endpoint_ListEndpoint_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndpointServer).ListEndpoint(ctx, req.(*ListEndpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Endpoint_SelectEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SelectEndpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndpointServer).SelectEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Endpoint_SelectEndpoint_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndpointServer).SelectEndpoint(ctx, req.(*SelectEndpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Endpoint_ServiceDesc is the grpc.ServiceDesc for Endpoint service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Endpoint_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.prom.endpoint.Endpoint",
	HandlerType: (*EndpointServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AppendEndpoint",
			Handler:    _Endpoint_AppendEndpoint_Handler,
		},
		{
			MethodName: "DeleteEndpoint",
			Handler:    _Endpoint_DeleteEndpoint_Handler,
		},
		{
			MethodName: "EditEndpoint",
			Handler:    _Endpoint_EditEndpoint_Handler,
		},
		{
			MethodName: "BatchEditEndpointStatus",
			Handler:    _Endpoint_BatchEditEndpointStatus_Handler,
		},
		{
			MethodName: "GetEndpoint",
			Handler:    _Endpoint_GetEndpoint_Handler,
		},
		{
			MethodName: "ListEndpoint",
			Handler:    _Endpoint_ListEndpoint_Handler,
		},
		{
			MethodName: "SelectEndpoint",
			Handler:    _Endpoint_SelectEndpoint_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "prom/endpoint/endpoint.proto",
}
