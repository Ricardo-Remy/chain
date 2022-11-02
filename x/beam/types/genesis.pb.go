// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lum-network/beam/genesis.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

// GenesisState defines the capability module's genesis state.
type GenesisState struct {
	ModuleAccountBalance types.Coin `protobuf:"bytes,1,opt,name=module_account_balance,json=moduleAccountBalance,proto3" json:"module_account_balance" yaml:"module_account_balance"`
	Beams                []*Beam    `protobuf:"bytes,2,rep,name=beams,proto3" json:"beams,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f65af11d46eec55, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetModuleAccountBalance() types.Coin {
	if m != nil {
		return m.ModuleAccountBalance
	}
	return types.Coin{}
}

func (m *GenesisState) GetBeams() []*Beam {
	if m != nil {
		return m.Beams
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "lum.network.beam.GenesisState")
}

func init() { proto.RegisterFile("lum-network/beam/genesis.proto", fileDescriptor_7f65af11d46eec55) }

var fileDescriptor_7f65af11d46eec55 = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xb1, 0x4e, 0xeb, 0x30,
	0x14, 0x86, 0xe3, 0x7b, 0x05, 0x43, 0xca, 0x80, 0xa2, 0xaa, 0x2a, 0x45, 0xb8, 0x55, 0x24, 0xa4,
	0x22, 0x81, 0xad, 0x96, 0x8d, 0x8d, 0x74, 0x60, 0x2f, 0x1b, 0x4b, 0x65, 0x9b, 0xa3, 0x34, 0x22,
	0xf6, 0xa9, 0x6a, 0xa7, 0xd0, 0xb7, 0xe0, 0x61, 0x78, 0x88, 0x8e, 0x1d, 0x99, 0x2a, 0x94, 0xbc,
	0x01, 0x4f, 0x80, 0x52, 0x67, 0x40, 0xc0, 0x62, 0x59, 0xe7, 0x3b, 0xf2, 0xe7, 0xff, 0x0f, 0x69,
	0x5e, 0xe8, 0x2b, 0x03, 0xee, 0x19, 0x97, 0x4f, 0x5c, 0x82, 0xd0, 0x3c, 0x05, 0x03, 0x36, 0xb3,
	0x6c, 0xb1, 0x44, 0x87, 0xd1, 0x71, 0x5e, 0x68, 0xd6, 0x70, 0x56, 0xf3, 0x5e, 0x3b, 0xc5, 0x14,
	0xf7, 0x90, 0xd7, 0x37, 0xbf, 0xd7, 0xa3, 0x0a, 0xad, 0x46, 0xcb, 0xa5, 0xb0, 0xc0, 0x57, 0x23,
	0x09, 0x4e, 0x8c, 0xb8, 0xc2, 0xcc, 0x34, 0xfc, 0xf4, 0x97, 0xa7, 0x3e, 0x3c, 0x8c, 0xdf, 0x48,
	0x78, 0x74, 0xe7, 0xb5, 0xf7, 0x4e, 0x38, 0x88, 0x56, 0x61, 0x47, 0xe3, 0x63, 0x91, 0xc3, 0x4c,
	0x28, 0x85, 0x85, 0x71, 0x33, 0x29, 0x72, 0x61, 0x14, 0x74, 0xc9, 0x80, 0x0c, 0x5b, 0xe3, 0x13,
	0xe6, 0x75, 0xac, 0xd6, 0xb1, 0x46, 0xc7, 0x26, 0x98, 0x99, 0xe4, 0x7c, 0xb3, 0xeb, 0x07, 0x9f,
	0xbb, 0xfe, 0xd9, 0x5a, 0xe8, 0xfc, 0x26, 0xfe, 0xfb, 0x99, 0x78, 0xda, 0xf6, 0xe0, 0xd6, 0xcf,
	0x13, 0x3f, 0x8e, 0x2e, 0xc3, 0x83, 0xfa, 0x5b, 0xb6, 0xfb, 0x6f, 0xf0, 0x7f, 0xd8, 0x1a, 0x77,
	0xd8, 0xcf, 0xf4, 0x2c, 0x01, 0xa1, 0xa7, 0x7e, 0x29, 0x99, 0x6c, 0x4a, 0x4a, 0xb6, 0x25, 0x25,
	0x1f, 0x25, 0x25, 0xaf, 0x15, 0x0d, 0xb6, 0x15, 0x0d, 0xde, 0x2b, 0x1a, 0x3c, 0x5c, 0xa4, 0x99,
	0x9b, 0x17, 0x92, 0x29, 0xd4, 0xfc, 0x7b, 0x70, 0x35, 0x17, 0x99, 0xe1, 0x2f, 0xbe, 0x00, 0xb7,
	0x5e, 0x80, 0x95, 0x87, 0xfb, 0x0a, 0xae, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x3d, 0x1b, 0x45,
	0xc4, 0x89, 0x01, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Beams) > 0 {
		for iNdEx := len(m.Beams) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Beams[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.ModuleAccountBalance.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ModuleAccountBalance.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.Beams) > 0 {
		for _, e := range m.Beams {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ModuleAccountBalance", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ModuleAccountBalance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Beams", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Beams = append(m.Beams, &Beam{})
			if err := m.Beams[len(m.Beams)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
