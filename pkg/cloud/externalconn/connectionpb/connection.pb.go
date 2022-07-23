// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cloud/externalconn/connectionpb/connection.proto

package connectionpb

import (
	fmt "fmt"
	cloudpb "github.com/cockroachdb/cockroach/pkg/cloud/cloudpb"
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

type ConnectionType int32

const (
	TypeUnspecified ConnectionType = 0
	TypeStorage     ConnectionType = 1
)

var ConnectionType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "STORAGE",
}

var ConnectionType_value = map[string]int32{
	"UNSPECIFIED": 0,
	"STORAGE":     1,
}

func (x ConnectionType) String() string {
	return proto.EnumName(ConnectionType_name, int32(x))
}

func (ConnectionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f2ab8ae20df1e701, []int{0}
}

type ConnectionDetails struct {
	// Types that are valid to be assigned to Details:
	//	*ConnectionDetails_Nodelocal
	Details isConnectionDetails_Details `protobuf_oneof:"details"`
}

func (m *ConnectionDetails) Reset()         { *m = ConnectionDetails{} }
func (m *ConnectionDetails) String() string { return proto.CompactTextString(m) }
func (*ConnectionDetails) ProtoMessage()    {}
func (*ConnectionDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2ab8ae20df1e701, []int{0}
}
func (m *ConnectionDetails) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConnectionDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *ConnectionDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectionDetails.Merge(m, src)
}
func (m *ConnectionDetails) XXX_Size() int {
	return m.Size()
}
func (m *ConnectionDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectionDetails.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectionDetails proto.InternalMessageInfo

type isConnectionDetails_Details interface {
	isConnectionDetails_Details()
	MarshalTo([]byte) (int, error)
	Size() int
}

type ConnectionDetails_Nodelocal struct {
	Nodelocal *NodelocalConnectionDetails `protobuf:"bytes,2,opt,name=nodelocal,proto3,oneof" json:"nodelocal,omitempty"`
}

func (*ConnectionDetails_Nodelocal) isConnectionDetails_Details() {}

func (m *ConnectionDetails) GetDetails() isConnectionDetails_Details {
	if m != nil {
		return m.Details
	}
	return nil
}

func (m *ConnectionDetails) GetNodelocal() *NodelocalConnectionDetails {
	if x, ok := m.GetDetails().(*ConnectionDetails_Nodelocal); ok {
		return x.Nodelocal
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ConnectionDetails) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ConnectionDetails_Nodelocal)(nil),
	}
}

type NodelocalConnectionDetails struct {
	Cfg cloudpb.LocalFileConfig `protobuf:"bytes,1,opt,name=cfg,proto3" json:"cfg"`
}

func (m *NodelocalConnectionDetails) Reset()         { *m = NodelocalConnectionDetails{} }
func (m *NodelocalConnectionDetails) String() string { return proto.CompactTextString(m) }
func (*NodelocalConnectionDetails) ProtoMessage()    {}
func (*NodelocalConnectionDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2ab8ae20df1e701, []int{1}
}
func (m *NodelocalConnectionDetails) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NodelocalConnectionDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *NodelocalConnectionDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodelocalConnectionDetails.Merge(m, src)
}
func (m *NodelocalConnectionDetails) XXX_Size() int {
	return m.Size()
}
func (m *NodelocalConnectionDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_NodelocalConnectionDetails.DiscardUnknown(m)
}

var xxx_messageInfo_NodelocalConnectionDetails proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("cockroach.cloud.externalconn.connectionpb.ConnectionType", ConnectionType_name, ConnectionType_value)
	proto.RegisterType((*ConnectionDetails)(nil), "cockroach.cloud.externalconn.connectionpb.ConnectionDetails")
	proto.RegisterType((*NodelocalConnectionDetails)(nil), "cockroach.cloud.externalconn.connectionpb.NodelocalConnectionDetails")
}

func init() {
	proto.RegisterFile("cloud/externalconn/connectionpb/connection.proto", fileDescriptor_f2ab8ae20df1e701)
}

var fileDescriptor_f2ab8ae20df1e701 = []byte{
	// 343 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x48, 0xce, 0xc9, 0x2f,
	0x4d, 0xd1, 0x4f, 0xad, 0x28, 0x49, 0x2d, 0xca, 0x4b, 0xcc, 0x49, 0xce, 0xcf, 0xcb, 0xd3, 0x07,
	0x11, 0xa9, 0xc9, 0x25, 0x99, 0xf9, 0x79, 0x05, 0x49, 0x48, 0x1c, 0xbd, 0x82, 0xa2, 0xfc, 0x92,
	0x7c, 0x21, 0xcd, 0xe4, 0xfc, 0xe4, 0xec, 0xa2, 0xfc, 0xc4, 0xe4, 0x0c, 0x3d, 0xb0, 0x5e, 0x3d,
	0x64, 0xbd, 0x7a, 0xc8, 0x7a, 0xa5, 0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0xba, 0xf4, 0x41, 0x2c,
	0x88, 0x01, 0x52, 0x2a, 0x10, 0x2b, 0xc1, 0x64, 0x41, 0x12, 0xdc, 0xea, 0xf8, 0xe2, 0x92, 0xfc,
	0xa2, 0xc4, 0xf4, 0x54, 0x88, 0x2a, 0xa5, 0x56, 0x46, 0x2e, 0x41, 0x67, 0xb8, 0x61, 0x2e, 0xa9,
	0x25, 0x89, 0x99, 0x39, 0xc5, 0x42, 0xa9, 0x5c, 0x9c, 0x79, 0xf9, 0x29, 0xa9, 0x39, 0xf9, 0xc9,
	0x89, 0x39, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0xae, 0x7a, 0x44, 0x3b, 0x48, 0xcf, 0x0f,
	0xa6, 0x17, 0xc3, 0x64, 0x0f, 0x86, 0x20, 0x84, 0xc9, 0x4e, 0x9c, 0x5c, 0xec, 0x29, 0x10, 0x71,
	0xa5, 0x38, 0x2e, 0x29, 0xdc, 0xba, 0x84, 0x1c, 0xb8, 0x98, 0x93, 0xd3, 0xd2, 0x25, 0x18, 0xc1,
	0x2e, 0xd1, 0xc0, 0x70, 0x09, 0xd4, 0x8f, 0x7a, 0x3e, 0x20, 0xdd, 0x6e, 0x99, 0x39, 0xa9, 0xce,
	0xf9, 0x79, 0x69, 0x99, 0xe9, 0x4e, 0x2c, 0x27, 0xee, 0xc9, 0x33, 0x04, 0x81, 0xb4, 0x6a, 0x45,
	0x71, 0xf1, 0x21, 0x8c, 0x0d, 0xa9, 0x2c, 0x48, 0x15, 0x52, 0xe1, 0xe2, 0x0e, 0xf5, 0x0b, 0x0e,
	0x70, 0x75, 0xf6, 0x74, 0xf3, 0x74, 0x75, 0x11, 0x60, 0x90, 0x12, 0xee, 0x9a, 0xab, 0xc0, 0x0f,
	0x92, 0x0a, 0xcd, 0x2b, 0x2e, 0x48, 0x4d, 0xce, 0x4c, 0xcb, 0x4c, 0x4d, 0x11, 0x92, 0xe1, 0x62,
	0x0f, 0x0e, 0xf1, 0x0f, 0x72, 0x74, 0x77, 0x15, 0x60, 0x94, 0xe2, 0xef, 0x9a, 0xab, 0xc0, 0x0d,
	0x52, 0x11, 0x0c, 0x09, 0x43, 0x29, 0x96, 0x8e, 0xc5, 0x72, 0x0c, 0x4e, 0x7a, 0x27, 0x1e, 0xca,
	0x31, 0x9c, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x8d, 0x47, 0x72, 0x8c, 0x0f, 0x1e,
	0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51,
	0x3c, 0xc8, 0xc1, 0x93, 0xc4, 0x06, 0x0e, 0x7a, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x03,
	0x03, 0xe1, 0xd5, 0x15, 0x02, 0x00, 0x00,
}

func (m *ConnectionDetails) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConnectionDetails) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ConnectionDetails) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Details != nil {
		{
			size := m.Details.Size()
			i -= size
			if _, err := m.Details.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *ConnectionDetails_Nodelocal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ConnectionDetails_Nodelocal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Nodelocal != nil {
		{
			size, err := m.Nodelocal.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintConnection(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *NodelocalConnectionDetails) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NodelocalConnectionDetails) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NodelocalConnectionDetails) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Cfg.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintConnection(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintConnection(dAtA []byte, offset int, v uint64) int {
	offset -= sovConnection(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ConnectionDetails) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Details != nil {
		n += m.Details.Size()
	}
	return n
}

func (m *ConnectionDetails_Nodelocal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Nodelocal != nil {
		l = m.Nodelocal.Size()
		n += 1 + l + sovConnection(uint64(l))
	}
	return n
}
func (m *NodelocalConnectionDetails) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Cfg.Size()
	n += 1 + l + sovConnection(uint64(l))
	return n
}

func sovConnection(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozConnection(x uint64) (n int) {
	return sovConnection(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ConnectionDetails) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConnection
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
			return fmt.Errorf("proto: ConnectionDetails: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConnectionDetails: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nodelocal", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConnection
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
				return ErrInvalidLengthConnection
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConnection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &NodelocalConnectionDetails{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Details = &ConnectionDetails_Nodelocal{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConnection(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthConnection
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
func (m *NodelocalConnectionDetails) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConnection
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
			return fmt.Errorf("proto: NodelocalConnectionDetails: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NodelocalConnectionDetails: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cfg", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConnection
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
				return ErrInvalidLengthConnection
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConnection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Cfg.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConnection(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthConnection
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
func skipConnection(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConnection
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
					return 0, ErrIntOverflowConnection
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
					return 0, ErrIntOverflowConnection
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
				return 0, ErrInvalidLengthConnection
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupConnection
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthConnection
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthConnection        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConnection          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupConnection = fmt.Errorf("proto: unexpected end of group")
)
