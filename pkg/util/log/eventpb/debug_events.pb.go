// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: util/log/eventpb/debug_events.proto

package eventpb

import (
	fmt "fmt"
	logpb "github.com/cockroachdb/cockroach/pkg/util/log/logpb"
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

// CommonDebugEventDetails contains the fields common to all debugging events.
type CommonDebugEventDetails struct {
	// The node ID where the event originated.
	NodeID int32 `protobuf:"varint,1,opt,name=node_id,json=nodeId,proto3" json:",omitempty"`
	// The user which performed the operation.
	User string `protobuf:"bytes,2,opt,name=user,proto3" json:",omitempty"`
}

func (m *CommonDebugEventDetails) Reset()         { *m = CommonDebugEventDetails{} }
func (m *CommonDebugEventDetails) String() string { return proto.CompactTextString(m) }
func (*CommonDebugEventDetails) ProtoMessage()    {}
func (*CommonDebugEventDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd5ad4bda5741116, []int{0}
}
func (m *CommonDebugEventDetails) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CommonDebugEventDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *CommonDebugEventDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommonDebugEventDetails.Merge(m, src)
}
func (m *CommonDebugEventDetails) XXX_Size() int {
	return m.Size()
}
func (m *CommonDebugEventDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_CommonDebugEventDetails.DiscardUnknown(m)
}

var xxx_messageInfo_CommonDebugEventDetails proto.InternalMessageInfo

// DebugSendKvBatch is recorded when an arbitrary KV BatchRequest is submitted
// to the cluster via the `debug send-kv-batch` CLI command.
type DebugSendKvBatch struct {
	logpb.CommonEventDetails `protobuf:"bytes,1,opt,name=common,proto3,embedded=common" json:""`
	CommonDebugEventDetails  `protobuf:"bytes,2,opt,name=debug,proto3,embedded=debug" json:""`
	BatchRequest             string `protobuf:"bytes,3,opt,name=batch_request,json=batchRequest,proto3" json:",omitempty"`
}

func (m *DebugSendKvBatch) Reset()         { *m = DebugSendKvBatch{} }
func (m *DebugSendKvBatch) String() string { return proto.CompactTextString(m) }
func (*DebugSendKvBatch) ProtoMessage()    {}
func (*DebugSendKvBatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd5ad4bda5741116, []int{1}
}
func (m *DebugSendKvBatch) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DebugSendKvBatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *DebugSendKvBatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DebugSendKvBatch.Merge(m, src)
}
func (m *DebugSendKvBatch) XXX_Size() int {
	return m.Size()
}
func (m *DebugSendKvBatch) XXX_DiscardUnknown() {
	xxx_messageInfo_DebugSendKvBatch.DiscardUnknown(m)
}

var xxx_messageInfo_DebugSendKvBatch proto.InternalMessageInfo

// DebugRecoverReplica is recorded when unsafe loss of quorum recovery is performed.
type DebugRecoverReplica struct {
	logpb.CommonEventDetails `protobuf:"bytes,1,opt,name=common,proto3,embedded=common" json:""`
	CommonDebugEventDetails  `protobuf:"bytes,2,opt,name=debug,proto3,embedded=debug" json:""`
	RangeID                  int64  `protobuf:"varint,3,opt,name=range_id,json=rangeId,proto3" json:"range_id,omitempty"`
	StoreID                  int64  `protobuf:"varint,4,opt,name=store_id,json=storeId,proto3" json:"store_id,omitempty"`
	SurvivorReplicaID        int32  `protobuf:"varint,5,opt,name=survivor_replica_id,json=survivorReplicaId,proto3" json:"survivor_replica_id,omitempty"`
	UpdatedReplicaID         int32  `protobuf:"varint,6,opt,name=updated_replica_id,json=updatedReplicaId,proto3" json:"updated_replica_id,omitempty"`
	StartKey                 string `protobuf:"bytes,7,opt,name=start_key,json=startKey,proto3" json:"start_key,omitempty"`
	EndKey                   string `protobuf:"bytes,8,opt,name=end_key,json=endKey,proto3" json:"end_key,omitempty"`
}

func (m *DebugRecoverReplica) Reset()         { *m = DebugRecoverReplica{} }
func (m *DebugRecoverReplica) String() string { return proto.CompactTextString(m) }
func (*DebugRecoverReplica) ProtoMessage()    {}
func (*DebugRecoverReplica) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd5ad4bda5741116, []int{2}
}
func (m *DebugRecoverReplica) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DebugRecoverReplica) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *DebugRecoverReplica) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DebugRecoverReplica.Merge(m, src)
}
func (m *DebugRecoverReplica) XXX_Size() int {
	return m.Size()
}
func (m *DebugRecoverReplica) XXX_DiscardUnknown() {
	xxx_messageInfo_DebugRecoverReplica.DiscardUnknown(m)
}

var xxx_messageInfo_DebugRecoverReplica proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CommonDebugEventDetails)(nil), "cockroach.util.log.eventpb.CommonDebugEventDetails")
	proto.RegisterType((*DebugSendKvBatch)(nil), "cockroach.util.log.eventpb.DebugSendKvBatch")
	proto.RegisterType((*DebugRecoverReplica)(nil), "cockroach.util.log.eventpb.DebugRecoverReplica")
}

func init() {
	proto.RegisterFile("util/log/eventpb/debug_events.proto", fileDescriptor_cd5ad4bda5741116)
}

var fileDescriptor_cd5ad4bda5741116 = []byte{
	// 521 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x93, 0x4f, 0x6a, 0xdb, 0x40,
	0x14, 0xc6, 0xa5, 0x3a, 0x96, 0x9d, 0x71, 0x5a, 0x9c, 0x49, 0x4a, 0x8c, 0x0b, 0x52, 0x70, 0x21,
	0x64, 0x51, 0x24, 0x1a, 0xd3, 0x0b, 0xa8, 0xce, 0x42, 0xa4, 0x74, 0x21, 0x53, 0x02, 0xdd, 0x18,
	0x49, 0x33, 0xc8, 0xc2, 0x92, 0x46, 0x1d, 0x8d, 0x04, 0xbe, 0x45, 0x8f, 0xe5, 0x4d, 0xc1, 0xcb,
	0xac, 0x44, 0x2b, 0xef, 0xb2, 0xe8, 0x19, 0xca, 0x3c, 0x99, 0x38, 0x35, 0xc9, 0x01, 0xb2, 0x93,
	0xde, 0xf7, 0x9b, 0xef, 0xfd, 0x99, 0x37, 0xe8, 0x7d, 0x21, 0xa2, 0xd8, 0x8a, 0x59, 0x68, 0xd1,
	0x92, 0xa6, 0x22, 0xf3, 0x2d, 0x42, 0xfd, 0x22, 0x9c, 0xc1, 0x5f, 0x6e, 0x66, 0x9c, 0x09, 0x86,
	0x87, 0x01, 0x0b, 0x16, 0x9c, 0x79, 0xc1, 0xdc, 0x94, 0xb8, 0x19, 0xb3, 0xd0, 0xdc, 0xe2, 0xc3,
	0xd3, 0x90, 0x85, 0x0c, 0x30, 0x4b, 0x7e, 0x35, 0x27, 0x86, 0xc3, 0x07, 0xdb, 0x98, 0x85, 0x99,
	0xdf, 0x98, 0x37, 0xda, 0x28, 0x43, 0x67, 0x9f, 0x59, 0x92, 0xb0, 0x74, 0x22, 0x33, 0x5d, 0x4b,
	0x65, 0x42, 0x85, 0x17, 0xc5, 0x39, 0xfe, 0x88, 0x3a, 0x29, 0x23, 0x74, 0x16, 0x91, 0x81, 0x7a,
	0xae, 0x5e, 0xb6, 0xed, 0x41, 0x5d, 0x19, 0xda, 0x57, 0x46, 0xa8, 0x33, 0xb9, 0xaf, 0x0c, 0xf4,
	0x81, 0x25, 0x91, 0xa0, 0x49, 0x26, 0x96, 0xae, 0x26, 0x41, 0x87, 0xe0, 0x11, 0x3a, 0x28, 0x72,
	0xca, 0x07, 0xaf, 0xce, 0xd5, 0xcb, 0x43, 0xfb, 0xcd, 0x1e, 0x05, 0xda, 0xe8, 0xaf, 0x8a, 0xfa,
	0x90, 0x6c, 0x4a, 0x53, 0x72, 0x53, 0xda, 0x9e, 0x08, 0xe6, 0xf8, 0x0b, 0xd2, 0x02, 0x28, 0x03,
	0x52, 0xf5, 0xae, 0x2e, 0xcc, 0x27, 0xba, 0x6c, 0x0a, 0x7d, 0x5c, 0xa3, 0x7d, 0xb4, 0xaa, 0x0c,
	0x65, 0x5d, 0x19, 0xea, 0x7d, 0x65, 0x28, 0xee, 0xd6, 0x03, 0xdf, 0xa2, 0x36, 0x0c, 0x0e, 0xea,
	0xe8, 0x5d, 0x8d, 0xcd, 0xe7, 0x47, 0x66, 0x3e, 0xd3, 0xfd, 0x9e, 0x73, 0xe3, 0x87, 0xc7, 0xe8,
	0xb5, 0x2f, 0xeb, 0x9d, 0x71, 0xfa, 0xa3, 0xa0, 0xb9, 0x18, 0xb4, 0x9e, 0x6c, 0xf4, 0x08, 0x20,
	0xb7, 0x61, 0x46, 0xbf, 0x5a, 0xe8, 0x04, 0xfc, 0x5d, 0x1a, 0xb0, 0x92, 0x72, 0x97, 0x66, 0x71,
	0x14, 0x78, 0x2f, 0xa5, 0xe7, 0x0b, 0xd4, 0xe5, 0x5e, 0x1a, 0xc2, 0x1e, 0xc8, 0x76, 0x5b, 0x76,
	0xaf, 0xae, 0x8c, 0x8e, 0x2b, 0x63, 0xce, 0xc4, 0xed, 0x80, 0xe8, 0x10, 0xc9, 0xe5, 0x82, 0x71,
	0xe0, 0x0e, 0x76, 0xdc, 0x54, 0xc6, 0x24, 0x07, 0xa2, 0x43, 0xf0, 0x35, 0x3a, 0xc9, 0x0b, 0x5e,
	0x46, 0x25, 0xe3, 0x33, 0xde, 0x8c, 0x42, 0x1e, 0x69, 0xc3, 0x8a, 0xbd, 0xad, 0x2b, 0xe3, 0x78,
	0xba, 0x95, 0xb7, 0x83, 0x72, 0x26, 0xee, 0x71, 0xbe, 0x17, 0x22, 0xd8, 0x46, 0xb8, 0xc8, 0x88,
	0x27, 0x28, 0x79, 0xec, 0xa2, 0x81, 0xcb, 0x69, 0x5d, 0x19, 0xfd, 0x6f, 0x8d, 0xba, 0x33, 0xe9,
	0x17, 0xff, 0x47, 0x08, 0x7e, 0x87, 0x0e, 0x73, 0xe1, 0x71, 0x31, 0x5b, 0xd0, 0xe5, 0xa0, 0x23,
	0xaf, 0xd2, 0xed, 0x42, 0xe0, 0x86, 0x2e, 0xf1, 0x19, 0xea, 0xd0, 0x94, 0x80, 0xd4, 0x05, 0x49,
	0x93, 0xfb, 0x4a, 0x97, 0xf6, 0xed, 0xea, 0x8f, 0xae, 0xac, 0x6a, 0x5d, 0x5d, 0xd7, 0xba, 0x7a,
	0x57, 0xeb, 0xea, 0xef, 0x5a, 0x57, 0x7f, 0x6e, 0x74, 0x65, 0xbd, 0xd1, 0x95, 0xbb, 0x8d, 0xae,
	0x7c, 0xff, 0x14, 0x46, 0x62, 0x5e, 0xf8, 0x66, 0xc0, 0x12, 0xeb, 0xe1, 0x1a, 0x88, 0xbf, 0xfb,
	0xb6, 0xb2, 0x45, 0x68, 0xed, 0x3f, 0x76, 0x5f, 0x83, 0x27, 0x39, 0xfe, 0x17, 0x00, 0x00, 0xff,
	0xff, 0xbf, 0xcc, 0xfd, 0x9c, 0x07, 0x04, 0x00, 0x00,
}

func (m *CommonDebugEventDetails) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CommonDebugEventDetails) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CommonDebugEventDetails) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.User) > 0 {
		i -= len(m.User)
		copy(dAtA[i:], m.User)
		i = encodeVarintDebugEvents(dAtA, i, uint64(len(m.User)))
		i--
		dAtA[i] = 0x12
	}
	if m.NodeID != 0 {
		i = encodeVarintDebugEvents(dAtA, i, uint64(m.NodeID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *DebugSendKvBatch) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DebugSendKvBatch) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DebugSendKvBatch) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BatchRequest) > 0 {
		i -= len(m.BatchRequest)
		copy(dAtA[i:], m.BatchRequest)
		i = encodeVarintDebugEvents(dAtA, i, uint64(len(m.BatchRequest)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size, err := m.CommonDebugEventDetails.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintDebugEvents(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.CommonEventDetails.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintDebugEvents(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *DebugRecoverReplica) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DebugRecoverReplica) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DebugRecoverReplica) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.EndKey) > 0 {
		i -= len(m.EndKey)
		copy(dAtA[i:], m.EndKey)
		i = encodeVarintDebugEvents(dAtA, i, uint64(len(m.EndKey)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.StartKey) > 0 {
		i -= len(m.StartKey)
		copy(dAtA[i:], m.StartKey)
		i = encodeVarintDebugEvents(dAtA, i, uint64(len(m.StartKey)))
		i--
		dAtA[i] = 0x3a
	}
	if m.UpdatedReplicaID != 0 {
		i = encodeVarintDebugEvents(dAtA, i, uint64(m.UpdatedReplicaID))
		i--
		dAtA[i] = 0x30
	}
	if m.SurvivorReplicaID != 0 {
		i = encodeVarintDebugEvents(dAtA, i, uint64(m.SurvivorReplicaID))
		i--
		dAtA[i] = 0x28
	}
	if m.StoreID != 0 {
		i = encodeVarintDebugEvents(dAtA, i, uint64(m.StoreID))
		i--
		dAtA[i] = 0x20
	}
	if m.RangeID != 0 {
		i = encodeVarintDebugEvents(dAtA, i, uint64(m.RangeID))
		i--
		dAtA[i] = 0x18
	}
	{
		size, err := m.CommonDebugEventDetails.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintDebugEvents(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.CommonEventDetails.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintDebugEvents(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintDebugEvents(dAtA []byte, offset int, v uint64) int {
	offset -= sovDebugEvents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CommonDebugEventDetails) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NodeID != 0 {
		n += 1 + sovDebugEvents(uint64(m.NodeID))
	}
	l = len(m.User)
	if l > 0 {
		n += 1 + l + sovDebugEvents(uint64(l))
	}
	return n
}

func (m *DebugSendKvBatch) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.CommonEventDetails.Size()
	n += 1 + l + sovDebugEvents(uint64(l))
	l = m.CommonDebugEventDetails.Size()
	n += 1 + l + sovDebugEvents(uint64(l))
	l = len(m.BatchRequest)
	if l > 0 {
		n += 1 + l + sovDebugEvents(uint64(l))
	}
	return n
}

func (m *DebugRecoverReplica) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.CommonEventDetails.Size()
	n += 1 + l + sovDebugEvents(uint64(l))
	l = m.CommonDebugEventDetails.Size()
	n += 1 + l + sovDebugEvents(uint64(l))
	if m.RangeID != 0 {
		n += 1 + sovDebugEvents(uint64(m.RangeID))
	}
	if m.StoreID != 0 {
		n += 1 + sovDebugEvents(uint64(m.StoreID))
	}
	if m.SurvivorReplicaID != 0 {
		n += 1 + sovDebugEvents(uint64(m.SurvivorReplicaID))
	}
	if m.UpdatedReplicaID != 0 {
		n += 1 + sovDebugEvents(uint64(m.UpdatedReplicaID))
	}
	l = len(m.StartKey)
	if l > 0 {
		n += 1 + l + sovDebugEvents(uint64(l))
	}
	l = len(m.EndKey)
	if l > 0 {
		n += 1 + l + sovDebugEvents(uint64(l))
	}
	return n
}

func sovDebugEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDebugEvents(x uint64) (n int) {
	return sovDebugEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CommonDebugEventDetails) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDebugEvents
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
			return fmt.Errorf("proto: CommonDebugEventDetails: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CommonDebugEventDetails: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeID", wireType)
			}
			m.NodeID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDebugEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NodeID |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDebugEvents
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
				return ErrInvalidLengthDebugEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDebugEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.User = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDebugEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDebugEvents
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
func (m *DebugSendKvBatch) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDebugEvents
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
			return fmt.Errorf("proto: DebugSendKvBatch: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DebugSendKvBatch: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommonEventDetails", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDebugEvents
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
				return ErrInvalidLengthDebugEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDebugEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CommonEventDetails.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommonDebugEventDetails", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDebugEvents
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
				return ErrInvalidLengthDebugEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDebugEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CommonDebugEventDetails.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BatchRequest", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDebugEvents
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
				return ErrInvalidLengthDebugEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDebugEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BatchRequest = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDebugEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDebugEvents
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
func (m *DebugRecoverReplica) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDebugEvents
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
			return fmt.Errorf("proto: DebugRecoverReplica: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DebugRecoverReplica: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommonEventDetails", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDebugEvents
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
				return ErrInvalidLengthDebugEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDebugEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CommonEventDetails.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommonDebugEventDetails", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDebugEvents
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
				return ErrInvalidLengthDebugEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDebugEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CommonDebugEventDetails.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RangeID", wireType)
			}
			m.RangeID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDebugEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RangeID |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StoreID", wireType)
			}
			m.StoreID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDebugEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StoreID |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SurvivorReplicaID", wireType)
			}
			m.SurvivorReplicaID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDebugEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SurvivorReplicaID |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedReplicaID", wireType)
			}
			m.UpdatedReplicaID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDebugEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UpdatedReplicaID |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDebugEvents
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
				return ErrInvalidLengthDebugEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDebugEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StartKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDebugEvents
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
				return ErrInvalidLengthDebugEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDebugEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EndKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDebugEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDebugEvents
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
func skipDebugEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDebugEvents
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
					return 0, ErrIntOverflowDebugEvents
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
					return 0, ErrIntOverflowDebugEvents
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
				return 0, ErrInvalidLengthDebugEvents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDebugEvents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDebugEvents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDebugEvents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDebugEvents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDebugEvents = fmt.Errorf("proto: unexpected end of group")
)

