// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.9
// source: statistic/config.proto

package statistic

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

type NetType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConnType       string `protobuf:"bytes,1,opt,name=conn_type,proto3" json:"conn_type,omitempty"`
	UnderlyingType string `protobuf:"bytes,2,opt,name=underlying_type,proto3" json:"underlying_type,omitempty"`
}

func (x *NetType) Reset() {
	*x = NetType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_statistic_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetType) ProtoMessage() {}

func (x *NetType) ProtoReflect() protoreflect.Message {
	mi := &file_statistic_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetType.ProtoReflect.Descriptor instead.
func (*NetType) Descriptor() ([]byte, []int) {
	return file_statistic_config_proto_rawDescGZIP(), []int{0}
}

func (x *NetType) GetConnType() string {
	if x != nil {
		return x.ConnType
	}
	return ""
}

func (x *NetType) GetUnderlyingType() string {
	if x != nil {
		return x.UnderlyingType
	}
	return ""
}

type Connection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addr   string            `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	Id     uint64            `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Remote string            `protobuf:"bytes,4,opt,name=remote,proto3" json:"remote,omitempty"`
	Type   *NetType          `protobuf:"bytes,6,opt,name=type,proto3" json:"type,omitempty"`
	Extra  map[string]string `protobuf:"bytes,9,rep,name=extra,proto3" json:"extra,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Connection) Reset() {
	*x = Connection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_statistic_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Connection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Connection) ProtoMessage() {}

func (x *Connection) ProtoReflect() protoreflect.Message {
	mi := &file_statistic_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Connection.ProtoReflect.Descriptor instead.
func (*Connection) Descriptor() ([]byte, []int) {
	return file_statistic_config_proto_rawDescGZIP(), []int{1}
}

func (x *Connection) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Connection) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Connection) GetRemote() string {
	if x != nil {
		return x.Remote
	}
	return ""
}

func (x *Connection) GetType() *NetType {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *Connection) GetExtra() map[string]string {
	if x != nil {
		return x.Extra
	}
	return nil
}

var File_statistic_config_proto protoreflect.FileDescriptor

var file_statistic_config_proto_rawDesc = []byte{
	0x0a, 0x16, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x79, 0x75, 0x68, 0x61, 0x69, 0x69,
	0x6e, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x22, 0x52, 0x0a, 0x08, 0x6e,
	0x65, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x6e, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x6e,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x75, 0x6e, 0x64, 0x65, 0x72, 0x6c, 0x79,
	0x69, 0x6e, 0x67, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x75, 0x6e, 0x64, 0x65, 0x72, 0x6c, 0x79, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22,
	0xf3, 0x01, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64,
	0x64, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x79, 0x75, 0x68, 0x61, 0x69,
	0x69, 0x6e, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x6e, 0x65, 0x74,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x3e, 0x0a, 0x05, 0x65,
	0x78, 0x74, 0x72, 0x61, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x79, 0x75, 0x68,
	0x61, 0x69, 0x69, 0x6e, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x1a, 0x38, 0x0a, 0x0a, 0x45,
	0x78, 0x74, 0x72, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x73, 0x75, 0x74, 0x6f, 0x72, 0x75, 0x66, 0x61, 0x2f, 0x79, 0x75,
	0x68, 0x61, 0x69, 0x69, 0x6e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_statistic_config_proto_rawDescOnce sync.Once
	file_statistic_config_proto_rawDescData = file_statistic_config_proto_rawDesc
)

func file_statistic_config_proto_rawDescGZIP() []byte {
	file_statistic_config_proto_rawDescOnce.Do(func() {
		file_statistic_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_statistic_config_proto_rawDescData)
	})
	return file_statistic_config_proto_rawDescData
}

var file_statistic_config_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_statistic_config_proto_goTypes = []interface{}{
	(*NetType)(nil),    // 0: yuhaiin.statistic.net_type
	(*Connection)(nil), // 1: yuhaiin.statistic.connection
	nil,                // 2: yuhaiin.statistic.connection.ExtraEntry
}
var file_statistic_config_proto_depIdxs = []int32{
	0, // 0: yuhaiin.statistic.connection.type:type_name -> yuhaiin.statistic.net_type
	2, // 1: yuhaiin.statistic.connection.extra:type_name -> yuhaiin.statistic.connection.ExtraEntry
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_statistic_config_proto_init() }
func file_statistic_config_proto_init() {
	if File_statistic_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_statistic_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetType); i {
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
		file_statistic_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Connection); i {
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
			RawDescriptor: file_statistic_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_statistic_config_proto_goTypes,
		DependencyIndexes: file_statistic_config_proto_depIdxs,
		MessageInfos:      file_statistic_config_proto_msgTypes,
	}.Build()
	File_statistic_config_proto = out.File
	file_statistic_config_proto_rawDesc = nil
	file_statistic_config_proto_goTypes = nil
	file_statistic_config_proto_depIdxs = nil
}
