// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/enums/target_impression_share_location.proto

package enums

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Enum describing possible goals.
type TargetImpressionShareLocationEnum_TargetImpressionShareLocation int32

const (
	// Not specified.
	TargetImpressionShareLocationEnum_UNSPECIFIED TargetImpressionShareLocationEnum_TargetImpressionShareLocation = 0
	// Used for return value only. Represents value unknown in this version.
	TargetImpressionShareLocationEnum_UNKNOWN TargetImpressionShareLocationEnum_TargetImpressionShareLocation = 1
	// Any location on the web page.
	TargetImpressionShareLocationEnum_ANYWHERE_ON_PAGE TargetImpressionShareLocationEnum_TargetImpressionShareLocation = 2
	// Top box of ads.
	TargetImpressionShareLocationEnum_TOP_OF_PAGE TargetImpressionShareLocationEnum_TargetImpressionShareLocation = 3
	// Top slot in the top box of ads.
	TargetImpressionShareLocationEnum_ABSOLUTE_TOP_OF_PAGE TargetImpressionShareLocationEnum_TargetImpressionShareLocation = 4
)

var TargetImpressionShareLocationEnum_TargetImpressionShareLocation_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "ANYWHERE_ON_PAGE",
	3: "TOP_OF_PAGE",
	4: "ABSOLUTE_TOP_OF_PAGE",
}

var TargetImpressionShareLocationEnum_TargetImpressionShareLocation_value = map[string]int32{
	"UNSPECIFIED":          0,
	"UNKNOWN":              1,
	"ANYWHERE_ON_PAGE":     2,
	"TOP_OF_PAGE":          3,
	"ABSOLUTE_TOP_OF_PAGE": 4,
}

func (x TargetImpressionShareLocationEnum_TargetImpressionShareLocation) String() string {
	return proto.EnumName(TargetImpressionShareLocationEnum_TargetImpressionShareLocation_name, int32(x))
}

func (TargetImpressionShareLocationEnum_TargetImpressionShareLocation) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_99554193aed0500a, []int{0, 0}
}

// Container for enum describing where on the first search results page the
// automated bidding system should target impressions for the
// TargetImpressionShare bidding strategy.
type TargetImpressionShareLocationEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TargetImpressionShareLocationEnum) Reset()         { *m = TargetImpressionShareLocationEnum{} }
func (m *TargetImpressionShareLocationEnum) String() string { return proto.CompactTextString(m) }
func (*TargetImpressionShareLocationEnum) ProtoMessage()    {}
func (*TargetImpressionShareLocationEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_99554193aed0500a, []int{0}
}

func (m *TargetImpressionShareLocationEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TargetImpressionShareLocationEnum.Unmarshal(m, b)
}
func (m *TargetImpressionShareLocationEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TargetImpressionShareLocationEnum.Marshal(b, m, deterministic)
}
func (m *TargetImpressionShareLocationEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TargetImpressionShareLocationEnum.Merge(m, src)
}
func (m *TargetImpressionShareLocationEnum) XXX_Size() int {
	return xxx_messageInfo_TargetImpressionShareLocationEnum.Size(m)
}
func (m *TargetImpressionShareLocationEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_TargetImpressionShareLocationEnum.DiscardUnknown(m)
}

var xxx_messageInfo_TargetImpressionShareLocationEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.enums.TargetImpressionShareLocationEnum_TargetImpressionShareLocation", TargetImpressionShareLocationEnum_TargetImpressionShareLocation_name, TargetImpressionShareLocationEnum_TargetImpressionShareLocation_value)
	proto.RegisterType((*TargetImpressionShareLocationEnum)(nil), "google.ads.googleads.v2.enums.TargetImpressionShareLocationEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/enums/target_impression_share_location.proto", fileDescriptor_99554193aed0500a)
}

var fileDescriptor_99554193aed0500a = []byte{
	// 343 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0xbd, 0x6a, 0xeb, 0x30,
	0x18, 0xbd, 0x76, 0x2e, 0xf7, 0x82, 0x32, 0xd4, 0x98, 0x0c, 0xa5, 0x34, 0x43, 0xf2, 0x00, 0x32,
	0xb8, 0x9b, 0x3a, 0xc9, 0x8d, 0x92, 0x86, 0x06, 0xdb, 0x34, 0x7f, 0xb4, 0x18, 0x84, 0x1a, 0x1b,
	0xd5, 0x10, 0x4b, 0xc6, 0x72, 0x32, 0xf6, 0x61, 0xda, 0xad, 0x8f, 0xd2, 0x47, 0xe9, 0x0b, 0x74,
	0x2d, 0x96, 0x12, 0xd3, 0xa5, 0x59, 0xc4, 0x41, 0xe7, 0x3b, 0xe7, 0x7c, 0x3f, 0x60, 0xc4, 0xa5,
	0xe4, 0xdb, 0xcc, 0x63, 0xa9, 0xf2, 0x0c, 0x6c, 0xd0, 0xde, 0xf7, 0x32, 0xb1, 0x2b, 0x94, 0x57,
	0xb3, 0x8a, 0x67, 0x35, 0xcd, 0x8b, 0xb2, 0xca, 0x94, 0xca, 0xa5, 0xa0, 0xea, 0x99, 0x55, 0x19,
	0xdd, 0xca, 0x0d, 0xab, 0x73, 0x29, 0x60, 0x59, 0xc9, 0x5a, 0xba, 0x7d, 0x23, 0x85, 0x2c, 0x55,
	0xb0, 0x75, 0x81, 0x7b, 0x1f, 0x6a, 0x97, 0x8b, 0xcb, 0x63, 0x48, 0x99, 0x7b, 0x4c, 0x08, 0x59,
	0x6b, 0xad, 0x32, 0xe2, 0xe1, 0x9b, 0x05, 0x06, 0x0b, 0x9d, 0x33, 0x6d, 0x63, 0xe6, 0x4d, 0xca,
	0xec, 0x10, 0x42, 0xc4, 0xae, 0x18, 0xbe, 0x80, 0xfe, 0xc9, 0x22, 0xf7, 0x0c, 0x74, 0x97, 0xe1,
	0x3c, 0x26, 0x37, 0xd3, 0xf1, 0x94, 0x8c, 0x9c, 0x3f, 0x6e, 0x17, 0xfc, 0x5f, 0x86, 0x77, 0x61,
	0xb4, 0x0e, 0x1d, 0xcb, 0xed, 0x01, 0x07, 0x87, 0x0f, 0xeb, 0x5b, 0x72, 0x4f, 0x68, 0x14, 0xd2,
	0x18, 0x4f, 0x88, 0x63, 0x37, 0x9a, 0x45, 0x14, 0xd3, 0x68, 0x6c, 0x3e, 0x3a, 0xee, 0x39, 0xe8,
	0xe1, 0x60, 0x1e, 0xcd, 0x96, 0x0b, 0x42, 0x7f, 0x32, 0x7f, 0x83, 0x2f, 0x0b, 0x0c, 0x36, 0xb2,
	0x80, 0x27, 0x27, 0x0d, 0x86, 0x27, 0x7b, 0x8c, 0x9b, 0x79, 0x63, 0xeb, 0x31, 0x38, 0x98, 0x70,
	0xb9, 0x65, 0x82, 0x43, 0x59, 0x71, 0x8f, 0x67, 0x42, 0x6f, 0xe3, 0x78, 0x84, 0x32, 0x57, 0xbf,
	0xdc, 0xe4, 0x5a, 0xbf, 0xaf, 0x76, 0x67, 0x82, 0xf1, 0xbb, 0xdd, 0x9f, 0x18, 0x2b, 0x9c, 0x2a,
	0x68, 0x60, 0x83, 0x56, 0x3e, 0x6c, 0x96, 0xa6, 0x3e, 0x8e, 0x7c, 0x82, 0x53, 0x95, 0xb4, 0x7c,
	0xb2, 0xf2, 0x13, 0xcd, 0x7f, 0xda, 0x03, 0xf3, 0x89, 0x10, 0x4e, 0x15, 0x42, 0x6d, 0x05, 0x42,
	0x2b, 0x1f, 0x21, 0x5d, 0xf3, 0xf4, 0x4f, 0x37, 0x76, 0xf5, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xd8,
	0x74, 0x3c, 0x06, 0x2b, 0x02, 0x00, 0x00,
}