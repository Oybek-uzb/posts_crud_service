// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.2
// source: posts_crud_service/posts_crud_service.proto

package posts_crud_service

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Post struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId int32  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Title  string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Body   string `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *Post) Reset() {
	*x = Post{}
	if protoimpl.UnsafeEnabled {
		mi := &file_posts_crud_service_posts_crud_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Post) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Post) ProtoMessage() {}

func (x *Post) ProtoReflect() protoreflect.Message {
	mi := &file_posts_crud_service_posts_crud_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Post.ProtoReflect.Descriptor instead.
func (*Post) Descriptor() ([]byte, []int) {
	return file_posts_crud_service_posts_crud_service_proto_rawDescGZIP(), []int{0}
}

func (x *Post) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Post) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Post) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Post) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

type GetAllPostsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllPostsRequest) Reset() {
	*x = GetAllPostsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_posts_crud_service_posts_crud_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllPostsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllPostsRequest) ProtoMessage() {}

func (x *GetAllPostsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_posts_crud_service_posts_crud_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllPostsRequest.ProtoReflect.Descriptor instead.
func (*GetAllPostsRequest) Descriptor() ([]byte, []int) {
	return file_posts_crud_service_posts_crud_service_proto_rawDescGZIP(), []int{1}
}

type GetAllPostsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Posts []*Post `protobuf:"bytes,1,rep,name=posts,proto3" json:"posts,omitempty"`
}

func (x *GetAllPostsResponse) Reset() {
	*x = GetAllPostsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_posts_crud_service_posts_crud_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllPostsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllPostsResponse) ProtoMessage() {}

func (x *GetAllPostsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_posts_crud_service_posts_crud_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllPostsResponse.ProtoReflect.Descriptor instead.
func (*GetAllPostsResponse) Descriptor() ([]byte, []int) {
	return file_posts_crud_service_posts_crud_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetAllPostsResponse) GetPosts() []*Post {
	if x != nil {
		return x.Posts
	}
	return nil
}

type GetPostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetPostRequest) Reset() {
	*x = GetPostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_posts_crud_service_posts_crud_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostRequest) ProtoMessage() {}

func (x *GetPostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_posts_crud_service_posts_crud_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostRequest.ProtoReflect.Descriptor instead.
func (*GetPostRequest) Descriptor() ([]byte, []int) {
	return file_posts_crud_service_posts_crud_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetPostRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetPostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Post *Post `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
}

func (x *GetPostResponse) Reset() {
	*x = GetPostResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_posts_crud_service_posts_crud_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostResponse) ProtoMessage() {}

func (x *GetPostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_posts_crud_service_posts_crud_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostResponse.ProtoReflect.Descriptor instead.
func (*GetPostResponse) Descriptor() ([]byte, []int) {
	return file_posts_crud_service_posts_crud_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetPostResponse) GetPost() *Post {
	if x != nil {
		return x.Post
	}
	return nil
}

type UpdatePartialPostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UpdateData *Post `protobuf:"bytes,1,opt,name=update_data,json=updateData,proto3" json:"update_data,omitempty"`
}

func (x *UpdatePartialPostRequest) Reset() {
	*x = UpdatePartialPostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_posts_crud_service_posts_crud_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePartialPostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePartialPostRequest) ProtoMessage() {}

func (x *UpdatePartialPostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_posts_crud_service_posts_crud_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePartialPostRequest.ProtoReflect.Descriptor instead.
func (*UpdatePartialPostRequest) Descriptor() ([]byte, []int) {
	return file_posts_crud_service_posts_crud_service_proto_rawDescGZIP(), []int{5}
}

func (x *UpdatePartialPostRequest) GetUpdateData() *Post {
	if x != nil {
		return x.UpdateData
	}
	return nil
}

type UpdatePartialPostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UpdatedPost *Post `protobuf:"bytes,1,opt,name=updated_post,json=updatedPost,proto3" json:"updated_post,omitempty"`
}

func (x *UpdatePartialPostResponse) Reset() {
	*x = UpdatePartialPostResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_posts_crud_service_posts_crud_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePartialPostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePartialPostResponse) ProtoMessage() {}

func (x *UpdatePartialPostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_posts_crud_service_posts_crud_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePartialPostResponse.ProtoReflect.Descriptor instead.
func (*UpdatePartialPostResponse) Descriptor() ([]byte, []int) {
	return file_posts_crud_service_posts_crud_service_proto_rawDescGZIP(), []int{6}
}

func (x *UpdatePartialPostResponse) GetUpdatedPost() *Post {
	if x != nil {
		return x.UpdatedPost
	}
	return nil
}

type DeletePostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeletePostRequest) Reset() {
	*x = DeletePostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_posts_crud_service_posts_crud_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePostRequest) ProtoMessage() {}

func (x *DeletePostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_posts_crud_service_posts_crud_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePostRequest.ProtoReflect.Descriptor instead.
func (*DeletePostRequest) Descriptor() ([]byte, []int) {
	return file_posts_crud_service_posts_crud_service_proto_rawDescGZIP(), []int{7}
}

func (x *DeletePostRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeletePostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeletedPost *Post `protobuf:"bytes,1,opt,name=deleted_post,json=deletedPost,proto3" json:"deleted_post,omitempty"`
}

func (x *DeletePostResponse) Reset() {
	*x = DeletePostResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_posts_crud_service_posts_crud_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePostResponse) ProtoMessage() {}

func (x *DeletePostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_posts_crud_service_posts_crud_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePostResponse.ProtoReflect.Descriptor instead.
func (*DeletePostResponse) Descriptor() ([]byte, []int) {
	return file_posts_crud_service_posts_crud_service_proto_rawDescGZIP(), []int{8}
}

func (x *DeletePostResponse) GetDeletedPost() *Post {
	if x != nil {
		return x.DeletedPost
	}
	return nil
}

var File_posts_crud_service_posts_crud_service_proto protoreflect.FileDescriptor

var file_posts_crud_service_posts_crud_service_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x5f, 0x63, 0x72, 0x75, 0x64, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x5f, 0x63, 0x72, 0x75, 0x64, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x70,
	0x6f, 0x73, 0x74, 0x73, 0x5f, 0x63, 0x72, 0x75, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x22, 0x59, 0x0a, 0x04, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x14, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x45, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x6f, 0x73, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x05, 0x70, 0x6f, 0x73,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x73,
	0x5f, 0x63, 0x72, 0x75, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x6f,
	0x73, 0x74, 0x52, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3f, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c,
	0x0a, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70,
	0x6f, 0x73, 0x74, 0x73, 0x5f, 0x63, 0x72, 0x75, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x22, 0x55, 0x0a, 0x18,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x50, 0x6f, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x70, 0x6f, 0x73, 0x74, 0x73, 0x5f, 0x63, 0x72, 0x75, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x22, 0x58, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72,
	0x74, 0x69, 0x61, 0x6c, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3b, 0x0a, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x70, 0x6f, 0x73, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x5f, 0x63,
	0x72, 0x75, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x6f, 0x73, 0x74,
	0x52, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x22, 0x23, 0x0a,
	0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x51, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x0c, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x5f, 0x63, 0x72, 0x75, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x0b, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x50, 0x6f, 0x73, 0x74, 0x32, 0x9d, 0x03, 0x0a, 0x10, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x43,
	0x52, 0x55, 0x44, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x60, 0x0a, 0x0b, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x26, 0x2e, 0x70, 0x6f, 0x73, 0x74,
	0x73, 0x5f, 0x63, 0x72, 0x75, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x27, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x5f, 0x63, 0x72, 0x75, 0x64, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x6f, 0x73,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x54, 0x0a, 0x07,
	0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x22, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x5f,
	0x63, 0x72, 0x75, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74,
	0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x70, 0x6f,
	0x73, 0x74, 0x73, 0x5f, 0x63, 0x72, 0x75, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x72, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x74,
	0x69, 0x61, 0x6c, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x2c, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x5f,
	0x63, 0x72, 0x75, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x5f, 0x63, 0x72,
	0x75, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x50, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5d, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x50, 0x6f, 0x73, 0x74, 0x12, 0x25, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x5f, 0x63, 0x72, 0x75,
	0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x70, 0x6f,
	0x73, 0x74, 0x73, 0x5f, 0x63, 0x72, 0x75, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x15, 0x5a, 0x13, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x5f,
	0x63, 0x72, 0x75, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_posts_crud_service_posts_crud_service_proto_rawDescOnce sync.Once
	file_posts_crud_service_posts_crud_service_proto_rawDescData = file_posts_crud_service_posts_crud_service_proto_rawDesc
)

func file_posts_crud_service_posts_crud_service_proto_rawDescGZIP() []byte {
	file_posts_crud_service_posts_crud_service_proto_rawDescOnce.Do(func() {
		file_posts_crud_service_posts_crud_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_posts_crud_service_posts_crud_service_proto_rawDescData)
	})
	return file_posts_crud_service_posts_crud_service_proto_rawDescData
}

var file_posts_crud_service_posts_crud_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_posts_crud_service_posts_crud_service_proto_goTypes = []interface{}{
	(*Post)(nil),                      // 0: posts_crud_service.Post
	(*GetAllPostsRequest)(nil),        // 1: posts_crud_service.GetAllPostsRequest
	(*GetAllPostsResponse)(nil),       // 2: posts_crud_service.GetAllPostsResponse
	(*GetPostRequest)(nil),            // 3: posts_crud_service.GetPostRequest
	(*GetPostResponse)(nil),           // 4: posts_crud_service.GetPostResponse
	(*UpdatePartialPostRequest)(nil),  // 5: posts_crud_service.UpdatePartialPostRequest
	(*UpdatePartialPostResponse)(nil), // 6: posts_crud_service.UpdatePartialPostResponse
	(*DeletePostRequest)(nil),         // 7: posts_crud_service.DeletePostRequest
	(*DeletePostResponse)(nil),        // 8: posts_crud_service.DeletePostResponse
}
var file_posts_crud_service_posts_crud_service_proto_depIdxs = []int32{
	0, // 0: posts_crud_service.GetAllPostsResponse.posts:type_name -> posts_crud_service.Post
	0, // 1: posts_crud_service.GetPostResponse.post:type_name -> posts_crud_service.Post
	0, // 2: posts_crud_service.UpdatePartialPostRequest.update_data:type_name -> posts_crud_service.Post
	0, // 3: posts_crud_service.UpdatePartialPostResponse.updated_post:type_name -> posts_crud_service.Post
	0, // 4: posts_crud_service.DeletePostResponse.deleted_post:type_name -> posts_crud_service.Post
	1, // 5: posts_crud_service.PostsCRUDService.GetAllPosts:input_type -> posts_crud_service.GetAllPostsRequest
	3, // 6: posts_crud_service.PostsCRUDService.GetPost:input_type -> posts_crud_service.GetPostRequest
	5, // 7: posts_crud_service.PostsCRUDService.UpdatePartialPost:input_type -> posts_crud_service.UpdatePartialPostRequest
	7, // 8: posts_crud_service.PostsCRUDService.DeletePost:input_type -> posts_crud_service.DeletePostRequest
	2, // 9: posts_crud_service.PostsCRUDService.GetAllPosts:output_type -> posts_crud_service.GetAllPostsResponse
	4, // 10: posts_crud_service.PostsCRUDService.GetPost:output_type -> posts_crud_service.GetPostResponse
	6, // 11: posts_crud_service.PostsCRUDService.UpdatePartialPost:output_type -> posts_crud_service.UpdatePartialPostResponse
	8, // 12: posts_crud_service.PostsCRUDService.DeletePost:output_type -> posts_crud_service.DeletePostResponse
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_posts_crud_service_posts_crud_service_proto_init() }
func file_posts_crud_service_posts_crud_service_proto_init() {
	if File_posts_crud_service_posts_crud_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_posts_crud_service_posts_crud_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Post); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_posts_crud_service_posts_crud_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllPostsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_posts_crud_service_posts_crud_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllPostsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_posts_crud_service_posts_crud_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPostRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_posts_crud_service_posts_crud_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPostResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_posts_crud_service_posts_crud_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePartialPostRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_posts_crud_service_posts_crud_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePartialPostResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_posts_crud_service_posts_crud_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePostRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_posts_crud_service_posts_crud_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePostResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_posts_crud_service_posts_crud_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_posts_crud_service_posts_crud_service_proto_goTypes,
		DependencyIndexes: file_posts_crud_service_posts_crud_service_proto_depIdxs,
		MessageInfos:      file_posts_crud_service_posts_crud_service_proto_msgTypes,
	}.Build()
	File_posts_crud_service_posts_crud_service_proto = out.File
	file_posts_crud_service_posts_crud_service_proto_rawDesc = nil
	file_posts_crud_service_posts_crud_service_proto_goTypes = nil
	file_posts_crud_service_posts_crud_service_proto_depIdxs = nil
}
