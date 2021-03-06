// Code generated by protoc-gen-go. DO NOT EDIT.
// source: type_metric.proto

package atlas

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type MetricType int32

const (
	MetricType_METRIC_TYPE_RESERVED    MetricType = 0
	MetricType_METRIC_TYPE_UNSPECIFIED MetricType = 100
	MetricType_METRIC_TYPE_CPU         MetricType = 200
	MetricType_METRIC_TYPE_RAM         MetricType = 300
)

var MetricType_name = map[int32]string{
	0:   "METRIC_TYPE_RESERVED",
	100: "METRIC_TYPE_UNSPECIFIED",
	200: "METRIC_TYPE_CPU",
	300: "METRIC_TYPE_RAM",
}

var MetricType_value = map[string]int32{
	"METRIC_TYPE_RESERVED":    0,
	"METRIC_TYPE_UNSPECIFIED": 100,
	"METRIC_TYPE_CPU":         200,
	"METRIC_TYPE_RAM":         300,
}

func (x MetricType) String() string {
	return proto.EnumName(MetricType_name, int32(x))
}

func (MetricType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fa4a9edf946d8607, []int{0}
}

type MetricsType int32

const (
	// Due to first enum value has to be zero in proto3
	MetricsType_METRICS_TYPE_RESERVED    MetricsType = 0
	MetricsType_METRICS_TYPE_UNSPECIFIED MetricsType = 100
	// Resource utilization, such as CPU, RAM, etc
	MetricsType_METRICS_TYPE_RESOURCE_UTILIZATION MetricsType = 200
)

var MetricsType_name = map[int32]string{
	0:   "METRICS_TYPE_RESERVED",
	100: "METRICS_TYPE_UNSPECIFIED",
	200: "METRICS_TYPE_RESOURCE_UTILIZATION",
}

var MetricsType_value = map[string]int32{
	"METRICS_TYPE_RESERVED":             0,
	"METRICS_TYPE_UNSPECIFIED":          100,
	"METRICS_TYPE_RESOURCE_UTILIZATION": 200,
}

func (x MetricsType) String() string {
	return proto.EnumName(MetricsType_name, int32(x))
}

func (MetricsType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fa4a9edf946d8607, []int{1}
}

// Metric is one metric tuple (metric identifier [optional], timestamp [optional], description [optional], metric value)
type Metric struct {
	// Types that are valid to be assigned to TypeOptional:
	//	*Metric_Type
	//	*Metric_Name
	TypeOptional isMetric_TypeOptional `protobuf_oneof:"type_optional"`
	// Types that are valid to be assigned to TimestampOptional:
	//	*Metric_Ts
	TimestampOptional isMetric_TimestampOptional `protobuf_oneof:"timestamp_optional"`
	// Types that are valid to be assigned to DescriptionOptional:
	//	*Metric_Description
	DescriptionOptional isMetric_DescriptionOptional `protobuf_oneof:"description_optional"`
	// Types that are valid to be assigned to Value:
	//	*Metric_StringValue
	//	*Metric_DoubleValue
	//	*Metric_Int32Value
	//	*Metric_Uint32Value
	//	*Metric_Int64Value
	//	*Metric_Uint64Value
	//	*Metric_BytesValue
	Value                isMetric_Value `protobuf_oneof:"value"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Metric) Reset()         { *m = Metric{} }
func (m *Metric) String() string { return proto.CompactTextString(m) }
func (*Metric) ProtoMessage()    {}
func (*Metric) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa4a9edf946d8607, []int{0}
}

func (m *Metric) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metric.Unmarshal(m, b)
}
func (m *Metric) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metric.Marshal(b, m, deterministic)
}
func (m *Metric) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metric.Merge(m, src)
}
func (m *Metric) XXX_Size() int {
	return xxx_messageInfo_Metric.Size(m)
}
func (m *Metric) XXX_DiscardUnknown() {
	xxx_messageInfo_Metric.DiscardUnknown(m)
}

var xxx_messageInfo_Metric proto.InternalMessageInfo

type isMetric_TypeOptional interface {
	isMetric_TypeOptional()
}

type Metric_Type struct {
	Type MetricType `protobuf:"varint,100,opt,name=type,proto3,enum=atlas.MetricType,oneof"`
}

type Metric_Name struct {
	Name string `protobuf:"bytes,110,opt,name=name,proto3,oneof"`
}

func (*Metric_Type) isMetric_TypeOptional() {}

func (*Metric_Name) isMetric_TypeOptional() {}

func (m *Metric) GetTypeOptional() isMetric_TypeOptional {
	if m != nil {
		return m.TypeOptional
	}
	return nil
}

func (m *Metric) GetType() MetricType {
	if x, ok := m.GetTypeOptional().(*Metric_Type); ok {
		return x.Type
	}
	return MetricType_METRIC_TYPE_RESERVED
}

func (m *Metric) GetName() string {
	if x, ok := m.GetTypeOptional().(*Metric_Name); ok {
		return x.Name
	}
	return ""
}

type isMetric_TimestampOptional interface {
	isMetric_TimestampOptional()
}

type Metric_Ts struct {
	Ts *timestamp.Timestamp `protobuf:"bytes,200,opt,name=ts,proto3,oneof"`
}

func (*Metric_Ts) isMetric_TimestampOptional() {}

func (m *Metric) GetTimestampOptional() isMetric_TimestampOptional {
	if m != nil {
		return m.TimestampOptional
	}
	return nil
}

func (m *Metric) GetTs() *timestamp.Timestamp {
	if x, ok := m.GetTimestampOptional().(*Metric_Ts); ok {
		return x.Ts
	}
	return nil
}

type isMetric_DescriptionOptional interface {
	isMetric_DescriptionOptional()
}

type Metric_Description struct {
	Description string `protobuf:"bytes,300,opt,name=description,proto3,oneof"`
}

func (*Metric_Description) isMetric_DescriptionOptional() {}

func (m *Metric) GetDescriptionOptional() isMetric_DescriptionOptional {
	if m != nil {
		return m.DescriptionOptional
	}
	return nil
}

func (m *Metric) GetDescription() string {
	if x, ok := m.GetDescriptionOptional().(*Metric_Description); ok {
		return x.Description
	}
	return ""
}

type isMetric_Value interface {
	isMetric_Value()
}

type Metric_StringValue struct {
	StringValue string `protobuf:"bytes,400,opt,name=string_value,json=stringValue,proto3,oneof"`
}

type Metric_DoubleValue struct {
	DoubleValue float64 `protobuf:"fixed64,410,opt,name=double_value,json=doubleValue,proto3,oneof"`
}

type Metric_Int32Value struct {
	Int32Value int32 `protobuf:"varint,420,opt,name=int32_value,json=int32Value,proto3,oneof"`
}

type Metric_Uint32Value struct {
	Uint32Value uint32 `protobuf:"varint,430,opt,name=uint32_value,json=uint32Value,proto3,oneof"`
}

type Metric_Int64Value struct {
	Int64Value int64 `protobuf:"varint,440,opt,name=int64_value,json=int64Value,proto3,oneof"`
}

type Metric_Uint64Value struct {
	Uint64Value uint64 `protobuf:"varint,450,opt,name=uint64_value,json=uint64Value,proto3,oneof"`
}

type Metric_BytesValue struct {
	BytesValue []byte `protobuf:"bytes,460,opt,name=bytes_value,json=bytesValue,proto3,oneof"`
}

func (*Metric_StringValue) isMetric_Value() {}

func (*Metric_DoubleValue) isMetric_Value() {}

func (*Metric_Int32Value) isMetric_Value() {}

func (*Metric_Uint32Value) isMetric_Value() {}

func (*Metric_Int64Value) isMetric_Value() {}

func (*Metric_Uint64Value) isMetric_Value() {}

func (*Metric_BytesValue) isMetric_Value() {}

func (m *Metric) GetValue() isMetric_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Metric) GetStringValue() string {
	if x, ok := m.GetValue().(*Metric_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (m *Metric) GetDoubleValue() float64 {
	if x, ok := m.GetValue().(*Metric_DoubleValue); ok {
		return x.DoubleValue
	}
	return 0
}

func (m *Metric) GetInt32Value() int32 {
	if x, ok := m.GetValue().(*Metric_Int32Value); ok {
		return x.Int32Value
	}
	return 0
}

func (m *Metric) GetUint32Value() uint32 {
	if x, ok := m.GetValue().(*Metric_Uint32Value); ok {
		return x.Uint32Value
	}
	return 0
}

func (m *Metric) GetInt64Value() int64 {
	if x, ok := m.GetValue().(*Metric_Int64Value); ok {
		return x.Int64Value
	}
	return 0
}

func (m *Metric) GetUint64Value() uint64 {
	if x, ok := m.GetValue().(*Metric_Uint64Value); ok {
		return x.Uint64Value
	}
	return 0
}

func (m *Metric) GetBytesValue() []byte {
	if x, ok := m.GetValue().(*Metric_BytesValue); ok {
		return x.BytesValue
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Metric) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Metric_Type)(nil),
		(*Metric_Name)(nil),
		(*Metric_Ts)(nil),
		(*Metric_Description)(nil),
		(*Metric_StringValue)(nil),
		(*Metric_DoubleValue)(nil),
		(*Metric_Int32Value)(nil),
		(*Metric_Uint32Value)(nil),
		(*Metric_Int64Value)(nil),
		(*Metric_Uint64Value)(nil),
		(*Metric_BytesValue)(nil),
	}
}

// Metrics is a set of Metric tuples
type Metrics struct {
	Header *Metadata `protobuf:"bytes,100,opt,name=header,proto3" json:"header,omitempty"`
	// MetricType can be specified once for the whole set of metrics, instead of specifying in each one of them
	//
	// Types that are valid to be assigned to TypeOptional:
	//	*Metrics_Type
	TypeOptional isMetrics_TypeOptional `protobuf_oneof:"type_optional"`
	// Types that are valid to be assigned to NameOptional:
	//	*Metrics_Name
	NameOptional isMetrics_NameOptional `protobuf_oneof:"name_optional"`
	// Metrics is the purpose of the whole Metrics data type, is expected to be present all the time
	Metrics              []*Metric `protobuf:"bytes,400,rep,name=metrics,proto3" json:"metrics,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Metrics) Reset()         { *m = Metrics{} }
func (m *Metrics) String() string { return proto.CompactTextString(m) }
func (*Metrics) ProtoMessage()    {}
func (*Metrics) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa4a9edf946d8607, []int{1}
}

func (m *Metrics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metrics.Unmarshal(m, b)
}
func (m *Metrics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metrics.Marshal(b, m, deterministic)
}
func (m *Metrics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metrics.Merge(m, src)
}
func (m *Metrics) XXX_Size() int {
	return xxx_messageInfo_Metrics.Size(m)
}
func (m *Metrics) XXX_DiscardUnknown() {
	xxx_messageInfo_Metrics.DiscardUnknown(m)
}

var xxx_messageInfo_Metrics proto.InternalMessageInfo

func (m *Metrics) GetHeader() *Metadata {
	if m != nil {
		return m.Header
	}
	return nil
}

type isMetrics_TypeOptional interface {
	isMetrics_TypeOptional()
}

type Metrics_Type struct {
	Type int32 `protobuf:"varint,200,opt,name=type,proto3,oneof"`
}

func (*Metrics_Type) isMetrics_TypeOptional() {}

func (m *Metrics) GetTypeOptional() isMetrics_TypeOptional {
	if m != nil {
		return m.TypeOptional
	}
	return nil
}

func (m *Metrics) GetType() int32 {
	if x, ok := m.GetTypeOptional().(*Metrics_Type); ok {
		return x.Type
	}
	return 0
}

type isMetrics_NameOptional interface {
	isMetrics_NameOptional()
}

type Metrics_Name struct {
	Name string `protobuf:"bytes,300,opt,name=name,proto3,oneof"`
}

func (*Metrics_Name) isMetrics_NameOptional() {}

func (m *Metrics) GetNameOptional() isMetrics_NameOptional {
	if m != nil {
		return m.NameOptional
	}
	return nil
}

func (m *Metrics) GetName() string {
	if x, ok := m.GetNameOptional().(*Metrics_Name); ok {
		return x.Name
	}
	return ""
}

func (m *Metrics) GetMetrics() []*Metric {
	if m != nil {
		return m.Metrics
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Metrics) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Metrics_Type)(nil),
		(*Metrics_Name)(nil),
	}
}

func init() {
	proto.RegisterEnum("atlas.MetricType", MetricType_name, MetricType_value)
	proto.RegisterEnum("atlas.MetricsType", MetricsType_name, MetricsType_value)
	proto.RegisterType((*Metric)(nil), "atlas.Metric")
	proto.RegisterType((*Metrics)(nil), "atlas.Metrics")
}

func init() {
	proto.RegisterFile("type_metric.proto", fileDescriptor_fa4a9edf946d8607)
}

var fileDescriptor_fa4a9edf946d8607 = []byte{
	// 534 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x93, 0x4f, 0x6f, 0x12, 0x41,
	0x18, 0xc6, 0x3b, 0x6c, 0x81, 0xf8, 0x6e, 0x11, 0x3a, 0x82, 0xae, 0x68, 0x22, 0xa2, 0xb1, 0xa4,
	0x89, 0xdb, 0x84, 0x36, 0xbd, 0xb7, 0x74, 0x0d, 0x9b, 0x48, 0x4b, 0x86, 0xa5, 0x89, 0x5e, 0xc8,
	0xd0, 0x1d, 0x71, 0x13, 0xd8, 0x25, 0xec, 0x60, 0xc2, 0xb7, 0xf0, 0xec, 0xd9, 0x23, 0xf1, 0xec,
	0xd9, 0x13, 0x07, 0xbf, 0x87, 0x5f, 0xc3, 0xcc, 0x9f, 0x5d, 0x96, 0xf6, 0xc8, 0xf3, 0xfe, 0xe6,
	0x99, 0xe1, 0x7d, 0x9e, 0x85, 0x43, 0xbe, 0x9a, 0xb3, 0xd1, 0x8c, 0xf1, 0x45, 0x70, 0x67, 0xcf,
	0x17, 0x11, 0x8f, 0x70, 0x9e, 0xf2, 0x29, 0x8d, 0xeb, 0xaf, 0x26, 0x51, 0x34, 0x99, 0xb2, 0x13,
	0x29, 0x8e, 0x97, 0x5f, 0x4e, 0x78, 0x30, 0x63, 0x31, 0xa7, 0xb3, 0xb9, 0xe2, 0xea, 0x4f, 0x92,
	0xa3, 0xd4, 0xa7, 0x9c, 0x2a, 0xb1, 0xf9, 0xcf, 0x80, 0x42, 0x4f, 0xba, 0xe1, 0x23, 0xd8, 0x17,
	0x84, 0xe5, 0x37, 0x50, 0xeb, 0x71, 0xfb, 0xd0, 0x96, 0xb6, 0xb6, 0x1a, 0x7a, 0xab, 0x39, 0xeb,
	0xee, 0x11, 0x09, 0xe0, 0x2a, 0xec, 0x87, 0x74, 0xc6, 0xac, 0xb0, 0x81, 0x5a, 0x8f, 0x84, 0x2a,
	0x7e, 0xe1, 0xf7, 0x90, 0xe3, 0xb1, 0xb5, 0x41, 0x0d, 0xd4, 0x32, 0xdb, 0x75, 0x5b, 0xbd, 0xc6,
	0x4e, 0x5e, 0x63, 0x7b, 0xc9, 0x6b, 0xba, 0x88, 0xe4, 0x78, 0x8c, 0xdf, 0x80, 0xe9, 0xb3, 0xf8,
	0x6e, 0x11, 0xcc, 0x79, 0x10, 0x85, 0xd6, 0x3a, 0x27, 0xcd, 0x72, 0x24, 0xab, 0xe2, 0xb7, 0x70,
	0x10, 0xf3, 0x45, 0x10, 0x4e, 0x46, 0xdf, 0xe8, 0x74, 0xc9, 0xac, 0xef, 0x86, 0xa4, 0x0c, 0x62,
	0x2a, 0xf9, 0x56, 0xa8, 0x82, 0xf2, 0xa3, 0xe5, 0x78, 0xca, 0x34, 0xf5, 0x43, 0x50, 0x48, 0x50,
	0x4a, 0x56, 0x54, 0x13, 0xcc, 0x20, 0xe4, 0xa7, 0x6d, 0x0d, 0xfd, 0x14, 0x50, 0xbe, 0x6b, 0x10,
	0x90, 0x6a, 0xea, 0xb4, 0xcc, 0x42, 0xbf, 0x04, 0x54, 0x12, 0x4e, 0xcb, 0x0c, 0xa5, 0x9c, 0xce,
	0xcf, 0x34, 0xf4, 0x5b, 0x40, 0x86, 0x76, 0x3a, 0x3f, 0xdb, 0x71, 0x4a, 0xa1, 0x3f, 0x02, 0xda,
	0x4f, 0x9c, 0x12, 0xaa, 0x09, 0xe6, 0x78, 0xc5, 0x59, 0xac, 0xa1, 0xbf, 0x02, 0x3a, 0x10, 0x4e,
	0x52, 0x95, 0xcc, 0x65, 0x19, 0x4a, 0x32, 0xb8, 0x48, 0xae, 0x84, 0x4e, 0x2f, 0xab, 0x80, 0xd3,
	0x68, 0xb7, 0xea, 0x53, 0xa8, 0x66, 0x36, 0xb7, 0xd5, 0x8b, 0x90, 0x97, 0xe6, 0xcd, 0x35, 0x82,
	0xa2, 0x0a, 0x33, 0xc6, 0x47, 0x50, 0xf8, 0xca, 0xa8, 0xcf, 0x16, 0x32, 0x6c, 0xb3, 0x5d, 0xde,
	0x86, 0x2d, 0xcb, 0x41, 0xf4, 0x18, 0xd7, 0x74, 0x27, 0x64, 0xac, 0xf9, 0xb4, 0x01, 0x35, 0xdd,
	0x00, 0x9d, 0x1a, 0xd2, 0x15, 0x68, 0x41, 0x51, 0x35, 0x33, 0x16, 0x49, 0x19, 0x2d, 0xb3, 0x5d,
	0xda, 0x69, 0x11, 0x49, 0xc6, 0x0f, 0xff, 0x54, 0x19, 0x4a, 0xc2, 0x22, 0x15, 0x8e, 0x23, 0x80,
	0x6d, 0xf5, 0xb0, 0x05, 0xd5, 0x9e, 0xe3, 0x11, 0xb7, 0x33, 0xf2, 0x3e, 0xf5, 0x9d, 0x11, 0x71,
	0x06, 0x0e, 0xb9, 0x75, 0xae, 0x2a, 0x7b, 0xf8, 0x05, 0x3c, 0xcb, 0x4e, 0x86, 0xd7, 0x83, 0xbe,
	0xd3, 0x71, 0x3f, 0xb8, 0xce, 0x55, 0xc5, 0xc7, 0x55, 0x28, 0x67, 0x87, 0x9d, 0xfe, 0xb0, 0xb2,
	0x41, 0xf7, 0x55, 0x72, 0xd1, 0xab, 0xac, 0x73, 0xc7, 0x21, 0x98, 0x7a, 0x3d, 0xf2, 0xc6, 0xe7,
	0x50, 0x53, 0xd0, 0xe0, 0xc1, 0x95, 0x2f, 0xc1, 0xda, 0x19, 0xed, 0xde, 0xf9, 0x0e, 0x5e, 0xdf,
	0x3f, 0x78, 0x33, 0x24, 0x1d, 0x67, 0x34, 0xf4, 0xdc, 0x8f, 0xee, 0xe7, 0x0b, 0xcf, 0xbd, 0xb9,
	0xae, 0x6c, 0xd0, 0xb8, 0x20, 0x3f, 0x8d, 0xd3, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0x8f, 0x83,
	0xc1, 0x27, 0xd2, 0x03, 0x00, 0x00,
}
