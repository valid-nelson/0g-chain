// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kava/committee/v1beta1/proposal.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// CommitteeChangeProposal is a gov proposal for creating a new committee or modifying an existing one.
type CommitteeChangeProposal struct {
	Title        string     `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description  string     `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	NewCommittee *types.Any `protobuf:"bytes,3,opt,name=new_committee,json=newCommittee,proto3" json:"new_committee,omitempty"`
}

func (m *CommitteeChangeProposal) Reset()         { *m = CommitteeChangeProposal{} }
func (m *CommitteeChangeProposal) String() string { return proto.CompactTextString(m) }
func (*CommitteeChangeProposal) ProtoMessage()    {}
func (*CommitteeChangeProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_4886de4a6c720e57, []int{0}
}
func (m *CommitteeChangeProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CommitteeChangeProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CommitteeChangeProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CommitteeChangeProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommitteeChangeProposal.Merge(m, src)
}
func (m *CommitteeChangeProposal) XXX_Size() int {
	return m.Size()
}
func (m *CommitteeChangeProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_CommitteeChangeProposal.DiscardUnknown(m)
}

var xxx_messageInfo_CommitteeChangeProposal proto.InternalMessageInfo

// CommitteeDeleteProposal is a gov proposal for removing a committee.
type CommitteeDeleteProposal struct {
	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	CommitteeID uint64 `protobuf:"varint,3,opt,name=committee_id,json=committeeId,proto3" json:"committee_id,omitempty"`
}

func (m *CommitteeDeleteProposal) Reset()         { *m = CommitteeDeleteProposal{} }
func (m *CommitteeDeleteProposal) String() string { return proto.CompactTextString(m) }
func (*CommitteeDeleteProposal) ProtoMessage()    {}
func (*CommitteeDeleteProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_4886de4a6c720e57, []int{1}
}
func (m *CommitteeDeleteProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CommitteeDeleteProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CommitteeDeleteProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CommitteeDeleteProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommitteeDeleteProposal.Merge(m, src)
}
func (m *CommitteeDeleteProposal) XXX_Size() int {
	return m.Size()
}
func (m *CommitteeDeleteProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_CommitteeDeleteProposal.DiscardUnknown(m)
}

var xxx_messageInfo_CommitteeDeleteProposal proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CommitteeChangeProposal)(nil), "kava.committee.v1beta1.CommitteeChangeProposal")
	proto.RegisterType((*CommitteeDeleteProposal)(nil), "kava.committee.v1beta1.CommitteeDeleteProposal")
}

func init() {
	proto.RegisterFile("kava/committee/v1beta1/proposal.proto", fileDescriptor_4886de4a6c720e57)
}

var fileDescriptor_4886de4a6c720e57 = []byte{
	// 353 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xbf, 0x6e, 0xf2, 0x30,
	0x14, 0xc5, 0xe3, 0xef, 0x9f, 0x44, 0x02, 0xfa, 0xa4, 0x08, 0xb5, 0xc0, 0xe0, 0x22, 0xa4, 0x4a,
	0x0c, 0xc5, 0x06, 0xba, 0x75, 0x2b, 0x30, 0x94, 0x4e, 0x15, 0x63, 0x17, 0xe4, 0x04, 0xd7, 0x58,
	0x0d, 0xbe, 0x11, 0x31, 0x50, 0xde, 0xa2, 0x2f, 0xd1, 0x37, 0x60, 0xeb, 0x0b, 0x20, 0x26, 0xc6,
	0x4e, 0x55, 0x1b, 0x5e, 0xa4, 0x22, 0x09, 0x16, 0x5b, 0x87, 0x6e, 0x3e, 0xe7, 0x1e, 0xeb, 0xfe,
	0x7c, 0x7d, 0xed, 0xf3, 0x47, 0x36, 0x67, 0xd4, 0x87, 0xc9, 0x44, 0x6a, 0xcd, 0x39, 0x9d, 0xb7,
	0x3c, 0xae, 0x59, 0x8b, 0x86, 0x53, 0x08, 0x21, 0x62, 0x01, 0x09, 0xa7, 0xa0, 0xc1, 0x3d, 0xd9,
	0xc7, 0x88, 0x89, 0x91, 0x2c, 0x56, 0x29, 0xfb, 0x10, 0x4d, 0x20, 0x1a, 0x26, 0x29, 0x9a, 0x8a,
	0xf4, 0x4a, 0xa5, 0x28, 0x40, 0x40, 0xea, 0xef, 0x4f, 0x99, 0x5b, 0x16, 0x00, 0x22, 0xe0, 0x34,
	0x51, 0xde, 0xec, 0x81, 0x32, 0xb5, 0x4c, 0x4b, 0xb5, 0x57, 0x64, 0x9f, 0x76, 0x0f, 0x1d, 0xba,
	0x63, 0xa6, 0x04, 0xbf, 0xcb, 0x28, 0xdc, 0xa2, 0xfd, 0x57, 0x4b, 0x1d, 0xf0, 0x12, 0xaa, 0xa2,
	0x7a, 0x6e, 0x90, 0x0a, 0xb7, 0x6a, 0x3b, 0x23, 0x1e, 0xf9, 0x53, 0x19, 0x6a, 0x09, 0xaa, 0xf4,
	0x2b, 0xa9, 0x1d, 0x5b, 0xee, 0x8d, 0x5d, 0x50, 0x7c, 0x31, 0x34, 0xe0, 0xa5, 0xdf, 0x55, 0x54,
	0x77, 0xda, 0x45, 0x92, 0x62, 0x90, 0x03, 0x06, 0xb9, 0x56, 0xcb, 0x4e, 0x61, 0xb3, 0x6a, 0xe4,
	0x0c, 0xc1, 0x20, 0xaf, 0xf8, 0xc2, 0xa8, 0x2b, 0xbc, 0x59, 0x35, 0x2a, 0xd9, 0x03, 0x05, 0xcc,
	0x0f, 0x13, 0x20, 0x5d, 0x50, 0x9a, 0x2b, 0x5d, 0x7b, 0x39, 0xa6, 0xef, 0xf1, 0x80, 0xeb, 0x9f,
	0xd3, 0xb7, 0xed, 0xbc, 0x21, 0x1f, 0xca, 0x51, 0x02, 0xff, 0xa7, 0xf3, 0x3f, 0x7e, 0x3f, 0x73,
	0x4c, 0xab, 0x7e, 0x6f, 0xe0, 0x98, 0x50, 0x7f, 0xf4, 0x1d, 0x67, 0xe7, 0x76, 0xfd, 0x89, 0xad,
	0x75, 0x8c, 0xd1, 0x36, 0xc6, 0xe8, 0x23, 0xc6, 0xe8, 0x79, 0x87, 0xad, 0xed, 0x0e, 0x5b, 0x6f,
	0x3b, 0x6c, 0xdd, 0x5f, 0x08, 0xa9, 0xc7, 0x33, 0x6f, 0xff, 0xd3, 0xb4, 0x29, 0x02, 0xe6, 0x45,
	0xb4, 0x29, 0x1a, 0xfe, 0x98, 0x49, 0x45, 0x9f, 0x8e, 0xd6, 0x44, 0x2f, 0x43, 0x1e, 0x79, 0xff,
	0x92, 0xf1, 0x5d, 0x7e, 0x05, 0x00, 0x00, 0xff, 0xff, 0xf9, 0x4a, 0x8b, 0xf9, 0x45, 0x02, 0x00,
	0x00,
}

func (m *CommitteeChangeProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CommitteeChangeProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CommitteeChangeProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.NewCommittee != nil {
		{
			size, err := m.NewCommittee.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintProposal(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CommitteeDeleteProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CommitteeDeleteProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CommitteeDeleteProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CommitteeID != 0 {
		i = encodeVarintProposal(dAtA, i, uint64(m.CommitteeID))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintProposal(dAtA []byte, offset int, v uint64) int {
	offset -= sovProposal(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CommitteeChangeProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	if m.NewCommittee != nil {
		l = m.NewCommittee.Size()
		n += 1 + l + sovProposal(uint64(l))
	}
	return n
}

func (m *CommitteeDeleteProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	if m.CommitteeID != 0 {
		n += 1 + sovProposal(uint64(m.CommitteeID))
	}
	return n
}

func sovProposal(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProposal(x uint64) (n int) {
	return sovProposal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CommitteeChangeProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposal
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
			return fmt.Errorf("proto: CommitteeChangeProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CommitteeChangeProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewCommittee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.NewCommittee == nil {
				m.NewCommittee = &types.Any{}
			}
			if err := m.NewCommittee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposal
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
func (m *CommitteeDeleteProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposal
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
			return fmt.Errorf("proto: CommitteeDeleteProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CommitteeDeleteProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommitteeID", wireType)
			}
			m.CommitteeID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CommitteeID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposal
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
func skipProposal(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProposal
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
					return 0, ErrIntOverflowProposal
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
					return 0, ErrIntOverflowProposal
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
				return 0, ErrInvalidLengthProposal
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProposal
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProposal
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProposal        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProposal          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProposal = fmt.Errorf("proto: unexpected end of group")
)
