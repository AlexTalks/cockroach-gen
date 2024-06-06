// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: jobs/jobspb/schedule.proto

package jobspb

import (
	fmt "fmt"
	clusterversion "github.com/cockroachdb/cockroach/pkg/clusterversion"
	github_com_cockroachdb_cockroach_pkg_util_uuid "github.com/cockroachdb/cockroach/pkg/util/uuid"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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

// WaitBehavior describes how to handle previously  started
// jobs that have not completed yet.
type ScheduleDetails_WaitBehavior int32

const (
	// Wait for the previous run to complete
	// before starting the next one.
	ScheduleDetails_WAIT ScheduleDetails_WaitBehavior = 0
	// Do not wait for the previous run to complete.
	ScheduleDetails_NO_WAIT ScheduleDetails_WaitBehavior = 1
	// If the previous run is still running, skip this run
	// and advance schedule to the next recurrence.
	ScheduleDetails_SKIP ScheduleDetails_WaitBehavior = 2
)

var ScheduleDetails_WaitBehavior_name = map[int32]string{
	0: "WAIT",
	1: "NO_WAIT",
	2: "SKIP",
}

var ScheduleDetails_WaitBehavior_value = map[string]int32{
	"WAIT":    0,
	"NO_WAIT": 1,
	"SKIP":    2,
}

func (x ScheduleDetails_WaitBehavior) String() string {
	return proto.EnumName(ScheduleDetails_WaitBehavior_name, int32(x))
}

func (ScheduleDetails_WaitBehavior) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_935d71200db86e25, []int{0, 0}
}

// ErrorHandlingBehavior describes how to handle failed job runs.
type ScheduleDetails_ErrorHandlingBehavior int32

const (
	// By default, failed jobs will run again, based on their schedule.
	ScheduleDetails_RETRY_SCHED ScheduleDetails_ErrorHandlingBehavior = 0
	// Retry failed jobs soon.
	ScheduleDetails_RETRY_SOON ScheduleDetails_ErrorHandlingBehavior = 1
	// Stop running this schedule
	ScheduleDetails_PAUSE_SCHED ScheduleDetails_ErrorHandlingBehavior = 2
)

var ScheduleDetails_ErrorHandlingBehavior_name = map[int32]string{
	0: "RETRY_SCHED",
	1: "RETRY_SOON",
	2: "PAUSE_SCHED",
}

var ScheduleDetails_ErrorHandlingBehavior_value = map[string]int32{
	"RETRY_SCHED": 0,
	"RETRY_SOON":  1,
	"PAUSE_SCHED": 2,
}

func (x ScheduleDetails_ErrorHandlingBehavior) String() string {
	return proto.EnumName(ScheduleDetails_ErrorHandlingBehavior_name, int32(x))
}

func (ScheduleDetails_ErrorHandlingBehavior) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_935d71200db86e25, []int{0, 1}
}

// ScheduleDetails describes how to schedule and execute the job.
type ScheduleDetails struct {
	// How to handle running jobs.
	Wait ScheduleDetails_WaitBehavior `protobuf:"varint,1,opt,name=wait,proto3,enum=cockroach.jobs.jobspb.ScheduleDetails_WaitBehavior" json:"wait,omitempty"`
	// How to handle failed jobs.
	OnError ScheduleDetails_ErrorHandlingBehavior `protobuf:"varint,2,opt,name=on_error,json=onError,proto3,enum=cockroach.jobs.jobspb.ScheduleDetails_ErrorHandlingBehavior" json:"on_error,omitempty"`
	// ClusterID is populated at schedule creation with the ClusterID, in case a
	// schedule resuming later, needs to use this information, e.g. to determine if it
	// has been restored into a different cluster, which might mean it should
	// terminate, pause or update some other state.
	//
	// A resuming schedule may change the cluster_id
	// on resumption.
	ClusterID github_com_cockroachdb_cockroach_pkg_util_uuid.UUID `protobuf:"bytes,3,opt,name=cluster_id,json=clusterId,proto3,customtype=github.com/cockroachdb/cockroach/pkg/util/uuid.UUID" json:"cluster_id"`
	// CreationClusterVersion documents the cluster version this schedule was created on.
	CreationClusterVersion clusterversion.ClusterVersion `protobuf:"bytes,4,opt,name=creation_cluster_version,json=creationClusterVersion,proto3" json:"creation_cluster_version"`
}

func (m *ScheduleDetails) Reset()         { *m = ScheduleDetails{} }
func (m *ScheduleDetails) String() string { return proto.CompactTextString(m) }
func (*ScheduleDetails) ProtoMessage()    {}
func (*ScheduleDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_935d71200db86e25, []int{0}
}
func (m *ScheduleDetails) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ScheduleDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *ScheduleDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScheduleDetails.Merge(m, src)
}
func (m *ScheduleDetails) XXX_Size() int {
	return m.Size()
}
func (m *ScheduleDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_ScheduleDetails.DiscardUnknown(m)
}

var xxx_messageInfo_ScheduleDetails proto.InternalMessageInfo

// ExecutionArguments describes data needed to execute scheduled jobs.
type ExecutionArguments struct {
	Args *types.Any `protobuf:"bytes,1,opt,name=args,proto3" json:"args,omitempty"`
}

func (m *ExecutionArguments) Reset()         { *m = ExecutionArguments{} }
func (m *ExecutionArguments) String() string { return proto.CompactTextString(m) }
func (*ExecutionArguments) ProtoMessage()    {}
func (*ExecutionArguments) Descriptor() ([]byte, []int) {
	return fileDescriptor_935d71200db86e25, []int{1}
}
func (m *ExecutionArguments) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExecutionArguments) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *ExecutionArguments) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecutionArguments.Merge(m, src)
}
func (m *ExecutionArguments) XXX_Size() int {
	return m.Size()
}
func (m *ExecutionArguments) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecutionArguments.DiscardUnknown(m)
}

var xxx_messageInfo_ExecutionArguments proto.InternalMessageInfo

// Message representing sql statement to execute.
type SqlStatementExecutionArg struct {
	Statement string `protobuf:"bytes,1,opt,name=statement,proto3" json:"statement,omitempty"`
}

func (m *SqlStatementExecutionArg) Reset()         { *m = SqlStatementExecutionArg{} }
func (m *SqlStatementExecutionArg) String() string { return proto.CompactTextString(m) }
func (*SqlStatementExecutionArg) ProtoMessage()    {}
func (*SqlStatementExecutionArg) Descriptor() ([]byte, []int) {
	return fileDescriptor_935d71200db86e25, []int{2}
}
func (m *SqlStatementExecutionArg) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SqlStatementExecutionArg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *SqlStatementExecutionArg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SqlStatementExecutionArg.Merge(m, src)
}
func (m *SqlStatementExecutionArg) XXX_Size() int {
	return m.Size()
}
func (m *SqlStatementExecutionArg) XXX_DiscardUnknown() {
	xxx_messageInfo_SqlStatementExecutionArg.DiscardUnknown(m)
}

var xxx_messageInfo_SqlStatementExecutionArg proto.InternalMessageInfo

// ScheduleState represents mutable schedule state.
// The members of this proto may be mutated during each schedule execution.
type ScheduleState struct {
	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (m *ScheduleState) Reset()         { *m = ScheduleState{} }
func (m *ScheduleState) String() string { return proto.CompactTextString(m) }
func (*ScheduleState) ProtoMessage()    {}
func (*ScheduleState) Descriptor() ([]byte, []int) {
	return fileDescriptor_935d71200db86e25, []int{3}
}
func (m *ScheduleState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ScheduleState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *ScheduleState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScheduleState.Merge(m, src)
}
func (m *ScheduleState) XXX_Size() int {
	return m.Size()
}
func (m *ScheduleState) XXX_DiscardUnknown() {
	xxx_messageInfo_ScheduleState.DiscardUnknown(m)
}

var xxx_messageInfo_ScheduleState proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("cockroach.jobs.jobspb.ScheduleDetails_WaitBehavior", ScheduleDetails_WaitBehavior_name, ScheduleDetails_WaitBehavior_value)
	proto.RegisterEnum("cockroach.jobs.jobspb.ScheduleDetails_ErrorHandlingBehavior", ScheduleDetails_ErrorHandlingBehavior_name, ScheduleDetails_ErrorHandlingBehavior_value)
	proto.RegisterType((*ScheduleDetails)(nil), "cockroach.jobs.jobspb.ScheduleDetails")
	proto.RegisterType((*ExecutionArguments)(nil), "cockroach.jobs.jobspb.ExecutionArguments")
	proto.RegisterType((*SqlStatementExecutionArg)(nil), "cockroach.jobs.jobspb.SqlStatementExecutionArg")
	proto.RegisterType((*ScheduleState)(nil), "cockroach.jobs.jobspb.ScheduleState")
}

func init() { proto.RegisterFile("jobs/jobspb/schedule.proto", fileDescriptor_935d71200db86e25) }

var fileDescriptor_935d71200db86e25 = []byte{
	// 531 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x4d, 0x6f, 0xda, 0x40,
	0x10, 0xb5, 0xa9, 0x95, 0x84, 0x25, 0x4d, 0xd0, 0x2a, 0x89, 0x5c, 0x54, 0x19, 0x64, 0x55, 0x2a,
	0xa7, 0xdd, 0x0a, 0x2e, 0x3d, 0x54, 0x95, 0x20, 0xa0, 0x62, 0x55, 0x0a, 0x91, 0x1d, 0x8a, 0xda,
	0x0b, 0xf2, 0xc7, 0xd6, 0x6c, 0xe3, 0x78, 0xe9, 0x7a, 0x9d, 0x36, 0xd7, 0xfe, 0x82, 0xfe, 0x2c,
	0x8e, 0x39, 0x46, 0x3d, 0xa0, 0x16, 0xfe, 0x48, 0xe5, 0xb5, 0x49, 0x48, 0x94, 0x43, 0x2e, 0xd6,
	0xbe, 0x79, 0x6f, 0x9e, 0x67, 0x66, 0x67, 0x41, 0xed, 0x1b, 0xf3, 0x12, 0x9c, 0x7d, 0x66, 0x1e,
	0x4e, 0xfc, 0x29, 0x09, 0xd2, 0x88, 0xa0, 0x19, 0x67, 0x82, 0xc1, 0x43, 0x9f, 0xf9, 0xe7, 0x9c,
	0xb9, 0xfe, 0x14, 0x65, 0x02, 0x94, 0xab, 0x6a, 0x07, 0x21, 0x0b, 0x99, 0x54, 0xe0, 0xec, 0x94,
	0x8b, 0x6b, 0x2f, 0x42, 0xc6, 0xc2, 0x88, 0x60, 0x89, 0xbc, 0xf4, 0x2b, 0x76, 0xe3, 0xab, 0x82,
	0x7a, 0xe5, 0x47, 0x69, 0x22, 0x08, 0xbf, 0x24, 0x3c, 0xa1, 0x2c, 0xc6, 0x05, 0x9c, 0x14, 0x38,
	0x57, 0x99, 0xbf, 0x34, 0xb0, 0xef, 0x14, 0x05, 0xf4, 0x88, 0x70, 0x69, 0x94, 0xc0, 0x0f, 0x40,
	0xfb, 0xe1, 0x52, 0xa1, 0xab, 0x0d, 0xb5, 0xb9, 0xd7, 0x6a, 0xa3, 0x47, 0x0b, 0x42, 0x0f, 0xb2,
	0xd0, 0xd8, 0xa5, 0xa2, 0x4b, 0xa6, 0xee, 0x25, 0x65, 0xdc, 0x96, 0x06, 0x70, 0x0c, 0x76, 0x58,
	0x3c, 0x21, 0x9c, 0x33, 0xae, 0x97, 0xa4, 0xd9, 0xbb, 0x27, 0x9a, 0xf5, 0xb3, 0x9c, 0x81, 0x1b,
	0x07, 0x11, 0x8d, 0xc3, 0x5b, 0xd7, 0x6d, 0x16, 0x4b, 0x02, 0x86, 0x00, 0xac, 0xdb, 0xa1, 0x81,
	0xfe, 0xac, 0xa1, 0x36, 0x77, 0xbb, 0x83, 0xf9, 0xa2, 0xae, 0xfc, 0x59, 0xd4, 0xdb, 0x21, 0x15,
	0xd3, 0xd4, 0x43, 0x3e, 0xbb, 0xc0, 0xb7, 0x3f, 0x0b, 0xbc, 0xbb, 0x33, 0x9e, 0x9d, 0x87, 0x38,
	0x15, 0x34, 0xc2, 0x69, 0x4a, 0x03, 0x34, 0x1a, 0x59, 0xbd, 0xe5, 0xa2, 0x5e, 0x3e, 0xce, 0x0d,
	0xad, 0x9e, 0x5d, 0x2e, 0xbc, 0xad, 0x00, 0x4e, 0x81, 0xee, 0x73, 0xe2, 0x0a, 0xca, 0xe2, 0xc9,
	0x83, 0x01, 0xea, 0x5a, 0x43, 0x6d, 0x56, 0x5a, 0xcd, 0x8d, 0x8e, 0xee, 0x4f, 0x1c, 0x15, 0x8e,
	0x9f, 0x72, 0xd8, 0xd5, 0xb2, 0x02, 0xed, 0xa3, 0xb5, 0xdf, 0x7d, 0xd6, 0xc4, 0x60, 0x77, 0x73,
	0x82, 0x70, 0x07, 0x68, 0xe3, 0x8e, 0x75, 0x56, 0x55, 0x60, 0x05, 0x6c, 0x9f, 0x0c, 0x27, 0x12,
	0xa8, 0x59, 0xd8, 0xf9, 0x68, 0x9d, 0x56, 0x4b, 0xa6, 0x05, 0x0e, 0x1f, 0x9d, 0x12, 0xdc, 0x07,
	0x15, 0xbb, 0x7f, 0x66, 0x7f, 0x9e, 0x38, 0xc7, 0x83, 0x7e, 0xaf, 0xaa, 0xc0, 0x3d, 0x00, 0x8a,
	0xc0, 0x70, 0x78, 0x52, 0x55, 0x33, 0xc1, 0x69, 0x67, 0xe4, 0xf4, 0x0b, 0x41, 0xc9, 0x7c, 0x0f,
	0x60, 0xff, 0x27, 0xf1, 0xd3, 0xac, 0xac, 0x0e, 0x0f, 0xd3, 0x0b, 0x12, 0x8b, 0x04, 0x36, 0x81,
	0xe6, 0xf2, 0x30, 0x91, 0x6b, 0x50, 0x69, 0x1d, 0xa0, 0x7c, 0xd5, 0xd0, 0x7a, 0xd5, 0x50, 0x27,
	0xbe, 0xb2, 0xa5, 0xc2, 0x7c, 0x0b, 0x74, 0xe7, 0x7b, 0xe4, 0x08, 0x57, 0x90, 0x2c, 0x75, 0xd3,
	0x0b, 0xbe, 0x04, 0xe5, 0x64, 0x4d, 0x48, 0xab, 0xb2, 0x7d, 0x17, 0x30, 0x5f, 0x83, 0xe7, 0xeb,
	0xab, 0x97, 0xe9, 0xf0, 0x08, 0x6c, 0x65, 0x6c, 0x9a, 0x14, 0xda, 0x02, 0x75, 0xed, 0xf9, 0x3f,
	0x43, 0x99, 0x2f, 0x0d, 0xf5, 0x7a, 0x69, 0xa8, 0x37, 0x4b, 0x43, 0xfd, 0xbb, 0x34, 0xd4, 0xdf,
	0x2b, 0x43, 0xb9, 0x5e, 0x19, 0xca, 0xcd, 0xca, 0x50, 0xbe, 0xbc, 0x79, 0xd2, 0xbd, 0x6f, 0x3c,
	0x3c, 0x6f, 0x4b, 0xb6, 0xd2, 0xfe, 0x1f, 0x00, 0x00, 0xff, 0xff, 0xf5, 0xb2, 0xc5, 0xeb, 0x8e,
	0x03, 0x00, 0x00,
}

func (m *ScheduleDetails) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ScheduleDetails) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ScheduleDetails) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.CreationClusterVersion.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintSchedule(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.ClusterID.Size()
		i -= size
		if _, err := m.ClusterID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintSchedule(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.OnError != 0 {
		i = encodeVarintSchedule(dAtA, i, uint64(m.OnError))
		i--
		dAtA[i] = 0x10
	}
	if m.Wait != 0 {
		i = encodeVarintSchedule(dAtA, i, uint64(m.Wait))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ExecutionArguments) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExecutionArguments) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ExecutionArguments) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Args != nil {
		{
			size, err := m.Args.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSchedule(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SqlStatementExecutionArg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SqlStatementExecutionArg) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SqlStatementExecutionArg) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Statement) > 0 {
		i -= len(m.Statement)
		copy(dAtA[i:], m.Statement)
		i = encodeVarintSchedule(dAtA, i, uint64(len(m.Statement)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ScheduleState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ScheduleState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ScheduleState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Status) > 0 {
		i -= len(m.Status)
		copy(dAtA[i:], m.Status)
		i = encodeVarintSchedule(dAtA, i, uint64(len(m.Status)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSchedule(dAtA []byte, offset int, v uint64) int {
	offset -= sovSchedule(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ScheduleDetails) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Wait != 0 {
		n += 1 + sovSchedule(uint64(m.Wait))
	}
	if m.OnError != 0 {
		n += 1 + sovSchedule(uint64(m.OnError))
	}
	l = m.ClusterID.Size()
	n += 1 + l + sovSchedule(uint64(l))
	l = m.CreationClusterVersion.Size()
	n += 1 + l + sovSchedule(uint64(l))
	return n
}

func (m *ExecutionArguments) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Args != nil {
		l = m.Args.Size()
		n += 1 + l + sovSchedule(uint64(l))
	}
	return n
}

func (m *SqlStatementExecutionArg) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Statement)
	if l > 0 {
		n += 1 + l + sovSchedule(uint64(l))
	}
	return n
}

func (m *ScheduleState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovSchedule(uint64(l))
	}
	return n
}

func sovSchedule(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSchedule(x uint64) (n int) {
	return sovSchedule(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ScheduleDetails) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSchedule
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
			return fmt.Errorf("proto: ScheduleDetails: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ScheduleDetails: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Wait", wireType)
			}
			m.Wait = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchedule
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Wait |= ScheduleDetails_WaitBehavior(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OnError", wireType)
			}
			m.OnError = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchedule
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OnError |= ScheduleDetails_ErrorHandlingBehavior(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClusterID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchedule
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSchedule
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSchedule
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ClusterID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreationClusterVersion", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchedule
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
				return ErrInvalidLengthSchedule
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSchedule
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CreationClusterVersion.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSchedule(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSchedule
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
func (m *ExecutionArguments) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSchedule
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
			return fmt.Errorf("proto: ExecutionArguments: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExecutionArguments: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Args", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchedule
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
				return ErrInvalidLengthSchedule
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSchedule
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Args == nil {
				m.Args = &types.Any{}
			}
			if err := m.Args.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSchedule(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSchedule
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
func (m *SqlStatementExecutionArg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSchedule
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
			return fmt.Errorf("proto: SqlStatementExecutionArg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SqlStatementExecutionArg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Statement", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchedule
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
				return ErrInvalidLengthSchedule
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSchedule
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Statement = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSchedule(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSchedule
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
func (m *ScheduleState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSchedule
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
			return fmt.Errorf("proto: ScheduleState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ScheduleState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchedule
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
				return ErrInvalidLengthSchedule
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSchedule
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Status = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSchedule(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSchedule
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
func skipSchedule(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSchedule
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
					return 0, ErrIntOverflowSchedule
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
					return 0, ErrIntOverflowSchedule
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
				return 0, ErrInvalidLengthSchedule
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSchedule
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSchedule
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSchedule        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSchedule          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSchedule = fmt.Errorf("proto: unexpected end of group")
)

