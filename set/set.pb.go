// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.6.1
// source: set/set.proto

package set

import (
	booster "github.com/stevezaluk/mtgjson-models/booster"
	meta "github.com/stevezaluk/mtgjson-models/meta"
	sealed "github.com/stevezaluk/mtgjson-models/sealed"
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

type Set struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseSetSize      int64                             `protobuf:"varint,1,opt,name=baseSetSize,proto3" json:"baseSetSize,omitempty"`
	Block            string                            `protobuf:"bytes,2,opt,name=block,proto3" json:"block,omitempty"`
	Booster          map[string]*booster.BoosterConfig `protobuf:"bytes,3,rep,name=booster,proto3" json:"booster,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Cards            []string                          `protobuf:"bytes,4,rep,name=cards,proto3" json:"cards,omitempty"`
	CardsphereId     int64                             `protobuf:"varint,5,opt,name=cardsphereId,proto3" json:"cardsphereId,omitempty"`
	Code             string                            `protobuf:"bytes,6,opt,name=code,proto3" json:"code,omitempty"`
	CodeV3           string                            `protobuf:"bytes,7,opt,name=codeV3,proto3" json:"codeV3,omitempty"`
	Decks            []string                          `protobuf:"bytes,8,rep,name=decks,proto3" json:"decks,omitempty"`
	IsForeignOnly    bool                              `protobuf:"varint,9,opt,name=isForeignOnly,proto3" json:"isForeignOnly,omitempty"`
	IsFoilOnly       bool                              `protobuf:"varint,10,opt,name=isFoilOnly,proto3" json:"isFoilOnly,omitempty"`
	IsNonFoilOnly    bool                              `protobuf:"varint,11,opt,name=isNonFoilOnly,proto3" json:"isNonFoilOnly,omitempty"`
	IsOnlineOnly     bool                              `protobuf:"varint,12,opt,name=isOnlineOnly,proto3" json:"isOnlineOnly,omitempty"`
	IsPaperOnly      bool                              `protobuf:"varint,13,opt,name=isPaperOnly,proto3" json:"isPaperOnly,omitempty"`
	IsPartialPreview bool                              `protobuf:"varint,14,opt,name=isPartialPreview,proto3" json:"isPartialPreview,omitempty"`
	KeyruneCode      string                            `protobuf:"bytes,15,opt,name=keyruneCode,proto3" json:"keyruneCode,omitempty"`
	Languages        []string                          `protobuf:"bytes,16,rep,name=languages,proto3" json:"languages,omitempty"`
	McmId            int64                             `protobuf:"varint,17,opt,name=mcmId,proto3" json:"mcmId,omitempty"`
	McmIdExtras      int64                             `protobuf:"varint,18,opt,name=mcmIdExtras,proto3" json:"mcmIdExtras,omitempty"`
	McmName          string                            `protobuf:"bytes,19,opt,name=mcmName,proto3" json:"mcmName,omitempty"`
	MtgoCode         string                            `protobuf:"bytes,20,opt,name=mtgoCode,proto3" json:"mtgoCode,omitempty"`
	Name             string                            `protobuf:"bytes,21,opt,name=name,proto3" json:"name,omitempty"`
	ParentCode       string                            `protobuf:"bytes,22,opt,name=parentCode,proto3" json:"parentCode,omitempty"`
	ReleaseDate      string                            `protobuf:"bytes,23,opt,name=releaseDate,proto3" json:"releaseDate,omitempty"`
	SealedProduct    []*sealed.SealedProduct           `protobuf:"bytes,24,rep,name=sealedProduct,proto3" json:"sealedProduct,omitempty"`
	TcgPlayerGroupId int64                             `protobuf:"varint,25,opt,name=tcgPlayerGroupId,proto3" json:"tcgPlayerGroupId,omitempty"`
	Tokens           []string                          `protobuf:"bytes,26,rep,name=tokens,proto3" json:"tokens,omitempty"`
	TokenSetCode     string                            `protobuf:"bytes,27,opt,name=tokenSetCode,proto3" json:"tokenSetCode,omitempty"`
	TokenSetSize     int64                             `protobuf:"varint,28,opt,name=tokenSetSize,proto3" json:"tokenSetSize,omitempty"`
	Translations     *meta.Translations                `protobuf:"bytes,29,opt,name=translations,proto3" json:"translations,omitempty"`
	Type             string                            `protobuf:"bytes,30,opt,name=type,proto3" json:"type,omitempty"`
	MtgjsonApiMeta   *meta.MTGJSONAPIMeta              `protobuf:"bytes,31,opt,name=mtgjsonApiMeta,proto3" json:"mtgjsonApiMeta,omitempty"`
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

func (x *Set) GetCards() []string {
	if x != nil {
		return x.Cards
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

var file_set_set_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x65, 0x74, 0x2f, 0x73, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x73, 0x65, 0x74, 0x1a, 0x1c, 0x62, 0x6f, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x62, 0x6f,
	0x6f, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x73, 0x65, 0x61, 0x6c, 0x65, 0x64, 0x2f, 0x73, 0x65, 0x61, 0x6c, 0x65,
	0x64, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x6d,
	0x74, 0x67, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xf3, 0x08, 0x0a, 0x03, 0x53, 0x65, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x62, 0x61, 0x73, 0x65,
	0x53, 0x65, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x62,
	0x61, 0x73, 0x65, 0x53, 0x65, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x12, 0x2f, 0x0a, 0x07, 0x62, 0x6f, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x73, 0x65, 0x74, 0x2e, 0x53, 0x65, 0x74, 0x2e, 0x42, 0x6f, 0x6f, 0x73,
	0x74, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x62, 0x6f, 0x6f, 0x73, 0x74, 0x65,
	0x72, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x61, 0x72, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x05, 0x63, 0x61, 0x72, 0x64, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x61, 0x72, 0x64, 0x73,
	0x70, 0x68, 0x65, 0x72, 0x65, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x63,
	0x61, 0x72, 0x64, 0x73, 0x70, 0x68, 0x65, 0x72, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x63, 0x6f, 0x64, 0x65, 0x56, 0x33, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x63, 0x6f, 0x64, 0x65, 0x56, 0x33, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x63, 0x6b, 0x73,
	0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x64, 0x65, 0x63, 0x6b, 0x73, 0x12, 0x24, 0x0a,
	0x0d, 0x69, 0x73, 0x46, 0x6f, 0x72, 0x65, 0x69, 0x67, 0x6e, 0x4f, 0x6e, 0x6c, 0x79, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x69, 0x73, 0x46, 0x6f, 0x72, 0x65, 0x69, 0x67, 0x6e, 0x4f,
	0x6e, 0x6c, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x73, 0x46, 0x6f, 0x69, 0x6c, 0x4f, 0x6e, 0x6c,
	0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x46, 0x6f, 0x69, 0x6c, 0x4f,
	0x6e, 0x6c, 0x79, 0x12, 0x24, 0x0a, 0x0d, 0x69, 0x73, 0x4e, 0x6f, 0x6e, 0x46, 0x6f, 0x69, 0x6c,
	0x4f, 0x6e, 0x6c, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x69, 0x73, 0x4e, 0x6f,
	0x6e, 0x46, 0x6f, 0x69, 0x6c, 0x4f, 0x6e, 0x6c, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x73, 0x4f,
	0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x4f, 0x6e, 0x6c, 0x79, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0c, 0x69, 0x73, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x4f, 0x6e, 0x6c, 0x79, 0x12, 0x20, 0x0a,
	0x0b, 0x69, 0x73, 0x50, 0x61, 0x70, 0x65, 0x72, 0x4f, 0x6e, 0x6c, 0x79, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x50, 0x61, 0x70, 0x65, 0x72, 0x4f, 0x6e, 0x6c, 0x79, 0x12,
	0x2a, 0x0a, 0x10, 0x69, 0x73, 0x50, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x50, 0x72, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x69, 0x73, 0x50, 0x61, 0x72,
	0x74, 0x69, 0x61, 0x6c, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x20, 0x0a, 0x0b, 0x6b,
	0x65, 0x79, 0x72, 0x75, 0x6e, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x6b, 0x65, 0x79, 0x72, 0x75, 0x6e, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73, 0x18, 0x10, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x09, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6d,
	0x63, 0x6d, 0x49, 0x64, 0x18, 0x11, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6d, 0x63, 0x6d, 0x49,
	0x64, 0x12, 0x20, 0x0a, 0x0b, 0x6d, 0x63, 0x6d, 0x49, 0x64, 0x45, 0x78, 0x74, 0x72, 0x61, 0x73,
	0x18, 0x12, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6d, 0x63, 0x6d, 0x49, 0x64, 0x45, 0x78, 0x74,
	0x72, 0x61, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x63, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x13,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x63, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x6d, 0x74, 0x67, 0x6f, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6d, 0x74, 0x67, 0x6f, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x16, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x44, 0x61, 0x74, 0x65, 0x18, 0x17, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x44, 0x61, 0x74, 0x65, 0x12,
	0x3b, 0x0a, 0x0d, 0x73, 0x65, 0x61, 0x6c, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x18, 0x18, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x65, 0x61, 0x6c, 0x65, 0x64, 0x2e,
	0x53, 0x65, 0x61, 0x6c, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x0d, 0x73,
	0x65, 0x61, 0x6c, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x2a, 0x0a, 0x10,
	0x74, 0x63, 0x67, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64,
	0x18, 0x19, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x74, 0x63, 0x67, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x73, 0x18, 0x1a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73,
	0x12, 0x22, 0x0a, 0x0c, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x1b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x65, 0x74,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x65, 0x74,
	0x53, 0x69, 0x7a, 0x65, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x53, 0x65, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x36, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x3c, 0x0a, 0x0e, 0x6d, 0x74, 0x67, 0x6a, 0x73, 0x6f, 0x6e, 0x41,
	0x70, 0x69, 0x4d, 0x65, 0x74, 0x61, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d,
	0x65, 0x74, 0x61, 0x2e, 0x4d, 0x54, 0x47, 0x4a, 0x53, 0x4f, 0x4e, 0x41, 0x50, 0x49, 0x4d, 0x65,
	0x74, 0x61, 0x52, 0x0e, 0x6d, 0x74, 0x67, 0x6a, 0x73, 0x6f, 0x6e, 0x41, 0x70, 0x69, 0x4d, 0x65,
	0x74, 0x61, 0x1a, 0x52, 0x0a, 0x0c, 0x42, 0x6f, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x2c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x62, 0x6f, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x42, 0x6f,
	0x6f, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x65, 0x76, 0x65, 0x7a, 0x61, 0x6c, 0x75, 0x6b, 0x2f,
	0x6d, 0x74, 0x67, 0x6a, 0x73, 0x6f, 0x6e, 0x2d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x73,
	0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_set_set_proto_rawDescOnce sync.Once
	file_set_set_proto_rawDescData = file_set_set_proto_rawDesc
)

func file_set_set_proto_rawDescGZIP() []byte {
	file_set_set_proto_rawDescOnce.Do(func() {
		file_set_set_proto_rawDescData = protoimpl.X.CompressGZIP(file_set_set_proto_rawDescData)
	})
	return file_set_set_proto_rawDescData
}

var file_set_set_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_set_set_proto_goTypes = []any{
	(*Set)(nil),                   // 0: set.Set
	nil,                           // 1: set.Set.BoosterEntry
	(*sealed.SealedProduct)(nil),  // 2: sealed.SealedProduct
	(*meta.Translations)(nil),     // 3: meta.Translations
	(*meta.MTGJSONAPIMeta)(nil),   // 4: meta.MTGJSONAPIMeta
	(*booster.BoosterConfig)(nil), // 5: booster.BoosterConfig
}
var file_set_set_proto_depIdxs = []int32{
	1, // 0: set.Set.booster:type_name -> set.Set.BoosterEntry
	2, // 1: set.Set.sealedProduct:type_name -> sealed.SealedProduct
	3, // 2: set.Set.translations:type_name -> meta.Translations
	4, // 3: set.Set.mtgjsonApiMeta:type_name -> meta.MTGJSONAPIMeta
	5, // 4: set.Set.BoosterEntry.value:type_name -> booster.BoosterConfig
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
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
			RawDescriptor: file_set_set_proto_rawDesc,
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
	file_set_set_proto_rawDesc = nil
	file_set_set_proto_goTypes = nil
	file_set_set_proto_depIdxs = nil
}