// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.12.4
// source: booster/booster_config.proto

package booster

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

type BoosterConfig struct {
	state               protoimpl.MessageState   `protogen:"open.v1"`
	Boosters            []*BoosterPack           `protobuf:"bytes,1,rep,name=boosters,proto3" json:"boosters,omitempty" bson:"boosters"`                                                                       // @gotags: bson:"boosters"
	BoostersTotalWeight int64                    `protobuf:"varint,2,opt,name=boostersTotalWeight,proto3" json:"boostersTotalWeight,omitempty" bson:"boostersTotalWeight"`                                                // @gotags: bson:"boostersTotalWeight"
	Name                string                   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty" bson:"name"`                                                                               // @gotags: bson:"name"
	Sheets              map[string]*BoosterSheet `protobuf:"bytes,4,rep,name=sheets,proto3" json:"sheets,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value" bson:"sheets"` // @gotags: bson:"sheets"
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
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

const file_booster_booster_config_proto_rawDesc = "" +
	"\n" +
	"\x1cbooster/booster_config.proto\x12\abooster\x1a\x15booster/booster.proto\x1a\x1bbooster/booster_sheet.proto\"\x95\x02\n" +
	"\rBoosterConfig\x120\n" +
	"\bboosters\x18\x01 \x03(\v2\x14.booster.BoosterPackR\bboosters\x120\n" +
	"\x13boostersTotalWeight\x18\x02 \x01(\x03R\x13boostersTotalWeight\x12\x12\n" +
	"\x04name\x18\x03 \x01(\tR\x04name\x12:\n" +
	"\x06sheets\x18\x04 \x03(\v2\".booster.BoosterConfig.SheetsEntryR\x06sheets\x1aP\n" +
	"\vSheetsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12+\n" +
	"\x05value\x18\x02 \x01(\v2\x15.booster.BoosterSheetR\x05value:\x028\x01B.Z,github.com/stevezaluk/mtgjson-models/boosterb\x06proto3"

var (
	file_booster_booster_config_proto_rawDescOnce sync.Once
	file_booster_booster_config_proto_rawDescData []byte
)

func file_booster_booster_config_proto_rawDescGZIP() []byte {
	file_booster_booster_config_proto_rawDescOnce.Do(func() {
		file_booster_booster_config_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_booster_booster_config_proto_rawDesc), len(file_booster_booster_config_proto_rawDesc)))
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
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_booster_booster_config_proto_rawDesc), len(file_booster_booster_config_proto_rawDesc)),
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
	file_booster_booster_config_proto_goTypes = nil
	file_booster_booster_config_proto_depIdxs = nil
}
