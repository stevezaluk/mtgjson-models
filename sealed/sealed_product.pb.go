// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.6.1
// source: sealed/sealed_product.proto

package sealed

import (
	meta "github.com/stevezaluk/mtgjson-models/meta"
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

type SealedProduct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CardCount    int64                  `protobuf:"varint,1,opt,name=cardCount,proto3" json:"cardCount,omitempty"`
	Category     string                 `protobuf:"bytes,2,opt,name=category,proto3" json:"category,omitempty"`
	Contents     *SealedProductContents `protobuf:"bytes,3,opt,name=contents,proto3" json:"contents,omitempty"`
	Identifiers  *meta.CardIdentifiers  `protobuf:"bytes,4,opt,name=identifiers,proto3" json:"identifiers,omitempty"`
	Name         string                 `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	ProductSize  int64                  `protobuf:"varint,6,opt,name=productSize,proto3" json:"productSize,omitempty"`
	PurchaseUrls *meta.PurchaseUrls     `protobuf:"bytes,7,opt,name=purchaseUrls,proto3" json:"purchaseUrls,omitempty"`
	ReleaseDate  string                 `protobuf:"bytes,8,opt,name=releaseDate,proto3" json:"releaseDate,omitempty"`
	Subtype      string                 `protobuf:"bytes,9,opt,name=subtype,proto3" json:"subtype,omitempty"`
	Uuid         string                 `protobuf:"bytes,10,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *SealedProduct) Reset() {
	*x = SealedProduct{}
	mi := &file_sealed_sealed_product_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SealedProduct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SealedProduct) ProtoMessage() {}

func (x *SealedProduct) ProtoReflect() protoreflect.Message {
	mi := &file_sealed_sealed_product_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SealedProduct.ProtoReflect.Descriptor instead.
func (*SealedProduct) Descriptor() ([]byte, []int) {
	return file_sealed_sealed_product_proto_rawDescGZIP(), []int{0}
}

func (x *SealedProduct) GetCardCount() int64 {
	if x != nil {
		return x.CardCount
	}
	return 0
}

func (x *SealedProduct) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *SealedProduct) GetContents() *SealedProductContents {
	if x != nil {
		return x.Contents
	}
	return nil
}

func (x *SealedProduct) GetIdentifiers() *meta.CardIdentifiers {
	if x != nil {
		return x.Identifiers
	}
	return nil
}

func (x *SealedProduct) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SealedProduct) GetProductSize() int64 {
	if x != nil {
		return x.ProductSize
	}
	return 0
}

func (x *SealedProduct) GetPurchaseUrls() *meta.PurchaseUrls {
	if x != nil {
		return x.PurchaseUrls
	}
	return nil
}

func (x *SealedProduct) GetReleaseDate() string {
	if x != nil {
		return x.ReleaseDate
	}
	return ""
}

func (x *SealedProduct) GetSubtype() string {
	if x != nil {
		return x.Subtype
	}
	return ""
}

func (x *SealedProduct) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

var File_sealed_sealed_product_proto protoreflect.FileDescriptor

var file_sealed_sealed_product_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x73, 0x65, 0x61, 0x6c, 0x65, 0x64, 0x2f, 0x73, 0x65, 0x61, 0x6c, 0x65, 0x64, 0x5f,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73,
	0x65, 0x61, 0x6c, 0x65, 0x64, 0x1a, 0x24, 0x73, 0x65, 0x61, 0x6c, 0x65, 0x64, 0x2f, 0x73, 0x65,
	0x61, 0x6c, 0x65, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x6d, 0x65, 0x74,
	0x61, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61,
	0x73, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfb, 0x02,
	0x0a, 0x0d, 0x53, 0x65, 0x61, 0x6c, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x63, 0x61, 0x72, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x63, 0x61, 0x72, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x39, 0x0a, 0x08, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x73, 0x65,
	0x61, 0x6c, 0x65, 0x64, 0x2e, 0x53, 0x65, 0x61, 0x6c, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x73, 0x12, 0x37, 0x0a, 0x0b, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6d, 0x65, 0x74, 0x61,
	0x2e, 0x43, 0x61, 0x72, 0x64, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73,
	0x52, 0x0b, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x69, 0x7a, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x36, 0x0a, 0x0c, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x55,
	0x72, 0x6c, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x65, 0x74, 0x61,
	0x2e, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x55, 0x72, 0x6c, 0x73, 0x52, 0x0c, 0x70,
	0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x55, 0x72, 0x6c, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x72,
	0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x44, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x44, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x75, 0x62, 0x74, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x73, 0x75, 0x62, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x42, 0x2d, 0x5a, 0x2b, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x65, 0x76, 0x65, 0x7a,
	0x61, 0x6c, 0x75, 0x6b, 0x2f, 0x6d, 0x74, 0x67, 0x6a, 0x73, 0x6f, 0x6e, 0x2d, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x2f, 0x73, 0x65, 0x61, 0x6c, 0x65, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_sealed_sealed_product_proto_rawDescOnce sync.Once
	file_sealed_sealed_product_proto_rawDescData = file_sealed_sealed_product_proto_rawDesc
)

func file_sealed_sealed_product_proto_rawDescGZIP() []byte {
	file_sealed_sealed_product_proto_rawDescOnce.Do(func() {
		file_sealed_sealed_product_proto_rawDescData = protoimpl.X.CompressGZIP(file_sealed_sealed_product_proto_rawDescData)
	})
	return file_sealed_sealed_product_proto_rawDescData
}

var file_sealed_sealed_product_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_sealed_sealed_product_proto_goTypes = []any{
	(*SealedProduct)(nil),         // 0: sealed.SealedProduct
	(*SealedProductContents)(nil), // 1: sealed.SealedProductContents
	(*meta.CardIdentifiers)(nil),  // 2: meta.CardIdentifiers
	(*meta.PurchaseUrls)(nil),     // 3: meta.PurchaseUrls
}
var file_sealed_sealed_product_proto_depIdxs = []int32{
	1, // 0: sealed.SealedProduct.contents:type_name -> sealed.SealedProductContents
	2, // 1: sealed.SealedProduct.identifiers:type_name -> meta.CardIdentifiers
	3, // 2: sealed.SealedProduct.purchaseUrls:type_name -> meta.PurchaseUrls
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_sealed_sealed_product_proto_init() }
func file_sealed_sealed_product_proto_init() {
	if File_sealed_sealed_product_proto != nil {
		return
	}
	file_sealed_sealed_product_contents_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sealed_sealed_product_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sealed_sealed_product_proto_goTypes,
		DependencyIndexes: file_sealed_sealed_product_proto_depIdxs,
		MessageInfos:      file_sealed_sealed_product_proto_msgTypes,
	}.Build()
	File_sealed_sealed_product_proto = out.File
	file_sealed_sealed_product_proto_rawDesc = nil
	file_sealed_sealed_product_proto_goTypes = nil
	file_sealed_sealed_product_proto_depIdxs = nil
}
