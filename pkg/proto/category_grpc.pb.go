// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.0
// source: pkg/proto/category.proto

package proto

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
	Categories_GetCategory_FullMethodName    = "/proto.Categories/GetCategory"
	Categories_GetCategories_FullMethodName  = "/proto.Categories/GetCategories"
	Categories_CreateCategory_FullMethodName = "/proto.Categories/CreateCategory"
	Categories_DeleteCategory_FullMethodName = "/proto.Categories/DeleteCategory"
)

// CategoriesClient is the client API for Categories service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CategoriesClient interface {
	GetCategory(ctx context.Context, in *GetCategoryRequest, opts ...grpc.CallOption) (*GetCategoryResponse, error)
	GetCategories(ctx context.Context, in *GetCategoriesRequest, opts ...grpc.CallOption) (*GetCategoriesResponse, error)
	CreateCategory(ctx context.Context, in *CreateCategoryRequest, opts ...grpc.CallOption) (*CreateCategoryResponse, error)
	DeleteCategory(ctx context.Context, in *DeleteCategoryRequest, opts ...grpc.CallOption) (*DeleteCategoryResponse, error)
}

type categoriesClient struct {
	cc grpc.ClientConnInterface
}

func NewCategoriesClient(cc grpc.ClientConnInterface) CategoriesClient {
	return &categoriesClient{cc}
}

func (c *categoriesClient) GetCategory(ctx context.Context, in *GetCategoryRequest, opts ...grpc.CallOption) (*GetCategoryResponse, error) {
	out := new(GetCategoryResponse)
	err := c.cc.Invoke(ctx, Categories_GetCategory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoriesClient) GetCategories(ctx context.Context, in *GetCategoriesRequest, opts ...grpc.CallOption) (*GetCategoriesResponse, error) {
	out := new(GetCategoriesResponse)
	err := c.cc.Invoke(ctx, Categories_GetCategories_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoriesClient) CreateCategory(ctx context.Context, in *CreateCategoryRequest, opts ...grpc.CallOption) (*CreateCategoryResponse, error) {
	out := new(CreateCategoryResponse)
	err := c.cc.Invoke(ctx, Categories_CreateCategory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoriesClient) DeleteCategory(ctx context.Context, in *DeleteCategoryRequest, opts ...grpc.CallOption) (*DeleteCategoryResponse, error) {
	out := new(DeleteCategoryResponse)
	err := c.cc.Invoke(ctx, Categories_DeleteCategory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CategoriesServer is the server API for Categories service.
// All implementations must embed UnimplementedCategoriesServer
// for forward compatibility
type CategoriesServer interface {
	GetCategory(context.Context, *GetCategoryRequest) (*GetCategoryResponse, error)
	GetCategories(context.Context, *GetCategoriesRequest) (*GetCategoriesResponse, error)
	CreateCategory(context.Context, *CreateCategoryRequest) (*CreateCategoryResponse, error)
	DeleteCategory(context.Context, *DeleteCategoryRequest) (*DeleteCategoryResponse, error)
	mustEmbedUnimplementedCategoriesServer()
}

// UnimplementedCategoriesServer must be embedded to have forward compatible implementations.
type UnimplementedCategoriesServer struct {
}

func (UnimplementedCategoriesServer) GetCategory(context.Context, *GetCategoryRequest) (*GetCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategory not implemented")
}
func (UnimplementedCategoriesServer) GetCategories(context.Context, *GetCategoriesRequest) (*GetCategoriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategories not implemented")
}
func (UnimplementedCategoriesServer) CreateCategory(context.Context, *CreateCategoryRequest) (*CreateCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCategory not implemented")
}
func (UnimplementedCategoriesServer) DeleteCategory(context.Context, *DeleteCategoryRequest) (*DeleteCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCategory not implemented")
}
func (UnimplementedCategoriesServer) mustEmbedUnimplementedCategoriesServer() {}

// UnsafeCategoriesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CategoriesServer will
// result in compilation errors.
type UnsafeCategoriesServer interface {
	mustEmbedUnimplementedCategoriesServer()
}

func RegisterCategoriesServer(s grpc.ServiceRegistrar, srv CategoriesServer) {
	s.RegisterService(&Categories_ServiceDesc, srv)
}

func _Categories_GetCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoriesServer).GetCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Categories_GetCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoriesServer).GetCategory(ctx, req.(*GetCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Categories_GetCategories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCategoriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoriesServer).GetCategories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Categories_GetCategories_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoriesServer).GetCategories(ctx, req.(*GetCategoriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Categories_CreateCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoriesServer).CreateCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Categories_CreateCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoriesServer).CreateCategory(ctx, req.(*CreateCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Categories_DeleteCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoriesServer).DeleteCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Categories_DeleteCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoriesServer).DeleteCategory(ctx, req.(*DeleteCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Categories_ServiceDesc is the grpc.ServiceDesc for Categories service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Categories_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Categories",
	HandlerType: (*CategoriesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCategory",
			Handler:    _Categories_GetCategory_Handler,
		},
		{
			MethodName: "GetCategories",
			Handler:    _Categories_GetCategories_Handler,
		},
		{
			MethodName: "CreateCategory",
			Handler:    _Categories_CreateCategory_Handler,
		},
		{
			MethodName: "DeleteCategory",
			Handler:    _Categories_DeleteCategory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/proto/category.proto",
}
