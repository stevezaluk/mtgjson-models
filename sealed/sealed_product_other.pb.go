// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: sealed/sealed_product_other.proto

package sealed

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SealedProductOther struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" bson:"name"` // @gotags: bson:"name"
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SealedProductOther) Reset() {
	*x = SealedProductOther{}
	mi := &file_sealed_sealed_product_other_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SealedProductOther) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SealedProductOther) ProtoMessage() {}

func (x *SealedProductOther) ProtoReflect() protoreflect.Message {
	mi := &file_sealed_sealed_product_other_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SealedProductOther.ProtoReflect.Descriptor instead.
func (*SealedProductOther) Descriptor() ([]byte, []int) {
	return file_sealed_sealed_product_other_proto_rawDescGZIP(), []int{0}
}

func (x *SealedProductOther) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_sealed_sealed_product_other_proto protoreflect.FileDescriptor

const file_sealed_sealed_product_other_proto_rawDesc = "" +
	"\n" +
	"!sealed/sealed_product_other.proto\x12\x06sealed\"(\n" +
	"\x12SealedProductOther\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04nameB-Z+github.com/stevezaluk/mtgjson-models/sealedb\x06proto3"

var (
	file_sealed_sealed_product_other_proto_rawDescOnce sync.Once
	file_sealed_sealed_product_other_proto_rawDescData []byte
)

func file_sealed_sealed_product_other_proto_rawDescGZIP() []byte {
	file_sealed_sealed_product_other_proto_rawDescOnce.Do(func() {
		file_sealed_sealed_product_other_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_sealed_sealed_product_other_proto_rawDesc), len(file_sealed_sealed_product_other_proto_rawDesc)))
	})
	return file_sealed_sealed_product_other_proto_rawDescData
}

var file_sealed_sealed_product_other_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_sealed_sealed_product_other_proto_goTypes = []any{
	(*SealedProductOther)(nil), // 0: sealed.SealedProductOther
}
var file_sealed_sealed_product_other_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sealed_sealed_product_other_proto_init() }
func file_sealed_sealed_product_other_proto_init() {
	if File_sealed_sealed_product_other_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_sealed_sealed_product_other_proto_rawDesc), len(file_sealed_sealed_product_other_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sealed_sealed_product_other_proto_goTypes,
		DependencyIndexes: file_sealed_sealed_product_other_proto_depIdxs,
		MessageInfos:      file_sealed_sealed_product_other_proto_msgTypes,
	}.Build()
	File_sealed_sealed_product_other_proto = out.File
	file_sealed_sealed_product_other_proto_goTypes = nil
	file_sealed_sealed_product_other_proto_depIdxs = nil
}
