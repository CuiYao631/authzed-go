// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: authze.proto

package pb

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

type AddPermRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddPermRequest) Reset() {
	*x = AddPermRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authze_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPermRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPermRequest) ProtoMessage() {}

func (x *AddPermRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authze_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPermRequest.ProtoReflect.Descriptor instead.
func (*AddPermRequest) Descriptor() ([]byte, []int) {
	return file_authze_proto_rawDescGZIP(), []int{0}
}

type AddPermResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Test string `protobuf:"bytes,1,opt,name=test,proto3" json:"test,omitempty"`
}

func (x *AddPermResponse) Reset() {
	*x = AddPermResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authze_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPermResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPermResponse) ProtoMessage() {}

func (x *AddPermResponse) ProtoReflect() protoreflect.Message {
	mi := &file_authze_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPermResponse.ProtoReflect.Descriptor instead.
func (*AddPermResponse) Descriptor() ([]byte, []int) {
	return file_authze_proto_rawDescGZIP(), []int{1}
}

func (x *AddPermResponse) GetTest() string {
	if x != nil {
		return x.Test
	}
	return ""
}

type UpdatePermRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdatePermRequest) Reset() {
	*x = UpdatePermRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authze_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePermRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePermRequest) ProtoMessage() {}

func (x *UpdatePermRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authze_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePermRequest.ProtoReflect.Descriptor instead.
func (*UpdatePermRequest) Descriptor() ([]byte, []int) {
	return file_authze_proto_rawDescGZIP(), []int{2}
}

type UpdatePermResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdatePermResponse) Reset() {
	*x = UpdatePermResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authze_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePermResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePermResponse) ProtoMessage() {}

func (x *UpdatePermResponse) ProtoReflect() protoreflect.Message {
	mi := &file_authze_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePermResponse.ProtoReflect.Descriptor instead.
func (*UpdatePermResponse) Descriptor() ([]byte, []int) {
	return file_authze_proto_rawDescGZIP(), []int{3}
}

type DelPermRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DelPermRequest) Reset() {
	*x = DelPermRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authze_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelPermRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelPermRequest) ProtoMessage() {}

func (x *DelPermRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authze_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelPermRequest.ProtoReflect.Descriptor instead.
func (*DelPermRequest) Descriptor() ([]byte, []int) {
	return file_authze_proto_rawDescGZIP(), []int{4}
}

type DelPermResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DelPermResponse) Reset() {
	*x = DelPermResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authze_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelPermResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelPermResponse) ProtoMessage() {}

func (x *DelPermResponse) ProtoReflect() protoreflect.Message {
	mi := &file_authze_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelPermResponse.ProtoReflect.Descriptor instead.
func (*DelPermResponse) Descriptor() ([]byte, []int) {
	return file_authze_proto_rawDescGZIP(), []int{5}
}

type ReadSchemaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReadSchemaRequest) Reset() {
	*x = ReadSchemaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authze_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadSchemaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadSchemaRequest) ProtoMessage() {}

func (x *ReadSchemaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authze_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadSchemaRequest.ProtoReflect.Descriptor instead.
func (*ReadSchemaRequest) Descriptor() ([]byte, []int) {
	return file_authze_proto_rawDescGZIP(), []int{6}
}

type ReadSchemaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Schema string `protobuf:"bytes,1,opt,name=schema,proto3" json:"schema,omitempty"`
}

func (x *ReadSchemaResponse) Reset() {
	*x = ReadSchemaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authze_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadSchemaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadSchemaResponse) ProtoMessage() {}

func (x *ReadSchemaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_authze_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadSchemaResponse.ProtoReflect.Descriptor instead.
func (*ReadSchemaResponse) Descriptor() ([]byte, []int) {
	return file_authze_proto_rawDescGZIP(), []int{7}
}

func (x *ReadSchemaResponse) GetSchema() string {
	if x != nil {
		return x.Schema
	}
	return ""
}

type WriteSchemaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Schema string `protobuf:"bytes,1,opt,name=schema,proto3" json:"schema,omitempty"`
}

func (x *WriteSchemaRequest) Reset() {
	*x = WriteSchemaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authze_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteSchemaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteSchemaRequest) ProtoMessage() {}

func (x *WriteSchemaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authze_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteSchemaRequest.ProtoReflect.Descriptor instead.
func (*WriteSchemaRequest) Descriptor() ([]byte, []int) {
	return file_authze_proto_rawDescGZIP(), []int{8}
}

func (x *WriteSchemaRequest) GetSchema() string {
	if x != nil {
		return x.Schema
	}
	return ""
}

type WriteSchemaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *WriteSchemaResponse) Reset() {
	*x = WriteSchemaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authze_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteSchemaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteSchemaResponse) ProtoMessage() {}

func (x *WriteSchemaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_authze_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteSchemaResponse.ProtoReflect.Descriptor instead.
func (*WriteSchemaResponse) Descriptor() ([]byte, []int) {
	return file_authze_proto_rawDescGZIP(), []int{9}
}

var File_authze_proto protoreflect.FileDescriptor

var file_authze_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a,
	0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x22, 0x10, 0x0a, 0x0e, 0x41, 0x64,
	0x64, 0x50, 0x65, 0x72, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x25, 0x0a, 0x0f,
	0x41, 0x64, 0x64, 0x50, 0x65, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x65, 0x73, 0x74, 0x22, 0x13, 0x0a, 0x11, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x14, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x10,
	0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x50, 0x65, 0x72, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x11, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x50, 0x65, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x13, 0x0a, 0x11, 0x52, 0x65, 0x61, 0x64, 0x53, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2c, 0x0a, 0x12, 0x52, 0x65, 0x61, 0x64,
	0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x22, 0x2c, 0x0a, 0x12, 0x57, 0x72, 0x69, 0x74, 0x65, 0x53,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x22, 0x15, 0x0a, 0x13, 0x57, 0x72, 0x69, 0x74, 0x65, 0x53, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x88, 0x03, 0x0a, 0x0a,
	0x41, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x53, 0x56, 0x43, 0x12, 0x44, 0x0a, 0x07, 0x41, 0x64,
	0x64, 0x50, 0x65, 0x72, 0x6d, 0x12, 0x1a, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x50, 0x65, 0x72, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x64, 0x64, 0x50, 0x65, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x4d, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x12, 0x1d,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x50, 0x65, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x44, 0x0a, 0x07, 0x44, 0x65, 0x6c, 0x50, 0x65, 0x72, 0x6d, 0x12, 0x1a, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x7a, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x50, 0x65, 0x72, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x50, 0x65, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x0a, 0x52, 0x65, 0x61, 0x64, 0x53, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x12, 0x1d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x61, 0x64, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x61, 0x64, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x0b, 0x57, 0x72, 0x69, 0x74, 0x65, 0x53, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x12, 0x1e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2e, 0x76, 0x31,
	0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2e, 0x76, 0x31,
	0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x16, 0x5a, 0x14, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65,
	0x64, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_authze_proto_rawDescOnce sync.Once
	file_authze_proto_rawDescData = file_authze_proto_rawDesc
)

func file_authze_proto_rawDescGZIP() []byte {
	file_authze_proto_rawDescOnce.Do(func() {
		file_authze_proto_rawDescData = protoimpl.X.CompressGZIP(file_authze_proto_rawDescData)
	})
	return file_authze_proto_rawDescData
}

var file_authze_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_authze_proto_goTypes = []interface{}{
	(*AddPermRequest)(nil),      // 0: authzed.v1.AddPermRequest
	(*AddPermResponse)(nil),     // 1: authzed.v1.AddPermResponse
	(*UpdatePermRequest)(nil),   // 2: authzed.v1.updatePermRequest
	(*UpdatePermResponse)(nil),  // 3: authzed.v1.UpdatePermResponse
	(*DelPermRequest)(nil),      // 4: authzed.v1.DelPermRequest
	(*DelPermResponse)(nil),     // 5: authzed.v1.DelPermResponse
	(*ReadSchemaRequest)(nil),   // 6: authzed.v1.ReadSchemaRequest
	(*ReadSchemaResponse)(nil),  // 7: authzed.v1.ReadSchemaResponse
	(*WriteSchemaRequest)(nil),  // 8: authzed.v1.WriteSchemaRequest
	(*WriteSchemaResponse)(nil), // 9: authzed.v1.WriteSchemaResponse
}
var file_authze_proto_depIdxs = []int32{
	0, // 0: authzed.v1.AuthzedSVC.AddPerm:input_type -> authzed.v1.AddPermRequest
	2, // 1: authzed.v1.AuthzedSVC.UpdatePerm:input_type -> authzed.v1.updatePermRequest
	4, // 2: authzed.v1.AuthzedSVC.DelPerm:input_type -> authzed.v1.DelPermRequest
	6, // 3: authzed.v1.AuthzedSVC.ReadSchema:input_type -> authzed.v1.ReadSchemaRequest
	8, // 4: authzed.v1.AuthzedSVC.WriteSchema:input_type -> authzed.v1.WriteSchemaRequest
	1, // 5: authzed.v1.AuthzedSVC.AddPerm:output_type -> authzed.v1.AddPermResponse
	3, // 6: authzed.v1.AuthzedSVC.UpdatePerm:output_type -> authzed.v1.UpdatePermResponse
	5, // 7: authzed.v1.AuthzedSVC.DelPerm:output_type -> authzed.v1.DelPermResponse
	7, // 8: authzed.v1.AuthzedSVC.ReadSchema:output_type -> authzed.v1.ReadSchemaResponse
	9, // 9: authzed.v1.AuthzedSVC.WriteSchema:output_type -> authzed.v1.WriteSchemaResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_authze_proto_init() }
func file_authze_proto_init() {
	if File_authze_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_authze_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddPermRequest); i {
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
		file_authze_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddPermResponse); i {
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
		file_authze_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePermRequest); i {
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
		file_authze_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePermResponse); i {
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
		file_authze_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelPermRequest); i {
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
		file_authze_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelPermResponse); i {
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
		file_authze_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadSchemaRequest); i {
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
		file_authze_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadSchemaResponse); i {
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
		file_authze_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteSchemaRequest); i {
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
		file_authze_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteSchemaResponse); i {
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
			RawDescriptor: file_authze_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_authze_proto_goTypes,
		DependencyIndexes: file_authze_proto_depIdxs,
		MessageInfos:      file_authze_proto_msgTypes,
	}.Build()
	File_authze_proto = out.File
	file_authze_proto_rawDesc = nil
	file_authze_proto_goTypes = nil
	file_authze_proto_depIdxs = nil
}
