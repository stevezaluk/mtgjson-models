// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.12.4
// source: meta/legalities.proto

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

type CardLegalities struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	Alchemy         string                 `protobuf:"bytes,1,opt,name=alchemy,proto3" json:"alchemy,omitempty" bson:"alchemy"`                  // @gotags: bson:"alchemy"
	Brawl           string                 `protobuf:"bytes,2,opt,name=brawl,proto3" json:"brawl,omitempty" bson:"brawl"`                      // @gotags: bson:"brawl"
	Commander       string                 `protobuf:"bytes,3,opt,name=commander,proto3" json:"commander,omitempty" bson:"commander"`              // @gotags: bson:"commander"
	Duel            string                 `protobuf:"bytes,4,opt,name=duel,proto3" json:"duel,omitempty" bson:"duel"`                        // @gotags: bson:"duel"
	Explorer        string                 `protobuf:"bytes,5,opt,name=explorer,proto3" json:"explorer,omitempty" bson:"explorer"`                // @gotags: bson:"explorer"
	Future          string                 `protobuf:"bytes,6,opt,name=future,proto3" json:"future,omitempty" bson:"future"`                    // @gotags: bson:"future"
	Gladiator       string                 `protobuf:"bytes,7,opt,name=gladiator,proto3" json:"gladiator,omitempty" bson:"gladiator"`              // @gotags: bson:"gladiator"
	Historic        string                 `protobuf:"bytes,8,opt,name=historic,proto3" json:"historic,omitempty" bson:"historic"`                // @gotags: bson:"historic"
	HistoricBrawl   string                 `protobuf:"bytes,9,opt,name=historicBrawl,proto3" json:"historicBrawl,omitempty" bson:"historicBrawl"`      // @gotags: bson:"historicBrawl"
	Legacy          string                 `protobuf:"bytes,10,opt,name=legacy,proto3" json:"legacy,omitempty" bson:"legacy"`                   // @gotags: bson:"legacy"
	Modern          string                 `protobuf:"bytes,11,opt,name=modern,proto3" json:"modern,omitempty" bson:"modern"`                   // @gotags: bson:"modern"
	Oathbreaker     string                 `protobuf:"bytes,12,opt,name=oathbreaker,proto3" json:"oathbreaker,omitempty" bson:"oathbreaker"`         // @gotags: bson:"oathbreaker"
	Oldschool       string                 `protobuf:"bytes,13,opt,name=oldschool,proto3" json:"oldschool,omitempty" bson:"oldschool"`             // @gotags: bson:"oldschool"
	Pauper          string                 `protobuf:"bytes,14,opt,name=pauper,proto3" json:"pauper,omitempty" bson:"pauper"`                   // @gotags: bson:"pauper"
	Paupercommander string                 `protobuf:"bytes,15,opt,name=paupercommander,proto3" json:"paupercommander,omitempty" bson:"paupercommander"` // @gotags: bson:"paupercommander"
	Penny           string                 `protobuf:"bytes,16,opt,name=penny,proto3" json:"penny,omitempty" bson:"penny"`                     // @gotags: bson:"penny"
	Pioneer         string                 `protobuf:"bytes,17,opt,name=pioneer,proto3" json:"pioneer,omitempty" bson:"pioneer"`                 // @gotags: bson:"pioneer"
	Predh           string                 `protobuf:"bytes,18,opt,name=predh,proto3" json:"predh,omitempty" bson:"predh"`                     // @gotags: bson:"predh"
	Premodern       string                 `protobuf:"bytes,19,opt,name=premodern,proto3" json:"premodern,omitempty" bson:"premodern"`             // @gotags: bson:"premodern"
	Standard        string                 `protobuf:"bytes,20,opt,name=standard,proto3" json:"standard,omitempty" bson:"standard"`               // @gotags: bson:"standard"
	Standardbrawl   string                 `protobuf:"bytes,21,opt,name=standardbrawl,proto3" json:"standardbrawl,omitempty" bson:"standardbrawl"`     // @gotags: bson:"standardbrawl"
	Timeless        string                 `protobuf:"bytes,22,opt,name=timeless,proto3" json:"timeless,omitempty" bson:"timeless"`               // @gotags: bson:"timeless"
	Vintage         string                 `protobuf:"bytes,23,opt,name=vintage,proto3" json:"vintage,omitempty" bson:"vintage"`                 // @gotags: bson:"vintage"
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *CardLegalities) Reset() {
	*x = CardLegalities{}
	mi := &file_meta_legalities_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CardLegalities) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CardLegalities) ProtoMessage() {}

func (x *CardLegalities) ProtoReflect() protoreflect.Message {
	mi := &file_meta_legalities_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CardLegalities.ProtoReflect.Descriptor instead.
func (*CardLegalities) Descriptor() ([]byte, []int) {
	return file_meta_legalities_proto_rawDescGZIP(), []int{0}
}

func (x *CardLegalities) GetAlchemy() string {
	if x != nil {
		return x.Alchemy
	}
	return ""
}

func (x *CardLegalities) GetBrawl() string {
	if x != nil {
		return x.Brawl
	}
	return ""
}

func (x *CardLegalities) GetCommander() string {
	if x != nil {
		return x.Commander
	}
	return ""
}

func (x *CardLegalities) GetDuel() string {
	if x != nil {
		return x.Duel
	}
	return ""
}

func (x *CardLegalities) GetExplorer() string {
	if x != nil {
		return x.Explorer
	}
	return ""
}

func (x *CardLegalities) GetFuture() string {
	if x != nil {
		return x.Future
	}
	return ""
}

func (x *CardLegalities) GetGladiator() string {
	if x != nil {
		return x.Gladiator
	}
	return ""
}

func (x *CardLegalities) GetHistoric() string {
	if x != nil {
		return x.Historic
	}
	return ""
}

func (x *CardLegalities) GetHistoricBrawl() string {
	if x != nil {
		return x.HistoricBrawl
	}
	return ""
}

func (x *CardLegalities) GetLegacy() string {
	if x != nil {
		return x.Legacy
	}
	return ""
}

func (x *CardLegalities) GetModern() string {
	if x != nil {
		return x.Modern
	}
	return ""
}

func (x *CardLegalities) GetOathbreaker() string {
	if x != nil {
		return x.Oathbreaker
	}
	return ""
}

func (x *CardLegalities) GetOldschool() string {
	if x != nil {
		return x.Oldschool
	}
	return ""
}

func (x *CardLegalities) GetPauper() string {
	if x != nil {
		return x.Pauper
	}
	return ""
}

func (x *CardLegalities) GetPaupercommander() string {
	if x != nil {
		return x.Paupercommander
	}
	return ""
}

func (x *CardLegalities) GetPenny() string {
	if x != nil {
		return x.Penny
	}
	return ""
}

func (x *CardLegalities) GetPioneer() string {
	if x != nil {
		return x.Pioneer
	}
	return ""
}

func (x *CardLegalities) GetPredh() string {
	if x != nil {
		return x.Predh
	}
	return ""
}

func (x *CardLegalities) GetPremodern() string {
	if x != nil {
		return x.Premodern
	}
	return ""
}

func (x *CardLegalities) GetStandard() string {
	if x != nil {
		return x.Standard
	}
	return ""
}

func (x *CardLegalities) GetStandardbrawl() string {
	if x != nil {
		return x.Standardbrawl
	}
	return ""
}

func (x *CardLegalities) GetTimeless() string {
	if x != nil {
		return x.Timeless
	}
	return ""
}

func (x *CardLegalities) GetVintage() string {
	if x != nil {
		return x.Vintage
	}
	return ""
}

var File_meta_legalities_proto protoreflect.FileDescriptor

const file_meta_legalities_proto_rawDesc = "" +
	"\n" +
	"\x15meta/legalities.proto\x12\x04meta\"\x94\x05\n" +
	"\x0eCardLegalities\x12\x18\n" +
	"\aalchemy\x18\x01 \x01(\tR\aalchemy\x12\x14\n" +
	"\x05brawl\x18\x02 \x01(\tR\x05brawl\x12\x1c\n" +
	"\tcommander\x18\x03 \x01(\tR\tcommander\x12\x12\n" +
	"\x04duel\x18\x04 \x01(\tR\x04duel\x12\x1a\n" +
	"\bexplorer\x18\x05 \x01(\tR\bexplorer\x12\x16\n" +
	"\x06future\x18\x06 \x01(\tR\x06future\x12\x1c\n" +
	"\tgladiator\x18\a \x01(\tR\tgladiator\x12\x1a\n" +
	"\bhistoric\x18\b \x01(\tR\bhistoric\x12$\n" +
	"\rhistoricBrawl\x18\t \x01(\tR\rhistoricBrawl\x12\x16\n" +
	"\x06legacy\x18\n" +
	" \x01(\tR\x06legacy\x12\x16\n" +
	"\x06modern\x18\v \x01(\tR\x06modern\x12 \n" +
	"\voathbreaker\x18\f \x01(\tR\voathbreaker\x12\x1c\n" +
	"\toldschool\x18\r \x01(\tR\toldschool\x12\x16\n" +
	"\x06pauper\x18\x0e \x01(\tR\x06pauper\x12(\n" +
	"\x0fpaupercommander\x18\x0f \x01(\tR\x0fpaupercommander\x12\x14\n" +
	"\x05penny\x18\x10 \x01(\tR\x05penny\x12\x18\n" +
	"\apioneer\x18\x11 \x01(\tR\apioneer\x12\x14\n" +
	"\x05predh\x18\x12 \x01(\tR\x05predh\x12\x1c\n" +
	"\tpremodern\x18\x13 \x01(\tR\tpremodern\x12\x1a\n" +
	"\bstandard\x18\x14 \x01(\tR\bstandard\x12$\n" +
	"\rstandardbrawl\x18\x15 \x01(\tR\rstandardbrawl\x12\x1a\n" +
	"\btimeless\x18\x16 \x01(\tR\btimeless\x12\x18\n" +
	"\avintage\x18\x17 \x01(\tR\avintageB+Z)github.com/stevezaluk/mtgjson-models/metab\x06proto3"

var (
	file_meta_legalities_proto_rawDescOnce sync.Once
	file_meta_legalities_proto_rawDescData []byte
)

func file_meta_legalities_proto_rawDescGZIP() []byte {
	file_meta_legalities_proto_rawDescOnce.Do(func() {
		file_meta_legalities_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_meta_legalities_proto_rawDesc), len(file_meta_legalities_proto_rawDesc)))
	})
	return file_meta_legalities_proto_rawDescData
}

var file_meta_legalities_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_meta_legalities_proto_goTypes = []any{
	(*CardLegalities)(nil), // 0: meta.CardLegalities
}
var file_meta_legalities_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_meta_legalities_proto_init() }
func file_meta_legalities_proto_init() {
	if File_meta_legalities_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_meta_legalities_proto_rawDesc), len(file_meta_legalities_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_meta_legalities_proto_goTypes,
		DependencyIndexes: file_meta_legalities_proto_depIdxs,
		MessageInfos:      file_meta_legalities_proto_msgTypes,
	}.Build()
	File_meta_legalities_proto = out.File
	file_meta_legalities_proto_goTypes = nil
	file_meta_legalities_proto_depIdxs = nil
}
