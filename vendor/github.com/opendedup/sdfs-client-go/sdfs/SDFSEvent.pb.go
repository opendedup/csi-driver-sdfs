// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: SDFSEvent.proto

package sdfs

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type SDFSEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartTime       int64             `protobuf:"varint,1,opt,name=startTime,proto3" json:"startTime,omitempty"`
	EndTime         int64             `protobuf:"varint,2,opt,name=endTime,proto3" json:"endTime,omitempty"`
	Level           string            `protobuf:"bytes,3,opt,name=level,proto3" json:"level,omitempty"`
	Type            string            `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Target          string            `protobuf:"bytes,5,opt,name=target,proto3" json:"target,omitempty"`
	ShortMsg        string            `protobuf:"bytes,6,opt,name=shortMsg,proto3" json:"shortMsg,omitempty"`
	LongMsg         string            `protobuf:"bytes,7,opt,name=longMsg,proto3" json:"longMsg,omitempty"`
	PercentComplete float64           `protobuf:"fixed64,8,opt,name=percentComplete,proto3" json:"percentComplete,omitempty"`
	MaxCount        int64             `protobuf:"varint,9,opt,name=maxCount,proto3" json:"maxCount,omitempty"`
	CurrentCount    int64             `protobuf:"varint,10,opt,name=currentCount,proto3" json:"currentCount,omitempty"`
	Uuid            string            `protobuf:"bytes,11,opt,name=uuid,proto3" json:"uuid,omitempty"`
	ParentUuid      string            `protobuf:"bytes,12,opt,name=parentUuid,proto3" json:"parentUuid,omitempty"`
	ExtendedInfo    string            `protobuf:"bytes,13,opt,name=extendedInfo,proto3" json:"extendedInfo,omitempty"`
	ChildrenUUid    []string          `protobuf:"bytes,14,rep,name=childrenUUid,proto3" json:"childrenUUid,omitempty"`
	Success         bool              `protobuf:"varint,15,opt,name=success,proto3" json:"success,omitempty"`
	Attributes      map[string]string `protobuf:"bytes,16,rep,name=attributes,proto3" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *SDFSEvent) Reset() {
	*x = SDFSEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SDFSEvent_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SDFSEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SDFSEvent) ProtoMessage() {}

func (x *SDFSEvent) ProtoReflect() protoreflect.Message {
	mi := &file_SDFSEvent_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SDFSEvent.ProtoReflect.Descriptor instead.
func (*SDFSEvent) Descriptor() ([]byte, []int) {
	return file_SDFSEvent_proto_rawDescGZIP(), []int{0}
}

func (x *SDFSEvent) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *SDFSEvent) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *SDFSEvent) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

func (x *SDFSEvent) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *SDFSEvent) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *SDFSEvent) GetShortMsg() string {
	if x != nil {
		return x.ShortMsg
	}
	return ""
}

func (x *SDFSEvent) GetLongMsg() string {
	if x != nil {
		return x.LongMsg
	}
	return ""
}

func (x *SDFSEvent) GetPercentComplete() float64 {
	if x != nil {
		return x.PercentComplete
	}
	return 0
}

func (x *SDFSEvent) GetMaxCount() int64 {
	if x != nil {
		return x.MaxCount
	}
	return 0
}

func (x *SDFSEvent) GetCurrentCount() int64 {
	if x != nil {
		return x.CurrentCount
	}
	return 0
}

func (x *SDFSEvent) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *SDFSEvent) GetParentUuid() string {
	if x != nil {
		return x.ParentUuid
	}
	return ""
}

func (x *SDFSEvent) GetExtendedInfo() string {
	if x != nil {
		return x.ExtendedInfo
	}
	return ""
}

func (x *SDFSEvent) GetChildrenUUid() []string {
	if x != nil {
		return x.ChildrenUUid
	}
	return nil
}

func (x *SDFSEvent) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *SDFSEvent) GetAttributes() map[string]string {
	if x != nil {
		return x.Attributes
	}
	return nil
}

type SDFSEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *SDFSEventRequest) Reset() {
	*x = SDFSEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SDFSEvent_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SDFSEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SDFSEventRequest) ProtoMessage() {}

func (x *SDFSEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_SDFSEvent_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SDFSEventRequest.ProtoReflect.Descriptor instead.
func (*SDFSEventRequest) Descriptor() ([]byte, []int) {
	return file_SDFSEvent_proto_rawDescGZIP(), []int{1}
}

func (x *SDFSEventRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type SDFSEventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error     string     `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	ErrorCode ErrorCodes `protobuf:"varint,2,opt,name=errorCode,proto3,enum=org.opendedup.grpc.ErrorCodes" json:"errorCode,omitempty"`
	Event     *SDFSEvent `protobuf:"bytes,3,opt,name=event,proto3" json:"event,omitempty"`
}

func (x *SDFSEventResponse) Reset() {
	*x = SDFSEventResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SDFSEvent_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SDFSEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SDFSEventResponse) ProtoMessage() {}

func (x *SDFSEventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_SDFSEvent_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SDFSEventResponse.ProtoReflect.Descriptor instead.
func (*SDFSEventResponse) Descriptor() ([]byte, []int) {
	return file_SDFSEvent_proto_rawDescGZIP(), []int{2}
}

func (x *SDFSEventResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *SDFSEventResponse) GetErrorCode() ErrorCodes {
	if x != nil {
		return x.ErrorCode
	}
	return ErrorCodes_NOERR
}

func (x *SDFSEventResponse) GetEvent() *SDFSEvent {
	if x != nil {
		return x.Event
	}
	return nil
}

type SDFSEventListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SDFSEventListRequest) Reset() {
	*x = SDFSEventListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SDFSEvent_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SDFSEventListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SDFSEventListRequest) ProtoMessage() {}

func (x *SDFSEventListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_SDFSEvent_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SDFSEventListRequest.ProtoReflect.Descriptor instead.
func (*SDFSEventListRequest) Descriptor() ([]byte, []int) {
	return file_SDFSEvent_proto_rawDescGZIP(), []int{3}
}

type SDFSEventListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error     string       `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	ErrorCode ErrorCodes   `protobuf:"varint,2,opt,name=errorCode,proto3,enum=org.opendedup.grpc.ErrorCodes" json:"errorCode,omitempty"`
	Events    []*SDFSEvent `protobuf:"bytes,3,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *SDFSEventListResponse) Reset() {
	*x = SDFSEventListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SDFSEvent_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SDFSEventListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SDFSEventListResponse) ProtoMessage() {}

func (x *SDFSEventListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_SDFSEvent_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SDFSEventListResponse.ProtoReflect.Descriptor instead.
func (*SDFSEventListResponse) Descriptor() ([]byte, []int) {
	return file_SDFSEvent_proto_rawDescGZIP(), []int{4}
}

func (x *SDFSEventListResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *SDFSEventListResponse) GetErrorCode() ErrorCodes {
	if x != nil {
		return x.ErrorCode
	}
	return ErrorCodes_NOERR
}

func (x *SDFSEventListResponse) GetEvents() []*SDFSEvent {
	if x != nil {
		return x.Events
	}
	return nil
}

var File_SDFSEvent_proto protoreflect.FileDescriptor

var file_SDFSEvent_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x53, 0x44, 0x46, 0x53, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x12, 0x6f, 0x72, 0x67, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x75, 0x70,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x1a, 0x0e, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc9, 0x04, 0x0a, 0x09, 0x53, 0x44, 0x46, 0x53, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x4d, 0x73, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x4d, 0x73, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x6f, 0x6e,
	0x67, 0x4d, 0x73, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x6f, 0x6e, 0x67,
	0x4d, 0x73, 0x67, 0x12, 0x28, 0x0a, 0x0f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x43, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0f, 0x70, 0x65,
	0x72, 0x63, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x6d, 0x61, 0x78, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x6d, 0x61, 0x78, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69,
	0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x55, 0x75, 0x69, 0x64, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x55, 0x75, 0x69,
	0x64, 0x12, 0x22, 0x0a, 0x0c, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x49, 0x6e, 0x66,
	0x6f, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65,
	0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65,
	0x6e, 0x55, 0x55, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x68, 0x69,
	0x6c, 0x64, 0x72, 0x65, 0x6e, 0x55, 0x55, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x12, 0x4d, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x73, 0x18, 0x10, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6f, 0x70,
	0x65, 0x6e, 0x64, 0x65, 0x64, 0x75, 0x70, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x44, 0x46,
	0x53, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x73, 0x1a, 0x3d, 0x0a, 0x0f, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0x26, 0x0a, 0x10, 0x53, 0x44, 0x46, 0x53, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0x9c, 0x01, 0x0a, 0x11, 0x53, 0x44,
	0x46, 0x53, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x3c, 0x0a, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f,
	0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6f,
	0x70, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x75, 0x70, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x33, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x64,
	0x75, 0x70, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x44, 0x46, 0x53, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x16, 0x0a, 0x14, 0x53, 0x44, 0x46, 0x53,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0xa2, 0x01, 0x0a, 0x15, 0x53, 0x44, 0x46, 0x53, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x12, 0x3c, 0x0a, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x64, 0x65,
	0x64, 0x75, 0x70, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f,
	0x64, 0x65, 0x73, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x35,
	0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x75, 0x70, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x53, 0x44, 0x46, 0x53, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x32, 0xaf, 0x02, 0x0a, 0x10, 0x53, 0x44, 0x46, 0x53, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x57, 0x0a, 0x08, 0x47, 0x65,
	0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x24, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6f, 0x70, 0x65,
	0x6e, 0x64, 0x65, 0x64, 0x75, 0x70, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x44, 0x46, 0x53,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x6f,
	0x72, 0x67, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x75, 0x70, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x53, 0x44, 0x46, 0x53, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x61, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x12, 0x28, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x75,
	0x70, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x44, 0x46, 0x53, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x6f, 0x72,
	0x67, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x75, 0x70, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x53, 0x44, 0x46, 0x53, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5f, 0x0a, 0x0e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x24, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6f,
	0x70, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x75, 0x70, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x44,
	0x46, 0x53, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25,
	0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x75, 0x70, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x53, 0x44, 0x46, 0x53, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x75, 0x70, 0x2f,
	0x73, 0x64, 0x66, 0x73, 0x2d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2d, 0x67, 0x6f, 0x2f, 0x73,
	0x64, 0x66, 0x73, 0x2f, 0x3b, 0x73, 0x64, 0x66, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_SDFSEvent_proto_rawDescOnce sync.Once
	file_SDFSEvent_proto_rawDescData = file_SDFSEvent_proto_rawDesc
)

func file_SDFSEvent_proto_rawDescGZIP() []byte {
	file_SDFSEvent_proto_rawDescOnce.Do(func() {
		file_SDFSEvent_proto_rawDescData = protoimpl.X.CompressGZIP(file_SDFSEvent_proto_rawDescData)
	})
	return file_SDFSEvent_proto_rawDescData
}

var file_SDFSEvent_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_SDFSEvent_proto_goTypes = []interface{}{
	(*SDFSEvent)(nil),             // 0: org.opendedup.grpc.SDFSEvent
	(*SDFSEventRequest)(nil),      // 1: org.opendedup.grpc.SDFSEventRequest
	(*SDFSEventResponse)(nil),     // 2: org.opendedup.grpc.SDFSEventResponse
	(*SDFSEventListRequest)(nil),  // 3: org.opendedup.grpc.SDFSEventListRequest
	(*SDFSEventListResponse)(nil), // 4: org.opendedup.grpc.SDFSEventListResponse
	nil,                           // 5: org.opendedup.grpc.SDFSEvent.AttributesEntry
	(ErrorCodes)(0),               // 6: org.opendedup.grpc.errorCodes
}
var file_SDFSEvent_proto_depIdxs = []int32{
	5, // 0: org.opendedup.grpc.SDFSEvent.attributes:type_name -> org.opendedup.grpc.SDFSEvent.AttributesEntry
	6, // 1: org.opendedup.grpc.SDFSEventResponse.errorCode:type_name -> org.opendedup.grpc.errorCodes
	0, // 2: org.opendedup.grpc.SDFSEventResponse.event:type_name -> org.opendedup.grpc.SDFSEvent
	6, // 3: org.opendedup.grpc.SDFSEventListResponse.errorCode:type_name -> org.opendedup.grpc.errorCodes
	0, // 4: org.opendedup.grpc.SDFSEventListResponse.events:type_name -> org.opendedup.grpc.SDFSEvent
	1, // 5: org.opendedup.grpc.SDFSEventService.GetEvent:input_type -> org.opendedup.grpc.SDFSEventRequest
	3, // 6: org.opendedup.grpc.SDFSEventService.ListEvents:input_type -> org.opendedup.grpc.SDFSEventListRequest
	1, // 7: org.opendedup.grpc.SDFSEventService.SubscribeEvent:input_type -> org.opendedup.grpc.SDFSEventRequest
	2, // 8: org.opendedup.grpc.SDFSEventService.GetEvent:output_type -> org.opendedup.grpc.SDFSEventResponse
	4, // 9: org.opendedup.grpc.SDFSEventService.ListEvents:output_type -> org.opendedup.grpc.SDFSEventListResponse
	2, // 10: org.opendedup.grpc.SDFSEventService.SubscribeEvent:output_type -> org.opendedup.grpc.SDFSEventResponse
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_SDFSEvent_proto_init() }
func file_SDFSEvent_proto_init() {
	if File_SDFSEvent_proto != nil {
		return
	}
	file_FileInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SDFSEvent_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SDFSEvent); i {
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
		file_SDFSEvent_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SDFSEventRequest); i {
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
		file_SDFSEvent_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SDFSEventResponse); i {
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
		file_SDFSEvent_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SDFSEventListRequest); i {
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
		file_SDFSEvent_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SDFSEventListResponse); i {
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
			RawDescriptor: file_SDFSEvent_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_SDFSEvent_proto_goTypes,
		DependencyIndexes: file_SDFSEvent_proto_depIdxs,
		MessageInfos:      file_SDFSEvent_proto_msgTypes,
	}.Build()
	File_SDFSEvent_proto = out.File
	file_SDFSEvent_proto_rawDesc = nil
	file_SDFSEvent_proto_goTypes = nil
	file_SDFSEvent_proto_depIdxs = nil
}
