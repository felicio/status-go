// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sync_settings.proto

package protobuf

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

type SyncSetting_Type int32

const (
	SyncSetting_UNKNOWN                     SyncSetting_Type = 0
	SyncSetting_CURRENCY                    SyncSetting_Type = 1
	SyncSetting_GIF_RECENTS                 SyncSetting_Type = 2
	SyncSetting_GIF_FAVOURITES              SyncSetting_Type = 3
	SyncSetting_MESSAGES_FROM_CONTACTS_ONLY SyncSetting_Type = 4
	SyncSetting_PREFERRED_NAME              SyncSetting_Type = 5
	SyncSetting_PREVIEW_PRIVACY             SyncSetting_Type = 6
	SyncSetting_PROFILE_PICTURES_SHOW_TO    SyncSetting_Type = 7
	SyncSetting_PROFILE_PICTURES_VISIBILITY SyncSetting_Type = 8
	SyncSetting_SEND_STATUS_UPDATES         SyncSetting_Type = 9
	SyncSetting_STICKERS_PACKS_INSTALLED    SyncSetting_Type = 10
	SyncSetting_STICKERS_PACKS_PENDING      SyncSetting_Type = 11
	SyncSetting_STICKERS_RECENT_STICKERS    SyncSetting_Type = 12
)

var SyncSetting_Type_name = map[int32]string{
	0:  "UNKNOWN",
	1:  "CURRENCY",
	2:  "GIF_RECENTS",
	3:  "GIF_FAVOURITES",
	4:  "MESSAGES_FROM_CONTACTS_ONLY",
	5:  "PREFERRED_NAME",
	6:  "PREVIEW_PRIVACY",
	7:  "PROFILE_PICTURES_SHOW_TO",
	8:  "PROFILE_PICTURES_VISIBILITY",
	9:  "SEND_STATUS_UPDATES",
	10: "STICKERS_PACKS_INSTALLED",
	11: "STICKERS_PACKS_PENDING",
	12: "STICKERS_RECENT_STICKERS",
}

var SyncSetting_Type_value = map[string]int32{
	"UNKNOWN":                     0,
	"CURRENCY":                    1,
	"GIF_RECENTS":                 2,
	"GIF_FAVOURITES":              3,
	"MESSAGES_FROM_CONTACTS_ONLY": 4,
	"PREFERRED_NAME":              5,
	"PREVIEW_PRIVACY":             6,
	"PROFILE_PICTURES_SHOW_TO":    7,
	"PROFILE_PICTURES_VISIBILITY": 8,
	"SEND_STATUS_UPDATES":         9,
	"STICKERS_PACKS_INSTALLED":    10,
	"STICKERS_PACKS_PENDING":      11,
	"STICKERS_RECENT_STICKERS":    12,
}

func (x SyncSetting_Type) String() string {
	return proto.EnumName(SyncSetting_Type_name, int32(x))
}

func (SyncSetting_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e2f7a0bce2873c78, []int{0, 0}
}

type SyncSetting struct {
	Type  SyncSetting_Type `protobuf:"varint,1,opt,name=type,proto3,enum=protobuf.SyncSetting_Type" json:"type,omitempty"`
	Clock uint64           `protobuf:"varint,2,opt,name=clock,proto3" json:"clock,omitempty"`
	// Types that are valid to be assigned to Value:
	//	*SyncSetting_ValueString
	//	*SyncSetting_ValueBytes
	//	*SyncSetting_ValueBool
	//	*SyncSetting_ValueInt64
	Value                isSyncSetting_Value `protobuf_oneof:"value"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *SyncSetting) Reset()         { *m = SyncSetting{} }
func (m *SyncSetting) String() string { return proto.CompactTextString(m) }
func (*SyncSetting) ProtoMessage()    {}
func (*SyncSetting) Descriptor() ([]byte, []int) {
	return fileDescriptor_e2f7a0bce2873c78, []int{0}
}

func (m *SyncSetting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncSetting.Unmarshal(m, b)
}
func (m *SyncSetting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncSetting.Marshal(b, m, deterministic)
}
func (m *SyncSetting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncSetting.Merge(m, src)
}
func (m *SyncSetting) XXX_Size() int {
	return xxx_messageInfo_SyncSetting.Size(m)
}
func (m *SyncSetting) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncSetting.DiscardUnknown(m)
}

var xxx_messageInfo_SyncSetting proto.InternalMessageInfo

func (m *SyncSetting) GetType() SyncSetting_Type {
	if m != nil {
		return m.Type
	}
	return SyncSetting_UNKNOWN
}

func (m *SyncSetting) GetClock() uint64 {
	if m != nil {
		return m.Clock
	}
	return 0
}

type isSyncSetting_Value interface {
	isSyncSetting_Value()
}

type SyncSetting_ValueString struct {
	ValueString string `protobuf:"bytes,3,opt,name=value_string,json=valueString,proto3,oneof"`
}

type SyncSetting_ValueBytes struct {
	ValueBytes []byte `protobuf:"bytes,4,opt,name=value_bytes,json=valueBytes,proto3,oneof"`
}

type SyncSetting_ValueBool struct {
	ValueBool bool `protobuf:"varint,5,opt,name=value_bool,json=valueBool,proto3,oneof"`
}

type SyncSetting_ValueInt64 struct {
	ValueInt64 int64 `protobuf:"varint,6,opt,name=value_int64,json=valueInt64,proto3,oneof"`
}

func (*SyncSetting_ValueString) isSyncSetting_Value() {}

func (*SyncSetting_ValueBytes) isSyncSetting_Value() {}

func (*SyncSetting_ValueBool) isSyncSetting_Value() {}

func (*SyncSetting_ValueInt64) isSyncSetting_Value() {}

func (m *SyncSetting) GetValue() isSyncSetting_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *SyncSetting) GetValueString() string {
	if x, ok := m.GetValue().(*SyncSetting_ValueString); ok {
		return x.ValueString
	}
	return ""
}

func (m *SyncSetting) GetValueBytes() []byte {
	if x, ok := m.GetValue().(*SyncSetting_ValueBytes); ok {
		return x.ValueBytes
	}
	return nil
}

func (m *SyncSetting) GetValueBool() bool {
	if x, ok := m.GetValue().(*SyncSetting_ValueBool); ok {
		return x.ValueBool
	}
	return false
}

func (m *SyncSetting) GetValueInt64() int64 {
	if x, ok := m.GetValue().(*SyncSetting_ValueInt64); ok {
		return x.ValueInt64
	}
	return 0
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SyncSetting) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SyncSetting_ValueString)(nil),
		(*SyncSetting_ValueBytes)(nil),
		(*SyncSetting_ValueBool)(nil),
		(*SyncSetting_ValueInt64)(nil),
	}
}

func init() {
	proto.RegisterEnum("protobuf.SyncSetting_Type", SyncSetting_Type_name, SyncSetting_Type_value)
	proto.RegisterType((*SyncSetting)(nil), "protobuf.SyncSetting")
}

func init() { proto.RegisterFile("sync_settings.proto", fileDescriptor_e2f7a0bce2873c78) }

var fileDescriptor_e2f7a0bce2873c78 = []byte{
	// 447 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0xcd, 0x8e, 0xd3, 0x30,
	0x10, 0x80, 0x9b, 0x36, 0xfd, 0x9b, 0x94, 0x5d, 0xcb, 0x45, 0x10, 0x2d, 0x48, 0x1b, 0x96, 0x4b,
	0x4e, 0x41, 0x02, 0xc4, 0x85, 0x53, 0x9a, 0x38, 0xad, 0xd5, 0xd4, 0x89, 0x6c, 0xa7, 0x55, 0xb9,
	0x58, 0xb4, 0x0a, 0xab, 0x8a, 0x2a, 0xa9, 0x36, 0x29, 0x52, 0x9e, 0x89, 0x97, 0xe0, 0xd1, 0x50,
	0x12, 0xca, 0xdf, 0x9e, 0xec, 0xf9, 0xe6, 0x9b, 0xf1, 0x68, 0x64, 0x98, 0x16, 0x55, 0xb6, 0x57,
	0x45, 0x5a, 0x96, 0x87, 0xec, 0xbe, 0x70, 0x4e, 0x0f, 0x79, 0x99, 0xe3, 0x51, 0x73, 0xec, 0xce,
	0x5f, 0xee, 0xbe, 0xeb, 0x60, 0x88, 0x2a, 0xdb, 0x8b, 0x56, 0xc0, 0x0e, 0xe8, 0x65, 0x75, 0x4a,
	0x4d, 0xcd, 0xd2, 0xec, 0xab, 0xb7, 0x37, 0xce, 0x45, 0x74, 0xfe, 0x92, 0x1c, 0x59, 0x9d, 0x52,
	0xde, 0x78, 0xf8, 0x29, 0xf4, 0xf7, 0xc7, 0x7c, 0xff, 0xd5, 0xec, 0x5a, 0x9a, 0xad, 0xf3, 0x36,
	0xc0, 0xaf, 0x61, 0xf2, 0xed, 0xf3, 0xf1, 0x9c, 0xaa, 0xa2, 0x7c, 0x38, 0x64, 0xf7, 0x66, 0xcf,
	0xd2, 0xec, 0xf1, 0xa2, 0xc3, 0x8d, 0x86, 0x8a, 0x06, 0xe2, 0x57, 0xd0, 0x86, 0x6a, 0x57, 0x95,
	0x69, 0x61, 0xea, 0x96, 0x66, 0x4f, 0x16, 0x1d, 0x0e, 0x0d, 0x9c, 0xd5, 0x0c, 0xdf, 0x02, 0xfc,
	0x52, 0xf2, 0xfc, 0x68, 0xf6, 0x2d, 0xcd, 0x1e, 0x2d, 0x3a, 0x7c, 0xdc, 0x1a, 0x79, 0x7e, 0xfc,
	0xd3, 0xe3, 0x90, 0x95, 0x1f, 0xde, 0x9b, 0x03, 0x4b, 0xb3, 0x7b, 0xbf, 0x7b, 0xd0, 0x9a, 0xdd,
	0xfd, 0xe8, 0x82, 0x5e, 0x0f, 0x8c, 0x0d, 0x18, 0x26, 0x6c, 0xc9, 0xa2, 0x0d, 0x43, 0x1d, 0x3c,
	0x81, 0x91, 0x97, 0x70, 0x4e, 0x98, 0xb7, 0x45, 0x1a, 0xbe, 0x06, 0x63, 0x4e, 0x03, 0xc5, 0x89,
	0x47, 0x98, 0x14, 0xa8, 0x8b, 0x31, 0x5c, 0xd5, 0x20, 0x70, 0xd7, 0x51, 0xc2, 0xa9, 0x24, 0x02,
	0xf5, 0xf0, 0x2d, 0xbc, 0x58, 0x11, 0x21, 0xdc, 0x39, 0x11, 0x2a, 0xe0, 0xd1, 0x4a, 0x79, 0x11,
	0x93, 0xae, 0x27, 0x85, 0x8a, 0x58, 0xb8, 0x45, 0x7a, 0x5d, 0x14, 0x73, 0x12, 0x10, 0xce, 0x89,
	0xaf, 0x98, 0xbb, 0x22, 0xa8, 0x8f, 0xa7, 0x70, 0x1d, 0x73, 0xb2, 0xa6, 0x64, 0xa3, 0x62, 0x4e,
	0xd7, 0xae, 0xb7, 0x45, 0x03, 0xfc, 0x12, 0xcc, 0x98, 0x47, 0x01, 0x0d, 0x89, 0x8a, 0xa9, 0x27,
	0x13, 0x4e, 0x84, 0x12, 0x8b, 0x68, 0xa3, 0x64, 0x84, 0x86, 0xf5, 0x3b, 0x8f, 0xb2, 0x6b, 0x2a,
	0xe8, 0x8c, 0x86, 0x54, 0x6e, 0xd1, 0x08, 0x3f, 0x87, 0xa9, 0x20, 0xcc, 0x57, 0x42, 0xba, 0x32,
	0x11, 0x2a, 0x89, 0x7d, 0xb7, 0x9e, 0x70, 0x5c, 0xf7, 0x15, 0x92, 0x7a, 0x4b, 0xc2, 0x85, 0x8a,
	0x5d, 0x6f, 0x29, 0x14, 0x65, 0x42, 0xba, 0x61, 0x48, 0x7c, 0x04, 0xf8, 0x06, 0x9e, 0xfd, 0x97,
	0x8d, 0x09, 0xf3, 0x29, 0x9b, 0x23, 0xe3, 0x9f, 0xca, 0x76, 0x0b, 0xea, 0x12, 0xa3, 0xc9, 0x6c,
	0x08, 0xfd, 0x76, 0xe5, 0x4f, 0x3e, 0x19, 0xce, 0x9b, 0x8f, 0x97, 0x3f, 0xb1, 0x1b, 0x34, 0xb7,
	0x77, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x91, 0x74, 0x80, 0x25, 0x64, 0x02, 0x00, 0x00,
}