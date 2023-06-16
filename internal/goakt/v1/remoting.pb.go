// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: goakt/v1/remoting.proto

package goaktv1

import (
	v1 "github.com/tochemey/goakt/messages/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// RemoteAsk is used to send a message to an actor remotely and expect a response
// immediately. With this type of message the receiver cannot communicate back to Sender
// except reply the message with a response. This one-way communication
type RemoteAskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RemoteMessage *v1.RemoteMessage `protobuf:"bytes,1,opt,name=remote_message,json=remoteMessage,proto3" json:"remote_message,omitempty"`
}

func (x *RemoteAskRequest) Reset() {
	*x = RemoteAskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goakt_v1_remoting_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteAskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteAskRequest) ProtoMessage() {}

func (x *RemoteAskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goakt_v1_remoting_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteAskRequest.ProtoReflect.Descriptor instead.
func (*RemoteAskRequest) Descriptor() ([]byte, []int) {
	return file_goakt_v1_remoting_proto_rawDescGZIP(), []int{0}
}

func (x *RemoteAskRequest) GetRemoteMessage() *v1.RemoteMessage {
	if x != nil {
		return x.RemoteMessage
	}
	return nil
}

type RemoteAskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Specifies the message to send to the actor
	// Any proto message is allowed to be sent
	Message *anypb.Any `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *RemoteAskResponse) Reset() {
	*x = RemoteAskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goakt_v1_remoting_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteAskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteAskResponse) ProtoMessage() {}

func (x *RemoteAskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goakt_v1_remoting_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteAskResponse.ProtoReflect.Descriptor instead.
func (*RemoteAskResponse) Descriptor() ([]byte, []int) {
	return file_goakt_v1_remoting_proto_rawDescGZIP(), []int{1}
}

func (x *RemoteAskResponse) GetMessage() *anypb.Any {
	if x != nil {
		return x.Message
	}
	return nil
}

// RemoteTell is used to send a message to an actor remotely by another actor
// This is the only way remote actors can interact with each other. The actor on the
// other line can reply to the sender by using the Sender in the message
type RemoteTellRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RemoteMessage *v1.RemoteMessage `protobuf:"bytes,1,opt,name=remote_message,json=remoteMessage,proto3" json:"remote_message,omitempty"`
}

func (x *RemoteTellRequest) Reset() {
	*x = RemoteTellRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goakt_v1_remoting_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteTellRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteTellRequest) ProtoMessage() {}

func (x *RemoteTellRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goakt_v1_remoting_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteTellRequest.ProtoReflect.Descriptor instead.
func (*RemoteTellRequest) Descriptor() ([]byte, []int) {
	return file_goakt_v1_remoting_proto_rawDescGZIP(), []int{2}
}

func (x *RemoteTellRequest) GetRemoteMessage() *v1.RemoteMessage {
	if x != nil {
		return x.RemoteMessage
	}
	return nil
}

type RemoteTellResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RemoteTellResponse) Reset() {
	*x = RemoteTellResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goakt_v1_remoting_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteTellResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteTellResponse) ProtoMessage() {}

func (x *RemoteTellResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goakt_v1_remoting_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteTellResponse.ProtoReflect.Descriptor instead.
func (*RemoteTellResponse) Descriptor() ([]byte, []int) {
	return file_goakt_v1_remoting_proto_rawDescGZIP(), []int{3}
}

// RemoteLookupRequest checks whether a given actor exists on a remote host
type RemoteLookupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Specifies the remote host address
	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	// Specifies the remote port
	Port int32 `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	// Specifies the actor name
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *RemoteLookupRequest) Reset() {
	*x = RemoteLookupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goakt_v1_remoting_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteLookupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteLookupRequest) ProtoMessage() {}

func (x *RemoteLookupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goakt_v1_remoting_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteLookupRequest.ProtoReflect.Descriptor instead.
func (*RemoteLookupRequest) Descriptor() ([]byte, []int) {
	return file_goakt_v1_remoting_proto_rawDescGZIP(), []int{4}
}

func (x *RemoteLookupRequest) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *RemoteLookupRequest) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *RemoteLookupRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type RemoteLookupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Specifies the actor address
	Address *v1.Address `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *RemoteLookupResponse) Reset() {
	*x = RemoteLookupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goakt_v1_remoting_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteLookupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteLookupResponse) ProtoMessage() {}

func (x *RemoteLookupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goakt_v1_remoting_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteLookupResponse.ProtoReflect.Descriptor instead.
func (*RemoteLookupResponse) Descriptor() ([]byte, []int) {
	return file_goakt_v1_remoting_proto_rawDescGZIP(), []int{5}
}

func (x *RemoteLookupResponse) GetAddress() *v1.Address {
	if x != nil {
		return x.Address
	}
	return nil
}

var File_goakt_v1_remoting_proto protoreflect.FileDescriptor

var file_goakt_v1_remoting_proto_rawDesc = []byte{
	0x0a, 0x17, 0x67, 0x6f, 0x61, 0x6b, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x6d, 0x6f, 0x74,
	0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x67, 0x6f, 0x61, 0x6b, 0x74,
	0x2e, 0x76, 0x31, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x55, 0x0a, 0x10, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x41,
	0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x41, 0x0a, 0x0e, 0x72, 0x65, 0x6d,
	0x6f, 0x74, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0d, 0x72,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x43, 0x0a, 0x11,
	0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x41, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2e, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x56, 0x0a, 0x11, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x54, 0x65, 0x6c, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x41, 0x0a, 0x0e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65,
	0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6d,
	0x6f, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0d, 0x72, 0x65, 0x6d, 0x6f,
	0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x14, 0x0a, 0x12, 0x52, 0x65, 0x6d,
	0x6f, 0x74, 0x65, 0x54, 0x65, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x51, 0x0a, 0x13, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f,
	0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x46, 0x0a, 0x14, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x4c, 0x6f, 0x6f, 0x6b,
	0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x32, 0xf6, 0x01, 0x0a, 0x16, 0x52,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x09, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x41,
	0x73, 0x6b, 0x12, 0x1a, 0x2e, 0x67, 0x6f, 0x61, 0x6b, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x6d, 0x6f, 0x74, 0x65, 0x41, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b,
	0x2e, 0x67, 0x6f, 0x61, 0x6b, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65,
	0x41, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x0a, 0x52,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x54, 0x65, 0x6c, 0x6c, 0x12, 0x1b, 0x2e, 0x67, 0x6f, 0x61, 0x6b,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x54, 0x65, 0x6c, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x67, 0x6f, 0x61, 0x6b, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x54, 0x65, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x0c, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x4c, 0x6f,
	0x6f, 0x6b, 0x75, 0x70, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x61, 0x6b, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x67, 0x6f, 0x61, 0x6b, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x95, 0x01, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x61, 0x6b,
	0x74, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x48, 0x02, 0x50, 0x01, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x6f, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x79, 0x2f, 0x67, 0x6f, 0x61,
	0x6b, 0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x6f, 0x61, 0x6b,
	0x74, 0x2f, 0x76, 0x31, 0x3b, 0x67, 0x6f, 0x61, 0x6b, 0x74, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x47,
	0x58, 0x58, 0xaa, 0x02, 0x08, 0x47, 0x6f, 0x61, 0x6b, 0x74, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x08,
	0x47, 0x6f, 0x61, 0x6b, 0x74, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x14, 0x47, 0x6f, 0x61, 0x6b, 0x74,
	0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x09, 0x47, 0x6f, 0x61, 0x6b, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_goakt_v1_remoting_proto_rawDescOnce sync.Once
	file_goakt_v1_remoting_proto_rawDescData = file_goakt_v1_remoting_proto_rawDesc
)

func file_goakt_v1_remoting_proto_rawDescGZIP() []byte {
	file_goakt_v1_remoting_proto_rawDescOnce.Do(func() {
		file_goakt_v1_remoting_proto_rawDescData = protoimpl.X.CompressGZIP(file_goakt_v1_remoting_proto_rawDescData)
	})
	return file_goakt_v1_remoting_proto_rawDescData
}

var file_goakt_v1_remoting_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_goakt_v1_remoting_proto_goTypes = []interface{}{
	(*RemoteAskRequest)(nil),     // 0: goakt.v1.RemoteAskRequest
	(*RemoteAskResponse)(nil),    // 1: goakt.v1.RemoteAskResponse
	(*RemoteTellRequest)(nil),    // 2: goakt.v1.RemoteTellRequest
	(*RemoteTellResponse)(nil),   // 3: goakt.v1.RemoteTellResponse
	(*RemoteLookupRequest)(nil),  // 4: goakt.v1.RemoteLookupRequest
	(*RemoteLookupResponse)(nil), // 5: goakt.v1.RemoteLookupResponse
	(*v1.RemoteMessage)(nil),     // 6: messages.v1.RemoteMessage
	(*anypb.Any)(nil),            // 7: google.protobuf.Any
	(*v1.Address)(nil),           // 8: messages.v1.Address
}
var file_goakt_v1_remoting_proto_depIdxs = []int32{
	6, // 0: goakt.v1.RemoteAskRequest.remote_message:type_name -> messages.v1.RemoteMessage
	7, // 1: goakt.v1.RemoteAskResponse.message:type_name -> google.protobuf.Any
	6, // 2: goakt.v1.RemoteTellRequest.remote_message:type_name -> messages.v1.RemoteMessage
	8, // 3: goakt.v1.RemoteLookupResponse.address:type_name -> messages.v1.Address
	0, // 4: goakt.v1.RemoteMessagingService.RemoteAsk:input_type -> goakt.v1.RemoteAskRequest
	2, // 5: goakt.v1.RemoteMessagingService.RemoteTell:input_type -> goakt.v1.RemoteTellRequest
	4, // 6: goakt.v1.RemoteMessagingService.RemoteLookup:input_type -> goakt.v1.RemoteLookupRequest
	1, // 7: goakt.v1.RemoteMessagingService.RemoteAsk:output_type -> goakt.v1.RemoteAskResponse
	3, // 8: goakt.v1.RemoteMessagingService.RemoteTell:output_type -> goakt.v1.RemoteTellResponse
	5, // 9: goakt.v1.RemoteMessagingService.RemoteLookup:output_type -> goakt.v1.RemoteLookupResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_goakt_v1_remoting_proto_init() }
func file_goakt_v1_remoting_proto_init() {
	if File_goakt_v1_remoting_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_goakt_v1_remoting_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoteAskRequest); i {
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
		file_goakt_v1_remoting_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoteAskResponse); i {
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
		file_goakt_v1_remoting_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoteTellRequest); i {
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
		file_goakt_v1_remoting_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoteTellResponse); i {
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
		file_goakt_v1_remoting_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoteLookupRequest); i {
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
		file_goakt_v1_remoting_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoteLookupResponse); i {
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
			RawDescriptor: file_goakt_v1_remoting_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_goakt_v1_remoting_proto_goTypes,
		DependencyIndexes: file_goakt_v1_remoting_proto_depIdxs,
		MessageInfos:      file_goakt_v1_remoting_proto_msgTypes,
	}.Build()
	File_goakt_v1_remoting_proto = out.File
	file_goakt_v1_remoting_proto_rawDesc = nil
	file_goakt_v1_remoting_proto_goTypes = nil
	file_goakt_v1_remoting_proto_depIdxs = nil
}