// Code generated by protoc-gen-go. DO NOT EDIT.
// source: services/greetings.proto

package services

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	models "github.com/scottcrawford03/grpc-service.felix/models"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GreetReq struct {
	Request              *models.Greeting `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GreetReq) Reset()         { *m = GreetReq{} }
func (m *GreetReq) String() string { return proto.CompactTextString(m) }
func (*GreetReq) ProtoMessage()    {}
func (*GreetReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_5835c466146ca2c2, []int{0}
}

func (m *GreetReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetReq.Unmarshal(m, b)
}
func (m *GreetReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetReq.Marshal(b, m, deterministic)
}
func (m *GreetReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetReq.Merge(m, src)
}
func (m *GreetReq) XXX_Size() int {
	return xxx_messageInfo_GreetReq.Size(m)
}
func (m *GreetReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetReq.DiscardUnknown(m)
}

var xxx_messageInfo_GreetReq proto.InternalMessageInfo

func (m *GreetReq) GetRequest() *models.Greeting {
	if m != nil {
		return m.Request
	}
	return nil
}

type GreetResp struct {
	Response             *models.Greeting `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GreetResp) Reset()         { *m = GreetResp{} }
func (m *GreetResp) String() string { return proto.CompactTextString(m) }
func (*GreetResp) ProtoMessage()    {}
func (*GreetResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_5835c466146ca2c2, []int{1}
}

func (m *GreetResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetResp.Unmarshal(m, b)
}
func (m *GreetResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetResp.Marshal(b, m, deterministic)
}
func (m *GreetResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetResp.Merge(m, src)
}
func (m *GreetResp) XXX_Size() int {
	return xxx_messageInfo_GreetResp.Size(m)
}
func (m *GreetResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetResp.DiscardUnknown(m)
}

var xxx_messageInfo_GreetResp proto.InternalMessageInfo

func (m *GreetResp) GetResponse() *models.Greeting {
	if m != nil {
		return m.Response
	}
	return nil
}

func init() {
	proto.RegisterType((*GreetReq)(nil), "services.GreetReq")
	proto.RegisterType((*GreetResp)(nil), "services.GreetResp")
}

func init() { proto.RegisterFile("services/greetings.proto", fileDescriptor_5835c466146ca2c2) }

var fileDescriptor_5835c466146ca2c2 = []byte{
	// 202 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x8f, 0x41, 0x4f, 0x84, 0x30,
	0x10, 0x46, 0xe5, 0xa0, 0x62, 0x4d, 0x8c, 0xa9, 0x17, 0xc2, 0xc9, 0x70, 0x32, 0x46, 0x5b, 0x03,
	0x89, 0x89, 0x77, 0x12, 0xee, 0xf8, 0x0b, 0xa4, 0x0c, 0xb5, 0x09, 0xd0, 0x32, 0x53, 0x74, 0x7f,
	0xfe, 0x66, 0x81, 0xee, 0x85, 0xec, 0xf5, 0xf5, 0x7b, 0xcd, 0x1b, 0x96, 0x10, 0xe0, 0x9f, 0x51,
	0x40, 0x52, 0x23, 0x80, 0x37, 0xa3, 0x26, 0xe1, 0xd0, 0x7a, 0xcb, 0xe3, 0xf0, 0x92, 0x96, 0xda,
	0xf8, 0xdf, 0xb9, 0x11, 0xca, 0x0e, 0x92, 0x94, 0xf5, 0x5e, 0xe1, 0xcf, 0x7f, 0x67, 0xb1, 0xfd,
	0x28, 0xa4, 0x46, 0xa7, 0xde, 0xb7, 0xa5, 0xe8, 0xa0, 0x37, 0x07, 0x39, 0xd8, 0x16, 0xfa, 0xdd,
	0x7f, 0xd9, 0x27, 0x8b, 0xab, 0x13, 0xaa, 0x61, 0xe2, 0xaf, 0xec, 0x16, 0x61, 0x9a, 0x81, 0x7c,
	0x12, 0x3d, 0x47, 0x2f, 0xf7, 0xf9, 0xa3, 0x58, 0x2d, 0x51, 0x6d, 0x56, 0x1d, 0x06, 0xd9, 0x17,
	0xbb, 0xdb, 0x3c, 0x72, 0xfc, 0x8d, 0xc5, 0x08, 0xe4, 0xec, 0x48, 0x70, 0xd1, 0x3c, 0x2f, 0xf2,
	0x92, 0x3d, 0x2c, 0x14, 0xf0, 0x7b, 0x2d, 0xe4, 0x39, 0xbb, 0x5e, 0x08, 0xe7, 0x22, 0x9c, 0x27,
	0x42, 0x55, 0xfa, 0xb4, 0x63, 0xe4, 0xb2, 0xab, 0xe6, 0x66, 0xe9, 0x2f, 0x8e, 0x01, 0x00, 0x00,
	0xff, 0xff, 0xf0, 0x3c, 0x03, 0xd0, 0x2b, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GreeterServiceClient is the client API for GreeterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreeterServiceClient interface {
	Greet(ctx context.Context, in *GreetReq, opts ...grpc.CallOption) (*GreetResp, error)
}

type greeterServiceClient struct {
	cc *grpc.ClientConn
}

func NewGreeterServiceClient(cc *grpc.ClientConn) GreeterServiceClient {
	return &greeterServiceClient{cc}
}

func (c *greeterServiceClient) Greet(ctx context.Context, in *GreetReq, opts ...grpc.CallOption) (*GreetResp, error) {
	out := new(GreetResp)
	err := c.cc.Invoke(ctx, "/services.GreeterService/Greet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreeterServiceServer is the server API for GreeterService service.
type GreeterServiceServer interface {
	Greet(context.Context, *GreetReq) (*GreetResp, error)
}

// UnimplementedGreeterServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGreeterServiceServer struct {
}

func (*UnimplementedGreeterServiceServer) Greet(ctx context.Context, req *GreetReq) (*GreetResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Greet not implemented")
}

func RegisterGreeterServiceServer(s *grpc.Server, srv GreeterServiceServer) {
	s.RegisterService(&_GreeterService_serviceDesc, srv)
}

func _GreeterService_Greet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GreetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServiceServer).Greet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.GreeterService/Greet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServiceServer).Greet(ctx, req.(*GreetReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _GreeterService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "services.GreeterService",
	HandlerType: (*GreeterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Greet",
			Handler:    _GreeterService_Greet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/greetings.proto",
}
