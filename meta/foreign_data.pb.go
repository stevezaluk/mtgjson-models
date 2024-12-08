// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.6.1
// source: meta/foreign_data.proto

package meta

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

type ForeignData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FaceName    string           `protobuf:"bytes,1,opt,name=faceName,proto3" json:"faceName,omitempty" bson:"faceName"`       // @gotags: bson:"faceName"
	FlavorText  string           `protobuf:"bytes,2,opt,name=flavorText,proto3" json:"flavorText,omitempty" bson:"flavorText"`   // @gotags: bson:"flavorText"
	Identifiers *CardIdentifiers `protobuf:"bytes,3,opt,name=identifiers,proto3" json:"identifiers,omitempty" bson:"identifiers"` // @gotags: bson:"identifiers"
	Language    string           `protobuf:"bytes,4,opt,name=language,proto3" json:"language,omitempty" bson:"language"`       // @gotags: bson:"language"
	Name        string           `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty" bson:"name"`               // @gotags: bson:"name"
	Text        string           `protobuf:"bytes,6,opt,name=text,proto3" json:"text,omitempty" bson:"text"`               // @gotags: bson:"text"
	Type        string           `protobuf:"bytes,7,opt,name=type,proto3" json:"type,omitempty" bson:"type"`               // @gotags: bson:"type"
}

func (x *ForeignData) Reset() {
	*x = ForeignData{}
	mi := &file_meta_foreign_data_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ForeignData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForeignData) ProtoMessage() {}

func (x *ForeignData) ProtoReflect() protoreflect.Message {
	mi := &file_meta_foreign_data_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForeignData.ProtoReflect.Descriptor instead.
func (*ForeignData) Descriptor() ([]byte, []int) {
	return file_meta_foreign_data_proto_rawDescGZIP(), []int{0}
}

func (x *ForeignData) GetFaceName() string {
	if x != nil {
		return x.FaceName
	}
	return ""
}

func (x *ForeignData) GetFlavorText() string {
	if x != nil {
		return x.FlavorText
	}
	return ""
}

func (x *ForeignData) GetIdentifiers() *CardIdentifiers {
	if x != nil {
		return x.Identifiers
	}
	return nil
}

func (x *ForeignData) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *ForeignData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ForeignData) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *ForeignData) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

var File_meta_foreign_data_proto protoreflect.FileDescriptor

var file_meta_foreign_data_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x66, 0x6f, 0x72, 0x65, 0x69, 0x67, 0x6e, 0x5f, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x1a,
	0x16, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xda, 0x01, 0x0a, 0x0b, 0x46, 0x6f, 0x72, 0x65,
	0x69, 0x67, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x61, 0x63, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x61, 0x63, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x6c, 0x61, 0x76, 0x6f, 0x72, 0x54, 0x65, 0x78,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x6c, 0x61, 0x76, 0x6f, 0x72, 0x54,
	0x65, 0x78, 0x74, 0x12, 0x37, 0x0a, 0x0b, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e,
	0x43, 0x61, 0x72, 0x64, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x52,
	0x0b, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x12, 0x1a, 0x0a, 0x08,
	0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x65, 0x76, 0x65, 0x7a, 0x61, 0x6c, 0x75, 0x6b, 0x2f, 0x6d, 0x74,
	0x67, 0x6a, 0x73, 0x6f, 0x6e, 0x2d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x65, 0x74,
	0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_meta_foreign_data_proto_rawDescOnce sync.Once
	file_meta_foreign_data_proto_rawDescData = file_meta_foreign_data_proto_rawDesc
)

func file_meta_foreign_data_proto_rawDescGZIP() []byte {
	file_meta_foreign_data_proto_rawDescOnce.Do(func() {
		file_meta_foreign_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_meta_foreign_data_proto_rawDescData)
	})
	return file_meta_foreign_data_proto_rawDescData
}

var file_meta_foreign_data_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_meta_foreign_data_proto_goTypes = []any{
	(*ForeignData)(nil),     // 0: meta.ForeignData
	(*CardIdentifiers)(nil), // 1: meta.CardIdentifiers
}
var file_meta_foreign_data_proto_depIdxs = []int32{
	1, // 0: meta.ForeignData.identifiers:type_name -> meta.CardIdentifiers
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_meta_foreign_data_proto_init() }
func file_meta_foreign_data_proto_init() {
	if File_meta_foreign_data_proto != nil {
		return
	}
	file_meta_identifiers_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_meta_foreign_data_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_meta_foreign_data_proto_goTypes,
		DependencyIndexes: file_meta_foreign_data_proto_depIdxs,
		MessageInfos:      file_meta_foreign_data_proto_msgTypes,
	}.Build()
	File_meta_foreign_data_proto = out.File
	file_meta_foreign_data_proto_rawDesc = nil
	file_meta_foreign_data_proto_goTypes = nil
	file_meta_foreign_data_proto_depIdxs = nil
}
