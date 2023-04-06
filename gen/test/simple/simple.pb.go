// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: test/simple/simple.proto

package simple

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Enum1 int32

const (
	Enum1_Foo Enum1 = 0
	Enum1_Bar Enum1 = 1
)

// Enum value maps for Enum1.
var (
	Enum1_name = map[int32]string{
		0: "Foo",
		1: "Bar",
	}
	Enum1_value = map[string]int32{
		"Foo": 0,
		"Bar": 1,
	}
)

func (x Enum1) Enum() *Enum1 {
	p := new(Enum1)
	*p = x
	return p
}

func (x Enum1) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Enum1) Descriptor() protoreflect.EnumDescriptor {
	return file_test_simple_simple_proto_enumTypes[0].Descriptor()
}

func (Enum1) Type() protoreflect.EnumType {
	return &file_test_simple_simple_proto_enumTypes[0]
}

func (x Enum1) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Enum1.Descriptor instead.
func (Enum1) EnumDescriptor() ([]byte, []int) {
	return file_test_simple_simple_proto_rawDescGZIP(), []int{0}
}

type Message1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	String_ string          `protobuf:"bytes,1,opt,name=string,proto3" json:"string,omitempty"`
	Int32   int32           `protobuf:"varint,2,opt,name=int32,proto3" json:"int32,omitempty"`
	Bytes   []byte          `protobuf:"bytes,3,opt,name=bytes,proto3" json:"bytes,omitempty"`
	Enum    Enum1           `protobuf:"varint,4,opt,name=enum,proto3,enum=test.simple.Enum1" json:"enum,omitempty"`
	Message *Message2       `protobuf:"bytes,5,opt,name=message,proto3" json:"message,omitempty"`
	Any     *anypb.Any      `protobuf:"bytes,6,opt,name=any,proto3" json:"any,omitempty"`
	Value   *structpb.Value `protobuf:"bytes,7,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Message1) Reset() {
	*x = Message1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_simple_simple_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message1) ProtoMessage() {}

func (x *Message1) ProtoReflect() protoreflect.Message {
	mi := &file_test_simple_simple_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message1.ProtoReflect.Descriptor instead.
func (*Message1) Descriptor() ([]byte, []int) {
	return file_test_simple_simple_proto_rawDescGZIP(), []int{0}
}

func (x *Message1) GetString_() string {
	if x != nil {
		return x.String_
	}
	return ""
}

func (x *Message1) GetInt32() int32 {
	if x != nil {
		return x.Int32
	}
	return 0
}

func (x *Message1) GetBytes() []byte {
	if x != nil {
		return x.Bytes
	}
	return nil
}

func (x *Message1) GetEnum() Enum1 {
	if x != nil {
		return x.Enum
	}
	return Enum1_Foo
}

func (x *Message1) GetMessage() *Message2 {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *Message1) GetAny() *anypb.Any {
	if x != nil {
		return x.Any
	}
	return nil
}

func (x *Message1) GetValue() *structpb.Value {
	if x != nil {
		return x.Value
	}
	return nil
}

type Message2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	String_ string    `protobuf:"bytes,1,opt,name=string,proto3" json:"string,omitempty"`
	Int32   int32     `protobuf:"varint,2,opt,name=int32,proto3" json:"int32,omitempty"`
	Bytes   []byte    `protobuf:"bytes,3,opt,name=bytes,proto3" json:"bytes,omitempty"`
	Enum    Enum1     `protobuf:"varint,4,opt,name=enum,proto3,enum=test.simple.Enum1" json:"enum,omitempty"`
	Message *Message3 `protobuf:"bytes,5,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Message2) Reset() {
	*x = Message2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_simple_simple_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message2) ProtoMessage() {}

func (x *Message2) ProtoReflect() protoreflect.Message {
	mi := &file_test_simple_simple_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message2.ProtoReflect.Descriptor instead.
func (*Message2) Descriptor() ([]byte, []int) {
	return file_test_simple_simple_proto_rawDescGZIP(), []int{1}
}

func (x *Message2) GetString_() string {
	if x != nil {
		return x.String_
	}
	return ""
}

func (x *Message2) GetInt32() int32 {
	if x != nil {
		return x.Int32
	}
	return 0
}

func (x *Message2) GetBytes() []byte {
	if x != nil {
		return x.Bytes
	}
	return nil
}

func (x *Message2) GetEnum() Enum1 {
	if x != nil {
		return x.Enum
	}
	return Enum1_Foo
}

func (x *Message2) GetMessage() *Message3 {
	if x != nil {
		return x.Message
	}
	return nil
}

type Message3 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	String_ string `protobuf:"bytes,1,opt,name=string,proto3" json:"string,omitempty"`
	Int32   int32  `protobuf:"varint,2,opt,name=int32,proto3" json:"int32,omitempty"`
	Bytes   []byte `protobuf:"bytes,3,opt,name=bytes,proto3" json:"bytes,omitempty"`
	Enum    Enum1  `protobuf:"varint,4,opt,name=enum,proto3,enum=test.simple.Enum1" json:"enum,omitempty"`
}

func (x *Message3) Reset() {
	*x = Message3{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_simple_simple_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message3) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message3) ProtoMessage() {}

func (x *Message3) ProtoReflect() protoreflect.Message {
	mi := &file_test_simple_simple_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message3.ProtoReflect.Descriptor instead.
func (*Message3) Descriptor() ([]byte, []int) {
	return file_test_simple_simple_proto_rawDescGZIP(), []int{2}
}

func (x *Message3) GetString_() string {
	if x != nil {
		return x.String_
	}
	return ""
}

func (x *Message3) GetInt32() int32 {
	if x != nil {
		return x.Int32
	}
	return 0
}

func (x *Message3) GetBytes() []byte {
	if x != nil {
		return x.Bytes
	}
	return nil
}

func (x *Message3) GetEnum() Enum1 {
	if x != nil {
		return x.Enum
	}
	return Enum1_Foo
}

var File_test_simple_simple_proto protoreflect.FileDescriptor

var file_test_simple_simple_proto_rawDesc = []byte{
	0x0a, 0x18, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x73, 0x69,
	0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x74, 0x65, 0x73, 0x74,
	0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xfd, 0x01, 0x0a, 0x08, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x31, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x14, 0x0a, 0x05, 0x62,
	0x79, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x62, 0x79, 0x74, 0x65,
	0x73, 0x12, 0x26, 0x0a, 0x04, 0x65, 0x6e, 0x75, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x12, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x45, 0x6e,
	0x75, 0x6d, 0x31, 0x52, 0x04, 0x65, 0x6e, 0x75, 0x6d, 0x12, 0x2f, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x32, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x26, 0x0a, 0x03, 0x61, 0x6e,
	0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x03, 0x61,
	0x6e, 0x79, 0x12, 0x2c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0xa7, 0x01, 0x0a, 0x08, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x14, 0x0a, 0x05, 0x62,
	0x79, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x62, 0x79, 0x74, 0x65,
	0x73, 0x12, 0x26, 0x0a, 0x04, 0x65, 0x6e, 0x75, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x12, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x45, 0x6e,
	0x75, 0x6d, 0x31, 0x52, 0x04, 0x65, 0x6e, 0x75, 0x6d, 0x12, 0x2f, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x33, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x76, 0x0a, 0x08, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x33, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x14,
	0x0a, 0x05, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69,
	0x6e, 0x74, 0x33, 0x32, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x04, 0x65, 0x6e,
	0x75, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e,
	0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x31, 0x52, 0x04, 0x65, 0x6e,
	0x75, 0x6d, 0x2a, 0x19, 0x0a, 0x05, 0x45, 0x6e, 0x75, 0x6d, 0x31, 0x12, 0x07, 0x0a, 0x03, 0x46,
	0x6f, 0x6f, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x42, 0x61, 0x72, 0x10, 0x01, 0x42, 0x9e, 0x01,
	0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c,
	0x65, 0x42, 0x0b, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x6f, 0x73,
	0x68, 0x63, 0x61, 0x72, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x74, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x73, 0x69, 0x6d,
	0x70, 0x6c, 0x65, 0xa2, 0x02, 0x03, 0x54, 0x53, 0x58, 0xaa, 0x02, 0x0b, 0x54, 0x65, 0x73, 0x74,
	0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0xca, 0x02, 0x0b, 0x54, 0x65, 0x73, 0x74, 0x5c, 0x53,
	0x69, 0x6d, 0x70, 0x6c, 0x65, 0xe2, 0x02, 0x17, 0x54, 0x65, 0x73, 0x74, 0x5c, 0x53, 0x69, 0x6d,
	0x70, 0x6c, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x0c, 0x54, 0x65, 0x73, 0x74, 0x3a, 0x3a, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_test_simple_simple_proto_rawDescOnce sync.Once
	file_test_simple_simple_proto_rawDescData = file_test_simple_simple_proto_rawDesc
)

func file_test_simple_simple_proto_rawDescGZIP() []byte {
	file_test_simple_simple_proto_rawDescOnce.Do(func() {
		file_test_simple_simple_proto_rawDescData = protoimpl.X.CompressGZIP(file_test_simple_simple_proto_rawDescData)
	})
	return file_test_simple_simple_proto_rawDescData
}

var file_test_simple_simple_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_test_simple_simple_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_test_simple_simple_proto_goTypes = []interface{}{
	(Enum1)(0),             // 0: test.simple.Enum1
	(*Message1)(nil),       // 1: test.simple.Message1
	(*Message2)(nil),       // 2: test.simple.Message2
	(*Message3)(nil),       // 3: test.simple.Message3
	(*anypb.Any)(nil),      // 4: google.protobuf.Any
	(*structpb.Value)(nil), // 5: google.protobuf.Value
}
var file_test_simple_simple_proto_depIdxs = []int32{
	0, // 0: test.simple.Message1.enum:type_name -> test.simple.Enum1
	2, // 1: test.simple.Message1.message:type_name -> test.simple.Message2
	4, // 2: test.simple.Message1.any:type_name -> google.protobuf.Any
	5, // 3: test.simple.Message1.value:type_name -> google.protobuf.Value
	0, // 4: test.simple.Message2.enum:type_name -> test.simple.Enum1
	3, // 5: test.simple.Message2.message:type_name -> test.simple.Message3
	0, // 6: test.simple.Message3.enum:type_name -> test.simple.Enum1
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_test_simple_simple_proto_init() }
func file_test_simple_simple_proto_init() {
	if File_test_simple_simple_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_test_simple_simple_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message1); i {
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
		file_test_simple_simple_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message2); i {
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
		file_test_simple_simple_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message3); i {
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
			RawDescriptor: file_test_simple_simple_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_test_simple_simple_proto_goTypes,
		DependencyIndexes: file_test_simple_simple_proto_depIdxs,
		EnumInfos:         file_test_simple_simple_proto_enumTypes,
		MessageInfos:      file_test_simple_simple_proto_msgTypes,
	}.Build()
	File_test_simple_simple_proto = out.File
	file_test_simple_simple_proto_rawDesc = nil
	file_test_simple_simple_proto_goTypes = nil
	file_test_simple_simple_proto_depIdxs = nil
}