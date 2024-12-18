// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: core/v1/review.proto

package core

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

type Review struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Rating uint32 `protobuf:"varint,2,opt,name=rating,proto3" json:"rating,omitempty"`
	Review string `protobuf:"bytes,3,opt,name=review,proto3" json:"review,omitempty"`
}

func (x *Review) Reset() {
	*x = Review{}
	mi := &file_core_v1_review_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Review) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Review) ProtoMessage() {}

func (x *Review) ProtoReflect() protoreflect.Message {
	mi := &file_core_v1_review_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Review.ProtoReflect.Descriptor instead.
func (*Review) Descriptor() ([]byte, []int) {
	return file_core_v1_review_proto_rawDescGZIP(), []int{0}
}

func (x *Review) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Review) GetRating() uint32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *Review) GetReview() string {
	if x != nil {
		return x.Review
	}
	return ""
}

var File_core_v1_review_proto protoreflect.FileDescriptor

var file_core_v1_review_proto_rawDesc = []byte{
	0x0a, 0x14, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x22,
	0x48, 0x0a, 0x06, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e,
	0x67, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x42, 0x85, 0x01, 0x0a, 0x0b, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x52, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x7a, 0x43, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2f, 0x70, 0x62, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2d, 0x67,
	0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0xa2, 0x02, 0x03, 0x43, 0x58, 0x58, 0xaa, 0x02, 0x07, 0x43,
	0x6f, 0x72, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x07, 0x43, 0x6f, 0x72, 0x65, 0x5c, 0x56, 0x31,
	0xe2, 0x02, 0x13, 0x43, 0x6f, 0x72, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x08, 0x43, 0x6f, 0x72, 0x65, 0x3a, 0x3a, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_core_v1_review_proto_rawDescOnce sync.Once
	file_core_v1_review_proto_rawDescData = file_core_v1_review_proto_rawDesc
)

func file_core_v1_review_proto_rawDescGZIP() []byte {
	file_core_v1_review_proto_rawDescOnce.Do(func() {
		file_core_v1_review_proto_rawDescData = protoimpl.X.CompressGZIP(file_core_v1_review_proto_rawDescData)
	})
	return file_core_v1_review_proto_rawDescData
}

var file_core_v1_review_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_core_v1_review_proto_goTypes = []any{
	(*Review)(nil), // 0: core.v1.Review
}
var file_core_v1_review_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_core_v1_review_proto_init() }
func file_core_v1_review_proto_init() {
	if File_core_v1_review_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_core_v1_review_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_core_v1_review_proto_goTypes,
		DependencyIndexes: file_core_v1_review_proto_depIdxs,
		MessageInfos:      file_core_v1_review_proto_msgTypes,
	}.Build()
	File_core_v1_review_proto = out.File
	file_core_v1_review_proto_rawDesc = nil
	file_core_v1_review_proto_goTypes = nil
	file_core_v1_review_proto_depIdxs = nil
}
