// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmosacademyexercisesimple/world_params_2.proto

package types

import (
	fmt "fmt"
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

type WorldParams2 struct {
	Id         uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Gravity    uint64 `protobuf:"varint,3,opt,name=gravity,proto3" json:"gravity,omitempty"`
	LandSupply uint64 `protobuf:"varint,4,opt,name=landSupply,proto3" json:"landSupply,omitempty"`
}

func (m *WorldParams2) Reset()         { *m = WorldParams2{} }
func (m *WorldParams2) String() string { return proto.CompactTextString(m) }
func (*WorldParams2) ProtoMessage()    {}
func (*WorldParams2) Descriptor() ([]byte, []int) {
	return fileDescriptor_f446316c9e61bc3a, []int{0}
}
func (m *WorldParams2) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WorldParams2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WorldParams2.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WorldParams2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorldParams2.Merge(m, src)
}
func (m *WorldParams2) XXX_Size() int {
	return m.Size()
}
func (m *WorldParams2) XXX_DiscardUnknown() {
	xxx_messageInfo_WorldParams2.DiscardUnknown(m)
}

var xxx_messageInfo_WorldParams2 proto.InternalMessageInfo

func (m *WorldParams2) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *WorldParams2) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *WorldParams2) GetGravity() uint64 {
	if m != nil {
		return m.Gravity
	}
	return 0
}

func (m *WorldParams2) GetLandSupply() uint64 {
	if m != nil {
		return m.LandSupply
	}
	return 0
}

func init() {
	proto.RegisterType((*WorldParams2)(nil), "b9lab.cosmosacademyexercisesimple.cosmosacademyexercisesimple.WorldParams2")
}

func init() {
	proto.RegisterFile("cosmosacademyexercisesimple/world_params_2.proto", fileDescriptor_f446316c9e61bc3a)
}

var fileDescriptor_f446316c9e61bc3a = []byte{
	// 249 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x48, 0xce, 0x2f, 0xce,
	0xcd, 0x2f, 0x4e, 0x4c, 0x4e, 0x4c, 0x49, 0xcd, 0xad, 0x4c, 0xad, 0x48, 0x2d, 0x4a, 0xce, 0x2c,
	0x4e, 0x2d, 0xce, 0xcc, 0x2d, 0xc8, 0x49, 0xd5, 0x2f, 0xcf, 0x2f, 0xca, 0x49, 0x89, 0x2f, 0x48,
	0x2c, 0x4a, 0xcc, 0x2d, 0x8e, 0x37, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xb2, 0x4d, 0xb2,
	0xcc, 0x49, 0x4c, 0xd2, 0xc3, 0xa3, 0x0f, 0x9f, 0x9c, 0x94, 0x48, 0x7a, 0x7e, 0x7a, 0x3e, 0xd8,
	0x24, 0x7d, 0x10, 0x0b, 0x62, 0xa8, 0x52, 0x0e, 0x17, 0x4f, 0x38, 0xc8, 0xb2, 0x00, 0xb0, 0x5d,
	0x46, 0x42, 0x7c, 0x5c, 0x4c, 0x99, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x2c, 0x41, 0x4c, 0x99,
	0x29, 0x42, 0x42, 0x5c, 0x2c, 0x79, 0x89, 0xb9, 0xa9, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41,
	0x60, 0xb6, 0x90, 0x04, 0x17, 0x7b, 0x7a, 0x51, 0x62, 0x59, 0x66, 0x49, 0xa5, 0x04, 0x33, 0x58,
	0x21, 0x8c, 0x2b, 0x24, 0xc7, 0xc5, 0x95, 0x93, 0x98, 0x97, 0x12, 0x5c, 0x5a, 0x50, 0x90, 0x53,
	0x29, 0xc1, 0x02, 0x96, 0x44, 0x12, 0x71, 0xca, 0x3d, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39,
	0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63,
	0x39, 0x86, 0xa8, 0xe0, 0xf4, 0xcc, 0x92, 0x8c, 0x52, 0x90, 0xef, 0x72, 0xf5, 0xc1, 0xfe, 0xd4,
	0x87, 0xf8, 0x45, 0x17, 0xea, 0x19, 0x5d, 0x98, 0x6f, 0x74, 0xa1, 0x41, 0x54, 0xa1, 0x8f, 0x2f,
	0x00, 0x4b, 0x2a, 0x0b, 0x52, 0x8b, 0x93, 0xd8, 0xc0, 0x7e, 0x34, 0x06, 0x04, 0x00, 0x00, 0xff,
	0xff, 0x2b, 0x55, 0xe5, 0x80, 0x6c, 0x01, 0x00, 0x00,
}

func (m *WorldParams2) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WorldParams2) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WorldParams2) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LandSupply != 0 {
		i = encodeVarintWorldParams_2(dAtA, i, uint64(m.LandSupply))
		i--
		dAtA[i] = 0x20
	}
	if m.Gravity != 0 {
		i = encodeVarintWorldParams_2(dAtA, i, uint64(m.Gravity))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintWorldParams_2(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintWorldParams_2(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintWorldParams_2(dAtA []byte, offset int, v uint64) int {
	offset -= sovWorldParams_2(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *WorldParams2) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovWorldParams_2(uint64(m.Id))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovWorldParams_2(uint64(l))
	}
	if m.Gravity != 0 {
		n += 1 + sovWorldParams_2(uint64(m.Gravity))
	}
	if m.LandSupply != 0 {
		n += 1 + sovWorldParams_2(uint64(m.LandSupply))
	}
	return n
}

func sovWorldParams_2(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozWorldParams_2(x uint64) (n int) {
	return sovWorldParams_2(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *WorldParams2) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWorldParams_2
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
			return fmt.Errorf("proto: WorldParams2: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WorldParams2: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorldParams_2
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorldParams_2
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
				return ErrInvalidLengthWorldParams_2
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWorldParams_2
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Gravity", wireType)
			}
			m.Gravity = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorldParams_2
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Gravity |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LandSupply", wireType)
			}
			m.LandSupply = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorldParams_2
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LandSupply |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipWorldParams_2(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthWorldParams_2
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
func skipWorldParams_2(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowWorldParams_2
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
					return 0, ErrIntOverflowWorldParams_2
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
					return 0, ErrIntOverflowWorldParams_2
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
				return 0, ErrInvalidLengthWorldParams_2
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupWorldParams_2
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthWorldParams_2
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthWorldParams_2        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowWorldParams_2          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupWorldParams_2 = fmt.Errorf("proto: unexpected end of group")
)
