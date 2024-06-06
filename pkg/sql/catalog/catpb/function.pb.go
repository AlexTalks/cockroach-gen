// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sql/catalog/catpb/function.proto

package catpb

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

type Function_Volatility int32

const (
	Function_UNKNOWN_VOLATILITY Function_Volatility = 0
	Function_VOLATILE           Function_Volatility = 1
	Function_IMMUTABLE          Function_Volatility = 2
	Function_STABLE             Function_Volatility = 3
)

var Function_Volatility_name = map[int32]string{
	0: "UNKNOWN_VOLATILITY",
	1: "VOLATILE",
	2: "IMMUTABLE",
	3: "STABLE",
}

var Function_Volatility_value = map[string]int32{
	"UNKNOWN_VOLATILITY": 0,
	"VOLATILE":           1,
	"IMMUTABLE":          2,
	"STABLE":             3,
}

func (x Function_Volatility) Enum() *Function_Volatility {
	p := new(Function_Volatility)
	*p = x
	return p
}

func (x Function_Volatility) String() string {
	return proto.EnumName(Function_Volatility_name, int32(x))
}

func (x *Function_Volatility) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Function_Volatility_value, data, "Function_Volatility")
	if err != nil {
		return err
	}
	*x = Function_Volatility(value)
	return nil
}

func (Function_Volatility) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e00fc847fc4944a7, []int{0, 0}
}

type Function_NullInputBehavior int32

const (
	Function_UNKNOWN_NULL_INPUT_BEHAVIOR Function_NullInputBehavior = 0
	Function_CALLED_ON_NULL_INPUT        Function_NullInputBehavior = 1
	Function_RETURNS_NULL_ON_NULL_INPUT  Function_NullInputBehavior = 2
	Function_STRICT                      Function_NullInputBehavior = 3
)

var Function_NullInputBehavior_name = map[int32]string{
	0: "UNKNOWN_NULL_INPUT_BEHAVIOR",
	1: "CALLED_ON_NULL_INPUT",
	2: "RETURNS_NULL_ON_NULL_INPUT",
	3: "STRICT",
}

var Function_NullInputBehavior_value = map[string]int32{
	"UNKNOWN_NULL_INPUT_BEHAVIOR": 0,
	"CALLED_ON_NULL_INPUT":        1,
	"RETURNS_NULL_ON_NULL_INPUT":  2,
	"STRICT":                      3,
}

func (x Function_NullInputBehavior) Enum() *Function_NullInputBehavior {
	p := new(Function_NullInputBehavior)
	*p = x
	return p
}

func (x Function_NullInputBehavior) String() string {
	return proto.EnumName(Function_NullInputBehavior_name, int32(x))
}

func (x *Function_NullInputBehavior) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Function_NullInputBehavior_value, data, "Function_NullInputBehavior")
	if err != nil {
		return err
	}
	*x = Function_NullInputBehavior(value)
	return nil
}

func (Function_NullInputBehavior) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e00fc847fc4944a7, []int{0, 1}
}

type Function_Language int32

const (
	Function_UNKNOWN_LANGUAGE Function_Language = 0
	Function_SQL              Function_Language = 1
	Function_PLPGSQL          Function_Language = 2
)

var Function_Language_name = map[int32]string{
	0: "UNKNOWN_LANGUAGE",
	1: "SQL",
	2: "PLPGSQL",
}

var Function_Language_value = map[string]int32{
	"UNKNOWN_LANGUAGE": 0,
	"SQL":              1,
	"PLPGSQL":          2,
}

func (x Function_Language) Enum() *Function_Language {
	p := new(Function_Language)
	*p = x
	return p
}

func (x Function_Language) String() string {
	return proto.EnumName(Function_Language_name, int32(x))
}

func (x *Function_Language) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Function_Language_value, data, "Function_Language")
	if err != nil {
		return err
	}
	*x = Function_Language(value)
	return nil
}

func (Function_Language) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e00fc847fc4944a7, []int{0, 2}
}

type Function_Param_Class int32

const (
	Function_Param_DEFAULT  Function_Param_Class = 0
	Function_Param_IN       Function_Param_Class = 1
	Function_Param_OUT      Function_Param_Class = 2
	Function_Param_IN_OUT   Function_Param_Class = 3
	Function_Param_VARIADIC Function_Param_Class = 4
)

var Function_Param_Class_name = map[int32]string{
	0: "DEFAULT",
	1: "IN",
	2: "OUT",
	3: "IN_OUT",
	4: "VARIADIC",
}

var Function_Param_Class_value = map[string]int32{
	"DEFAULT":  0,
	"IN":       1,
	"OUT":      2,
	"IN_OUT":   3,
	"VARIADIC": 4,
}

func (x Function_Param_Class) Enum() *Function_Param_Class {
	p := new(Function_Param_Class)
	*p = x
	return p
}

func (x Function_Param_Class) String() string {
	return proto.EnumName(Function_Param_Class_name, int32(x))
}

func (x *Function_Param_Class) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Function_Param_Class_value, data, "Function_Param_Class")
	if err != nil {
		return err
	}
	*x = Function_Param_Class(value)
	return nil
}

func (Function_Param_Class) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e00fc847fc4944a7, []int{0, 0, 0}
}

// Function contains a few enum types of function properties
type Function struct {
}

func (m *Function) Reset()         { *m = Function{} }
func (m *Function) String() string { return proto.CompactTextString(m) }
func (*Function) ProtoMessage()    {}
func (*Function) Descriptor() ([]byte, []int) {
	return fileDescriptor_e00fc847fc4944a7, []int{0}
}
func (m *Function) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Function) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *Function) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Function.Merge(m, src)
}
func (m *Function) XXX_Size() int {
	return m.Size()
}
func (m *Function) XXX_DiscardUnknown() {
	xxx_messageInfo_Function.DiscardUnknown(m)
}

var xxx_messageInfo_Function proto.InternalMessageInfo

type Function_Param struct {
}

func (m *Function_Param) Reset()         { *m = Function_Param{} }
func (m *Function_Param) String() string { return proto.CompactTextString(m) }
func (*Function_Param) ProtoMessage()    {}
func (*Function_Param) Descriptor() ([]byte, []int) {
	return fileDescriptor_e00fc847fc4944a7, []int{0, 0}
}
func (m *Function_Param) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Function_Param) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *Function_Param) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Function_Param.Merge(m, src)
}
func (m *Function_Param) XXX_Size() int {
	return m.Size()
}
func (m *Function_Param) XXX_DiscardUnknown() {
	xxx_messageInfo_Function_Param.DiscardUnknown(m)
}

var xxx_messageInfo_Function_Param proto.InternalMessageInfo

// These wrappers are for the convenience of referencing the enum types from a
// proto3 definition.
type FunctionVolatility struct {
	Volatility Function_Volatility `protobuf:"varint,1,opt,name=volatility,enum=cockroach.sql.catalog.catpb.Function_Volatility" json:"volatility"`
}

func (m *FunctionVolatility) Reset()         { *m = FunctionVolatility{} }
func (m *FunctionVolatility) String() string { return proto.CompactTextString(m) }
func (*FunctionVolatility) ProtoMessage()    {}
func (*FunctionVolatility) Descriptor() ([]byte, []int) {
	return fileDescriptor_e00fc847fc4944a7, []int{1}
}
func (m *FunctionVolatility) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FunctionVolatility) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *FunctionVolatility) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FunctionVolatility.Merge(m, src)
}
func (m *FunctionVolatility) XXX_Size() int {
	return m.Size()
}
func (m *FunctionVolatility) XXX_DiscardUnknown() {
	xxx_messageInfo_FunctionVolatility.DiscardUnknown(m)
}

var xxx_messageInfo_FunctionVolatility proto.InternalMessageInfo

type FunctionNullInputBehavior struct {
	NullInputBehavior Function_NullInputBehavior `protobuf:"varint,1,opt,name=nullInputBehavior,enum=cockroach.sql.catalog.catpb.Function_NullInputBehavior" json:"nullInputBehavior"`
}

func (m *FunctionNullInputBehavior) Reset()         { *m = FunctionNullInputBehavior{} }
func (m *FunctionNullInputBehavior) String() string { return proto.CompactTextString(m) }
func (*FunctionNullInputBehavior) ProtoMessage()    {}
func (*FunctionNullInputBehavior) Descriptor() ([]byte, []int) {
	return fileDescriptor_e00fc847fc4944a7, []int{2}
}
func (m *FunctionNullInputBehavior) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FunctionNullInputBehavior) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *FunctionNullInputBehavior) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FunctionNullInputBehavior.Merge(m, src)
}
func (m *FunctionNullInputBehavior) XXX_Size() int {
	return m.Size()
}
func (m *FunctionNullInputBehavior) XXX_DiscardUnknown() {
	xxx_messageInfo_FunctionNullInputBehavior.DiscardUnknown(m)
}

var xxx_messageInfo_FunctionNullInputBehavior proto.InternalMessageInfo

type FunctionLanguage struct {
	Lang Function_Language `protobuf:"varint,1,opt,name=lang,enum=cockroach.sql.catalog.catpb.Function_Language" json:"lang"`
}

func (m *FunctionLanguage) Reset()         { *m = FunctionLanguage{} }
func (m *FunctionLanguage) String() string { return proto.CompactTextString(m) }
func (*FunctionLanguage) ProtoMessage()    {}
func (*FunctionLanguage) Descriptor() ([]byte, []int) {
	return fileDescriptor_e00fc847fc4944a7, []int{3}
}
func (m *FunctionLanguage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FunctionLanguage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *FunctionLanguage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FunctionLanguage.Merge(m, src)
}
func (m *FunctionLanguage) XXX_Size() int {
	return m.Size()
}
func (m *FunctionLanguage) XXX_DiscardUnknown() {
	xxx_messageInfo_FunctionLanguage.DiscardUnknown(m)
}

var xxx_messageInfo_FunctionLanguage proto.InternalMessageInfo

type FunctionParamClass struct {
	Class Function_Param_Class `protobuf:"varint,1,opt,name=class,enum=cockroach.sql.catalog.catpb.Function_Param_Class" json:"class"`
}

func (m *FunctionParamClass) Reset()         { *m = FunctionParamClass{} }
func (m *FunctionParamClass) String() string { return proto.CompactTextString(m) }
func (*FunctionParamClass) ProtoMessage()    {}
func (*FunctionParamClass) Descriptor() ([]byte, []int) {
	return fileDescriptor_e00fc847fc4944a7, []int{4}
}
func (m *FunctionParamClass) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FunctionParamClass) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *FunctionParamClass) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FunctionParamClass.Merge(m, src)
}
func (m *FunctionParamClass) XXX_Size() int {
	return m.Size()
}
func (m *FunctionParamClass) XXX_DiscardUnknown() {
	xxx_messageInfo_FunctionParamClass.DiscardUnknown(m)
}

var xxx_messageInfo_FunctionParamClass proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("cockroach.sql.catalog.catpb.Function_Volatility", Function_Volatility_name, Function_Volatility_value)
	proto.RegisterEnum("cockroach.sql.catalog.catpb.Function_NullInputBehavior", Function_NullInputBehavior_name, Function_NullInputBehavior_value)
	proto.RegisterEnum("cockroach.sql.catalog.catpb.Function_Language", Function_Language_name, Function_Language_value)
	proto.RegisterEnum("cockroach.sql.catalog.catpb.Function_Param_Class", Function_Param_Class_name, Function_Param_Class_value)
	proto.RegisterType((*Function)(nil), "cockroach.sql.catalog.catpb.Function")
	proto.RegisterType((*Function_Param)(nil), "cockroach.sql.catalog.catpb.Function.Param")
	proto.RegisterType((*FunctionVolatility)(nil), "cockroach.sql.catalog.catpb.FunctionVolatility")
	proto.RegisterType((*FunctionNullInputBehavior)(nil), "cockroach.sql.catalog.catpb.FunctionNullInputBehavior")
	proto.RegisterType((*FunctionLanguage)(nil), "cockroach.sql.catalog.catpb.FunctionLanguage")
	proto.RegisterType((*FunctionParamClass)(nil), "cockroach.sql.catalog.catpb.FunctionParamClass")
}

func init() { proto.RegisterFile("sql/catalog/catpb/function.proto", fileDescriptor_e00fc847fc4944a7) }

var fileDescriptor_e00fc847fc4944a7 = []byte{
	// 538 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xcd, 0x8e, 0xd2, 0x50,
	0x14, 0xc7, 0x7b, 0xf9, 0x98, 0xc1, 0xe3, 0x47, 0xee, 0xdc, 0x10, 0x33, 0x32, 0x49, 0x99, 0x74,
	0xe5, 0xaa, 0x55, 0x17, 0x63, 0xe2, 0xc6, 0x14, 0xe8, 0x40, 0xe3, 0xa5, 0x60, 0x69, 0xf1, 0x63,
	0x43, 0x4a, 0x1d, 0x4b, 0x43, 0xa7, 0x97, 0x81, 0x32, 0x89, 0x3e, 0x84, 0xf1, 0x11, 0x7c, 0x1c,
	0x96, 0xb3, 0x73, 0x56, 0x46, 0x61, 0xe3, 0x63, 0x98, 0xb6, 0xb7, 0x80, 0x92, 0x18, 0x56, 0x3d,
	0xe7, 0xf6, 0x7f, 0x7e, 0xff, 0x93, 0x7f, 0x6f, 0xe1, 0x74, 0x76, 0x15, 0x28, 0xae, 0x13, 0x39,
	0x01, 0xf3, 0xe2, 0xe7, 0x64, 0xa8, 0x7c, 0x9c, 0x87, 0x6e, 0xe4, 0xb3, 0x50, 0x9e, 0x4c, 0x59,
	0xc4, 0xc8, 0x89, 0xcb, 0xdc, 0xf1, 0x94, 0x39, 0xee, 0x48, 0x9e, 0x5d, 0x05, 0x32, 0xd7, 0xca,
	0x89, 0xb6, 0x52, 0xf6, 0x98, 0xc7, 0x12, 0x9d, 0x12, 0x57, 0xe9, 0x88, 0xf4, 0x3d, 0x07, 0xa5,
	0x73, 0x4e, 0xa9, 0xb4, 0xa0, 0xd8, 0x75, 0xa6, 0xce, 0xa5, 0xf4, 0x12, 0x8a, 0xf5, 0xc0, 0x99,
	0xcd, 0xc8, 0x5d, 0x38, 0x6c, 0x68, 0xe7, 0xaa, 0x4d, 0x2d, 0x2c, 0x90, 0x03, 0xc8, 0xe9, 0x06,
	0x46, 0xe4, 0x10, 0xf2, 0x1d, 0xdb, 0xc2, 0x39, 0x02, 0x70, 0xa0, 0x1b, 0x83, 0xb8, 0xce, 0x93,
	0x7b, 0x50, 0xea, 0xab, 0xa6, 0xae, 0x36, 0xf4, 0x3a, 0x2e, 0x48, 0x6d, 0x80, 0x3e, 0x0b, 0x9c,
	0xc8, 0x0f, 0xfc, 0xe8, 0x13, 0x79, 0x08, 0xc4, 0x36, 0x5e, 0x19, 0x9d, 0x37, 0xc6, 0xa0, 0xdf,
	0xa1, 0xaa, 0xa5, 0x53, 0xdd, 0x7a, 0x87, 0x85, 0x64, 0x26, 0xed, 0x35, 0x8c, 0xc8, 0x7d, 0xb8,
	0xa3, 0xb7, 0xdb, 0xb6, 0xa5, 0xd6, 0xa8, 0x96, 0xc2, 0x7b, 0x69, 0x9d, 0x97, 0x3e, 0xc3, 0x91,
	0x31, 0x0f, 0x02, 0x3d, 0x9c, 0xcc, 0xa3, 0xda, 0xc5, 0xc8, 0xb9, 0xf6, 0xd9, 0x94, 0x54, 0xe1,
	0x24, 0xa3, 0x1a, 0x36, 0xa5, 0x03, 0xdd, 0xe8, 0xda, 0xd6, 0xa0, 0xa6, 0xb5, 0xd4, 0xbe, 0xde,
	0x31, 0xb1, 0x40, 0x8e, 0xa1, 0x5c, 0x57, 0x29, 0xd5, 0x1a, 0x83, 0xce, 0xb6, 0x04, 0x23, 0x22,
	0x42, 0xc5, 0xd4, 0x2c, 0xdb, 0x34, 0x7a, 0xe9, 0xf9, 0xdf, 0xef, 0xb9, 0xb7, 0xa9, 0xd7, 0x2d,
	0x9c, 0x97, 0xce, 0xa0, 0x44, 0x9d, 0xd0, 0x9b, 0x3b, 0xde, 0x05, 0x29, 0x03, 0xce, 0x2c, 0xa9,
	0x6a, 0x34, 0x6d, 0xb5, 0xa9, 0x61, 0x21, 0xce, 0xa3, 0xf7, 0x9a, 0x62, 0x14, 0xa7, 0xd5, 0xa5,
	0xdd, 0x66, 0xdc, 0xe4, 0xa4, 0x29, 0x90, 0x2c, 0xd8, 0xad, 0x28, 0xfa, 0x00, 0xd7, 0xeb, 0xee,
	0x18, 0x9d, 0xa2, 0xc7, 0x0f, 0x9e, 0x3d, 0x91, 0xff, 0xf3, 0xdd, 0xe4, 0x0c, 0x22, 0x6f, 0x28,
	0xb5, 0xc2, 0xe2, 0x47, 0x55, 0x30, 0xb7, 0x48, 0x2f, 0x0a, 0xbf, 0xbf, 0x55, 0x91, 0xf4, 0x05,
	0xc1, 0xa3, 0x4c, 0xbf, 0x1b, 0xd8, 0x18, 0x8e, 0xc2, 0x7f, 0x0f, 0xf9, 0x0a, 0xcf, 0xf7, 0x5b,
	0x61, 0x87, 0xc9, 0x37, 0xd9, 0xe5, 0xf2, 0x85, 0x86, 0x80, 0xb3, 0xe1, 0x75, 0x88, 0x2d, 0x28,
	0x04, 0x4e, 0xe8, 0x71, 0x67, 0x79, 0x3f, 0xe7, 0x6c, 0x9a, 0x1b, 0x26, 0x04, 0xee, 0xe1, 0x6f,
	0x82, 0x4e, 0x6e, 0x6f, 0x7a, 0x73, 0xdb, 0x50, 0x74, 0xe3, 0x82, 0xdb, 0x3c, 0xdd, 0xcf, 0x26,
	0x01, 0xc8, 0x09, 0x81, 0x3b, 0xa5, 0x94, 0xd4, 0xaa, 0xf6, 0x76, 0xf1, 0x4b, 0x14, 0x16, 0x4b,
	0x11, 0xdd, 0x2c, 0x45, 0x74, 0xbb, 0x14, 0xd1, 0xcf, 0xa5, 0x88, 0xbe, 0xae, 0x44, 0xe1, 0x66,
	0x25, 0x0a, 0xb7, 0x2b, 0x51, 0x78, 0x7f, 0xe6, 0xf9, 0xd1, 0x68, 0x3e, 0x94, 0x5d, 0x76, 0xa9,
	0xac, 0x1d, 0x3f, 0x0c, 0x37, 0xb5, 0x32, 0x19, 0x7b, 0xca, 0xce, 0x9f, 0xfc, 0x27, 0x00, 0x00,
	0xff, 0xff, 0x9c, 0x3d, 0x31, 0xb1, 0xdd, 0x03, 0x00, 0x00,
}

func (this *FunctionVolatility) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FunctionVolatility)
	if !ok {
		that2, ok := that.(FunctionVolatility)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Volatility != that1.Volatility {
		return false
	}
	return true
}
func (this *FunctionNullInputBehavior) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FunctionNullInputBehavior)
	if !ok {
		that2, ok := that.(FunctionNullInputBehavior)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.NullInputBehavior != that1.NullInputBehavior {
		return false
	}
	return true
}
func (this *FunctionLanguage) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FunctionLanguage)
	if !ok {
		that2, ok := that.(FunctionLanguage)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Lang != that1.Lang {
		return false
	}
	return true
}
func (this *FunctionParamClass) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FunctionParamClass)
	if !ok {
		that2, ok := that.(FunctionParamClass)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Class != that1.Class {
		return false
	}
	return true
}
func (m *Function) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Function) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Function) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *Function_Param) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Function_Param) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Function_Param) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *FunctionVolatility) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FunctionVolatility) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FunctionVolatility) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	i = encodeVarintFunction(dAtA, i, uint64(m.Volatility))
	i--
	dAtA[i] = 0x8
	return len(dAtA) - i, nil
}

func (m *FunctionNullInputBehavior) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FunctionNullInputBehavior) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FunctionNullInputBehavior) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	i = encodeVarintFunction(dAtA, i, uint64(m.NullInputBehavior))
	i--
	dAtA[i] = 0x8
	return len(dAtA) - i, nil
}

func (m *FunctionLanguage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FunctionLanguage) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FunctionLanguage) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	i = encodeVarintFunction(dAtA, i, uint64(m.Lang))
	i--
	dAtA[i] = 0x8
	return len(dAtA) - i, nil
}

func (m *FunctionParamClass) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FunctionParamClass) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FunctionParamClass) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	i = encodeVarintFunction(dAtA, i, uint64(m.Class))
	i--
	dAtA[i] = 0x8
	return len(dAtA) - i, nil
}

func encodeVarintFunction(dAtA []byte, offset int, v uint64) int {
	offset -= sovFunction(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Function) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *Function_Param) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *FunctionVolatility) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovFunction(uint64(m.Volatility))
	return n
}

func (m *FunctionNullInputBehavior) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovFunction(uint64(m.NullInputBehavior))
	return n
}

func (m *FunctionLanguage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovFunction(uint64(m.Lang))
	return n
}

func (m *FunctionParamClass) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovFunction(uint64(m.Class))
	return n
}

func sovFunction(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFunction(x uint64) (n int) {
	return sovFunction(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Function) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFunction
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
			return fmt.Errorf("proto: Function: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Function: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipFunction(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFunction
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
func (m *Function_Param) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFunction
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
			return fmt.Errorf("proto: Param: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Param: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipFunction(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFunction
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
func (m *FunctionVolatility) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFunction
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
			return fmt.Errorf("proto: FunctionVolatility: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FunctionVolatility: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Volatility", wireType)
			}
			m.Volatility = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFunction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Volatility |= Function_Volatility(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipFunction(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFunction
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
func (m *FunctionNullInputBehavior) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFunction
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
			return fmt.Errorf("proto: FunctionNullInputBehavior: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FunctionNullInputBehavior: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NullInputBehavior", wireType)
			}
			m.NullInputBehavior = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFunction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NullInputBehavior |= Function_NullInputBehavior(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipFunction(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFunction
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
func (m *FunctionLanguage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFunction
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
			return fmt.Errorf("proto: FunctionLanguage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FunctionLanguage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lang", wireType)
			}
			m.Lang = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFunction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Lang |= Function_Language(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipFunction(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFunction
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
func (m *FunctionParamClass) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFunction
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
			return fmt.Errorf("proto: FunctionParamClass: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FunctionParamClass: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Class", wireType)
			}
			m.Class = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFunction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Class |= Function_Param_Class(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipFunction(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFunction
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
func skipFunction(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFunction
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
					return 0, ErrIntOverflowFunction
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
					return 0, ErrIntOverflowFunction
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
				return 0, ErrInvalidLengthFunction
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupFunction
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthFunction
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthFunction        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFunction          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupFunction = fmt.Errorf("proto: unexpected end of group")
)

