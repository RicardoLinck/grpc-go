// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calculator/calculatorpb/calculator.proto

package calculatorpb

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
	NumA                 int32    `protobuf:"varint,1,opt,name=numA,proto3" json:"numA,omitempty"`
	NumB                 int32    `protobuf:"varint,2,opt,name=numB,proto3" json:"numB,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumRequest) Reset()         { *m = SumRequest{} }
func (m *SumRequest) String() string { return proto.CompactTextString(m) }
func (*SumRequest) ProtoMessage()    {}
func (*SumRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{0}
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

func (m *SumRequest) GetNumA() int32 {
	if m != nil {
		return m.NumA
	}
	return 0
}

func (m *SumRequest) GetNumB() int32 {
	if m != nil {
		return m.NumB
	}
	return 0
}

type SumResponse struct {
	Result               int32    `protobuf:"varint,1,opt,name=Result,proto3" json:"Result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumResponse) Reset()         { *m = SumResponse{} }
func (m *SumResponse) String() string { return proto.CompactTextString(m) }
func (*SumResponse) ProtoMessage()    {}
func (*SumResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{1}
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

func (m *SumResponse) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

type PrimeNumberDecompositionRequest struct {
	Input                int32    `protobuf:"varint,1,opt,name=input,proto3" json:"input,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrimeNumberDecompositionRequest) Reset()         { *m = PrimeNumberDecompositionRequest{} }
func (m *PrimeNumberDecompositionRequest) String() string { return proto.CompactTextString(m) }
func (*PrimeNumberDecompositionRequest) ProtoMessage()    {}
func (*PrimeNumberDecompositionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{2}
}

func (m *PrimeNumberDecompositionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrimeNumberDecompositionRequest.Unmarshal(m, b)
}
func (m *PrimeNumberDecompositionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrimeNumberDecompositionRequest.Marshal(b, m, deterministic)
}
func (m *PrimeNumberDecompositionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrimeNumberDecompositionRequest.Merge(m, src)
}
func (m *PrimeNumberDecompositionRequest) XXX_Size() int {
	return xxx_messageInfo_PrimeNumberDecompositionRequest.Size(m)
}
func (m *PrimeNumberDecompositionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PrimeNumberDecompositionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PrimeNumberDecompositionRequest proto.InternalMessageInfo

func (m *PrimeNumberDecompositionRequest) GetInput() int32 {
	if m != nil {
		return m.Input
	}
	return 0
}

type PrimeNumberDecompositionResponse struct {
	Result               int32    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrimeNumberDecompositionResponse) Reset()         { *m = PrimeNumberDecompositionResponse{} }
func (m *PrimeNumberDecompositionResponse) String() string { return proto.CompactTextString(m) }
func (*PrimeNumberDecompositionResponse) ProtoMessage()    {}
func (*PrimeNumberDecompositionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{3}
}

func (m *PrimeNumberDecompositionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrimeNumberDecompositionResponse.Unmarshal(m, b)
}
func (m *PrimeNumberDecompositionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrimeNumberDecompositionResponse.Marshal(b, m, deterministic)
}
func (m *PrimeNumberDecompositionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrimeNumberDecompositionResponse.Merge(m, src)
}
func (m *PrimeNumberDecompositionResponse) XXX_Size() int {
	return xxx_messageInfo_PrimeNumberDecompositionResponse.Size(m)
}
func (m *PrimeNumberDecompositionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PrimeNumberDecompositionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PrimeNumberDecompositionResponse proto.InternalMessageInfo

func (m *PrimeNumberDecompositionResponse) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

type ComputeAverageRequest struct {
	Input                float64  `protobuf:"fixed64,1,opt,name=input,proto3" json:"input,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ComputeAverageRequest) Reset()         { *m = ComputeAverageRequest{} }
func (m *ComputeAverageRequest) String() string { return proto.CompactTextString(m) }
func (*ComputeAverageRequest) ProtoMessage()    {}
func (*ComputeAverageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{4}
}

func (m *ComputeAverageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComputeAverageRequest.Unmarshal(m, b)
}
func (m *ComputeAverageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComputeAverageRequest.Marshal(b, m, deterministic)
}
func (m *ComputeAverageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComputeAverageRequest.Merge(m, src)
}
func (m *ComputeAverageRequest) XXX_Size() int {
	return xxx_messageInfo_ComputeAverageRequest.Size(m)
}
func (m *ComputeAverageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ComputeAverageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ComputeAverageRequest proto.InternalMessageInfo

func (m *ComputeAverageRequest) GetInput() float64 {
	if m != nil {
		return m.Input
	}
	return 0
}

type ComputeAverageResponse struct {
	Result               float64  `protobuf:"fixed64,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ComputeAverageResponse) Reset()         { *m = ComputeAverageResponse{} }
func (m *ComputeAverageResponse) String() string { return proto.CompactTextString(m) }
func (*ComputeAverageResponse) ProtoMessage()    {}
func (*ComputeAverageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{5}
}

func (m *ComputeAverageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComputeAverageResponse.Unmarshal(m, b)
}
func (m *ComputeAverageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComputeAverageResponse.Marshal(b, m, deterministic)
}
func (m *ComputeAverageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComputeAverageResponse.Merge(m, src)
}
func (m *ComputeAverageResponse) XXX_Size() int {
	return xxx_messageInfo_ComputeAverageResponse.Size(m)
}
func (m *ComputeAverageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ComputeAverageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ComputeAverageResponse proto.InternalMessageInfo

func (m *ComputeAverageResponse) GetResult() float64 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*SumRequest)(nil), "calculator.SumRequest")
	proto.RegisterType((*SumResponse)(nil), "calculator.SumResponse")
	proto.RegisterType((*PrimeNumberDecompositionRequest)(nil), "calculator.PrimeNumberDecompositionRequest")
	proto.RegisterType((*PrimeNumberDecompositionResponse)(nil), "calculator.PrimeNumberDecompositionResponse")
	proto.RegisterType((*ComputeAverageRequest)(nil), "calculator.ComputeAverageRequest")
	proto.RegisterType((*ComputeAverageResponse)(nil), "calculator.ComputeAverageResponse")
}

func init() {
	proto.RegisterFile("calculator/calculatorpb/calculator.proto", fileDescriptor_7f42938f8c8365cf)
}

var fileDescriptor_7f42938f8c8365cf = []byte{
	// 291 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0xc6, 0xcd, 0x74, 0x3b, 0x3c, 0x65, 0x60, 0xd0, 0x5a, 0x7a, 0x71, 0x06, 0x84, 0x82, 0x3a,
	0x87, 0x0a, 0x8a, 0xb7, 0x75, 0x9e, 0x45, 0xda, 0x9b, 0x9e, 0xda, 0xf2, 0x90, 0x42, 0xd3, 0xc4,
	0x34, 0x19, 0xf8, 0x07, 0xf9, 0x7f, 0x0a, 0x6d, 0xb3, 0x76, 0xda, 0xe9, 0x6e, 0xf9, 0x3e, 0xde,
	0x97, 0x1f, 0xef, 0xe3, 0x81, 0x9f, 0xc6, 0x79, 0x6a, 0xf2, 0x58, 0x0b, 0x75, 0xdd, 0x3e, 0x65,
	0xd2, 0x11, 0x53, 0xa9, 0x84, 0x16, 0x14, 0x5a, 0x87, 0xdd, 0x01, 0x44, 0x86, 0x87, 0xf8, 0x61,
	0xb0, 0xd4, 0x94, 0xc2, 0x5e, 0x61, 0xf8, 0xdc, 0x25, 0x13, 0xe2, 0x0f, 0xc3, 0xea, 0xdd, 0x78,
	0x81, 0x3b, 0x58, 0x79, 0x01, 0x3b, 0x87, 0xfd, 0x2a, 0x55, 0x4a, 0x51, 0x94, 0x48, 0x1d, 0x18,
	0x85, 0x58, 0x9a, 0x5c, 0x37, 0xc1, 0x46, 0xb1, 0x7b, 0x38, 0x7d, 0x51, 0x19, 0xc7, 0x67, 0xc3,
	0x13, 0x54, 0x4f, 0x98, 0x0a, 0x2e, 0x45, 0x99, 0xe9, 0x4c, 0x14, 0x96, 0x78, 0x04, 0xc3, 0xac,
	0x90, 0xc6, 0x26, 0x6b, 0xc1, 0x1e, 0x61, 0xb2, 0x39, 0xd8, 0x42, 0xd5, 0x1a, 0xb4, 0x56, 0xec,
	0x0a, 0x8e, 0x17, 0x82, 0x4b, 0xa3, 0x71, 0xbe, 0x44, 0x15, 0xbf, 0x63, 0x2f, 0x8a, 0x58, 0xd4,
	0x0c, 0x9c, 0x9f, 0xe3, 0xbd, 0x00, 0x62, 0x01, 0x37, 0x5f, 0x03, 0x38, 0x5c, 0xac, 0x1a, 0x8c,
	0x50, 0x2d, 0xb3, 0x14, 0xe9, 0x03, 0xec, 0x46, 0x86, 0x53, 0x67, 0xda, 0xa9, 0xbb, 0x6d, 0xd6,
	0x3b, 0xf9, 0xe5, 0xd7, 0x14, 0xb6, 0x43, 0x3f, 0xc1, 0xdd, 0xb4, 0x2c, 0xbd, 0xe8, 0xc6, 0xfe,
	0xe9, 0xd2, 0xbb, 0xdc, 0x6e, 0xd8, 0x82, 0x67, 0x84, 0xbe, 0xc1, 0x78, 0x7d, 0x79, 0x7a, 0xd6,
	0xfd, 0xa3, 0xb7, 0x47, 0x8f, 0xfd, 0x35, 0x62, 0x3f, 0xf7, 0x49, 0x30, 0x7e, 0x3d, 0xe8, 0xde,
	0x61, 0x32, 0xaa, 0xae, 0xef, 0xf6, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x0d, 0x0b, 0xf1, 0xe9, 0xa9,
	0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CalculatorServiceClient is the client API for CalculatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalculatorServiceClient interface {
	Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error)
	PrimeNumberDecomposition(ctx context.Context, in *PrimeNumberDecompositionRequest, opts ...grpc.CallOption) (CalculatorService_PrimeNumberDecompositionClient, error)
	ComputeAverage(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_ComputeAverageClient, error)
}

type calculatorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculatorServiceClient(cc grpc.ClientConnInterface) CalculatorServiceClient {
	return &calculatorServiceClient{cc}
}

func (c *calculatorServiceClient) Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error) {
	out := new(SumResponse)
	err := c.cc.Invoke(ctx, "/calculator.CalculatorService/Sum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) PrimeNumberDecomposition(ctx context.Context, in *PrimeNumberDecompositionRequest, opts ...grpc.CallOption) (CalculatorService_PrimeNumberDecompositionClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculatorService_serviceDesc.Streams[0], "/calculator.CalculatorService/PrimeNumberDecomposition", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServicePrimeNumberDecompositionClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CalculatorService_PrimeNumberDecompositionClient interface {
	Recv() (*PrimeNumberDecompositionResponse, error)
	grpc.ClientStream
}

type calculatorServicePrimeNumberDecompositionClient struct {
	grpc.ClientStream
}

func (x *calculatorServicePrimeNumberDecompositionClient) Recv() (*PrimeNumberDecompositionResponse, error) {
	m := new(PrimeNumberDecompositionResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorServiceClient) ComputeAverage(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_ComputeAverageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculatorService_serviceDesc.Streams[1], "/calculator.CalculatorService/ComputeAverage", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceComputeAverageClient{stream}
	return x, nil
}

type CalculatorService_ComputeAverageClient interface {
	Send(*ComputeAverageRequest) error
	CloseAndRecv() (*ComputeAverageResponse, error)
	grpc.ClientStream
}

type calculatorServiceComputeAverageClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceComputeAverageClient) Send(m *ComputeAverageRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorServiceComputeAverageClient) CloseAndRecv() (*ComputeAverageResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ComputeAverageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalculatorServiceServer is the server API for CalculatorService service.
type CalculatorServiceServer interface {
	Sum(context.Context, *SumRequest) (*SumResponse, error)
	PrimeNumberDecomposition(*PrimeNumberDecompositionRequest, CalculatorService_PrimeNumberDecompositionServer) error
	ComputeAverage(CalculatorService_ComputeAverageServer) error
}

// UnimplementedCalculatorServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCalculatorServiceServer struct {
}

func (*UnimplementedCalculatorServiceServer) Sum(ctx context.Context, req *SumRequest) (*SumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sum not implemented")
}
func (*UnimplementedCalculatorServiceServer) PrimeNumberDecomposition(req *PrimeNumberDecompositionRequest, srv CalculatorService_PrimeNumberDecompositionServer) error {
	return status.Errorf(codes.Unimplemented, "method PrimeNumberDecomposition not implemented")
}
func (*UnimplementedCalculatorServiceServer) ComputeAverage(srv CalculatorService_ComputeAverageServer) error {
	return status.Errorf(codes.Unimplemented, "method ComputeAverage not implemented")
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
		FullMethod: "/calculator.CalculatorService/Sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Sum(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_PrimeNumberDecomposition_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PrimeNumberDecompositionRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalculatorServiceServer).PrimeNumberDecomposition(m, &calculatorServicePrimeNumberDecompositionServer{stream})
}

type CalculatorService_PrimeNumberDecompositionServer interface {
	Send(*PrimeNumberDecompositionResponse) error
	grpc.ServerStream
}

type calculatorServicePrimeNumberDecompositionServer struct {
	grpc.ServerStream
}

func (x *calculatorServicePrimeNumberDecompositionServer) Send(m *PrimeNumberDecompositionResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CalculatorService_ComputeAverage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServiceServer).ComputeAverage(&calculatorServiceComputeAverageServer{stream})
}

type CalculatorService_ComputeAverageServer interface {
	SendAndClose(*ComputeAverageResponse) error
	Recv() (*ComputeAverageRequest, error)
	grpc.ServerStream
}

type calculatorServiceComputeAverageServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceComputeAverageServer) SendAndClose(m *ComputeAverageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorServiceComputeAverageServer) Recv() (*ComputeAverageRequest, error) {
	m := new(ComputeAverageRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _CalculatorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.CalculatorService",
	HandlerType: (*CalculatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _CalculatorService_Sum_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PrimeNumberDecomposition",
			Handler:       _CalculatorService_PrimeNumberDecomposition_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ComputeAverage",
			Handler:       _CalculatorService_ComputeAverage_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "calculator/calculatorpb/calculator.proto",
}
