// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calcurator/calcpb/calc.proto

package calcpb

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
	FirstNumber          int32    `protobuf:"varint,1,opt,name=first_number,json=firstNumber,proto3" json:"first_number,omitempty"`
	SecondNumber         int32    `protobuf:"varint,2,opt,name=second_number,json=secondNumber,proto3" json:"second_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumRequest) Reset()         { *m = SumRequest{} }
func (m *SumRequest) String() string { return proto.CompactTextString(m) }
func (*SumRequest) ProtoMessage()    {}
func (*SumRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_879c0539adf68aed, []int{0}
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

func (m *SumRequest) GetFirstNumber() int32 {
	if m != nil {
		return m.FirstNumber
	}
	return 0
}

func (m *SumRequest) GetSecondNumber() int32 {
	if m != nil {
		return m.SecondNumber
	}
	return 0
}

type SumResponse struct {
	SumResult            int32    `protobuf:"varint,1,opt,name=sum_result,json=sumResult,proto3" json:"sum_result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumResponse) Reset()         { *m = SumResponse{} }
func (m *SumResponse) String() string { return proto.CompactTextString(m) }
func (*SumResponse) ProtoMessage()    {}
func (*SumResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_879c0539adf68aed, []int{1}
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

func (m *SumResponse) GetSumResult() int32 {
	if m != nil {
		return m.SumResult
	}
	return 0
}

func init() {
	proto.RegisterType((*SumRequest)(nil), "calc.SumRequest")
	proto.RegisterType((*SumResponse)(nil), "calc.SumResponse")
}

func init() { proto.RegisterFile("calcurator/calcpb/calc.proto", fileDescriptor_879c0539adf68aed) }

var fileDescriptor_879c0539adf68aed = []byte{
	// 201 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0x4f, 0x4b, 0xc4, 0x30,
	0x10, 0xc5, 0xad, 0x7f, 0x8a, 0x4e, 0x2b, 0xd8, 0x9c, 0x44, 0x14, 0xb4, 0x5e, 0x3c, 0x94, 0x0a,
	0xfa, 0x09, 0xd4, 0xbb, 0x87, 0xd6, 0x93, 0x97, 0xd2, 0xc6, 0x11, 0x0a, 0x49, 0x53, 0x33, 0x19,
	0x3f, 0xff, 0x92, 0xc9, 0x2e, 0xbb, 0xa7, 0x19, 0x1e, 0x3f, 0x1e, 0xbf, 0x07, 0xb7, 0x7a, 0x34,
	0x9a, 0xfd, 0x18, 0x9c, 0x7f, 0x8e, 0xef, 0x3a, 0xc9, 0x69, 0x57, 0xef, 0x82, 0x53, 0xa7, 0xf1,
	0xaf, 0xbf, 0x00, 0x7a, 0xb6, 0x1d, 0xfe, 0x31, 0x52, 0x50, 0x0f, 0x50, 0xfe, 0xce, 0x9e, 0xc2,
	0xb0, 0xb0, 0x9d, 0xd0, 0x5f, 0x67, 0xf7, 0xd9, 0xd3, 0x59, 0x57, 0x48, 0xf6, 0x29, 0x91, 0x7a,
	0x84, 0x4b, 0x42, 0xed, 0x96, 0x9f, 0x1d, 0x73, 0x2c, 0x4c, 0x99, 0xc2, 0x04, 0xd5, 0x0d, 0x14,
	0xd2, 0x4a, 0xab, 0x5b, 0x08, 0xd5, 0x1d, 0x00, 0xb1, 0x1d, 0x3c, 0x12, 0x9b, 0xb0, 0x2d, 0xbd,
	0x20, 0x01, 0xd8, 0x84, 0x97, 0x37, 0xa8, 0x3e, 0xa2, 0xa9, 0x89, 0xa6, 0x3d, 0xfa, 0xff, 0x59,
	0xa3, 0x6a, 0xe0, 0xa4, 0x67, 0xab, 0xae, 0x5a, 0x51, 0xde, 0x3b, 0xde, 0x54, 0x07, 0x49, 0xea,
	0xaf, 0x8f, 0xde, 0xcf, 0xbf, 0xf3, 0xb4, 0x70, 0xca, 0x65, 0xdd, 0xeb, 0x26, 0x00, 0x00, 0xff,
	0xff, 0x96, 0xb4, 0xe6, 0x45, 0xfd, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CalculatorServiceClient is the client API for CalculatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalculatorServiceClient interface {
	Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error)
}

type calculatorServiceClient struct {
	cc *grpc.ClientConn
}

func NewCalculatorServiceClient(cc *grpc.ClientConn) CalculatorServiceClient {
	return &calculatorServiceClient{cc}
}

func (c *calculatorServiceClient) Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error) {
	out := new(SumResponse)
	err := c.cc.Invoke(ctx, "/calc.CalculatorService/Sum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalculatorServiceServer is the server API for CalculatorService service.
type CalculatorServiceServer interface {
	Sum(context.Context, *SumRequest) (*SumResponse, error)
}

// UnimplementedCalculatorServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCalculatorServiceServer struct {
}

func (*UnimplementedCalculatorServiceServer) Sum(ctx context.Context, req *SumRequest) (*SumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sum not implemented")
}

func RegisterCalculatorServiceServer(s *grpc.Server, srv CalculatorServiceServer) {
	s.RegisterService(&_CalculatorService_serviceDesc, srv)
}

func _CalculatorService_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calc.CalculatorService/Sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Sum(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CalculatorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calc.CalculatorService",
	HandlerType: (*CalculatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _CalculatorService_Sum_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "calcurator/calcpb/calc.proto",
}