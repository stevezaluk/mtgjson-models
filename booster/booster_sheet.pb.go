// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.6.1
// source: booster/booster_sheet.proto

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

type BoosterSheet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AllowDuplicates bool             `protobuf:"varint,1,opt,name=allowDuplicates,proto3" json:"allowDuplicates,omitempty" bson:"allowDuplicates"`                                                                     // @gotags: bson:"allowDuplicates"
	BalanceColors   bool             `protobuf:"varint,2,opt,name=balanceColors,proto3" json:"balanceColors,omitempty" bson:"balanceColors"`                                                                         // @gotags: bson:"balanceColors"
	Cards           map[string]int64 `protobuf:"bytes,3,rep,name=cards,proto3" json:"cards,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3" bson:"cards"` // @gotags: bson:"cards"
	Foil            bool             `protobuf:"varint,4,opt,name=foil,proto3" json:"foil,omitempty" bson:"foil"`                                                                                           // @gotags: bson:"foil"
	Fixed           bool             `protobuf:"varint,5,opt,name=fixed,proto3" json:"fixed,omitempty" bson:"fixed"`                                                                                         // @gotags: bson:"fixed"
	TotalWeight     int64            `protobuf:"varint,6,opt,name=totalWeight,proto3" json:"totalWeight,omitempty" bson:"totalWeight"`                                                                             // @gotags: bson:"totalWeight"
}

func (x *BoosterSheet) Reset() {
	*x = BoosterSheet{}
	mi := &file_booster_booster_sheet_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BoosterSheet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoosterSheet) ProtoMessage() {}

func (x *BoosterSheet) ProtoReflect() protoreflect.Message {
	mi := &file_booster_booster_sheet_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoosterSheet.ProtoReflect.Descriptor instead.
func (*BoosterSheet) Descriptor() ([]byte, []int) {
	return file_booster_booster_sheet_proto_rawDescGZIP(), []int{0}
}

func (x *BoosterSheet) GetAllowDuplicates() bool {
	if x != nil {
		return x.AllowDuplicates
	}
	return false
}

func (x *BoosterSheet) GetBalanceColors() bool {
	if x != nil {
		return x.BalanceColors
	}
	return false
}

func (x *BoosterSheet) GetCards() map[string]int64 {
	if x != nil {
		return x.Cards
	}
	return nil
}

func (x *BoosterSheet) GetFoil() bool {
	if x != nil {
		return x.Foil
	}
	return false
}

func (x *BoosterSheet) GetFixed() bool {
	if x != nil {
		return x.Fixed
	}
	return false
}

func (x *BoosterSheet) GetTotalWeight() int64 {
	if x != nil {
		return x.TotalWeight
	}
	return 0
}

var File_booster_booster_sheet_proto protoreflect.FileDescriptor

var file_booster_booster_sheet_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x62, 0x6f, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x62, 0x6f, 0x6f, 0x73, 0x74, 0x65,
	0x72, 0x5f, 0x73, 0x68, 0x65, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x62,
	0x6f, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x22, 0x9c, 0x02, 0x0a, 0x0c, 0x42, 0x6f, 0x6f, 0x73, 0x74,
	0x65, 0x72, 0x53, 0x68, 0x65, 0x65, 0x74, 0x12, 0x28, 0x0a, 0x0f, 0x61, 0x6c, 0x6c, 0x6f, 0x77,
	0x44, 0x75, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0f, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x44, 0x75, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x73, 0x12, 0x24, 0x0a, 0x0d, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x43, 0x6f, 0x6c, 0x6f,
	0x72, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x73, 0x12, 0x36, 0x0a, 0x05, 0x63, 0x61, 0x72, 0x64, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x62, 0x6f, 0x6f, 0x73, 0x74, 0x65, 0x72,
	0x2e, 0x42, 0x6f, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x53, 0x68, 0x65, 0x65, 0x74, 0x2e, 0x43, 0x61,
	0x72, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x63, 0x61, 0x72, 0x64, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x66, 0x6f, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x66,
	0x6f, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x78, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x05, 0x66, 0x69, 0x78, 0x65, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x1a, 0x38, 0x0a, 0x0a, 0x43,
	0x61, 0x72, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x65, 0x76, 0x65, 0x7a, 0x61, 0x6c, 0x75, 0x6b, 0x2f, 0x6d,
	0x74, 0x67, 0x6a, 0x73, 0x6f, 0x6e, 0x2d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x62, 0x6f,
	0x6f, 0x73, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_booster_booster_sheet_proto_rawDescOnce sync.Once
	file_booster_booster_sheet_proto_rawDescData = file_booster_booster_sheet_proto_rawDesc
)

func file_booster_booster_sheet_proto_rawDescGZIP() []byte {
	file_booster_booster_sheet_proto_rawDescOnce.Do(func() {
		file_booster_booster_sheet_proto_rawDescData = protoimpl.X.CompressGZIP(file_booster_booster_sheet_proto_rawDescData)
	})
	return file_booster_booster_sheet_proto_rawDescData
}

var file_booster_booster_sheet_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_booster_booster_sheet_proto_goTypes = []any{
	(*BoosterSheet)(nil), // 0: booster.BoosterSheet
	nil,                  // 1: booster.BoosterSheet.CardsEntry
}
var file_booster_booster_sheet_proto_depIdxs = []int32{
	1, // 0: booster.BoosterSheet.cards:type_name -> booster.BoosterSheet.CardsEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_booster_booster_sheet_proto_init() }
func file_booster_booster_sheet_proto_init() {
	if File_booster_booster_sheet_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_booster_booster_sheet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_booster_booster_sheet_proto_goTypes,
		DependencyIndexes: file_booster_booster_sheet_proto_depIdxs,
		MessageInfos:      file_booster_booster_sheet_proto_msgTypes,
	}.Build()
	File_booster_booster_sheet_proto = out.File
	file_booster_booster_sheet_proto_rawDesc = nil
	file_booster_booster_sheet_proto_goTypes = nil
	file_booster_booster_sheet_proto_depIdxs = nil
}
