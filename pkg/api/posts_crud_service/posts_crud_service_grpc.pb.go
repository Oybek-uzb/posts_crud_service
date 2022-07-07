// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.2
// source: posts_crud_service/posts_crud_service.proto

package posts_crud_service

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

// PostsCRUDServiceClient is the client API for PostsCRUDService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PostsCRUDServiceClient interface {
	GetAllPosts(ctx context.Context, in *GetAllPostsRequest, opts ...grpc.CallOption) (*GetAllPostsResponse, error)
	GetPost(ctx context.Context, in *GetPostRequest, opts ...grpc.CallOption) (*GetPostResponse, error)
	UpdatePartialPost(ctx context.Context, in *UpdatePartialPostRequest, opts ...grpc.CallOption) (*UpdatePartialPostResponse, error)
	DeletePost(ctx context.Context, in *DeletePostRequest, opts ...grpc.CallOption) (*DeletePostResponse, error)
}

type postsCRUDServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPostsCRUDServiceClient(cc grpc.ClientConnInterface) PostsCRUDServiceClient {
	return &postsCRUDServiceClient{cc}
}

func (c *postsCRUDServiceClient) GetAllPosts(ctx context.Context, in *GetAllPostsRequest, opts ...grpc.CallOption) (*GetAllPostsResponse, error) {
	out := new(GetAllPostsResponse)
	err := c.cc.Invoke(ctx, "/posts_crud_service.PostsCRUDService/GetAllPosts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postsCRUDServiceClient) GetPost(ctx context.Context, in *GetPostRequest, opts ...grpc.CallOption) (*GetPostResponse, error) {
	out := new(GetPostResponse)
	err := c.cc.Invoke(ctx, "/posts_crud_service.PostsCRUDService/GetPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postsCRUDServiceClient) UpdatePartialPost(ctx context.Context, in *UpdatePartialPostRequest, opts ...grpc.CallOption) (*UpdatePartialPostResponse, error) {
	out := new(UpdatePartialPostResponse)
	err := c.cc.Invoke(ctx, "/posts_crud_service.PostsCRUDService/UpdatePartialPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postsCRUDServiceClient) DeletePost(ctx context.Context, in *DeletePostRequest, opts ...grpc.CallOption) (*DeletePostResponse, error) {
	out := new(DeletePostResponse)
	err := c.cc.Invoke(ctx, "/posts_crud_service.PostsCRUDService/DeletePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PostsCRUDServiceServer is the server API for PostsCRUDService service.
// All implementations must embed UnimplementedPostsCRUDServiceServer
// for forward compatibility
type PostsCRUDServiceServer interface {
	GetAllPosts(context.Context, *GetAllPostsRequest) (*GetAllPostsResponse, error)
	GetPost(context.Context, *GetPostRequest) (*GetPostResponse, error)
	UpdatePartialPost(context.Context, *UpdatePartialPostRequest) (*UpdatePartialPostResponse, error)
	DeletePost(context.Context, *DeletePostRequest) (*DeletePostResponse, error)
	mustEmbedUnimplementedPostsCRUDServiceServer()
}

// UnimplementedPostsCRUDServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPostsCRUDServiceServer struct {
}

func (UnimplementedPostsCRUDServiceServer) GetAllPosts(context.Context, *GetAllPostsRequest) (*GetAllPostsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllPosts not implemented")
}
func (UnimplementedPostsCRUDServiceServer) GetPost(context.Context, *GetPostRequest) (*GetPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPost not implemented")
}
func (UnimplementedPostsCRUDServiceServer) UpdatePartialPost(context.Context, *UpdatePartialPostRequest) (*UpdatePartialPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePartialPost not implemented")
}
func (UnimplementedPostsCRUDServiceServer) DeletePost(context.Context, *DeletePostRequest) (*DeletePostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePost not implemented")
}
func (UnimplementedPostsCRUDServiceServer) mustEmbedUnimplementedPostsCRUDServiceServer() {}

// UnsafePostsCRUDServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PostsCRUDServiceServer will
// result in compilation errors.
type UnsafePostsCRUDServiceServer interface {
	mustEmbedUnimplementedPostsCRUDServiceServer()
}

func RegisterPostsCRUDServiceServer(s grpc.ServiceRegistrar, srv PostsCRUDServiceServer) {
	s.RegisterService(&PostsCRUDService_ServiceDesc, srv)
}

func _PostsCRUDService_GetAllPosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllPostsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostsCRUDServiceServer).GetAllPosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/posts_crud_service.PostsCRUDService/GetAllPosts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostsCRUDServiceServer).GetAllPosts(ctx, req.(*GetAllPostsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostsCRUDService_GetPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostsCRUDServiceServer).GetPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/posts_crud_service.PostsCRUDService/GetPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostsCRUDServiceServer).GetPost(ctx, req.(*GetPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostsCRUDService_UpdatePartialPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePartialPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostsCRUDServiceServer).UpdatePartialPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/posts_crud_service.PostsCRUDService/UpdatePartialPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostsCRUDServiceServer).UpdatePartialPost(ctx, req.(*UpdatePartialPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostsCRUDService_DeletePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostsCRUDServiceServer).DeletePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/posts_crud_service.PostsCRUDService/DeletePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostsCRUDServiceServer).DeletePost(ctx, req.(*DeletePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PostsCRUDService_ServiceDesc is the grpc.ServiceDesc for PostsCRUDService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PostsCRUDService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "posts_crud_service.PostsCRUDService",
	HandlerType: (*PostsCRUDServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllPosts",
			Handler:    _PostsCRUDService_GetAllPosts_Handler,
		},
		{
			MethodName: "GetPost",
			Handler:    _PostsCRUDService_GetPost_Handler,
		},
		{
			MethodName: "UpdatePartialPost",
			Handler:    _PostsCRUDService_UpdatePartialPost_Handler,
		},
		{
			MethodName: "DeletePost",
			Handler:    _PostsCRUDService_DeletePost_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "posts_crud_service/posts_crud_service.proto",
}
