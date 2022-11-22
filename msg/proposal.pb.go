// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: msg/proposal.proto

package msg

import (
	chain  "github.com/adithyabhatkajake/libsynchs/chain"
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

// All configuration aggregated here
type Proposal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A proposal contains the following
	// The block being proposed
	Block *chain.ProtoBlock `protobuf:"bytes,1,opt,name=Block,proto3" json:"Block,omitempty"`
	// The forward sender of this propsoal innitial value is nil
	ForwardSender uint64 `protobuf:"varint,2,opt,name=ForwardSender,proto3" json:"ForwardSender,omitempty"`
	// The signature of sender
	ForwardSig []byte `protobuf:"bytes,3,opt,name=ForwardSig,proto3" json:"ForwardSig,omitempty"`
	// The miner of the block
	Miner uint64 `protobuf:"varint,4,opt,name=Miner,proto3" json:"Miner,omitempty"`
	// The signature for the block
	MiningProof []byte `protobuf:"bytes,5,opt,name=MiningProof,proto3" json:"MiningProof,omitempty"`
	// The view number
	View uint64 `protobuf:"varint,6,opt,name=View,proto3" json:"View,omitempty"`
	// Propose Evidence - Evidence why this node should propose this block
	// Can be a plain signature for closed world systems
	// Can be certificate for previous block as in Sync HotStuff/HotStuff
	// Can be a nonce in PoW systems
	ProposeEvidence []byte `protobuf:"bytes,7,opt,name=ProposeEvidence,proto3" json:"ProposeEvidence,omitempty"`
}

func (x *Proposal) Reset() {
	*x = Proposal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_proposal_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Proposal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Proposal) ProtoMessage() {}

func (x *Proposal) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proposal_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Proposal.ProtoReflect.Descriptor instead.
func (*Proposal) Descriptor() ([]byte, []int) {
	return file_msg_proposal_proto_rawDescGZIP(), []int{0}
}

func (x *Proposal) GetBlock() *chain.ProtoBlock {
	if x != nil {
		return x.Block
	}
	return nil
}

func (x *Proposal) GetForwardSender() uint64 {
	if x != nil {
		return x.ForwardSender
	}
	return 0
}

func (x *Proposal) GetForwardSig() []byte {
	if x != nil {
		return x.ForwardSig
	}
	return nil
}

func (x *Proposal) GetMiner() uint64 {
	if x != nil {
		return x.Miner
	}
	return 0
}

func (x *Proposal) GetMiningProof() []byte {
	if x != nil {
		return x.MiningProof
	}
	return nil
}

func (x *Proposal) GetView() uint64 {
	if x != nil {
		return x.View
	}
	return 0
}

func (x *Proposal) GetProposeEvidence() []byte {
	if x != nil {
		return x.ProposeEvidence
	}
	return nil
}

var File_msg_proposal_proto protoreflect.FileDescriptor

var file_msg_proposal_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6d, 0x73, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6d, 0x73, 0x67, 0x1a, 0x11, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xef, 0x01, 0x0a,
	0x08, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x12, 0x27, 0x0a, 0x05, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x05, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x12, 0x24, 0x0a, 0x0d, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x53, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x46, 0x6f, 0x72, 0x77, 0x61,
	0x72, 0x64, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x46, 0x6f, 0x72, 0x77,
	0x61, 0x72, 0x64, 0x53, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x46, 0x6f,
	0x72, 0x77, 0x61, 0x72, 0x64, 0x53, 0x69, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x4d, 0x69, 0x6e, 0x65,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x4d, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x20,
	0x0a, 0x0b, 0x4d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x0b, 0x4d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x6f, 0x66,
	0x12, 0x12, 0x0a, 0x04, 0x56, 0x69, 0x65, 0x77, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04,
	0x56, 0x69, 0x65, 0x77, 0x12, 0x28, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x45,
	0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x50,
	0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x45, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x42, 0x3c,
	0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x65, 0x70,
	0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x62, 0x61, 0x73, 0x65, 0x64, 0x2d, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x2d, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2d, 0x72, 0x65, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x2f, 0x6d, 0x73, 0x67, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_msg_proposal_proto_rawDescOnce sync.Once
	file_msg_proposal_proto_rawDescData = file_msg_proposal_proto_rawDesc
)

func file_msg_proposal_proto_rawDescGZIP() []byte {
	file_msg_proposal_proto_rawDescOnce.Do(func() {
		file_msg_proposal_proto_rawDescData = protoimpl.X.CompressGZIP(file_msg_proposal_proto_rawDescData)
	})
	return file_msg_proposal_proto_rawDescData
}

var file_msg_proposal_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_msg_proposal_proto_goTypes = []interface{}{
	(*Proposal)(nil),         // 0: msg.Proposal
	(*chain.ProtoBlock)(nil), // 1: chain.ProtoBlock
}
var file_msg_proposal_proto_depIdxs = []int32{
	1, // 0: msg.Proposal.Block:type_name -> chain.ProtoBlock
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_msg_proposal_proto_init() }
func file_msg_proposal_proto_init() {
	if File_msg_proposal_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_msg_proposal_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Proposal); i {
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
			RawDescriptor: file_msg_proposal_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_msg_proposal_proto_goTypes,
		DependencyIndexes: file_msg_proposal_proto_depIdxs,
		MessageInfos:      file_msg_proposal_proto_msgTypes,
	}.Build()
	File_msg_proposal_proto = out.File
	file_msg_proposal_proto_rawDesc = nil
	file_msg_proposal_proto_goTypes = nil
	file_msg_proposal_proto_depIdxs = nil
}
