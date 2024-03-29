// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.4
// source: msg/vote.proto

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

type ProtoVoteBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Voter     uint64 `protobuf:"varint,1,opt,name=Voter,proto3" json:"Voter,omitempty"`
	Signature []byte `protobuf:"bytes,2,opt,name=Signature,proto3" json:"Signature,omitempty"`
}

func (x *ProtoVoteBody) Reset() {
	*x = ProtoVoteBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_vote_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtoVoteBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoVoteBody) ProtoMessage() {}

func (x *ProtoVoteBody) ProtoReflect() protoreflect.Message {
	mi := &file_msg_vote_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoVoteBody.ProtoReflect.Descriptor instead.
func (*ProtoVoteBody) Descriptor() ([]byte, []int) {
	return file_msg_vote_proto_rawDescGZIP(), []int{0}
}

func (x *ProtoVoteBody) GetVoter() uint64 {
	if x != nil {
		return x.Voter
	}
	return 0
}

func (x *ProtoVoteBody) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type ProtoVote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *ProtoVoteData `protobuf:"bytes,1,opt,name=Data,proto3" json:"Data,omitempty"`
	Body *ProtoVoteBody `protobuf:"bytes,2,opt,name=Body,proto3" json:"Body,omitempty"`
}

func (x *ProtoVote) Reset() {
	*x = ProtoVote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_vote_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtoVote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoVote) ProtoMessage() {}

func (x *ProtoVote) ProtoReflect() protoreflect.Message {
	mi := &file_msg_vote_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoVote.ProtoReflect.Descriptor instead.
func (*ProtoVote) Descriptor() ([]byte, []int) {
	return file_msg_vote_proto_rawDescGZIP(), []int{1}
}

func (x *ProtoVote) GetData() *ProtoVoteData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ProtoVote) GetBody() *ProtoVoteBody {
	if x != nil {
		return x.Body
	}
	return nil
}

// VoteData are the parts that are independent of the voter
type ProtoVoteData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockHash []byte `protobuf:"bytes,1,opt,name=BlockHash,proto3" json:"BlockHash,omitempty"` // Bk
	View      uint64 `protobuf:"varint,2,opt,name=View,proto3" json:"View,omitempty"`          // v
}

func (x *ProtoVoteData) Reset() {
	*x = ProtoVoteData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_vote_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtoVoteData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoVoteData) ProtoMessage() {}

func (x *ProtoVoteData) ProtoReflect() protoreflect.Message {
	mi := &file_msg_vote_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoVoteData.ProtoReflect.Descriptor instead.
func (*ProtoVoteData) Descriptor() ([]byte, []int) {
	return file_msg_vote_proto_rawDescGZIP(), []int{2}
}

func (x *ProtoVoteData) GetBlockHash() []byte {
	if x != nil {
		return x.BlockHash
	}
	return nil
}

func (x *ProtoVoteData) GetView() uint64 {
	if x != nil {
		return x.View
	}
	return 0
}

var File_msg_vote_proto protoreflect.FileDescriptor

var file_msg_vote_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6d, 0x73, 0x67, 0x2f, 0x76, 0x6f, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x43, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x56, 0x6f,
	0x74, 0x65, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x6f, 0x74, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x56, 0x6f, 0x74, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09,
	0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x09, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x5b, 0x0a, 0x09, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x26, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x56, 0x6f, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x26, 0x0a, 0x04, 0x42, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x6d, 0x73, 0x67, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x56, 0x6f, 0x74, 0x65, 0x42, 0x6f, 0x64,
	0x79, 0x52, 0x04, 0x42, 0x6f, 0x64, 0x79, 0x22, 0x41, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x56, 0x6f, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x48, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x56, 0x69, 0x65, 0x77, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x56, 0x69, 0x65, 0x77, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x64, 0x69, 0x74, 0x68, 0x79, 0x61,
	0x62, 0x68, 0x61, 0x74, 0x6b, 0x61, 0x6a, 0x61, 0x6b, 0x65, 0x2f, 0x6c, 0x69, 0x62, 0x73, 0x79,
	0x6e, 0x63, 0x68, 0x73, 0x2f, 0x6d, 0x73, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_msg_vote_proto_rawDescOnce sync.Once
	file_msg_vote_proto_rawDescData = file_msg_vote_proto_rawDesc
)

func file_msg_vote_proto_rawDescGZIP() []byte {
	file_msg_vote_proto_rawDescOnce.Do(func() {
		file_msg_vote_proto_rawDescData = protoimpl.X.CompressGZIP(file_msg_vote_proto_rawDescData)
	})
	return file_msg_vote_proto_rawDescData
}

var file_msg_vote_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_msg_vote_proto_goTypes = []interface{}{
	(*ProtoVoteBody)(nil), // 0: msg.ProtoVoteBody
	(*ProtoVote)(nil),     // 1: msg.ProtoVote
	(*ProtoVoteData)(nil), // 2: msg.ProtoVoteData
}
var file_msg_vote_proto_depIdxs = []int32{
	2, // 0: msg.ProtoVote.Data:type_name -> msg.ProtoVoteData
	0, // 1: msg.ProtoVote.Body:type_name -> msg.ProtoVoteBody
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_msg_vote_proto_init() }
func file_msg_vote_proto_init() {
	if File_msg_vote_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_msg_vote_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtoVoteBody); i {
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
		file_msg_vote_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtoVote); i {
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
		file_msg_vote_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtoVoteData); i {
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
			RawDescriptor: file_msg_vote_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_msg_vote_proto_goTypes,
		DependencyIndexes: file_msg_vote_proto_depIdxs,
		MessageInfos:      file_msg_vote_proto_msgTypes,
	}.Build()
	File_msg_vote_proto = out.File
	file_msg_vote_proto_rawDesc = nil
	file_msg_vote_proto_goTypes = nil
	file_msg_vote_proto_depIdxs = nil
}
