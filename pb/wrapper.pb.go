// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: wrapper.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type Flat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	TopicId       string                 `protobuf:"bytes,2,opt,name=topic_id,json=topicId,proto3" json:"topic_id,omitempty"`
	Mimetype      uint32                 `protobuf:"varint,3,opt,name=mimetype,proto3" json:"mimetype,omitempty"`
	Type          *Type                  `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Key           []byte                 `protobuf:"bytes,5,opt,name=key,proto3" json:"key,omitempty"`
	Data          []byte                 `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
	Encryption    *Encryption            `protobuf:"bytes,7,opt,name=encryption,proto3" json:"encryption,omitempty"`
	Compression   *Compression           `protobuf:"bytes,8,opt,name=compression,proto3" json:"compression,omitempty"`
	Geography     *Region                `protobuf:"bytes,9,opt,name=geography,proto3" json:"geography,omitempty"`
	Publisher     *Publisher             `protobuf:"bytes,10,opt,name=publisher,proto3" json:"publisher,omitempty"`
	UserDefinedId string                 `protobuf:"bytes,11,opt,name=user_defined_id,json=userDefinedId,proto3" json:"user_defined_id,omitempty"`
	Created       *timestamppb.Timestamp `protobuf:"bytes,14,opt,name=created,proto3" json:"created,omitempty"`
	Committed     *timestamppb.Timestamp `protobuf:"bytes,15,opt,name=committed,proto3" json:"committed,omitempty"`
}

func (x *Flat) Reset() {
	*x = Flat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wrapper_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Flat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Flat) ProtoMessage() {}

func (x *Flat) ProtoReflect() protoreflect.Message {
	mi := &file_wrapper_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Flat.ProtoReflect.Descriptor instead.
func (*Flat) Descriptor() ([]byte, []int) {
	return file_wrapper_proto_rawDescGZIP(), []int{0}
}

func (x *Flat) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Flat) GetTopicId() string {
	if x != nil {
		return x.TopicId
	}
	return ""
}

func (x *Flat) GetMimetype() uint32 {
	if x != nil {
		return x.Mimetype
	}
	return 0
}

func (x *Flat) GetType() *Type {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *Flat) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *Flat) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Flat) GetEncryption() *Encryption {
	if x != nil {
		return x.Encryption
	}
	return nil
}

func (x *Flat) GetCompression() *Compression {
	if x != nil {
		return x.Compression
	}
	return nil
}

func (x *Flat) GetGeography() *Region {
	if x != nil {
		return x.Geography
	}
	return nil
}

func (x *Flat) GetPublisher() *Publisher {
	if x != nil {
		return x.Publisher
	}
	return nil
}

func (x *Flat) GetUserDefinedId() string {
	if x != nil {
		return x.UserDefinedId
	}
	return ""
}

func (x *Flat) GetCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *Flat) GetCommitted() *timestamppb.Timestamp {
	if x != nil {
		return x.Committed
	}
	return nil
}

type Wrapped struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	TopicId   string                 `protobuf:"bytes,2,opt,name=topic_id,json=topicId,proto3" json:"topic_id,omitempty"`
	Geography *Region                `protobuf:"bytes,3,opt,name=geography,proto3" json:"geography,omitempty"`
	Publisher *Publisher             `protobuf:"bytes,4,opt,name=publisher,proto3" json:"publisher,omitempty"`
	Key       []byte                 `protobuf:"bytes,5,opt,name=key,proto3" json:"key,omitempty"`
	Inner     []byte                 `protobuf:"bytes,6,opt,name=inner,proto3" json:"inner,omitempty"`
	Committed *timestamppb.Timestamp `protobuf:"bytes,15,opt,name=committed,proto3" json:"committed,omitempty"`
}

func (x *Wrapped) Reset() {
	*x = Wrapped{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wrapper_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Wrapped) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Wrapped) ProtoMessage() {}

func (x *Wrapped) ProtoReflect() protoreflect.Message {
	mi := &file_wrapper_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Wrapped.ProtoReflect.Descriptor instead.
func (*Wrapped) Descriptor() ([]byte, []int) {
	return file_wrapper_proto_rawDescGZIP(), []int{1}
}

func (x *Wrapped) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Wrapped) GetTopicId() string {
	if x != nil {
		return x.TopicId
	}
	return ""
}

func (x *Wrapped) GetGeography() *Region {
	if x != nil {
		return x.Geography
	}
	return nil
}

func (x *Wrapped) GetPublisher() *Publisher {
	if x != nil {
		return x.Publisher
	}
	return nil
}

func (x *Wrapped) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *Wrapped) GetInner() []byte {
	if x != nil {
		return x.Inner
	}
	return nil
}

func (x *Wrapped) GetCommitted() *timestamppb.Timestamp {
	if x != nil {
		return x.Committed
	}
	return nil
}

type Inner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mimetype      uint32                 `protobuf:"varint,1,opt,name=mimetype,proto3" json:"mimetype,omitempty"`
	Type          *Type                  `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Encryption    *Encryption            `protobuf:"bytes,3,opt,name=encryption,proto3" json:"encryption,omitempty"`
	Compression   *Compression           `protobuf:"bytes,4,opt,name=compression,proto3" json:"compression,omitempty"`
	UserDefinedId string                 `protobuf:"bytes,5,opt,name=user_defined_id,json=userDefinedId,proto3" json:"user_defined_id,omitempty"`
	Data          []byte                 `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
	Created       *timestamppb.Timestamp `protobuf:"bytes,15,opt,name=created,proto3" json:"created,omitempty"`
}

func (x *Inner) Reset() {
	*x = Inner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wrapper_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Inner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Inner) ProtoMessage() {}

func (x *Inner) ProtoReflect() protoreflect.Message {
	mi := &file_wrapper_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Inner.ProtoReflect.Descriptor instead.
func (*Inner) Descriptor() ([]byte, []int) {
	return file_wrapper_proto_rawDescGZIP(), []int{2}
}

func (x *Inner) GetMimetype() uint32 {
	if x != nil {
		return x.Mimetype
	}
	return 0
}

func (x *Inner) GetType() *Type {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *Inner) GetEncryption() *Encryption {
	if x != nil {
		return x.Encryption
	}
	return nil
}

func (x *Inner) GetCompression() *Compression {
	if x != nil {
		return x.Compression
	}
	return nil
}

func (x *Inner) GetUserDefinedId() string {
	if x != nil {
		return x.UserDefinedId
	}
	return ""
}

func (x *Inner) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Inner) GetCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

type Type struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	MajorVersion uint32 `protobuf:"varint,2,opt,name=major_version,json=majorVersion,proto3" json:"major_version,omitempty"`
	MinorVersion uint32 `protobuf:"varint,3,opt,name=minor_version,json=minorVersion,proto3" json:"minor_version,omitempty"`
	PatchVersion uint32 `protobuf:"varint,4,opt,name=patch_version,json=patchVersion,proto3" json:"patch_version,omitempty"`
}

func (x *Type) Reset() {
	*x = Type{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wrapper_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Type) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Type) ProtoMessage() {}

func (x *Type) ProtoReflect() protoreflect.Message {
	mi := &file_wrapper_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Type.ProtoReflect.Descriptor instead.
func (*Type) Descriptor() ([]byte, []int) {
	return file_wrapper_proto_rawDescGZIP(), []int{3}
}

func (x *Type) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Type) GetMajorVersion() uint32 {
	if x != nil {
		return x.MajorVersion
	}
	return 0
}

func (x *Type) GetMinorVersion() uint32 {
	if x != nil {
		return x.MinorVersion
	}
	return 0
}

func (x *Type) GetPatchVersion() uint32 {
	if x != nil {
		return x.PatchVersion
	}
	return 0
}

type Encryption struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Algorithm     uint32 `protobuf:"varint,1,opt,name=algorithm,proto3" json:"algorithm,omitempty"`
	KeyId         string `protobuf:"bytes,2,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
	EncryptionKey []byte `protobuf:"bytes,3,opt,name=encryption_key,json=encryptionKey,proto3" json:"encryption_key,omitempty"`
	HmacSecret    []byte `protobuf:"bytes,4,opt,name=hmac_secret,json=hmacSecret,proto3" json:"hmac_secret,omitempty"`
	Signature     []byte `protobuf:"bytes,5,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *Encryption) Reset() {
	*x = Encryption{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wrapper_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Encryption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Encryption) ProtoMessage() {}

func (x *Encryption) ProtoReflect() protoreflect.Message {
	mi := &file_wrapper_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Encryption.ProtoReflect.Descriptor instead.
func (*Encryption) Descriptor() ([]byte, []int) {
	return file_wrapper_proto_rawDescGZIP(), []int{4}
}

func (x *Encryption) GetAlgorithm() uint32 {
	if x != nil {
		return x.Algorithm
	}
	return 0
}

func (x *Encryption) GetKeyId() string {
	if x != nil {
		return x.KeyId
	}
	return ""
}

func (x *Encryption) GetEncryptionKey() []byte {
	if x != nil {
		return x.EncryptionKey
	}
	return nil
}

func (x *Encryption) GetHmacSecret() []byte {
	if x != nil {
		return x.HmacSecret
	}
	return nil
}

func (x *Encryption) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type Compression struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Algorithm uint32 `protobuf:"varint,1,opt,name=algorithm,proto3" json:"algorithm,omitempty"`
}

func (x *Compression) Reset() {
	*x = Compression{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wrapper_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Compression) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Compression) ProtoMessage() {}

func (x *Compression) ProtoReflect() protoreflect.Message {
	mi := &file_wrapper_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Compression.ProtoReflect.Descriptor instead.
func (*Compression) Descriptor() ([]byte, []int) {
	return file_wrapper_proto_rawDescGZIP(), []int{5}
}

func (x *Compression) GetAlgorithm() uint32 {
	if x != nil {
		return x.Algorithm
	}
	return 0
}

type Region struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Region  string `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	Zone    string `protobuf:"bytes,2,opt,name=zone,proto3" json:"zone,omitempty"`
	Country string `protobuf:"bytes,3,opt,name=country,proto3" json:"country,omitempty"`
}

func (x *Region) Reset() {
	*x = Region{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wrapper_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Region) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Region) ProtoMessage() {}

func (x *Region) ProtoReflect() protoreflect.Message {
	mi := &file_wrapper_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Region.ProtoReflect.Descriptor instead.
func (*Region) Descriptor() ([]byte, []int) {
	return file_wrapper_proto_rawDescGZIP(), []int{6}
}

func (x *Region) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *Region) GetZone() string {
	if x != nil {
		return x.Zone
	}
	return ""
}

func (x *Region) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

type Publisher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PublisherId string `protobuf:"bytes,1,opt,name=publisher_id,json=publisherId,proto3" json:"publisher_id,omitempty"`
	Ipaddr      string `protobuf:"bytes,2,opt,name=ipaddr,proto3" json:"ipaddr,omitempty"`
	ClientId    string `protobuf:"bytes,3,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	UserAgent   string `protobuf:"bytes,4,opt,name=user_agent,json=userAgent,proto3" json:"user_agent,omitempty"`
}

func (x *Publisher) Reset() {
	*x = Publisher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wrapper_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Publisher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Publisher) ProtoMessage() {}

func (x *Publisher) ProtoReflect() protoreflect.Message {
	mi := &file_wrapper_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Publisher.ProtoReflect.Descriptor instead.
func (*Publisher) Descriptor() ([]byte, []int) {
	return file_wrapper_proto_rawDescGZIP(), []int{7}
}

func (x *Publisher) GetPublisherId() string {
	if x != nil {
		return x.PublisherId
	}
	return ""
}

func (x *Publisher) GetIpaddr() string {
	if x != nil {
		return x.Ipaddr
	}
	return ""
}

func (x *Publisher) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *Publisher) GetUserAgent() string {
	if x != nil {
		return x.UserAgent
	}
	return ""
}

var File_wrapper_proto protoreflect.FileDescriptor

var file_wrapper_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf2, 0x03, 0x0a, 0x04, 0x46, 0x6c, 0x61, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x19, 0x0a, 0x08, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6d,
	0x69, 0x6d, 0x65, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x6d,
	0x69, 0x6d, 0x65, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x2e, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x31,
	0x0a, 0x0a, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x2e, 0x45, 0x6e, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x34, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x2e, 0x43,
	0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70,
	0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x0a, 0x09, 0x67, 0x65, 0x6f, 0x67, 0x72,
	0x61, 0x70, 0x68, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6f, 0x6e, 0x65,
	0x6f, 0x66, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x67, 0x65, 0x6f, 0x67, 0x72,
	0x61, 0x70, 0x68, 0x79, 0x12, 0x2e, 0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65,
	0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x2e,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x66,
	0x69, 0x6e, 0x65, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x75,
	0x73, 0x65, 0x72, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x49, 0x64, 0x12, 0x34, 0x0a, 0x07,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x64, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x64, 0x22, 0xf3, 0x01, 0x0a,
	0x07, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x6f, 0x70, 0x69,
	0x63, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x6f, 0x70, 0x69,
	0x63, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x09, 0x67, 0x65, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x2e, 0x52,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x67, 0x65, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x79,
	0x12, 0x2e, 0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x2e, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x65, 0x72, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x05, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x6d,
	0x69, 0x74, 0x74, 0x65, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74,
	0x65, 0x64, 0x22, 0x9f, 0x02, 0x0a, 0x05, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08,
	0x6d, 0x69, 0x6d, 0x65, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08,
	0x6d, 0x69, 0x6d, 0x65, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x2e, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x31, 0x0a, 0x0a, 0x65, 0x6e, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x2e, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0a, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x0b,
	0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e,
	0x65, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x75, 0x73, 0x65,
	0x72, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x34,
	0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x22, 0x89, 0x01, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x69, 0x6e, 0x6f, 0x72, 0x5f,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x6d,
	0x69, 0x6e, 0x6f, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x70,
	0x61, 0x74, 0x63, 0x68, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0c, 0x70, 0x61, 0x74, 0x63, 0x68, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x22, 0xa7, 0x01, 0x0a, 0x0a, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1c, 0x0a, 0x09, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x09, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x12, 0x15, 0x0a,
	0x06, 0x6b, 0x65, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6b,
	0x65, 0x79, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x65, 0x6e,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x68,
	0x6d, 0x61, 0x63, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x0a, 0x68, 0x6d, 0x61, 0x63, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x2b, 0x0a, 0x0b, 0x43, 0x6f,
	0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x6c, 0x67,
	0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x61, 0x6c,
	0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x22, 0x4e, 0x0a, 0x06, 0x52, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x7a, 0x6f, 0x6e,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x22, 0x82, 0x01, 0x0a, 0x09, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x70, 0x61, 0x64,
	0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x70, 0x61, 0x64, 0x64, 0x72,
	0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x42, 0x1f, 0x5a, 0x1d,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x62, 0x65, 0x6e, 0x67,
	0x66, 0x6f, 0x72, 0x74, 0x2f, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wrapper_proto_rawDescOnce sync.Once
	file_wrapper_proto_rawDescData = file_wrapper_proto_rawDesc
)

func file_wrapper_proto_rawDescGZIP() []byte {
	file_wrapper_proto_rawDescOnce.Do(func() {
		file_wrapper_proto_rawDescData = protoimpl.X.CompressGZIP(file_wrapper_proto_rawDescData)
	})
	return file_wrapper_proto_rawDescData
}

var file_wrapper_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_wrapper_proto_goTypes = []interface{}{
	(*Flat)(nil),                  // 0: oneof.Flat
	(*Wrapped)(nil),               // 1: oneof.Wrapped
	(*Inner)(nil),                 // 2: oneof.Inner
	(*Type)(nil),                  // 3: oneof.Type
	(*Encryption)(nil),            // 4: oneof.Encryption
	(*Compression)(nil),           // 5: oneof.Compression
	(*Region)(nil),                // 6: oneof.Region
	(*Publisher)(nil),             // 7: oneof.Publisher
	(*timestamppb.Timestamp)(nil), // 8: google.protobuf.Timestamp
}
var file_wrapper_proto_depIdxs = []int32{
	3,  // 0: oneof.Flat.type:type_name -> oneof.Type
	4,  // 1: oneof.Flat.encryption:type_name -> oneof.Encryption
	5,  // 2: oneof.Flat.compression:type_name -> oneof.Compression
	6,  // 3: oneof.Flat.geography:type_name -> oneof.Region
	7,  // 4: oneof.Flat.publisher:type_name -> oneof.Publisher
	8,  // 5: oneof.Flat.created:type_name -> google.protobuf.Timestamp
	8,  // 6: oneof.Flat.committed:type_name -> google.protobuf.Timestamp
	6,  // 7: oneof.Wrapped.geography:type_name -> oneof.Region
	7,  // 8: oneof.Wrapped.publisher:type_name -> oneof.Publisher
	8,  // 9: oneof.Wrapped.committed:type_name -> google.protobuf.Timestamp
	3,  // 10: oneof.Inner.type:type_name -> oneof.Type
	4,  // 11: oneof.Inner.encryption:type_name -> oneof.Encryption
	5,  // 12: oneof.Inner.compression:type_name -> oneof.Compression
	8,  // 13: oneof.Inner.created:type_name -> google.protobuf.Timestamp
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_wrapper_proto_init() }
func file_wrapper_proto_init() {
	if File_wrapper_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wrapper_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Flat); i {
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
		file_wrapper_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Wrapped); i {
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
		file_wrapper_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Inner); i {
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
		file_wrapper_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Type); i {
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
		file_wrapper_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Encryption); i {
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
		file_wrapper_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Compression); i {
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
		file_wrapper_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Region); i {
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
		file_wrapper_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Publisher); i {
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
			RawDescriptor: file_wrapper_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wrapper_proto_goTypes,
		DependencyIndexes: file_wrapper_proto_depIdxs,
		MessageInfos:      file_wrapper_proto_msgTypes,
	}.Build()
	File_wrapper_proto = out.File
	file_wrapper_proto_rawDesc = nil
	file_wrapper_proto_goTypes = nil
	file_wrapper_proto_depIdxs = nil
}