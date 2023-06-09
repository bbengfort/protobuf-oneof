// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: messages.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EnumMessage_MType int32

const (
	EnumMessage_UNKNOWN EnumMessage_MType = 0
	EnumMessage_INIT    EnumMessage_MType = 1
	EnumMessage_DATA    EnumMessage_MType = 2
	EnumMessage_ACK     EnumMessage_MType = 3
	EnumMessage_NACK    EnumMessage_MType = 4
)

// Enum value maps for EnumMessage_MType.
var (
	EnumMessage_MType_name = map[int32]string{
		0: "UNKNOWN",
		1: "INIT",
		2: "DATA",
		3: "ACK",
		4: "NACK",
	}
	EnumMessage_MType_value = map[string]int32{
		"UNKNOWN": 0,
		"INIT":    1,
		"DATA":    2,
		"ACK":     3,
		"NACK":    4,
	}
)

func (x EnumMessage_MType) Enum() *EnumMessage_MType {
	p := new(EnumMessage_MType)
	*p = x
	return p
}

func (x EnumMessage_MType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EnumMessage_MType) Descriptor() protoreflect.EnumDescriptor {
	return file_messages_proto_enumTypes[0].Descriptor()
}

func (EnumMessage_MType) Type() protoreflect.EnumType {
	return &file_messages_proto_enumTypes[0]
}

func (x EnumMessage_MType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EnumMessage_MType.Descriptor instead.
func (EnumMessage_MType) EnumDescriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{7, 0}
}

type Init struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Region   string `protobuf:"bytes,2,opt,name=region,proto3" json:"region,omitempty"`
}

func (x *Init) Reset() {
	*x = Init{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Init) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Init) ProtoMessage() {}

func (x *Init) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Init.ProtoReflect.Descriptor instead.
func (*Init) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{0}
}

func (x *Init) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *Init) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{1}
}

func (x *Data) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Data) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type Ack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Ts *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=ts,proto3" json:"ts,omitempty"`
}

func (x *Ack) Reset() {
	*x = Ack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ack) ProtoMessage() {}

func (x *Ack) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ack.ProtoReflect.Descriptor instead.
func (*Ack) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{2}
}

func (x *Ack) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Ack) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

type Nack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Code  uint32 `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Error string `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *Nack) Reset() {
	*x = Nack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Nack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Nack) ProtoMessage() {}

func (x *Nack) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Nack.ProtoReflect.Descriptor instead.
func (*Nack) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{3}
}

func (x *Nack) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Nack) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Nack) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type Done struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerId   string `protobuf:"bytes,1,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	TotalBytes int64  `protobuf:"varint,2,opt,name=total_bytes,json=totalBytes,proto3" json:"total_bytes,omitempty"`
	Acks       int64  `protobuf:"varint,3,opt,name=acks,proto3" json:"acks,omitempty"`
	Nacks      int64  `protobuf:"varint,4,opt,name=nacks,proto3" json:"nacks,omitempty"`
	Messages   int64  `protobuf:"varint,5,opt,name=messages,proto3" json:"messages,omitempty"`
}

func (x *Done) Reset() {
	*x = Done{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Done) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Done) ProtoMessage() {}

func (x *Done) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Done.ProtoReflect.Descriptor instead.
func (*Done) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{4}
}

func (x *Done) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

func (x *Done) GetTotalBytes() int64 {
	if x != nil {
		return x.TotalBytes
	}
	return 0
}

func (x *Done) GetAcks() int64 {
	if x != nil {
		return x.Acks
	}
	return 0
}

func (x *Done) GetNacks() int64 {
	if x != nil {
		return x.Nacks
	}
	return 0
}

func (x *Done) GetMessages() int64 {
	if x != nil {
		return x.Messages
	}
	return 0
}

type OnefMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Embed:
	//
	//	*OnefMessage_Init
	//	*OnefMessage_Data
	//	*OnefMessage_Ack
	//	*OnefMessage_Nack
	Embed isOnefMessage_Embed `protobuf_oneof:"embed"`
}

func (x *OnefMessage) Reset() {
	*x = OnefMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OnefMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnefMessage) ProtoMessage() {}

func (x *OnefMessage) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnefMessage.ProtoReflect.Descriptor instead.
func (*OnefMessage) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{5}
}

func (m *OnefMessage) GetEmbed() isOnefMessage_Embed {
	if m != nil {
		return m.Embed
	}
	return nil
}

func (x *OnefMessage) GetInit() *Init {
	if x, ok := x.GetEmbed().(*OnefMessage_Init); ok {
		return x.Init
	}
	return nil
}

func (x *OnefMessage) GetData() *Data {
	if x, ok := x.GetEmbed().(*OnefMessage_Data); ok {
		return x.Data
	}
	return nil
}

func (x *OnefMessage) GetAck() *Ack {
	if x, ok := x.GetEmbed().(*OnefMessage_Ack); ok {
		return x.Ack
	}
	return nil
}

func (x *OnefMessage) GetNack() *Nack {
	if x, ok := x.GetEmbed().(*OnefMessage_Nack); ok {
		return x.Nack
	}
	return nil
}

type isOnefMessage_Embed interface {
	isOnefMessage_Embed()
}

type OnefMessage_Init struct {
	Init *Init `protobuf:"bytes,1,opt,name=init,proto3,oneof"`
}

type OnefMessage_Data struct {
	Data *Data `protobuf:"bytes,2,opt,name=data,proto3,oneof"`
}

type OnefMessage_Ack struct {
	Ack *Ack `protobuf:"bytes,3,opt,name=ack,proto3,oneof"`
}

type OnefMessage_Nack struct {
	Nack *Nack `protobuf:"bytes,4,opt,name=nack,proto3,oneof"`
}

func (*OnefMessage_Init) isOnefMessage_Embed() {}

func (*OnefMessage_Data) isOnefMessage_Embed() {}

func (*OnefMessage_Ack) isOnefMessage_Embed() {}

func (*OnefMessage_Nack) isOnefMessage_Embed() {}

type AnyfMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Embed *anypb.Any `protobuf:"bytes,1,opt,name=embed,proto3" json:"embed,omitempty"`
}

func (x *AnyfMessage) Reset() {
	*x = AnyfMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnyfMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnyfMessage) ProtoMessage() {}

func (x *AnyfMessage) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnyfMessage.ProtoReflect.Descriptor instead.
func (*AnyfMessage) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{6}
}

func (x *AnyfMessage) GetEmbed() *anypb.Any {
	if x != nil {
		return x.Embed
	}
	return nil
}

type EnumMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mtype EnumMessage_MType `protobuf:"varint,1,opt,name=mtype,proto3,enum=oneof.EnumMessage_MType" json:"mtype,omitempty"`
	Embed []byte            `protobuf:"bytes,2,opt,name=embed,proto3" json:"embed,omitempty"`
}

func (x *EnumMessage) Reset() {
	*x = EnumMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnumMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnumMessage) ProtoMessage() {}

func (x *EnumMessage) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnumMessage.ProtoReflect.Descriptor instead.
func (*EnumMessage) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{7}
}

func (x *EnumMessage) GetMtype() EnumMessage_MType {
	if x != nil {
		return x.Mtype
	}
	return EnumMessage_UNKNOWN
}

func (x *EnumMessage) GetEmbed() []byte {
	if x != nil {
		return x.Embed
	}
	return nil
}

var File_messages_proto protoreflect.FileDescriptor

var file_messages_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x3b, 0x0a, 0x04, 0x49, 0x6e, 0x69, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x22, 0x2a, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x41, 0x0a, 0x03,
	0x41, 0x63, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x22,
	0x40, 0x0a, 0x04, 0x4e, 0x61, 0x63, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x22, 0x8a, 0x01, 0x0a, 0x04, 0x44, 0x6f, 0x6e, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x63, 0x6b, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x61, 0x63, 0x6b, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x6e, 0x61, 0x63, 0x6b, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6e, 0x61, 0x63,
	0x6b, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x22, 0x9f,
	0x01, 0x0a, 0x0b, 0x4f, 0x6e, 0x65, 0x66, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x21,
	0x0a, 0x04, 0x69, 0x6e, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x6f,
	0x6e, 0x65, 0x6f, 0x66, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x48, 0x00, 0x52, 0x04, 0x69, 0x6e, 0x69,
	0x74, 0x12, 0x21, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0b, 0x2e, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x1e, 0x0a, 0x03, 0x61, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0a, 0x2e, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x2e, 0x41, 0x63, 0x6b, 0x48, 0x00, 0x52,
	0x03, 0x61, 0x63, 0x6b, 0x12, 0x21, 0x0a, 0x04, 0x6e, 0x61, 0x63, 0x6b, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x2e, 0x4e, 0x61, 0x63, 0x6b, 0x48,
	0x00, 0x52, 0x04, 0x6e, 0x61, 0x63, 0x6b, 0x42, 0x07, 0x0a, 0x05, 0x65, 0x6d, 0x62, 0x65, 0x64,
	0x22, 0x39, 0x0a, 0x0b, 0x41, 0x6e, 0x79, 0x66, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x2a, 0x0a, 0x05, 0x65, 0x6d, 0x62, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05, 0x65, 0x6d, 0x62, 0x65, 0x64, 0x22, 0x90, 0x01, 0x0a, 0x0b,
	0x45, 0x6e, 0x75, 0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2e, 0x0a, 0x05, 0x6d,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x6f, 0x6e, 0x65,
	0x6f, 0x66, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4d,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x05, 0x6d, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x6d, 0x62, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x65, 0x6d, 0x62, 0x65,
	0x64, 0x22, 0x3b, 0x0a, 0x05, 0x4d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x4e, 0x49, 0x54, 0x10,
	0x01, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x41, 0x54, 0x41, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x41,
	0x43, 0x4b, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x41, 0x43, 0x4b, 0x10, 0x04, 0x32, 0x91,
	0x01, 0x0a, 0x08, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x04, 0x4f,
	0x6e, 0x65, 0x66, 0x12, 0x12, 0x2e, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x2e, 0x4f, 0x6e, 0x65, 0x66,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0b, 0x2e, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x2e,
	0x44, 0x6f, 0x6e, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x2b, 0x0a, 0x04, 0x41, 0x6e, 0x79, 0x66,
	0x12, 0x12, 0x2e, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x66, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x1a, 0x0b, 0x2e, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x2e, 0x44, 0x6f, 0x6e,
	0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x2b, 0x0a, 0x04, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x12, 0x2e,
	0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x1a, 0x0b, 0x2e, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x2e, 0x44, 0x6f, 0x6e, 0x65, 0x22, 0x00,
	0x28, 0x01, 0x42, 0x1f, 0x5a, 0x1d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x62, 0x62, 0x65, 0x6e, 0x67, 0x66, 0x6f, 0x72, 0x74, 0x2f, 0x6f, 0x6e, 0x65, 0x6f, 0x66,
	0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_messages_proto_rawDescOnce sync.Once
	file_messages_proto_rawDescData = file_messages_proto_rawDesc
)

func file_messages_proto_rawDescGZIP() []byte {
	file_messages_proto_rawDescOnce.Do(func() {
		file_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_messages_proto_rawDescData)
	})
	return file_messages_proto_rawDescData
}

var file_messages_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_messages_proto_goTypes = []interface{}{
	(EnumMessage_MType)(0),        // 0: oneof.EnumMessage.MType
	(*Init)(nil),                  // 1: oneof.Init
	(*Data)(nil),                  // 2: oneof.Data
	(*Ack)(nil),                   // 3: oneof.Ack
	(*Nack)(nil),                  // 4: oneof.Nack
	(*Done)(nil),                  // 5: oneof.Done
	(*OnefMessage)(nil),           // 6: oneof.OnefMessage
	(*AnyfMessage)(nil),           // 7: oneof.AnyfMessage
	(*EnumMessage)(nil),           // 8: oneof.EnumMessage
	(*timestamppb.Timestamp)(nil), // 9: google.protobuf.Timestamp
	(*anypb.Any)(nil),             // 10: google.protobuf.Any
}
var file_messages_proto_depIdxs = []int32{
	9,  // 0: oneof.Ack.ts:type_name -> google.protobuf.Timestamp
	1,  // 1: oneof.OnefMessage.init:type_name -> oneof.Init
	2,  // 2: oneof.OnefMessage.data:type_name -> oneof.Data
	3,  // 3: oneof.OnefMessage.ack:type_name -> oneof.Ack
	4,  // 4: oneof.OnefMessage.nack:type_name -> oneof.Nack
	10, // 5: oneof.AnyfMessage.embed:type_name -> google.protobuf.Any
	0,  // 6: oneof.EnumMessage.mtype:type_name -> oneof.EnumMessage.MType
	6,  // 7: oneof.Streamer.Onef:input_type -> oneof.OnefMessage
	7,  // 8: oneof.Streamer.Anyf:input_type -> oneof.AnyfMessage
	8,  // 9: oneof.Streamer.Enum:input_type -> oneof.EnumMessage
	5,  // 10: oneof.Streamer.Onef:output_type -> oneof.Done
	5,  // 11: oneof.Streamer.Anyf:output_type -> oneof.Done
	5,  // 12: oneof.Streamer.Enum:output_type -> oneof.Done
	10, // [10:13] is the sub-list for method output_type
	7,  // [7:10] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_messages_proto_init() }
func file_messages_proto_init() {
	if File_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Init); i {
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
		file_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
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
		file_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ack); i {
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
		file_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Nack); i {
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
		file_messages_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Done); i {
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
		file_messages_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OnefMessage); i {
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
		file_messages_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnyfMessage); i {
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
		file_messages_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnumMessage); i {
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
	file_messages_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*OnefMessage_Init)(nil),
		(*OnefMessage_Data)(nil),
		(*OnefMessage_Ack)(nil),
		(*OnefMessage_Nack)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_messages_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_messages_proto_goTypes,
		DependencyIndexes: file_messages_proto_depIdxs,
		EnumInfos:         file_messages_proto_enumTypes,
		MessageInfos:      file_messages_proto_msgTypes,
	}.Build()
	File_messages_proto = out.File
	file_messages_proto_rawDesc = nil
	file_messages_proto_goTypes = nil
	file_messages_proto_depIdxs = nil
}
