// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: internal/service/pb/wikitable.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Row struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Columns map[int64]string `protobuf:"bytes,1,rep,name=columns,proto3" json:"columns,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Row) Reset() {
	*x = Row{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_service_pb_wikitable_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Row) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Row) ProtoMessage() {}

func (x *Row) ProtoReflect() protoreflect.Message {
	mi := &file_internal_service_pb_wikitable_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Row.ProtoReflect.Descriptor instead.
func (*Row) Descriptor() ([]byte, []int) {
	return file_internal_service_pb_wikitable_proto_rawDescGZIP(), []int{0}
}

func (x *Row) GetColumns() map[int64]string {
	if x != nil {
		return x.Columns
	}
	return nil
}

type Table struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Caption string         `protobuf:"bytes,1,opt,name=caption,proto3" json:"caption,omitempty"`
	Rows    map[int64]*Row `protobuf:"bytes,2,rep,name=rows,proto3" json:"rows,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Table) Reset() {
	*x = Table{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_service_pb_wikitable_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Table) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Table) ProtoMessage() {}

func (x *Table) ProtoReflect() protoreflect.Message {
	mi := &file_internal_service_pb_wikitable_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Table.ProtoReflect.Descriptor instead.
func (*Table) Descriptor() ([]byte, []int) {
	return file_internal_service_pb_wikitable_proto_rawDescGZIP(), []int{1}
}

func (x *Table) GetCaption() string {
	if x != nil {
		return x.Caption
	}
	return ""
}

func (x *Table) GetRows() map[int64]*Row {
	if x != nil {
		return x.Rows
	}
	return nil
}

type TablesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page  string  `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	Lang  string  `protobuf:"bytes,2,opt,name=lang,proto3" json:"lang,omitempty"`
	Table []int32 `protobuf:"varint,3,rep,packed,name=table,proto3" json:"table,omitempty"`
}

func (x *TablesRequest) Reset() {
	*x = TablesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_service_pb_wikitable_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TablesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TablesRequest) ProtoMessage() {}

func (x *TablesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_service_pb_wikitable_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TablesRequest.ProtoReflect.Descriptor instead.
func (*TablesRequest) Descriptor() ([]byte, []int) {
	return file_internal_service_pb_wikitable_proto_rawDescGZIP(), []int{2}
}

func (x *TablesRequest) GetPage() string {
	if x != nil {
		return x.Page
	}
	return ""
}

func (x *TablesRequest) GetLang() string {
	if x != nil {
		return x.Lang
	}
	return ""
}

func (x *TablesRequest) GetTable() []int32 {
	if x != nil {
		return x.Table
	}
	return nil
}

type TablesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tables []*Table `protobuf:"bytes,1,rep,name=tables,proto3" json:"tables,omitempty"`
}

func (x *TablesResponse) Reset() {
	*x = TablesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_service_pb_wikitable_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TablesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TablesResponse) ProtoMessage() {}

func (x *TablesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_service_pb_wikitable_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TablesResponse.ProtoReflect.Descriptor instead.
func (*TablesResponse) Descriptor() ([]byte, []int) {
	return file_internal_service_pb_wikitable_proto_rawDescGZIP(), []int{3}
}

func (x *TablesResponse) GetTables() []*Table {
	if x != nil {
		return x.Tables
	}
	return nil
}

var File_internal_service_pb_wikitable_proto protoreflect.FileDescriptor

var file_internal_service_pb_wikitable_proto_rawDesc = []byte{
	0x0a, 0x23, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x77, 0x69, 0x6b, 0x69, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x77, 0x69, 0x6b, 0x69, 0x70, 0x65, 0x64, 0x69, 0x61,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61,
	0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x78,
	0x0a, 0x03, 0x52, 0x6f, 0x77, 0x12, 0x35, 0x0a, 0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x77, 0x69, 0x6b, 0x69, 0x70, 0x65, 0x64,
	0x69, 0x61, 0x2e, 0x52, 0x6f, 0x77, 0x2e, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x1a, 0x3a, 0x0a, 0x0c,
	0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x9a, 0x01, 0x0a, 0x05, 0x54, 0x61, 0x62,
	0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x61, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x61, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x04,
	0x72, 0x6f, 0x77, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x77, 0x69, 0x6b,
	0x69, 0x70, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x52, 0x6f, 0x77,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x1a, 0x47, 0x0a, 0x09,
	0x52, 0x6f, 0x77, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x24, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x77, 0x69, 0x6b,
	0x69, 0x70, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x52, 0x6f, 0x77, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x4d, 0x0a, 0x0d, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x61,
	0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x05, 0x52, 0x05, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x22, 0x3a, 0x0a, 0x0e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x06, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x77, 0x69, 0x6b, 0x69, 0x70, 0x65, 0x64,
	0x69, 0x61, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x06, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73,
	0x32, 0x71, 0x0a, 0x15, 0x57, 0x69, 0x6b, 0x69, 0x70, 0x65, 0x64, 0x69, 0x61, 0x54, 0x61, 0x62,
	0x6c, 0x65, 0x4a, 0x53, 0x4f, 0x4e, 0x41, 0x50, 0x49, 0x12, 0x58, 0x0a, 0x09, 0x47, 0x65, 0x74,
	0x54, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x12, 0x18, 0x2e, 0x77, 0x69, 0x6b, 0x69, 0x70, 0x65, 0x64,
	0x69, 0x61, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x77, 0x69, 0x6b, 0x69, 0x70, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x54, 0x61, 0x62,
	0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x10, 0x12, 0x0e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x7b, 0x70, 0x61, 0x67, 0x65, 0x3d,
	0x2a, 0x2a, 0x7d, 0x42, 0x4a, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x61, 0x74, 0x79, 0x65, 0x2f, 0x77, 0x69, 0x6b, 0x69, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x2d, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x62,
	0x92, 0x41, 0x1c, 0x12, 0x1a, 0x0a, 0x18, 0x57, 0x69, 0x6b, 0x69, 0x70, 0x65, 0x64, 0x69, 0x61,
	0x20, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x20, 0x4a, 0x53, 0x4f, 0x4e, 0x20, 0x41, 0x50, 0x49, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_service_pb_wikitable_proto_rawDescOnce sync.Once
	file_internal_service_pb_wikitable_proto_rawDescData = file_internal_service_pb_wikitable_proto_rawDesc
)

func file_internal_service_pb_wikitable_proto_rawDescGZIP() []byte {
	file_internal_service_pb_wikitable_proto_rawDescOnce.Do(func() {
		file_internal_service_pb_wikitable_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_service_pb_wikitable_proto_rawDescData)
	})
	return file_internal_service_pb_wikitable_proto_rawDescData
}

var file_internal_service_pb_wikitable_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_internal_service_pb_wikitable_proto_goTypes = []interface{}{
	(*Row)(nil),            // 0: wikipedia.Row
	(*Table)(nil),          // 1: wikipedia.Table
	(*TablesRequest)(nil),  // 2: wikipedia.TablesRequest
	(*TablesResponse)(nil), // 3: wikipedia.TablesResponse
	nil,                    // 4: wikipedia.Row.ColumnsEntry
	nil,                    // 5: wikipedia.Table.RowsEntry
}
var file_internal_service_pb_wikitable_proto_depIdxs = []int32{
	4, // 0: wikipedia.Row.columns:type_name -> wikipedia.Row.ColumnsEntry
	5, // 1: wikipedia.Table.rows:type_name -> wikipedia.Table.RowsEntry
	1, // 2: wikipedia.TablesResponse.tables:type_name -> wikipedia.Table
	0, // 3: wikipedia.Table.RowsEntry.value:type_name -> wikipedia.Row
	2, // 4: wikipedia.WikipediaTableJSONAPI.GetTables:input_type -> wikipedia.TablesRequest
	3, // 5: wikipedia.WikipediaTableJSONAPI.GetTables:output_type -> wikipedia.TablesResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_internal_service_pb_wikitable_proto_init() }
func file_internal_service_pb_wikitable_proto_init() {
	if File_internal_service_pb_wikitable_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_service_pb_wikitable_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Row); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_service_pb_wikitable_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Table); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_service_pb_wikitable_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TablesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_service_pb_wikitable_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TablesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_service_pb_wikitable_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_service_pb_wikitable_proto_goTypes,
		DependencyIndexes: file_internal_service_pb_wikitable_proto_depIdxs,
		MessageInfos:      file_internal_service_pb_wikitable_proto_msgTypes,
	}.Build()
	File_internal_service_pb_wikitable_proto = out.File
	file_internal_service_pb_wikitable_proto_rawDesc = nil
	file_internal_service_pb_wikitable_proto_goTypes = nil
	file_internal_service_pb_wikitable_proto_depIdxs = nil
}
