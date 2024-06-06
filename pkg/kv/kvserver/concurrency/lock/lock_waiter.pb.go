// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kv/kvserver/concurrency/lock/lock_waiter.proto

package lock

import (
	fmt "fmt"
	enginepb "github.com/cockroachdb/cockroach/pkg/storage/enginepb"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Waiter represents a transaction (or non-transactional operation) that is
// waiting in the wait queue of readers or writers on an individual lock.
type Waiter struct {
	// The transaction associated with this waiter, or nil in the case of a
	// non-transactional waiter.
	WaitingTxn *enginepb.TxnMeta `protobuf:"bytes,1,opt,name=waiting_txn,json=waitingTxn,proto3" json:"waiting_txn,omitempty"`
	// Represents if this operation is actively waiting on the lock.  While all
	// readers are active waiters, there are some cases in which writers may not
	// be actively waiting, for instance in the case of a broken reservation.
	ActiveWaiter bool `protobuf:"varint,2,opt,name=active_waiter,json=activeWaiter,proto3" json:"active_waiter,omitempty"`
	// The strength at which this waiter is attempting to acquire the lock.
	Strength Strength `protobuf:"varint,3,opt,name=strength,proto3,enum=cockroach.kv.kvserver.concurrency.lock.Strength" json:"strength,omitempty"`
	// The wall clock duration since this operation began waiting on the lock.
	WaitDuration time.Duration `protobuf:"bytes,4,opt,name=wait_duration,json=waitDuration,proto3,stdduration" json:"wait_duration"`
}

func (m *Waiter) Reset()         { *m = Waiter{} }
func (m *Waiter) String() string { return proto.CompactTextString(m) }
func (*Waiter) ProtoMessage()    {}
func (*Waiter) Descriptor() ([]byte, []int) {
	return fileDescriptor_f15b14e23af219f5, []int{0}
}
func (m *Waiter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Waiter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *Waiter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Waiter.Merge(m, src)
}
func (m *Waiter) XXX_Size() int {
	return m.Size()
}
func (m *Waiter) XXX_DiscardUnknown() {
	xxx_messageInfo_Waiter.DiscardUnknown(m)
}

var xxx_messageInfo_Waiter proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Waiter)(nil), "cockroach.kv.kvserver.concurrency.lock.Waiter")
}

func init() {
	proto.RegisterFile("kv/kvserver/concurrency/lock/lock_waiter.proto", fileDescriptor_f15b14e23af219f5)
}

var fileDescriptor_f15b14e23af219f5 = []byte{
	// 364 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0xbd, 0x6e, 0xea, 0x30,
	0x14, 0x8e, 0xb9, 0x57, 0x08, 0x05, 0xb8, 0x43, 0x74, 0x87, 0x5c, 0x74, 0x65, 0x50, 0x91, 0x2a,
	0xd4, 0xe1, 0xb8, 0x82, 0x27, 0x68, 0xc5, 0xd0, 0xa1, 0x5d, 0x52, 0xa4, 0x4a, 0x5d, 0x50, 0x62,
	0x5c, 0x13, 0x05, 0x6c, 0x64, 0x9c, 0x94, 0x3e, 0x43, 0x97, 0x8e, 0x7d, 0x24, 0x46, 0x46, 0xa6,
	0xfe, 0x84, 0x17, 0xa9, 0x92, 0x38, 0xd0, 0xa9, 0xed, 0x62, 0x1d, 0x9f, 0xe3, 0xef, 0xef, 0xd8,
	0x86, 0x28, 0x21, 0x51, 0xb2, 0x64, 0x2a, 0x61, 0x8a, 0x50, 0x29, 0x68, 0xac, 0x14, 0x13, 0xf4,
	0x81, 0xcc, 0x24, 0x8d, 0xf2, 0x63, 0x7c, 0xef, 0x87, 0x9a, 0x29, 0x58, 0x28, 0xa9, 0xa5, 0x73,
	0x4c, 0x25, 0x8d, 0x94, 0xf4, 0xe9, 0x14, 0xa2, 0x04, 0x4a, 0x24, 0x7c, 0x42, 0x42, 0x06, 0x6a,
	0x9d, 0x7c, 0xcb, 0x1b, 0x0a, 0x5e, 0x70, 0xb6, 0xfe, 0x2f, 0xb5, 0x54, 0x3e, 0x67, 0x84, 0x09,
	0x1e, 0x0a, 0xb6, 0x08, 0xc8, 0x3c, 0xa1, 0x74, 0x60, 0xa6, 0x7f, 0xb9, 0xe4, 0x32, 0x2f, 0x49,
	0x56, 0x99, 0x2e, 0xe6, 0x52, 0xf2, 0x19, 0x23, 0xf9, 0x2d, 0x88, 0xef, 0xc8, 0x24, 0x56, 0xbe,
	0x0e, 0xa5, 0x28, 0xe6, 0x47, 0x8f, 0x15, 0xbb, 0x7a, 0x93, 0x1b, 0x77, 0x86, 0x76, 0x3d, 0x8b,
	0x10, 0x0a, 0x3e, 0xd6, 0x2b, 0xe1, 0xa2, 0x0e, 0xea, 0xd5, 0xfb, 0x5d, 0x38, 0x04, 0x31, 0xf2,
	0x50, 0xca, 0xc3, 0x68, 0x25, 0xae, 0x98, 0xf6, 0x3d, 0xdb, 0xe0, 0x46, 0x2b, 0xe1, 0x74, 0xed,
	0xa6, 0x4f, 0x75, 0x98, 0x30, 0xb3, 0x0f, 0xb7, 0xd2, 0x41, 0xbd, 0x9a, 0xd7, 0x28, 0x9a, 0x46,
	0xea, 0xd2, 0xae, 0x2d, 0xb5, 0x62, 0x82, 0xeb, 0xa9, 0xfb, 0xab, 0x83, 0x7a, 0x7f, 0xfa, 0xa7,
	0xf0, 0xb3, 0x85, 0xc1, 0xb5, 0xc1, 0x79, 0x7b, 0x06, 0xe7, 0xc2, 0x6e, 0x66, 0x5a, 0xe3, 0x32,
	0x9a, 0xfb, 0x3b, 0xb7, 0xfe, 0x0f, 0x8a, 0xec, 0x50, 0x66, 0x87, 0xa1, 0x79, 0x70, 0x5e, 0x5b,
	0xbf, 0xb4, 0xad, 0xe7, 0xd7, 0x36, 0xf2, 0x1a, 0x19, 0x72, 0xdf, 0xa7, 0xeb, 0x77, 0x6c, 0xad,
	0x53, 0x8c, 0x36, 0x29, 0x46, 0xdb, 0x14, 0xa3, 0xb7, 0x14, 0xa3, 0xa7, 0x1d, 0xb6, 0x36, 0x3b,
	0x6c, 0x6d, 0x77, 0xd8, 0xba, 0x3d, 0xe3, 0xa1, 0x9e, 0xc6, 0x01, 0x50, 0x39, 0x27, 0x7b, 0xc7,
	0x93, 0xe0, 0x50, 0x93, 0x45, 0xc4, 0xc9, 0x57, 0x9f, 0x1a, 0x54, 0x73, 0x3f, 0x83, 0x8f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xa4, 0xf9, 0x94, 0x3b, 0x53, 0x02, 0x00, 0x00,
}

func (m *Waiter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Waiter) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Waiter) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.WaitDuration, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.WaitDuration):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintLockWaiter(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x22
	if m.Strength != 0 {
		i = encodeVarintLockWaiter(dAtA, i, uint64(m.Strength))
		i--
		dAtA[i] = 0x18
	}
	if m.ActiveWaiter {
		i--
		if m.ActiveWaiter {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.WaitingTxn != nil {
		{
			size, err := m.WaitingTxn.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintLockWaiter(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintLockWaiter(dAtA []byte, offset int, v uint64) int {
	offset -= sovLockWaiter(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Waiter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.WaitingTxn != nil {
		l = m.WaitingTxn.Size()
		n += 1 + l + sovLockWaiter(uint64(l))
	}
	if m.ActiveWaiter {
		n += 2
	}
	if m.Strength != 0 {
		n += 1 + sovLockWaiter(uint64(m.Strength))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.WaitDuration)
	n += 1 + l + sovLockWaiter(uint64(l))
	return n
}

func sovLockWaiter(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLockWaiter(x uint64) (n int) {
	return sovLockWaiter(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Waiter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLockWaiter
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
			return fmt.Errorf("proto: Waiter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Waiter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WaitingTxn", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLockWaiter
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
				return ErrInvalidLengthLockWaiter
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLockWaiter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.WaitingTxn == nil {
				m.WaitingTxn = &enginepb.TxnMeta{}
			}
			if err := m.WaitingTxn.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActiveWaiter", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLockWaiter
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
			m.ActiveWaiter = bool(v != 0)
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Strength", wireType)
			}
			m.Strength = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLockWaiter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Strength |= Strength(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WaitDuration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLockWaiter
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
				return ErrInvalidLengthLockWaiter
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLockWaiter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.WaitDuration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLockWaiter(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLockWaiter
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
func skipLockWaiter(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLockWaiter
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
					return 0, ErrIntOverflowLockWaiter
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
					return 0, ErrIntOverflowLockWaiter
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
				return 0, ErrInvalidLengthLockWaiter
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLockWaiter
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLockWaiter
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLockWaiter        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLockWaiter          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLockWaiter = fmt.Errorf("proto: unexpected end of group")
)

