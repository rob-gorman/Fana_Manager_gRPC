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
	0x02, 0x49, 0x44, 0x32, 0xfd, 0x04, 0x0a, 0x04, 0x46, 0x61, 0x6e, 0x61, 0x12, 0x1d, 0x0a, 0x07,
	0x47, 0x65, 0x74, 0x46, 0x6c, 0x61, 0x67, 0x12, 0x03, 0x2e, 0x49, 0x44, 0x1a, 0x0d, 0x2e, 0x46,
	0x6c, 0x61, 0x67, 0x46, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x47,
	0x65, 0x74, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x06, 0x2e, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x12, 0x25, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x75,
	0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x03, 0x2e, 0x49, 0x44, 0x1a, 0x11, 0x2e, 0x41, 0x75,
	0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x46, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x12, 0x22,
	0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x06,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0a, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63,
	0x65, 0x73, 0x12, 0x23, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x12, 0x03, 0x2e, 0x49, 0x44, 0x1a, 0x0e, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x24, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x0b, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x1e, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x53, 0x44, 0x4b, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x06, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x08, 0x2e, 0x53, 0x44, 0x4b, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x25, 0x0a,
	0x0c, 0x47, 0x65, 0x74, 0x41, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x12, 0x06, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0d, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x28, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x6c,
	0x61, 0x67, 0x12, 0x0b, 0x2e, 0x46, 0x6c, 0x61, 0x67, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x1a,
	0x0d, 0x2e, 0x46, 0x6c, 0x61, 0x67, 0x46, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2f,
	0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65,
	0x12, 0x0a, 0x2e, 0x41, 0x75, 0x64, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x1a, 0x11, 0x2e, 0x41,
	0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x46, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x2e, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x12, 0x0b, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x1a,
	0x0e, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x28, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x6c, 0x61, 0x67, 0x12, 0x0b, 0x2e,
	0x46, 0x6c, 0x61, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x1a, 0x0d, 0x2e, 0x46, 0x6c, 0x61,
	0x67, 0x46, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2f, 0x0a, 0x0e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x41, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x0a, 0x2e, 0x41, 0x75,
	0x64, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x1a, 0x11, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x65, 0x6e,
	0x63, 0x65, 0x46, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1d, 0x0a, 0x0d, 0x52, 0x65,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x53, 0x44, 0x4b, 0x12, 0x03, 0x2e, 0x49, 0x44,
	0x1a, 0x07, 0x2e, 0x53, 0x44, 0x4b, 0x4b, 0x65, 0x79, 0x12, 0x19, 0x0a, 0x0a, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x46, 0x6c, 0x61, 0x67, 0x12, 0x03, 0x2e, 0x49, 0x44, 0x1a, 0x06, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x12, 0x1d, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x75,
	0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x03, 0x2e, 0x49, 0x44, 0x1a, 0x06, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x12, 0x1e, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x12, 0x03, 0x2e, 0x49, 0x44, 0x1a, 0x06, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
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
	(*ID)(nil),               // 0: ID
	(*Empty)(nil),            // 1: Empty
	(*FlagSubmit)(nil),       // 2: FlagSubmit
	(*AudSubmit)(nil),        // 3: AudSubmit
	(*AttrSubmit)(nil),       // 4: AttrSubmit
	(*FlagUpdate)(nil),       // 5: FlagUpdate
	(*AudUpdate)(nil),        // 6: AudUpdate
	(*FlagFullResp)(nil),     // 7: FlagFullResp
	(*Flags)(nil),            // 8: Flags
	(*AudienceFullResp)(nil), // 9: AudienceFullResp
	(*Audiences)(nil),        // 10: Audiences
	(*AttributeResp)(nil),    // 11: AttributeResp
	(*Attributes)(nil),       // 12: Attributes
	(*SDKKeys)(nil),          // 13: SDKKeys
	(*AuditLogResp)(nil),     // 14: AuditLogResp
	(*SDKKey)(nil),           // 15: SDKKey
}
var file_fana_proto_depIdxs = []int32{
	0,  // 0: Fana.GetFlag:input_type -> ID
	1,  // 1: Fana.GetFlags:input_type -> Empty
	0,  // 2: Fana.GetAudience:input_type -> ID
	1,  // 3: Fana.GetAudiences:input_type -> Empty
	0,  // 4: Fana.GetAttribute:input_type -> ID
	1,  // 5: Fana.GetAttributes:input_type -> Empty
	1,  // 6: Fana.GetSDKKeys:input_type -> Empty
	1,  // 7: Fana.GetAuditLogs:input_type -> Empty
	2,  // 8: Fana.CreateFlag:input_type -> FlagSubmit
	3,  // 9: Fana.CreateAudience:input_type -> AudSubmit
	4,  // 10: Fana.CreateAttribute:input_type -> AttrSubmit
	5,  // 11: Fana.UpdateFlag:input_type -> FlagUpdate
	6,  // 12: Fana.UpdateAudience:input_type -> AudUpdate
	0,  // 13: Fana.RegenerateSDK:input_type -> ID
	0,  // 14: Fana.DeleteFlag:input_type -> ID
	0,  // 15: Fana.DeleteAudience:input_type -> ID
	0,  // 16: Fana.DeleteAttribute:input_type -> ID
	7,  // 17: Fana.GetFlag:output_type -> FlagFullResp
	8,  // 18: Fana.GetFlags:output_type -> Flags
	9,  // 19: Fana.GetAudience:output_type -> AudienceFullResp
	10, // 20: Fana.GetAudiences:output_type -> Audiences
	11, // 21: Fana.GetAttribute:output_type -> AttributeResp
	12, // 22: Fana.GetAttributes:output_type -> Attributes
	13, // 23: Fana.GetSDKKeys:output_type -> SDKKeys
	14, // 24: Fana.GetAuditLogs:output_type -> AuditLogResp
	7,  // 25: Fana.CreateFlag:output_type -> FlagFullResp
	9,  // 26: Fana.CreateAudience:output_type -> AudienceFullResp
	11, // 27: Fana.CreateAttribute:output_type -> AttributeResp
	7,  // 28: Fana.UpdateFlag:output_type -> FlagFullResp
	9,  // 29: Fana.UpdateAudience:output_type -> AudienceFullResp
	15, // 30: Fana.RegenerateSDK:output_type -> SDKKey
	1,  // 31: Fana.DeleteFlag:output_type -> Empty
	1,  // 32: Fana.DeleteAudience:output_type -> Empty
	1,  // 33: Fana.DeleteAttribute:output_type -> Empty
	17, // [17:34] is the sub-list for method output_type
	0,  // [0:17] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
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
