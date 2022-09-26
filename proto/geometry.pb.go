// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: geometry.proto

package proto

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

type RectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Height float32 `protobuf:"fixed32,1,opt,name=height,proto3" json:"height,omitempty"`
	Width  float32 `protobuf:"fixed32,2,opt,name=width,proto3" json:"width,omitempty"`
}

func (x *RectRequest) Reset() {
	*x = RectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_geometry_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RectRequest) ProtoMessage() {}

func (x *RectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_geometry_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RectRequest.ProtoReflect.Descriptor instead.
func (*RectRequest) Descriptor() ([]byte, []int) {
	return file_geometry_proto_rawDescGZIP(), []int{0}
}

func (x *RectRequest) GetHeight() float32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *RectRequest) GetWidth() float32 {
	if x != nil {
		return x.Width
	}
	return 0
}

type AreaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result float32 `protobuf:"fixed32,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *AreaResponse) Reset() {
	*x = AreaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_geometry_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AreaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AreaResponse) ProtoMessage() {}

func (x *AreaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_geometry_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AreaResponse.ProtoReflect.Descriptor instead.
func (*AreaResponse) Descriptor() ([]byte, []int) {
	return file_geometry_proto_rawDescGZIP(), []int{1}
}

func (x *AreaResponse) GetResult() float32 {
	if x != nil {
		return x.Result
	}
	return 0
}

type PermiterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result float32 `protobuf:"fixed32,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *PermiterResponse) Reset() {
	*x = PermiterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_geometry_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PermiterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermiterResponse) ProtoMessage() {}

func (x *PermiterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_geometry_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermiterResponse.ProtoReflect.Descriptor instead.
func (*PermiterResponse) Descriptor() ([]byte, []int) {
	return file_geometry_proto_rawDescGZIP(), []int{2}
}

func (x *PermiterResponse) GetResult() float32 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_geometry_proto protoreflect.FileDescriptor

var file_geometry_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x67, 0x65, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x67, 0x65, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x22, 0x3b, 0x0a, 0x0b, 0x52, 0x65,
	0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x22, 0x26, 0x0a, 0x0c, 0x41, 0x72, 0x65, 0x61, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22,
	0x2a, 0x0a, 0x10, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x88, 0x01, 0x0a, 0x0f,
	0x47, 0x65, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x35, 0x0a, 0x04, 0x41, 0x72, 0x65, 0x61, 0x12, 0x15, 0x2e, 0x67, 0x65, 0x6f, 0x6d, 0x65, 0x74,
	0x72, 0x79, 0x2e, 0x52, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x65, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x41, 0x72, 0x65, 0x61, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x09, 0x50, 0x65, 0x72, 0x69, 0x6d, 0x65,
	0x74, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x67, 0x65, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x52,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x67, 0x65, 0x6f,
	0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x79, 0x2d, 0x73, 0x61, 0x62, 0x62, 0x69, 0x72, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_geometry_proto_rawDescOnce sync.Once
	file_geometry_proto_rawDescData = file_geometry_proto_rawDesc
)

func file_geometry_proto_rawDescGZIP() []byte {
	file_geometry_proto_rawDescOnce.Do(func() {
		file_geometry_proto_rawDescData = protoimpl.X.CompressGZIP(file_geometry_proto_rawDescData)
	})
	return file_geometry_proto_rawDescData
}

var file_geometry_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_geometry_proto_goTypes = []interface{}{
	(*RectRequest)(nil),      // 0: geometry.RectRequest
	(*AreaResponse)(nil),     // 1: geometry.AreaResponse
	(*PermiterResponse)(nil), // 2: geometry.PermiterResponse
}
var file_geometry_proto_depIdxs = []int32{
	0, // 0: geometry.GeometryService.Area:input_type -> geometry.RectRequest
	0, // 1: geometry.GeometryService.Perimeter:input_type -> geometry.RectRequest
	1, // 2: geometry.GeometryService.Area:output_type -> geometry.AreaResponse
	2, // 3: geometry.GeometryService.Perimeter:output_type -> geometry.PermiterResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_geometry_proto_init() }
func file_geometry_proto_init() {
	if File_geometry_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_geometry_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RectRequest); i {
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
		file_geometry_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AreaResponse); i {
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
		file_geometry_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PermiterResponse); i {
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
			RawDescriptor: file_geometry_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_geometry_proto_goTypes,
		DependencyIndexes: file_geometry_proto_depIdxs,
		MessageInfos:      file_geometry_proto_msgTypes,
	}.Build()
	File_geometry_proto = out.File
	file_geometry_proto_rawDesc = nil
	file_geometry_proto_goTypes = nil
	file_geometry_proto_depIdxs = nil
}
