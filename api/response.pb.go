// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.6.1
// source: api/response.proto

package api

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

// APIResponse - Represents a generic API response that can get returned from the API.
// Some fields are not used and will be nulled out on response
type APIResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message       string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty" bson:"message"`             // @gotags: bson:"message"
	Err           string   `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty" bson:"err"`                     // @gotags: bson:"err"
	RequiredScope string   `protobuf:"bytes,3,opt,name=requiredScope,proto3" json:"requiredScope,omitempty" bson:"requiredScope"` // @gotags: bson:"requiredScope"
	NoExistCards  []string `protobuf:"bytes,4,rep,name=noExistCards,proto3" json:"noExistCards,omitempty" bson:"noExistCards"`   // @gotags: bson:"noExistCards"
	InvalidCards  []string `protobuf:"bytes,5,rep,name=invalidCards,proto3" json:"invalidCards,omitempty" bson:"invalidCards"`   // @gotags: bson:"invalidCards"
	DeckCode      string   `protobuf:"bytes,6,opt,name=deckCode,proto3" json:"deckCode,omitempty" bson:"deckCode"`           // @gotags: bson:"deckCode"
	SetCode       string   `protobuf:"bytes,7,opt,name=setCode,proto3" json:"setCode,omitempty" bson:"setCode"`             // @gotags: bson:"setCode"
}

func (x *APIResponse) Reset() {
	*x = APIResponse{}
	mi := &file_api_response_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *APIResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*APIResponse) ProtoMessage() {}

func (x *APIResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_response_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use APIResponse.ProtoReflect.Descriptor instead.
func (*APIResponse) Descriptor() ([]byte, []int) {
	return file_api_response_proto_rawDescGZIP(), []int{0}
}

func (x *APIResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *APIResponse) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

func (x *APIResponse) GetRequiredScope() string {
	if x != nil {
		return x.RequiredScope
	}
	return ""
}

func (x *APIResponse) GetNoExistCards() []string {
	if x != nil {
		return x.NoExistCards
	}
	return nil
}

func (x *APIResponse) GetInvalidCards() []string {
	if x != nil {
		return x.InvalidCards
	}
	return nil
}

func (x *APIResponse) GetDeckCode() string {
	if x != nil {
		return x.DeckCode
	}
	return ""
}

func (x *APIResponse) GetSetCode() string {
	if x != nil {
		return x.SetCode
	}
	return ""
}

var File_api_response_proto protoreflect.FileDescriptor

var file_api_response_proto_rawDesc = []byte{
	0x0a, 0x12, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x22, 0xdd, 0x01, 0x0a, 0x0b, 0x41, 0x50,
	0x49, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x65, 0x72, 0x72, 0x12, 0x24, 0x0a, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65,
	0x64, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65,
	0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x6e,
	0x6f, 0x45, 0x78, 0x69, 0x73, 0x74, 0x43, 0x61, 0x72, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0c, 0x6e, 0x6f, 0x45, 0x78, 0x69, 0x73, 0x74, 0x43, 0x61, 0x72, 0x64, 0x73, 0x12,
	0x22, 0x0a, 0x0c, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x43, 0x61, 0x72, 0x64, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x43, 0x61,
	0x72, 0x64, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x63, 0x6b, 0x43, 0x6f, 0x64, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x63, 0x6b, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x73, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x65, 0x76, 0x65, 0x7a, 0x61, 0x6c,
	0x75, 0x6b, 0x2f, 0x6d, 0x74, 0x67, 0x6a, 0x73, 0x6f, 0x6e, 0x2d, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_response_proto_rawDescOnce sync.Once
	file_api_response_proto_rawDescData = file_api_response_proto_rawDesc
)

func file_api_response_proto_rawDescGZIP() []byte {
	file_api_response_proto_rawDescOnce.Do(func() {
		file_api_response_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_response_proto_rawDescData)
	})
	return file_api_response_proto_rawDescData
}

var file_api_response_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_response_proto_goTypes = []any{
	(*APIResponse)(nil), // 0: api.APIResponse
}
var file_api_response_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_response_proto_init() }
func file_api_response_proto_init() {
	if File_api_response_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_response_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_response_proto_goTypes,
		DependencyIndexes: file_api_response_proto_depIdxs,
		MessageInfos:      file_api_response_proto_msgTypes,
	}.Build()
	File_api_response_proto = out.File
	file_api_response_proto_rawDesc = nil
	file_api_response_proto_goTypes = nil
	file_api_response_proto_depIdxs = nil
}
