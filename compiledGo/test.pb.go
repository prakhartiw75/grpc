// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.4
// source: test.proto

package compiledGo

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

type SumRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	First  int32 `protobuf:"varint,1,opt,name=first,proto3" json:"first,omitempty"`
	Second int32 `protobuf:"varint,2,opt,name=second,proto3" json:"second,omitempty"`
}

func (x *SumRequest) Reset() {
	*x = SumRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SumRequest) ProtoMessage() {}

func (x *SumRequest) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SumRequest.ProtoReflect.Descriptor instead.
func (*SumRequest) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{0}
}

func (x *SumRequest) GetFirst() int32 {
	if x != nil {
		return x.First
	}
	return 0
}

func (x *SumRequest) GetSecond() int32 {
	if x != nil {
		return x.Second
	}
	return 0
}

type SumResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalSum int32 `protobuf:"varint,1,opt,name=totalSum,proto3" json:"totalSum,omitempty"`
}

func (x *SumResponse) Reset() {
	*x = SumResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SumResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SumResponse) ProtoMessage() {}

func (x *SumResponse) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SumResponse.ProtoReflect.Descriptor instead.
func (*SumResponse) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{1}
}

func (x *SumResponse) GetTotalSum() int32 {
	if x != nil {
		return x.TotalSum
	}
	return 0
}

type PrimeNumberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val int32 `protobuf:"varint,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *PrimeNumberRequest) Reset() {
	*x = PrimeNumberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrimeNumberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrimeNumberRequest) ProtoMessage() {}

func (x *PrimeNumberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrimeNumberRequest.ProtoReflect.Descriptor instead.
func (*PrimeNumberRequest) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{2}
}

func (x *PrimeNumberRequest) GetVal() int32 {
	if x != nil {
		return x.Val
	}
	return 0
}

type PrimeNumberResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Prime int32 `protobuf:"varint,1,opt,name=prime,proto3" json:"prime,omitempty"`
}

func (x *PrimeNumberResponse) Reset() {
	*x = PrimeNumberResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrimeNumberResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrimeNumberResponse) ProtoMessage() {}

func (x *PrimeNumberResponse) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrimeNumberResponse.ProtoReflect.Descriptor instead.
func (*PrimeNumberResponse) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{3}
}

func (x *PrimeNumberResponse) GetPrime() int32 {
	if x != nil {
		return x.Prime
	}
	return 0
}

type ComputeAverageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val int32 `protobuf:"varint,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *ComputeAverageRequest) Reset() {
	*x = ComputeAverageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComputeAverageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComputeAverageRequest) ProtoMessage() {}

func (x *ComputeAverageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComputeAverageRequest.ProtoReflect.Descriptor instead.
func (*ComputeAverageRequest) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{4}
}

func (x *ComputeAverageRequest) GetVal() int32 {
	if x != nil {
		return x.Val
	}
	return 0
}

type ComputeAverageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Average float32 `protobuf:"fixed32,1,opt,name=average,proto3" json:"average,omitempty"`
}

func (x *ComputeAverageResponse) Reset() {
	*x = ComputeAverageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComputeAverageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComputeAverageResponse) ProtoMessage() {}

func (x *ComputeAverageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComputeAverageResponse.ProtoReflect.Descriptor instead.
func (*ComputeAverageResponse) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{5}
}

func (x *ComputeAverageResponse) GetAverage() float32 {
	if x != nil {
		return x.Average
	}
	return 0
}

type FindMaxRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int32 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *FindMaxRequest) Reset() {
	*x = FindMaxRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindMaxRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindMaxRequest) ProtoMessage() {}

func (x *FindMaxRequest) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindMaxRequest.ProtoReflect.Descriptor instead.
func (*FindMaxRequest) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{6}
}

func (x *FindMaxRequest) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

type FindMaxResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Maximum int32 `protobuf:"varint,1,opt,name=maximum,proto3" json:"maximum,omitempty"`
}

func (x *FindMaxResponse) Reset() {
	*x = FindMaxResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindMaxResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindMaxResponse) ProtoMessage() {}

func (x *FindMaxResponse) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindMaxResponse.ProtoReflect.Descriptor instead.
func (*FindMaxResponse) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{7}
}

func (x *FindMaxResponse) GetMaximum() int32 {
	if x != nil {
		return x.Maximum
	}
	return 0
}

var File_test_proto protoreflect.FileDescriptor

var file_test_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3a, 0x0a, 0x0a,
	0x53, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x22, 0x29, 0x0a, 0x0b, 0x53, 0x75, 0x6d, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x53, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x53, 0x75, 0x6d, 0x22, 0x26, 0x0a, 0x12, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x61, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x22, 0x2b, 0x0a, 0x13, 0x50,
	0x72, 0x69, 0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x22, 0x29, 0x0a, 0x15, 0x43, 0x6f, 0x6d, 0x70,
	0x75, 0x74, 0x65, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x76, 0x61, 0x6c, 0x22, 0x32, 0x0a, 0x16, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x41, 0x76,
	0x65, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07,
	0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x22, 0x28, 0x0a, 0x0e, 0x46, 0x69, 0x6e, 0x64, 0x4d,
	0x61, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x22, 0x2b, 0x0a, 0x0f, 0x46, 0x69, 0x6e, 0x64, 0x4d, 0x61, 0x78, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x32, 0xeb,
	0x01, 0x0a, 0x0e, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x20, 0x0a, 0x03, 0x53, 0x75, 0x6d, 0x12, 0x0b, 0x2e, 0x53, 0x75, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x0b, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x12, 0x13, 0x2e, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12,
	0x43, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67,
	0x65, 0x12, 0x16, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x41, 0x76, 0x65, 0x72, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x43, 0x6f, 0x6d, 0x70,
	0x75, 0x74, 0x65, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x28, 0x01, 0x12, 0x36, 0x0a, 0x0d, 0x46, 0x69, 0x6e, 0x64, 0x4d, 0x61, 0x78, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x0f, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4d, 0x61, 0x78, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4d, 0x61, 0x78,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42, 0x0d, 0x5a, 0x0b,
	0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x64, 0x47, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_test_proto_rawDescOnce sync.Once
	file_test_proto_rawDescData = file_test_proto_rawDesc
)

func file_test_proto_rawDescGZIP() []byte {
	file_test_proto_rawDescOnce.Do(func() {
		file_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_test_proto_rawDescData)
	})
	return file_test_proto_rawDescData
}

var file_test_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_test_proto_goTypes = []interface{}{
	(*SumRequest)(nil),             // 0: SumRequest
	(*SumResponse)(nil),            // 1: SumResponse
	(*PrimeNumberRequest)(nil),     // 2: PrimeNumberRequest
	(*PrimeNumberResponse)(nil),    // 3: PrimeNumberResponse
	(*ComputeAverageRequest)(nil),  // 4: ComputeAverageRequest
	(*ComputeAverageResponse)(nil), // 5: ComputeAverageResponse
	(*FindMaxRequest)(nil),         // 6: FindMaxRequest
	(*FindMaxResponse)(nil),        // 7: FindMaxResponse
}
var file_test_proto_depIdxs = []int32{
	0, // 0: RandomeRequest.Sum:input_type -> SumRequest
	2, // 1: RandomeRequest.PrimeNumber:input_type -> PrimeNumberRequest
	4, // 2: RandomeRequest.ComputeAverage:input_type -> ComputeAverageRequest
	6, // 3: RandomeRequest.FindMaxNumber:input_type -> FindMaxRequest
	1, // 4: RandomeRequest.Sum:output_type -> SumResponse
	3, // 5: RandomeRequest.PrimeNumber:output_type -> PrimeNumberResponse
	5, // 6: RandomeRequest.ComputeAverage:output_type -> ComputeAverageResponse
	7, // 7: RandomeRequest.FindMaxNumber:output_type -> FindMaxResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_test_proto_init() }
func file_test_proto_init() {
	if File_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SumRequest); i {
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
		file_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SumResponse); i {
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
		file_test_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrimeNumberRequest); i {
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
		file_test_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrimeNumberResponse); i {
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
		file_test_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComputeAverageRequest); i {
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
		file_test_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComputeAverageResponse); i {
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
		file_test_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindMaxRequest); i {
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
		file_test_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindMaxResponse); i {
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
			RawDescriptor: file_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_test_proto_goTypes,
		DependencyIndexes: file_test_proto_depIdxs,
		MessageInfos:      file_test_proto_msgTypes,
	}.Build()
	File_test_proto = out.File
	file_test_proto_rawDesc = nil
	file_test_proto_goTypes = nil
	file_test_proto_depIdxs = nil
}
