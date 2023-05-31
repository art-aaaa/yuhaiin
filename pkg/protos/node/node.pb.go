// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: node/node.proto

package node

import (
	point "github.com/Asutorufa/yuhaiin/pkg/protos/node/point"
	subscribe "github.com/Asutorufa/yuhaiin/pkg/protos/node/subscribe"
	tag "github.com/Asutorufa/yuhaiin/pkg/protos/node/tag"
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

type Node struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tcp     *point.Point               `protobuf:"bytes,4,opt,name=tcp,proto3" json:"tcp,omitempty"`
	Udp     *point.Point               `protobuf:"bytes,5,opt,name=udp,proto3" json:"udp,omitempty"`
	Links   map[string]*subscribe.Link `protobuf:"bytes,2,rep,name=links,proto3" json:"links,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Manager *Manager                   `protobuf:"bytes,3,opt,name=manager,proto3" json:"manager,omitempty"`
}

func (x *Node) Reset() {
	*x = Node{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_node_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Node) ProtoMessage() {}

func (x *Node) ProtoReflect() protoreflect.Message {
	mi := &file_node_node_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Node.ProtoReflect.Descriptor instead.
func (*Node) Descriptor() ([]byte, []int) {
	return file_node_node_proto_rawDescGZIP(), []int{0}
}

func (x *Node) GetTcp() *point.Point {
	if x != nil {
		return x.Tcp
	}
	return nil
}

func (x *Node) GetUdp() *point.Point {
	if x != nil {
		return x.Udp
	}
	return nil
}

func (x *Node) GetLinks() map[string]*subscribe.Link {
	if x != nil {
		return x.Links
	}
	return nil
}

func (x *Node) GetManager() *Manager {
	if x != nil {
		return x.Manager
	}
	return nil
}

type Nodes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodesV2 map[string]string `protobuf:"bytes,3,rep,name=nodesV2,json=node_hash_map,proto3" json:"nodesV2,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Nodes) Reset() {
	*x = Nodes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_node_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Nodes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Nodes) ProtoMessage() {}

func (x *Nodes) ProtoReflect() protoreflect.Message {
	mi := &file_node_node_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Nodes.ProtoReflect.Descriptor instead.
func (*Nodes) Descriptor() ([]byte, []int) {
	return file_node_node_proto_rawDescGZIP(), []int{1}
}

func (x *Nodes) GetNodesV2() map[string]string {
	if x != nil {
		return x.NodesV2
	}
	return nil
}

type Manager struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupsV2 map[string]*Nodes       `protobuf:"bytes,2,rep,name=groupsV2,json=group_nodes_map,proto3" json:"groupsV2,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Nodes    map[string]*point.Point `protobuf:"bytes,3,rep,name=nodes,proto3" json:"nodes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Tags     map[string]*tag.Tags    `protobuf:"bytes,4,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Manager) Reset() {
	*x = Manager{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_node_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Manager) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Manager) ProtoMessage() {}

func (x *Manager) ProtoReflect() protoreflect.Message {
	mi := &file_node_node_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Manager.ProtoReflect.Descriptor instead.
func (*Manager) Descriptor() ([]byte, []int) {
	return file_node_node_proto_rawDescGZIP(), []int{2}
}

func (x *Manager) GetGroupsV2() map[string]*Nodes {
	if x != nil {
		return x.GroupsV2
	}
	return nil
}

func (x *Manager) GetNodes() map[string]*point.Point {
	if x != nil {
		return x.Nodes
	}
	return nil
}

func (x *Manager) GetTags() map[string]*tag.Tags {
	if x != nil {
		return x.Tags
	}
	return nil
}

var File_node_node_proto protoreflect.FileDescriptor

var file_node_node_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x79, 0x75, 0x68, 0x61, 0x69, 0x69, 0x6e, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x1a,
	0x16, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2f, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x73, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x74, 0x61,
	0x67, 0x2f, 0x74, 0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8f, 0x02, 0x0a, 0x04,
	0x6e, 0x6f, 0x64, 0x65, 0x12, 0x26, 0x0a, 0x03, 0x74, 0x63, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x79, 0x75, 0x68, 0x61, 0x69, 0x69, 0x6e, 0x2e, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x2e, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x03, 0x74, 0x63, 0x70, 0x12, 0x26, 0x0a, 0x03,
	0x75, 0x64, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x79, 0x75, 0x68, 0x61,
	0x69, 0x69, 0x6e, 0x2e, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52,
	0x03, 0x75, 0x64, 0x70, 0x12, 0x33, 0x0a, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x79, 0x75, 0x68, 0x61, 0x69, 0x69, 0x6e, 0x2e, 0x6e, 0x6f,
	0x64, 0x65, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x12, 0x2f, 0x0a, 0x07, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x79, 0x75, 0x68,
	0x61, 0x69, 0x69, 0x6e, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x52, 0x07, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x1a, 0x51, 0x0a, 0x0a, 0x4c, 0x69,
	0x6e, 0x6b, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2d, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x79, 0x75, 0x68, 0x61,
	0x69, 0x69, 0x6e, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x2e, 0x6c, 0x69,
	0x6e, 0x6b, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x85, 0x01,
	0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x40, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x73,
	0x56, 0x32, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x79, 0x75, 0x68, 0x61, 0x69,
	0x69, 0x6e, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x2e, 0x4e, 0x6f,
	0x64, 0x65, 0x73, 0x56, 0x32, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d, 0x6e, 0x6f, 0x64, 0x65,
	0x5f, 0x68, 0x61, 0x73, 0x68, 0x5f, 0x6d, 0x61, 0x70, 0x1a, 0x3a, 0x0a, 0x0c, 0x4e, 0x6f, 0x64,
	0x65, 0x73, 0x56, 0x32, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xac, 0x03, 0x0a, 0x07, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x12, 0x46, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x56, 0x32, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x79, 0x75, 0x68, 0x61, 0x69, 0x69, 0x6e, 0x2e, 0x6e, 0x6f,
	0x64, 0x65, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x73, 0x56, 0x32, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f,
	0x6e, 0x6f, 0x64, 0x65, 0x73, 0x5f, 0x6d, 0x61, 0x70, 0x12, 0x36, 0x0a, 0x05, 0x6e, 0x6f, 0x64,
	0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x79, 0x75, 0x68, 0x61, 0x69,
	0x69, 0x6e, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e,
	0x4e, 0x6f, 0x64, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x6e, 0x6f, 0x64, 0x65,
	0x73, 0x12, 0x33, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1f, 0x2e, 0x79, 0x75, 0x68, 0x61, 0x69, 0x69, 0x6e, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x54, 0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x1a, 0x50, 0x0a, 0x0d, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73,
	0x56, 0x32, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x29, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x79, 0x75, 0x68, 0x61, 0x69,
	0x69, 0x6e, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x4e, 0x0a, 0x0a, 0x4e, 0x6f, 0x64, 0x65,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x79, 0x75, 0x68, 0x61, 0x69, 0x69,
	0x6e, 0x2e, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x4a, 0x0a, 0x09, 0x54, 0x61, 0x67, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x27, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x79, 0x75, 0x68, 0x61, 0x69, 0x69, 0x6e,
	0x2e, 0x74, 0x61, 0x67, 0x2e, 0x74, 0x61, 0x67, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x41, 0x73, 0x75, 0x74, 0x6f, 0x72, 0x75, 0x66, 0x61, 0x2f, 0x79, 0x75, 0x68,
	0x61, 0x69, 0x69, 0x6e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f,
	0x6e, 0x6f, 0x64, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_node_node_proto_rawDescOnce sync.Once
	file_node_node_proto_rawDescData = file_node_node_proto_rawDesc
)

func file_node_node_proto_rawDescGZIP() []byte {
	file_node_node_proto_rawDescOnce.Do(func() {
		file_node_node_proto_rawDescData = protoimpl.X.CompressGZIP(file_node_node_proto_rawDescData)
	})
	return file_node_node_proto_rawDescData
}

var file_node_node_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_node_node_proto_goTypes = []interface{}{
	(*Node)(nil),           // 0: yuhaiin.node.node
	(*Nodes)(nil),          // 1: yuhaiin.node.nodes
	(*Manager)(nil),        // 2: yuhaiin.node.manager
	nil,                    // 3: yuhaiin.node.node.LinksEntry
	nil,                    // 4: yuhaiin.node.nodes.NodesV2Entry
	nil,                    // 5: yuhaiin.node.manager.GroupsV2Entry
	nil,                    // 6: yuhaiin.node.manager.NodesEntry
	nil,                    // 7: yuhaiin.node.manager.TagsEntry
	(*point.Point)(nil),    // 8: yuhaiin.point.point
	(*subscribe.Link)(nil), // 9: yuhaiin.subscribe.link
	(*tag.Tags)(nil),       // 10: yuhaiin.tag.tags
}
var file_node_node_proto_depIdxs = []int32{
	8,  // 0: yuhaiin.node.node.tcp:type_name -> yuhaiin.point.point
	8,  // 1: yuhaiin.node.node.udp:type_name -> yuhaiin.point.point
	3,  // 2: yuhaiin.node.node.links:type_name -> yuhaiin.node.node.LinksEntry
	2,  // 3: yuhaiin.node.node.manager:type_name -> yuhaiin.node.manager
	4,  // 4: yuhaiin.node.nodes.nodesV2:type_name -> yuhaiin.node.nodes.NodesV2Entry
	5,  // 5: yuhaiin.node.manager.groupsV2:type_name -> yuhaiin.node.manager.GroupsV2Entry
	6,  // 6: yuhaiin.node.manager.nodes:type_name -> yuhaiin.node.manager.NodesEntry
	7,  // 7: yuhaiin.node.manager.tags:type_name -> yuhaiin.node.manager.TagsEntry
	9,  // 8: yuhaiin.node.node.LinksEntry.value:type_name -> yuhaiin.subscribe.link
	1,  // 9: yuhaiin.node.manager.GroupsV2Entry.value:type_name -> yuhaiin.node.nodes
	8,  // 10: yuhaiin.node.manager.NodesEntry.value:type_name -> yuhaiin.point.point
	10, // 11: yuhaiin.node.manager.TagsEntry.value:type_name -> yuhaiin.tag.tags
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_node_node_proto_init() }
func file_node_node_proto_init() {
	if File_node_node_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_node_node_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Node); i {
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
		file_node_node_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Nodes); i {
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
		file_node_node_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Manager); i {
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
			RawDescriptor: file_node_node_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_node_node_proto_goTypes,
		DependencyIndexes: file_node_node_proto_depIdxs,
		MessageInfos:      file_node_node_proto_msgTypes,
	}.Build()
	File_node_node_proto = out.File
	file_node_node_proto_rawDesc = nil
	file_node_node_proto_goTypes = nil
	file_node_node_proto_depIdxs = nil
}
