// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: fana.proto

package pb

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

type ID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID int32 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *ID) Reset() {
	*x = ID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fana_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ID) ProtoMessage() {}

func (x *ID) ProtoReflect() protoreflect.Message {
	mi := &file_fana_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ID.ProtoReflect.Descriptor instead.
func (*ID) Descriptor() ([]byte, []int) {
	return file_fana_proto_rawDescGZIP(), []int{0}
}

func (x *ID) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

var File_fana_proto protoreflect.FileDescriptor

var file_fana_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x66, 0x61, 0x6e, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x14, 0x0a,
	0x02, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x49, 0x44, 0x32, 0x8b, 0x01, 0x0a, 0x04, 0x46, 0x61, 0x6e, 0x61, 0x12, 0x1d, 0x0a, 0x07,
	0x47, 0x65, 0x74, 0x46, 0x6c, 0x61, 0x67, 0x12, 0x03, 0x2e, 0x49, 0x44, 0x1a, 0x0d, 0x2e, 0x46,
	0x6c, 0x61, 0x67, 0x46, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x47,
	0x65, 0x74, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x06, 0x2e, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41, 0x75,
	0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x0a, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x47,
	0x65, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x06, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0b, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x73, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_fana_proto_rawDescOnce sync.Once
	file_fana_proto_rawDescData = file_fana_proto_rawDesc
)

func file_fana_proto_rawDescGZIP() []byte {
	file_fana_proto_rawDescOnce.Do(func() {
		file_fana_proto_rawDescData = protoimpl.X.CompressGZIP(file_fana_proto_rawDescData)
	})
	return file_fana_proto_rawDescData
}

var file_fana_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_fana_proto_goTypes = []interface{}{
	(*ID)(nil),           // 0: ID
	(*Empty)(nil),        // 1: Empty
	(*FlagFullResp)(nil), // 2: FlagFullResp
	(*Flags)(nil),        // 3: Flags
	(*Audiences)(nil),    // 4: Audiences
	(*Attributes)(nil),   // 5: Attributes
}
var file_fana_proto_depIdxs = []int32{
	0, // 0: Fana.GetFlag:input_type -> ID
	1, // 1: Fana.GetFlags:input_type -> Empty
	1, // 2: Fana.GetAudiences:input_type -> Empty
	1, // 3: Fana.GetAttributes:input_type -> Empty
	2, // 4: Fana.GetFlag:output_type -> FlagFullResp
	3, // 5: Fana.GetFlags:output_type -> Flags
	4, // 6: Fana.GetAudiences:output_type -> Audiences
	5, // 7: Fana.GetAttributes:output_type -> Attributes
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_fana_proto_init() }
func file_fana_proto_init() {
	if File_fana_proto != nil {
		return
	}
	file_responses_proto_init()
	file_requests_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_fana_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ID); i {
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
			RawDescriptor: file_fana_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_fana_proto_goTypes,
		DependencyIndexes: file_fana_proto_depIdxs,
		MessageInfos:      file_fana_proto_msgTypes,
	}.Build()
	File_fana_proto = out.File
	file_fana_proto_rawDesc = nil
	file_fana_proto_goTypes = nil
	file_fana_proto_depIdxs = nil
}
