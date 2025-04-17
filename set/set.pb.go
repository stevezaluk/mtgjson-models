// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: set/set.proto

package set

import (
	booster "github.com/stevezaluk/mtgjson-models/booster"
	card "github.com/stevezaluk/mtgjson-models/card"
	meta "github.com/stevezaluk/mtgjson-models/meta"
	sealed "github.com/stevezaluk/mtgjson-models/sealed"
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

type Set struct {
	state            protoimpl.MessageState            `protogen:"open.v1"`
	BaseSetSize      int64                             `protobuf:"varint,1,opt,name=baseSetSize,proto3" json:"baseSetSize,omitempty" bson:"baseSetSize"`                                                                  // @gotags: bson:"baseSetSize"
	Block            string                            `protobuf:"bytes,2,opt,name=block,proto3" json:"block,omitempty" bson:"block"`                                                                               // @gotags: bson:"block"
	Booster          map[string]*booster.BoosterConfig `protobuf:"bytes,3,rep,name=booster,proto3" json:"booster,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value" bson:"booster"` // @gotags: bson:"booster"
	Contents         []*card.CardSet                   `protobuf:"bytes,4,rep,name=contents,proto3" json:"contents,omitempty" bson:"contents"`                                                                         // @gotags: bson:"contents"
	ContentIds       []string                          `protobuf:"bytes,5,rep,name=contentIds,proto3" json:"contentIds,omitempty" bson:"contentIds"`                                                                     // @gotags: bson:"contentIds"
	CardsphereId     int64                             `protobuf:"varint,6,opt,name=cardsphereId,proto3" json:"cardsphereId,omitempty" bson:"cardsphereId"`                                                                // @gotags: bson:"cardsphereId"
	Code             string                            `protobuf:"bytes,7,opt,name=code,proto3" json:"code,omitempty" bson:"code"`                                                                                 // @gotags: bson:"code"
	CodeV3           string                            `protobuf:"bytes,8,opt,name=codeV3,proto3" json:"codeV3,omitempty" bson:"codeV3"`                                                                             // @gotags: bson:"codeV3"
	Decks            []string                          `protobuf:"bytes,9,rep,name=decks,proto3" json:"decks,omitempty" bson:"decks"`                                                                               // @gotags: bson:"decks"
	IsForeignOnly    bool                              `protobuf:"varint,10,opt,name=isForeignOnly,proto3" json:"isForeignOnly,omitempty" bson:"isForeignOnly"`                                                             // @gotags: bson:"isForeignOnly"
	IsFoilOnly       bool                              `protobuf:"varint,11,opt,name=isFoilOnly,proto3" json:"isFoilOnly,omitempty" bson:"isFoilOnly"`                                                                   // @gotags: bson:"isFoilOnly"
	IsNonFoilOnly    bool                              `protobuf:"varint,12,opt,name=isNonFoilOnly,proto3" json:"isNonFoilOnly,omitempty" bson:"isNonFoilOnly"`                                                             // @gotags: bson:"isNonFoilOnly"
	IsOnlineOnly     bool                              `protobuf:"varint,13,opt,name=isOnlineOnly,proto3" json:"isOnlineOnly,omitempty" bson:"isOnlineOnly"`                                                               // @gotags: bson:"isOnlineOnly"
	IsPaperOnly      bool                              `protobuf:"varint,14,opt,name=isPaperOnly,proto3" json:"isPaperOnly,omitempty" bson:"isPaperOnly"`                                                                 // @gotags: bson:"isPaperOnly"
	IsPartialPreview bool                              `protobuf:"varint,15,opt,name=isPartialPreview,proto3" json:"isPartialPreview,omitempty" bson:"isPartialPreview"`                                                       // @gotags: bson:"isPartialPreview"
	KeyruneCode      string                            `protobuf:"bytes,16,opt,name=keyruneCode,proto3" json:"keyruneCode,omitempty" bson:"keyrunCode"`                                                                  // @gotags: bson:"keyrunCode"
	Languages        []string                          `protobuf:"bytes,17,rep,name=languages,proto3" json:"languages,omitempty" bson:"languages"`                                                                      // @gotags: bson:"languages"
	McmId            int64                             `protobuf:"varint,18,opt,name=mcmId,proto3" json:"mcmId,omitempty" bson:"mcmId"`                                                                             // @gotags: bson:"mcmId"
	McmIdExtras      int64                             `protobuf:"varint,19,opt,name=mcmIdExtras,proto3" json:"mcmIdExtras,omitempty" bson:"mcmIdExtras"`                                                                 // @gotags: bson:"mcmIdExtras"
	McmName          string                            `protobuf:"bytes,20,opt,name=mcmName,proto3" json:"mcmName,omitempty" bson:"mcmName"`                                                                          // @gotags: bson:"mcmName"
	MtgoCode         string                            `protobuf:"bytes,21,opt,name=mtgoCode,proto3" json:"mtgoCode,omitempty" bson:"mtgoCode"`                                                                        // @gotags: bson:"mtgoCode"
	Name             string                            `protobuf:"bytes,22,opt,name=name,proto3" json:"name,omitempty" bson:"name"`                                                                                // @gotags: bson:"name"
	ParentCode       string                            `protobuf:"bytes,23,opt,name=parentCode,proto3" json:"parentCode,omitempty" bson:"parentCode"`                                                                    // @gotags: bson:"parentCode"
	ReleaseDate      string                            `protobuf:"bytes,24,opt,name=releaseDate,proto3" json:"releaseDate,omitempty" bson:"releaseDate"`                                                                  // @gotags: bson:"releaseDate"
	SealedProduct    []*sealed.SealedProduct           `protobuf:"bytes,25,rep,name=sealedProduct,proto3" json:"sealedProduct,omitempty" bson:"sealedProduct"`                                                              // @gotags: bson:"sealedProduct"
	TcgPlayerGroupId int64                             `protobuf:"varint,26,opt,name=tcgPlayerGroupId,proto3" json:"tcgPlayerGroupId,omitempty" bson:"tcgPlayerGroupId"`                                                       // @gotags: bson:"tcgPlayerGroupId"
	Tokens           []string                          `protobuf:"bytes,27,rep,name=tokens,proto3" json:"tokens,omitempty" bson:"tokens"`                                                                            // @gotags: bson:"tokens"
	TokenSetCode     string                            `protobuf:"bytes,28,opt,name=tokenSetCode,proto3" json:"tokenSetCode,omitempty" bson:"tokenSetCode"`                                                                // @gotags: bson:"tokenSetCode"
	TokenSetSize     int64                             `protobuf:"varint,29,opt,name=tokenSetSize,proto3" json:"tokenSetSize,omitempty" bson:"tokenSetSize"`                                                               // @gotags: bson:"tokenSetSize"
	Translations     *meta.Translations                `protobuf:"bytes,30,opt,name=translations,proto3" json:"translations,omitempty" bson:"translations"`                                                                // @gotags: bson:"translations"
	Type             string                            `protobuf:"bytes,31,opt,name=type,proto3" json:"type,omitempty" bson:"type"`                                                                                // @gotags: bson:"type"
	MtgjsonApiMeta   *meta.MTGJSONAPIMeta              `protobuf:"bytes,32,opt,name=mtgjsonApiMeta,proto3" json:"mtgjsonApiMeta,omitempty" bson:"mtgjsonApiMeta"`                                                            // @gotags: bson:"mtgjsonApiMeta"
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *Set) Reset() {
	*x = Set{}
	mi := &file_set_set_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Set) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Set) ProtoMessage() {}

func (x *Set) ProtoReflect() protoreflect.Message {
	mi := &file_set_set_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Set.ProtoReflect.Descriptor instead.
func (*Set) Descriptor() ([]byte, []int) {
	return file_set_set_proto_rawDescGZIP(), []int{0}
}

func (x *Set) GetBaseSetSize() int64 {
	if x != nil {
		return x.BaseSetSize
	}
	return 0
}

func (x *Set) GetBlock() string {
	if x != nil {
		return x.Block
	}
	return ""
}

func (x *Set) GetBooster() map[string]*booster.BoosterConfig {
	if x != nil {
		return x.Booster
	}
	return nil
}

func (x *Set) GetContents() []*card.CardSet {
	if x != nil {
		return x.Contents
	}
	return nil
}

func (x *Set) GetContentIds() []string {
	if x != nil {
		return x.ContentIds
	}
	return nil
}

func (x *Set) GetCardsphereId() int64 {
	if x != nil {
		return x.CardsphereId
	}
	return 0
}

func (x *Set) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Set) GetCodeV3() string {
	if x != nil {
		return x.CodeV3
	}
	return ""
}

func (x *Set) GetDecks() []string {
	if x != nil {
		return x.Decks
	}
	return nil
}

func (x *Set) GetIsForeignOnly() bool {
	if x != nil {
		return x.IsForeignOnly
	}
	return false
}

func (x *Set) GetIsFoilOnly() bool {
	if x != nil {
		return x.IsFoilOnly
	}
	return false
}

func (x *Set) GetIsNonFoilOnly() bool {
	if x != nil {
		return x.IsNonFoilOnly
	}
	return false
}

func (x *Set) GetIsOnlineOnly() bool {
	if x != nil {
		return x.IsOnlineOnly
	}
	return false
}

func (x *Set) GetIsPaperOnly() bool {
	if x != nil {
		return x.IsPaperOnly
	}
	return false
}

func (x *Set) GetIsPartialPreview() bool {
	if x != nil {
		return x.IsPartialPreview
	}
	return false
}

func (x *Set) GetKeyruneCode() string {
	if x != nil {
		return x.KeyruneCode
	}
	return ""
}

func (x *Set) GetLanguages() []string {
	if x != nil {
		return x.Languages
	}
	return nil
}

func (x *Set) GetMcmId() int64 {
	if x != nil {
		return x.McmId
	}
	return 0
}

func (x *Set) GetMcmIdExtras() int64 {
	if x != nil {
		return x.McmIdExtras
	}
	return 0
}

func (x *Set) GetMcmName() string {
	if x != nil {
		return x.McmName
	}
	return ""
}

func (x *Set) GetMtgoCode() string {
	if x != nil {
		return x.MtgoCode
	}
	return ""
}

func (x *Set) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Set) GetParentCode() string {
	if x != nil {
		return x.ParentCode
	}
	return ""
}

func (x *Set) GetReleaseDate() string {
	if x != nil {
		return x.ReleaseDate
	}
	return ""
}

func (x *Set) GetSealedProduct() []*sealed.SealedProduct {
	if x != nil {
		return x.SealedProduct
	}
	return nil
}

func (x *Set) GetTcgPlayerGroupId() int64 {
	if x != nil {
		return x.TcgPlayerGroupId
	}
	return 0
}

func (x *Set) GetTokens() []string {
	if x != nil {
		return x.Tokens
	}
	return nil
}

func (x *Set) GetTokenSetCode() string {
	if x != nil {
		return x.TokenSetCode
	}
	return ""
}

func (x *Set) GetTokenSetSize() int64 {
	if x != nil {
		return x.TokenSetSize
	}
	return 0
}

func (x *Set) GetTranslations() *meta.Translations {
	if x != nil {
		return x.Translations
	}
	return nil
}

func (x *Set) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Set) GetMtgjsonApiMeta() *meta.MTGJSONAPIMeta {
	if x != nil {
		return x.MtgjsonApiMeta
	}
	return nil
}

var File_set_set_proto protoreflect.FileDescriptor

const file_set_set_proto_rawDesc = "" +
	"\n" +
	"\rset/set.proto\x12\x03set\x1a\x1cbooster/booster_config.proto\x1a\x1bsealed/sealed_product.proto\x1a\x17meta/translations.proto\x1a\x16meta/mtgjson_api.proto\x1a\x13card/card_set.proto\"\xa8\t\n" +
	"\x03Set\x12 \n" +
	"\vbaseSetSize\x18\x01 \x01(\x03R\vbaseSetSize\x12\x14\n" +
	"\x05block\x18\x02 \x01(\tR\x05block\x12/\n" +
	"\abooster\x18\x03 \x03(\v2\x15.set.Set.BoosterEntryR\abooster\x12)\n" +
	"\bcontents\x18\x04 \x03(\v2\r.card.CardSetR\bcontents\x12\x1e\n" +
	"\n" +
	"contentIds\x18\x05 \x03(\tR\n" +
	"contentIds\x12\"\n" +
	"\fcardsphereId\x18\x06 \x01(\x03R\fcardsphereId\x12\x12\n" +
	"\x04code\x18\a \x01(\tR\x04code\x12\x16\n" +
	"\x06codeV3\x18\b \x01(\tR\x06codeV3\x12\x14\n" +
	"\x05decks\x18\t \x03(\tR\x05decks\x12$\n" +
	"\risForeignOnly\x18\n" +
	" \x01(\bR\risForeignOnly\x12\x1e\n" +
	"\n" +
	"isFoilOnly\x18\v \x01(\bR\n" +
	"isFoilOnly\x12$\n" +
	"\risNonFoilOnly\x18\f \x01(\bR\risNonFoilOnly\x12\"\n" +
	"\fisOnlineOnly\x18\r \x01(\bR\fisOnlineOnly\x12 \n" +
	"\visPaperOnly\x18\x0e \x01(\bR\visPaperOnly\x12*\n" +
	"\x10isPartialPreview\x18\x0f \x01(\bR\x10isPartialPreview\x12 \n" +
	"\vkeyruneCode\x18\x10 \x01(\tR\vkeyruneCode\x12\x1c\n" +
	"\tlanguages\x18\x11 \x03(\tR\tlanguages\x12\x14\n" +
	"\x05mcmId\x18\x12 \x01(\x03R\x05mcmId\x12 \n" +
	"\vmcmIdExtras\x18\x13 \x01(\x03R\vmcmIdExtras\x12\x18\n" +
	"\amcmName\x18\x14 \x01(\tR\amcmName\x12\x1a\n" +
	"\bmtgoCode\x18\x15 \x01(\tR\bmtgoCode\x12\x12\n" +
	"\x04name\x18\x16 \x01(\tR\x04name\x12\x1e\n" +
	"\n" +
	"parentCode\x18\x17 \x01(\tR\n" +
	"parentCode\x12 \n" +
	"\vreleaseDate\x18\x18 \x01(\tR\vreleaseDate\x12;\n" +
	"\rsealedProduct\x18\x19 \x03(\v2\x15.sealed.SealedProductR\rsealedProduct\x12*\n" +
	"\x10tcgPlayerGroupId\x18\x1a \x01(\x03R\x10tcgPlayerGroupId\x12\x16\n" +
	"\x06tokens\x18\x1b \x03(\tR\x06tokens\x12\"\n" +
	"\ftokenSetCode\x18\x1c \x01(\tR\ftokenSetCode\x12\"\n" +
	"\ftokenSetSize\x18\x1d \x01(\x03R\ftokenSetSize\x126\n" +
	"\ftranslations\x18\x1e \x01(\v2\x12.meta.TranslationsR\ftranslations\x12\x12\n" +
	"\x04type\x18\x1f \x01(\tR\x04type\x12<\n" +
	"\x0emtgjsonApiMeta\x18  \x01(\v2\x14.meta.MTGJSONAPIMetaR\x0emtgjsonApiMeta\x1aR\n" +
	"\fBoosterEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12,\n" +
	"\x05value\x18\x02 \x01(\v2\x16.booster.BoosterConfigR\x05value:\x028\x01B*Z(github.com/stevezaluk/mtgjson-models/setb\x06proto3"

var (
	file_set_set_proto_rawDescOnce sync.Once
	file_set_set_proto_rawDescData []byte
)

func file_set_set_proto_rawDescGZIP() []byte {
	file_set_set_proto_rawDescOnce.Do(func() {
		file_set_set_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_set_set_proto_rawDesc), len(file_set_set_proto_rawDesc)))
	})
	return file_set_set_proto_rawDescData
}

var file_set_set_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_set_set_proto_goTypes = []any{
	(*Set)(nil),                   // 0: set.Set
	nil,                           // 1: set.Set.BoosterEntry
	(*card.CardSet)(nil),          // 2: card.CardSet
	(*sealed.SealedProduct)(nil),  // 3: sealed.SealedProduct
	(*meta.Translations)(nil),     // 4: meta.Translations
	(*meta.MTGJSONAPIMeta)(nil),   // 5: meta.MTGJSONAPIMeta
	(*booster.BoosterConfig)(nil), // 6: booster.BoosterConfig
}
var file_set_set_proto_depIdxs = []int32{
	1, // 0: set.Set.booster:type_name -> set.Set.BoosterEntry
	2, // 1: set.Set.contents:type_name -> card.CardSet
	3, // 2: set.Set.sealedProduct:type_name -> sealed.SealedProduct
	4, // 3: set.Set.translations:type_name -> meta.Translations
	5, // 4: set.Set.mtgjsonApiMeta:type_name -> meta.MTGJSONAPIMeta
	6, // 5: set.Set.BoosterEntry.value:type_name -> booster.BoosterConfig
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_set_set_proto_init() }
func file_set_set_proto_init() {
	if File_set_set_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_set_set_proto_rawDesc), len(file_set_set_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_set_set_proto_goTypes,
		DependencyIndexes: file_set_set_proto_depIdxs,
		MessageInfos:      file_set_set_proto_msgTypes,
	}.Build()
	File_set_set_proto = out.File
	file_set_set_proto_goTypes = nil
	file_set_set_proto_depIdxs = nil
}
