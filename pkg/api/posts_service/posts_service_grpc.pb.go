// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.2
// source: posts_service.proto

package posts_service

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

// PostsServiceClient is the client API for PostsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PostsServiceClient interface {
	GetAllPosts(ctx context.Context, in *GetAllPostsRequest, opts ...grpc.CallOption) (*GetAllPostsResponse, error)
	GetPost(ctx context.Context, in *GetPostRequest, opts ...grpc.CallOption) (*GetPostResponse, error)
	UpdatePartialPost(ctx context.Context, in *UpdatePartialPostRequest, opts ...grpc.CallOption) (*UpdatePartialPostResponse, error)
	DeletePost(ctx context.Context, in *DeletePostRequest, opts ...grpc.CallOption) (*DeletePostResponse, error)
}

type postsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPostsServiceClient(cc grpc.ClientConnInterface) PostsServiceClient {
	return &postsServiceClient{cc}
}

func (c *postsServiceClient) GetAllPosts(ctx context.Context, in *GetAllPostsRequest, opts ...grpc.CallOption) (*GetAllPostsResponse, error) {
	out := new(GetAllPostsResponse)
	err := c.cc.Invoke(ctx, "/api.PostsService/GetAllPosts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postsServiceClient) GetPost(ctx context.Context, in *GetPostRequest, opts ...grpc.CallOption) (*GetPostResponse, error) {
	out := new(GetPostResponse)
	err := c.cc.Invoke(ctx, "/api.PostsService/GetPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postsServiceClient) UpdatePartialPost(ctx context.Context, in *UpdatePartialPostRequest, opts ...grpc.CallOption) (*UpdatePartialPostResponse, error) {
	out := new(UpdatePartialPostResponse)
	err := c.cc.Invoke(ctx, "/api.PostsService/UpdatePartialPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postsServiceClient) DeletePost(ctx context.Context, in *DeletePostRequest, opts ...grpc.CallOption) (*DeletePostResponse, error) {
	out := new(DeletePostResponse)
	err := c.cc.Invoke(ctx, "/api.PostsService/DeletePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PostsServiceServer is the server API for PostsService service.
// All implementations must embed UnimplementedPostsServiceServer
// for forward compatibility
type PostsServiceServer interface {
	GetAllPosts(context.Context, *GetAllPostsRequest) (*GetAllPostsResponse, error)
	GetPost(context.Context, *GetPostRequest) (*GetPostResponse, error)
	UpdatePartialPost(context.Context, *UpdatePartialPostRequest) (*UpdatePartialPostResponse, error)
	DeletePost(context.Context, *DeletePostRequest) (*DeletePostResponse, error)
	mustEmbedUnimplementedPostsServiceServer()
}

// UnimplementedPostsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPostsServiceServer struct {
}

func (UnimplementedPostsServiceServer) GetAllPosts(context.Context, *GetAllPostsRequest) (*GetAllPostsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllPosts not implemented")
}
func (UnimplementedPostsServiceServer) GetPost(context.Context, *GetPostRequest) (*GetPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPost not implemented")
}
func (UnimplementedPostsServiceServer) UpdatePartialPost(context.Context, *UpdatePartialPostRequest) (*UpdatePartialPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePartialPost not implemented")
}
func (UnimplementedPostsServiceServer) DeletePost(context.Context, *DeletePostRequest) (*DeletePostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePost not implemented")
}
func (UnimplementedPostsServiceServer) mustEmbedUnimplementedPostsServiceServer() {}

// UnsafePostsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PostsServiceServer will
// result in compilation errors.
type UnsafePostsServiceServer interface {
	mustEmbedUnimplementedPostsServiceServer()
}

func RegisterPostsServiceServer(s grpc.ServiceRegistrar, srv PostsServiceServer) {
	s.RegisterService(&PostsService_ServiceDesc, srv)
}

func _PostsService_GetAllPosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllPostsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostsServiceServer).GetAllPosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.PostsService/GetAllPosts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostsServiceServer).GetAllPosts(ctx, req.(*GetAllPostsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostsService_GetPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostsServiceServer).GetPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.PostsService/GetPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostsServiceServer).GetPost(ctx, req.(*GetPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostsService_UpdatePartialPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePartialPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostsServiceServer).UpdatePartialPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.PostsService/UpdatePartialPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostsServiceServer).UpdatePartialPost(ctx, req.(*UpdatePartialPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostsService_DeletePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostsServiceServer).DeletePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.PostsService/DeletePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostsServiceServer).DeletePost(ctx, req.(*DeletePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PostsService_ServiceDesc is the grpc.ServiceDesc for PostsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PostsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.PostsService",
	HandlerType: (*PostsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllPosts",
			Handler:    _PostsService_GetAllPosts_Handler,
		},
		{
			MethodName: "GetPost",
			Handler:    _PostsService_GetPost_Handler,
		},
		{
			MethodName: "UpdatePartialPost",
			Handler:    _PostsService_UpdatePartialPost_Handler,
		},
		{
			MethodName: "DeletePost",
			Handler:    _PostsService_DeletePost_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "posts_service.proto",
}
