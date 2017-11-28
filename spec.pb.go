// Code generated by protoc-gen-go. DO NOT EDIT.
// source: spec.proto

/*
Package prefecturejp is a generated protocol buffer package.

It is generated from these files:
	spec.proto

It has these top-level messages:
*/
package prefecturejp

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PrefectureField int32

const (
	PrefectureField_CODE       PrefectureField = 0
	PrefectureField_NAME       PrefectureField = 1
	PrefectureField_JAPANESE   PrefectureField = 2
	PrefectureField_SHORT_NAME PrefectureField = 3
	PrefectureField_REGION     PrefectureField = 4
)

var PrefectureField_name = map[int32]string{
	0: "CODE",
	1: "NAME",
	2: "JAPANESE",
	3: "SHORT_NAME",
	4: "REGION",
}
var PrefectureField_value = map[string]int32{
	"CODE":       0,
	"NAME":       1,
	"JAPANESE":   2,
	"SHORT_NAME": 3,
	"REGION":     4,
}

func (x PrefectureField) String() string {
	return proto.EnumName(PrefectureField_name, int32(x))
}
func (PrefectureField) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterEnum("prefecturejp.PrefectureField", PrefectureField_name, PrefectureField_value)
}

func init() { proto.RegisterFile("spec.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 124 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x2e, 0x48, 0x4d,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x29, 0x28, 0x4a, 0x4d, 0x4b, 0x4d, 0x2e, 0x29,
	0x2d, 0x4a, 0xcd, 0x2a, 0xd0, 0xf2, 0xe7, 0xe2, 0x0f, 0x80, 0xf3, 0xdd, 0x32, 0x53, 0x73, 0x52,
	0x84, 0x38, 0xb8, 0x58, 0x9c, 0xfd, 0x5d, 0x5c, 0x05, 0x18, 0x40, 0x2c, 0x3f, 0x47, 0x5f, 0x57,
	0x01, 0x46, 0x21, 0x1e, 0x2e, 0x0e, 0x2f, 0xc7, 0x00, 0x47, 0x3f, 0xd7, 0x60, 0x57, 0x01, 0x26,
	0x21, 0x3e, 0x2e, 0xae, 0x60, 0x0f, 0xff, 0xa0, 0x90, 0x78, 0xb0, 0x2c, 0xb3, 0x10, 0x17, 0x17,
	0x5b, 0x90, 0xab, 0xbb, 0xa7, 0xbf, 0x9f, 0x00, 0x4b, 0x12, 0x1b, 0xd8, 0x16, 0x63, 0x40, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x2c, 0x81, 0xb4, 0x3f, 0x73, 0x00, 0x00, 0x00,
}
