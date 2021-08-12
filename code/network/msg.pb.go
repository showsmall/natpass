// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: msg.proto

package network

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

type MsgType int32

const (
	Msg_handshake   MsgType = 0
	Msg_keepalive   MsgType = 1
	Msg_connect_req MsgType = 2
	Msg_connect_rep MsgType = 3
	Msg_disconnect  MsgType = 4
	Msg_forward     MsgType = 5
)

// Enum value maps for MsgType.
var (
	MsgType_name = map[int32]string{
		0: "handshake",
		1: "keepalive",
		2: "connect_req",
		3: "connect_rep",
		4: "disconnect",
		5: "forward",
	}
	MsgType_value = map[string]int32{
		"handshake":   0,
		"keepalive":   1,
		"connect_req": 2,
		"connect_rep": 3,
		"disconnect":  4,
		"forward":     5,
	}
)

func (x MsgType) Enum() *MsgType {
	p := new(MsgType)
	*p = x
	return p
}

func (x MsgType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MsgType) Descriptor() protoreflect.EnumDescriptor {
	return file_msg_proto_enumTypes[0].Descriptor()
}

func (MsgType) Type() protoreflect.EnumType {
	return &file_msg_proto_enumTypes[0]
}

func (x MsgType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MsgType.Descriptor instead.
func (MsgType) EnumDescriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{1, 0}
}

type HandshakePayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enc []byte `protobuf:"bytes,1,opt,name=enc,proto3" json:"enc,omitempty"`
}

func (x *HandshakePayload) Reset() {
	*x = HandshakePayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HandshakePayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HandshakePayload) ProtoMessage() {}

func (x *HandshakePayload) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HandshakePayload.ProtoReflect.Descriptor instead.
func (*HandshakePayload) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{0}
}

func (x *HandshakePayload) GetEnc() []byte {
	if x != nil {
		return x.Enc
	}
	return nil
}

type Msg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	XType MsgType `protobuf:"varint,1,opt,name=_type,json=Type,proto3,enum=network.MsgType" json:"_type,omitempty"`
	From  string  `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	To    string  `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	// Types that are assignable to Payload:
	//	*Msg_Hsp
	//	*Msg_Creq
	//	*Msg_Crep
	//	*Msg_XDisconnect
	//	*Msg_XData
	Payload isMsg_Payload `protobuf_oneof:"payload"`
}

func (x *Msg) Reset() {
	*x = Msg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Msg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Msg) ProtoMessage() {}

func (x *Msg) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Msg.ProtoReflect.Descriptor instead.
func (*Msg) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{1}
}

func (x *Msg) GetXType() MsgType {
	if x != nil {
		return x.XType
	}
	return Msg_handshake
}

func (x *Msg) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *Msg) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (m *Msg) GetPayload() isMsg_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *Msg) GetHsp() *HandshakePayload {
	if x, ok := x.GetPayload().(*Msg_Hsp); ok {
		return x.Hsp
	}
	return nil
}

func (x *Msg) GetCreq() *ConnectRequest {
	if x, ok := x.GetPayload().(*Msg_Creq); ok {
		return x.Creq
	}
	return nil
}

func (x *Msg) GetCrep() *ConnectResponse {
	if x, ok := x.GetPayload().(*Msg_Crep); ok {
		return x.Crep
	}
	return nil
}

func (x *Msg) GetXDisconnect() *Disconnect {
	if x, ok := x.GetPayload().(*Msg_XDisconnect); ok {
		return x.XDisconnect
	}
	return nil
}

func (x *Msg) GetXData() *Data {
	if x, ok := x.GetPayload().(*Msg_XData); ok {
		return x.XData
	}
	return nil
}

type isMsg_Payload interface {
	isMsg_Payload()
}

type Msg_Hsp struct {
	Hsp *HandshakePayload `protobuf:"bytes,10,opt,name=hsp,proto3,oneof"`
}

type Msg_Creq struct {
	Creq *ConnectRequest `protobuf:"bytes,11,opt,name=creq,proto3,oneof"`
}

type Msg_Crep struct {
	Crep *ConnectResponse `protobuf:"bytes,12,opt,name=crep,proto3,oneof"`
}

type Msg_XDisconnect struct {
	XDisconnect *Disconnect `protobuf:"bytes,13,opt,name=_disconnect,json=Disconnect,proto3,oneof"`
}

type Msg_XData struct {
	XData *Data `protobuf:"bytes,14,opt,name=_data,json=Data,proto3,oneof"`
}

func (*Msg_Hsp) isMsg_Payload() {}

func (*Msg_Creq) isMsg_Payload() {}

func (*Msg_Crep) isMsg_Payload() {}

func (*Msg_XDisconnect) isMsg_Payload() {}

func (*Msg_XData) isMsg_Payload() {}

var File_msg_proto protoreflect.FileDescriptor

var file_msg_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6d, 0x73, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x1a, 0x0d, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x25, 0x0a, 0x11, 0x68, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x5f,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x63, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x65, 0x6e, 0x63, 0x22, 0xb0, 0x03, 0x0a, 0x03, 0x6d, 0x73,
	0x67, 0x12, 0x26, 0x0a, 0x05, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x11, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f,
	0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a,
	0x02, 0x74, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x2e, 0x0a,
	0x03, 0x68, 0x73, 0x70, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x68, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x5f, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x48, 0x00, 0x52, 0x03, 0x68, 0x73, 0x70, 0x12, 0x2e, 0x0a,
	0x04, 0x63, 0x72, 0x65, 0x71, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x5f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x04, 0x63, 0x72, 0x65, 0x71, 0x12, 0x2f, 0x0a,
	0x04, 0x63, 0x72, 0x65, 0x70, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x5f, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x04, 0x63, 0x72, 0x65, 0x70, 0x12, 0x36,
	0x0a, 0x0b, 0x5f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x48, 0x00, 0x52, 0x0a, 0x44, 0x69, 0x73, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x24, 0x0a, 0x05, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x22, 0x63, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x0d, 0x0a, 0x09, 0x68, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b,
	0x65, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x6b, 0x65, 0x65, 0x70, 0x61, 0x6c, 0x69, 0x76, 0x65,
	0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x5f, 0x72, 0x65,
	0x71, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x5f, 0x72,
	0x65, 0x70, 0x10, 0x03, 0x12, 0x0e, 0x0a, 0x0a, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x10,
	0x05, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x0c, 0x5a, 0x0a,
	0x2e, 0x2f, 0x3b, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_msg_proto_rawDescOnce sync.Once
	file_msg_proto_rawDescData = file_msg_proto_rawDesc
)

func file_msg_proto_rawDescGZIP() []byte {
	file_msg_proto_rawDescOnce.Do(func() {
		file_msg_proto_rawDescData = protoimpl.X.CompressGZIP(file_msg_proto_rawDescData)
	})
	return file_msg_proto_rawDescData
}

var file_msg_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_msg_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_msg_proto_goTypes = []interface{}{
	(MsgType)(0),             // 0: network.msg.type
	(*HandshakePayload)(nil), // 1: network.handshake_payload
	(*Msg)(nil),              // 2: network.msg
	(*ConnectRequest)(nil),   // 3: network.connect_request
	(*ConnectResponse)(nil),  // 4: network.connect_response
	(*Disconnect)(nil),       // 5: network.disconnect
	(*Data)(nil),             // 6: network.data
}
var file_msg_proto_depIdxs = []int32{
	0, // 0: network.msg._type:type_name -> network.msg.type
	1, // 1: network.msg.hsp:type_name -> network.handshake_payload
	3, // 2: network.msg.creq:type_name -> network.connect_request
	4, // 3: network.msg.crep:type_name -> network.connect_response
	5, // 4: network.msg._disconnect:type_name -> network.disconnect
	6, // 5: network.msg._data:type_name -> network.data
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_msg_proto_init() }
func file_msg_proto_init() {
	if File_msg_proto != nil {
		return
	}
	file_connect_proto_init()
	file_forward_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_msg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HandshakePayload); i {
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
		file_msg_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Msg); i {
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
	file_msg_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Msg_Hsp)(nil),
		(*Msg_Creq)(nil),
		(*Msg_Crep)(nil),
		(*Msg_XDisconnect)(nil),
		(*Msg_XData)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_msg_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_msg_proto_goTypes,
		DependencyIndexes: file_msg_proto_depIdxs,
		EnumInfos:         file_msg_proto_enumTypes,
		MessageInfos:      file_msg_proto_msgTypes,
	}.Build()
	File_msg_proto = out.File
	file_msg_proto_rawDesc = nil
	file_msg_proto_goTypes = nil
	file_msg_proto_depIdxs = nil
}