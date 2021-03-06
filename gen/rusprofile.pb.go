// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.14.0
// source: pb/rusprofile.proto

package rusprofile

import (
	context "context"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type FirmByINNRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Inn string `protobuf:"bytes,1,opt,name=inn,proto3" json:"inn,omitempty"`
}

func (x *FirmByINNRequest) Reset() {
	*x = FirmByINNRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_rusprofile_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FirmByINNRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FirmByINNRequest) ProtoMessage() {}

func (x *FirmByINNRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_rusprofile_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FirmByINNRequest.ProtoReflect.Descriptor instead.
func (*FirmByINNRequest) Descriptor() ([]byte, []int) {
	return file_pb_rusprofile_proto_rawDescGZIP(), []int{0}
}

func (x *FirmByINNRequest) GetInn() string {
	if x != nil {
		return x.Inn
	}
	return ""
}

type FirmInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Kpp  string `protobuf:"bytes,2,opt,name=kpp,proto3" json:"kpp,omitempty"`
	Inn  string `protobuf:"bytes,3,opt,name=inn,proto3" json:"inn,omitempty"`
	Boss string `protobuf:"bytes,4,opt,name=boss,proto3" json:"boss,omitempty"`
}

func (x *FirmInfoResponse) Reset() {
	*x = FirmInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_rusprofile_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FirmInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FirmInfoResponse) ProtoMessage() {}

func (x *FirmInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_rusprofile_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FirmInfoResponse.ProtoReflect.Descriptor instead.
func (*FirmInfoResponse) Descriptor() ([]byte, []int) {
	return file_pb_rusprofile_proto_rawDescGZIP(), []int{1}
}

func (x *FirmInfoResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FirmInfoResponse) GetKpp() string {
	if x != nil {
		return x.Kpp
	}
	return ""
}

func (x *FirmInfoResponse) GetInn() string {
	if x != nil {
		return x.Inn
	}
	return ""
}

func (x *FirmInfoResponse) GetBoss() string {
	if x != nil {
		return x.Boss
	}
	return ""
}

var File_pb_rusprofile_proto protoreflect.FileDescriptor

var file_pb_rusprofile_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x62, 0x2f, 0x72, 0x75, 0x73, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x72, 0x75, 0x73, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x24, 0x0a, 0x10, 0x46, 0x69, 0x72, 0x6d, 0x42, 0x79, 0x49, 0x4e, 0x4e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x6e, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x69, 0x6e, 0x6e, 0x22, 0x5e, 0x0a, 0x10, 0x46, 0x69, 0x72, 0x6d, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x70, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x70, 0x70, 0x12,
	0x10, 0x0a, 0x03, 0x69, 0x6e, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x69, 0x6e,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x62, 0x6f, 0x73, 0x73, 0x32, 0x65, 0x0a, 0x07, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72,
	0x12, 0x5a, 0x0a, 0x0b, 0x46, 0x69, 0x72, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x47, 0x65, 0x74, 0x12,
	0x1c, 0x2e, 0x72, 0x75, 0x73, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x46, 0x69, 0x72,
	0x6d, 0x42, 0x79, 0x49, 0x4e, 0x4e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e,
	0x72, 0x75, 0x73, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x46, 0x69, 0x72, 0x6d, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0f, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x09, 0x22, 0x04, 0x2f, 0x69, 0x6e, 0x6e, 0x3a, 0x01, 0x2a, 0x42, 0x12, 0x5a, 0x10,
	0x2e, 0x2f, 0x67, 0x65, 0x6e, 0x3b, 0x72, 0x75, 0x73, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_rusprofile_proto_rawDescOnce sync.Once
	file_pb_rusprofile_proto_rawDescData = file_pb_rusprofile_proto_rawDesc
)

func file_pb_rusprofile_proto_rawDescGZIP() []byte {
	file_pb_rusprofile_proto_rawDescOnce.Do(func() {
		file_pb_rusprofile_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_rusprofile_proto_rawDescData)
	})
	return file_pb_rusprofile_proto_rawDescData
}

var file_pb_rusprofile_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pb_rusprofile_proto_goTypes = []interface{}{
	(*FirmByINNRequest)(nil), // 0: rusprofile.FirmByINNRequest
	(*FirmInfoResponse)(nil), // 1: rusprofile.FirmInfoResponse
}
var file_pb_rusprofile_proto_depIdxs = []int32{
	0, // 0: rusprofile.Greeter.FirmInfoGet:input_type -> rusprofile.FirmByINNRequest
	1, // 1: rusprofile.Greeter.FirmInfoGet:output_type -> rusprofile.FirmInfoResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_rusprofile_proto_init() }
func file_pb_rusprofile_proto_init() {
	if File_pb_rusprofile_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_rusprofile_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FirmByINNRequest); i {
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
		file_pb_rusprofile_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FirmInfoResponse); i {
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
			RawDescriptor: file_pb_rusprofile_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_rusprofile_proto_goTypes,
		DependencyIndexes: file_pb_rusprofile_proto_depIdxs,
		MessageInfos:      file_pb_rusprofile_proto_msgTypes,
	}.Build()
	File_pb_rusprofile_proto = out.File
	file_pb_rusprofile_proto_rawDesc = nil
	file_pb_rusprofile_proto_goTypes = nil
	file_pb_rusprofile_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreeterClient interface {
	FirmInfoGet(ctx context.Context, in *FirmByINNRequest, opts ...grpc.CallOption) (*FirmInfoResponse, error)
}

type greeterClient struct {
	cc grpc.ClientConnInterface
}

func NewGreeterClient(cc grpc.ClientConnInterface) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) FirmInfoGet(ctx context.Context, in *FirmByINNRequest, opts ...grpc.CallOption) (*FirmInfoResponse, error) {
	out := new(FirmInfoResponse)
	err := c.cc.Invoke(ctx, "/rusprofile.Greeter/FirmInfoGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreeterServer is the server API for Greeter service.
type GreeterServer interface {
	FirmInfoGet(context.Context, *FirmByINNRequest) (*FirmInfoResponse, error)
}

// UnimplementedGreeterServer can be embedded to have forward compatible implementations.
type UnimplementedGreeterServer struct {
}

func (*UnimplementedGreeterServer) FirmInfoGet(context.Context, *FirmByINNRequest) (*FirmInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FirmInfoGet not implemented")
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_FirmInfoGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FirmByINNRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).FirmInfoGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rusprofile.Greeter/FirmInfoGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).FirmInfoGet(ctx, req.(*FirmByINNRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rusprofile.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FirmInfoGet",
			Handler:    _Greeter_FirmInfoGet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/rusprofile.proto",
}
