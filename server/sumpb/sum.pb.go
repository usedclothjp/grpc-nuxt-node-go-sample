// Code generated by protoc-gen-go. DO NOT EDIT.
// source: server/proto/sum.proto

package sumpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type SumRequest struct {
	Num1                 int32    `protobuf:"varint,1,opt,name=num1,proto3" json:"num1,omitempty"`
	Num2                 int32    `protobuf:"varint,2,opt,name=num2,proto3" json:"num2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumRequest) Reset()         { *m = SumRequest{} }
func (m *SumRequest) String() string { return proto.CompactTextString(m) }
func (*SumRequest) ProtoMessage()    {}
func (*SumRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9d499b5cc5f8f64, []int{0}
}

func (m *SumRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumRequest.Unmarshal(m, b)
}
func (m *SumRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumRequest.Marshal(b, m, deterministic)
}
func (m *SumRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumRequest.Merge(m, src)
}
func (m *SumRequest) XXX_Size() int {
	return xxx_messageInfo_SumRequest.Size(m)
}
func (m *SumRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SumRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SumRequest proto.InternalMessageInfo

func (m *SumRequest) GetNum1() int32 {
	if m != nil {
		return m.Num1
	}
	return 0
}

func (m *SumRequest) GetNum2() int32 {
	if m != nil {
		return m.Num2
	}
	return 0
}

type SumResponse struct {
	Num                  int32    `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumResponse) Reset()         { *m = SumResponse{} }
func (m *SumResponse) String() string { return proto.CompactTextString(m) }
func (*SumResponse) ProtoMessage()    {}
func (*SumResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9d499b5cc5f8f64, []int{1}
}

func (m *SumResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumResponse.Unmarshal(m, b)
}
func (m *SumResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumResponse.Marshal(b, m, deterministic)
}
func (m *SumResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumResponse.Merge(m, src)
}
func (m *SumResponse) XXX_Size() int {
	return xxx_messageInfo_SumResponse.Size(m)
}
func (m *SumResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SumResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SumResponse proto.InternalMessageInfo

func (m *SumResponse) GetNum() int32 {
	if m != nil {
		return m.Num
	}
	return 0
}

func init() {
	proto.RegisterType((*SumRequest)(nil), "sumpb.SumRequest")
	proto.RegisterType((*SumResponse)(nil), "sumpb.SumResponse")
}

func init() { proto.RegisterFile("server/proto/sum.proto", fileDescriptor_d9d499b5cc5f8f64) }

var fileDescriptor_d9d499b5cc5f8f64 = []byte{
	// 151 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2b, 0x4e, 0x2d, 0x2a,
	0x4b, 0x2d, 0xd2, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x2f, 0x2e, 0xcd, 0xd5, 0x03, 0xb3, 0x84,
	0x58, 0x8b, 0x4b, 0x73, 0x0b, 0x92, 0x94, 0x4c, 0xb8, 0xb8, 0x82, 0x4b, 0x73, 0x83, 0x52, 0x0b,
	0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x84, 0xb8, 0x58, 0xf2, 0x4a, 0x73, 0x0d, 0x25, 0x18, 0x15, 0x18,
	0x35, 0x58, 0x83, 0xc0, 0x6c, 0xa8, 0x98, 0x91, 0x04, 0x13, 0x5c, 0xcc, 0x48, 0x49, 0x9e, 0x8b,
	0x1b, 0xac, 0xab, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0x48, 0x80, 0x8b, 0x39, 0xaf, 0x34, 0x17,
	0xaa, 0x0b, 0xc4, 0x34, 0xb2, 0x02, 0x1b, 0x1b, 0x9c, 0x5a, 0x54, 0x96, 0x99, 0x9c, 0x2a, 0xa4,
	0xc3, 0xc5, 0x1c, 0x5c, 0x9a, 0x2b, 0x24, 0xa8, 0x07, 0xb6, 0x53, 0x0f, 0x61, 0xa1, 0x94, 0x10,
	0xb2, 0x10, 0xc4, 0xb4, 0x24, 0x36, 0xb0, 0x03, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb7,
	0xbf, 0xdd, 0xfb, 0xba, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SumServiceClient is the client API for SumService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SumServiceClient interface {
	Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error)
}

type sumServiceClient struct {
	cc *grpc.ClientConn
}

func NewSumServiceClient(cc *grpc.ClientConn) SumServiceClient {
	return &sumServiceClient{cc}
}

func (c *sumServiceClient) Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error) {
	out := new(SumResponse)
	err := c.cc.Invoke(ctx, "/sumpb.SumService/Sum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SumServiceServer is the server API for SumService service.
type SumServiceServer interface {
	Sum(context.Context, *SumRequest) (*SumResponse, error)
}

// UnimplementedSumServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSumServiceServer struct {
}

func (*UnimplementedSumServiceServer) Sum(ctx context.Context, req *SumRequest) (*SumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sum not implemented")
}

func RegisterSumServiceServer(s *grpc.Server, srv SumServiceServer) {
	s.RegisterService(&_SumService_serviceDesc, srv)
}

func _SumService_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SumServiceServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sumpb.SumService/Sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SumServiceServer).Sum(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SumService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sumpb.SumService",
	HandlerType: (*SumServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _SumService_Sum_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server/proto/sum.proto",
}
