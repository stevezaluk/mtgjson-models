// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.12.4
// source: meta/translations.proto

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

type Translations struct {
	state              protoimpl.MessageState `protogen:"open.v1"`
	AncientGreek       string                 `protobuf:"bytes,1,opt,name=ancientGreek,json=Ancient Greek,proto3" json:"ancientGreek,omitempty" bson:"ancientGreek"`                   // @gotags: bson:"ancientGreek"
	Arabic             string                 `protobuf:"bytes,2,opt,name=arabic,json=Arabic,proto3" json:"arabic,omitempty" bson:"arabic"`                                      // @gotags: bson:"arabic"
	ChineseSimplified  string                 `protobuf:"bytes,3,opt,name=chineseSimplified,json=Chinese Simplified,proto3" json:"chineseSimplified,omitempty" bson:"chineseSimplified"`    // @gotags: bson:"chineseSimplified"
	ChineseTraditional string                 `protobuf:"bytes,4,opt,name=chineseTraditional,json=Chinese Traditional,proto3" json:"chineseTraditional,omitempty" bson:"chineseTraditional"` // @gotags: bson:"chineseTraditional"
	French             string                 `protobuf:"bytes,5,opt,name=french,json=French,proto3" json:"french,omitempty" bson:"french"`                                      // @gotags: bson:"french"
	German             string                 `protobuf:"bytes,6,opt,name=german,json=German,proto3" json:"german,omitempty" bson:"german"`                                      // @gotags: bson:"german"
	Hebrew             string                 `protobuf:"bytes,7,opt,name=hebrew,json=Hebrew,proto3" json:"hebrew,omitempty" bson:"hebrew"`                                      // @gotags: bson:"hebrew"
	Italian            string                 `protobuf:"bytes,8,opt,name=italian,json=Italian,proto3" json:"italian,omitempty" bson:"italian"`                                   // @gotags: bson:"italian"
	Japanese           string                 `protobuf:"bytes,9,opt,name=japanese,json=Japanese,proto3" json:"japanese,omitempty" bson:"japanese"`                                // @gotags: bson:"japanese"
	Korean             string                 `protobuf:"bytes,10,opt,name=korean,json=Korean,proto3" json:"korean,omitempty" bson:"korean"`                                     // @gotags: bson:"korean"
	Latin              string                 `protobuf:"bytes,11,opt,name=latin,json=Latin,proto3" json:"latin,omitempty" bson:"latin"`                                        // @gotags: bson:"latin"
	Phyrexian          string                 `protobuf:"bytes,12,opt,name=phyrexian,json=Phyrexian,proto3" json:"phyrexian,omitempty" bson:"phyrexian"`                            // @gotags: bson:"phyrexian"
	PortugeseBrazil    string                 `protobuf:"bytes,13,opt,name=portugeseBrazil,json=Portugese (Brazil),proto3" json:"portugeseBrazil,omitempty" bson:"portugeseBrazil"`       // @gotags: bson:"portugeseBrazil"
	Russian            string                 `protobuf:"bytes,14,opt,name=russian,json=Russian,proto3" json:"russian,omitempty" bson:"russian"`                                  // @gotags: bson:"russian"
	Sanskrit           string                 `protobuf:"bytes,15,opt,name=sanskrit,json=Sanskrit,proto3" json:"sanskrit,omitempty" bson:"sanskrit"`                               // @gotags: bson:"sanskrit"
	Spanish            string                 `protobuf:"bytes,16,opt,name=spanish,json=Spanish,proto3" json:"spanish,omitempty" bson:"spanish"`                                  // @gotags: bson:"spanish"
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *Translations) Reset() {
	*x = Translations{}
	mi := &file_meta_translations_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Translations) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Translations) ProtoMessage() {}

func (x *Translations) ProtoReflect() protoreflect.Message {
	mi := &file_meta_translations_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Translations.ProtoReflect.Descriptor instead.
func (*Translations) Descriptor() ([]byte, []int) {
	return file_meta_translations_proto_rawDescGZIP(), []int{0}
}

func (x *Translations) GetAncientGreek() string {
	if x != nil {
		return x.AncientGreek
	}
	return ""
}

func (x *Translations) GetArabic() string {
	if x != nil {
		return x.Arabic
	}
	return ""
}

func (x *Translations) GetChineseSimplified() string {
	if x != nil {
		return x.ChineseSimplified
	}
	return ""
}

func (x *Translations) GetChineseTraditional() string {
	if x != nil {
		return x.ChineseTraditional
	}
	return ""
}

func (x *Translations) GetFrench() string {
	if x != nil {
		return x.French
	}
	return ""
}

func (x *Translations) GetGerman() string {
	if x != nil {
		return x.German
	}
	return ""
}

func (x *Translations) GetHebrew() string {
	if x != nil {
		return x.Hebrew
	}
	return ""
}

func (x *Translations) GetItalian() string {
	if x != nil {
		return x.Italian
	}
	return ""
}

func (x *Translations) GetJapanese() string {
	if x != nil {
		return x.Japanese
	}
	return ""
}

func (x *Translations) GetKorean() string {
	if x != nil {
		return x.Korean
	}
	return ""
}

func (x *Translations) GetLatin() string {
	if x != nil {
		return x.Latin
	}
	return ""
}

func (x *Translations) GetPhyrexian() string {
	if x != nil {
		return x.Phyrexian
	}
	return ""
}

func (x *Translations) GetPortugeseBrazil() string {
	if x != nil {
		return x.PortugeseBrazil
	}
	return ""
}

func (x *Translations) GetRussian() string {
	if x != nil {
		return x.Russian
	}
	return ""
}

func (x *Translations) GetSanskrit() string {
	if x != nil {
		return x.Sanskrit
	}
	return ""
}

func (x *Translations) GetSpanish() string {
	if x != nil {
		return x.Spanish
	}
	return ""
}

var File_meta_translations_proto protoreflect.FileDescriptor

const file_meta_translations_proto_rawDesc = "" +
	"\n" +
	"\x17meta/translations.proto\x12\x04meta\"\xf2\x03\n" +
	"\fTranslations\x12#\n" +
	"\fancientGreek\x18\x01 \x01(\tR\rAncient Greek\x12\x16\n" +
	"\x06arabic\x18\x02 \x01(\tR\x06Arabic\x12-\n" +
	"\x11chineseSimplified\x18\x03 \x01(\tR\x12Chinese Simplified\x12/\n" +
	"\x12chineseTraditional\x18\x04 \x01(\tR\x13Chinese Traditional\x12\x16\n" +
	"\x06french\x18\x05 \x01(\tR\x06French\x12\x16\n" +
	"\x06german\x18\x06 \x01(\tR\x06German\x12\x16\n" +
	"\x06hebrew\x18\a \x01(\tR\x06Hebrew\x12\x18\n" +
	"\aitalian\x18\b \x01(\tR\aItalian\x12\x1a\n" +
	"\bjapanese\x18\t \x01(\tR\bJapanese\x12\x16\n" +
	"\x06korean\x18\n" +
	" \x01(\tR\x06Korean\x12\x14\n" +
	"\x05latin\x18\v \x01(\tR\x05Latin\x12\x1c\n" +
	"\tphyrexian\x18\f \x01(\tR\tPhyrexian\x12+\n" +
	"\x0fportugeseBrazil\x18\r \x01(\tR\x12Portugese (Brazil)\x12\x18\n" +
	"\arussian\x18\x0e \x01(\tR\aRussian\x12\x1a\n" +
	"\bsanskrit\x18\x0f \x01(\tR\bSanskrit\x12\x18\n" +
	"\aspanish\x18\x10 \x01(\tR\aSpanishB+Z)github.com/stevezaluk/mtgjson-models/metab\x06proto3"

var (
	file_meta_translations_proto_rawDescOnce sync.Once
	file_meta_translations_proto_rawDescData []byte
)

func file_meta_translations_proto_rawDescGZIP() []byte {
	file_meta_translations_proto_rawDescOnce.Do(func() {
		file_meta_translations_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_meta_translations_proto_rawDesc), len(file_meta_translations_proto_rawDesc)))
	})
	return file_meta_translations_proto_rawDescData
}

var file_meta_translations_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_meta_translations_proto_goTypes = []any{
	(*Translations)(nil), // 0: meta.Translations
}
var file_meta_translations_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_meta_translations_proto_init() }
func file_meta_translations_proto_init() {
	if File_meta_translations_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_meta_translations_proto_rawDesc), len(file_meta_translations_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_meta_translations_proto_goTypes,
		DependencyIndexes: file_meta_translations_proto_depIdxs,
		MessageInfos:      file_meta_translations_proto_msgTypes,
	}.Build()
	File_meta_translations_proto = out.File
	file_meta_translations_proto_goTypes = nil
	file_meta_translations_proto_depIdxs = nil
}
