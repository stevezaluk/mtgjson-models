// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.12.4
// source: meta/related_cards.proto

package meta

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

type RelatedCards struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	ReverseRelated []string               `protobuf:"bytes,1,rep,name=reverseRelated,proto3" json:"reverseRelated,omitempty" bson:"reverseRelated"` // @gotags: bson:"reverseRelated"
	Spellbook      []string               `protobuf:"bytes,2,rep,name=spellbook,proto3" json:"spellbook,omitempty" bson:"spellbook"`           // @gotags: bson:"spellbook"
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *RelatedCards) Reset() {
	*x = RelatedCards{}
	mi := &file_meta_related_cards_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RelatedCards) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelatedCards) ProtoMessage() {}

func (x *RelatedCards) ProtoReflect() protoreflect.Message {
	mi := &file_meta_related_cards_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelatedCards.ProtoReflect.Descriptor instead.
func (*RelatedCards) Descriptor() ([]byte, []int) {
	return file_meta_related_cards_proto_rawDescGZIP(), []int{0}
}

func (x *RelatedCards) GetReverseRelated() []string {
	if x != nil {
		return x.ReverseRelated
	}
	return nil
}

func (x *RelatedCards) GetSpellbook() []string {
	if x != nil {
		return x.Spellbook
	}
	return nil
}

var File_meta_related_cards_proto protoreflect.FileDescriptor

const file_meta_related_cards_proto_rawDesc = "" +
	"\n" +
	"\x18meta/related_cards.proto\x12\x04meta\"T\n" +
	"\fRelatedCards\x12&\n" +
	"\x0ereverseRelated\x18\x01 \x03(\tR\x0ereverseRelated\x12\x1c\n" +
	"\tspellbook\x18\x02 \x03(\tR\tspellbookB+Z)github.com/stevezaluk/mtgjson-models/metab\x06proto3"

var (
	file_meta_related_cards_proto_rawDescOnce sync.Once
	file_meta_related_cards_proto_rawDescData []byte
)

func file_meta_related_cards_proto_rawDescGZIP() []byte {
	file_meta_related_cards_proto_rawDescOnce.Do(func() {
		file_meta_related_cards_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_meta_related_cards_proto_rawDesc), len(file_meta_related_cards_proto_rawDesc)))
	})
	return file_meta_related_cards_proto_rawDescData
}

var file_meta_related_cards_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_meta_related_cards_proto_goTypes = []any{
	(*RelatedCards)(nil), // 0: meta.RelatedCards
}
var file_meta_related_cards_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_meta_related_cards_proto_init() }
func file_meta_related_cards_proto_init() {
	if File_meta_related_cards_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_meta_related_cards_proto_rawDesc), len(file_meta_related_cards_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_meta_related_cards_proto_goTypes,
		DependencyIndexes: file_meta_related_cards_proto_depIdxs,
		MessageInfos:      file_meta_related_cards_proto_msgTypes,
	}.Build()
	File_meta_related_cards_proto = out.File
	file_meta_related_cards_proto_goTypes = nil
	file_meta_related_cards_proto_depIdxs = nil
}
