// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: user/user.proto

package user

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

// User - Represents an MTGJSON API user. The fields: decks, cards, and sets are a list of
// MTGJSON V4 UUID's that represent the objects that they own
type User struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Username      string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty" bson:"username"`     // @gotags: bson:"username"
	Email         string                 `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty" bson:"email"`           // @gotags: bson:"email"
	Auth0Id       string                 `protobuf:"bytes,3,opt,name=auth0Id,proto3" json:"auth0Id,omitempty" bson:"auth0Id"`       // @gotags: bson:"auth0Id"
	OwnedDecks    []string               `protobuf:"bytes,4,rep,name=ownedDecks,proto3" json:"ownedDecks,omitempty" bson:"ownedDecks"` // @gotags: bson:"ownedDecks"
	OwnedCards    []string               `protobuf:"bytes,5,rep,name=ownedCards,proto3" json:"ownedCards,omitempty" bson:"ownedCards"` // @gotags: bson:"ownedCards"
	OwnedSets     []string               `protobuf:"bytes,6,rep,name=ownedSets,proto3" json:"ownedSets,omitempty" bson:"ownedSets"`   // @gotags: bson:"ownedSets"
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *User) Reset() {
	*x = User{}
	mi := &file_user_user_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_user_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetAuth0Id() string {
	if x != nil {
		return x.Auth0Id
	}
	return ""
}

func (x *User) GetOwnedDecks() []string {
	if x != nil {
		return x.OwnedDecks
	}
	return nil
}

func (x *User) GetOwnedCards() []string {
	if x != nil {
		return x.OwnedCards
	}
	return nil
}

func (x *User) GetOwnedSets() []string {
	if x != nil {
		return x.OwnedSets
	}
	return nil
}

var File_user_user_proto protoreflect.FileDescriptor

const file_user_user_proto_rawDesc = "" +
	"\n" +
	"\x0fuser/user.proto\x12\x04user\"\xb0\x01\n" +
	"\x04User\x12\x1a\n" +
	"\busername\x18\x01 \x01(\tR\busername\x12\x14\n" +
	"\x05email\x18\x02 \x01(\tR\x05email\x12\x18\n" +
	"\aauth0Id\x18\x03 \x01(\tR\aauth0Id\x12\x1e\n" +
	"\n" +
	"ownedDecks\x18\x04 \x03(\tR\n" +
	"ownedDecks\x12\x1e\n" +
	"\n" +
	"ownedCards\x18\x05 \x03(\tR\n" +
	"ownedCards\x12\x1c\n" +
	"\townedSets\x18\x06 \x03(\tR\townedSetsB+Z)github.com/stevezaluk/mtgjson-models/userb\x06proto3"

var (
	file_user_user_proto_rawDescOnce sync.Once
	file_user_user_proto_rawDescData []byte
)

func file_user_user_proto_rawDescGZIP() []byte {
	file_user_user_proto_rawDescOnce.Do(func() {
		file_user_user_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_user_user_proto_rawDesc), len(file_user_user_proto_rawDesc)))
	})
	return file_user_user_proto_rawDescData
}

var file_user_user_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_user_user_proto_goTypes = []any{
	(*User)(nil), // 0: user.User
}
var file_user_user_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_user_user_proto_init() }
func file_user_user_proto_init() {
	if File_user_user_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_user_user_proto_rawDesc), len(file_user_user_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_user_user_proto_goTypes,
		DependencyIndexes: file_user_user_proto_depIdxs,
		MessageInfos:      file_user_user_proto_msgTypes,
	}.Build()
	File_user_user_proto = out.File
	file_user_user_proto_goTypes = nil
	file_user_user_proto_depIdxs = nil
}
