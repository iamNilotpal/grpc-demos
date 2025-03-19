// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.3
// source: proto/stream.proto

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

type LogLevel int32

const (
	LogLevel_UNSPECIFIED LogLevel = 0
	LogLevel_DEBUG       LogLevel = 1
	LogLevel_INFO        LogLevel = 2
	LogLevel_WARNING     LogLevel = 3
	LogLevel_ERROR       LogLevel = 4
)

// Enum value maps for LogLevel.
var (
	LogLevel_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "DEBUG",
		2: "INFO",
		3: "WARNING",
		4: "ERROR",
	}
	LogLevel_value = map[string]int32{
		"UNSPECIFIED": 0,
		"DEBUG":       1,
		"INFO":        2,
		"WARNING":     3,
		"ERROR":       4,
	}
)

func (x LogLevel) Enum() *LogLevel {
	p := new(LogLevel)
	*p = x
	return p
}

func (x LogLevel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LogLevel) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_stream_proto_enumTypes[0].Descriptor()
}

func (LogLevel) Type() protoreflect.EnumType {
	return &file_proto_stream_proto_enumTypes[0]
}

func (x LogLevel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LogLevel.Descriptor instead.
func (LogLevel) EnumDescriptor() ([]byte, []int) {
	return file_proto_stream_proto_rawDescGZIP(), []int{0}
}

type StreamTimeRequest struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	IntervalSeconds int32                  `protobuf:"zigzag32,1,opt,name=intervalSeconds,proto3" json:"intervalSeconds,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *StreamTimeRequest) Reset() {
	*x = StreamTimeRequest{}
	mi := &file_proto_stream_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StreamTimeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamTimeRequest) ProtoMessage() {}

func (x *StreamTimeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_stream_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamTimeRequest.ProtoReflect.Descriptor instead.
func (*StreamTimeRequest) Descriptor() ([]byte, []int) {
	return file_proto_stream_proto_rawDescGZIP(), []int{0}
}

func (x *StreamTimeRequest) GetIntervalSeconds() int32 {
	if x != nil {
		return x.IntervalSeconds
	}
	return 0
}

type StreamTimeResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CurrentTime   *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=currentTime,proto3" json:"currentTime,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StreamTimeResponse) Reset() {
	*x = StreamTimeResponse{}
	mi := &file_proto_stream_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StreamTimeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamTimeResponse) ProtoMessage() {}

func (x *StreamTimeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_stream_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamTimeResponse.ProtoReflect.Descriptor instead.
func (*StreamTimeResponse) Descriptor() ([]byte, []int) {
	return file_proto_stream_proto_rawDescGZIP(), []int{1}
}

func (x *StreamTimeResponse) GetCurrentTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CurrentTime
	}
	return nil
}

type LogStreamRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Timestamp     *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	LogLevel      LogLevel               `protobuf:"varint,2,opt,name=logLevel,proto3,enum=stream.LogLevel" json:"logLevel,omitempty"`
	Message       string                 `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LogStreamRequest) Reset() {
	*x = LogStreamRequest{}
	mi := &file_proto_stream_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogStreamRequest) ProtoMessage() {}

func (x *LogStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_stream_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogStreamRequest.ProtoReflect.Descriptor instead.
func (*LogStreamRequest) Descriptor() ([]byte, []int) {
	return file_proto_stream_proto_rawDescGZIP(), []int{2}
}

func (x *LogStreamRequest) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *LogStreamRequest) GetLogLevel() LogLevel {
	if x != nil {
		return x.LogLevel
	}
	return LogLevel_UNSPECIFIED
}

func (x *LogStreamRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type LogStreamResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	EntiresLogged int32                  `protobuf:"varint,1,opt,name=entiresLogged,proto3" json:"entiresLogged,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LogStreamResponse) Reset() {
	*x = LogStreamResponse{}
	mi := &file_proto_stream_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogStreamResponse) ProtoMessage() {}

func (x *LogStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_stream_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogStreamResponse.ProtoReflect.Descriptor instead.
func (*LogStreamResponse) Descriptor() ([]byte, []int) {
	return file_proto_stream_proto_rawDescGZIP(), []int{3}
}

func (x *LogStreamResponse) GetEntiresLogged() int32 {
	if x != nil {
		return x.EntiresLogged
	}
	return 0
}

var File_proto_stream_proto protoreflect.FileDescriptor

var file_proto_stream_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3d, 0x0a,
	0x11, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x28, 0x0a, 0x0f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x53, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x11, 0x52, 0x0f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x76, 0x61, 0x6c, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x22, 0x52, 0x0a, 0x12,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x22, 0x94, 0x01, 0x0a, 0x10, 0x4c, 0x6f, 0x67, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12,
	0x2c, 0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x10, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x4c, 0x6f, 0x67, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x52, 0x08, 0x6c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x39, 0x0a, 0x11, 0x4c, 0x6f, 0x67, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x0d,
	0x65, 0x6e, 0x74, 0x69, 0x72, 0x65, 0x73, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0d, 0x65, 0x6e, 0x74, 0x69, 0x72, 0x65, 0x73, 0x4c, 0x6f, 0x67, 0x67,
	0x65, 0x64, 0x2a, 0x48, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x0f,
	0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x09, 0x0a, 0x05, 0x44, 0x45, 0x42, 0x55, 0x47, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x4e,
	0x46, 0x4f, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x57, 0x41, 0x52, 0x4e, 0x49, 0x4e, 0x47, 0x10,
	0x03, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x04, 0x32, 0xaa, 0x01, 0x0a,
	0x11, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x4b, 0x0a, 0x10, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x19, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1a, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12,
	0x48, 0x0a, 0x0f, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4c,
	0x6f, 0x67, 0x12, 0x18, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x4c, 0x6f, 0x67, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x4c, 0x6f, 0x67, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x61, 0x6d, 0x4e, 0x69, 0x6c, 0x6f, 0x74,
	0x70, 0x61, 0x6c, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_stream_proto_rawDescOnce sync.Once
	file_proto_stream_proto_rawDescData = file_proto_stream_proto_rawDesc
)

func file_proto_stream_proto_rawDescGZIP() []byte {
	file_proto_stream_proto_rawDescOnce.Do(func() {
		file_proto_stream_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_stream_proto_rawDescData)
	})
	return file_proto_stream_proto_rawDescData
}

var file_proto_stream_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_stream_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_stream_proto_goTypes = []any{
	(LogLevel)(0),                 // 0: stream.LogLevel
	(*StreamTimeRequest)(nil),     // 1: stream.StreamTimeRequest
	(*StreamTimeResponse)(nil),    // 2: stream.StreamTimeResponse
	(*LogStreamRequest)(nil),      // 3: stream.LogStreamRequest
	(*LogStreamResponse)(nil),     // 4: stream.LogStreamResponse
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_proto_stream_proto_depIdxs = []int32{
	5, // 0: stream.StreamTimeResponse.currentTime:type_name -> google.protobuf.Timestamp
	5, // 1: stream.LogStreamRequest.timestamp:type_name -> google.protobuf.Timestamp
	0, // 2: stream.LogStreamRequest.logLevel:type_name -> stream.LogLevel
	1, // 3: stream.StreamTimeService.StreamServerTime:input_type -> stream.StreamTimeRequest
	3, // 4: stream.StreamTimeService.StreamServerLog:input_type -> stream.LogStreamRequest
	2, // 5: stream.StreamTimeService.StreamServerTime:output_type -> stream.StreamTimeResponse
	4, // 6: stream.StreamTimeService.StreamServerLog:output_type -> stream.LogStreamResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_stream_proto_init() }
func file_proto_stream_proto_init() {
	if File_proto_stream_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_stream_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_stream_proto_goTypes,
		DependencyIndexes: file_proto_stream_proto_depIdxs,
		EnumInfos:         file_proto_stream_proto_enumTypes,
		MessageInfos:      file_proto_stream_proto_msgTypes,
	}.Build()
	File_proto_stream_proto = out.File
	file_proto_stream_proto_rawDesc = nil
	file_proto_stream_proto_goTypes = nil
	file_proto_stream_proto_depIdxs = nil
}
