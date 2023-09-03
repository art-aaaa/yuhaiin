// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.24.2
// source: config/dns/dns.proto

package dns

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

type Type int32

const (
	Type_reserve Type = 0
	Type_udp     Type = 1
	Type_tcp     Type = 2
	Type_doh     Type = 3
	Type_dot     Type = 4
	Type_doq     Type = 5
	Type_doh3    Type = 6
)

// Enum value maps for Type.
var (
	Type_name = map[int32]string{
		0: "reserve",
		1: "udp",
		2: "tcp",
		3: "doh",
		4: "dot",
		5: "doq",
		6: "doh3",
	}
	Type_value = map[string]int32{
		"reserve": 0,
		"udp":     1,
		"tcp":     2,
		"doh":     3,
		"dot":     4,
		"doq":     5,
		"doh3":    6,
	}
)

func (x Type) Enum() *Type {
	p := new(Type)
	*p = x
	return p
}

func (x Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Type) Descriptor() protoreflect.EnumDescriptor {
	return file_config_dns_dns_proto_enumTypes[0].Descriptor()
}

func (Type) Type() protoreflect.EnumType {
	return &file_config_dns_dns_proto_enumTypes[0]
}

func (x Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Type.Descriptor instead.
func (Type) EnumDescriptor() ([]byte, []int) {
	return file_config_dns_dns_proto_rawDescGZIP(), []int{0}
}

type Dns struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host          string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Type          Type   `protobuf:"varint,5,opt,name=type,proto3,enum=yuhaiin.dns.Type" json:"type,omitempty"`
	Subnet        string `protobuf:"bytes,4,opt,name=subnet,proto3" json:"subnet,omitempty"`
	TlsServername string `protobuf:"bytes,2,opt,name=tls_servername,proto3" json:"tls_servername,omitempty"`
}

func (x *Dns) Reset() {
	*x = Dns{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_dns_dns_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dns) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dns) ProtoMessage() {}

func (x *Dns) ProtoReflect() protoreflect.Message {
	mi := &file_config_dns_dns_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dns.ProtoReflect.Descriptor instead.
func (*Dns) Descriptor() ([]byte, []int) {
	return file_config_dns_dns_proto_rawDescGZIP(), []int{0}
}

func (x *Dns) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *Dns) GetType() Type {
	if x != nil {
		return x.Type
	}
	return Type_reserve
}

func (x *Dns) GetSubnet() string {
	if x != nil {
		return x.Subnet
	}
	return ""
}

func (x *Dns) GetTlsServername() string {
	if x != nil {
		return x.TlsServername
	}
	return ""
}

type DnsConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Server              string            `protobuf:"bytes,4,opt,name=server,proto3" json:"server,omitempty"`
	Fakedns             bool              `protobuf:"varint,5,opt,name=fakedns,proto3" json:"fakedns,omitempty"`
	FakednsIpRange      string            `protobuf:"bytes,6,opt,name=fakedns_ip_range,proto3" json:"fakedns_ip_range,omitempty"`
	ResolveRemoteDomain bool              `protobuf:"varint,7,opt,name=resolve_remote_domain,proto3" json:"resolve_remote_domain,omitempty"`
	Remote              *Dns              `protobuf:"bytes,1,opt,name=remote,proto3" json:"remote,omitempty"`
	Local               *Dns              `protobuf:"bytes,2,opt,name=local,proto3" json:"local,omitempty"`
	Bootstrap           *Dns              `protobuf:"bytes,3,opt,name=bootstrap,proto3" json:"bootstrap,omitempty"`
	Hosts               map[string]string `protobuf:"bytes,8,rep,name=hosts,proto3" json:"hosts,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *DnsConfig) Reset() {
	*x = DnsConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_dns_dns_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DnsConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DnsConfig) ProtoMessage() {}

func (x *DnsConfig) ProtoReflect() protoreflect.Message {
	mi := &file_config_dns_dns_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DnsConfig.ProtoReflect.Descriptor instead.
func (*DnsConfig) Descriptor() ([]byte, []int) {
	return file_config_dns_dns_proto_rawDescGZIP(), []int{1}
}

func (x *DnsConfig) GetServer() string {
	if x != nil {
		return x.Server
	}
	return ""
}

func (x *DnsConfig) GetFakedns() bool {
	if x != nil {
		return x.Fakedns
	}
	return false
}

func (x *DnsConfig) GetFakednsIpRange() string {
	if x != nil {
		return x.FakednsIpRange
	}
	return ""
}

func (x *DnsConfig) GetResolveRemoteDomain() bool {
	if x != nil {
		return x.ResolveRemoteDomain
	}
	return false
}

func (x *DnsConfig) GetRemote() *Dns {
	if x != nil {
		return x.Remote
	}
	return nil
}

func (x *DnsConfig) GetLocal() *Dns {
	if x != nil {
		return x.Local
	}
	return nil
}

func (x *DnsConfig) GetBootstrap() *Dns {
	if x != nil {
		return x.Bootstrap
	}
	return nil
}

func (x *DnsConfig) GetHosts() map[string]string {
	if x != nil {
		return x.Hosts
	}
	return nil
}

var File_config_dns_dns_proto protoreflect.FileDescriptor

var file_config_dns_dns_proto_rawDesc = []byte{
	0x0a, 0x14, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x64, 0x6e, 0x73, 0x2f, 0x64, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x79, 0x75, 0x68, 0x61, 0x69, 0x69, 0x6e, 0x2e,
	0x64, 0x6e, 0x73, 0x22, 0x80, 0x01, 0x0a, 0x03, 0x64, 0x6e, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x68,
	0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12,
	0x25, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e,
	0x79, 0x75, 0x68, 0x61, 0x69, 0x69, 0x6e, 0x2e, 0x64, 0x6e, 0x73, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x12, 0x26,
	0x0a, 0x0e, 0x74, 0x6c, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x74, 0x6c, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x96, 0x03, 0x0a, 0x0a, 0x64, 0x6e, 0x73, 0x5f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x18, 0x0a,
	0x07, 0x66, 0x61, 0x6b, 0x65, 0x64, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x66, 0x61, 0x6b, 0x65, 0x64, 0x6e, 0x73, 0x12, 0x2a, 0x0a, 0x10, 0x66, 0x61, 0x6b, 0x65, 0x64,
	0x6e, 0x73, 0x5f, 0x69, 0x70, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x10, 0x66, 0x61, 0x6b, 0x65, 0x64, 0x6e, 0x73, 0x5f, 0x69, 0x70, 0x5f, 0x72, 0x61,
	0x6e, 0x67, 0x65, 0x12, 0x34, 0x0a, 0x15, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x5f, 0x72,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x15, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x5f, 0x72, 0x65, 0x6d, 0x6f,
	0x74, 0x65, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x28, 0x0a, 0x06, 0x72, 0x65, 0x6d,
	0x6f, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x79, 0x75, 0x68, 0x61,
	0x69, 0x69, 0x6e, 0x2e, 0x64, 0x6e, 0x73, 0x2e, 0x64, 0x6e, 0x73, 0x52, 0x06, 0x72, 0x65, 0x6d,
	0x6f, 0x74, 0x65, 0x12, 0x26, 0x0a, 0x05, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x79, 0x75, 0x68, 0x61, 0x69, 0x69, 0x6e, 0x2e, 0x64, 0x6e, 0x73,
	0x2e, 0x64, 0x6e, 0x73, 0x52, 0x05, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x12, 0x2e, 0x0a, 0x09, 0x62,
	0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x79, 0x75, 0x68, 0x61, 0x69, 0x69, 0x6e, 0x2e, 0x64, 0x6e, 0x73, 0x2e, 0x64, 0x6e, 0x73,
	0x52, 0x09, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x12, 0x38, 0x0a, 0x05, 0x68,
	0x6f, 0x73, 0x74, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x79, 0x75, 0x68,
	0x61, 0x69, 0x69, 0x6e, 0x2e, 0x64, 0x6e, 0x73, 0x2e, 0x64, 0x6e, 0x73, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05,
	0x68, 0x6f, 0x73, 0x74, 0x73, 0x1a, 0x38, 0x0a, 0x0a, 0x48, 0x6f, 0x73, 0x74, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x2a,
	0x4a, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x75, 0x64, 0x70, 0x10, 0x01, 0x12, 0x07, 0x0a,
	0x03, 0x74, 0x63, 0x70, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x64, 0x6f, 0x68, 0x10, 0x03, 0x12,
	0x07, 0x0a, 0x03, 0x64, 0x6f, 0x74, 0x10, 0x04, 0x12, 0x07, 0x0a, 0x03, 0x64, 0x6f, 0x71, 0x10,
	0x05, 0x12, 0x08, 0x0a, 0x04, 0x64, 0x6f, 0x68, 0x33, 0x10, 0x06, 0x42, 0x34, 0x5a, 0x32, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x73, 0x75, 0x74, 0x6f, 0x72,
	0x75, 0x66, 0x61, 0x2f, 0x79, 0x75, 0x68, 0x61, 0x69, 0x69, 0x6e, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x64, 0x6e,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_config_dns_dns_proto_rawDescOnce sync.Once
	file_config_dns_dns_proto_rawDescData = file_config_dns_dns_proto_rawDesc
)

func file_config_dns_dns_proto_rawDescGZIP() []byte {
	file_config_dns_dns_proto_rawDescOnce.Do(func() {
		file_config_dns_dns_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_dns_dns_proto_rawDescData)
	})
	return file_config_dns_dns_proto_rawDescData
}

var file_config_dns_dns_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_config_dns_dns_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_config_dns_dns_proto_goTypes = []interface{}{
	(Type)(0),         // 0: yuhaiin.dns.type
	(*Dns)(nil),       // 1: yuhaiin.dns.dns
	(*DnsConfig)(nil), // 2: yuhaiin.dns.dns_config
	nil,               // 3: yuhaiin.dns.dns_config.HostsEntry
}
var file_config_dns_dns_proto_depIdxs = []int32{
	0, // 0: yuhaiin.dns.dns.type:type_name -> yuhaiin.dns.type
	1, // 1: yuhaiin.dns.dns_config.remote:type_name -> yuhaiin.dns.dns
	1, // 2: yuhaiin.dns.dns_config.local:type_name -> yuhaiin.dns.dns
	1, // 3: yuhaiin.dns.dns_config.bootstrap:type_name -> yuhaiin.dns.dns
	3, // 4: yuhaiin.dns.dns_config.hosts:type_name -> yuhaiin.dns.dns_config.HostsEntry
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_config_dns_dns_proto_init() }
func file_config_dns_dns_proto_init() {
	if File_config_dns_dns_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_config_dns_dns_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Dns); i {
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
		file_config_dns_dns_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DnsConfig); i {
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
			RawDescriptor: file_config_dns_dns_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_config_dns_dns_proto_goTypes,
		DependencyIndexes: file_config_dns_dns_proto_depIdxs,
		EnumInfos:         file_config_dns_dns_proto_enumTypes,
		MessageInfos:      file_config_dns_dns_proto_msgTypes,
	}.Build()
	File_config_dns_dns_proto = out.File
	file_config_dns_dns_proto_rawDesc = nil
	file_config_dns_dns_proto_goTypes = nil
	file_config_dns_dns_proto_depIdxs = nil
}
