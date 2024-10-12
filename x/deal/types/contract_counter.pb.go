// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: example/deal/contract_counter.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ContractCounter struct {
	DealId  string `protobuf:"bytes,1,opt,name=dealId,proto3" json:"dealId,omitempty"`
	IdValue uint64 `protobuf:"varint,2,opt,name=idValue,proto3" json:"idValue,omitempty"`
}

func (m *ContractCounter) Reset()         { *m = ContractCounter{} }
func (m *ContractCounter) String() string { return proto.CompactTextString(m) }
func (*ContractCounter) ProtoMessage()    {}
func (*ContractCounter) Descriptor() ([]byte, []int) {
	return fileDescriptor_51be192cf6eae170, []int{0}
}
func (m *ContractCounter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ContractCounter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ContractCounter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ContractCounter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContractCounter.Merge(m, src)
}
func (m *ContractCounter) XXX_Size() int {
	return m.Size()
}
func (m *ContractCounter) XXX_DiscardUnknown() {
	xxx_messageInfo_ContractCounter.DiscardUnknown(m)
}

var xxx_messageInfo_ContractCounter proto.InternalMessageInfo

func (m *ContractCounter) GetDealId() string {
	if m != nil {
		return m.DealId
	}
	return ""
}

func (m *ContractCounter) GetIdValue() uint64 {
	if m != nil {
		return m.IdValue
	}
	return 0
}

func init() {
	proto.RegisterType((*ContractCounter)(nil), "example.deal.ContractCounter")
}

func init() {
	proto.RegisterFile("example/deal/contract_counter.proto", fileDescriptor_51be192cf6eae170)
}

var fileDescriptor_51be192cf6eae170 = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8e, 0xb1, 0x4a, 0xfc, 0x40,
	0x10, 0xc6, 0xb3, 0x7f, 0xfe, 0x9c, 0x18, 0x04, 0x31, 0x1c, 0x12, 0x0f, 0x59, 0x0e, 0x6d, 0x0e,
	0x8b, 0x0c, 0x87, 0x6f, 0x60, 0x2a, 0xdb, 0x2b, 0x2c, 0x6c, 0x64, 0x92, 0x0c, 0x61, 0x21, 0xd9,
	0x59, 0x77, 0x37, 0x72, 0xf7, 0x16, 0x3e, 0x96, 0xe5, 0x95, 0x96, 0x92, 0xbc, 0x88, 0x24, 0xd9,
	0x2b, 0x6c, 0x86, 0xef, 0x9b, 0xf9, 0x31, 0xdf, 0x17, 0xdf, 0xd3, 0x1e, 0x5b, 0xd3, 0x10, 0x54,
	0x84, 0x0d, 0x94, 0xac, 0xbd, 0xc5, 0xd2, 0xbf, 0x95, 0xdc, 0x69, 0x4f, 0x36, 0x33, 0x96, 0x3d,
	0x27, 0x17, 0x01, 0xca, 0x46, 0x68, 0x75, 0x85, 0xad, 0xd2, 0x0c, 0xd3, 0x9c, 0x81, 0xd5, 0xb2,
	0xe6, 0x9a, 0x27, 0x09, 0xa3, 0x0a, 0xdb, 0xdb, 0x9a, 0xb9, 0x6e, 0x08, 0xd0, 0x28, 0x40, 0xad,
	0xd9, 0xa3, 0x57, 0xac, 0x5d, 0xb8, 0x3e, 0x94, 0xec, 0x5a, 0x76, 0x50, 0xa0, 0x23, 0x78, 0xef,
	0xc8, 0x1e, 0xe0, 0x63, 0x5b, 0x90, 0xc7, 0x2d, 0x18, 0xac, 0x95, 0x9e, 0xe0, 0xc0, 0xde, 0xfc,
	0x69, 0x69, 0xd0, 0x62, 0x1b, 0xde, 0xdc, 0xe5, 0xf1, 0x65, 0x1e, 0x5a, 0xe7, 0x73, 0xe9, 0xe4,
	0x3a, 0x5e, 0x8c, 0xdc, 0x73, 0x95, 0x8a, 0xb5, 0xd8, 0x9c, 0xef, 0x82, 0x4b, 0xd2, 0xf8, 0x4c,
	0x55, 0x2f, 0xd8, 0x74, 0x94, 0xfe, 0x5b, 0x8b, 0xcd, 0xff, 0xdd, 0xc9, 0x3e, 0x65, 0x5f, 0xbd,
	0x14, 0xc7, 0x5e, 0x8a, 0x9f, 0x5e, 0x8a, 0xcf, 0x41, 0x46, 0xc7, 0x41, 0x46, 0xdf, 0x83, 0x8c,
	0x5e, 0x97, 0xa7, 0xe4, 0xfd, 0x9c, 0xed, 0x0f, 0x86, 0x5c, 0xb1, 0x98, 0xb2, 0x1f, 0x7f, 0x03,
	0x00, 0x00, 0xff, 0xff, 0x48, 0x1b, 0xee, 0xd7, 0x3e, 0x01, 0x00, 0x00,
}

func (m *ContractCounter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContractCounter) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ContractCounter) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.IdValue != 0 {
		i = encodeVarintContractCounter(dAtA, i, uint64(m.IdValue))
		i--
		dAtA[i] = 0x10
	}
	if len(m.DealId) > 0 {
		i -= len(m.DealId)
		copy(dAtA[i:], m.DealId)
		i = encodeVarintContractCounter(dAtA, i, uint64(len(m.DealId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintContractCounter(dAtA []byte, offset int, v uint64) int {
	offset -= sovContractCounter(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ContractCounter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DealId)
	if l > 0 {
		n += 1 + l + sovContractCounter(uint64(l))
	}
	if m.IdValue != 0 {
		n += 1 + sovContractCounter(uint64(m.IdValue))
	}
	return n
}

func sovContractCounter(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozContractCounter(x uint64) (n int) {
	return sovContractCounter(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ContractCounter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowContractCounter
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
			return fmt.Errorf("proto: ContractCounter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContractCounter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DealId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContractCounter
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
				return ErrInvalidLengthContractCounter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthContractCounter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DealId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IdValue", wireType)
			}
			m.IdValue = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContractCounter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.IdValue |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipContractCounter(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthContractCounter
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
func skipContractCounter(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowContractCounter
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
					return 0, ErrIntOverflowContractCounter
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
					return 0, ErrIntOverflowContractCounter
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
				return 0, ErrInvalidLengthContractCounter
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupContractCounter
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthContractCounter
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthContractCounter        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowContractCounter          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupContractCounter = fmt.Errorf("proto: unexpected end of group")
)
