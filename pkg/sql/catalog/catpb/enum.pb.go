// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sql/catalog/catpb/enum.proto

package catpb

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	math "math"
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

// SystemColumnKind is an enum representing the different kind of system
// columns that can be synthesized by the execution engine.
type SystemColumnKind int32

const (
	// Default value, unused.
	SystemColumnKind_NONE SystemColumnKind = 0
	// A system column containing the value of the MVCC timestamp associated
	// with the kv's corresponding to the row.
	SystemColumnKind_MVCCTIMESTAMP SystemColumnKind = 1
	// A system column containing the OID of the table that the row came from.
	SystemColumnKind_TABLEOID SystemColumnKind = 2
)

var SystemColumnKind_name = map[int32]string{
	0: "NONE",
	1: "MVCCTIMESTAMP",
	2: "TABLEOID",
}

var SystemColumnKind_value = map[string]int32{
	"NONE":          0,
	"MVCCTIMESTAMP": 1,
	"TABLEOID":      2,
}

func (x SystemColumnKind) String() string {
	return proto.EnumName(SystemColumnKind_name, int32(x))
}

func (SystemColumnKind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_31fce16dbfe42fb4, []int{0}
}

// InvertedIndexColumnKind is the kind of the inverted index on a column. The
// reason this needs to be stored is that we need to be able to check that the
// "opclass" passed into an inverted index declaration (for example,
// gin_trgm_ops) is compatible with the datatype of a particular column
// (gin_tgrm_ops is only valid on text). A future reason is that it's possible
// to desire having more than one type of inverted index on a particular
// datatype - for example, you might want to create a "stemming" inverted index
// on text. And without this extra kind, it wouldn't be possible to distinguish
// a text inverted index that uses trigrams, vs a text inverted index that uses
// stemming.
type InvertedIndexColumnKind int32

const (
	// DEFAULT is the default kind of inverted index column. JSON, Array, and
	// geo inverted indexes all are DEFAULT, though prior to 22.2 they had no
	// kind at all.
	InvertedIndexColumnKind_DEFAULT InvertedIndexColumnKind = 0
	// TRIGRAM is the trigram kind of inverted index column. It's only valid on
	// text columns.
	InvertedIndexColumnKind_TRIGRAM InvertedIndexColumnKind = 1
)

var InvertedIndexColumnKind_name = map[int32]string{
	0: "DEFAULT",
	1: "TRIGRAM",
}

var InvertedIndexColumnKind_value = map[string]int32{
	"DEFAULT": 0,
	"TRIGRAM": 1,
}

func (x InvertedIndexColumnKind) String() string {
	return proto.EnumName(InvertedIndexColumnKind_name, int32(x))
}

func (InvertedIndexColumnKind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_31fce16dbfe42fb4, []int{1}
}

func init() {
	proto.RegisterEnum("cockroach.sql.catalog.catpb.SystemColumnKind", SystemColumnKind_name, SystemColumnKind_value)
	proto.RegisterEnum("cockroach.sql.catalog.catpb.InvertedIndexColumnKind", InvertedIndexColumnKind_name, InvertedIndexColumnKind_value)
}

func init() { proto.RegisterFile("sql/catalog/catpb/enum.proto", fileDescriptor_31fce16dbfe42fb4) }

var fileDescriptor_31fce16dbfe42fb4 = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x29, 0x2e, 0xcc, 0xd1,
	0x4f, 0x4e, 0x2c, 0x49, 0xcc, 0xc9, 0x4f, 0x07, 0xd1, 0x05, 0x49, 0xfa, 0xa9, 0x79, 0xa5, 0xb9,
	0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xd2, 0xc9, 0xf9, 0xc9, 0xd9, 0x45, 0xf9, 0x89, 0xc9,
	0x19, 0x7a, 0xc5, 0x85, 0x39, 0x7a, 0x50, 0x75, 0x7a, 0x60, 0x75, 0x52, 0x22, 0xe9, 0xf9, 0xe9,
	0xf9, 0x60, 0x75, 0xfa, 0x20, 0x16, 0x44, 0x8b, 0x96, 0x2d, 0x97, 0x40, 0x70, 0x65, 0x71, 0x49,
	0x6a, 0xae, 0x73, 0x7e, 0x4e, 0x69, 0x6e, 0x9e, 0x77, 0x66, 0x5e, 0x8a, 0x10, 0x07, 0x17, 0x8b,
	0x9f, 0xbf, 0x9f, 0xab, 0x00, 0x83, 0x90, 0x20, 0x17, 0xaf, 0x6f, 0x98, 0xb3, 0x73, 0x88, 0xa7,
	0xaf, 0x6b, 0x70, 0x88, 0xa3, 0x6f, 0x80, 0x00, 0xa3, 0x10, 0x0f, 0x17, 0x47, 0x88, 0xa3, 0x93,
	0x8f, 0xab, 0xbf, 0xa7, 0x8b, 0x00, 0x93, 0x96, 0x31, 0x97, 0xb8, 0x67, 0x5e, 0x59, 0x6a, 0x51,
	0x49, 0x6a, 0x8a, 0x67, 0x5e, 0x4a, 0x6a, 0x05, 0x92, 0x29, 0xdc, 0x5c, 0xec, 0x2e, 0xae, 0x6e,
	0x8e, 0xa1, 0x3e, 0x21, 0x02, 0x0c, 0x20, 0x4e, 0x48, 0x90, 0xa7, 0x7b, 0x90, 0xa3, 0xaf, 0x00,
	0xa3, 0x53, 0xc4, 0x89, 0x87, 0x72, 0x0c, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0x78,
	0xe3, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7,
	0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x59, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e,
	0xae, 0x3e, 0xdc, 0x4f, 0x29, 0x49, 0x08, 0xb6, 0x7e, 0x41, 0x76, 0xba, 0x3e, 0x46, 0x58, 0x24,
	0xb1, 0x81, 0x3d, 0x65, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xdf, 0x81, 0x7d, 0x36, 0x27, 0x01,
	0x00, 0x00,
}

