// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.12.4
// source: deck/deck.proto

package deck

import (
	card "github.com/stevezaluk/mtgjson-models/card"
	meta "github.com/stevezaluk/mtgjson-models/meta"
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

// Deck - Represents an MTGJSON deck
//
// See here for more info: https://mtgjson.com/data-models/deck
//
// This model has been changed slightly from the original model to seperate out lists for
// both the card ids and for cards themselves. This way we can store just the ID's in the database
// and fill the contents field before it gets returned to the user
type Deck struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	Name           string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" bson:"name"`                                                                                      // @gotags: bson:"name"
	Code           string                 `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty" bson:"code"`                                                                                      // @gotags: bson:"code"
	Type           string                 `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty" bson:"type"`                                                                                      // @gotags: bson:"type"
	ReleaseDate    string                 `protobuf:"bytes,4,opt,name=releaseDate,proto3" json:"releaseDate,omitempty" bson:"releaseDate"`                                                                        // @gotags: bson:"releaseDate"
	MainBoard      map[string]int64       `protobuf:"bytes,5,rep,name=mainBoard,proto3" json:"mainBoard,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value" bson:"mainBoard"` // @gotags: bson:"mainBoard"
	SideBoard      map[string]int64       `protobuf:"bytes,6,rep,name=sideBoard,proto3" json:"sideBoard,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value" bson:"sideBoard"` // @gotags: bson:"sideBoard"
	Commander      map[string]int64       `protobuf:"bytes,7,rep,name=commander,proto3" json:"commander,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value" bson:"commander"` // @gotags: bson:"commander"
	MtgjsonApiMeta *meta.MTGJSONAPIMeta   `protobuf:"bytes,8,opt,name=mtgjsonApiMeta,proto3" json:"mtgjsonApiMeta,omitempty" bson:"mtgjsonApiMeta"`                                                                  // @gotags: bson:"mtgjsonApiMeta"
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *Deck) Reset() {
	*x = Deck{}
	mi := &file_deck_deck_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Deck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Deck) ProtoMessage() {}

func (x *Deck) ProtoReflect() protoreflect.Message {
	mi := &file_deck_deck_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Deck.ProtoReflect.Descriptor instead.
func (*Deck) Descriptor() ([]byte, []int) {
	return file_deck_deck_proto_rawDescGZIP(), []int{0}
}

func (x *Deck) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Deck) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Deck) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Deck) GetReleaseDate() string {
	if x != nil {
		return x.ReleaseDate
	}
	return ""
}

func (x *Deck) GetMainBoard() map[string]int64 {
	if x != nil {
		return x.MainBoard
	}
	return nil
}

func (x *Deck) GetSideBoard() map[string]int64 {
	if x != nil {
		return x.SideBoard
	}
	return nil
}

func (x *Deck) GetCommander() map[string]int64 {
	if x != nil {
		return x.Commander
	}
	return nil
}

func (x *Deck) GetMtgjsonApiMeta() *meta.MTGJSONAPIMeta {
	if x != nil {
		return x.MtgjsonApiMeta
	}
	return nil
}

// DeckContents - Represents the contents of deck that are returned to the caller
type DeckContents struct {
	state         protoimpl.MessageState       `protogen:"open.v1"`
	MainBoard     map[string]*DeckContentEntry `protobuf:"bytes,1,rep,name=mainBoard,proto3" json:"mainBoard,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value" bson:"mainBoard"` // @gotags: bson:"mainBoard"
	SideBoard     map[string]*DeckContentEntry `protobuf:"bytes,2,rep,name=sideBoard,proto3" json:"sideBoard,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value" bson:"sideBoard"` // @gotags: bson:"sideBoard"
	Commander     map[string]*DeckContentEntry `protobuf:"bytes,3,rep,name=commander,proto3" json:"commander,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value" bson:"commander"` // @gotags: bson:"commander"
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeckContents) Reset() {
	*x = DeckContents{}
	mi := &file_deck_deck_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeckContents) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeckContents) ProtoMessage() {}

func (x *DeckContents) ProtoReflect() protoreflect.Message {
	mi := &file_deck_deck_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeckContents.ProtoReflect.Descriptor instead.
func (*DeckContents) Descriptor() ([]byte, []int) {
	return file_deck_deck_proto_rawDescGZIP(), []int{1}
}

func (x *DeckContents) GetMainBoard() map[string]*DeckContentEntry {
	if x != nil {
		return x.MainBoard
	}
	return nil
}

func (x *DeckContents) GetSideBoard() map[string]*DeckContentEntry {
	if x != nil {
		return x.SideBoard
	}
	return nil
}

func (x *DeckContents) GetCommander() map[string]*DeckContentEntry {
	if x != nil {
		return x.Commander
	}
	return nil
}

// DeckContentEntry - Represents a single entry in the DeckContents structure
type DeckContentEntry struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Quantity      int64                  `protobuf:"varint,1,opt,name=quantity,proto3" json:"quantity,omitempty" bson:"quantity"` // @gotags: bson:"quantity"
	Card          *card.CardSet          `protobuf:"bytes,2,opt,name=card,proto3" json:"card,omitempty" bson:"card"`          // @gotags: bson:"card"
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeckContentEntry) Reset() {
	*x = DeckContentEntry{}
	mi := &file_deck_deck_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeckContentEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeckContentEntry) ProtoMessage() {}

func (x *DeckContentEntry) ProtoReflect() protoreflect.Message {
	mi := &file_deck_deck_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeckContentEntry.ProtoReflect.Descriptor instead.
func (*DeckContentEntry) Descriptor() ([]byte, []int) {
	return file_deck_deck_proto_rawDescGZIP(), []int{2}
}

func (x *DeckContentEntry) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *DeckContentEntry) GetCard() *card.CardSet {
	if x != nil {
		return x.Card
	}
	return nil
}

var File_deck_deck_proto protoreflect.FileDescriptor

const file_deck_deck_proto_rawDesc = "" +
	"\n" +
	"\x0fdeck/deck.proto\x12\x04deck\x1a\x13card/card_set.proto\x1a\x16meta/mtgjson_api.proto\"\x87\x04\n" +
	"\x04Deck\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x12\n" +
	"\x04code\x18\x02 \x01(\tR\x04code\x12\x12\n" +
	"\x04type\x18\x03 \x01(\tR\x04type\x12 \n" +
	"\vreleaseDate\x18\x04 \x01(\tR\vreleaseDate\x127\n" +
	"\tmainBoard\x18\x05 \x03(\v2\x19.deck.Deck.MainBoardEntryR\tmainBoard\x127\n" +
	"\tsideBoard\x18\x06 \x03(\v2\x19.deck.Deck.SideBoardEntryR\tsideBoard\x127\n" +
	"\tcommander\x18\a \x03(\v2\x19.deck.Deck.CommanderEntryR\tcommander\x12<\n" +
	"\x0emtgjsonApiMeta\x18\b \x01(\v2\x14.meta.MTGJSONAPIMetaR\x0emtgjsonApiMeta\x1a<\n" +
	"\x0eMainBoardEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\x03R\x05value:\x028\x01\x1a<\n" +
	"\x0eSideBoardEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\x03R\x05value:\x028\x01\x1a<\n" +
	"\x0eCommanderEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\x03R\x05value:\x028\x01\"\xd3\x03\n" +
	"\fDeckContents\x12?\n" +
	"\tmainBoard\x18\x01 \x03(\v2!.deck.DeckContents.MainBoardEntryR\tmainBoard\x12?\n" +
	"\tsideBoard\x18\x02 \x03(\v2!.deck.DeckContents.SideBoardEntryR\tsideBoard\x12?\n" +
	"\tcommander\x18\x03 \x03(\v2!.deck.DeckContents.CommanderEntryR\tcommander\x1aT\n" +
	"\x0eMainBoardEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12,\n" +
	"\x05value\x18\x02 \x01(\v2\x16.deck.DeckContentEntryR\x05value:\x028\x01\x1aT\n" +
	"\x0eSideBoardEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12,\n" +
	"\x05value\x18\x02 \x01(\v2\x16.deck.DeckContentEntryR\x05value:\x028\x01\x1aT\n" +
	"\x0eCommanderEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12,\n" +
	"\x05value\x18\x02 \x01(\v2\x16.deck.DeckContentEntryR\x05value:\x028\x01\"Q\n" +
	"\x10DeckContentEntry\x12\x1a\n" +
	"\bquantity\x18\x01 \x01(\x03R\bquantity\x12!\n" +
	"\x04card\x18\x02 \x01(\v2\r.card.CardSetR\x04cardB+Z)github.com/stevezaluk/mtgjson-models/deckb\x06proto3"

var (
	file_deck_deck_proto_rawDescOnce sync.Once
	file_deck_deck_proto_rawDescData []byte
)

func file_deck_deck_proto_rawDescGZIP() []byte {
	file_deck_deck_proto_rawDescOnce.Do(func() {
		file_deck_deck_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_deck_deck_proto_rawDesc), len(file_deck_deck_proto_rawDesc)))
	})
	return file_deck_deck_proto_rawDescData
}

var file_deck_deck_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_deck_deck_proto_goTypes = []any{
	(*Deck)(nil),                // 0: deck.Deck
	(*DeckContents)(nil),        // 1: deck.DeckContents
	(*DeckContentEntry)(nil),    // 2: deck.DeckContentEntry
	nil,                         // 3: deck.Deck.MainBoardEntry
	nil,                         // 4: deck.Deck.SideBoardEntry
	nil,                         // 5: deck.Deck.CommanderEntry
	nil,                         // 6: deck.DeckContents.MainBoardEntry
	nil,                         // 7: deck.DeckContents.SideBoardEntry
	nil,                         // 8: deck.DeckContents.CommanderEntry
	(*meta.MTGJSONAPIMeta)(nil), // 9: meta.MTGJSONAPIMeta
	(*card.CardSet)(nil),        // 10: card.CardSet
}
var file_deck_deck_proto_depIdxs = []int32{
	3,  // 0: deck.Deck.mainBoard:type_name -> deck.Deck.MainBoardEntry
	4,  // 1: deck.Deck.sideBoard:type_name -> deck.Deck.SideBoardEntry
	5,  // 2: deck.Deck.commander:type_name -> deck.Deck.CommanderEntry
	9,  // 3: deck.Deck.mtgjsonApiMeta:type_name -> meta.MTGJSONAPIMeta
	6,  // 4: deck.DeckContents.mainBoard:type_name -> deck.DeckContents.MainBoardEntry
	7,  // 5: deck.DeckContents.sideBoard:type_name -> deck.DeckContents.SideBoardEntry
	8,  // 6: deck.DeckContents.commander:type_name -> deck.DeckContents.CommanderEntry
	10, // 7: deck.DeckContentEntry.card:type_name -> card.CardSet
	2,  // 8: deck.DeckContents.MainBoardEntry.value:type_name -> deck.DeckContentEntry
	2,  // 9: deck.DeckContents.SideBoardEntry.value:type_name -> deck.DeckContentEntry
	2,  // 10: deck.DeckContents.CommanderEntry.value:type_name -> deck.DeckContentEntry
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_deck_deck_proto_init() }
func file_deck_deck_proto_init() {
	if File_deck_deck_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_deck_deck_proto_rawDesc), len(file_deck_deck_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_deck_deck_proto_goTypes,
		DependencyIndexes: file_deck_deck_proto_depIdxs,
		MessageInfos:      file_deck_deck_proto_msgTypes,
	}.Build()
	File_deck_deck_proto = out.File
	file_deck_deck_proto_goTypes = nil
	file_deck_deck_proto_depIdxs = nil
}
