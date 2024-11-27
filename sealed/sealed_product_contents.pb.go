// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.6.1
// source: sealed/sealed_product_contents.proto

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

type SealedProductContents struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Card     *SealedProductCard                `protobuf:"bytes,1,opt,name=card,proto3" json:"card,omitempty"`
	Deck     *SealedProductDeck                `protobuf:"bytes,2,opt,name=deck,proto3" json:"deck,omitempty"`
	Other    *SealedProductOther               `protobuf:"bytes,3,opt,name=other,proto3" json:"other,omitempty"`
	Pack     *SealedProductPack                `protobuf:"bytes,4,opt,name=pack,proto3" json:"pack,omitempty"`
	Sealed   *SealedProductSealed              `protobuf:"bytes,5,opt,name=sealed,proto3" json:"sealed,omitempty"`
	Variable map[string]*SealedProductContents `protobuf:"bytes,6,rep,name=variable,proto3" json:"variable,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *SealedProductContents) Reset() {
	*x = SealedProductContents{}
	mi := &file_sealed_sealed_product_contents_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SealedProductContents) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SealedProductContents) ProtoMessage() {}

func (x *SealedProductContents) ProtoReflect() protoreflect.Message {
	mi := &file_sealed_sealed_product_contents_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SealedProductContents.ProtoReflect.Descriptor instead.
func (*SealedProductContents) Descriptor() ([]byte, []int) {
	return file_sealed_sealed_product_contents_proto_rawDescGZIP(), []int{0}
}

func (x *SealedProductContents) GetCard() *SealedProductCard {
	if x != nil {
		return x.Card
	}
	return nil
}

func (x *SealedProductContents) GetDeck() *SealedProductDeck {
	if x != nil {
		return x.Deck
	}
	return nil
}

func (x *SealedProductContents) GetOther() *SealedProductOther {
	if x != nil {
		return x.Other
	}
	return nil
}

func (x *SealedProductContents) GetPack() *SealedProductPack {
	if x != nil {
		return x.Pack
	}
	return nil
}

func (x *SealedProductContents) GetSealed() *SealedProductSealed {
	if x != nil {
		return x.Sealed
	}
	return nil
}

func (x *SealedProductContents) GetVariable() map[string]*SealedProductContents {
	if x != nil {
		return x.Variable
	}
	return nil
}

var File_sealed_sealed_product_contents_proto protoreflect.FileDescriptor

var file_sealed_sealed_product_contents_proto_rawDesc = []byte{
	0x0a, 0x24, 0x73, 0x65, 0x61, 0x6c, 0x65, 0x64, 0x2f, 0x73, 0x65, 0x61, 0x6c, 0x65, 0x64, 0x5f,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x65, 0x61, 0x6c, 0x65, 0x64, 0x1a, 0x20,
	0x73, 0x65, 0x61, 0x6c, 0x65, 0x64, 0x2f, 0x73, 0x65, 0x61, 0x6c, 0x65, 0x64, 0x5f, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x20, 0x73, 0x65, 0x61, 0x6c, 0x65, 0x64, 0x2f, 0x73, 0x65, 0x61, 0x6c, 0x65, 0x64, 0x5f,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x64, 0x65, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x21, 0x73, 0x65, 0x61, 0x6c, 0x65, 0x64, 0x2f, 0x73, 0x65, 0x61, 0x6c, 0x65,
	0x64, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x73, 0x65, 0x61, 0x6c, 0x65, 0x64, 0x2f, 0x73, 0x65,
	0x61, 0x6c, 0x65, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x70, 0x61, 0x63,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x73, 0x65, 0x61, 0x6c, 0x65, 0x64, 0x2f,
	0x73, 0x65, 0x61, 0x6c, 0x65, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x73,
	0x65, 0x61, 0x6c, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb0, 0x03, 0x0a, 0x15,
	0x53, 0x65, 0x61, 0x6c, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x2d, 0x0a, 0x04, 0x63, 0x61, 0x72, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x65, 0x61, 0x6c, 0x65, 0x64, 0x2e, 0x53, 0x65, 0x61,
	0x6c, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x72, 0x64, 0x52, 0x04,
	0x63, 0x61, 0x72, 0x64, 0x12, 0x2d, 0x0a, 0x04, 0x64, 0x65, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x65, 0x61, 0x6c, 0x65, 0x64, 0x2e, 0x53, 0x65, 0x61, 0x6c,
	0x65, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x65, 0x63, 0x6b, 0x52, 0x04, 0x64,
	0x65, 0x63, 0x6b, 0x12, 0x30, 0x0a, 0x05, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x65, 0x61, 0x6c, 0x65, 0x64, 0x2e, 0x53, 0x65, 0x61, 0x6c,
	0x65, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x52, 0x05,
	0x6f, 0x74, 0x68, 0x65, 0x72, 0x12, 0x2d, 0x0a, 0x04, 0x70, 0x61, 0x63, 0x6b, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x65, 0x61, 0x6c, 0x65, 0x64, 0x2e, 0x53, 0x65, 0x61,
	0x6c, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x52, 0x04,
	0x70, 0x61, 0x63, 0x6b, 0x12, 0x33, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x6c, 0x65, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x65, 0x61, 0x6c, 0x65, 0x64, 0x2e, 0x53, 0x65,
	0x61, 0x6c, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x61, 0x6c, 0x65,
	0x64, 0x52, 0x06, 0x73, 0x65, 0x61, 0x6c, 0x65, 0x64, 0x12, 0x47, 0x0a, 0x08, 0x76, 0x61, 0x72,
	0x69, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x73, 0x65,
	0x61, 0x6c, 0x65, 0x64, 0x2e, 0x53, 0x65, 0x61, 0x6c, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x56, 0x61, 0x72, 0x69, 0x61,
	0x62, 0x6c, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62,
	0x6c, 0x65, 0x1a, 0x5a, 0x0a, 0x0d, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x33, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x73, 0x65, 0x61, 0x6c, 0x65, 0x64, 0x2e, 0x53, 0x65,
	0x61, 0x6c, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x2d,
	0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x65,
	0x76, 0x65, 0x7a, 0x61, 0x6c, 0x75, 0x6b, 0x2f, 0x6d, 0x74, 0x67, 0x6a, 0x73, 0x6f, 0x6e, 0x2d,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x73, 0x65, 0x61, 0x6c, 0x65, 0x64, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sealed_sealed_product_contents_proto_rawDescOnce sync.Once
	file_sealed_sealed_product_contents_proto_rawDescData = file_sealed_sealed_product_contents_proto_rawDesc
)

func file_sealed_sealed_product_contents_proto_rawDescGZIP() []byte {
	file_sealed_sealed_product_contents_proto_rawDescOnce.Do(func() {
		file_sealed_sealed_product_contents_proto_rawDescData = protoimpl.X.CompressGZIP(file_sealed_sealed_product_contents_proto_rawDescData)
	})
	return file_sealed_sealed_product_contents_proto_rawDescData
}

var file_sealed_sealed_product_contents_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_sealed_sealed_product_contents_proto_goTypes = []any{
	(*SealedProductContents)(nil), // 0: sealed.SealedProductContents
	nil,                           // 1: sealed.SealedProductContents.VariableEntry
	(*SealedProductCard)(nil),     // 2: sealed.SealedProductCard
	(*SealedProductDeck)(nil),     // 3: sealed.SealedProductDeck
	(*SealedProductOther)(nil),    // 4: sealed.SealedProductOther
	(*SealedProductPack)(nil),     // 5: sealed.SealedProductPack
	(*SealedProductSealed)(nil),   // 6: sealed.SealedProductSealed
}
var file_sealed_sealed_product_contents_proto_depIdxs = []int32{
	2, // 0: sealed.SealedProductContents.card:type_name -> sealed.SealedProductCard
	3, // 1: sealed.SealedProductContents.deck:type_name -> sealed.SealedProductDeck
	4, // 2: sealed.SealedProductContents.other:type_name -> sealed.SealedProductOther
	5, // 3: sealed.SealedProductContents.pack:type_name -> sealed.SealedProductPack
	6, // 4: sealed.SealedProductContents.sealed:type_name -> sealed.SealedProductSealed
	1, // 5: sealed.SealedProductContents.variable:type_name -> sealed.SealedProductContents.VariableEntry
	0, // 6: sealed.SealedProductContents.VariableEntry.value:type_name -> sealed.SealedProductContents
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_sealed_sealed_product_contents_proto_init() }
func file_sealed_sealed_product_contents_proto_init() {
	if File_sealed_sealed_product_contents_proto != nil {
		return
	}
	file_sealed_sealed_product_card_proto_init()
	file_sealed_sealed_product_deck_proto_init()
	file_sealed_sealed_product_other_proto_init()
	file_sealed_sealed_product_pack_proto_init()
	file_sealed_sealed_product_sealed_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sealed_sealed_product_contents_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sealed_sealed_product_contents_proto_goTypes,
		DependencyIndexes: file_sealed_sealed_product_contents_proto_depIdxs,
		MessageInfos:      file_sealed_sealed_product_contents_proto_msgTypes,
	}.Build()
	File_sealed_sealed_product_contents_proto = out.File
	file_sealed_sealed_product_contents_proto_rawDesc = nil
	file_sealed_sealed_product_contents_proto_goTypes = nil
	file_sealed_sealed_product_contents_proto_depIdxs = nil
}