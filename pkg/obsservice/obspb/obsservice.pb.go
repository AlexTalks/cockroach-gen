// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: obsservice/obspb/obsservice.proto

package obspb

import (
	fmt "fmt"
	v11 "github.com/cockroachdb/cockroach/pkg/obsservice/obspb/opentelemetry-proto/common/v1"
	v12 "github.com/cockroachdb/cockroach/pkg/obsservice/obspb/opentelemetry-proto/logs/v1"
	v1 "github.com/cockroachdb/cockroach/pkg/obsservice/obspb/opentelemetry-proto/resource/v1"
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

// A generic event record used within the Observability Service.
// Generally, the data within log_record is eventually transformed
// into an event-specific protobuf message for further processing,
// but this message represents the event in its raw form.
type Event struct {
	// The resource for the event.
	// If this field is not set then resource info is unknown.
	// Contains information referring to the source of the event.
	// For example, cluster ID, node ID, etc.
	Resource *v1.Resource `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	// The instrumentation scope information for the event. Contains
	// event-specific information. For example, event type and version.
	Scope *v11.InstrumentationScope `protobuf:"bytes,2,opt,name=scope,proto3" json:"scope,omitempty"`
	// The LogRecord containing the specific event information.
	LogRecord *v12.LogRecord `protobuf:"bytes,3,opt,name=log_record,json=logRecord,proto3" json:"log_record,omitempty"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca722c7377006edd, []int{0}
}
func (m *Event) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return m.Size()
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

type EventInfo struct {
	Timestamp *time.Time `protobuf:"bytes,1,opt,name=timestamp,proto3,stdtime" json:"timestamp,omitempty"`
	EventID   string     `protobuf:"bytes,2,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	OrgID     string     `protobuf:"bytes,3,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	ClusterID string     `protobuf:"bytes,4,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	TenantID  string     `protobuf:"bytes,5,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
}

func (m *EventInfo) Reset()         { *m = EventInfo{} }
func (m *EventInfo) String() string { return proto.CompactTextString(m) }
func (*EventInfo) ProtoMessage()    {}
func (*EventInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca722c7377006edd, []int{1}
}
func (m *EventInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *EventInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventInfo.Merge(m, src)
}
func (m *EventInfo) XXX_Size() int {
	return m.Size()
}
func (m *EventInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_EventInfo.DiscardUnknown(m)
}

var xxx_messageInfo_EventInfo proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Event)(nil), "cockroach.obspb.Event")
	proto.RegisterType((*EventInfo)(nil), "cockroach.obspb.EventInfo")
}

func init() { proto.RegisterFile("obsservice/obspb/obsservice.proto", fileDescriptor_ca722c7377006edd) }

var fileDescriptor_ca722c7377006edd = []byte{
	// 480 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x6b, 0xdb, 0x30,
	0x14, 0xc7, 0xe3, 0xb5, 0x6e, 0x63, 0x75, 0x63, 0x60, 0x76, 0x08, 0x39, 0xd8, 0x5d, 0x0f, 0x65,
	0x85, 0x4d, 0x26, 0x2b, 0x63, 0x87, 0xc1, 0x06, 0x59, 0x72, 0x10, 0x0c, 0x06, 0x5a, 0x61, 0xb0,
	0x4b, 0x89, 0x65, 0x55, 0x35, 0xb5, 0xfd, 0x8c, 0x24, 0x1b, 0xf6, 0x2d, 0xfa, 0xb1, 0x7a, 0xec,
	0xb1, 0xa7, 0xac, 0x73, 0xee, 0xfb, 0x0c, 0xc3, 0x92, 0x9d, 0xb0, 0xb5, 0x87, 0x9c, 0xfc, 0xb7,
	0xde, 0xef, 0xff, 0x7f, 0x3c, 0x3d, 0xa1, 0x97, 0x10, 0x2b, 0xc5, 0x65, 0x9d, 0x32, 0x1e, 0x41,
	0xac, 0xca, 0x38, 0xda, 0x1c, 0xe0, 0x52, 0x82, 0x06, 0xff, 0x39, 0x03, 0x76, 0x25, 0x61, 0xc1,
	0x2e, 0xb1, 0x21, 0xc6, 0x2f, 0x04, 0x08, 0x30, 0xb5, 0xa8, 0x55, 0x16, 0x1b, 0x87, 0x02, 0x40,
	0x64, 0x3c, 0x32, 0x7f, 0x71, 0x75, 0x11, 0xe9, 0x34, 0xe7, 0x4a, 0x2f, 0xf2, 0xb2, 0x03, 0xde,
	0x3f, 0x6c, 0x55, 0xf2, 0x42, 0xf3, 0x8c, 0xe7, 0x5c, 0xcb, 0x9f, 0x6f, 0x6c, 0x62, 0x06, 0x42,
	0x45, 0xf5, 0xc4, 0x7c, 0x3b, 0xe3, 0x87, 0xad, 0x8c, 0x0c, 0xf2, 0x1c, 0x8a, 0xd6, 0x6a, 0x55,
	0x67, 0xfe, 0xb4, 0x95, 0x59, 0x72, 0x05, 0x95, 0x64, 0xbc, 0xb5, 0xf7, 0xda, 0x06, 0x1c, 0xdd,
	0x3b, 0xc8, 0x9d, 0xd7, 0xbc, 0xd0, 0xfe, 0x1c, 0x0d, 0xfb, 0xda, 0xc8, 0x39, 0x74, 0x5e, 0x1d,
	0xbc, 0x3d, 0xc1, 0xff, 0x84, 0x59, 0x07, 0x5e, 0x07, 0xd4, 0x13, 0x4c, 0x3b, 0x4d, 0xd7, 0x56,
	0x9f, 0x20, 0x57, 0x31, 0x28, 0xf9, 0xe8, 0x89, 0xc9, 0x38, 0x7d, 0x34, 0xa3, 0x9b, 0xa1, 0x9e,
	0x60, 0x52, 0x28, 0x2d, 0xab, 0x9c, 0x17, 0x7a, 0xa1, 0x53, 0x28, 0xbe, 0xb5, 0x56, 0x6a, 0x13,
	0xfc, 0x39, 0x42, 0x19, 0x88, 0x73, 0xc9, 0x19, 0xc8, 0x64, 0xb4, 0x63, 0xf2, 0x8e, 0x1f, 0xcd,
	0x33, 0xd7, 0x59, 0x4f, 0xf0, 0x17, 0x10, 0xd4, 0xd0, 0xd4, 0xcb, 0x7a, 0x79, 0xf4, 0xc7, 0x41,
	0x9e, 0x19, 0x91, 0x14, 0x17, 0xe0, 0x7f, 0x44, 0xde, 0x7a, 0x75, 0xdd, 0x9c, 0x63, 0x6c, 0x97,
	0x8b, 0xfb, 0xe5, 0xe2, 0xb3, 0x9e, 0x98, 0xee, 0x5e, 0xff, 0x0a, 0x1d, 0xba, 0xb1, 0xf8, 0xc7,
	0x68, 0xc8, 0xdb, 0xb0, 0xf3, 0x34, 0x31, 0x23, 0x7a, 0xd3, 0x83, 0x66, 0x19, 0xee, 0xdb, 0x06,
	0x33, 0xba, 0x6f, 0x8a, 0x24, 0xf1, 0x0f, 0xd1, 0x1e, 0x48, 0xd1, 0x52, 0x3b, 0x86, 0xf2, 0x9a,
	0x65, 0xe8, 0x7e, 0x95, 0x82, 0xcc, 0xa8, 0x0b, 0x52, 0x90, 0xc4, 0x7f, 0x8d, 0x10, 0xcb, 0x2a,
	0xa5, 0xb9, 0x6c, 0xa9, 0x5d, 0x43, 0x3d, 0x6b, 0x96, 0xa1, 0xf7, 0xd9, 0x9e, 0x92, 0x19, 0xf5,
	0x3a, 0x80, 0x24, 0xfe, 0x09, 0xf2, 0x34, 0x2f, 0x16, 0xb6, 0xb1, 0x6b, 0xe0, 0xa7, 0xcd, 0x32,
	0x1c, 0x9e, 0x99, 0x43, 0x32, 0xa3, 0x43, 0x5b, 0x26, 0xc9, 0xf4, 0xfb, 0xcd, 0xef, 0x60, 0x70,
	0xd3, 0x04, 0xce, 0x6d, 0x13, 0x38, 0x77, 0x4d, 0xe0, 0xdc, 0x37, 0x81, 0x73, 0xbd, 0x0a, 0x06,
	0xb7, 0xab, 0x60, 0x70, 0xb7, 0x0a, 0x06, 0x3f, 0xde, 0x89, 0x54, 0x5f, 0x56, 0x71, 0xbb, 0x87,
	0x68, 0xfd, 0xfe, 0x93, 0x78, 0xa3, 0xa3, 0xf2, 0x4a, 0x44, 0xff, 0xbf, 0xad, 0x78, 0xcf, 0x5c,
	0xd0, 0xe9, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x53, 0x87, 0x13, 0x3b, 0x57, 0x03, 0x00, 0x00,
}

func (m *Event) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Event) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Event) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LogRecord != nil {
		{
			size, err := m.LogRecord.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintObsservice(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.Scope != nil {
		{
			size, err := m.Scope.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintObsservice(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Resource != nil {
		{
			size, err := m.Resource.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintObsservice(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TenantID) > 0 {
		i -= len(m.TenantID)
		copy(dAtA[i:], m.TenantID)
		i = encodeVarintObsservice(dAtA, i, uint64(len(m.TenantID)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.ClusterID) > 0 {
		i -= len(m.ClusterID)
		copy(dAtA[i:], m.ClusterID)
		i = encodeVarintObsservice(dAtA, i, uint64(len(m.ClusterID)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.OrgID) > 0 {
		i -= len(m.OrgID)
		copy(dAtA[i:], m.OrgID)
		i = encodeVarintObsservice(dAtA, i, uint64(len(m.OrgID)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.EventID) > 0 {
		i -= len(m.EventID)
		copy(dAtA[i:], m.EventID)
		i = encodeVarintObsservice(dAtA, i, uint64(len(m.EventID)))
		i--
		dAtA[i] = 0x12
	}
	if m.Timestamp != nil {
		n4, err4 := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.Timestamp, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(*m.Timestamp):])
		if err4 != nil {
			return 0, err4
		}
		i -= n4
		i = encodeVarintObsservice(dAtA, i, uint64(n4))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintObsservice(dAtA []byte, offset int, v uint64) int {
	offset -= sovObsservice(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Event) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Resource != nil {
		l = m.Resource.Size()
		n += 1 + l + sovObsservice(uint64(l))
	}
	if m.Scope != nil {
		l = m.Scope.Size()
		n += 1 + l + sovObsservice(uint64(l))
	}
	if m.LogRecord != nil {
		l = m.LogRecord.Size()
		n += 1 + l + sovObsservice(uint64(l))
	}
	return n
}

func (m *EventInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Timestamp != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.Timestamp)
		n += 1 + l + sovObsservice(uint64(l))
	}
	l = len(m.EventID)
	if l > 0 {
		n += 1 + l + sovObsservice(uint64(l))
	}
	l = len(m.OrgID)
	if l > 0 {
		n += 1 + l + sovObsservice(uint64(l))
	}
	l = len(m.ClusterID)
	if l > 0 {
		n += 1 + l + sovObsservice(uint64(l))
	}
	l = len(m.TenantID)
	if l > 0 {
		n += 1 + l + sovObsservice(uint64(l))
	}
	return n
}

func sovObsservice(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozObsservice(x uint64) (n int) {
	return sovObsservice(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Event) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowObsservice
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
			return fmt.Errorf("proto: Event: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Event: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Resource", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObsservice
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
				return ErrInvalidLengthObsservice
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthObsservice
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Resource == nil {
				m.Resource = &v1.Resource{}
			}
			if err := m.Resource.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Scope", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObsservice
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
				return ErrInvalidLengthObsservice
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthObsservice
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Scope == nil {
				m.Scope = &v11.InstrumentationScope{}
			}
			if err := m.Scope.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LogRecord", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObsservice
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
				return ErrInvalidLengthObsservice
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthObsservice
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LogRecord == nil {
				m.LogRecord = &v12.LogRecord{}
			}
			if err := m.LogRecord.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipObsservice(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthObsservice
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
func (m *EventInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowObsservice
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
			return fmt.Errorf("proto: EventInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObsservice
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
				return ErrInvalidLengthObsservice
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthObsservice
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Timestamp == nil {
				m.Timestamp = new(time.Time)
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(m.Timestamp, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EventID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObsservice
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
				return ErrInvalidLengthObsservice
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthObsservice
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EventID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrgID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObsservice
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
				return ErrInvalidLengthObsservice
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthObsservice
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OrgID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClusterID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObsservice
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
				return ErrInvalidLengthObsservice
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthObsservice
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClusterID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TenantID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObsservice
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
				return ErrInvalidLengthObsservice
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthObsservice
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TenantID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipObsservice(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthObsservice
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
func skipObsservice(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowObsservice
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
					return 0, ErrIntOverflowObsservice
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
					return 0, ErrIntOverflowObsservice
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
				return 0, ErrInvalidLengthObsservice
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupObsservice
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthObsservice
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthObsservice        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowObsservice          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupObsservice = fmt.Errorf("proto: unexpected end of group")
)

