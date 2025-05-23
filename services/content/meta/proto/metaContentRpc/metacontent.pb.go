// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.29.0--rc1
// source: metacontent.proto

package metaContentRpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PublishReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId       int64    `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Title        string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	ShortText    string   `protobuf:"bytes,3,opt,name=shortText,proto3" json:"shortText,omitempty"`
	LongTextUri  string   `protobuf:"bytes,4,opt,name=longTextUri,proto3" json:"longTextUri,omitempty"`
	PhotoUriList []string `protobuf:"bytes,5,rep,name=photoUriList,proto3" json:"photoUriList,omitempty"`
	VideoUriList []string `protobuf:"bytes,6,rep,name=videoUriList,proto3" json:"videoUriList,omitempty"`
}

func (x *PublishReq) Reset() {
	*x = PublishReq{}
	mi := &file_metacontent_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PublishReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishReq) ProtoMessage() {}

func (x *PublishReq) ProtoReflect() protoreflect.Message {
	mi := &file_metacontent_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishReq.ProtoReflect.Descriptor instead.
func (*PublishReq) Descriptor() ([]byte, []int) {
	return file_metacontent_proto_rawDescGZIP(), []int{0}
}

func (x *PublishReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *PublishReq) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *PublishReq) GetShortText() string {
	if x != nil {
		return x.ShortText
	}
	return ""
}

func (x *PublishReq) GetLongTextUri() string {
	if x != nil {
		return x.LongTextUri
	}
	return ""
}

func (x *PublishReq) GetPhotoUriList() []string {
	if x != nil {
		return x.PhotoUriList
	}
	return nil
}

func (x *PublishReq) GetVideoUriList() []string {
	if x != nil {
		return x.VideoUriList
	}
	return nil
}

type UpdateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContentId    int64    `protobuf:"varint,1,opt,name=contentId,proto3" json:"contentId,omitempty"`
	UserId       int64    `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	Title        string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	ShortText    string   `protobuf:"bytes,4,opt,name=shortText,proto3" json:"shortText,omitempty"`
	LongTextUri  string   `protobuf:"bytes,5,opt,name=longTextUri,proto3" json:"longTextUri,omitempty"`
	PhotoUriList []string `protobuf:"bytes,6,rep,name=photoUriList,proto3" json:"photoUriList,omitempty"`
	VideoUriList []string `protobuf:"bytes,7,rep,name=videoUriList,proto3" json:"videoUriList,omitempty"`
}

func (x *UpdateReq) Reset() {
	*x = UpdateReq{}
	mi := &file_metacontent_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateReq) ProtoMessage() {}

func (x *UpdateReq) ProtoReflect() protoreflect.Message {
	mi := &file_metacontent_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateReq.ProtoReflect.Descriptor instead.
func (*UpdateReq) Descriptor() ([]byte, []int) {
	return file_metacontent_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateReq) GetContentId() int64 {
	if x != nil {
		return x.ContentId
	}
	return 0
}

func (x *UpdateReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UpdateReq) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateReq) GetShortText() string {
	if x != nil {
		return x.ShortText
	}
	return ""
}

func (x *UpdateReq) GetLongTextUri() string {
	if x != nil {
		return x.LongTextUri
	}
	return ""
}

func (x *UpdateReq) GetPhotoUriList() []string {
	if x != nil {
		return x.PhotoUriList
	}
	return nil
}

func (x *UpdateReq) GetVideoUriList() []string {
	if x != nil {
		return x.VideoUriList
	}
	return nil
}

type DeleteReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContentId int64 `protobuf:"varint,1,opt,name=contentId,proto3" json:"contentId,omitempty"`
	UserId    int64 `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *DeleteReq) Reset() {
	*x = DeleteReq{}
	mi := &file_metacontent_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteReq) ProtoMessage() {}

func (x *DeleteReq) ProtoReflect() protoreflect.Message {
	mi := &file_metacontent_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteReq.ProtoReflect.Descriptor instead.
func (*DeleteReq) Descriptor() ([]byte, []int) {
	return file_metacontent_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteReq) GetContentId() int64 {
	if x != nil {
		return x.ContentId
	}
	return 0
}

func (x *DeleteReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type StatusSearchReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    int64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	ContentId int64 `protobuf:"varint,2,opt,name=contentId,proto3" json:"contentId,omitempty"`
}

func (x *StatusSearchReq) Reset() {
	*x = StatusSearchReq{}
	mi := &file_metacontent_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StatusSearchReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusSearchReq) ProtoMessage() {}

func (x *StatusSearchReq) ProtoReflect() protoreflect.Message {
	mi := &file_metacontent_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusSearchReq.ProtoReflect.Descriptor instead.
func (*StatusSearchReq) Descriptor() ([]byte, []int) {
	return file_metacontent_proto_rawDescGZIP(), []int{3}
}

func (x *StatusSearchReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *StatusSearchReq) GetContentId() int64 {
	if x != nil {
		return x.ContentId
	}
	return 0
}

type StatusSearchResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int32  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Desc   string `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty"`
}

func (x *StatusSearchResp) Reset() {
	*x = StatusSearchResp{}
	mi := &file_metacontent_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StatusSearchResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusSearchResp) ProtoMessage() {}

func (x *StatusSearchResp) ProtoReflect() protoreflect.Message {
	mi := &file_metacontent_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusSearchResp.ProtoReflect.Descriptor instead.
func (*StatusSearchResp) Descriptor() ([]byte, []int) {
	return file_metacontent_proto_rawDescGZIP(), []int{4}
}

func (x *StatusSearchResp) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *StatusSearchResp) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	mi := &file_metacontent_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_metacontent_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_metacontent_proto_rawDescGZIP(), []int{5}
}

var File_metacontent_proto protoreflect.FileDescriptor

var file_metacontent_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6d, 0x65, 0x74, 0x61, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x6d, 0x65, 0x74, 0x61, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x52, 0x70, 0x63, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xc2, 0x01, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x71, 0x12,
	0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x54, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x54, 0x65, 0x78, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x6c,
	0x6f, 0x6e, 0x67, 0x54, 0x65, 0x78, 0x74, 0x55, 0x72, 0x69, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x6c, 0x6f, 0x6e, 0x67, 0x54, 0x65, 0x78, 0x74, 0x55, 0x72, 0x69, 0x12, 0x22, 0x0a,
	0x0c, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x55, 0x72, 0x69, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x55, 0x72, 0x69, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x22, 0x0a, 0x0c, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x55, 0x72, 0x69, 0x4c, 0x69, 0x73,
	0x74, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x55, 0x72,
	0x69, 0x4c, 0x69, 0x73, 0x74, 0x22, 0xdf, 0x01, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x54, 0x65, 0x78, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x54, 0x65, 0x78, 0x74, 0x12, 0x20, 0x0a,
	0x0b, 0x6c, 0x6f, 0x6e, 0x67, 0x54, 0x65, 0x78, 0x74, 0x55, 0x72, 0x69, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x6c, 0x6f, 0x6e, 0x67, 0x54, 0x65, 0x78, 0x74, 0x55, 0x72, 0x69, 0x12,
	0x22, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x55, 0x72, 0x69, 0x4c, 0x69, 0x73, 0x74, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x55, 0x72, 0x69, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x55, 0x72, 0x69, 0x4c,
	0x69, 0x73, 0x74, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x55, 0x72, 0x69, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x41, 0x0a, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x47, 0x0a, 0x0f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x22, 0x3e, 0x0a, 0x10, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64,
	0x65, 0x73, 0x63, 0x22, 0x07, 0x0a, 0x05, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x32, 0x9d, 0x02, 0x0a,
	0x12, 0x4d, 0x65, 0x74, 0x61, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x07, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x12, 0x1a,
	0x2e, 0x6d, 0x65, 0x74, 0x61, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x70, 0x63, 0x2e,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x6d, 0x65, 0x74,
	0x61, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x70, 0x63, 0x2e, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x12, 0x3a, 0x0a, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x19, 0x2e, 0x6d, 0x65,
	0x74, 0x61, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x70, 0x63, 0x2e, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x52, 0x70, 0x63, 0x2e, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3a, 0x0a,
	0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x19, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x70, 0x63, 0x2e, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x1a, 0x15, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x52, 0x70, 0x63, 0x2e, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x51, 0x0a, 0x0c, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x1f, 0x2e, 0x6d, 0x65, 0x74, 0x61,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x70, 0x63, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x1a, 0x20, 0x2e, 0x6d, 0x65, 0x74,
	0x61, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x70, 0x63, 0x2e, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x42, 0x12, 0x5a, 0x10,
	0x2e, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x70, 0x63,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_metacontent_proto_rawDescOnce sync.Once
	file_metacontent_proto_rawDescData = file_metacontent_proto_rawDesc
)

func file_metacontent_proto_rawDescGZIP() []byte {
	file_metacontent_proto_rawDescOnce.Do(func() {
		file_metacontent_proto_rawDescData = protoimpl.X.CompressGZIP(file_metacontent_proto_rawDescData)
	})
	return file_metacontent_proto_rawDescData
}

var file_metacontent_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_metacontent_proto_goTypes = []any{
	(*PublishReq)(nil),       // 0: metaContentRpc.publishReq
	(*UpdateReq)(nil),        // 1: metaContentRpc.updateReq
	(*DeleteReq)(nil),        // 2: metaContentRpc.deleteReq
	(*StatusSearchReq)(nil),  // 3: metaContentRpc.statusSearchReq
	(*StatusSearchResp)(nil), // 4: metaContentRpc.statusSearchResp
	(*Empty)(nil),            // 5: metaContentRpc.empty
}
var file_metacontent_proto_depIdxs = []int32{
	0, // 0: metaContentRpc.MetaContentService.publish:input_type -> metaContentRpc.publishReq
	1, // 1: metaContentRpc.MetaContentService.update:input_type -> metaContentRpc.updateReq
	2, // 2: metaContentRpc.MetaContentService.delete:input_type -> metaContentRpc.deleteReq
	3, // 3: metaContentRpc.MetaContentService.statusSearch:input_type -> metaContentRpc.statusSearchReq
	5, // 4: metaContentRpc.MetaContentService.publish:output_type -> metaContentRpc.empty
	5, // 5: metaContentRpc.MetaContentService.update:output_type -> metaContentRpc.empty
	5, // 6: metaContentRpc.MetaContentService.delete:output_type -> metaContentRpc.empty
	4, // 7: metaContentRpc.MetaContentService.statusSearch:output_type -> metaContentRpc.statusSearchResp
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_metacontent_proto_init() }
func file_metacontent_proto_init() {
	if File_metacontent_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_metacontent_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_metacontent_proto_goTypes,
		DependencyIndexes: file_metacontent_proto_depIdxs,
		MessageInfos:      file_metacontent_proto_msgTypes,
	}.Build()
	File_metacontent_proto = out.File
	file_metacontent_proto_rawDesc = nil
	file_metacontent_proto_goTypes = nil
	file_metacontent_proto_depIdxs = nil
}
