// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/securitycenter/v1p1beta1/organization_settings.proto

package securitycenter

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

// The mode of inclusion when running Asset Discovery.
// Asset discovery can be limited by explicitly identifying projects to be
// included or excluded. If INCLUDE_ONLY is set, then only those projects
// within the organization and their children are discovered during asset
// discovery. If EXCLUDE is set, then projects that don't match those
// projects are discovered during asset discovery. If neither are set, then
// all projects within the organization are discovered during asset
// discovery.
type OrganizationSettings_AssetDiscoveryConfig_InclusionMode int32

const (
	// Unspecified. Setting the mode with this value will disable
	// inclusion/exclusion filtering for Asset Discovery.
	OrganizationSettings_AssetDiscoveryConfig_INCLUSION_MODE_UNSPECIFIED OrganizationSettings_AssetDiscoveryConfig_InclusionMode = 0
	// Asset Discovery will capture only the resources within the projects
	// specified. All other resources will be ignored.
	OrganizationSettings_AssetDiscoveryConfig_INCLUDE_ONLY OrganizationSettings_AssetDiscoveryConfig_InclusionMode = 1
	// Asset Discovery will ignore all resources under the projects specified.
	// All other resources will be retrieved.
	OrganizationSettings_AssetDiscoveryConfig_EXCLUDE OrganizationSettings_AssetDiscoveryConfig_InclusionMode = 2
)

var OrganizationSettings_AssetDiscoveryConfig_InclusionMode_name = map[int32]string{
	0: "INCLUSION_MODE_UNSPECIFIED",
	1: "INCLUDE_ONLY",
	2: "EXCLUDE",
}

var OrganizationSettings_AssetDiscoveryConfig_InclusionMode_value = map[string]int32{
	"INCLUSION_MODE_UNSPECIFIED": 0,
	"INCLUDE_ONLY":               1,
	"EXCLUDE":                    2,
}

func (x OrganizationSettings_AssetDiscoveryConfig_InclusionMode) String() string {
	return proto.EnumName(OrganizationSettings_AssetDiscoveryConfig_InclusionMode_name, int32(x))
}

func (OrganizationSettings_AssetDiscoveryConfig_InclusionMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3c5f40c5918aa7c6, []int{0, 0, 0}
}

// User specified settings that are attached to the Cloud Security Command
// Center (Cloud SCC) organization.
type OrganizationSettings struct {
	// The relative resource name of the settings. See:
	// https://cloud.google.com/apis/design/resource_names#relative_resource_name
	// Example:
	// "organizations/{organization_id}/organizationSettings".
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// A flag that indicates if Asset Discovery should be enabled. If the flag is
	// set to `true`, then discovery of assets will occur. If it is set to `false,
	// all historical assets will remain, but discovery of future assets will not
	// occur.
	EnableAssetDiscovery bool `protobuf:"varint,2,opt,name=enable_asset_discovery,json=enableAssetDiscovery,proto3" json:"enable_asset_discovery,omitempty"`
	// The configuration used for Asset Discovery runs.
	AssetDiscoveryConfig *OrganizationSettings_AssetDiscoveryConfig `protobuf:"bytes,3,opt,name=asset_discovery_config,json=assetDiscoveryConfig,proto3" json:"asset_discovery_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                   `json:"-"`
	XXX_unrecognized     []byte                                     `json:"-"`
	XXX_sizecache        int32                                      `json:"-"`
}

func (m *OrganizationSettings) Reset()         { *m = OrganizationSettings{} }
func (m *OrganizationSettings) String() string { return proto.CompactTextString(m) }
func (*OrganizationSettings) ProtoMessage()    {}
func (*OrganizationSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c5f40c5918aa7c6, []int{0}
}

func (m *OrganizationSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrganizationSettings.Unmarshal(m, b)
}
func (m *OrganizationSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrganizationSettings.Marshal(b, m, deterministic)
}
func (m *OrganizationSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrganizationSettings.Merge(m, src)
}
func (m *OrganizationSettings) XXX_Size() int {
	return xxx_messageInfo_OrganizationSettings.Size(m)
}
func (m *OrganizationSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_OrganizationSettings.DiscardUnknown(m)
}

var xxx_messageInfo_OrganizationSettings proto.InternalMessageInfo

func (m *OrganizationSettings) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *OrganizationSettings) GetEnableAssetDiscovery() bool {
	if m != nil {
		return m.EnableAssetDiscovery
	}
	return false
}

func (m *OrganizationSettings) GetAssetDiscoveryConfig() *OrganizationSettings_AssetDiscoveryConfig {
	if m != nil {
		return m.AssetDiscoveryConfig
	}
	return nil
}

// The configuration used for Asset Discovery runs.
type OrganizationSettings_AssetDiscoveryConfig struct {
	// The project ids to use for filtering asset discovery.
	ProjectIds []string `protobuf:"bytes,1,rep,name=project_ids,json=projectIds,proto3" json:"project_ids,omitempty"`
	// The mode to use for filtering asset discovery.
	InclusionMode        OrganizationSettings_AssetDiscoveryConfig_InclusionMode `protobuf:"varint,2,opt,name=inclusion_mode,json=inclusionMode,proto3,enum=google.cloud.securitycenter.v1p1beta1.OrganizationSettings_AssetDiscoveryConfig_InclusionMode" json:"inclusion_mode,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                `json:"-"`
	XXX_unrecognized     []byte                                                  `json:"-"`
	XXX_sizecache        int32                                                   `json:"-"`
}

func (m *OrganizationSettings_AssetDiscoveryConfig) Reset() {
	*m = OrganizationSettings_AssetDiscoveryConfig{}
}
func (m *OrganizationSettings_AssetDiscoveryConfig) String() string { return proto.CompactTextString(m) }
func (*OrganizationSettings_AssetDiscoveryConfig) ProtoMessage()    {}
func (*OrganizationSettings_AssetDiscoveryConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c5f40c5918aa7c6, []int{0, 0}
}

func (m *OrganizationSettings_AssetDiscoveryConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrganizationSettings_AssetDiscoveryConfig.Unmarshal(m, b)
}
func (m *OrganizationSettings_AssetDiscoveryConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrganizationSettings_AssetDiscoveryConfig.Marshal(b, m, deterministic)
}
func (m *OrganizationSettings_AssetDiscoveryConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrganizationSettings_AssetDiscoveryConfig.Merge(m, src)
}
func (m *OrganizationSettings_AssetDiscoveryConfig) XXX_Size() int {
	return xxx_messageInfo_OrganizationSettings_AssetDiscoveryConfig.Size(m)
}
func (m *OrganizationSettings_AssetDiscoveryConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_OrganizationSettings_AssetDiscoveryConfig.DiscardUnknown(m)
}

var xxx_messageInfo_OrganizationSettings_AssetDiscoveryConfig proto.InternalMessageInfo

func (m *OrganizationSettings_AssetDiscoveryConfig) GetProjectIds() []string {
	if m != nil {
		return m.ProjectIds
	}
	return nil
}

func (m *OrganizationSettings_AssetDiscoveryConfig) GetInclusionMode() OrganizationSettings_AssetDiscoveryConfig_InclusionMode {
	if m != nil {
		return m.InclusionMode
	}
	return OrganizationSettings_AssetDiscoveryConfig_INCLUSION_MODE_UNSPECIFIED
}

func init() {
	proto.RegisterEnum("google.cloud.securitycenter.v1p1beta1.OrganizationSettings_AssetDiscoveryConfig_InclusionMode", OrganizationSettings_AssetDiscoveryConfig_InclusionMode_name, OrganizationSettings_AssetDiscoveryConfig_InclusionMode_value)
	proto.RegisterType((*OrganizationSettings)(nil), "google.cloud.securitycenter.v1p1beta1.OrganizationSettings")
	proto.RegisterType((*OrganizationSettings_AssetDiscoveryConfig)(nil), "google.cloud.securitycenter.v1p1beta1.OrganizationSettings.AssetDiscoveryConfig")
}

func init() {
	proto.RegisterFile("google/cloud/securitycenter/v1p1beta1/organization_settings.proto", fileDescriptor_3c5f40c5918aa7c6)
}

var fileDescriptor_3c5f40c5918aa7c6 = []byte{
	// 502 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xd1, 0x6a, 0xdb, 0x3e,
	0x14, 0xc6, 0xff, 0x4a, 0xcb, 0x7f, 0xab, 0xb2, 0x96, 0x20, 0x42, 0xc9, 0xc2, 0xd8, 0x42, 0xa1,
	0x90, 0xdd, 0xc8, 0x38, 0xdb, 0x95, 0x77, 0x95, 0x26, 0xd9, 0x30, 0xb4, 0x49, 0xb0, 0x69, 0xd9,
	0x46, 0x98, 0x51, 0x64, 0x4d, 0xa8, 0x38, 0x92, 0xb1, 0x94, 0x42, 0x37, 0x76, 0xb9, 0xdd, 0xee,
	0x3d, 0xf6, 0x28, 0x7b, 0x91, 0x41, 0xdf, 0x61, 0x30, 0x22, 0xdb, 0x9b, 0x5d, 0x02, 0xeb, 0xc5,
	0xee, 0xac, 0xf3, 0x7d, 0xfa, 0x7d, 0xe7, 0xd8, 0xc7, 0x70, 0xc8, 0x95, 0xe2, 0x09, 0x73, 0x68,
	0xa2, 0xd6, 0xb1, 0xa3, 0x19, 0x5d, 0x67, 0xc2, 0x5c, 0x53, 0x26, 0x0d, 0xcb, 0x9c, 0x2b, 0x37,
	0x75, 0x97, 0xcc, 0x10, 0xd7, 0x51, 0x19, 0x27, 0x52, 0x7c, 0x20, 0x46, 0x28, 0x19, 0x69, 0x66,
	0x8c, 0x90, 0x5c, 0xe3, 0x34, 0x53, 0x46, 0xa1, 0xe3, 0x1c, 0x81, 0x2d, 0x02, 0xd7, 0x11, 0xf8,
	0x37, 0xa2, 0xfb, 0xa8, 0x48, 0x22, 0xa9, 0x70, 0x88, 0x94, 0xca, 0x58, 0x58, 0x01, 0xe9, 0x3e,
	0xac, 0xa8, 0x19, 0xd3, 0x6a, 0x9d, 0x51, 0x96, 0x4b, 0x47, 0x3f, 0x76, 0x61, 0x7b, 0x56, 0xc9,
	0x0f, 0x8b, 0x78, 0x84, 0xe0, 0xae, 0x24, 0x2b, 0xd6, 0x01, 0x3d, 0xd0, 0xdf, 0x0b, 0xec, 0x33,
	0x7a, 0x0e, 0x0f, 0x99, 0x24, 0xcb, 0x84, 0x45, 0x44, 0x6b, 0x66, 0xa2, 0x58, 0x68, 0xaa, 0xae,
	0x58, 0x76, 0xdd, 0x69, 0xf4, 0x40, 0xff, 0x7e, 0xd0, 0xce, 0xd5, 0xe1, 0x46, 0x1c, 0x97, 0x1a,
	0xfa, 0x02, 0xe0, 0xe1, 0x2d, 0x7f, 0x44, 0x95, 0x7c, 0x2f, 0x78, 0x67, 0xa7, 0x07, 0xfa, 0xcd,
	0xc1, 0x1c, 0xdf, 0x69, 0x48, 0xbc, 0xad, 0x4f, 0x5c, 0x0f, 0x1b, 0x59, 0x6e, 0xd0, 0x26, 0x5b,
	0xaa, 0xdd, 0xaf, 0x0d, 0xd8, 0xde, 0x66, 0x47, 0x4f, 0x60, 0x33, 0xcd, 0xd4, 0x25, 0xa3, 0x26,
	0x12, 0xb1, 0xee, 0x80, 0xde, 0x4e, 0x7f, 0x2f, 0x80, 0x45, 0xc9, 0x8f, 0x35, 0xfa, 0x0c, 0xe0,
	0x81, 0x90, 0x34, 0x59, 0xeb, 0xcd, 0x27, 0x5a, 0xa9, 0x98, 0xd9, 0x89, 0x0f, 0x06, 0xef, 0xfe,
	0x75, 0xeb, 0xd8, 0x2f, 0x63, 0xce, 0x54, 0xcc, 0x82, 0x7d, 0x51, 0x3d, 0x1e, 0x4d, 0xe1, 0x7e,
	0x4d, 0x47, 0x8f, 0x61, 0xd7, 0x9f, 0x8e, 0x4e, 0xcf, 0x43, 0x7f, 0x36, 0x8d, 0xce, 0x66, 0xe3,
	0x49, 0x74, 0x3e, 0x0d, 0xe7, 0x93, 0x91, 0xff, 0xd2, 0x9f, 0x8c, 0x5b, 0xff, 0xa1, 0x16, 0x7c,
	0x60, 0xf5, 0xf1, 0x24, 0x9a, 0x4d, 0x4f, 0xdf, 0xb4, 0x00, 0x6a, 0xc2, 0x7b, 0x93, 0xd7, 0xb6,
	0xd2, 0x6a, 0x78, 0x97, 0x37, 0x43, 0x0e, 0x07, 0xb7, 0xba, 0xce, 0x27, 0x22, 0xa9, 0xd0, 0x98,
	0xaa, 0x95, 0xb3, 0x75, 0x3b, 0xdc, 0xea, 0xce, 0x6a, 0xe7, 0x63, 0xf5, 0xf8, 0xa9, 0xb6, 0xd0,
	0xe5, 0x95, 0x93, 0x9f, 0x00, 0x3e, 0xa5, 0x6a, 0x75, 0xb7, 0x17, 0x36, 0x07, 0x6f, 0xc3, 0xc2,
	0xc8, 0x55, 0x42, 0x24, 0xc7, 0x2a, 0xe3, 0x0e, 0x67, 0xd2, 0x6e, 0xad, 0xf3, 0xa7, 0xc5, 0xbf,
	0xfc, 0x5b, 0x2f, 0xea, 0xc2, 0xb7, 0xc6, 0xf1, 0xab, 0x9c, 0x3a, 0xb2, 0xf1, 0x61, 0xa1, 0x8e,
	0xf2, 0xf8, 0x0b, 0x77, 0xee, 0x9e, 0x6c, 0xae, 0x7d, 0x2f, 0x7d, 0x0b, 0xeb, 0x5b, 0xd4, 0x7d,
	0x8b, 0x8b, 0x12, 0x7f, 0xd3, 0xe8, 0xe7, 0x3e, 0xcf, 0xb3, 0x46, 0xcf, 0xab, 0x3b, 0x3d, 0x6f,
	0x63, 0xb5, 0xc8, 0xe5, 0xff, 0xb6, 0xf5, 0x67, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0xc3, 0xe0,
	0xd5, 0x83, 0x15, 0x04, 0x00, 0x00,
}