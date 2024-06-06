// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kv/kvserver/kvserverpb/lease_status.proto

package kvserverpb

import (
	fmt "fmt"
	livenesspb "github.com/cockroachdb/cockroach/pkg/kv/kvserver/liveness/livenesspb"
	roachpb "github.com/cockroachdb/cockroach/pkg/roachpb"
	github_com_cockroachdb_cockroach_pkg_util_hlc "github.com/cockroachdb/cockroach/pkg/util/hlc"
	hlc "github.com/cockroachdb/cockroach/pkg/util/hlc"
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

type LeaseState int32

const (
	// ERROR indicates that the lease can't be used or acquired.
	LeaseState_ERROR LeaseState = 0
	// VALID indicates that the lease is not expired at the current clock
	// time and can be used to serve a given request.
	LeaseState_VALID LeaseState = 1
	// UNUSABLE indicates that a lease has not expired at the current clock
	// time, but cannot be used to serve a given request. A lease may be
	// unusable for one of two reasons.
	//
	// First, if the request operates at a timestamp in the future, it is
	// possible for the request's timestamp to fall outside of the lease's
	// validity window, even if the lease is not yet expired at the current
	// clock time. In such cases, the lease must be extended past the
	// request's timestamp before the request can be served under the lease.
	//
	// Second, even if the request does not operate at a timestamp in the
	// future and operates fully within the lease's validity window, it may
	// operate at a time too close to the lease's expiration to be served
	// safely due to clock uncertainty. We refer to the period at the end of
	// each lease, immediately before its expiration, as its stasis period.
	//
	// The point of the stasis period is to prevent reads on the old
	// leaseholder (the one whose stasis we're talking about) from missing
	// to see writes performed under the next lease (held by someone else)
	// when these writes should fall in the uncertainty window. Even without
	// the stasis, writes performed by the new leaseholder are guaranteed to
	// have higher timestamps than any reads served by the old leaseholder.
	// However, a read at timestamp T needs to observe all writes at
	// timestamps [T, T+maxOffset] and so, without the stasis, only the new
	// leaseholder might have some of these writes. In other words, without
	// the stasis, a new leaseholder with a fast clock could start
	// performing writes ordered in real time before the old leaseholder
	// considers its lease to have expired.
	//
	// An UNUSABLE lease may become VALID for the same leaseholder after a
	// successful RequestLease (for expiration-based leases) or Heartbeat
	// (for epoch-based leases), each of which serve as forms of "lease
	// extension".
	LeaseState_UNUSABLE LeaseState = 2
	// EXPIRED indicates that the current clock time is past the lease's
	// expiration time. An expired lease may become VALID for the same
	// leaseholder on RequestLease or Heartbeat, or it may be replaced by a
	// new leaseholder with a RequestLease (for expiration-based leases) or
	// IncrementEpoch+RequestLease (for epoch-based leases).
	//
	// Only an EXPIRED lease may change hands non-cooperatively.
	LeaseState_EXPIRED LeaseState = 3
	// PROSCRIBED indicates that the lease's proposed timestamp is earlier
	// than allowed and can't be used to serve a request. This is used to
	// detect node restarts: a node that has restarted will see its former
	// incarnation's leases as PROSCRIBED so it will renew them before using
	// them. This state also used during a lease transfer, to prevent the
	// outgoing leaseholder from serving any other requests under its old
	// lease. Note that the PROSCRIBED state is only visible to the
	// leaseholder; other nodes may see this as a VALID lease.
	LeaseState_PROSCRIBED LeaseState = 4
)

var LeaseState_name = map[int32]string{
	0: "ERROR",
	1: "VALID",
	2: "UNUSABLE",
	3: "EXPIRED",
	4: "PROSCRIBED",
}

var LeaseState_value = map[string]int32{
	"ERROR":      0,
	"VALID":      1,
	"UNUSABLE":   2,
	"EXPIRED":    3,
	"PROSCRIBED": 4,
}

func (x LeaseState) String() string {
	return proto.EnumName(LeaseState_name, int32(x))
}

func (LeaseState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fb959f350412bc90, []int{0}
}

// LeaseStatus holds the lease state, the current clock time at which the
// state is accurate, the request time at which the status is accurate, the
// lease iself, and optionally the liveness if the lease is epoch-based.
type LeaseStatus struct {
	// Lease which this status describes.
	Lease roachpb.Lease `protobuf:"bytes,1,opt,name=lease,proto3" json:"lease"`
	// Clock timestamp that the lease was evaluated at.
	Now github_com_cockroachdb_cockroach_pkg_util_hlc.ClockTimestamp `protobuf:"bytes,2,opt,name=now,proto3,casttype=github.com/cockroachdb/cockroach/pkg/util/hlc.ClockTimestamp" json:"now"`
	// Timestamp for the request operating under the lease.
	RequestTime hlc.Timestamp `protobuf:"bytes,5,opt,name=request_time,json=requestTime,proto3" json:"request_time"`
	// State of the lease at now for a request at request_time.
	State LeaseState `protobuf:"varint,3,opt,name=state,proto3,enum=cockroach.kv.kvserver.storagepb.LeaseState" json:"state,omitempty"`
	// If state == ERROR, this provides more info about the error.
	ErrInfo string `protobuf:"bytes,6,opt,name=err_info,json=errInfo,proto3" json:"err_info,omitempty"`
	// Liveness if this is an epoch-based lease.
	Liveness livenesspb.Liveness `protobuf:"bytes,4,opt,name=liveness,proto3" json:"liveness"`
	// The minimum observed timestamp on a transaction that is respected by this lease.
	MinValidObservedTimestamp github_com_cockroachdb_cockroach_pkg_util_hlc.ClockTimestamp `protobuf:"bytes,7,opt,name=min_valid_observed_timestamp,json=minValidObservedTimestamp,proto3,casttype=github.com/cockroachdb/cockroach/pkg/util/hlc.ClockTimestamp" json:"min_valid_observed_timestamp"`
}

func (m *LeaseStatus) Reset()         { *m = LeaseStatus{} }
func (m *LeaseStatus) String() string { return proto.CompactTextString(m) }
func (*LeaseStatus) ProtoMessage()    {}
func (*LeaseStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb959f350412bc90, []int{0}
}
func (m *LeaseStatus) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LeaseStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *LeaseStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeaseStatus.Merge(m, src)
}
func (m *LeaseStatus) XXX_Size() int {
	return m.Size()
}
func (m *LeaseStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_LeaseStatus.DiscardUnknown(m)
}

var xxx_messageInfo_LeaseStatus proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("cockroach.kv.kvserver.storagepb.LeaseState", LeaseState_name, LeaseState_value)
	proto.RegisterType((*LeaseStatus)(nil), "cockroach.kv.kvserver.storagepb.LeaseStatus")
}

func init() {
	proto.RegisterFile("kv/kvserver/kvserverpb/lease_status.proto", fileDescriptor_fb959f350412bc90)
}

var fileDescriptor_fb959f350412bc90 = []byte{
	// 506 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x93, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0xe3, 0xb5, 0x5d, 0x3b, 0x77, 0x9a, 0x2a, 0x8b, 0x43, 0x36, 0x41, 0x5a, 0x71, 0x2a,
	0x20, 0x39, 0xd2, 0xc6, 0x0d, 0x2e, 0xed, 0x1a, 0xa4, 0x4a, 0x85, 0x4e, 0x2e, 0x9d, 0x10, 0x12,
	0x8a, 0x92, 0xd4, 0x4b, 0xa3, 0xfc, 0x71, 0x70, 0x9c, 0xf0, 0x35, 0xf8, 0x10, 0x7c, 0x98, 0x1e,
	0x77, 0xdc, 0xa9, 0x82, 0xf6, 0x5b, 0x70, 0x42, 0x4e, 0xd2, 0x74, 0x07, 0x10, 0x3b, 0xed, 0xf6,
	0xc8, 0x7e, 0x9f, 0xf7, 0xfd, 0xbd, 0x7e, 0x64, 0xf8, 0xc2, 0xcf, 0x74, 0x3f, 0x4b, 0x28, 0xcf,
	0x28, 0xaf, 0x44, 0x6c, 0xeb, 0x01, 0xb5, 0x12, 0x6a, 0x26, 0xc2, 0x12, 0x69, 0x82, 0x63, 0xce,
	0x04, 0x43, 0x5d, 0x87, 0x39, 0x3e, 0x67, 0x96, 0xb3, 0xc4, 0x7e, 0x86, 0x77, 0xb5, 0x38, 0x11,
	0x8c, 0x5b, 0x2e, 0x8d, 0xed, 0x33, 0x94, 0x5f, 0xc6, 0xb6, 0xbe, 0xb0, 0x84, 0x55, 0x98, 0xce,
	0xf0, 0xfd, 0xfe, 0x81, 0x97, 0xd1, 0x88, 0x26, 0x49, 0x25, 0xe4, 0xa0, 0x52, 0x96, 0xf5, 0x6a,
	0x2a, 0xbc, 0x40, 0x5f, 0x06, 0x8e, 0x2e, 0xbc, 0x90, 0x26, 0xc2, 0x0a, 0xe3, 0xf2, 0xe6, 0x89,
	0xcb, 0x5c, 0x96, 0x4b, 0x5d, 0xaa, 0xe2, 0xf4, 0xf9, 0xba, 0x0e, 0xdb, 0x13, 0xc9, 0x3a, 0xcb,
	0x51, 0xd1, 0x6b, 0xd8, 0xc8, 0xd1, 0x55, 0xd0, 0x03, 0xfd, 0xf6, 0xb9, 0x8a, 0xf7, 0xd0, 0x25,
	0x1d, 0xce, 0xcb, 0x87, 0xf5, 0xd5, 0xba, 0xab, 0x90, 0xa2, 0x18, 0xa5, 0xb0, 0x16, 0xb1, 0x6f,
	0xea, 0x41, 0xee, 0x79, 0x76, 0xcf, 0x23, 0x69, 0xf0, 0x32, 0x70, 0xf0, 0xc7, 0x1d, 0xcd, 0x70,
	0x24, 0x8d, 0xbf, 0xd7, 0xdd, 0xb7, 0xae, 0x27, 0x96, 0xa9, 0x8d, 0x1d, 0x16, 0xea, 0x95, 0x61,
	0x61, 0xef, 0xb5, 0x1e, 0xfb, 0xae, 0xbe, 0x5b, 0x07, 0x5f, 0x06, 0xcc, 0xf1, 0xab, 0x2e, 0x44,
	0xce, 0x43, 0xef, 0xe0, 0x31, 0xa7, 0x5f, 0x53, 0x9a, 0x08, 0x53, 0x6e, 0xab, 0x36, 0x1e, 0x32,
	0xbf, 0x00, 0x6f, 0x97, 0x46, 0x79, 0x8e, 0x06, 0xb0, 0x21, 0x93, 0xa2, 0x6a, 0xad, 0x07, 0xfa,
	0x27, 0xe7, 0xaf, 0xf0, 0x7f, 0x92, 0xc2, 0xd5, 0x8b, 0x51, 0x52, 0x38, 0xd1, 0x29, 0x6c, 0x51,
	0xce, 0x4d, 0x2f, 0xba, 0x61, 0xea, 0x61, 0x0f, 0xf4, 0x8f, 0x48, 0x93, 0x72, 0x3e, 0x8e, 0x6e,
	0x18, 0x9a, 0xc3, 0xd6, 0x2e, 0x24, 0xb5, 0x9e, 0x13, 0x5e, 0xfc, 0x63, 0x40, 0x95, 0xe5, 0x3e,
	0x5f, 0x3c, 0x29, 0x65, 0xc9, 0x5d, 0xb5, 0x42, 0x3f, 0x00, 0x7c, 0x1a, 0x7a, 0x91, 0x99, 0x59,
	0x81, 0xb7, 0x30, 0x99, 0x9d, 0xf7, 0x58, 0x98, 0x55, 0xec, 0x6a, 0xf3, 0xf1, 0xd2, 0x38, 0x0d,
	0xbd, 0xe8, 0x5a, 0x72, 0x4c, 0x4b, 0x8c, 0xea, 0xea, 0xe5, 0x7b, 0x08, 0xf7, 0xaf, 0x85, 0x8e,
	0x60, 0xc3, 0x20, 0x64, 0x4a, 0x3a, 0x8a, 0x94, 0xd7, 0x83, 0xc9, 0x78, 0xd4, 0x01, 0xe8, 0x18,
	0xb6, 0xe6, 0x1f, 0xe6, 0xb3, 0xc1, 0x70, 0x62, 0x74, 0x0e, 0x50, 0x1b, 0x36, 0x8d, 0x4f, 0x57,
	0x63, 0x62, 0x8c, 0x3a, 0x35, 0x74, 0x02, 0xe1, 0x15, 0x99, 0xce, 0x2e, 0xc9, 0x78, 0x68, 0x8c,
	0x3a, 0xf5, 0xe1, 0x97, 0xd5, 0x2f, 0x4d, 0x59, 0x6d, 0x34, 0x70, 0xbb, 0xd1, 0xc0, 0xdd, 0x46,
	0x03, 0x3f, 0x37, 0x1a, 0xf8, 0xbe, 0xd5, 0x94, 0xdb, 0xad, 0xa6, 0xdc, 0x6d, 0x35, 0xe5, 0xf3,
	0x9b, 0x07, 0x6d, 0xf2, 0xf7, 0x6f, 0x6b, 0x1f, 0xe6, 0xbf, 0xe2, 0xe2, 0x4f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x47, 0xab, 0x53, 0x94, 0xd7, 0x03, 0x00, 0x00,
}

func (m *LeaseStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LeaseStatus) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LeaseStatus) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.MinValidObservedTimestamp.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintLeaseStatus(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	if len(m.ErrInfo) > 0 {
		i -= len(m.ErrInfo)
		copy(dAtA[i:], m.ErrInfo)
		i = encodeVarintLeaseStatus(dAtA, i, uint64(len(m.ErrInfo)))
		i--
		dAtA[i] = 0x32
	}
	{
		size, err := m.RequestTime.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintLeaseStatus(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size, err := m.Liveness.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintLeaseStatus(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.State != 0 {
		i = encodeVarintLeaseStatus(dAtA, i, uint64(m.State))
		i--
		dAtA[i] = 0x18
	}
	{
		size, err := m.Now.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintLeaseStatus(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.Lease.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintLeaseStatus(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintLeaseStatus(dAtA []byte, offset int, v uint64) int {
	offset -= sovLeaseStatus(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LeaseStatus) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Lease.Size()
	n += 1 + l + sovLeaseStatus(uint64(l))
	l = m.Now.Size()
	n += 1 + l + sovLeaseStatus(uint64(l))
	if m.State != 0 {
		n += 1 + sovLeaseStatus(uint64(m.State))
	}
	l = m.Liveness.Size()
	n += 1 + l + sovLeaseStatus(uint64(l))
	l = m.RequestTime.Size()
	n += 1 + l + sovLeaseStatus(uint64(l))
	l = len(m.ErrInfo)
	if l > 0 {
		n += 1 + l + sovLeaseStatus(uint64(l))
	}
	l = m.MinValidObservedTimestamp.Size()
	n += 1 + l + sovLeaseStatus(uint64(l))
	return n
}

func sovLeaseStatus(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLeaseStatus(x uint64) (n int) {
	return sovLeaseStatus(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LeaseStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLeaseStatus
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
			return fmt.Errorf("proto: LeaseStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LeaseStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lease", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLeaseStatus
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
				return ErrInvalidLengthLeaseStatus
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLeaseStatus
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Lease.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Now", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLeaseStatus
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
				return ErrInvalidLengthLeaseStatus
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLeaseStatus
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Now.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			m.State = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLeaseStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.State |= LeaseState(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Liveness", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLeaseStatus
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
				return ErrInvalidLengthLeaseStatus
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLeaseStatus
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Liveness.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLeaseStatus
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
				return ErrInvalidLengthLeaseStatus
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLeaseStatus
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RequestTime.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ErrInfo", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLeaseStatus
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
				return ErrInvalidLengthLeaseStatus
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLeaseStatus
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ErrInfo = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinValidObservedTimestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLeaseStatus
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
				return ErrInvalidLengthLeaseStatus
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLeaseStatus
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinValidObservedTimestamp.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLeaseStatus(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLeaseStatus
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
func skipLeaseStatus(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLeaseStatus
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
					return 0, ErrIntOverflowLeaseStatus
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
					return 0, ErrIntOverflowLeaseStatus
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
				return 0, ErrInvalidLengthLeaseStatus
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLeaseStatus
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLeaseStatus
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLeaseStatus        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLeaseStatus          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLeaseStatus = fmt.Errorf("proto: unexpected end of group")
)

