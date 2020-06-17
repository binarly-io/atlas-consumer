// Code generated by protoc-gen-go. DO NOT EDIT.
// source: type_command.proto

package atlas

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type CommandType int32

const (
	// Due to first enum value has to be zero in proto3
	CommandType_COMMAND_RESERVED         CommandType = 0
	CommandType_COMMAND_UNSPECIFIED      CommandType = 100
	CommandType_COMMAND_ECHO_REQUEST     CommandType = 200
	CommandType_COMMAND_ECHO_REPLY       CommandType = 300
	CommandType_COMMAND_CONFIG_REQUEST   CommandType = 400
	CommandType_COMMAND_CONFIG           CommandType = 500
	CommandType_COMMAND_METRICS_SCHEDULE CommandType = 600
	CommandType_COMMAND_METRICS_REQUEST  CommandType = 700
	CommandType_COMMAND_METRICS          CommandType = 800
	CommandType_COMMAND_DATA_SCHEDULE    CommandType = 900
	CommandType_COMMAND_DATA_REQUEST     CommandType = 1000
	CommandType_COMMAND_DATA             CommandType = 1100
)

var CommandType_name = map[int32]string{
	0:    "COMMAND_RESERVED",
	100:  "COMMAND_UNSPECIFIED",
	200:  "COMMAND_ECHO_REQUEST",
	300:  "COMMAND_ECHO_REPLY",
	400:  "COMMAND_CONFIG_REQUEST",
	500:  "COMMAND_CONFIG",
	600:  "COMMAND_METRICS_SCHEDULE",
	700:  "COMMAND_METRICS_REQUEST",
	800:  "COMMAND_METRICS",
	900:  "COMMAND_DATA_SCHEDULE",
	1000: "COMMAND_DATA_REQUEST",
	1100: "COMMAND_DATA",
}

var CommandType_value = map[string]int32{
	"COMMAND_RESERVED":         0,
	"COMMAND_UNSPECIFIED":      100,
	"COMMAND_ECHO_REQUEST":     200,
	"COMMAND_ECHO_REPLY":       300,
	"COMMAND_CONFIG_REQUEST":   400,
	"COMMAND_CONFIG":           500,
	"COMMAND_METRICS_SCHEDULE": 600,
	"COMMAND_METRICS_REQUEST":  700,
	"COMMAND_METRICS":          800,
	"COMMAND_DATA_SCHEDULE":    900,
	"COMMAND_DATA_REQUEST":     1000,
	"COMMAND_DATA":             1100,
}

func (x CommandType) String() string {
	return proto.EnumName(CommandType_name, int32(x))
}

func (CommandType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e84d1630985fc0a4, []int{0}
}

type Command struct {
	Header               *Metadata `protobuf:"bytes,100,opt,name=header,proto3" json:"header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Command) Reset()         { *m = Command{} }
func (m *Command) String() string { return proto.CompactTextString(m) }
func (*Command) ProtoMessage()    {}
func (*Command) Descriptor() ([]byte, []int) {
	return fileDescriptor_e84d1630985fc0a4, []int{0}
}

func (m *Command) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Command.Unmarshal(m, b)
}
func (m *Command) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Command.Marshal(b, m, deterministic)
}
func (m *Command) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Command.Merge(m, src)
}
func (m *Command) XXX_Size() int {
	return xxx_messageInfo_Command.Size(m)
}
func (m *Command) XXX_DiscardUnknown() {
	xxx_messageInfo_Command.DiscardUnknown(m)
}

var xxx_messageInfo_Command proto.InternalMessageInfo

func (m *Command) GetHeader() *Metadata {
	if m != nil {
		return m.Header
	}
	return nil
}

func init() {
	proto.RegisterEnum("atlas.CommandType", CommandType_name, CommandType_value)
	proto.RegisterType((*Command)(nil), "atlas.Command")
}

func init() {
	proto.RegisterFile("type_command.proto", fileDescriptor_e84d1630985fc0a4)
}

var fileDescriptor_e84d1630985fc0a4 = []byte{
	// 291 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0xa9, 0x2c, 0x48,
	0x8d, 0x4f, 0xce, 0xcf, 0xcd, 0x4d, 0xcc, 0x4b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62,
	0x4d, 0x2c, 0xc9, 0x49, 0x2c, 0x96, 0x12, 0x06, 0x4b, 0xe5, 0xa6, 0x96, 0x24, 0xa6, 0x24, 0x96,
	0x24, 0x42, 0xe4, 0x94, 0x8c, 0xb8, 0xd8, 0x9d, 0x21, 0x8a, 0x85, 0xd4, 0xb9, 0xd8, 0x32, 0x52,
	0x13, 0x53, 0x52, 0x8b, 0x24, 0x52, 0x14, 0x18, 0x35, 0xb8, 0x8d, 0xf8, 0xf5, 0xc0, 0xfa, 0xf4,
	0x7c, 0xa1, 0x3a, 0x82, 0xa0, 0xd2, 0x5a, 0xfb, 0x99, 0xb8, 0xb8, 0xa1, 0x9a, 0x42, 0x2a, 0x0b,
	0x52, 0x85, 0x44, 0xb8, 0x04, 0x9c, 0xfd, 0x7d, 0x7d, 0x1d, 0xfd, 0x5c, 0xe2, 0x83, 0x5c, 0x83,
	0x5d, 0x83, 0xc2, 0x5c, 0x5d, 0x04, 0x18, 0x84, 0xc4, 0xb9, 0x84, 0x61, 0xa2, 0xa1, 0x7e, 0xc1,
	0x01, 0xae, 0xce, 0x9e, 0x6e, 0x9e, 0xae, 0x2e, 0x02, 0x29, 0x42, 0x92, 0x5c, 0x22, 0x30, 0x09,
	0x57, 0x67, 0x0f, 0xff, 0xf8, 0x20, 0xd7, 0xc0, 0x50, 0xd7, 0xe0, 0x10, 0x81, 0x13, 0x8c, 0x42,
	0xe2, 0x5c, 0x42, 0x68, 0x52, 0x01, 0x3e, 0x91, 0x02, 0x6b, 0x98, 0x84, 0xa4, 0xb9, 0xc4, 0x60,
	0x12, 0xce, 0xfe, 0x7e, 0x6e, 0x9e, 0xee, 0x70, 0x5d, 0x13, 0x98, 0x85, 0x84, 0xb9, 0xf8, 0x50,
	0x25, 0x05, 0xbe, 0x30, 0x0b, 0xc9, 0x72, 0x49, 0xc0, 0x04, 0x7d, 0x5d, 0x43, 0x82, 0x3c, 0x9d,
	0x83, 0xe3, 0x83, 0x9d, 0x3d, 0x5c, 0x5d, 0x42, 0x7d, 0x5c, 0x05, 0x6e, 0xb0, 0x08, 0xc9, 0x70,
	0x89, 0xa3, 0x4b, 0xc3, 0x4c, 0xdc, 0xc3, 0x2a, 0x24, 0xc2, 0xc5, 0x8f, 0x26, 0x2b, 0xb0, 0x80,
	0x4d, 0x48, 0x8a, 0x4b, 0x14, 0x26, 0xea, 0xe2, 0x18, 0xe2, 0x88, 0x30, 0xaf, 0x85, 0x1d, 0xd9,
	0x53, 0x60, 0x39, 0x98, 0x61, 0x2f, 0xd8, 0x85, 0x04, 0xb9, 0x78, 0x90, 0xa5, 0x04, 0xce, 0x70,
	0x24, 0xb1, 0x81, 0x03, 0xdf, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xc8, 0xb8, 0x26, 0xd1, 0xae,
	0x01, 0x00, 0x00,
}
