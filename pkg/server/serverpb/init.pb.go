// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: server/serverpb/init.proto

package serverpb

import (
	context "context"
	fmt "fmt"
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

type InitType int32

const (
	InitType_DEFAULT           InitType = 0
	InitType_NONE              InitType = 1
	InitType_VIRTUALIZED_EMPTY InitType = 2
	InitType_VIRTUALIZED       InitType = 3
)

var InitType_name = map[int32]string{
	0: "DEFAULT",
	1: "NONE",
	2: "VIRTUALIZED_EMPTY",
	3: "VIRTUALIZED",
}

var InitType_value = map[string]int32{
	"DEFAULT":           0,
	"NONE":              1,
	"VIRTUALIZED_EMPTY": 2,
	"VIRTUALIZED":       3,
}

func (x InitType) String() string {
	return proto.EnumName(InitType_name, int32(x))
}

func (InitType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f8cc8040be9166ce, []int{0}
}

type BootstrapRequest struct {
	InitType InitType `protobuf:"varint,1,opt,name=init_type,json=initType,proto3,enum=cockroach.server.serverpb.InitType" json:"init_type,omitempty"`
}

func (m *BootstrapRequest) Reset()         { *m = BootstrapRequest{} }
func (m *BootstrapRequest) String() string { return proto.CompactTextString(m) }
func (*BootstrapRequest) ProtoMessage()    {}
func (*BootstrapRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8cc8040be9166ce, []int{0}
}
func (m *BootstrapRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BootstrapRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *BootstrapRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BootstrapRequest.Merge(m, src)
}
func (m *BootstrapRequest) XXX_Size() int {
	return m.Size()
}
func (m *BootstrapRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BootstrapRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BootstrapRequest proto.InternalMessageInfo

type BootstrapResponse struct {
}

func (m *BootstrapResponse) Reset()         { *m = BootstrapResponse{} }
func (m *BootstrapResponse) String() string { return proto.CompactTextString(m) }
func (*BootstrapResponse) ProtoMessage()    {}
func (*BootstrapResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8cc8040be9166ce, []int{1}
}
func (m *BootstrapResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BootstrapResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *BootstrapResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BootstrapResponse.Merge(m, src)
}
func (m *BootstrapResponse) XXX_Size() int {
	return m.Size()
}
func (m *BootstrapResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BootstrapResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BootstrapResponse proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("cockroach.server.serverpb.InitType", InitType_name, InitType_value)
	proto.RegisterType((*BootstrapRequest)(nil), "cockroach.server.serverpb.BootstrapRequest")
	proto.RegisterType((*BootstrapResponse)(nil), "cockroach.server.serverpb.BootstrapResponse")
}

func init() { proto.RegisterFile("server/serverpb/init.proto", fileDescriptor_f8cc8040be9166ce) }

var fileDescriptor_f8cc8040be9166ce = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2a, 0x4e, 0x2d, 0x2a,
	0x4b, 0x2d, 0xd2, 0x87, 0x50, 0x05, 0x49, 0xfa, 0x99, 0x79, 0x99, 0x25, 0x7a, 0x05, 0x45, 0xf9,
	0x25, 0xf9, 0x42, 0x92, 0xc9, 0xf9, 0xc9, 0xd9, 0x45, 0xf9, 0x89, 0xc9, 0x19, 0x7a, 0x10, 0x69,
	0x3d, 0x98, 0x2a, 0xa5, 0x10, 0x2e, 0x01, 0xa7, 0xfc, 0xfc, 0x92, 0xe2, 0x92, 0xa2, 0xc4, 0x82,
	0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x07, 0x2e, 0x4e, 0x90, 0xe6, 0xf8, 0x92, 0xca,
	0x82, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x3e, 0x23, 0x65, 0x3d, 0x9c, 0x46, 0xe8, 0x79, 0xe6,
	0x65, 0x96, 0x84, 0x54, 0x16, 0xa4, 0x06, 0x71, 0x64, 0x42, 0x59, 0x4a, 0xc2, 0x5c, 0x82, 0x48,
	0xa6, 0x16, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x6a, 0x79, 0x72, 0x71, 0xc0, 0x94, 0x0a, 0x71, 0x73,
	0xb1, 0xbb, 0xb8, 0xba, 0x39, 0x86, 0xfa, 0x84, 0x08, 0x30, 0x08, 0x71, 0x70, 0xb1, 0xf8, 0xf9,
	0xfb, 0xb9, 0x0a, 0x30, 0x0a, 0x89, 0x72, 0x09, 0x86, 0x79, 0x06, 0x85, 0x84, 0x3a, 0xfa, 0x78,
	0x46, 0xb9, 0xba, 0xc4, 0xbb, 0xfa, 0x06, 0x84, 0x44, 0x0a, 0x30, 0x09, 0xf1, 0x73, 0x71, 0x23,
	0x09, 0x0b, 0x30, 0x1b, 0x15, 0x70, 0xb1, 0x80, 0x8c, 0x12, 0xca, 0xe0, 0xe2, 0x84, 0xdb, 0x23,
	0xa4, 0x8d, 0xc7, 0x8d, 0xe8, 0x7e, 0x94, 0xd2, 0x21, 0x4e, 0x31, 0xc4, 0xe9, 0x4a, 0x0c, 0x4e,
	0x61, 0x27, 0x1e, 0xca, 0x31, 0x9c, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x8d, 0x47,
	0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d,
	0xc7, 0x72, 0x0c, 0x51, 0x26, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa,
	0x70, 0xb3, 0x53, 0x92, 0x10, 0x6c, 0xfd, 0x82, 0xec, 0x74, 0x7d, 0xb4, 0x58, 0x4a, 0x62, 0x03,
	0xc7, 0x90, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x86, 0xdc, 0xf3, 0x38, 0xbf, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// InitClient is the client API for Init service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type InitClient interface {
	// Bootstrap bootstraps an uninitialized cluster. This is primarily used by
	// cockroach init. It writes the data for a single-node cluster comprised of
	// this node (with NodeID "1") to the first store (StoreID "1"), after which
	// the node can start and allow other nodes to join the cluster.
	Bootstrap(ctx context.Context, in *BootstrapRequest, opts ...grpc.CallOption) (*BootstrapResponse, error)
}

type initClient struct {
	cc *grpc.ClientConn
}

func NewInitClient(cc *grpc.ClientConn) InitClient {
	return &initClient{cc}
}

func (c *initClient) Bootstrap(ctx context.Context, in *BootstrapRequest, opts ...grpc.CallOption) (*BootstrapResponse, error) {
	out := new(BootstrapResponse)
	err := c.cc.Invoke(ctx, "/cockroach.server.serverpb.Init/Bootstrap", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InitServer is the server API for Init service.
type InitServer interface {
	// Bootstrap bootstraps an uninitialized cluster. This is primarily used by
	// cockroach init. It writes the data for a single-node cluster comprised of
	// this node (with NodeID "1") to the first store (StoreID "1"), after which
	// the node can start and allow other nodes to join the cluster.
	Bootstrap(context.Context, *BootstrapRequest) (*BootstrapResponse, error)
}

// UnimplementedInitServer can be embedded to have forward compatible implementations.
type UnimplementedInitServer struct {
}

func (*UnimplementedInitServer) Bootstrap(ctx context.Context, req *BootstrapRequest) (*BootstrapResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Bootstrap not implemented")
}

func RegisterInitServer(s *grpc.Server, srv InitServer) {
	s.RegisterService(&_Init_serviceDesc, srv)
}

func _Init_Bootstrap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BootstrapRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InitServer).Bootstrap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cockroach.server.serverpb.Init/Bootstrap",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InitServer).Bootstrap(ctx, req.(*BootstrapRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Init_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cockroach.server.serverpb.Init",
	HandlerType: (*InitServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Bootstrap",
			Handler:    _Init_Bootstrap_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server/serverpb/init.proto",
}

func (m *BootstrapRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BootstrapRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BootstrapRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.InitType != 0 {
		i = encodeVarintInit(dAtA, i, uint64(m.InitType))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *BootstrapResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BootstrapResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BootstrapResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintInit(dAtA []byte, offset int, v uint64) int {
	offset -= sovInit(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BootstrapRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.InitType != 0 {
		n += 1 + sovInit(uint64(m.InitType))
	}
	return n
}

func (m *BootstrapResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovInit(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozInit(x uint64) (n int) {
	return sovInit(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BootstrapRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInit
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
			return fmt.Errorf("proto: BootstrapRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BootstrapRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitType", wireType)
			}
			m.InitType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.InitType |= InitType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipInit(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInit
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
func (m *BootstrapResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInit
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
			return fmt.Errorf("proto: BootstrapResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BootstrapResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipInit(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInit
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
func skipInit(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowInit
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
					return 0, ErrIntOverflowInit
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
					return 0, ErrIntOverflowInit
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
				return 0, ErrInvalidLengthInit
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupInit
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthInit
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthInit        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowInit          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupInit = fmt.Errorf("proto: unexpected end of group")
)

