// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: comdex/bandoracle/v1beta1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// QueryParamsRequest is request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_51b7d3cb3f498405, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params holds all the parameters of this module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_51b7d3cb3f498405, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

type QueryFetchPriceRequest struct {
	RequestId int64 `protobuf:"varint,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
}

func (m *QueryFetchPriceRequest) Reset()         { *m = QueryFetchPriceRequest{} }
func (m *QueryFetchPriceRequest) String() string { return proto.CompactTextString(m) }
func (*QueryFetchPriceRequest) ProtoMessage()    {}
func (*QueryFetchPriceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_51b7d3cb3f498405, []int{2}
}
func (m *QueryFetchPriceRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryFetchPriceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryFetchPriceRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryFetchPriceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryFetchPriceRequest.Merge(m, src)
}
func (m *QueryFetchPriceRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryFetchPriceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryFetchPriceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryFetchPriceRequest proto.InternalMessageInfo

func (m *QueryFetchPriceRequest) GetRequestId() int64 {
	if m != nil {
		return m.RequestId
	}
	return 0
}

type QueryFetchPriceResponse struct {
	Result *FetchPriceResult `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (m *QueryFetchPriceResponse) Reset()         { *m = QueryFetchPriceResponse{} }
func (m *QueryFetchPriceResponse) String() string { return proto.CompactTextString(m) }
func (*QueryFetchPriceResponse) ProtoMessage()    {}
func (*QueryFetchPriceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_51b7d3cb3f498405, []int{3}
}
func (m *QueryFetchPriceResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryFetchPriceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryFetchPriceResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryFetchPriceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryFetchPriceResponse.Merge(m, src)
}
func (m *QueryFetchPriceResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryFetchPriceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryFetchPriceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryFetchPriceResponse proto.InternalMessageInfo

func (m *QueryFetchPriceResponse) GetResult() *FetchPriceResult {
	if m != nil {
		return m.Result
	}
	return nil
}

type QueryLastFetchPriceIdRequest struct {
}

func (m *QueryLastFetchPriceIdRequest) Reset()         { *m = QueryLastFetchPriceIdRequest{} }
func (m *QueryLastFetchPriceIdRequest) String() string { return proto.CompactTextString(m) }
func (*QueryLastFetchPriceIdRequest) ProtoMessage()    {}
func (*QueryLastFetchPriceIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_51b7d3cb3f498405, []int{4}
}
func (m *QueryLastFetchPriceIdRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryLastFetchPriceIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryLastFetchPriceIdRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryLastFetchPriceIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryLastFetchPriceIdRequest.Merge(m, src)
}
func (m *QueryLastFetchPriceIdRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryLastFetchPriceIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryLastFetchPriceIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryLastFetchPriceIdRequest proto.InternalMessageInfo

type QueryLastFetchPriceIdResponse struct {
	RequestId int64 `protobuf:"varint,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
}

func (m *QueryLastFetchPriceIdResponse) Reset()         { *m = QueryLastFetchPriceIdResponse{} }
func (m *QueryLastFetchPriceIdResponse) String() string { return proto.CompactTextString(m) }
func (*QueryLastFetchPriceIdResponse) ProtoMessage()    {}
func (*QueryLastFetchPriceIdResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_51b7d3cb3f498405, []int{5}
}
func (m *QueryLastFetchPriceIdResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryLastFetchPriceIdResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryLastFetchPriceIdResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryLastFetchPriceIdResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryLastFetchPriceIdResponse.Merge(m, src)
}
func (m *QueryLastFetchPriceIdResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryLastFetchPriceIdResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryLastFetchPriceIdResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryLastFetchPriceIdResponse proto.InternalMessageInfo

func (m *QueryLastFetchPriceIdResponse) GetRequestId() int64 {
	if m != nil {
		return m.RequestId
	}
	return 0
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "comdex.bandoracle.v1beta1.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "comdex.bandoracle.v1beta1.QueryParamsResponse")
	proto.RegisterType((*QueryFetchPriceRequest)(nil), "comdex.bandoracle.v1beta1.QueryFetchPriceRequest")
	proto.RegisterType((*QueryFetchPriceResponse)(nil), "comdex.bandoracle.v1beta1.QueryFetchPriceResponse")
	proto.RegisterType((*QueryLastFetchPriceIdRequest)(nil), "comdex.bandoracle.v1beta1.QueryLastFetchPriceIdRequest")
	proto.RegisterType((*QueryLastFetchPriceIdResponse)(nil), "comdex.bandoracle.v1beta1.QueryLastFetchPriceIdResponse")
}

func init() {
	proto.RegisterFile("comdex/bandoracle/v1beta1/query.proto", fileDescriptor_51b7d3cb3f498405)
}

var fileDescriptor_51b7d3cb3f498405 = []byte{
	// 487 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x4f, 0x8b, 0xd3, 0x40,
	0x18, 0xc6, 0x13, 0x5d, 0x03, 0x8e, 0x17, 0x19, 0x17, 0xff, 0x84, 0xdd, 0xe8, 0x46, 0x14, 0x75,
	0xd9, 0x8c, 0xad, 0xe2, 0x8a, 0x88, 0x42, 0x05, 0x61, 0x41, 0x64, 0xcd, 0xc1, 0x83, 0x07, 0xc3,
	0x24, 0x99, 0x66, 0x07, 0xd2, 0x4c, 0x9a, 0x99, 0x88, 0x45, 0xbc, 0xf8, 0x09, 0x44, 0xaf, 0x7e,
	0x15, 0xef, 0x3d, 0x16, 0xbc, 0x78, 0x12, 0x69, 0xfd, 0x20, 0x92, 0x99, 0xa9, 0xb5, 0x2d, 0x49,
	0xdb, 0x5b, 0x98, 0x79, 0x9e, 0xf7, 0xf9, 0xbd, 0xef, 0x3b, 0x01, 0x37, 0x22, 0xd6, 0x8b, 0xc9,
	0x7b, 0x14, 0xe2, 0x2c, 0x66, 0x05, 0x8e, 0x52, 0x82, 0xde, 0xb5, 0x42, 0x22, 0x70, 0x0b, 0xf5,
	0x4b, 0x52, 0x0c, 0xbc, 0xbc, 0x60, 0x82, 0xc1, 0x2b, 0x4a, 0xe6, 0xcd, 0x64, 0x9e, 0x96, 0xd9,
	0xdb, 0x09, 0x4b, 0x98, 0x54, 0xa1, 0xea, 0x4b, 0x19, 0xec, 0x9d, 0x84, 0xb1, 0x24, 0x25, 0x08,
	0xe7, 0x14, 0xe1, 0x2c, 0x63, 0x02, 0x0b, 0xca, 0x32, 0xae, 0x6f, 0xef, 0x44, 0x8c, 0xf7, 0x18,
	0x47, 0x21, 0xe6, 0x44, 0xe5, 0xfc, 0x4b, 0xcd, 0x71, 0x42, 0x33, 0x29, 0xd6, 0xda, 0xfd, 0x7a,
	0xc2, 0x2e, 0x11, 0xd1, 0x49, 0x90, 0x17, 0x34, 0x22, 0x5a, 0x7c, 0xb3, 0x5e, 0x9c, 0xe3, 0x02,
	0xf7, 0x34, 0x80, 0xbb, 0x0d, 0xe0, 0xab, 0x2a, 0xf6, 0x58, 0x1e, 0xfa, 0xa4, 0x5f, 0x12, 0x2e,
	0xdc, 0xd7, 0xe0, 0xc2, 0xdc, 0x29, 0xcf, 0x59, 0xc6, 0x09, 0x7c, 0x0a, 0x2c, 0x65, 0xbe, 0x6c,
	0x5e, 0x33, 0x6f, 0x9d, 0x6b, 0xef, 0x79, 0xb5, 0xd3, 0xf0, 0x94, 0xb5, 0xb3, 0x35, 0xfc, 0x75,
	0xd5, 0xf0, 0xb5, 0xcd, 0x3d, 0x04, 0x17, 0x65, 0xdd, 0xe7, 0x15, 0xef, 0x71, 0x85, 0xab, 0x13,
	0xe1, 0x2e, 0x00, 0x85, 0xfa, 0x0c, 0x68, 0x2c, 0xcb, 0x9f, 0xf6, 0xcf, 0xea, 0x93, 0xa3, 0xd8,
	0x7d, 0x0b, 0x2e, 0x2d, 0x19, 0x35, 0xd4, 0x33, 0x60, 0x15, 0x84, 0x97, 0xa9, 0xd0, 0x50, 0xfb,
	0x0d, 0x50, 0x73, 0xf6, 0x32, 0x15, 0xbe, 0xb6, 0xba, 0x0e, 0xd8, 0x91, 0xf5, 0x5f, 0x60, 0x2e,
	0x66, 0xa2, 0xa3, 0x78, 0x3a, 0x90, 0x27, 0x60, 0xb7, 0xe6, 0x5e, 0x53, 0x34, 0xf3, 0xb7, 0xbf,
	0x6d, 0x81, 0x33, 0xb2, 0x00, 0xfc, 0x62, 0x02, 0x4b, 0xcd, 0x06, 0x1e, 0x34, 0x90, 0x2e, 0x2f,
	0xc5, 0xf6, 0xd6, 0x95, 0x2b, 0x24, 0xf7, 0xf6, 0xa7, 0x1f, 0x7f, 0xbe, 0x9e, 0xba, 0x0e, 0xf7,
	0xd0, 0xaa, 0xb7, 0x00, 0xbf, 0x9b, 0xe0, 0xfc, 0xe2, 0x6c, 0x60, 0x6b, 0x55, 0xde, 0xd2, 0x16,
	0xed, 0xf6, 0x26, 0x16, 0x8d, 0xd9, 0x91, 0x98, 0x8f, 0xe1, 0x23, 0xb4, 0xd6, 0xfb, 0x0e, 0xd4,
	0xc6, 0xd0, 0x87, 0xd9, 0xb8, 0x3f, 0x4a, 0xfe, 0xc5, 0xd5, 0xc0, 0xc3, 0x55, 0x30, 0x35, 0xcb,
	0xb6, 0x1f, 0x6e, 0x6e, 0xd4, 0xbd, 0x3c, 0x90, 0xbd, 0xdc, 0x85, 0x5e, 0x43, 0x2f, 0x29, 0xe6,
	0x22, 0xf8, 0xbf, 0x21, 0x1a, 0x77, 0x5e, 0x0e, 0xc7, 0x8e, 0x39, 0x1a, 0x3b, 0xe6, 0xef, 0xb1,
	0x63, 0x7e, 0x9e, 0x38, 0xc6, 0x68, 0xe2, 0x18, 0x3f, 0x27, 0x8e, 0xf1, 0xe6, 0x7e, 0x42, 0xc5,
	0x49, 0x19, 0x56, 0x44, 0xba, 0xe6, 0x01, 0xeb, 0x76, 0x69, 0x44, 0x71, 0x3a, 0xcd, 0x98, 0x4b,
	0x11, 0x83, 0x9c, 0xf0, 0xd0, 0x92, 0x3f, 0xf7, 0xbd, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd9,
	0x55, 0xe3, 0xc9, 0xd5, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// FetchPriceResult defines a rpc handler method for MsgFetchPriceData.
	FetchPriceResult(ctx context.Context, in *QueryFetchPriceRequest, opts ...grpc.CallOption) (*QueryFetchPriceResponse, error)
	// LastFetchPriceId query the last FetchPrice result id
	LastFetchPriceId(ctx context.Context, in *QueryLastFetchPriceIdRequest, opts ...grpc.CallOption) (*QueryLastFetchPriceIdResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/comdex.bandoracle.v1beta1.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) FetchPriceResult(ctx context.Context, in *QueryFetchPriceRequest, opts ...grpc.CallOption) (*QueryFetchPriceResponse, error) {
	out := new(QueryFetchPriceResponse)
	err := c.cc.Invoke(ctx, "/comdex.bandoracle.v1beta1.Query/FetchPriceResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) LastFetchPriceId(ctx context.Context, in *QueryLastFetchPriceIdRequest, opts ...grpc.CallOption) (*QueryLastFetchPriceIdResponse, error) {
	out := new(QueryLastFetchPriceIdResponse)
	err := c.cc.Invoke(ctx, "/comdex.bandoracle.v1beta1.Query/LastFetchPriceId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// FetchPriceResult defines a rpc handler method for MsgFetchPriceData.
	FetchPriceResult(context.Context, *QueryFetchPriceRequest) (*QueryFetchPriceResponse, error)
	// LastFetchPriceId query the last FetchPrice result id
	LastFetchPriceId(context.Context, *QueryLastFetchPriceIdRequest) (*QueryLastFetchPriceIdResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) FetchPriceResult(ctx context.Context, req *QueryFetchPriceRequest) (*QueryFetchPriceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchPriceResult not implemented")
}
func (*UnimplementedQueryServer) LastFetchPriceId(ctx context.Context, req *QueryLastFetchPriceIdRequest) (*QueryLastFetchPriceIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LastFetchPriceId not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comdex.bandoracle.v1beta1.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_FetchPriceResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryFetchPriceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).FetchPriceResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comdex.bandoracle.v1beta1.Query/FetchPriceResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).FetchPriceResult(ctx, req.(*QueryFetchPriceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_LastFetchPriceId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryLastFetchPriceIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).LastFetchPriceId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comdex.bandoracle.v1beta1.Query/LastFetchPriceId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).LastFetchPriceId(ctx, req.(*QueryLastFetchPriceIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "comdex.bandoracle.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "FetchPriceResult",
			Handler:    _Query_FetchPriceResult_Handler,
		},
		{
			MethodName: "LastFetchPriceId",
			Handler:    _Query_LastFetchPriceId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "comdex/bandoracle/v1beta1/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryFetchPriceRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryFetchPriceRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryFetchPriceRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.RequestId != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.RequestId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryFetchPriceResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryFetchPriceResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryFetchPriceResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Result != nil {
		{
			size, err := m.Result.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryLastFetchPriceIdRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryLastFetchPriceIdRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryLastFetchPriceIdRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryLastFetchPriceIdResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryLastFetchPriceIdResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryLastFetchPriceIdResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.RequestId != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.RequestId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryFetchPriceRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RequestId != 0 {
		n += 1 + sovQuery(uint64(m.RequestId))
	}
	return n
}

func (m *QueryFetchPriceResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Result != nil {
		l = m.Result.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryLastFetchPriceIdRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryLastFetchPriceIdResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RequestId != 0 {
		n += 1 + sovQuery(uint64(m.RequestId))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryFetchPriceRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryFetchPriceRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryFetchPriceRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestId", wireType)
			}
			m.RequestId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RequestId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryFetchPriceResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryFetchPriceResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryFetchPriceResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Result == nil {
				m.Result = &FetchPriceResult{}
			}
			if err := m.Result.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryLastFetchPriceIdRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryLastFetchPriceIdRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryLastFetchPriceIdRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryLastFetchPriceIdResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryLastFetchPriceIdResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryLastFetchPriceIdResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestId", wireType)
			}
			m.RequestId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RequestId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
