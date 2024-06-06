// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: util/log/eventpb/events.proto

package eventpb

import (
	fmt "fmt"
	logpb "github.com/cockroachdb/cockroach/pkg/util/log/logpb"
	github_com_cockroachdb_redact "github.com/cockroachdb/redact"
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

// CommonSQLEventDetails contains the fields common to all
// SQL events.
type CommonSQLEventDetails struct {
	// A normalized copy of the SQL statement that triggered the event.
	// The statement string contains a mix of sensitive and non-sensitive details (it is redactable).
	Statement github_com_cockroachdb_redact.RedactableString `protobuf:"bytes,1,opt,name=statement,proto3,customtype=github.com/cockroachdb/redact.RedactableString" json:",omitempty" redact:"mixed"`
	// The statement tag. This is separate from the statement string,
	// since the statement string can contain sensitive information. The
	// tag is guaranteed not to.
	Tag string `protobuf:"bytes,6,opt,name=tag,proto3" json:",omitempty" redact:"nonsensitive"`
	// The user account that triggered the event.
	// The special usernames `root` and `node` are not considered sensitive.
	User string `protobuf:"bytes,2,opt,name=user,proto3" json:",omitempty" redact:"safeif:root|node"`
	// The primary object descriptor affected by the operation. Set to zero for operations
	// that don't affect descriptors.
	DescriptorID uint32 `protobuf:"varint,3,opt,name=descriptor_id,json=descriptorId,proto3" json:",omitempty"`
	// The application name for the session where the event was emitted.
	// This is included in the event to ease filtering of logging output
	// by application.
	ApplicationName string `protobuf:"bytes,4,opt,name=application_name,json=applicationName,proto3" json:",omitempty" redact:"nonsensitive"`
	// The mapping of SQL placeholders to their values, for prepared statements.
	PlaceholderValues []string `protobuf:"bytes,5,rep,name=placeholder_values,json=placeholderValues,proto3" json:",omitempty"`
}

func (m *CommonSQLEventDetails) Reset()         { *m = CommonSQLEventDetails{} }
func (m *CommonSQLEventDetails) String() string { return proto.CompactTextString(m) }
func (*CommonSQLEventDetails) ProtoMessage()    {}
func (*CommonSQLEventDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_656955fd5b536468, []int{0}
}
func (m *CommonSQLEventDetails) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CommonSQLEventDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *CommonSQLEventDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommonSQLEventDetails.Merge(m, src)
}
func (m *CommonSQLEventDetails) XXX_Size() int {
	return m.Size()
}
func (m *CommonSQLEventDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_CommonSQLEventDetails.DiscardUnknown(m)
}

var xxx_messageInfo_CommonSQLEventDetails proto.InternalMessageInfo

// CommonJobEventDetails contains the fields common to all job events.
type CommonJobEventDetails struct {
	// The ID of the job that triggered the event.
	JobID int64 `protobuf:"varint,1,opt,name=job_id,json=jobId,proto3" json:",omitempty"`
	// The type of the job that triggered the event.
	JobType string `protobuf:"bytes,2,opt,name=job_type,json=jobType,proto3" json:",omitempty" redact:"nonsensitive"`
	// A description of the job that triggered the event. Some jobs populate the
	// description with an approximate representation of the SQL statement run to
	// create the job.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:",omitempty"`
	// The user account that triggered the event.
	User string `protobuf:"bytes,4,opt,name=user,proto3" json:",omitempty"`
	// The object descriptors affected by the job. Set to zero for operations
	// that don't affect descriptors.
	DescriptorIDs []uint32 `protobuf:"varint,5,rep,packed,name=descriptor_ids,json=descriptorIds,proto3" json:",omitempty"`
	// The status of the job that triggered the event. This allows the job to
	// indicate which phase execution it is in when the event is triggered.
	Status string `protobuf:"bytes,6,opt,name=status,proto3" json:",omitempty" redact:"nonsensitive"`
}

func (m *CommonJobEventDetails) Reset()         { *m = CommonJobEventDetails{} }
func (m *CommonJobEventDetails) String() string { return proto.CompactTextString(m) }
func (*CommonJobEventDetails) ProtoMessage()    {}
func (*CommonJobEventDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_656955fd5b536468, []int{1}
}
func (m *CommonJobEventDetails) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CommonJobEventDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *CommonJobEventDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommonJobEventDetails.Merge(m, src)
}
func (m *CommonJobEventDetails) XXX_Size() int {
	return m.Size()
}
func (m *CommonJobEventDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_CommonJobEventDetails.DiscardUnknown(m)
}

var xxx_messageInfo_CommonJobEventDetails proto.InternalMessageInfo

// CommonChangefeedEventDetails contains the fields common to all changefeed events
type CommonChangefeedEventDetails struct {
	logpb.CommonEventDetails `protobuf:"bytes,1,opt,name=common,proto3,embedded=common" json:""`
	// The description of that would show up in the job's description field, redacted
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:",omitempty"`
	// The type of sink being emitted to (ex: kafka, nodelocal, webhook-https).
	SinkType string `protobuf:"bytes,3,opt,name=sink_type,json=sinkType,proto3" json:",omitempty" redact:"nonsensitive"`
	// The number of tables listed in the query that the changefeed is to run on.
	NumTables int32 `protobuf:"varint,4,opt,name=num_tables,json=numTables,proto3" json:",omitempty"`
	// The behavior of emitted resolved spans (ex: yes, no, 10s)
	Resolved string `protobuf:"bytes,5,opt,name=resolved,proto3" json:",omitempty" redact:"nonsensitive"`
	// The desired behavior of initial scans (ex: yes, no, only)
	InitialScan string `protobuf:"bytes,6,opt,name=initial_scan,json=initialScan,proto3" json:",omitempty" redact:"nonsensitive"`
	// The data format being emitted (ex: JSON, Avro).
	Format string `protobuf:"bytes,7,opt,name=format,proto3" json:",omitempty" redact:"nonsensitive"`
}

func (m *CommonChangefeedEventDetails) Reset()         { *m = CommonChangefeedEventDetails{} }
func (m *CommonChangefeedEventDetails) String() string { return proto.CompactTextString(m) }
func (*CommonChangefeedEventDetails) ProtoMessage()    {}
func (*CommonChangefeedEventDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_656955fd5b536468, []int{2}
}
func (m *CommonChangefeedEventDetails) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CommonChangefeedEventDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *CommonChangefeedEventDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommonChangefeedEventDetails.Merge(m, src)
}
func (m *CommonChangefeedEventDetails) XXX_Size() int {
	return m.Size()
}
func (m *CommonChangefeedEventDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_CommonChangefeedEventDetails.DiscardUnknown(m)
}

var xxx_messageInfo_CommonChangefeedEventDetails proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CommonSQLEventDetails)(nil), "cockroach.util.log.eventpb.CommonSQLEventDetails")
	proto.RegisterType((*CommonJobEventDetails)(nil), "cockroach.util.log.eventpb.CommonJobEventDetails")
	proto.RegisterType((*CommonChangefeedEventDetails)(nil), "cockroach.util.log.eventpb.CommonChangefeedEventDetails")
}

func init() { proto.RegisterFile("util/log/eventpb/events.proto", fileDescriptor_656955fd5b536468) }

var fileDescriptor_656955fd5b536468 = []byte{
	// 686 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0x5f, 0x6b, 0xd3, 0x50,
	0x18, 0xc6, 0x9b, 0x75, 0xed, 0xd6, 0xb3, 0x76, 0x6a, 0x70, 0x18, 0x8b, 0x26, 0x23, 0x17, 0x3a,
	0x50, 0x53, 0x51, 0xbc, 0x70, 0x20, 0x83, 0xb4, 0x0a, 0x1b, 0x43, 0xb0, 0x1b, 0x13, 0xbc, 0x29,
	0x27, 0xc9, 0x69, 0x7a, 0xb6, 0xe4, 0xbc, 0x21, 0xe7, 0xb4, 0x38, 0xf0, 0x43, 0xf8, 0xb1, 0x76,
	0xb9, 0xcb, 0xe1, 0x45, 0x70, 0x19, 0xde, 0xec, 0xd2, 0x4f, 0x20, 0x39, 0x69, 0xbb, 0xfe, 0x99,
	0x68, 0xaf, 0x12, 0x78, 0x9f, 0xdf, 0xfb, 0xe6, 0x7d, 0x9e, 0x9c, 0x83, 0x1e, 0xf7, 0x05, 0x0d,
	0x1a, 0x01, 0xf8, 0x0d, 0x32, 0x20, 0x4c, 0x44, 0x4e, 0xfe, 0xe4, 0x56, 0x14, 0x83, 0x00, 0xb5,
	0xee, 0x82, 0x7b, 0x12, 0x03, 0x76, 0x7b, 0x56, 0x26, 0xb4, 0x02, 0xf0, 0xad, 0xa1, 0xb0, 0x7e,
	0xdf, 0x07, 0x1f, 0xa4, 0xac, 0x91, 0xbd, 0xe5, 0x44, 0xbd, 0x3e, 0x6e, 0x18, 0x80, 0x3f, 0x6a,
	0x97, 0xd7, 0xcc, 0xcb, 0x22, 0xda, 0x68, 0x42, 0x18, 0x02, 0x3b, 0xf8, 0xb4, 0xff, 0x3e, 0x2b,
	0xb4, 0x88, 0xc0, 0x34, 0xe0, 0xaa, 0x40, 0x15, 0x2e, 0xb0, 0x20, 0x21, 0x61, 0x42, 0x53, 0x36,
	0x95, 0xad, 0x8a, 0x7d, 0x74, 0x96, 0x18, 0x85, 0x1f, 0x89, 0x61, 0xf9, 0x54, 0xf4, 0xfa, 0x8e,
	0xe5, 0x42, 0xd8, 0x18, 0x7f, 0x8d, 0xe7, 0x34, 0x62, 0xe2, 0x61, 0x57, 0x58, 0x6d, 0xf9, 0xc0,
	0x4e, 0x40, 0x0e, 0x44, 0x4c, 0x99, 0x7f, 0x9d, 0x18, 0xe8, 0x39, 0x84, 0x54, 0x90, 0x30, 0x12,
	0xa7, 0xbf, 0x13, 0x63, 0x3d, 0x17, 0x6e, 0x9b, 0x21, 0xfd, 0x4a, 0x3c, 0xb3, 0x7d, 0x33, 0x48,
	0x7d, 0x8b, 0x8a, 0x02, 0xfb, 0x5a, 0x59, 0xce, 0x7b, 0x3a, 0x47, 0x6e, 0x8c, 0x48, 0x06, 0x8c,
	0x13, 0xc6, 0xa9, 0xa0, 0x03, 0x62, 0xb6, 0x33, 0x46, 0xdd, 0x41, 0xcb, 0x7d, 0x4e, 0x62, 0x6d,
	0x49, 0xb2, 0xcf, 0xe6, 0xd8, 0x87, 0x23, 0x96, 0xe3, 0x2e, 0xa1, 0xdd, 0xed, 0x18, 0x40, 0x7c,
	0x63, 0xe0, 0x11, 0xb3, 0x2d, 0x41, 0xb5, 0x89, 0x6a, 0x1e, 0xe1, 0x6e, 0x4c, 0x23, 0x01, 0x71,
	0x87, 0x7a, 0x5a, 0x71, 0x53, 0xd9, 0xaa, 0xd9, 0x7a, 0x9a, 0x18, 0xd5, 0xd6, 0xb8, 0xb0, 0xdb,
	0x9a, 0xee, 0xdc, 0xae, 0xde, 0x40, 0xbb, 0x9e, 0xda, 0x46, 0x77, 0x71, 0x14, 0x05, 0xd4, 0xc5,
	0x82, 0x02, 0xeb, 0x30, 0x1c, 0x12, 0x6d, 0x79, 0xb1, 0x6d, 0xee, 0x4c, 0x34, 0xf8, 0x88, 0x43,
	0xa2, 0xbe, 0x43, 0x6a, 0x14, 0x60, 0x97, 0xf4, 0x20, 0xf0, 0x48, 0xdc, 0x19, 0xe0, 0xa0, 0x4f,
	0xb8, 0x56, 0xda, 0x2c, 0x6e, 0x55, 0xec, 0xf5, 0x99, 0xaf, 0xb9, 0x37, 0xa1, 0x3c, 0x92, 0x42,
	0xf3, 0x7a, 0x69, 0x94, 0xf1, 0x1e, 0x38, 0x53, 0x19, 0x5b, 0xa8, 0x7c, 0x0c, 0x4e, 0xb6, 0x6a,
	0x16, 0x70, 0xd1, 0x7e, 0x90, 0x26, 0x46, 0x69, 0x0f, 0x9c, 0xb9, 0x1d, 0x4b, 0xc7, 0xe0, 0xec,
	0x7a, 0xaa, 0x8d, 0x56, 0x33, 0xbd, 0x38, 0x8d, 0xc8, 0xd0, 0xe6, 0xff, 0x5e, 0x6a, 0xe5, 0x18,
	0x9c, 0xc3, 0xd3, 0x88, 0xa8, 0x2f, 0xd1, 0xda, 0xc8, 0x30, 0x0a, 0x4c, 0x7a, 0x3c, 0xbf, 0xc5,
	0xa4, 0x44, 0x35, 0x87, 0xc1, 0x2e, 0xdf, 0x2a, 0xcd, 0xb3, 0xfb, 0x80, 0xd6, 0xa7, 0xb2, 0xcb,
	0xed, 0xa9, 0xd9, 0x46, 0x9a, 0x18, 0xb5, 0xc9, 0xf0, 0xf8, 0x0c, 0x5e, 0x9b, 0x4c, 0x8f, 0xab,
	0x3b, 0xa8, 0x9c, 0xfd, 0x8c, 0x7d, 0xbe, 0xe8, 0x2f, 0x38, 0xc4, 0xcc, 0x5f, 0x45, 0xf4, 0x28,
	0x37, 0xbb, 0xd9, 0xc3, 0xcc, 0x27, 0x5d, 0x42, 0xbc, 0x29, 0xcf, 0xf7, 0x51, 0xd9, 0x95, 0x75,
	0xe9, 0xf9, 0xda, 0xab, 0x27, 0xd6, 0x2d, 0x07, 0x3a, 0xef, 0x30, 0xc9, 0xd9, 0xd5, 0xec, 0xf0,
	0x9d, 0x27, 0x86, 0x72, 0x9d, 0x18, 0x85, 0xf6, 0xb0, 0xc7, 0xac, 0x9b, 0x4b, 0xff, 0x76, 0xb3,
	0x85, 0x2a, 0x9c, 0xb2, 0x93, 0x3c, 0xc4, 0xe2, 0x62, 0x4b, 0xae, 0x66, 0xa4, 0x4c, 0xf1, 0x05,
	0x42, 0xac, 0x1f, 0x76, 0xe4, 0x29, 0xe7, 0x32, 0x99, 0xd2, 0xdc, 0xd8, 0x0a, 0xeb, 0x87, 0x87,
	0x52, 0xa0, 0x36, 0xd1, 0x6a, 0x4c, 0x38, 0x04, 0x03, 0xe2, 0x69, 0xa5, 0x05, 0x67, 0x8e, 0x40,
	0x75, 0x0f, 0x55, 0x29, 0xa3, 0x82, 0xe2, 0xa0, 0xc3, 0x5d, 0xcc, 0x16, 0x4d, 0x68, 0x6d, 0x08,
	0x1f, 0xb8, 0x98, 0x65, 0x39, 0x77, 0x21, 0x0e, 0xb1, 0xd0, 0x56, 0x16, 0xcc, 0x39, 0xc7, 0xec,
	0xcf, 0x67, 0x97, 0x7a, 0xe1, 0x2c, 0xd5, 0x95, 0xf3, 0x54, 0x57, 0x2e, 0x52, 0x5d, 0xf9, 0x99,
	0xea, 0xca, 0xf7, 0x2b, 0xbd, 0x70, 0x7e, 0xa5, 0x17, 0x2e, 0xae, 0xf4, 0xc2, 0x97, 0x37, 0x7f,
	0xb9, 0x25, 0xc7, 0xef, 0x8d, 0xe8, 0xc4, 0x6f, 0xcc, 0x5e, 0xf6, 0x4e, 0x59, 0x5e, 0xcc, 0xaf,
	0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x99, 0x37, 0x4a, 0xbb, 0x07, 0x06, 0x00, 0x00,
}

func (m *CommonSQLEventDetails) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CommonSQLEventDetails) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CommonSQLEventDetails) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Tag) > 0 {
		i -= len(m.Tag)
		copy(dAtA[i:], m.Tag)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Tag)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.PlaceholderValues) > 0 {
		for iNdEx := len(m.PlaceholderValues) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.PlaceholderValues[iNdEx])
			copy(dAtA[i:], m.PlaceholderValues[iNdEx])
			i = encodeVarintEvents(dAtA, i, uint64(len(m.PlaceholderValues[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.ApplicationName) > 0 {
		i -= len(m.ApplicationName)
		copy(dAtA[i:], m.ApplicationName)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.ApplicationName)))
		i--
		dAtA[i] = 0x22
	}
	if m.DescriptorID != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.DescriptorID))
		i--
		dAtA[i] = 0x18
	}
	if len(m.User) > 0 {
		i -= len(m.User)
		copy(dAtA[i:], m.User)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.User)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Statement) > 0 {
		i -= len(m.Statement)
		copy(dAtA[i:], m.Statement)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Statement)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CommonJobEventDetails) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CommonJobEventDetails) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CommonJobEventDetails) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Status) > 0 {
		i -= len(m.Status)
		copy(dAtA[i:], m.Status)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Status)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.DescriptorIDs) > 0 {
		dAtA2 := make([]byte, len(m.DescriptorIDs)*10)
		var j1 int
		for _, num := range m.DescriptorIDs {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintEvents(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.User) > 0 {
		i -= len(m.User)
		copy(dAtA[i:], m.User)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.User)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.JobType) > 0 {
		i -= len(m.JobType)
		copy(dAtA[i:], m.JobType)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.JobType)))
		i--
		dAtA[i] = 0x12
	}
	if m.JobID != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.JobID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *CommonChangefeedEventDetails) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CommonChangefeedEventDetails) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CommonChangefeedEventDetails) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Format) > 0 {
		i -= len(m.Format)
		copy(dAtA[i:], m.Format)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Format)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.InitialScan) > 0 {
		i -= len(m.InitialScan)
		copy(dAtA[i:], m.InitialScan)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.InitialScan)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Resolved) > 0 {
		i -= len(m.Resolved)
		copy(dAtA[i:], m.Resolved)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Resolved)))
		i--
		dAtA[i] = 0x2a
	}
	if m.NumTables != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.NumTables))
		i--
		dAtA[i] = 0x20
	}
	if len(m.SinkType) > 0 {
		i -= len(m.SinkType)
		copy(dAtA[i:], m.SinkType)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.SinkType)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.CommonEventDetails.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintEvents(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CommonSQLEventDetails) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Statement)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.User)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.DescriptorID != 0 {
		n += 1 + sovEvents(uint64(m.DescriptorID))
	}
	l = len(m.ApplicationName)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if len(m.PlaceholderValues) > 0 {
		for _, s := range m.PlaceholderValues {
			l = len(s)
			n += 1 + l + sovEvents(uint64(l))
		}
	}
	l = len(m.Tag)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *CommonJobEventDetails) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.JobID != 0 {
		n += 1 + sovEvents(uint64(m.JobID))
	}
	l = len(m.JobType)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.User)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if len(m.DescriptorIDs) > 0 {
		l = 0
		for _, e := range m.DescriptorIDs {
			l += sovEvents(uint64(e))
		}
		n += 1 + sovEvents(uint64(l)) + l
	}
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *CommonChangefeedEventDetails) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.CommonEventDetails.Size()
	n += 1 + l + sovEvents(uint64(l))
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.SinkType)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.NumTables != 0 {
		n += 1 + sovEvents(uint64(m.NumTables))
	}
	l = len(m.Resolved)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.InitialScan)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Format)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CommonSQLEventDetails) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: CommonSQLEventDetails: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CommonSQLEventDetails: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Statement", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Statement = github_com_cockroachdb_redact.RedactableString(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.User = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DescriptorID", wireType)
			}
			m.DescriptorID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DescriptorID |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApplicationName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ApplicationName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlaceholderValues", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PlaceholderValues = append(m.PlaceholderValues, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tag", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tag = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *CommonJobEventDetails) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: CommonJobEventDetails: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CommonJobEventDetails: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field JobID", wireType)
			}
			m.JobID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.JobID |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field JobType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.JobType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.User = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType == 0 {
				var v uint32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowEvents
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.DescriptorIDs = append(m.DescriptorIDs, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowEvents
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthEvents
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthEvents
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.DescriptorIDs) == 0 {
					m.DescriptorIDs = make([]uint32, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowEvents
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint32(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.DescriptorIDs = append(m.DescriptorIDs, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field DescriptorIDs", wireType)
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Status = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *CommonChangefeedEventDetails) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: CommonChangefeedEventDetails: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CommonChangefeedEventDetails: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommonEventDetails", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
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
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SinkType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SinkType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumTables", wireType)
			}
			m.NumTables = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumTables |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Resolved", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Resolved = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialScan", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InitialScan = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Format", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Format = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func skipEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
				return 0, ErrInvalidLengthEvents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvents = fmt.Errorf("proto: unexpected end of group")
)

