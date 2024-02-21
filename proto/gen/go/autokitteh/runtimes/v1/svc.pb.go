// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: autokitteh/runtimes/v1/svc.proto

package runtimesv1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
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

type DescribeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DescribeRequest) Reset() {
	*x = DescribeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_autokitteh_runtimes_v1_svc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeRequest) ProtoMessage() {}

func (x *DescribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_autokitteh_runtimes_v1_svc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeRequest.ProtoReflect.Descriptor instead.
func (*DescribeRequest) Descriptor() ([]byte, []int) {
	return file_autokitteh_runtimes_v1_svc_proto_rawDescGZIP(), []int{0}
}

func (x *DescribeRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DescribeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Runtime *Runtime `protobuf:"bytes,1,opt,name=runtime,proto3" json:"runtime,omitempty"` // empty if not found.
}

func (x *DescribeResponse) Reset() {
	*x = DescribeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_autokitteh_runtimes_v1_svc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeResponse) ProtoMessage() {}

func (x *DescribeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_autokitteh_runtimes_v1_svc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeResponse.ProtoReflect.Descriptor instead.
func (*DescribeResponse) Descriptor() ([]byte, []int) {
	return file_autokitteh_runtimes_v1_svc_proto_rawDescGZIP(), []int{1}
}

func (x *DescribeResponse) GetRuntime() *Runtime {
	if x != nil {
		return x.Runtime
	}
	return nil
}

type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_autokitteh_runtimes_v1_svc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_autokitteh_runtimes_v1_svc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_autokitteh_runtimes_v1_svc_proto_rawDescGZIP(), []int{2}
}

type ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Runtimes []*Runtime `protobuf:"bytes,1,rep,name=runtimes,proto3" json:"runtimes,omitempty"`
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_autokitteh_runtimes_v1_svc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_autokitteh_runtimes_v1_svc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_autokitteh_runtimes_v1_svc_proto_rawDescGZIP(), []int{3}
}

func (x *ListResponse) GetRuntimes() []*Runtime {
	if x != nil {
		return x.Runtimes
	}
	return nil
}

var File_autokitteh_runtimes_v1_svc_proto protoreflect.FileDescriptor

var file_autokitteh_runtimes_v1_svc_proto_rawDesc = []byte{
	0x0a, 0x20, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x2f, 0x72, 0x75, 0x6e,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x76, 0x63, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x16, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x2e, 0x72,
	0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x24, 0x61, 0x75, 0x74, 0x6f,
	0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2f, 0x0a,
	0x0f, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1c, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08,
	0xfa, 0xf7, 0x18, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x4d,
	0x0a, 0x10, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x39, 0x0a, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68,
	0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x6e,
	0x74, 0x69, 0x6d, 0x65, 0x52, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x0d, 0x0a,
	0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x54, 0x0a, 0x0c,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x08,
	0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x2e, 0x72, 0x75, 0x6e, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x42,
	0x07, 0xfa, 0xf7, 0x18, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x08, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x32, 0xc3, 0x01, 0x0a, 0x0f, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5d, 0x0a, 0x08, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x12, 0x27, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x2e,
	0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x61, 0x75,
	0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x23, 0x2e,
	0x61, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x24, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x2e,
	0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xed, 0x01, 0x0a, 0x1a, 0x63, 0x6f, 0x6d,
	0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x2e, 0x72, 0x75, 0x6e, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x08, 0x53, 0x76, 0x63, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x4b, 0x67, 0x6f, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74,
	0x65, 0x68, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65,
	0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x61,
	0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x76, 0x31,
	0xa2, 0x02, 0x03, 0x41, 0x52, 0x58, 0xaa, 0x02, 0x16, 0x41, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74,
	0x74, 0x65, 0x68, 0x2e, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x2e, 0x56, 0x31, 0xca,
	0x02, 0x16, 0x41, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x5c, 0x52, 0x75, 0x6e,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x22, 0x41, 0x75, 0x74, 0x6f, 0x6b,
	0x69, 0x74, 0x74, 0x65, 0x68, 0x5c, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x5c, 0x56,
	0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x18,
	0x41, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x3a, 0x3a, 0x52, 0x75, 0x6e, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_autokitteh_runtimes_v1_svc_proto_rawDescOnce sync.Once
	file_autokitteh_runtimes_v1_svc_proto_rawDescData = file_autokitteh_runtimes_v1_svc_proto_rawDesc
)

func file_autokitteh_runtimes_v1_svc_proto_rawDescGZIP() []byte {
	file_autokitteh_runtimes_v1_svc_proto_rawDescOnce.Do(func() {
		file_autokitteh_runtimes_v1_svc_proto_rawDescData = protoimpl.X.CompressGZIP(file_autokitteh_runtimes_v1_svc_proto_rawDescData)
	})
	return file_autokitteh_runtimes_v1_svc_proto_rawDescData
}

var file_autokitteh_runtimes_v1_svc_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_autokitteh_runtimes_v1_svc_proto_goTypes = []interface{}{
	(*DescribeRequest)(nil),  // 0: autokitteh.runtimes.v1.DescribeRequest
	(*DescribeResponse)(nil), // 1: autokitteh.runtimes.v1.DescribeResponse
	(*ListRequest)(nil),      // 2: autokitteh.runtimes.v1.ListRequest
	(*ListResponse)(nil),     // 3: autokitteh.runtimes.v1.ListResponse
	(*Runtime)(nil),          // 4: autokitteh.runtimes.v1.Runtime
}
var file_autokitteh_runtimes_v1_svc_proto_depIdxs = []int32{
	4, // 0: autokitteh.runtimes.v1.DescribeResponse.runtime:type_name -> autokitteh.runtimes.v1.Runtime
	4, // 1: autokitteh.runtimes.v1.ListResponse.runtimes:type_name -> autokitteh.runtimes.v1.Runtime
	0, // 2: autokitteh.runtimes.v1.RuntimesService.Describe:input_type -> autokitteh.runtimes.v1.DescribeRequest
	2, // 3: autokitteh.runtimes.v1.RuntimesService.List:input_type -> autokitteh.runtimes.v1.ListRequest
	1, // 4: autokitteh.runtimes.v1.RuntimesService.Describe:output_type -> autokitteh.runtimes.v1.DescribeResponse
	3, // 5: autokitteh.runtimes.v1.RuntimesService.List:output_type -> autokitteh.runtimes.v1.ListResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_autokitteh_runtimes_v1_svc_proto_init() }
func file_autokitteh_runtimes_v1_svc_proto_init() {
	if File_autokitteh_runtimes_v1_svc_proto != nil {
		return
	}
	file_autokitteh_runtimes_v1_runtime_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_autokitteh_runtimes_v1_svc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeRequest); i {
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
		file_autokitteh_runtimes_v1_svc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeResponse); i {
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
		file_autokitteh_runtimes_v1_svc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequest); i {
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
		file_autokitteh_runtimes_v1_svc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResponse); i {
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
			RawDescriptor: file_autokitteh_runtimes_v1_svc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_autokitteh_runtimes_v1_svc_proto_goTypes,
		DependencyIndexes: file_autokitteh_runtimes_v1_svc_proto_depIdxs,
		MessageInfos:      file_autokitteh_runtimes_v1_svc_proto_msgTypes,
	}.Build()
	File_autokitteh_runtimes_v1_svc_proto = out.File
	file_autokitteh_runtimes_v1_svc_proto_rawDesc = nil
	file_autokitteh_runtimes_v1_svc_proto_goTypes = nil
	file_autokitteh_runtimes_v1_svc_proto_depIdxs = nil
}
