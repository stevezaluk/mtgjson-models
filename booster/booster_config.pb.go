// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.6.1
// source: booster/booster_config.proto

package booster

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

type BoosterConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Boosters            []*BoosterPack           `protobuf:"bytes,1,rep,name=boosters,proto3" json:"boosters,omitempty" bson:"boosters"`                                                                                     // @gotags: bson:"boosters"
	BoostersTotalWeight int64                    `protobuf:"varint,2,opt,name=boostersTotalWeight,proto3" json:"boostersTotalWeight,omitempty" bson:"boostersTotalWeight"`                                                              // @gotags: bson:"boostersTotalWeight"
	Name                string                   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty" bson:"name"`                                                                                             // @gotags: bson:"name"
	Sheets              map[string]*BoosterSheet `protobuf:"bytes,4,rep,name=sheets,proto3" json:"sheets,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" bson:"sheets"` // @gotags: bson:"sheets"
}

func (x *BoosterConfig) Reset() {
	*x = BoosterConfig{}
	mi := &file_booster_booster_config_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BoosterConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoosterConfig) ProtoMessage() {}

func (x *BoosterConfig) ProtoReflect() protoreflect.Message {
	mi := &file_booster_booster_config_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoosterConfig.ProtoReflect.Descriptor instead.
func (*BoosterConfig) Descriptor() ([]byte, []int) {
	return file_booster_booster_config_proto_rawDescGZIP(), []int{0}
}

func (x *BoosterConfig) GetBoosters() []*BoosterPack {
	if x != nil {
		return x.Boosters
	}
	return nil
}

func (x *BoosterConfig) GetBoostersTotalWeight() int64 {
	if x != nil {
		return x.BoostersTotalWeight
	}
	return 0
}

func (x *BoosterConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BoosterConfig) GetSheets() map[string]*BoosterSheet {
	if x != nil {
		return x.Sheets
	}
	return nil
}

var File_booster_booster_config_proto protoreflect.FileDescriptor

var file_booster_booster_config_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x62, 0x6f, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x62, 0x6f, 0x6f, 0x73, 0x74, 0x65,
	0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07,
	0x62, 0x6f, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x1a, 0x15, 0x62, 0x6f, 0x6f, 0x73, 0x74, 0x65, 0x72,
	0x2f, 0x62, 0x6f, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x62, 0x6f, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x62, 0x6f, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x5f,
	0x73, 0x68, 0x65, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x95, 0x02, 0x0a, 0x0d,
	0x42, 0x6f, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x30, 0x0a,
	0x08, 0x62, 0x6f, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x62, 0x6f, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x42, 0x6f, 0x6f, 0x73, 0x74, 0x65,
	0x72, 0x50, 0x61, 0x63, 0x6b, 0x52, 0x08, 0x62, 0x6f, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x73, 0x12,
	0x30, 0x0a, 0x13, 0x62, 0x6f, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x73, 0x54, 0x6f, 0x74, 0x61, 0x6c,
	0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x13, 0x62, 0x6f,
	0x6f, 0x73, 0x74, 0x65, 0x72, 0x73, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x57, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x06, 0x73, 0x68, 0x65, 0x65, 0x74, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x62, 0x6f, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x2e,
	0x42, 0x6f, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x53, 0x68,
	0x65, 0x65, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x73, 0x68, 0x65, 0x65, 0x74,
	0x73, 0x1a, 0x50, 0x0a, 0x0b, 0x53, 0x68, 0x65, 0x65, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x2b, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x62, 0x6f, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x42, 0x6f, 0x6f, 0x73,
	0x74, 0x65, 0x72, 0x53, 0x68, 0x65, 0x65, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x74, 0x65, 0x76, 0x65, 0x7a, 0x61, 0x6c, 0x75, 0x6b, 0x2f, 0x6d, 0x74, 0x67,
	0x6a, 0x73, 0x6f, 0x6e, 0x2d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x62, 0x6f, 0x6f, 0x73,
	0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_booster_booster_config_proto_rawDescOnce sync.Once
	file_booster_booster_config_proto_rawDescData = file_booster_booster_config_proto_rawDesc
)

func file_booster_booster_config_proto_rawDescGZIP() []byte {
	file_booster_booster_config_proto_rawDescOnce.Do(func() {
		file_booster_booster_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_booster_booster_config_proto_rawDescData)
	})
	return file_booster_booster_config_proto_rawDescData
}

var file_booster_booster_config_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_booster_booster_config_proto_goTypes = []any{
	(*BoosterConfig)(nil), // 0: booster.BoosterConfig
	nil,                   // 1: booster.BoosterConfig.SheetsEntry
	(*BoosterPack)(nil),   // 2: booster.BoosterPack
	(*BoosterSheet)(nil),  // 3: booster.BoosterSheet
}
var file_booster_booster_config_proto_depIdxs = []int32{
	2, // 0: booster.BoosterConfig.boosters:type_name -> booster.BoosterPack
	1, // 1: booster.BoosterConfig.sheets:type_name -> booster.BoosterConfig.SheetsEntry
	3, // 2: booster.BoosterConfig.SheetsEntry.value:type_name -> booster.BoosterSheet
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_booster_booster_config_proto_init() }
func file_booster_booster_config_proto_init() {
	if File_booster_booster_config_proto != nil {
		return
	}
	file_booster_booster_proto_init()
	file_booster_booster_sheet_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_booster_booster_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_booster_booster_config_proto_goTypes,
		DependencyIndexes: file_booster_booster_config_proto_depIdxs,
		MessageInfos:      file_booster_booster_config_proto_msgTypes,
	}.Build()
	File_booster_booster_config_proto = out.File
	file_booster_booster_config_proto_rawDesc = nil
	file_booster_booster_config_proto_goTypes = nil
	file_booster_booster_config_proto_depIdxs = nil
}
