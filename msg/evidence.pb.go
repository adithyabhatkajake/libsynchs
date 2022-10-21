// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: msg/evidence.proto

package msg

import (
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

type EvidenceData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MisbehaviourTarget uint64 `protobuf:"varint,1,opt,name=MisbehaviourTarget,proto3" json:"MisbehaviourTarget,omitempty"` //The misbehaviour target
	View               uint64 `protobuf:"varint,2,opt,name=view,proto3" json:"view,omitempty"`                             // The view number which produce this evidence
}

func (x *EvidenceData) Reset() {
	*x = EvidenceData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_evidence_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EvidenceData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvidenceData) ProtoMessage() {}

func (x *EvidenceData) ProtoReflect() protoreflect.Message {
	mi := &file_msg_evidence_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EvidenceData.ProtoReflect.Descriptor instead.
func (*EvidenceData) Descriptor() ([]byte, []int) {
	return file_msg_evidence_proto_rawDescGZIP(), []int{0}
}

func (x *EvidenceData) GetMisbehaviourTarget() uint64 {
	if x != nil {
		return x.MisbehaviourTarget
	}
	return 0
}

func (x *EvidenceData) GetView() uint64 {
	if x != nil {
		return x.View
	}
	return 0
}

type Evidence struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EvidenceData *EvidenceData `protobuf:"bytes,1,opt,name=EvidenceData,proto3" json:"EvidenceData,omitempty"`
	EvOrigin     uint64        `protobuf:"varint,2,opt,name=EvOrigin,proto3" json:"EvOrigin,omitempty"`      // The sender of this evidence
	OrSignature  []byte        `protobuf:"bytes,3,opt,name=OrSignature,proto3" json:"OrSignature,omitempty"` //The sigature of this sender
}

func (x *Evidence) Reset() {
	*x = Evidence{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_evidence_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Evidence) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Evidence) ProtoMessage() {}

func (x *Evidence) ProtoReflect() protoreflect.Message {
	mi := &file_msg_evidence_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Evidence.ProtoReflect.Descriptor instead.
func (*Evidence) Descriptor() ([]byte, []int) {
	return file_msg_evidence_proto_rawDescGZIP(), []int{1}
}

func (x *Evidence) GetEvidenceData() *EvidenceData {
	if x != nil {
		return x.EvidenceData
	}
	return nil
}

func (x *Evidence) GetEvOrigin() uint64 {
	if x != nil {
		return x.EvOrigin
	}
	return 0
}

func (x *Evidence) GetOrSignature() []byte {
	if x != nil {
		return x.OrSignature
	}
	return nil
}

type EquivocationEvidence struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Evidence *Evidence `protobuf:"bytes,1,opt,name=Evidence,proto3" json:"Evidence,omitempty"`
	E1       *Proposal `protobuf:"bytes,2,opt,name=e1,proto3" json:"e1,omitempty"` //the content of this evidence is
	E2       *Proposal `protobuf:"bytes,3,opt,name=e2,proto3" json:"e2,omitempty"` //(two blocks and corresponding signatures and the pubkey of this leader)
}

func (x *EquivocationEvidence) Reset() {
	*x = EquivocationEvidence{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_evidence_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EquivocationEvidence) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EquivocationEvidence) ProtoMessage() {}

func (x *EquivocationEvidence) ProtoReflect() protoreflect.Message {
	mi := &file_msg_evidence_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EquivocationEvidence.ProtoReflect.Descriptor instead.
func (*EquivocationEvidence) Descriptor() ([]byte, []int) {
	return file_msg_evidence_proto_rawDescGZIP(), []int{2}
}

func (x *EquivocationEvidence) GetEvidence() *Evidence {
	if x != nil {
		return x.Evidence
	}
	return nil
}

func (x *EquivocationEvidence) GetE1() *Proposal {
	if x != nil {
		return x.E1
	}
	return nil
}

func (x *EquivocationEvidence) GetE2() *Proposal {
	if x != nil {
		return x.E2
	}
	return nil
}

type MalicousProposalEvidence struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Evidence *Evidence `protobuf:"bytes,1,opt,name=Evidence,proto3" json:"Evidence,omitempty"` //but sent the proposal (the content of this )
	E        *Proposal `protobuf:"bytes,2,opt,name=e,proto3" json:"e,omitempty"`
}

func (x *MalicousProposalEvidence) Reset() {
	*x = MalicousProposalEvidence{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_evidence_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MalicousProposalEvidence) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MalicousProposalEvidence) ProtoMessage() {}

func (x *MalicousProposalEvidence) ProtoReflect() protoreflect.Message {
	mi := &file_msg_evidence_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MalicousProposalEvidence.ProtoReflect.Descriptor instead.
func (*MalicousProposalEvidence) Descriptor() ([]byte, []int) {
	return file_msg_evidence_proto_rawDescGZIP(), []int{3}
}

func (x *MalicousProposalEvidence) GetEvidence() *Evidence {
	if x != nil {
		return x.Evidence
	}
	return nil
}

func (x *MalicousProposalEvidence) GetE() *Proposal {
	if x != nil {
		return x.E
	}
	return nil
}

type MalicousVoteEvidence struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Evidence *Evidence  `protobuf:"bytes,1,opt,name=Evidence,proto3" json:"Evidence,omitempty"`
	E        *ProtoVote `protobuf:"bytes,2,opt,name=e,proto3" json:"e,omitempty"`
}

func (x *MalicousVoteEvidence) Reset() {
	*x = MalicousVoteEvidence{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_evidence_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MalicousVoteEvidence) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MalicousVoteEvidence) ProtoMessage() {}

func (x *MalicousVoteEvidence) ProtoReflect() protoreflect.Message {
	mi := &file_msg_evidence_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MalicousVoteEvidence.ProtoReflect.Descriptor instead.
func (*MalicousVoteEvidence) Descriptor() ([]byte, []int) {
	return file_msg_evidence_proto_rawDescGZIP(), []int{4}
}

func (x *MalicousVoteEvidence) GetEvidence() *Evidence {
	if x != nil {
		return x.Evidence
	}
	return nil
}

func (x *MalicousVoteEvidence) GetE() *ProtoVote {
	if x != nil {
		return x.E
	}
	return nil
}

var File_msg_evidence_proto protoreflect.FileDescriptor

var file_msg_evidence_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6d, 0x73, 0x67, 0x2f, 0x65, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6d, 0x73, 0x67, 0x1a, 0x12, 0x6d, 0x73, 0x67, 0x2f, 0x70,
	0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x6d,
	0x73, 0x67, 0x2f, 0x76, 0x6f, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x52, 0x0a,
	0x0c, 0x45, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2e, 0x0a,
	0x12, 0x4d, 0x69, 0x73, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x75, 0x72, 0x54, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x12, 0x4d, 0x69, 0x73, 0x62, 0x65,
	0x68, 0x61, 0x76, 0x69, 0x6f, 0x75, 0x72, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x76, 0x69, 0x65, 0x77, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x76, 0x69, 0x65,
	0x77, 0x22, 0x7f, 0x0a, 0x08, 0x45, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x35, 0x0a,
	0x0c, 0x45, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x45, 0x76, 0x69, 0x64, 0x65, 0x6e,
	0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0c, 0x45, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x45, 0x76, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x45, 0x76, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x12, 0x20, 0x0a, 0x0b, 0x4f, 0x72, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x4f, 0x72, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x22, 0x7f, 0x0a, 0x14, 0x45, 0x71, 0x75, 0x69, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x45, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x29, 0x0a, 0x08, 0x45, 0x76,
	0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6d,
	0x73, 0x67, 0x2e, 0x45, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x08, 0x45, 0x76, 0x69,
	0x64, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x02, 0x65, 0x31, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c,
	0x52, 0x02, 0x65, 0x31, 0x12, 0x1d, 0x0a, 0x02, 0x65, 0x32, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x52,
	0x02, 0x65, 0x32, 0x22, 0x62, 0x0a, 0x18, 0x4d, 0x61, 0x6c, 0x69, 0x63, 0x6f, 0x75, 0x73, 0x50,
	0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x45, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x12,
	0x29, 0x0a, 0x08, 0x45, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x45, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65,
	0x52, 0x08, 0x45, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x01, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x50, 0x72, 0x6f, 0x70,
	0x6f, 0x73, 0x61, 0x6c, 0x52, 0x01, 0x65, 0x22, 0x5f, 0x0a, 0x14, 0x4d, 0x61, 0x6c, 0x69, 0x63,
	0x6f, 0x75, 0x73, 0x56, 0x6f, 0x74, 0x65, 0x45, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x12,
	0x29, 0x0a, 0x08, 0x45, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x45, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65,
	0x52, 0x08, 0x45, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x01, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x01, 0x65, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x65, 0x70, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2d, 0x62, 0x61, 0x73, 0x65, 0x64, 0x2d, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2d, 0x6d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x2d, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2d, 0x2f, 0x6d, 0x73, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_msg_evidence_proto_rawDescOnce sync.Once
	file_msg_evidence_proto_rawDescData = file_msg_evidence_proto_rawDesc
)

func file_msg_evidence_proto_rawDescGZIP() []byte {
	file_msg_evidence_proto_rawDescOnce.Do(func() {
		file_msg_evidence_proto_rawDescData = protoimpl.X.CompressGZIP(file_msg_evidence_proto_rawDescData)
	})
	return file_msg_evidence_proto_rawDescData
}

var file_msg_evidence_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_msg_evidence_proto_goTypes = []interface{}{
	(*EvidenceData)(nil),             // 0: msg.EvidenceData
	(*Evidence)(nil),                 // 1: msg.Evidence
	(*EquivocationEvidence)(nil),     // 2: msg.EquivocationEvidence
	(*MalicousProposalEvidence)(nil), // 3: msg.MalicousProposalEvidence
	(*MalicousVoteEvidence)(nil),     // 4: msg.MalicousVoteEvidence
	(*Proposal)(nil),                 // 5: msg.Proposal
	(*ProtoVote)(nil),                // 6: msg.ProtoVote
}
var file_msg_evidence_proto_depIdxs = []int32{
	0, // 0: msg.Evidence.EvidenceData:type_name -> msg.EvidenceData
	1, // 1: msg.EquivocationEvidence.Evidence:type_name -> msg.Evidence
	5, // 2: msg.EquivocationEvidence.e1:type_name -> msg.Proposal
	5, // 3: msg.EquivocationEvidence.e2:type_name -> msg.Proposal
	1, // 4: msg.MalicousProposalEvidence.Evidence:type_name -> msg.Evidence
	5, // 5: msg.MalicousProposalEvidence.e:type_name -> msg.Proposal
	1, // 6: msg.MalicousVoteEvidence.Evidence:type_name -> msg.Evidence
	6, // 7: msg.MalicousVoteEvidence.e:type_name -> msg.ProtoVote
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_msg_evidence_proto_init() }
func file_msg_evidence_proto_init() {
	if File_msg_evidence_proto != nil {
		return
	}
	file_msg_proposal_proto_init()
	file_msg_vote_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_msg_evidence_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EvidenceData); i {
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
		file_msg_evidence_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Evidence); i {
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
		file_msg_evidence_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EquivocationEvidence); i {
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
		file_msg_evidence_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MalicousProposalEvidence); i {
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
		file_msg_evidence_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MalicousVoteEvidence); i {
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
			RawDescriptor: file_msg_evidence_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_msg_evidence_proto_goTypes,
		DependencyIndexes: file_msg_evidence_proto_depIdxs,
		MessageInfos:      file_msg_evidence_proto_msgTypes,
	}.Build()
	File_msg_evidence_proto = out.File
	file_msg_evidence_proto_rawDesc = nil
	file_msg_evidence_proto_goTypes = nil
	file_msg_evidence_proto_depIdxs = nil
}
