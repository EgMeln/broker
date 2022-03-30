// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: protocol/positionService.proto

package protocol

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

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID         string  `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	PriceOpen  float32 `protobuf:"fixed32,2,opt,name=PriceOpen,proto3" json:"PriceOpen,omitempty"`
	IsBay      bool    `protobuf:"varint,3,opt,name=IsBay,proto3" json:"IsBay,omitempty"`
	Symbol     string  `protobuf:"bytes,4,opt,name=Symbol,proto3" json:"Symbol,omitempty"`
	PriceClose float32 `protobuf:"fixed32,5,opt,name=PriceClose,proto3" json:"PriceClose,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_positionService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_positionService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_protocol_positionService_proto_rawDescGZIP(), []int{0}
}

func (x *Transaction) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Transaction) GetPriceOpen() float32 {
	if x != nil {
		return x.PriceOpen
	}
	return 0
}

func (x *Transaction) GetIsBay() bool {
	if x != nil {
		return x.IsBay
	}
	return false
}

func (x *Transaction) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *Transaction) GetPriceClose() float32 {
	if x != nil {
		return x.PriceClose
	}
	return 0
}

type OpenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Trans *Transaction `protobuf:"bytes,1,opt,name=trans,proto3" json:"trans,omitempty"`
}

func (x *OpenRequest) Reset() {
	*x = OpenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_positionService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenRequest) ProtoMessage() {}

func (x *OpenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_positionService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenRequest.ProtoReflect.Descriptor instead.
func (*OpenRequest) Descriptor() ([]byte, []int) {
	return file_protocol_positionService_proto_rawDescGZIP(), []int{1}
}

func (x *OpenRequest) GetTrans() *Transaction {
	if x != nil {
		return x.Trans
	}
	return nil
}

type OpenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *OpenResponse) Reset() {
	*x = OpenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_positionService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenResponse) ProtoMessage() {}

func (x *OpenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_positionService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenResponse.ProtoReflect.Descriptor instead.
func (*OpenResponse) Descriptor() ([]byte, []int) {
	return file_protocol_positionService_proto_rawDescGZIP(), []int{2}
}

func (x *OpenResponse) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type CloseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID         string  `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Symbol     string  `protobuf:"bytes,3,opt,name=Symbol,proto3" json:"Symbol,omitempty"`
	PriceClose float32 `protobuf:"fixed32,2,opt,name=PriceClose,proto3" json:"PriceClose,omitempty"`
}

func (x *CloseRequest) Reset() {
	*x = CloseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_positionService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloseRequest) ProtoMessage() {}

func (x *CloseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_positionService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloseRequest.ProtoReflect.Descriptor instead.
func (*CloseRequest) Descriptor() ([]byte, []int) {
	return file_protocol_positionService_proto_rawDescGZIP(), []int{3}
}

func (x *CloseRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *CloseRequest) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *CloseRequest) GetPriceClose() float32 {
	if x != nil {
		return x.PriceClose
	}
	return 0
}

type CloseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error string `protobuf:"bytes,1,opt,name=Error,proto3" json:"Error,omitempty"`
}

func (x *CloseResponse) Reset() {
	*x = CloseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_positionService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloseResponse) ProtoMessage() {}

func (x *CloseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_positionService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloseResponse.ProtoReflect.Descriptor instead.
func (*CloseResponse) Descriptor() ([]byte, []int) {
	return file_protocol_positionService_proto_rawDescGZIP(), []int{4}
}

func (x *CloseResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_protocol_positionService_proto protoreflect.FileDescriptor

var file_protocol_positionService_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x22, 0x89, 0x01, 0x0a, 0x0b, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x4f, 0x70, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x73, 0x42, 0x61,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x49, 0x73, 0x42, 0x61, 0x79, 0x12, 0x16,
	0x0a, 0x06, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x72, 0x69, 0x63, 0x65, 0x43,
	0x6c, 0x6f, 0x73, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x22, 0x3a, 0x0a, 0x0b, 0x4f, 0x70, 0x65, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x05, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x22, 0x1e, 0x0a, 0x0c, 0x4f, 0x70, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x49, 0x44, 0x22, 0x56, 0x0a, 0x0c, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x22, 0x25, 0x0a, 0x0d, 0x43, 0x6c,
	0x6f, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x32, 0xa7, 0x02, 0x0a, 0x0f, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x0f, 0x4f, 0x70, 0x65, 0x6e, 0x50, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x73, 0x6b, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x0f, 0x4f, 0x70, 0x65,
	0x6e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x69, 0x64, 0x12, 0x15, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4f,
	0x70, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a,
	0x10, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x73,
	0x6b, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x43, 0x6c, 0x6f,
	0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x10, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x50, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x69, 0x64, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x43, 0x6c, 0x6f, 0x73,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_protocol_positionService_proto_rawDescOnce sync.Once
	file_protocol_positionService_proto_rawDescData = file_protocol_positionService_proto_rawDesc
)

func file_protocol_positionService_proto_rawDescGZIP() []byte {
	file_protocol_positionService_proto_rawDescOnce.Do(func() {
		file_protocol_positionService_proto_rawDescData = protoimpl.X.CompressGZIP(file_protocol_positionService_proto_rawDescData)
	})
	return file_protocol_positionService_proto_rawDescData
}

var file_protocol_positionService_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_protocol_positionService_proto_goTypes = []interface{}{
	(*Transaction)(nil),   // 0: protobuf.Transaction
	(*OpenRequest)(nil),   // 1: protobuf.OpenRequest
	(*OpenResponse)(nil),  // 2: protobuf.OpenResponse
	(*CloseRequest)(nil),  // 3: protobuf.CloseRequest
	(*CloseResponse)(nil), // 4: protobuf.CloseResponse
}
var file_protocol_positionService_proto_depIdxs = []int32{
	0, // 0: protobuf.OpenRequest.trans:type_name -> protobuf.Transaction
	1, // 1: protobuf.PositionService.OpenPositionAsk:input_type -> protobuf.OpenRequest
	1, // 2: protobuf.PositionService.OpenPositionBid:input_type -> protobuf.OpenRequest
	3, // 3: protobuf.PositionService.ClosePositionAsk:input_type -> protobuf.CloseRequest
	3, // 4: protobuf.PositionService.ClosePositionBid:input_type -> protobuf.CloseRequest
	2, // 5: protobuf.PositionService.OpenPositionAsk:output_type -> protobuf.OpenResponse
	2, // 6: protobuf.PositionService.OpenPositionBid:output_type -> protobuf.OpenResponse
	4, // 7: protobuf.PositionService.ClosePositionAsk:output_type -> protobuf.CloseResponse
	4, // 8: protobuf.PositionService.ClosePositionBid:output_type -> protobuf.CloseResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protocol_positionService_proto_init() }
func file_protocol_positionService_proto_init() {
	if File_protocol_positionService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protocol_positionService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
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
		file_protocol_positionService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenRequest); i {
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
		file_protocol_positionService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenResponse); i {
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
		file_protocol_positionService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloseRequest); i {
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
		file_protocol_positionService_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloseResponse); i {
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
			RawDescriptor: file_protocol_positionService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protocol_positionService_proto_goTypes,
		DependencyIndexes: file_protocol_positionService_proto_depIdxs,
		MessageInfos:      file_protocol_positionService_proto_msgTypes,
	}.Build()
	File_protocol_positionService_proto = out.File
	file_protocol_positionService_proto_rawDesc = nil
	file_protocol_positionService_proto_goTypes = nil
	file_protocol_positionService_proto_depIdxs = nil
}
