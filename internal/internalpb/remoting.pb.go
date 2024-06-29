// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: internal/remoting.proto

package internalpb

import (
	goaktpb "github.com/tochemey/goakt/v2/goaktpb"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
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
// immediately.
type RemoteAskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Specifies the remote message to send
	RemoteMessage *RemoteMessage `protobuf:"bytes,1,opt,name=remote_message,json=remoteMessage,proto3" json:"remote_message,omitempty"`
	// Specifies the timeout(how long to wait for a reply)
	Timeout *durationpb.Duration `protobuf:"bytes,2,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

func (x *RemoteAskRequest) Reset() {
	*x = RemoteAskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_remoting_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteAskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteAskRequest) ProtoMessage() {}

func (x *RemoteAskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_remoting_proto_msgTypes[0]
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
	return file_internal_remoting_proto_rawDescGZIP(), []int{0}
}

func (x *RemoteAskRequest) GetRemoteMessage() *RemoteMessage {
	if x != nil {
		return x.RemoteMessage
	}
	return nil
}

func (x *RemoteAskRequest) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
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
		mi := &file_internal_remoting_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteAskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteAskResponse) ProtoMessage() {}

func (x *RemoteAskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_remoting_proto_msgTypes[1]
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
	return file_internal_remoting_proto_rawDescGZIP(), []int{1}
}

func (x *RemoteAskResponse) GetMessage() *anypb.Any {
	if x != nil {
		return x.Message
	}
	return nil
}

// RemoteTell is used to send a message to an actor remotely
type RemoteTellRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Specifies the remote message to send
	RemoteMessage *RemoteMessage `protobuf:"bytes,1,opt,name=remote_message,json=remoteMessage,proto3" json:"remote_message,omitempty"`
}

func (x *RemoteTellRequest) Reset() {
	*x = RemoteTellRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_remoting_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteTellRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteTellRequest) ProtoMessage() {}

func (x *RemoteTellRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_remoting_proto_msgTypes[2]
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
	return file_internal_remoting_proto_rawDescGZIP(), []int{2}
}

func (x *RemoteTellRequest) GetRemoteMessage() *RemoteMessage {
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
		mi := &file_internal_remoting_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteTellResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteTellResponse) ProtoMessage() {}

func (x *RemoteTellResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_remoting_proto_msgTypes[3]
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
	return file_internal_remoting_proto_rawDescGZIP(), []int{3}
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
		mi := &file_internal_remoting_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteLookupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteLookupRequest) ProtoMessage() {}

func (x *RemoteLookupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_remoting_proto_msgTypes[4]
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
	return file_internal_remoting_proto_rawDescGZIP(), []int{4}
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
	Address *goaktpb.Address `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *RemoteLookupResponse) Reset() {
	*x = RemoteLookupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_remoting_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteLookupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteLookupResponse) ProtoMessage() {}

func (x *RemoteLookupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_remoting_proto_msgTypes[5]
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
	return file_internal_remoting_proto_rawDescGZIP(), []int{5}
}

func (x *RemoteLookupResponse) GetAddress() *goaktpb.Address {
	if x != nil {
		return x.Address
	}
	return nil
}

// RemoteMessage will be used by Actors to communicate remotely
type RemoteMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Specifies the sender' address
	Sender *goaktpb.Address `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	// Specifies the actor address
	Receiver *goaktpb.Address `protobuf:"bytes,2,opt,name=receiver,proto3" json:"receiver,omitempty"`
	// Specifies the message to send to the actor
	// Any proto message is allowed to be sent
	Message *anypb.Any `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *RemoteMessage) Reset() {
	*x = RemoteMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_remoting_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteMessage) ProtoMessage() {}

func (x *RemoteMessage) ProtoReflect() protoreflect.Message {
	mi := &file_internal_remoting_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteMessage.ProtoReflect.Descriptor instead.
func (*RemoteMessage) Descriptor() ([]byte, []int) {
	return file_internal_remoting_proto_rawDescGZIP(), []int{6}
}

func (x *RemoteMessage) GetSender() *goaktpb.Address {
	if x != nil {
		return x.Sender
	}
	return nil
}

func (x *RemoteMessage) GetReceiver() *goaktpb.Address {
	if x != nil {
		return x.Receiver
	}
	return nil
}

func (x *RemoteMessage) GetMessage() *anypb.Any {
	if x != nil {
		return x.Message
	}
	return nil
}

type RemoteReSpawnRequest struct {
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

func (x *RemoteReSpawnRequest) Reset() {
	*x = RemoteReSpawnRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_remoting_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteReSpawnRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteReSpawnRequest) ProtoMessage() {}

func (x *RemoteReSpawnRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_remoting_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteReSpawnRequest.ProtoReflect.Descriptor instead.
func (*RemoteReSpawnRequest) Descriptor() ([]byte, []int) {
	return file_internal_remoting_proto_rawDescGZIP(), []int{7}
}

func (x *RemoteReSpawnRequest) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *RemoteReSpawnRequest) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *RemoteReSpawnRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type RemoteReSpawnResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RemoteReSpawnResponse) Reset() {
	*x = RemoteReSpawnResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_remoting_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteReSpawnResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteReSpawnResponse) ProtoMessage() {}

func (x *RemoteReSpawnResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_remoting_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteReSpawnResponse.ProtoReflect.Descriptor instead.
func (*RemoteReSpawnResponse) Descriptor() ([]byte, []int) {
	return file_internal_remoting_proto_rawDescGZIP(), []int{8}
}

type RemoteStopRequest struct {
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

func (x *RemoteStopRequest) Reset() {
	*x = RemoteStopRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_remoting_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteStopRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteStopRequest) ProtoMessage() {}

func (x *RemoteStopRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_remoting_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteStopRequest.ProtoReflect.Descriptor instead.
func (*RemoteStopRequest) Descriptor() ([]byte, []int) {
	return file_internal_remoting_proto_rawDescGZIP(), []int{9}
}

func (x *RemoteStopRequest) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *RemoteStopRequest) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *RemoteStopRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type RemoteStopResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RemoteStopResponse) Reset() {
	*x = RemoteStopResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_remoting_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteStopResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteStopResponse) ProtoMessage() {}

func (x *RemoteStopResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_remoting_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteStopResponse.ProtoReflect.Descriptor instead.
func (*RemoteStopResponse) Descriptor() ([]byte, []int) {
	return file_internal_remoting_proto_rawDescGZIP(), []int{10}
}

type RemoteSpawnRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Specifies the remote host address
	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	// Specifies the remote port
	Port int32 `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	// Specifies the actor name.
	ActorName string `protobuf:"bytes,3,opt,name=actor_name,json=actorName,proto3" json:"actor_name,omitempty"`
	// Specifies the actor type
	ActorType string `protobuf:"bytes,4,opt,name=actor_type,json=actorType,proto3" json:"actor_type,omitempty"`
}

func (x *RemoteSpawnRequest) Reset() {
	*x = RemoteSpawnRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_remoting_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteSpawnRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteSpawnRequest) ProtoMessage() {}

func (x *RemoteSpawnRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_remoting_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteSpawnRequest.ProtoReflect.Descriptor instead.
func (*RemoteSpawnRequest) Descriptor() ([]byte, []int) {
	return file_internal_remoting_proto_rawDescGZIP(), []int{11}
}

func (x *RemoteSpawnRequest) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *RemoteSpawnRequest) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *RemoteSpawnRequest) GetActorName() string {
	if x != nil {
		return x.ActorName
	}
	return ""
}

func (x *RemoteSpawnRequest) GetActorType() string {
	if x != nil {
		return x.ActorType
	}
	return ""
}

type RemoteSpawnResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RemoteSpawnResponse) Reset() {
	*x = RemoteSpawnResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_remoting_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteSpawnResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteSpawnResponse) ProtoMessage() {}

func (x *RemoteSpawnResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_remoting_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteSpawnResponse.ProtoReflect.Descriptor instead.
func (*RemoteSpawnResponse) Descriptor() ([]byte, []int) {
	return file_internal_remoting_proto_rawDescGZIP(), []int{12}
}

var File_internal_remoting_proto protoreflect.FileDescriptor

var file_internal_remoting_proto_rawDesc = []byte{
	0x0a, 0x17, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x72, 0x65, 0x6d, 0x6f, 0x74,
	0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x70, 0x62, 0x1a, 0x11, 0x67, 0x6f, 0x61, 0x6b, 0x74, 0x2f, 0x67, 0x6f, 0x61,
	0x6b, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x89, 0x01, 0x0a, 0x10, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x41, 0x73,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x40, 0x0a, 0x0e, 0x72, 0x65, 0x6d, 0x6f,
	0x74, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x70, 0x62, 0x2e, 0x52, 0x65,
	0x6d, 0x6f, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0d, 0x72, 0x65, 0x6d,
	0x6f, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x33, 0x0a, 0x07, 0x74, 0x69,
	0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x22,
	0x43, 0x0a, 0x11, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x41, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x55, 0x0a, 0x11, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x54, 0x65,
	0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x40, 0x0a, 0x0e, 0x72, 0x65, 0x6d,
	0x6f, 0x74, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x70, 0x62, 0x2e, 0x52,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0d, 0x72, 0x65,
	0x6d, 0x6f, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x14, 0x0a, 0x12, 0x52,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x54, 0x65, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x51, 0x0a, 0x13, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x4c, 0x6f, 0x6f, 0x6b, 0x75,
	0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x42, 0x0a, 0x14, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x4c, 0x6f,
	0x6f, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x67, 0x6f, 0x61, 0x6b, 0x74, 0x70, 0x62, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x97, 0x01, 0x0a, 0x0d, 0x52, 0x65, 0x6d,
	0x6f, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x28, 0x0a, 0x06, 0x73, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67, 0x6f, 0x61,
	0x6b, 0x74, 0x70, 0x62, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x06, 0x73, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67, 0x6f, 0x61, 0x6b, 0x74, 0x70, 0x62,
	0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x72, 0x12, 0x2e, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x52, 0x0a, 0x14, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x53, 0x70,
	0x61, 0x77, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f,
	0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x17, 0x0a, 0x15, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65,
	0x52, 0x65, 0x53, 0x70, 0x61, 0x77, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x4f, 0x0a, 0x11, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x14, 0x0a, 0x12, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x7a, 0x0a, 0x12, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65,
	0x53, 0x70, 0x61, 0x77, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x70, 0x6f, 0x72, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x54, 0x79,
	0x70, 0x65, 0x22, 0x15, 0x0a, 0x13, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x53, 0x70, 0x61, 0x77,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xf4, 0x03, 0x0a, 0x0f, 0x52, 0x65,
	0x6d, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4c, 0x0a,
	0x09, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x41, 0x73, 0x6b, 0x12, 0x1c, 0x2e, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x41, 0x73,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x41, 0x73, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x12, 0x4d, 0x0a, 0x0a, 0x52,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x54, 0x65, 0x6c, 0x6c, 0x12, 0x1d, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x54, 0x65, 0x6c,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x54, 0x65, 0x6c, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x12, 0x51, 0x0a, 0x0c, 0x52, 0x65,
	0x6d, 0x6f, 0x74, 0x65, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x12, 0x1f, 0x2e, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x4c, 0x6f,
	0x6f, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x4c,
	0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x54, 0x0a,
	0x0d, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x53, 0x70, 0x61, 0x77, 0x6e, 0x12, 0x20,
	0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x6d, 0x6f,
	0x74, 0x65, 0x52, 0x65, 0x53, 0x70, 0x61, 0x77, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x21, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x70, 0x62, 0x2e, 0x52, 0x65,
	0x6d, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x53, 0x70, 0x61, 0x77, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x0a, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x53, 0x74, 0x6f,
	0x70, 0x12, 0x1d, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x70, 0x62, 0x2e, 0x52,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1e, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x70, 0x62, 0x2e, 0x52, 0x65,
	0x6d, 0x6f, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x4e, 0x0a, 0x0b, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x53, 0x70, 0x61, 0x77, 0x6e, 0x12,
	0x1e, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x6d,
	0x6f, 0x74, 0x65, 0x53, 0x70, 0x61, 0x77, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1f, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x6d,
	0x6f, 0x74, 0x65, 0x53, 0x70, 0x61, 0x77, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0xa6, 0x01, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x70, 0x62, 0x42, 0x0d, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x48, 0x02, 0x50, 0x01, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x74, 0x6f, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x79, 0x2f, 0x67, 0x6f, 0x61, 0x6b,
	0x74, 0x2f, 0x76, 0x32, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x70, 0x62, 0x3b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x70, 0x62, 0xa2, 0x02, 0x03, 0x49, 0x58, 0x58, 0xaa, 0x02, 0x0a, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x70, 0x62, 0xca, 0x02, 0x0a, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x70, 0x62, 0xe2, 0x02, 0x16, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x70, 0x62,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_internal_remoting_proto_rawDescOnce sync.Once
	file_internal_remoting_proto_rawDescData = file_internal_remoting_proto_rawDesc
)

func file_internal_remoting_proto_rawDescGZIP() []byte {
	file_internal_remoting_proto_rawDescOnce.Do(func() {
		file_internal_remoting_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_remoting_proto_rawDescData)
	})
	return file_internal_remoting_proto_rawDescData
}

var file_internal_remoting_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_internal_remoting_proto_goTypes = []interface{}{
	(*RemoteAskRequest)(nil),      // 0: internalpb.RemoteAskRequest
	(*RemoteAskResponse)(nil),     // 1: internalpb.RemoteAskResponse
	(*RemoteTellRequest)(nil),     // 2: internalpb.RemoteTellRequest
	(*RemoteTellResponse)(nil),    // 3: internalpb.RemoteTellResponse
	(*RemoteLookupRequest)(nil),   // 4: internalpb.RemoteLookupRequest
	(*RemoteLookupResponse)(nil),  // 5: internalpb.RemoteLookupResponse
	(*RemoteMessage)(nil),         // 6: internalpb.RemoteMessage
	(*RemoteReSpawnRequest)(nil),  // 7: internalpb.RemoteReSpawnRequest
	(*RemoteReSpawnResponse)(nil), // 8: internalpb.RemoteReSpawnResponse
	(*RemoteStopRequest)(nil),     // 9: internalpb.RemoteStopRequest
	(*RemoteStopResponse)(nil),    // 10: internalpb.RemoteStopResponse
	(*RemoteSpawnRequest)(nil),    // 11: internalpb.RemoteSpawnRequest
	(*RemoteSpawnResponse)(nil),   // 12: internalpb.RemoteSpawnResponse
	(*durationpb.Duration)(nil),   // 13: google.protobuf.Duration
	(*anypb.Any)(nil),             // 14: google.protobuf.Any
	(*goaktpb.Address)(nil),       // 15: goaktpb.Address
}
var file_internal_remoting_proto_depIdxs = []int32{
	6,  // 0: internalpb.RemoteAskRequest.remote_message:type_name -> internalpb.RemoteMessage
	13, // 1: internalpb.RemoteAskRequest.timeout:type_name -> google.protobuf.Duration
	14, // 2: internalpb.RemoteAskResponse.message:type_name -> google.protobuf.Any
	6,  // 3: internalpb.RemoteTellRequest.remote_message:type_name -> internalpb.RemoteMessage
	15, // 4: internalpb.RemoteLookupResponse.address:type_name -> goaktpb.Address
	15, // 5: internalpb.RemoteMessage.sender:type_name -> goaktpb.Address
	15, // 6: internalpb.RemoteMessage.receiver:type_name -> goaktpb.Address
	14, // 7: internalpb.RemoteMessage.message:type_name -> google.protobuf.Any
	0,  // 8: internalpb.RemotingService.RemoteAsk:input_type -> internalpb.RemoteAskRequest
	2,  // 9: internalpb.RemotingService.RemoteTell:input_type -> internalpb.RemoteTellRequest
	4,  // 10: internalpb.RemotingService.RemoteLookup:input_type -> internalpb.RemoteLookupRequest
	7,  // 11: internalpb.RemotingService.RemoteReSpawn:input_type -> internalpb.RemoteReSpawnRequest
	9,  // 12: internalpb.RemotingService.RemoteStop:input_type -> internalpb.RemoteStopRequest
	11, // 13: internalpb.RemotingService.RemoteSpawn:input_type -> internalpb.RemoteSpawnRequest
	1,  // 14: internalpb.RemotingService.RemoteAsk:output_type -> internalpb.RemoteAskResponse
	3,  // 15: internalpb.RemotingService.RemoteTell:output_type -> internalpb.RemoteTellResponse
	5,  // 16: internalpb.RemotingService.RemoteLookup:output_type -> internalpb.RemoteLookupResponse
	8,  // 17: internalpb.RemotingService.RemoteReSpawn:output_type -> internalpb.RemoteReSpawnResponse
	10, // 18: internalpb.RemotingService.RemoteStop:output_type -> internalpb.RemoteStopResponse
	12, // 19: internalpb.RemotingService.RemoteSpawn:output_type -> internalpb.RemoteSpawnResponse
	14, // [14:20] is the sub-list for method output_type
	8,  // [8:14] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_internal_remoting_proto_init() }
func file_internal_remoting_proto_init() {
	if File_internal_remoting_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_remoting_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_internal_remoting_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_internal_remoting_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_internal_remoting_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_internal_remoting_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_internal_remoting_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_internal_remoting_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoteMessage); i {
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
		file_internal_remoting_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoteReSpawnRequest); i {
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
		file_internal_remoting_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoteReSpawnResponse); i {
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
		file_internal_remoting_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoteStopRequest); i {
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
		file_internal_remoting_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoteStopResponse); i {
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
		file_internal_remoting_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoteSpawnRequest); i {
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
		file_internal_remoting_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoteSpawnResponse); i {
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
			RawDescriptor: file_internal_remoting_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_remoting_proto_goTypes,
		DependencyIndexes: file_internal_remoting_proto_depIdxs,
		MessageInfos:      file_internal_remoting_proto_msgTypes,
	}.Build()
	File_internal_remoting_proto = out.File
	file_internal_remoting_proto_rawDesc = nil
	file_internal_remoting_proto_goTypes = nil
	file_internal_remoting_proto_depIdxs = nil
}
