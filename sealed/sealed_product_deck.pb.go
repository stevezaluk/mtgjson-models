// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.6.1
// source: sealed/sealed_product_deck.proto

package sealed

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

type SealedProductDeck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" bson:"name"` // @gotags: bson:"name"
	Set  string `protobuf:"bytes,2,opt,name=set,proto3" json:"set,omitempty" bson:"set"`   // @gotags: bson:"set"
}

func (x *SealedProductDeck) Reset() {
	*x = SealedProductDeck{}
	mi := &file_sealed_sealed_product_deck_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SealedProductDeck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SealedProductDeck) ProtoMessage() {}

func (x *SealedProductDeck) ProtoReflect() protoreflect.Message {
	mi := &file_sealed_sealed_product_deck_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SealedProductDeck.ProtoReflect.Descriptor instead.
func (*SealedProductDeck) Descriptor() ([]byte, []int) {
	return file_sealed_sealed_product_deck_proto_rawDescGZIP(), []int{0}
}

func (x *SealedProductDeck) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SealedProductDeck) GetSet() string {
	if x != nil {
		return x.Set
	}
	return ""
}

var File_sealed_sealed_product_deck_proto protoreflect.FileDescriptor

var file_sealed_sealed_product_deck_proto_rawDesc = []byte{
	0x0a, 0x20, 0x73, 0x65, 0x61, 0x6c, 0x65, 0x64, 0x2f, 0x73, 0x65, 0x61, 0x6c, 0x65, 0x64, 0x5f,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x64, 0x65, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x06, 0x73, 0x65, 0x61, 0x6c, 0x65, 0x64, 0x22, 0x39, 0x0a, 0x11, 0x53, 0x65,
	0x61, 0x6c, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x65, 0x63, 0x6b, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x73, 0x65, 0x74, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x65, 0x76, 0x65, 0x7a, 0x61, 0x6c, 0x75, 0x6b, 0x2f, 0x6d,
	0x74, 0x67, 0x6a, 0x73, 0x6f, 0x6e, 0x2d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x73, 0x65,
	0x61, 0x6c, 0x65, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sealed_sealed_product_deck_proto_rawDescOnce sync.Once
	file_sealed_sealed_product_deck_proto_rawDescData = file_sealed_sealed_product_deck_proto_rawDesc
)

func file_sealed_sealed_product_deck_proto_rawDescGZIP() []byte {
	file_sealed_sealed_product_deck_proto_rawDescOnce.Do(func() {
		file_sealed_sealed_product_deck_proto_rawDescData = protoimpl.X.CompressGZIP(file_sealed_sealed_product_deck_proto_rawDescData)
	})
	return file_sealed_sealed_product_deck_proto_rawDescData
}

var file_sealed_sealed_product_deck_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_sealed_sealed_product_deck_proto_goTypes = []any{
	(*SealedProductDeck)(nil), // 0: sealed.SealedProductDeck
}
var file_sealed_sealed_product_deck_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sealed_sealed_product_deck_proto_init() }
func file_sealed_sealed_product_deck_proto_init() {
	if File_sealed_sealed_product_deck_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sealed_sealed_product_deck_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sealed_sealed_product_deck_proto_goTypes,
		DependencyIndexes: file_sealed_sealed_product_deck_proto_depIdxs,
		MessageInfos:      file_sealed_sealed_product_deck_proto_msgTypes,
	}.Build()
	File_sealed_sealed_product_deck_proto = out.File
	file_sealed_sealed_product_deck_proto_rawDesc = nil
	file_sealed_sealed_product_deck_proto_goTypes = nil
	file_sealed_sealed_product_deck_proto_depIdxs = nil
}
