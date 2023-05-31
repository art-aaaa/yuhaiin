// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: config/bypass/bypass.proto

package bypass

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

type Mode int32

const (
	Mode_bypass Mode = 0
	Mode_direct Mode = 1
	Mode_proxy  Mode = 2
	Mode_block  Mode = 3
)

// Enum value maps for Mode.
var (
	Mode_name = map[int32]string{
		0: "bypass",
		1: "direct",
		2: "proxy",
		3: "block",
	}
	Mode_value = map[string]int32{
		"bypass": 0,
		"direct": 1,
		"proxy":  2,
		"block":  3,
	}
)

func (x Mode) Enum() *Mode {
	p := new(Mode)
	*p = x
	return p
}

func (x Mode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Mode) Descriptor() protoreflect.EnumDescriptor {
	return file_config_bypass_bypass_proto_enumTypes[0].Descriptor()
}

func (Mode) Type() protoreflect.EnumType {
	return &file_config_bypass_bypass_proto_enumTypes[0]
}

func (x Mode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Mode.Descriptor instead.
func (Mode) EnumDescriptor() ([]byte, []int) {
	return file_config_bypass_bypass_proto_rawDescGZIP(), []int{0}
}

type ResolveStrategy int32

const (
	ResolveStrategy_default     ResolveStrategy = 0
	ResolveStrategy_prefer_ipv4 ResolveStrategy = 1
	ResolveStrategy_only_ipv4   ResolveStrategy = 2
	ResolveStrategy_prefer_ipv6 ResolveStrategy = 3
	ResolveStrategy_only_ipv6   ResolveStrategy = 4
)

// Enum value maps for ResolveStrategy.
var (
	ResolveStrategy_name = map[int32]string{
		0: "default",
		1: "prefer_ipv4",
		2: "only_ipv4",
		3: "prefer_ipv6",
		4: "only_ipv6",
	}
	ResolveStrategy_value = map[string]int32{
		"default":     0,
		"prefer_ipv4": 1,
		"only_ipv4":   2,
		"prefer_ipv6": 3,
		"only_ipv6":   4,
	}
)

func (x ResolveStrategy) Enum() *ResolveStrategy {
	p := new(ResolveStrategy)
	*p = x
	return p
}

func (x ResolveStrategy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResolveStrategy) Descriptor() protoreflect.EnumDescriptor {
	return file_config_bypass_bypass_proto_enumTypes[1].Descriptor()
}

func (ResolveStrategy) Type() protoreflect.EnumType {
	return &file_config_bypass_bypass_proto_enumTypes[1]
}

func (x ResolveStrategy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResolveStrategy.Descriptor instead.
func (ResolveStrategy) EnumDescriptor() ([]byte, []int) {
	return file_config_bypass_bypass_proto_rawDescGZIP(), []int{1}
}

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tcp          Mode          `protobuf:"varint,3,opt,name=tcp,proto3,enum=yuhaiin.bypass.Mode" json:"tcp,omitempty"`
	Udp          Mode          `protobuf:"varint,4,opt,name=udp,proto3,enum=yuhaiin.bypass.Mode" json:"udp,omitempty"`
	BypassFile   string        `protobuf:"bytes,2,opt,name=bypass_file,proto3" json:"bypass_file,omitempty"`
	CustomRuleV3 []*ModeConfig `protobuf:"bytes,7,rep,name=custom_rule_v3,proto3" json:"custom_rule_v3,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_bypass_bypass_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_config_bypass_bypass_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_config_bypass_bypass_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetTcp() Mode {
	if x != nil {
		return x.Tcp
	}
	return Mode_bypass
}

func (x *Config) GetUdp() Mode {
	if x != nil {
		return x.Udp
	}
	return Mode_bypass
}

func (x *Config) GetBypassFile() string {
	if x != nil {
		return x.BypassFile
	}
	return ""
}

func (x *Config) GetCustomRuleV3() []*ModeConfig {
	if x != nil {
		return x.CustomRuleV3
	}
	return nil
}

type ModeConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hostname        []string        `protobuf:"bytes,3,rep,name=hostname,proto3" json:"hostname,omitempty"`
	Mode            Mode            `protobuf:"varint,1,opt,name=mode,proto3,enum=yuhaiin.bypass.Mode" json:"mode,omitempty"`
	Tag             string          `protobuf:"bytes,2,opt,name=tag,proto3" json:"tag,omitempty"`
	ResolveStrategy ResolveStrategy `protobuf:"varint,4,opt,name=resolve_strategy,proto3,enum=yuhaiin.bypass.ResolveStrategy" json:"resolve_strategy,omitempty"`
}

func (x *ModeConfig) Reset() {
	*x = ModeConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_bypass_bypass_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModeConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModeConfig) ProtoMessage() {}

func (x *ModeConfig) ProtoReflect() protoreflect.Message {
	mi := &file_config_bypass_bypass_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModeConfig.ProtoReflect.Descriptor instead.
func (*ModeConfig) Descriptor() ([]byte, []int) {
	return file_config_bypass_bypass_proto_rawDescGZIP(), []int{1}
}

func (x *ModeConfig) GetHostname() []string {
	if x != nil {
		return x.Hostname
	}
	return nil
}

func (x *ModeConfig) GetMode() Mode {
	if x != nil {
		return x.Mode
	}
	return Mode_bypass
}

func (x *ModeConfig) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *ModeConfig) GetResolveStrategy() ResolveStrategy {
	if x != nil {
		return x.ResolveStrategy
	}
	return ResolveStrategy_default
}

var File_config_bypass_bypass_proto protoreflect.FileDescriptor

var file_config_bypass_bypass_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x62, 0x79, 0x70, 0x61, 0x73, 0x73, 0x2f,
	0x62, 0x79, 0x70, 0x61, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x79, 0x75,
	0x68, 0x61, 0x69, 0x69, 0x6e, 0x2e, 0x62, 0x79, 0x70, 0x61, 0x73, 0x73, 0x22, 0xbf, 0x01, 0x0a,
	0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x26, 0x0a, 0x03, 0x74, 0x63, 0x70, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x79, 0x75, 0x68, 0x61, 0x69, 0x69, 0x6e, 0x2e, 0x62,
	0x79, 0x70, 0x61, 0x73, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x52, 0x03, 0x74, 0x63, 0x70, 0x12,
	0x26, 0x0a, 0x03, 0x75, 0x64, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x79,
	0x75, 0x68, 0x61, 0x69, 0x69, 0x6e, 0x2e, 0x62, 0x79, 0x70, 0x61, 0x73, 0x73, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x52, 0x03, 0x75, 0x64, 0x70, 0x12, 0x20, 0x0a, 0x0b, 0x62, 0x79, 0x70, 0x61, 0x73,
	0x73, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x79,
	0x70, 0x61, 0x73, 0x73, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x43, 0x0a, 0x0e, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x76, 0x33, 0x18, 0x07, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x79, 0x75, 0x68, 0x61, 0x69, 0x69, 0x6e, 0x2e, 0x62, 0x79, 0x70, 0x61,
	0x73, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0e,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x76, 0x33, 0x22, 0xb3,
	0x01, 0x0a, 0x0b, 0x6d, 0x6f, 0x64, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1a,
	0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x6d, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x79, 0x75, 0x68, 0x61, 0x69,
	0x69, 0x6e, 0x2e, 0x62, 0x79, 0x70, 0x61, 0x73, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x52, 0x04,
	0x6d, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x12, 0x4c, 0x0a, 0x10, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76,
	0x65, 0x5f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x20, 0x2e, 0x79, 0x75, 0x68, 0x61, 0x69, 0x69, 0x6e, 0x2e, 0x62, 0x79, 0x70, 0x61, 0x73,
	0x73, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x5f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65,
	0x67, 0x79, 0x52, 0x10, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x5f, 0x73, 0x74, 0x72, 0x61,
	0x74, 0x65, 0x67, 0x79, 0x2a, 0x34, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x0a, 0x0a, 0x06,
	0x62, 0x79, 0x70, 0x61, 0x73, 0x73, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x64, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x10, 0x02, 0x12,
	0x09, 0x0a, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x10, 0x03, 0x2a, 0x5f, 0x0a, 0x10, 0x72, 0x65,
	0x73, 0x6f, 0x6c, 0x76, 0x65, 0x5f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x12, 0x0b,
	0x0a, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x70,
	0x72, 0x65, 0x66, 0x65, 0x72, 0x5f, 0x69, 0x70, 0x76, 0x34, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09,
	0x6f, 0x6e, 0x6c, 0x79, 0x5f, 0x69, 0x70, 0x76, 0x34, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x70,
	0x72, 0x65, 0x66, 0x65, 0x72, 0x5f, 0x69, 0x70, 0x76, 0x36, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09,
	0x6f, 0x6e, 0x6c, 0x79, 0x5f, 0x69, 0x70, 0x76, 0x36, 0x10, 0x04, 0x42, 0x37, 0x5a, 0x35, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x73, 0x75, 0x74, 0x6f, 0x72,
	0x75, 0x66, 0x61, 0x2f, 0x79, 0x75, 0x68, 0x61, 0x69, 0x69, 0x6e, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x62, 0x79,
	0x70, 0x61, 0x73, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_config_bypass_bypass_proto_rawDescOnce sync.Once
	file_config_bypass_bypass_proto_rawDescData = file_config_bypass_bypass_proto_rawDesc
)

func file_config_bypass_bypass_proto_rawDescGZIP() []byte {
	file_config_bypass_bypass_proto_rawDescOnce.Do(func() {
		file_config_bypass_bypass_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_bypass_bypass_proto_rawDescData)
	})
	return file_config_bypass_bypass_proto_rawDescData
}

var file_config_bypass_bypass_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_config_bypass_bypass_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_config_bypass_bypass_proto_goTypes = []interface{}{
	(Mode)(0),            // 0: yuhaiin.bypass.mode
	(ResolveStrategy)(0), // 1: yuhaiin.bypass.resolve_strategy
	(*Config)(nil),       // 2: yuhaiin.bypass.config
	(*ModeConfig)(nil),   // 3: yuhaiin.bypass.mode_config
}
var file_config_bypass_bypass_proto_depIdxs = []int32{
	0, // 0: yuhaiin.bypass.config.tcp:type_name -> yuhaiin.bypass.mode
	0, // 1: yuhaiin.bypass.config.udp:type_name -> yuhaiin.bypass.mode
	3, // 2: yuhaiin.bypass.config.custom_rule_v3:type_name -> yuhaiin.bypass.mode_config
	0, // 3: yuhaiin.bypass.mode_config.mode:type_name -> yuhaiin.bypass.mode
	1, // 4: yuhaiin.bypass.mode_config.resolve_strategy:type_name -> yuhaiin.bypass.resolve_strategy
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_config_bypass_bypass_proto_init() }
func file_config_bypass_bypass_proto_init() {
	if File_config_bypass_bypass_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_config_bypass_bypass_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
		file_config_bypass_bypass_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModeConfig); i {
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
			RawDescriptor: file_config_bypass_bypass_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_config_bypass_bypass_proto_goTypes,
		DependencyIndexes: file_config_bypass_bypass_proto_depIdxs,
		EnumInfos:         file_config_bypass_bypass_proto_enumTypes,
		MessageInfos:      file_config_bypass_bypass_proto_msgTypes,
	}.Build()
	File_config_bypass_bypass_proto = out.File
	file_config_bypass_bypass_proto_rawDesc = nil
	file_config_bypass_bypass_proto_goTypes = nil
	file_config_bypass_bypass_proto_depIdxs = nil
}
