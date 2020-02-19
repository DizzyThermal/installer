// Code generated by protoc-gen-go.
// source: ots_filter.proto
// DO NOT EDIT!

package otsprotocol

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type VariantType int32

const (
	VariantType_VT_INTEGER VariantType = 0
	VariantType_VT_DOUBLE  VariantType = 1
	// VT_BOOLEAN = 2;
	VariantType_VT_STRING VariantType = 3
	VariantType_VT_NULL   VariantType = 6
	VariantType_VT_BLOB   VariantType = 7
)

var VariantType_name = map[int32]string{
	0: "VT_INTEGER",
	1: "VT_DOUBLE",
	3: "VT_STRING",
	6: "VT_NULL",
	7: "VT_BLOB",
}
var VariantType_value = map[string]int32{
	"VT_INTEGER": 0,
	"VT_DOUBLE":  1,
	"VT_STRING":  3,
	"VT_NULL":    6,
	"VT_BLOB":    7,
}

func (x VariantType) Enum() *VariantType {
	p := new(VariantType)
	*p = x
	return p
}
func (x VariantType) String() string {
	return proto.EnumName(VariantType_name, int32(x))
}
func (x *VariantType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(VariantType_value, data, "VariantType")
	if err != nil {
		return err
	}
	*x = VariantType(value)
	return nil
}
func (VariantType) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type FilterType int32

const (
	FilterType_FT_SINGLE_COLUMN_VALUE    FilterType = 1
	FilterType_FT_COMPOSITE_COLUMN_VALUE FilterType = 2
	FilterType_FT_COLUMN_PAGINATION      FilterType = 3
)

var FilterType_name = map[int32]string{
	1: "FT_SINGLE_COLUMN_VALUE",
	2: "FT_COMPOSITE_COLUMN_VALUE",
	3: "FT_COLUMN_PAGINATION",
}
var FilterType_value = map[string]int32{
	"FT_SINGLE_COLUMN_VALUE":    1,
	"FT_COMPOSITE_COLUMN_VALUE": 2,
	"FT_COLUMN_PAGINATION":      3,
}

func (x FilterType) Enum() *FilterType {
	p := new(FilterType)
	*p = x
	return p
}
func (x FilterType) String() string {
	return proto.EnumName(FilterType_name, int32(x))
}
func (x *FilterType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(FilterType_value, data, "FilterType")
	if err != nil {
		return err
	}
	*x = FilterType(value)
	return nil
}
func (FilterType) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

type ComparatorType int32

const (
	ComparatorType_CT_EQUAL         ComparatorType = 1
	ComparatorType_CT_NOT_EQUAL     ComparatorType = 2
	ComparatorType_CT_GREATER_THAN  ComparatorType = 3
	ComparatorType_CT_GREATER_EQUAL ComparatorType = 4
	ComparatorType_CT_LESS_THAN     ComparatorType = 5
	ComparatorType_CT_LESS_EQUAL    ComparatorType = 6
)

var ComparatorType_name = map[int32]string{
	1: "CT_EQUAL",
	2: "CT_NOT_EQUAL",
	3: "CT_GREATER_THAN",
	4: "CT_GREATER_EQUAL",
	5: "CT_LESS_THAN",
	6: "CT_LESS_EQUAL",
}
var ComparatorType_value = map[string]int32{
	"CT_EQUAL":         1,
	"CT_NOT_EQUAL":     2,
	"CT_GREATER_THAN":  3,
	"CT_GREATER_EQUAL": 4,
	"CT_LESS_THAN":     5,
	"CT_LESS_EQUAL":    6,
}

func (x ComparatorType) Enum() *ComparatorType {
	p := new(ComparatorType)
	*p = x
	return p
}
func (x ComparatorType) String() string {
	return proto.EnumName(ComparatorType_name, int32(x))
}
func (x *ComparatorType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ComparatorType_value, data, "ComparatorType")
	if err != nil {
		return err
	}
	*x = ComparatorType(value)
	return nil
}
func (ComparatorType) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

type LogicalOperator int32

const (
	LogicalOperator_LO_NOT LogicalOperator = 1
	LogicalOperator_LO_AND LogicalOperator = 2
	LogicalOperator_LO_OR  LogicalOperator = 3
)

var LogicalOperator_name = map[int32]string{
	1: "LO_NOT",
	2: "LO_AND",
	3: "LO_OR",
}
var LogicalOperator_value = map[string]int32{
	"LO_NOT": 1,
	"LO_AND": 2,
	"LO_OR":  3,
}

func (x LogicalOperator) Enum() *LogicalOperator {
	p := new(LogicalOperator)
	*p = x
	return p
}
func (x LogicalOperator) String() string {
	return proto.EnumName(LogicalOperator_name, int32(x))
}
func (x *LogicalOperator) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(LogicalOperator_value, data, "LogicalOperator")
	if err != nil {
		return err
	}
	*x = LogicalOperator(value)
	return nil
}
func (LogicalOperator) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

type ValueTransferRule struct {
	Regex            *string      `protobuf:"bytes,1,req,name=regex" json:"regex,omitempty"`
	CastType         *VariantType `protobuf:"varint,2,opt,name=cast_type,enum=otsprotocol.VariantType" json:"cast_type,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *ValueTransferRule) Reset()                    { *m = ValueTransferRule{} }
func (m *ValueTransferRule) String() string            { return proto.CompactTextString(m) }
func (*ValueTransferRule) ProtoMessage()               {}
func (*ValueTransferRule) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *ValueTransferRule) GetRegex() string {
	if m != nil && m.Regex != nil {
		return *m.Regex
	}
	return ""
}

func (m *ValueTransferRule) GetCastType() VariantType {
	if m != nil && m.CastType != nil {
		return *m.CastType
	}
	return VariantType_VT_INTEGER
}

type SingleColumnValueFilter struct {
	Comparator        *ComparatorType    `protobuf:"varint,1,req,name=comparator,enum=otsprotocol.ComparatorType" json:"comparator,omitempty"`
	ColumnName        *string            `protobuf:"bytes,2,req,name=column_name" json:"column_name,omitempty"`
	ColumnValue       []byte             `protobuf:"bytes,3,req,name=column_value" json:"column_value,omitempty"`
	FilterIfMissing   *bool              `protobuf:"varint,4,req,name=filter_if_missing" json:"filter_if_missing,omitempty"`
	LatestVersionOnly *bool              `protobuf:"varint,5,req,name=latest_version_only" json:"latest_version_only,omitempty"`
	ValueTransRule    *ValueTransferRule `protobuf:"bytes,6,opt,name=value_trans_rule" json:"value_trans_rule,omitempty"`
	XXX_unrecognized  []byte             `json:"-"`
}

func (m *SingleColumnValueFilter) Reset()                    { *m = SingleColumnValueFilter{} }
func (m *SingleColumnValueFilter) String() string            { return proto.CompactTextString(m) }
func (*SingleColumnValueFilter) ProtoMessage()               {}
func (*SingleColumnValueFilter) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *SingleColumnValueFilter) GetComparator() ComparatorType {
	if m != nil && m.Comparator != nil {
		return *m.Comparator
	}
	return ComparatorType_CT_EQUAL
}

func (m *SingleColumnValueFilter) GetColumnName() string {
	if m != nil && m.ColumnName != nil {
		return *m.ColumnName
	}
	return ""
}

func (m *SingleColumnValueFilter) GetColumnValue() []byte {
	if m != nil {
		return m.ColumnValue
	}
	return nil
}

func (m *SingleColumnValueFilter) GetFilterIfMissing() bool {
	if m != nil && m.FilterIfMissing != nil {
		return *m.FilterIfMissing
	}
	return false
}

func (m *SingleColumnValueFilter) GetLatestVersionOnly() bool {
	if m != nil && m.LatestVersionOnly != nil {
		return *m.LatestVersionOnly
	}
	return false
}

func (m *SingleColumnValueFilter) GetValueTransRule() *ValueTransferRule {
	if m != nil {
		return m.ValueTransRule
	}
	return nil
}

type CompositeColumnValueFilter struct {
	Combinator       *LogicalOperator `protobuf:"varint,1,req,name=combinator,enum=otsprotocol.LogicalOperator" json:"combinator,omitempty"`
	SubFilters       []*Filter        `protobuf:"bytes,2,rep,name=sub_filters" json:"sub_filters,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *CompositeColumnValueFilter) Reset()                    { *m = CompositeColumnValueFilter{} }
func (m *CompositeColumnValueFilter) String() string            { return proto.CompactTextString(m) }
func (*CompositeColumnValueFilter) ProtoMessage()               {}
func (*CompositeColumnValueFilter) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *CompositeColumnValueFilter) GetCombinator() LogicalOperator {
	if m != nil && m.Combinator != nil {
		return *m.Combinator
	}
	return LogicalOperator_LO_NOT
}

func (m *CompositeColumnValueFilter) GetSubFilters() []*Filter {
	if m != nil {
		return m.SubFilters
	}
	return nil
}

type ColumnPaginationFilter struct {
	Offset           *int32 `protobuf:"varint,1,req,name=offset" json:"offset,omitempty"`
	Limit            *int32 `protobuf:"varint,2,req,name=limit" json:"limit,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ColumnPaginationFilter) Reset()                    { *m = ColumnPaginationFilter{} }
func (m *ColumnPaginationFilter) String() string            { return proto.CompactTextString(m) }
func (*ColumnPaginationFilter) ProtoMessage()               {}
func (*ColumnPaginationFilter) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *ColumnPaginationFilter) GetOffset() int32 {
	if m != nil && m.Offset != nil {
		return *m.Offset
	}
	return 0
}

func (m *ColumnPaginationFilter) GetLimit() int32 {
	if m != nil && m.Limit != nil {
		return *m.Limit
	}
	return 0
}

type Filter struct {
	Type             *FilterType `protobuf:"varint,1,req,name=type,enum=otsprotocol.FilterType" json:"type,omitempty"`
	Filter           []byte      `protobuf:"bytes,2,req,name=filter" json:"filter,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *Filter) Reset()                    { *m = Filter{} }
func (m *Filter) String() string            { return proto.CompactTextString(m) }
func (*Filter) ProtoMessage()               {}
func (*Filter) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *Filter) GetType() FilterType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return FilterType_FT_SINGLE_COLUMN_VALUE
}

func (m *Filter) GetFilter() []byte {
	if m != nil {
		return m.Filter
	}
	return nil
}

func init() {
	proto.RegisterType((*ValueTransferRule)(nil), "otsprotocol.ValueTransferRule")
	proto.RegisterType((*SingleColumnValueFilter)(nil), "otsprotocol.SingleColumnValueFilter")
	proto.RegisterType((*CompositeColumnValueFilter)(nil), "otsprotocol.CompositeColumnValueFilter")
	proto.RegisterType((*ColumnPaginationFilter)(nil), "otsprotocol.ColumnPaginationFilter")
	proto.RegisterType((*Filter)(nil), "otsprotocol.Filter")
	proto.RegisterEnum("otsprotocol.VariantType", VariantType_name, VariantType_value)
	proto.RegisterEnum("otsprotocol.FilterType", FilterType_name, FilterType_value)
	proto.RegisterEnum("otsprotocol.ComparatorType", ComparatorType_name, ComparatorType_value)
	proto.RegisterEnum("otsprotocol.LogicalOperator", LogicalOperator_name, LogicalOperator_value)
}

func init() { proto.RegisterFile("ots_filter.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 585 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x92, 0x51, 0x6b, 0xdb, 0x30,
	0x14, 0x85, 0x67, 0xa7, 0x71, 0x9b, 0xeb, 0x34, 0x55, 0x95, 0xd2, 0xba, 0xed, 0x36, 0x42, 0x60,
	0x60, 0x32, 0xe8, 0x46, 0x18, 0x6c, 0x6f, 0xc3, 0x75, 0xdd, 0x2c, 0xe0, 0xda, 0x5d, 0xa2, 0xf8,
	0x55, 0xb8, 0x41, 0x09, 0x02, 0xc7, 0x0a, 0x96, 0x52, 0xda, 0xb7, 0xfd, 0xdb, 0xfd, 0x8d, 0x61,
	0xd9, 0x1d, 0x69, 0xe8, 0x9b, 0x75, 0xef, 0xd5, 0x77, 0xee, 0xd1, 0x31, 0x20, 0xa1, 0x24, 0x5d,
	0xf0, 0x4c, 0xb1, 0xe2, 0x6a, 0x5d, 0x08, 0x25, 0xb0, 0x2d, 0x94, 0xd4, 0x5f, 0x73, 0x91, 0xf5,
	0x63, 0x38, 0x4e, 0xd2, 0x6c, 0xc3, 0x48, 0x91, 0xe6, 0x72, 0xc1, 0x8a, 0xc9, 0x26, 0x63, 0xf8,
	0x10, 0x9a, 0x05, 0x5b, 0xb2, 0x27, 0xc7, 0xe8, 0x99, 0x6e, 0x0b, 0x7f, 0x86, 0xd6, 0x3c, 0x95,
	0x8a, 0xaa, 0xe7, 0x35, 0x73, 0xcc, 0x9e, 0xe1, 0x76, 0x86, 0xce, 0xd5, 0x16, 0xe4, 0x2a, 0x49,
	0x0b, 0x9e, 0xe6, 0x8a, 0x3c, 0xaf, 0x59, 0xff, 0xaf, 0x01, 0x67, 0x53, 0x9e, 0x2f, 0x33, 0xe6,
	0x8b, 0x6c, 0xb3, 0xca, 0x35, 0xfd, 0x56, 0xeb, 0xe3, 0x2f, 0x00, 0x73, 0xb1, 0x5a, 0xa7, 0x45,
	0xaa, 0x44, 0xa1, 0xe1, 0x9d, 0xe1, 0xe5, 0x2b, 0x92, 0xff, 0xbf, 0x5d, 0xc2, 0x70, 0x17, 0xec,
	0xb9, 0xa6, 0xd0, 0x3c, 0x5d, 0x95, 0xda, 0xe5, 0x3a, 0x27, 0xd0, 0xae, 0x8b, 0x8f, 0x25, 0xdb,
	0x69, 0xf4, 0x4c, 0xb7, 0x8d, 0xcf, 0xe1, 0xb8, 0x72, 0x49, 0xf9, 0x82, 0xae, 0xb8, 0x94, 0x3c,
	0x5f, 0x3a, 0x7b, 0x3d, 0xd3, 0x3d, 0xc0, 0x97, 0xd0, 0xcd, 0x52, 0xc5, 0xa4, 0xa2, 0x8f, 0xac,
	0x90, 0x5c, 0xe4, 0x54, 0xe4, 0xd9, 0xb3, 0xd3, 0xd4, 0xcd, 0x1f, 0x80, 0x34, 0x86, 0xaa, 0xf2,
	0x05, 0x68, 0xb1, 0xc9, 0x98, 0x63, 0xf5, 0x0c, 0xd7, 0x1e, 0x7e, 0xdc, 0xf1, 0xb8, 0xf3, 0x4a,
	0xfd, 0x27, 0xb8, 0x28, 0xd7, 0x15, 0x92, 0xab, 0x37, 0xbc, 0x7e, 0xd5, 0x5e, 0x1f, 0x78, 0xbe,
	0xe5, 0xf5, 0xfd, 0x2b, 0x62, 0x28, 0x96, 0x7c, 0x9e, 0x66, 0xf1, 0x9a, 0x69, 0xc3, 0xd8, 0x05,
	0x5b, 0x6e, 0x1e, 0xea, 0xac, 0xa4, 0x63, 0xf6, 0x1a, 0xae, 0x3d, 0xec, 0xbe, 0xba, 0x52, 0xb1,
	0xfb, 0xdf, 0xe1, 0xb4, 0x12, 0xbc, 0x4f, 0x97, 0xa5, 0x00, 0x17, 0x79, 0xad, 0xda, 0x01, 0x4b,
	0x2c, 0x16, 0x92, 0x29, 0xad, 0xd8, 0x2c, 0x93, 0xcc, 0xf8, 0x8a, 0x2b, 0xfd, 0x74, 0xcd, 0xfe,
	0x4f, 0xb0, 0xea, 0xc1, 0x4f, 0xb0, 0xa7, 0xe3, 0xac, 0x16, 0x3b, 0x7b, 0x43, 0x45, 0x07, 0xd0,
	0x01, 0xab, 0xda, 0x47, 0x03, 0xda, 0x83, 0x19, 0xd8, 0x5b, 0x61, 0xe3, 0x0e, 0x40, 0x42, 0xe8,
	0x38, 0x22, 0xc1, 0x28, 0x98, 0xa0, 0x77, 0xf8, 0x10, 0x5a, 0x09, 0xa1, 0x37, 0xf1, 0xec, 0x3a,
	0x0c, 0x90, 0x51, 0x1f, 0xa7, 0x64, 0x32, 0x8e, 0x46, 0xa8, 0x81, 0x6d, 0xd8, 0x4f, 0x08, 0x8d,
	0x66, 0x61, 0x88, 0xac, 0xfa, 0x70, 0x1d, 0xc6, 0xd7, 0x68, 0x7f, 0x90, 0x02, 0x6c, 0x89, 0x5e,
	0xc0, 0xe9, 0x2d, 0xa1, 0xd3, 0x71, 0x34, 0x0a, 0x03, 0xea, 0xc7, 0xe1, 0xec, 0x2e, 0xa2, 0x89,
	0x17, 0xce, 0x4a, 0xe4, 0x07, 0x38, 0xbf, 0x25, 0xd4, 0x8f, 0xef, 0xee, 0xe3, 0xe9, 0x98, 0xec,
	0xb4, 0x4d, 0xec, 0xc0, 0x89, 0x6e, 0xeb, 0xe2, 0xbd, 0x37, 0x1a, 0x47, 0x1e, 0x19, 0xc7, 0x11,
	0x6a, 0x0c, 0xfe, 0x18, 0xd0, 0xd9, 0xf9, 0xbb, 0xda, 0x70, 0xe0, 0x13, 0x1a, 0xfc, 0x9e, 0x79,
	0x21, 0x32, 0x30, 0x82, 0xb6, 0x4f, 0x68, 0x14, 0xbf, 0x54, 0x4c, 0xdc, 0x85, 0x23, 0x9f, 0xd0,
	0xd1, 0x24, 0xf0, 0x48, 0x30, 0xa1, 0xe4, 0x97, 0x17, 0xa1, 0x06, 0x3e, 0x01, 0xb4, 0x55, 0xac,
	0x46, 0xf7, 0xea, 0xcb, 0x61, 0x30, 0x9d, 0x56, 0x73, 0x4d, 0x7c, 0x0c, 0x87, 0x2f, 0x95, 0x6a,
	0xc8, 0x1a, 0x7c, 0x83, 0xa3, 0xdd, 0xcc, 0x01, 0xac, 0x30, 0x2e, 0x45, 0x91, 0x51, 0x7f, 0x7b,
	0xd1, 0x0d, 0x32, 0x71, 0x0b, 0x9a, 0x61, 0x4c, 0xe3, 0x09, 0x6a, 0xfc, 0x0b, 0x00, 0x00, 0xff,
	0xff, 0xb3, 0x10, 0x19, 0xa7, 0xc1, 0x03, 0x00, 0x00,
}
