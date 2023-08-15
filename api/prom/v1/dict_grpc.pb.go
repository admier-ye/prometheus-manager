// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: prom/v1/dict.proto

package v1

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

// DictClient is the client API for Dict service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DictClient interface {
	// CreateDict creates a new dict.
	CreateDict(ctx context.Context, in *CreateDictRequest, opts ...grpc.CallOption) (*CreateDictReply, error)
	// UpdateDict updates a dict by id.
	UpdateDict(ctx context.Context, in *UpdateDictRequest, opts ...grpc.CallOption) (*UpdateDictReply, error)
	// UpdateDictsStatus updates dicts status by ids.
	UpdateDictsStatus(ctx context.Context, in *UpdateDictsStatusRequest, opts ...grpc.CallOption) (*UpdateDictsStatusReply, error)
	// DeleteDict deletes a dict by id.
	DeleteDict(ctx context.Context, in *DeleteDictRequest, opts ...grpc.CallOption) (*DeleteDictReply, error)
	// GetDict gets a dict by id.
	GetDict(ctx context.Context, in *GetDictRequest, opts ...grpc.CallOption) (*GetDictReply, error)
	// ListDict lists dicts.
	ListDict(ctx context.Context, in *ListDictRequest, opts ...grpc.CallOption) (*ListDictReply, error)
}

type dictClient struct {
	cc grpc.ClientConnInterface
}

func NewDictClient(cc grpc.ClientConnInterface) DictClient {
	return &dictClient{cc}
}

func (c *dictClient) CreateDict(ctx context.Context, in *CreateDictRequest, opts ...grpc.CallOption) (*CreateDictReply, error) {
	out := new(CreateDictReply)
	err := c.cc.Invoke(ctx, "/api.prom.v1.Dict/CreateDict", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dictClient) UpdateDict(ctx context.Context, in *UpdateDictRequest, opts ...grpc.CallOption) (*UpdateDictReply, error) {
	out := new(UpdateDictReply)
	err := c.cc.Invoke(ctx, "/api.prom.v1.Dict/UpdateDict", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dictClient) UpdateDictsStatus(ctx context.Context, in *UpdateDictsStatusRequest, opts ...grpc.CallOption) (*UpdateDictsStatusReply, error) {
	out := new(UpdateDictsStatusReply)
	err := c.cc.Invoke(ctx, "/api.prom.v1.Dict/UpdateDictsStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dictClient) DeleteDict(ctx context.Context, in *DeleteDictRequest, opts ...grpc.CallOption) (*DeleteDictReply, error) {
	out := new(DeleteDictReply)
	err := c.cc.Invoke(ctx, "/api.prom.v1.Dict/DeleteDict", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dictClient) GetDict(ctx context.Context, in *GetDictRequest, opts ...grpc.CallOption) (*GetDictReply, error) {
	out := new(GetDictReply)
	err := c.cc.Invoke(ctx, "/api.prom.v1.Dict/GetDict", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dictClient) ListDict(ctx context.Context, in *ListDictRequest, opts ...grpc.CallOption) (*ListDictReply, error) {
	out := new(ListDictReply)
	err := c.cc.Invoke(ctx, "/api.prom.v1.Dict/ListDict", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DictServer is the server API for Dict service.
// All implementations must embed UnimplementedDictServer
// for forward compatibility
type DictServer interface {
	// CreateDict creates a new dict.
	CreateDict(context.Context, *CreateDictRequest) (*CreateDictReply, error)
	// UpdateDict updates a dict by id.
	UpdateDict(context.Context, *UpdateDictRequest) (*UpdateDictReply, error)
	// UpdateDictsStatus updates dicts status by ids.
	UpdateDictsStatus(context.Context, *UpdateDictsStatusRequest) (*UpdateDictsStatusReply, error)
	// DeleteDict deletes a dict by id.
	DeleteDict(context.Context, *DeleteDictRequest) (*DeleteDictReply, error)
	// GetDict gets a dict by id.
	GetDict(context.Context, *GetDictRequest) (*GetDictReply, error)
	// ListDict lists dicts.
	ListDict(context.Context, *ListDictRequest) (*ListDictReply, error)
	mustEmbedUnimplementedDictServer()
}

// UnimplementedDictServer must be embedded to have forward compatible implementations.
type UnimplementedDictServer struct {
}

func (UnimplementedDictServer) CreateDict(context.Context, *CreateDictRequest) (*CreateDictReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDict not implemented")
}
func (UnimplementedDictServer) UpdateDict(context.Context, *UpdateDictRequest) (*UpdateDictReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDict not implemented")
}
func (UnimplementedDictServer) UpdateDictsStatus(context.Context, *UpdateDictsStatusRequest) (*UpdateDictsStatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDictsStatus not implemented")
}
func (UnimplementedDictServer) DeleteDict(context.Context, *DeleteDictRequest) (*DeleteDictReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDict not implemented")
}
func (UnimplementedDictServer) GetDict(context.Context, *GetDictRequest) (*GetDictReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDict not implemented")
}
func (UnimplementedDictServer) ListDict(context.Context, *ListDictRequest) (*ListDictReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDict not implemented")
}
func (UnimplementedDictServer) mustEmbedUnimplementedDictServer() {}

// UnsafeDictServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DictServer will
// result in compilation errors.
type UnsafeDictServer interface {
	mustEmbedUnimplementedDictServer()
}

func RegisterDictServer(s grpc.ServiceRegistrar, srv DictServer) {
	s.RegisterService(&Dict_ServiceDesc, srv)
}

func _Dict_CreateDict_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDictRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DictServer).CreateDict(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.prom.v1.Dict/CreateDict",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DictServer).CreateDict(ctx, req.(*CreateDictRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dict_UpdateDict_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDictRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DictServer).UpdateDict(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.prom.v1.Dict/UpdateDict",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DictServer).UpdateDict(ctx, req.(*UpdateDictRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dict_UpdateDictsStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDictsStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DictServer).UpdateDictsStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.prom.v1.Dict/UpdateDictsStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DictServer).UpdateDictsStatus(ctx, req.(*UpdateDictsStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dict_DeleteDict_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDictRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DictServer).DeleteDict(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.prom.v1.Dict/DeleteDict",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DictServer).DeleteDict(ctx, req.(*DeleteDictRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dict_GetDict_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDictRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DictServer).GetDict(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.prom.v1.Dict/GetDict",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DictServer).GetDict(ctx, req.(*GetDictRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dict_ListDict_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDictRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DictServer).ListDict(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.prom.v1.Dict/ListDict",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DictServer).ListDict(ctx, req.(*ListDictRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Dict_ServiceDesc is the grpc.ServiceDesc for Dict service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Dict_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.prom.v1.Dict",
	HandlerType: (*DictServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDict",
			Handler:    _Dict_CreateDict_Handler,
		},
		{
			MethodName: "UpdateDict",
			Handler:    _Dict_UpdateDict_Handler,
		},
		{
			MethodName: "UpdateDictsStatus",
			Handler:    _Dict_UpdateDictsStatus_Handler,
		},
		{
			MethodName: "DeleteDict",
			Handler:    _Dict_DeleteDict_Handler,
		},
		{
			MethodName: "GetDict",
			Handler:    _Dict_GetDict_Handler,
		},
		{
			MethodName: "ListDict",
			Handler:    _Dict_ListDict_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "prom/v1/dict.proto",
}
