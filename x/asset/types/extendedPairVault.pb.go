// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: comdex/asset/v1beta1/extendedPairVault.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type ExtendedPairVault struct {
	Id                  uint64                                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AppMappingId        uint64                                 `protobuf:"varint,2,opt,name=app_mapping_id,json=appMappingId,proto3" json:"app_mapping_id,omitempty" yaml:"app_mapping_id"`
	PairId              uint64                                 `protobuf:"varint,3,opt,name=pair_id,json=pairId,proto3" json:"pair_id,omitempty" yaml:"pair_id"`
	LiquidationRatio    github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=liquidation_ratio,json=liquidationRatio,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"liquidation_ratio" yaml:"liquidation_ratio"`
	StabilityFee        github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,5,opt,name=stability_fee,json=stabilityFee,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"stability_fee" yaml:"stability_fee"`
	ClosingFee          github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,6,opt,name=closing_fee,json=closingFee,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"closing_fee" yaml:"closing_fee"`
	LiquidationPenalty  github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,7,opt,name=liquidation_penalty,json=liquidationPenalty,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"liquidation_penalty" yaml:"liquidation_penalty"`
	DrawDownFee         github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,8,opt,name=draw_down_fee,json=drawDownFee,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"draw_down_fee" yaml:"draw_down_fee"`
	IsVaultActive       bool                                   `protobuf:"varint,9,opt,name=is_vault_active,json=isVaultActive,proto3" json:"is_vault_active,omitempty" yaml:"active_flag"`
	DebtCeiling         github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,10,opt,name=debt_ceiling,json=debtCeiling,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"debt_ceiling" yaml:"debt_ceiling"`
	DebtFloor           github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,11,opt,name=debt_floor,json=debtFloor,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"debt_floor" yaml:"debt_floor"`
	IsPsmPair           bool                                   `protobuf:"varint,12,opt,name=is_psm_pair,json=isPsmPair,proto3" json:"is_psm_pair,omitempty" yaml:"is_psm_pair"`
	MinCr               github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,13,opt,name=min_cr,json=minCr,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"min_cr" yaml:"min_cr"`
	PairName            string                                 `protobuf:"bytes,14,opt,name=pair_name,json=pairName,proto3" json:"pair_name,omitempty" yaml:"pair_name"`
	AssetOutOraclePrice bool                                   `protobuf:"varint,15,opt,name=asset_out_oracle_price,json=assetOutOraclePrice,proto3" json:"asset_out_oracle_price,omitempty" yaml:"asset_out_oracle_price"`
	AssetOutPrice       uint64                                 `protobuf:"varint,16,opt,name=asset_out_price,json=assetOutPrice,proto3" json:"asset_out_price,omitempty" yaml:"asset_out_price"`
}

func (m *ExtendedPairVault) Reset()         { *m = ExtendedPairVault{} }
func (m *ExtendedPairVault) String() string { return proto.CompactTextString(m) }
func (*ExtendedPairVault) ProtoMessage()    {}
func (*ExtendedPairVault) Descriptor() ([]byte, []int) {
	return fileDescriptor_23dd38fcddb231cd, []int{0}
}
func (m *ExtendedPairVault) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExtendedPairVault) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ExtendedPairVault.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExtendedPairVault) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtendedPairVault.Merge(m, src)
}
func (m *ExtendedPairVault) XXX_Size() int {
	return m.Size()
}
func (m *ExtendedPairVault) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtendedPairVault.DiscardUnknown(m)
}

var xxx_messageInfo_ExtendedPairVault proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ExtendedPairVault)(nil), "comdex.asset.v1beta1.ExtendedPairVault")
}

func init() {
	proto.RegisterFile("comdex/asset/v1beta1/extendedPairVault.proto", fileDescriptor_23dd38fcddb231cd)
}

var fileDescriptor_23dd38fcddb231cd = []byte{
	// 709 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0xcb, 0x6e, 0xd3, 0x4c,
	0x14, 0xc7, 0x93, 0x7e, 0x6d, 0xda, 0x4c, 0x2e, 0x6d, 0xa7, 0xf9, 0x2a, 0x53, 0x09, 0xbb, 0xcc,
	0x02, 0x55, 0x82, 0xc6, 0xaa, 0x90, 0x58, 0xb0, 0xa0, 0x22, 0xbd, 0x48, 0x45, 0x40, 0xc3, 0x2c,
	0xba, 0x60, 0x33, 0x9a, 0xd8, 0x93, 0x74, 0xa8, 0xed, 0x31, 0xf6, 0xa4, 0x69, 0x16, 0xbc, 0x03,
	0x8f, 0xc1, 0x13, 0xf0, 0x0c, 0x5d, 0x76, 0x89, 0x58, 0x58, 0x90, 0xbe, 0x81, 0x9f, 0x00, 0xcd,
	0xd8, 0xa5, 0x4e, 0xcb, 0x26, 0x62, 0x63, 0xcf, 0xb9, 0xcc, 0xff, 0x77, 0xe6, 0xd8, 0x73, 0xc0,
	0x53, 0x47, 0xf8, 0x2e, 0xbb, 0xb0, 0x69, 0x1c, 0x33, 0x69, 0x9f, 0xef, 0xf4, 0x98, 0xa4, 0x3b,
	0x36, 0xbb, 0x90, 0x2c, 0x70, 0x99, 0xdb, 0xa5, 0x3c, 0x3a, 0xa1, 0x43, 0x4f, 0xb6, 0xc3, 0x48,
	0x48, 0x01, 0x5b, 0x59, 0x76, 0x5b, 0x67, 0xb7, 0xf3, 0xec, 0x8d, 0xd6, 0x40, 0x0c, 0x84, 0x4e,
	0xb0, 0xd5, 0x2a, 0xcb, 0x45, 0xdf, 0x00, 0x58, 0x3d, 0xb8, 0xab, 0x03, 0x9b, 0x60, 0x8e, 0xbb,
	0x46, 0x79, 0xb3, 0xbc, 0x35, 0x8f, 0xe7, 0xb8, 0x0b, 0x77, 0x41, 0x93, 0x86, 0x21, 0xf1, 0x69,
	0x18, 0xf2, 0x60, 0x40, 0xb8, 0x6b, 0xcc, 0xa9, 0x58, 0xe7, 0x41, 0x9a, 0x58, 0xff, 0x8f, 0xa9,
	0xef, 0xbd, 0x40, 0xd3, 0x71, 0x84, 0xeb, 0x34, 0x0c, 0xdf, 0x66, 0xf6, 0x91, 0x0b, 0x9f, 0x80,
	0xc5, 0x90, 0xf2, 0x48, 0xed, 0xfc, 0x4f, 0xef, 0x84, 0x69, 0x62, 0x35, 0xb3, 0x9d, 0x79, 0x00,
	0xe1, 0x8a, 0x5a, 0x1d, 0xb9, 0x70, 0x04, 0x56, 0x3d, 0xfe, 0x69, 0xc8, 0x5d, 0x2a, 0xb9, 0x08,
	0x48, 0xa4, 0x5e, 0xc6, 0xfc, 0x66, 0x79, 0xab, 0xda, 0x79, 0x7d, 0x99, 0x58, 0xa5, 0x1f, 0x89,
	0xf5, 0x78, 0xc0, 0xe5, 0xe9, 0xb0, 0xd7, 0x76, 0x84, 0x6f, 0x3b, 0x22, 0xf6, 0x45, 0x9c, 0xbf,
	0xb6, 0x63, 0xf7, 0xcc, 0x96, 0xe3, 0x90, 0xc5, 0xed, 0x7d, 0xe6, 0xa4, 0x89, 0x65, 0x64, 0x90,
	0x7b, 0x82, 0x08, 0xaf, 0x14, 0x7c, 0x58, 0x3d, 0xe1, 0x19, 0x68, 0xc4, 0x92, 0xf6, 0xb8, 0xc7,
	0xe5, 0x98, 0xf4, 0x19, 0x33, 0x16, 0x34, 0xf4, 0x70, 0x66, 0x68, 0x2b, 0x83, 0x4e, 0x89, 0x21,
	0x5c, 0xff, 0x63, 0x1f, 0x32, 0x06, 0x19, 0xa8, 0x39, 0x9e, 0x88, 0x55, 0xbf, 0x14, 0xaa, 0xa2,
	0x51, 0xfb, 0x33, 0xa3, 0x60, 0x86, 0x2a, 0x48, 0x21, 0x0c, 0x72, 0x4b, 0x61, 0x3e, 0x83, 0xb5,
	0xe2, 0xd9, 0x43, 0x16, 0x50, 0x4f, 0x8e, 0x8d, 0x45, 0x8d, 0x7b, 0x33, 0x33, 0x6e, 0xe3, 0x7e,
	0x3b, 0x73, 0x49, 0x84, 0x61, 0xc1, 0xdb, 0xcd, 0x9c, 0xf0, 0x23, 0x68, 0xb8, 0x11, 0x1d, 0x11,
	0x57, 0x8c, 0x02, 0x7d, 0xce, 0xa5, 0x7f, 0x6b, 0xe9, 0x94, 0x18, 0xc2, 0x35, 0x65, 0xef, 0x8b,
	0x51, 0xa0, 0x8e, 0xfa, 0x12, 0x2c, 0xf3, 0x98, 0x9c, 0xab, 0x3f, 0x98, 0x50, 0x47, 0xf2, 0x73,
	0x66, 0x54, 0x37, 0xcb, 0x5b, 0x4b, 0x9d, 0xf5, 0xdb, 0x3e, 0x65, 0x7e, 0xd2, 0xf7, 0xe8, 0x00,
	0xe1, 0x06, 0x8f, 0xf5, 0xff, 0xfe, 0x4a, 0x3b, 0xe1, 0x29, 0xa8, 0xbb, 0xac, 0x27, 0x89, 0xc3,
	0xb8, 0xc7, 0x83, 0x81, 0x01, 0x74, 0xa9, 0x07, 0x33, 0x94, 0x7a, 0x14, 0xc8, 0x34, 0xb1, 0xd6,
	0xf2, 0x52, 0x0b, 0x5a, 0xaa, 0x52, 0xd6, 0x93, 0x7b, 0x99, 0x05, 0x7b, 0x00, 0xe8, 0x68, 0xdf,
	0x13, 0x22, 0x32, 0x6a, 0x9a, 0xb3, 0x37, 0x33, 0x67, 0xb5, 0xc0, 0xd1, 0x4a, 0x08, 0x57, 0x95,
	0x71, 0xa8, 0xd6, 0xf0, 0x39, 0xa8, 0xf1, 0x98, 0x84, 0xb1, 0x4f, 0xd4, 0xb5, 0x32, 0xea, 0x77,
	0x3b, 0x51, 0x08, 0x22, 0x5c, 0xe5, 0x71, 0x37, 0xf6, 0xd5, 0x00, 0x80, 0x27, 0xa0, 0xe2, 0xf3,
	0x80, 0x38, 0x91, 0xd1, 0xd0, 0x75, 0xed, 0xce, 0xfc, 0xa9, 0x1a, 0x19, 0x20, 0x53, 0x41, 0x78,
	0xc1, 0xe7, 0xc1, 0x5e, 0x04, 0x77, 0x40, 0x55, 0xdf, 0xf4, 0x80, 0xfa, 0xcc, 0x68, 0x6a, 0xe9,
	0x56, 0x9a, 0x58, 0x2b, 0x85, 0x21, 0xa0, 0x42, 0x08, 0x2f, 0xa9, 0xf5, 0x3b, 0xea, 0x33, 0x78,
	0x02, 0xd6, 0xf5, 0x0c, 0x23, 0x62, 0x28, 0x89, 0x88, 0xa8, 0xe3, 0x31, 0x12, 0x46, 0xdc, 0x61,
	0xc6, 0xb2, 0x3e, 0xcd, 0xa3, 0x34, 0xb1, 0x1e, 0xe6, 0xdf, 0xf5, 0xaf, 0x79, 0x08, 0xaf, 0xe9,
	0xc0, 0xf1, 0x50, 0x1e, 0x6b, 0x77, 0x57, 0x79, 0x61, 0x07, 0x2c, 0xdf, 0xe6, 0x67, 0x82, 0x2b,
	0x7a, 0x2a, 0x6d, 0xa4, 0x89, 0xb5, 0x7e, 0x57, 0x30, 0x57, 0x6a, 0xdc, 0x28, 0x69, 0x8d, 0xce,
	0xfb, 0xcb, 0x5f, 0x66, 0xe9, 0xeb, 0xc4, 0x2c, 0x5d, 0x4e, 0xcc, 0xf2, 0xd5, 0xc4, 0x2c, 0xff,
	0x9c, 0x98, 0xe5, 0x2f, 0xd7, 0x66, 0xe9, 0xea, 0xda, 0x2c, 0x7d, 0xbf, 0x36, 0x4b, 0x1f, 0xec,
	0xa9, 0x86, 0xa9, 0x89, 0xbc, 0x2d, 0xfa, 0x7d, 0xee, 0x70, 0xea, 0xe5, 0xb6, 0x7d, 0x33, 0xd1,
	0x75, 0xf7, 0x7a, 0x15, 0x3d, 0x92, 0x9f, 0xfd, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x56, 0x21, 0x85,
	0x4e, 0xee, 0x05, 0x00, 0x00,
}

func (m *ExtendedPairVault) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExtendedPairVault) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ExtendedPairVault) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.AssetOutPrice != 0 {
		i = encodeVarintExtendedPairVault(dAtA, i, uint64(m.AssetOutPrice))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x80
	}
	if m.AssetOutOraclePrice {
		i--
		if m.AssetOutOraclePrice {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x78
	}
	if len(m.PairName) > 0 {
		i -= len(m.PairName)
		copy(dAtA[i:], m.PairName)
		i = encodeVarintExtendedPairVault(dAtA, i, uint64(len(m.PairName)))
		i--
		dAtA[i] = 0x72
	}
	{
		size := m.MinCr.Size()
		i -= size
		if _, err := m.MinCr.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintExtendedPairVault(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x6a
	if m.IsPsmPair {
		i--
		if m.IsPsmPair {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x60
	}
	{
		size := m.DebtFloor.Size()
		i -= size
		if _, err := m.DebtFloor.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintExtendedPairVault(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x5a
	{
		size := m.DebtCeiling.Size()
		i -= size
		if _, err := m.DebtCeiling.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintExtendedPairVault(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x52
	if m.IsVaultActive {
		i--
		if m.IsVaultActive {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x48
	}
	{
		size := m.DrawDownFee.Size()
		i -= size
		if _, err := m.DrawDownFee.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintExtendedPairVault(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	{
		size := m.LiquidationPenalty.Size()
		i -= size
		if _, err := m.LiquidationPenalty.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintExtendedPairVault(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	{
		size := m.ClosingFee.Size()
		i -= size
		if _, err := m.ClosingFee.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintExtendedPairVault(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size := m.StabilityFee.Size()
		i -= size
		if _, err := m.StabilityFee.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintExtendedPairVault(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.LiquidationRatio.Size()
		i -= size
		if _, err := m.LiquidationRatio.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintExtendedPairVault(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.PairId != 0 {
		i = encodeVarintExtendedPairVault(dAtA, i, uint64(m.PairId))
		i--
		dAtA[i] = 0x18
	}
	if m.AppMappingId != 0 {
		i = encodeVarintExtendedPairVault(dAtA, i, uint64(m.AppMappingId))
		i--
		dAtA[i] = 0x10
	}
	if m.Id != 0 {
		i = encodeVarintExtendedPairVault(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintExtendedPairVault(dAtA []byte, offset int, v uint64) int {
	offset -= sovExtendedPairVault(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ExtendedPairVault) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovExtendedPairVault(uint64(m.Id))
	}
	if m.AppMappingId != 0 {
		n += 1 + sovExtendedPairVault(uint64(m.AppMappingId))
	}
	if m.PairId != 0 {
		n += 1 + sovExtendedPairVault(uint64(m.PairId))
	}
	l = m.LiquidationRatio.Size()
	n += 1 + l + sovExtendedPairVault(uint64(l))
	l = m.StabilityFee.Size()
	n += 1 + l + sovExtendedPairVault(uint64(l))
	l = m.ClosingFee.Size()
	n += 1 + l + sovExtendedPairVault(uint64(l))
	l = m.LiquidationPenalty.Size()
	n += 1 + l + sovExtendedPairVault(uint64(l))
	l = m.DrawDownFee.Size()
	n += 1 + l + sovExtendedPairVault(uint64(l))
	if m.IsVaultActive {
		n += 2
	}
	l = m.DebtCeiling.Size()
	n += 1 + l + sovExtendedPairVault(uint64(l))
	l = m.DebtFloor.Size()
	n += 1 + l + sovExtendedPairVault(uint64(l))
	if m.IsPsmPair {
		n += 2
	}
	l = m.MinCr.Size()
	n += 1 + l + sovExtendedPairVault(uint64(l))
	l = len(m.PairName)
	if l > 0 {
		n += 1 + l + sovExtendedPairVault(uint64(l))
	}
	if m.AssetOutOraclePrice {
		n += 2
	}
	if m.AssetOutPrice != 0 {
		n += 2 + sovExtendedPairVault(uint64(m.AssetOutPrice))
	}
	return n
}

func sovExtendedPairVault(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozExtendedPairVault(x uint64) (n int) {
	return sovExtendedPairVault(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ExtendedPairVault) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExtendedPairVault
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
			return fmt.Errorf("proto: ExtendedPairVault: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExtendedPairVault: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtendedPairVault
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
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppMappingId", wireType)
			}
			m.AppMappingId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtendedPairVault
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AppMappingId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PairId", wireType)
			}
			m.PairId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtendedPairVault
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PairId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LiquidationRatio", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtendedPairVault
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
				return ErrInvalidLengthExtendedPairVault
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExtendedPairVault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LiquidationRatio.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StabilityFee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtendedPairVault
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
				return ErrInvalidLengthExtendedPairVault
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExtendedPairVault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.StabilityFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClosingFee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtendedPairVault
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
				return ErrInvalidLengthExtendedPairVault
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExtendedPairVault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ClosingFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LiquidationPenalty", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtendedPairVault
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
				return ErrInvalidLengthExtendedPairVault
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExtendedPairVault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LiquidationPenalty.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DrawDownFee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtendedPairVault
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
				return ErrInvalidLengthExtendedPairVault
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExtendedPairVault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DrawDownFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsVaultActive", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtendedPairVault
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsVaultActive = bool(v != 0)
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DebtCeiling", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtendedPairVault
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
				return ErrInvalidLengthExtendedPairVault
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExtendedPairVault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DebtCeiling.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DebtFloor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtendedPairVault
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
				return ErrInvalidLengthExtendedPairVault
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExtendedPairVault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DebtFloor.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsPsmPair", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtendedPairVault
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsPsmPair = bool(v != 0)
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinCr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtendedPairVault
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
				return ErrInvalidLengthExtendedPairVault
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExtendedPairVault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinCr.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PairName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtendedPairVault
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
				return ErrInvalidLengthExtendedPairVault
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExtendedPairVault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PairName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 15:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetOutOraclePrice", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtendedPairVault
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.AssetOutOraclePrice = bool(v != 0)
		case 16:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetOutPrice", wireType)
			}
			m.AssetOutPrice = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtendedPairVault
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AssetOutPrice |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipExtendedPairVault(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthExtendedPairVault
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
func skipExtendedPairVault(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowExtendedPairVault
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
					return 0, ErrIntOverflowExtendedPairVault
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
					return 0, ErrIntOverflowExtendedPairVault
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
				return 0, ErrInvalidLengthExtendedPairVault
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupExtendedPairVault
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthExtendedPairVault
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthExtendedPairVault        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowExtendedPairVault          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupExtendedPairVault = fmt.Errorf("proto: unexpected end of group")
)