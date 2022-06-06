// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: comdex/asset/v1beta1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

type MsgAddAssetRequest struct {
	From     string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty" yaml:"from"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" yaml:"name"`
	Denom    string `protobuf:"bytes,3,opt,name=denom,proto3" json:"denom,omitempty" yaml:"denom"`
	Decimals int64  `protobuf:"varint,4,opt,name=decimals,proto3" json:"decimals,omitempty" yaml:"decimals"`
}

func (m *MsgAddAssetRequest) Reset()         { *m = MsgAddAssetRequest{} }
func (m *MsgAddAssetRequest) String() string { return proto.CompactTextString(m) }
func (*MsgAddAssetRequest) ProtoMessage()    {}
func (*MsgAddAssetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_78d9fd76d8e93e29, []int{0}
}
func (m *MsgAddAssetRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAddAssetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAddAssetRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAddAssetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAddAssetRequest.Merge(m, src)
}
func (m *MsgAddAssetRequest) XXX_Size() int {
	return m.Size()
}
func (m *MsgAddAssetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAddAssetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAddAssetRequest proto.InternalMessageInfo

type MsgAddAssetResponse struct {
}

func (m *MsgAddAssetResponse) Reset()         { *m = MsgAddAssetResponse{} }
func (m *MsgAddAssetResponse) String() string { return proto.CompactTextString(m) }
func (*MsgAddAssetResponse) ProtoMessage()    {}
func (*MsgAddAssetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_78d9fd76d8e93e29, []int{1}
}
func (m *MsgAddAssetResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAddAssetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAddAssetResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAddAssetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAddAssetResponse.Merge(m, src)
}
func (m *MsgAddAssetResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgAddAssetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAddAssetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAddAssetResponse proto.InternalMessageInfo

type MsgUpdateAssetRequest struct {
	From     string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty" yaml:"from"`
	Id       uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty" yaml:"name"`
	Denom    string `protobuf:"bytes,4,opt,name=denom,proto3" json:"denom,omitempty" yaml:"denom"`
	Decimals int64  `protobuf:"varint,5,opt,name=decimals,proto3" json:"decimals,omitempty" yaml:"decimals"`
}

func (m *MsgUpdateAssetRequest) Reset()         { *m = MsgUpdateAssetRequest{} }
func (m *MsgUpdateAssetRequest) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateAssetRequest) ProtoMessage()    {}
func (*MsgUpdateAssetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_78d9fd76d8e93e29, []int{2}
}
func (m *MsgUpdateAssetRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateAssetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateAssetRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateAssetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateAssetRequest.Merge(m, src)
}
func (m *MsgUpdateAssetRequest) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateAssetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateAssetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateAssetRequest proto.InternalMessageInfo

type MsgUpdateAssetResponse struct {
}

func (m *MsgUpdateAssetResponse) Reset()         { *m = MsgUpdateAssetResponse{} }
func (m *MsgUpdateAssetResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateAssetResponse) ProtoMessage()    {}
func (*MsgUpdateAssetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_78d9fd76d8e93e29, []int{3}
}
func (m *MsgUpdateAssetResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateAssetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateAssetResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateAssetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateAssetResponse.Merge(m, src)
}
func (m *MsgUpdateAssetResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateAssetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateAssetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateAssetResponse proto.InternalMessageInfo

type MsgAddPairRequest struct {
	From     string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty" yaml:"from"`
	AssetIn  uint64 `protobuf:"varint,2,opt,name=asset_in,json=assetIn,proto3" json:"asset_in,omitempty" yaml:"asset_in"`
	AssetOut uint64 `protobuf:"varint,3,opt,name=asset_out,json=assetOut,proto3" json:"asset_out,omitempty" yaml:"asset_out"`
}

func (m *MsgAddPairRequest) Reset()         { *m = MsgAddPairRequest{} }
func (m *MsgAddPairRequest) String() string { return proto.CompactTextString(m) }
func (*MsgAddPairRequest) ProtoMessage()    {}
func (*MsgAddPairRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_78d9fd76d8e93e29, []int{4}
}
func (m *MsgAddPairRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAddPairRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAddPairRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAddPairRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAddPairRequest.Merge(m, src)
}
func (m *MsgAddPairRequest) XXX_Size() int {
	return m.Size()
}
func (m *MsgAddPairRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAddPairRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAddPairRequest proto.InternalMessageInfo

type MsgAddPairResponse struct {
}

func (m *MsgAddPairResponse) Reset()         { *m = MsgAddPairResponse{} }
func (m *MsgAddPairResponse) String() string { return proto.CompactTextString(m) }
func (*MsgAddPairResponse) ProtoMessage()    {}
func (*MsgAddPairResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_78d9fd76d8e93e29, []int{5}
}
func (m *MsgAddPairResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAddPairResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAddPairResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAddPairResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAddPairResponse.Merge(m, src)
}
func (m *MsgAddPairResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgAddPairResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAddPairResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAddPairResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgAddAssetRequest)(nil), "comdex.asset.v1beta1.MsgAddAssetRequest")
	proto.RegisterType((*MsgAddAssetResponse)(nil), "comdex.asset.v1beta1.MsgAddAssetResponse")
	proto.RegisterType((*MsgUpdateAssetRequest)(nil), "comdex.asset.v1beta1.MsgUpdateAssetRequest")
	proto.RegisterType((*MsgUpdateAssetResponse)(nil), "comdex.asset.v1beta1.MsgUpdateAssetResponse")
	proto.RegisterType((*MsgAddPairRequest)(nil), "comdex.asset.v1beta1.MsgAddPairRequest")
	proto.RegisterType((*MsgAddPairResponse)(nil), "comdex.asset.v1beta1.MsgAddPairResponse")
}

func init() { proto.RegisterFile("comdex/asset/v1beta1/tx.proto", fileDescriptor_78d9fd76d8e93e29) }

var fileDescriptor_78d9fd76d8e93e29 = []byte{
	// 475 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xc1, 0x6e, 0xd3, 0x30,
	0x18, 0xc7, 0xe3, 0x36, 0x83, 0xed, 0x1b, 0xda, 0x86, 0xd7, 0xa1, 0xa8, 0x12, 0xe9, 0x64, 0x24,
	0x28, 0x02, 0x12, 0x15, 0x6e, 0xdc, 0xd6, 0x1b, 0x87, 0x0a, 0xb0, 0xc4, 0x85, 0xcb, 0xe4, 0xd6,
	0x6e, 0xb0, 0x68, 0xe2, 0x52, 0x3b, 0x68, 0x7b, 0x0b, 0xee, 0xbc, 0x00, 0x6f, 0xb0, 0x57, 0x98,
	0x38, 0xed, 0xc8, 0xa9, 0x82, 0xf6, 0x0d, 0xfa, 0x04, 0x28, 0x76, 0x3b, 0xb2, 0x0d, 0x58, 0x76,
	0x73, 0xbe, 0xff, 0xdf, 0xdf, 0xf7, 0xfd, 0xec, 0x2f, 0x86, 0xfb, 0x03, 0x95, 0x72, 0x71, 0x14,
	0x33, 0xad, 0x85, 0x89, 0x3f, 0x77, 0xfa, 0xc2, 0xb0, 0x4e, 0x6c, 0x8e, 0xa2, 0xf1, 0x44, 0x19,
	0x85, 0x1b, 0x4e, 0x8e, 0xac, 0x1c, 0x2d, 0xe5, 0x66, 0x23, 0x51, 0x89, 0xb2, 0x86, 0xb8, 0x58,
	0x39, 0x2f, 0x39, 0x41, 0x80, 0x7b, 0x3a, 0x39, 0xe0, 0xfc, 0xa0, 0x70, 0x53, 0xf1, 0x29, 0x17,
	0xda, 0xe0, 0x07, 0xe0, 0x0f, 0x27, 0x2a, 0x0d, 0xd0, 0x3e, 0x6a, 0x6f, 0x74, 0xb7, 0x17, 0xd3,
	0xd6, 0xe6, 0x31, 0x4b, 0x47, 0x2f, 0x49, 0x11, 0x25, 0xd4, 0x8a, 0x85, 0x29, 0x63, 0xa9, 0x08,
	0x6a, 0x97, 0x4d, 0x45, 0x94, 0x50, 0x2b, 0xe2, 0x87, 0xb0, 0xc6, 0x45, 0xa6, 0xd2, 0xa0, 0x6e,
	0x5d, 0x3b, 0x8b, 0x69, 0xeb, 0x8e, 0x73, 0xd9, 0x30, 0xa1, 0x4e, 0xc6, 0x31, 0xac, 0x73, 0x31,
	0x90, 0x29, 0x1b, 0xe9, 0xc0, 0xdf, 0x47, 0xed, 0x7a, 0x77, 0x77, 0x31, 0x6d, 0x6d, 0xaf, 0xac,
	0x4e, 0x21, 0xf4, 0xdc, 0x44, 0xf6, 0x60, 0xf7, 0x42, 0xe3, 0x7a, 0xac, 0x32, 0x2d, 0xc8, 0x77,
	0x04, 0x7b, 0x3d, 0x9d, 0xbc, 0x1b, 0x73, 0x66, 0xc4, 0xcd, 0x99, 0xb6, 0xa0, 0x26, 0xb9, 0x25,
	0xf2, 0x69, 0x4d, 0xf2, 0x73, 0xc6, 0x7a, 0x25, 0x46, 0xbf, 0x3a, 0xe3, 0x5a, 0x15, 0xc6, 0x00,
	0xee, 0x5d, 0x66, 0x59, 0x62, 0x7e, 0x45, 0x70, 0xd7, 0xe1, 0xbf, 0x61, 0x72, 0x72, 0x23, 0xc4,
	0x08, 0xd6, 0xed, 0x64, 0x1c, 0xca, 0xcc, 0x81, 0x96, 0xbb, 0x58, 0x29, 0x84, 0xde, 0xb6, 0xcb,
	0x57, 0x19, 0xee, 0xc0, 0x86, 0x8b, 0xaa, 0xdc, 0xd8, 0x73, 0xf0, 0xbb, 0x8d, 0xc5, 0xb4, 0xb5,
	0x53, 0xde, 0xa0, 0x72, 0x43, 0xa8, 0x4b, 0xfb, 0x3a, 0x37, 0xa4, 0xb1, 0x1a, 0x2a, 0xd7, 0x9c,
	0xeb, 0xf9, 0xf9, 0x49, 0x0d, 0xea, 0x3d, 0x9d, 0xe0, 0x3e, 0x6c, 0x96, 0x6e, 0x0e, 0xb7, 0xa3,
	0xbf, 0xcd, 0x6b, 0x74, 0x75, 0x2a, 0x9b, 0x8f, 0x2b, 0x38, 0x5d, 0x2d, 0xfc, 0x11, 0xb6, 0x2e,
	0x9e, 0x1c, 0x7e, 0xf2, 0xcf, 0xcd, 0x57, 0x67, 0xa5, 0xf9, 0xb4, 0x9a, 0x79, 0x59, 0xec, 0x10,
	0xe0, 0x0f, 0x2e, 0x7e, 0xf4, 0xbf, 0x2e, 0x4b, 0xb7, 0xd5, 0x6c, 0x5f, 0x6f, 0x74, 0x05, 0xba,
	0x6f, 0x4f, 0x7f, 0x85, 0xde, 0xb7, 0x59, 0xe8, 0x9d, 0xce, 0x42, 0x74, 0x36, 0x0b, 0xd1, 0xcf,
	0x59, 0x88, 0xbe, 0xcc, 0x43, 0xef, 0x6c, 0x1e, 0x7a, 0x3f, 0xe6, 0xa1, 0xf7, 0x3e, 0x4e, 0xa4,
	0xf9, 0x90, 0xf7, 0x8b, 0x8c, 0xb1, 0xcb, 0xfa, 0x4c, 0x0d, 0x87, 0x72, 0x20, 0xd9, 0x68, 0xf9,
	0x1d, 0xaf, 0xde, 0x0b, 0x73, 0x3c, 0x16, 0xba, 0x7f, 0xcb, 0xfe, 0xff, 0x2f, 0x7e, 0x07, 0x00,
	0x00, 0xff, 0xff, 0xb6, 0xa2, 0xad, 0x8f, 0x4c, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	MsgAddAsset(ctx context.Context, in *MsgAddAssetRequest, opts ...grpc.CallOption) (*MsgAddAssetResponse, error)
	MsgUpdateAsset(ctx context.Context, in *MsgUpdateAssetRequest, opts ...grpc.CallOption) (*MsgUpdateAssetResponse, error)
	MsgAddPair(ctx context.Context, in *MsgAddPairRequest, opts ...grpc.CallOption) (*MsgAddPairResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) MsgAddAsset(ctx context.Context, in *MsgAddAssetRequest, opts ...grpc.CallOption) (*MsgAddAssetResponse, error) {
	out := new(MsgAddAssetResponse)
	err := c.cc.Invoke(ctx, "/comdex.asset.v1beta1.Msg/MsgAddAsset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) MsgUpdateAsset(ctx context.Context, in *MsgUpdateAssetRequest, opts ...grpc.CallOption) (*MsgUpdateAssetResponse, error) {
	out := new(MsgUpdateAssetResponse)
	err := c.cc.Invoke(ctx, "/comdex.asset.v1beta1.Msg/MsgUpdateAsset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) MsgAddPair(ctx context.Context, in *MsgAddPairRequest, opts ...grpc.CallOption) (*MsgAddPairResponse, error) {
	out := new(MsgAddPairResponse)
	err := c.cc.Invoke(ctx, "/comdex.asset.v1beta1.Msg/MsgAddPair", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	MsgAddAsset(context.Context, *MsgAddAssetRequest) (*MsgAddAssetResponse, error)
	MsgUpdateAsset(context.Context, *MsgUpdateAssetRequest) (*MsgUpdateAssetResponse, error)
	MsgAddPair(context.Context, *MsgAddPairRequest) (*MsgAddPairResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) MsgAddAsset(ctx context.Context, req *MsgAddAssetRequest) (*MsgAddAssetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MsgAddAsset not implemented")
}
func (*UnimplementedMsgServer) MsgUpdateAsset(ctx context.Context, req *MsgUpdateAssetRequest) (*MsgUpdateAssetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MsgUpdateAsset not implemented")
}
func (*UnimplementedMsgServer) MsgAddPair(ctx context.Context, req *MsgAddPairRequest) (*MsgAddPairResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MsgAddPair not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_MsgAddAsset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgAddAssetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).MsgAddAsset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comdex.asset.v1beta1.Msg/MsgAddAsset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).MsgAddAsset(ctx, req.(*MsgAddAssetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_MsgUpdateAsset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateAssetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).MsgUpdateAsset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comdex.asset.v1beta1.Msg/MsgUpdateAsset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).MsgUpdateAsset(ctx, req.(*MsgUpdateAssetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_MsgAddPair_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgAddPairRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).MsgAddPair(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comdex.asset.v1beta1.Msg/MsgAddPair",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).MsgAddPair(ctx, req.(*MsgAddPairRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "comdex.asset.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MsgAddAsset",
			Handler:    _Msg_MsgAddAsset_Handler,
		},
		{
			MethodName: "MsgUpdateAsset",
			Handler:    _Msg_MsgUpdateAsset_Handler,
		},
		{
			MethodName: "MsgAddPair",
			Handler:    _Msg_MsgAddPair_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "comdex/asset/v1beta1/tx.proto",
}

func (m *MsgAddAssetRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAddAssetRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAddAssetRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Decimals != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Decimals))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintTx(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgAddAssetResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAddAssetResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAddAssetResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgUpdateAssetRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateAssetRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateAssetRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Decimals != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Decimals))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Id != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintTx(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateAssetResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateAssetResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateAssetResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgAddPairRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAddPairRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAddPairRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.AssetOut != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.AssetOut))
		i--
		dAtA[i] = 0x18
	}
	if m.AssetIn != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.AssetIn))
		i--
		dAtA[i] = 0x10
	}
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintTx(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgAddPairResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAddPairResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAddPairResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgAddAssetRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Decimals != 0 {
		n += 1 + sovTx(uint64(m.Decimals))
	}
	return n
}

func (m *MsgAddAssetResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgUpdateAssetRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovTx(uint64(m.Id))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Decimals != 0 {
		n += 1 + sovTx(uint64(m.Decimals))
	}
	return n
}

func (m *MsgUpdateAssetResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgAddPairRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.AssetIn != 0 {
		n += 1 + sovTx(uint64(m.AssetIn))
	}
	if m.AssetOut != 0 {
		n += 1 + sovTx(uint64(m.AssetOut))
	}
	return n
}

func (m *MsgAddPairResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgAddAssetRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgAddAssetRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAddAssetRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Decimals", wireType)
			}
			m.Decimals = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Decimals |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgAddAssetResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgAddAssetResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAddAssetResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgUpdateAssetRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgUpdateAssetRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateAssetRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Decimals", wireType)
			}
			m.Decimals = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Decimals |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgUpdateAssetResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgUpdateAssetResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateAssetResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgAddPairRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgAddPairRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAddPairRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetIn", wireType)
			}
			m.AssetIn = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AssetIn |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetOut", wireType)
			}
			m.AssetOut = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AssetOut |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgAddPairResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgAddPairResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAddPairResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)