// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: api/proto/mock.proto

package mock

import (
	any1 "github.com/golang/protobuf/ptypes/any"
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

type MockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Method string    `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	Body   *any1.Any `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	Sleep  int32     `protobuf:"varint,3,opt,name=sleep,proto3" json:"sleep,omitempty"`
}

func (x *MockRequest) Reset() {
	*x = MockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_mock_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MockRequest) ProtoMessage() {}

func (x *MockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_mock_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MockRequest.ProtoReflect.Descriptor instead.
func (*MockRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_mock_proto_rawDescGZIP(), []int{0}
}

func (x *MockRequest) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *MockRequest) GetBody() *any1.Any {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *MockRequest) GetSleep() int32 {
	if x != nil {
		return x.Sleep
	}
	return 0
}

type MockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body *any1.Any `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *MockResponse) Reset() {
	*x = MockResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_mock_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MockResponse) ProtoMessage() {}

func (x *MockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_mock_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MockResponse.ProtoReflect.Descriptor instead.
func (*MockResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_mock_proto_rawDescGZIP(), []int{1}
}

func (x *MockResponse) GetBody() *any1.Any {
	if x != nil {
		return x.Body
	}
	return nil
}

var File_api_proto_mock_proto protoreflect.FileDescriptor

var file_api_proto_mock_proto_rawDesc = []byte{
	0x0a, 0x14, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x6f, 0x63, 0x6b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6d, 0x6f, 0x63, 0x6b, 0x1a, 0x19, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x65, 0x0a, 0x0b, 0x4d, 0x6f, 0x63, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x28,
	0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41,
	0x6e, 0x79, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x6c, 0x65, 0x65,
	0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x6c, 0x65, 0x65, 0x70, 0x22, 0x38,
	0x0a, 0x0c, 0x4d, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28,
	0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41,
	0x6e, 0x79, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x32, 0x3e, 0x0a, 0x0b, 0x4d, 0x6f, 0x63, 0x6b,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x4d, 0x6f, 0x63, 0x6b, 0x12,
	0x11, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x2e, 0x4d, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x12, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x2e, 0x4d, 0x6f, 0x63, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x61, 0x70, 0x69, 0x2f,
	0x6d, 0x6f, 0x63, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_mock_proto_rawDescOnce sync.Once
	file_api_proto_mock_proto_rawDescData = file_api_proto_mock_proto_rawDesc
)

func file_api_proto_mock_proto_rawDescGZIP() []byte {
	file_api_proto_mock_proto_rawDescOnce.Do(func() {
		file_api_proto_mock_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_mock_proto_rawDescData)
	})
	return file_api_proto_mock_proto_rawDescData
}

var file_api_proto_mock_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_proto_mock_proto_goTypes = []interface{}{
	(*MockRequest)(nil),  // 0: mock.MockRequest
	(*MockResponse)(nil), // 1: mock.MockResponse
	(*any1.Any)(nil),     // 2: google.protobuf.Any
}
var file_api_proto_mock_proto_depIdxs = []int32{
	2, // 0: mock.MockRequest.body:type_name -> google.protobuf.Any
	2, // 1: mock.MockResponse.body:type_name -> google.protobuf.Any
	0, // 2: mock.MockService.Mock:input_type -> mock.MockRequest
	1, // 3: mock.MockService.Mock:output_type -> mock.MockResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_proto_mock_proto_init() }
func file_api_proto_mock_proto_init() {
	if File_api_proto_mock_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_mock_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MockRequest); i {
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
		file_api_proto_mock_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MockResponse); i {
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
			RawDescriptor: file_api_proto_mock_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_mock_proto_goTypes,
		DependencyIndexes: file_api_proto_mock_proto_depIdxs,
		MessageInfos:      file_api_proto_mock_proto_msgTypes,
	}.Build()
	File_api_proto_mock_proto = out.File
	file_api_proto_mock_proto_rawDesc = nil
	file_api_proto_mock_proto_goTypes = nil
	file_api_proto_mock_proto_depIdxs = nil
}