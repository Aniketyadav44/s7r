// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: getter.proto

package getter

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type UrlItem struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId        int64                  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Url           string                 `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	Short         string                 `protobuf:"bytes,4,opt,name=short,proto3" json:"short,omitempty"`
	Clicks        int64                  `protobuf:"varint,5,opt,name=clicks,proto3" json:"clicks,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UrlItem) Reset() {
	*x = UrlItem{}
	mi := &file_getter_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UrlItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UrlItem) ProtoMessage() {}

func (x *UrlItem) ProtoReflect() protoreflect.Message {
	mi := &file_getter_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UrlItem.ProtoReflect.Descriptor instead.
func (*UrlItem) Descriptor() ([]byte, []int) {
	return file_getter_proto_rawDescGZIP(), []int{0}
}

func (x *UrlItem) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UrlItem) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UrlItem) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *UrlItem) GetShort() string {
	if x != nil {
		return x.Short
	}
	return ""
}

func (x *UrlItem) GetClicks() int64 {
	if x != nil {
		return x.Clicks
	}
	return 0
}

func (x *UrlItem) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *UrlItem) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type AllRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        int64                  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AllRequest) Reset() {
	*x = AllRequest{}
	mi := &file_getter_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllRequest) ProtoMessage() {}

func (x *AllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_getter_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllRequest.ProtoReflect.Descriptor instead.
func (*AllRequest) Descriptor() ([]byte, []int) {
	return file_getter_proto_rawDescGZIP(), []int{1}
}

func (x *AllRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type AllResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Urls          []*UrlItem             `protobuf:"bytes,3,rep,name=urls,proto3" json:"urls,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AllResponse) Reset() {
	*x = AllResponse{}
	mi := &file_getter_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AllResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllResponse) ProtoMessage() {}

func (x *AllResponse) ProtoReflect() protoreflect.Message {
	mi := &file_getter_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllResponse.ProtoReflect.Descriptor instead.
func (*AllResponse) Descriptor() ([]byte, []int) {
	return file_getter_proto_rawDescGZIP(), []int{2}
}

func (x *AllResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *AllResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *AllResponse) GetUrls() []*UrlItem {
	if x != nil {
		return x.Urls
	}
	return nil
}

type ClickRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ShortCode     string                 `protobuf:"bytes,1,opt,name=short_code,json=shortCode,proto3" json:"short_code,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ClickRequest) Reset() {
	*x = ClickRequest{}
	mi := &file_getter_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ClickRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClickRequest) ProtoMessage() {}

func (x *ClickRequest) ProtoReflect() protoreflect.Message {
	mi := &file_getter_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClickRequest.ProtoReflect.Descriptor instead.
func (*ClickRequest) Descriptor() ([]byte, []int) {
	return file_getter_proto_rawDescGZIP(), []int{3}
}

func (x *ClickRequest) GetShortCode() string {
	if x != nil {
		return x.ShortCode
	}
	return ""
}

type ClickResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Count         int64                  `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ClickResponse) Reset() {
	*x = ClickResponse{}
	mi := &file_getter_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ClickResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClickResponse) ProtoMessage() {}

func (x *ClickResponse) ProtoReflect() protoreflect.Message {
	mi := &file_getter_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClickResponse.ProtoReflect.Descriptor instead.
func (*ClickResponse) Descriptor() ([]byte, []int) {
	return file_getter_proto_rawDescGZIP(), []int{4}
}

func (x *ClickResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *ClickResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ClickResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type HistoryRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ShortCode     string                 `protobuf:"bytes,1,opt,name=short_code,json=shortCode,proto3" json:"short_code,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HistoryRequest) Reset() {
	*x = HistoryRequest{}
	mi := &file_getter_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HistoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HistoryRequest) ProtoMessage() {}

func (x *HistoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_getter_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HistoryRequest.ProtoReflect.Descriptor instead.
func (*HistoryRequest) Descriptor() ([]byte, []int) {
	return file_getter_proto_rawDescGZIP(), []int{5}
}

func (x *HistoryRequest) GetShortCode() string {
	if x != nil {
		return x.ShortCode
	}
	return ""
}

type HistoryResponse struct {
	state         protoimpl.MessageState   `protogen:"open.v1"`
	Success       bool                     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message       string                   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	History       []*timestamppb.Timestamp `protobuf:"bytes,3,rep,name=history,proto3" json:"history,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HistoryResponse) Reset() {
	*x = HistoryResponse{}
	mi := &file_getter_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HistoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HistoryResponse) ProtoMessage() {}

func (x *HistoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_getter_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HistoryResponse.ProtoReflect.Descriptor instead.
func (*HistoryResponse) Descriptor() ([]byte, []int) {
	return file_getter_proto_rawDescGZIP(), []int{6}
}

func (x *HistoryResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *HistoryResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *HistoryResponse) GetHistory() []*timestamppb.Timestamp {
	if x != nil {
		return x.History
	}
	return nil
}

var File_getter_proto protoreflect.FileDescriptor

const file_getter_proto_rawDesc = "" +
	"\n" +
	"\fgetter.proto\x12\x0egetter_service\x1a\x1fgoogle/protobuf/timestamp.proto\"\xe8\x01\n" +
	"\aUrlItem\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x17\n" +
	"\auser_id\x18\x02 \x01(\x03R\x06userId\x12\x10\n" +
	"\x03url\x18\x03 \x01(\tR\x03url\x12\x14\n" +
	"\x05short\x18\x04 \x01(\tR\x05short\x12\x16\n" +
	"\x06clicks\x18\x05 \x01(\x03R\x06clicks\x129\n" +
	"\n" +
	"created_at\x18\x06 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x129\n" +
	"\n" +
	"updated_at\x18\a \x01(\v2\x1a.google.protobuf.TimestampR\tupdatedAt\"%\n" +
	"\n" +
	"AllRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\x03R\x06userId\"n\n" +
	"\vAllResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage\x12+\n" +
	"\x04urls\x18\x03 \x03(\v2\x17.getter_service.UrlItemR\x04urls\"-\n" +
	"\fClickRequest\x12\x1d\n" +
	"\n" +
	"short_code\x18\x01 \x01(\tR\tshortCode\"Y\n" +
	"\rClickResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage\x12\x14\n" +
	"\x05count\x18\x03 \x01(\x03R\x05count\"/\n" +
	"\x0eHistoryRequest\x12\x1d\n" +
	"\n" +
	"short_code\x18\x01 \x01(\tR\tshortCode\"{\n" +
	"\x0fHistoryResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage\x124\n" +
	"\ahistory\x18\x03 \x03(\v2\x1a.google.protobuf.TimestampR\ahistory2\xee\x01\n" +
	"\x06Getter\x12G\n" +
	"\n" +
	"GetAllUrls\x12\x1a.getter_service.AllRequest\x1a\x1b.getter_service.AllResponse\"\x00\x12J\n" +
	"\tGetClicks\x12\x1c.getter_service.ClickRequest\x1a\x1d.getter_service.ClickResponse\"\x00\x12O\n" +
	"\n" +
	"GetHistory\x12\x1e.getter_service.HistoryRequest\x1a\x1f.getter_service.HistoryResponse\"\x00B%Z#github.com/Aniketyadav44/s7r/getterb\x06proto3"

var (
	file_getter_proto_rawDescOnce sync.Once
	file_getter_proto_rawDescData []byte
)

func file_getter_proto_rawDescGZIP() []byte {
	file_getter_proto_rawDescOnce.Do(func() {
		file_getter_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_getter_proto_rawDesc), len(file_getter_proto_rawDesc)))
	})
	return file_getter_proto_rawDescData
}

var file_getter_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_getter_proto_goTypes = []any{
	(*UrlItem)(nil),               // 0: getter_service.UrlItem
	(*AllRequest)(nil),            // 1: getter_service.AllRequest
	(*AllResponse)(nil),           // 2: getter_service.AllResponse
	(*ClickRequest)(nil),          // 3: getter_service.ClickRequest
	(*ClickResponse)(nil),         // 4: getter_service.ClickResponse
	(*HistoryRequest)(nil),        // 5: getter_service.HistoryRequest
	(*HistoryResponse)(nil),       // 6: getter_service.HistoryResponse
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_getter_proto_depIdxs = []int32{
	7, // 0: getter_service.UrlItem.created_at:type_name -> google.protobuf.Timestamp
	7, // 1: getter_service.UrlItem.updated_at:type_name -> google.protobuf.Timestamp
	0, // 2: getter_service.AllResponse.urls:type_name -> getter_service.UrlItem
	7, // 3: getter_service.HistoryResponse.history:type_name -> google.protobuf.Timestamp
	1, // 4: getter_service.Getter.GetAllUrls:input_type -> getter_service.AllRequest
	3, // 5: getter_service.Getter.GetClicks:input_type -> getter_service.ClickRequest
	5, // 6: getter_service.Getter.GetHistory:input_type -> getter_service.HistoryRequest
	2, // 7: getter_service.Getter.GetAllUrls:output_type -> getter_service.AllResponse
	4, // 8: getter_service.Getter.GetClicks:output_type -> getter_service.ClickResponse
	6, // 9: getter_service.Getter.GetHistory:output_type -> getter_service.HistoryResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_getter_proto_init() }
func file_getter_proto_init() {
	if File_getter_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_getter_proto_rawDesc), len(file_getter_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_getter_proto_goTypes,
		DependencyIndexes: file_getter_proto_depIdxs,
		MessageInfos:      file_getter_proto_msgTypes,
	}.Build()
	File_getter_proto = out.File
	file_getter_proto_goTypes = nil
	file_getter_proto_depIdxs = nil
}
