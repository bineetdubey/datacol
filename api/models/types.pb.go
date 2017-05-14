// Code generated by protoc-gen-go.
// source: types.proto
// DO NOT EDIT!

/*
Package models is a generated protocol buffer package.

It is generated from these files:
	types.proto

It has these top-level messages:
	App
	Build
	Release
	Resource
	ResourceVar
	EnvConfig
*/
package models

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "github.com/golang/protobuf/ptypes/timestamp"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Status int32

const (
	Status_CREATED Status = 0
)

var Status_name = map[int32]string{
	0: "CREATED",
}
var Status_value = map[string]int32{
	"CREATED": 0,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}
func (Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type App struct {
	Name      string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Status    Status `protobuf:"varint,2,opt,name=status,enum=models.Status" json:"status,omitempty"`
	ReleaseId string `protobuf:"bytes,3,opt,name=release_id,json=releaseId" json:"release_id,omitempty"`
	Endpoint  string `protobuf:"bytes,4,opt,name=endpoint" json:"endpoint,omitempty"`
}

func (m *App) Reset()                    { *m = App{} }
func (m *App) String() string            { return proto.CompactTextString(m) }
func (*App) ProtoMessage()               {}
func (*App) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *App) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *App) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_CREATED
}

func (m *App) GetReleaseId() string {
	if m != nil {
		return m.ReleaseId
	}
	return ""
}

func (m *App) GetEndpoint() string {
	if m != nil {
		return m.Endpoint
	}
	return ""
}

type Build struct {
	Id        string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	App       string `protobuf:"bytes,2,opt,name=app" json:"app,omitempty"`
	RemoteId  string `protobuf:"bytes,3,opt,name=remote_id,json=remoteId" json:"remote_id,omitempty"`
	Status    Status `protobuf:"varint,4,opt,name=status,enum=models.Status" json:"status,omitempty"`
	CreatedAt int32  `protobuf:"varint,5,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
}

func (m *Build) Reset()                    { *m = Build{} }
func (m *Build) String() string            { return proto.CompactTextString(m) }
func (*Build) ProtoMessage()               {}
func (*Build) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Build) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Build) GetApp() string {
	if m != nil {
		return m.App
	}
	return ""
}

func (m *Build) GetRemoteId() string {
	if m != nil {
		return m.RemoteId
	}
	return ""
}

func (m *Build) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_CREATED
}

func (m *Build) GetCreatedAt() int32 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

type Release struct {
	Id        string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	App       string `protobuf:"bytes,2,opt,name=app" json:"app,omitempty"`
	BuildId   string `protobuf:"bytes,3,opt,name=build_id,json=buildId" json:"build_id,omitempty"`
	Status    Status `protobuf:"varint,4,opt,name=status,enum=models.Status" json:"status,omitempty"`
	CreatedAt int32  `protobuf:"varint,5,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
}

func (m *Release) Reset()                    { *m = Release{} }
func (m *Release) String() string            { return proto.CompactTextString(m) }
func (*Release) ProtoMessage()               {}
func (*Release) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Release) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Release) GetApp() string {
	if m != nil {
		return m.App
	}
	return ""
}

func (m *Release) GetBuildId() string {
	if m != nil {
		return m.BuildId
	}
	return ""
}

func (m *Release) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_CREATED
}

func (m *Release) GetCreatedAt() int32 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

type Resource struct {
	Name       string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Status     Status   `protobuf:"varint,2,opt,name=status,enum=models.Status" json:"status,omitempty"`
	Kind       string   `protobuf:"bytes,3,opt,name=kind" json:"kind,omitempty"`
	URL        string   `protobuf:"bytes,4,opt,name=URL" json:"URL,omitempty"`
	Apps       []string `protobuf:"bytes,5,rep,name=apps" json:"apps,omitempty"`
	Exports    []byte   `protobuf:"bytes,6,opt,name=exports,proto3" json:"exports,omitempty"`
	Parameters []byte   `protobuf:"bytes,7,opt,name=parameters,proto3" json:"parameters,omitempty"`
}

func (m *Resource) Reset()                    { *m = Resource{} }
func (m *Resource) String() string            { return proto.CompactTextString(m) }
func (*Resource) ProtoMessage()               {}
func (*Resource) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Resource) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Resource) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_CREATED
}

func (m *Resource) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *Resource) GetURL() string {
	if m != nil {
		return m.URL
	}
	return ""
}

func (m *Resource) GetApps() []string {
	if m != nil {
		return m.Apps
	}
	return nil
}

func (m *Resource) GetExports() []byte {
	if m != nil {
		return m.Exports
	}
	return nil
}

func (m *Resource) GetParameters() []byte {
	if m != nil {
		return m.Parameters
	}
	return nil
}

type ResourceVar struct {
	Key   string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *ResourceVar) Reset()                    { *m = ResourceVar{} }
func (m *ResourceVar) String() string            { return proto.CompactTextString(m) }
func (*ResourceVar) ProtoMessage()               {}
func (*ResourceVar) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ResourceVar) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ResourceVar) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type EnvConfig struct {
	Data map[string]string `protobuf:"bytes,1,rep,name=data" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *EnvConfig) Reset()                    { *m = EnvConfig{} }
func (m *EnvConfig) String() string            { return proto.CompactTextString(m) }
func (*EnvConfig) ProtoMessage()               {}
func (*EnvConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *EnvConfig) GetData() map[string]string {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*App)(nil), "models.App")
	proto.RegisterType((*Build)(nil), "models.Build")
	proto.RegisterType((*Release)(nil), "models.Release")
	proto.RegisterType((*Resource)(nil), "models.Resource")
	proto.RegisterType((*ResourceVar)(nil), "models.ResourceVar")
	proto.RegisterType((*EnvConfig)(nil), "models.EnvConfig")
	proto.RegisterEnum("models.Status", Status_name, Status_value)
}

func init() { proto.RegisterFile("types.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 616 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x94, 0xcd, 0x6e, 0xd4, 0x3c,
	0x14, 0x86, 0xbf, 0x64, 0x7e, 0x73, 0xe6, 0x53, 0xa9, 0xdc, 0x69, 0x09, 0x03, 0x6a, 0x46, 0x96,
	0x90, 0x22, 0x24, 0x66, 0x44, 0xcb, 0x6f, 0x17, 0x48, 0x4d, 0x3b, 0x8b, 0x4a, 0x5d, 0x19, 0xca,
	0x82, 0x4d, 0xe5, 0x69, 0xdc, 0x10, 0x75, 0x12, 0x5b, 0x89, 0x53, 0xd1, 0x3d, 0xdc, 0x1b, 0x4b,
	0xae, 0x20, 0xaa, 0x7a, 0x09, 0xb9, 0x02, 0x64, 0x27, 0xe9, 0x84, 0x81, 0x05, 0x42, 0x08, 0x76,
	0xe7, 0xf8, 0x7d, 0x8e, 0x4f, 0xce, 0x6b, 0xc7, 0x30, 0x90, 0x57, 0x82, 0xa5, 0x13, 0x91, 0x70,
	0xc9, 0x51, 0x37, 0xe2, 0x3e, 0x5b, 0xa4, 0xa3, 0x07, 0x01, 0xe7, 0xc1, 0x82, 0x4d, 0xa9, 0x08,
	0xa7, 0x34, 0x8e, 0xb9, 0xa4, 0x32, 0xe4, 0x71, 0x45, 0x8d, 0x9c, 0x4a, 0xd5, 0xd9, 0x3c, 0x3b,
	0x9f, 0xca, 0x30, 0x62, 0xa9, 0xa4, 0x91, 0xa8, 0x80, 0xc7, 0x41, 0x28, 0x3f, 0x64, 0xf3, 0xc9,
	0x19, 0x8f, 0xa6, 0x01, 0x0f, 0xf8, 0x92, 0x54, 0x99, 0x4e, 0x74, 0x54, 0xe2, 0xf8, 0xda, 0x80,
	0xd6, 0xbe, 0x10, 0xc8, 0x85, 0x76, 0x4c, 0x23, 0x66, 0x1b, 0x63, 0xc3, 0xb5, 0xbc, 0x61, 0x91,
	0x3b, 0xeb, 0x3e, 0x95, 0x34, 0x95, 0x3c, 0x61, 0x7b, 0x58, 0x49, 0x98, 0x68, 0x02, 0xbd, 0x86,
	0x6e, 0x2a, 0xa9, 0xcc, 0x52, 0xdb, 0x1c, 0x1b, 0xee, 0xda, 0xce, 0xda, 0xa4, 0xfc, 0xf0, 0xc9,
	0x1b, 0xbd, 0xea, 0x6d, 0x15, 0xb9, 0x83, 0x1a, 0xb5, 0x25, 0x8c, 0x49, 0x55, 0x85, 0x5e, 0x01,
	0x24, 0x6c, 0xc1, 0x68, 0xca, 0x4e, 0x43, 0xdf, 0x6e, 0xe9, 0x7e, 0xa3, 0x22, 0x77, 0xb6, 0x1a,
	0x35, 0x4b, 0x00, 0x13, 0xab, 0x4a, 0x8e, 0x7c, 0xf4, 0x14, 0xfa, 0x2c, 0xf6, 0x05, 0x0f, 0x63,
	0x69, 0xb7, 0x75, 0xa1, 0x5d, 0xe4, 0xce, 0xb0, 0x51, 0x58, 0xcb, 0x98, 0xdc, 0x92, 0xf8, 0xb3,
	0x09, 0x1d, 0x2f, 0x0b, 0x17, 0x3e, 0xc2, 0x60, 0x86, 0x7e, 0x35, 0x22, 0x2a, 0x72, 0x67, 0xad,
	0x51, 0xa9, 0x5a, 0x99, 0xa1, 0x8f, 0x1e, 0x42, 0x8b, 0x0a, 0xa1, 0x67, 0xb3, 0xbc, 0x8d, 0x22,
	0x77, 0xee, 0x34, 0x20, 0x2a, 0x04, 0x26, 0x4a, 0x47, 0xcf, 0xc1, 0x4a, 0x58, 0xc4, 0x65, 0x63,
	0x88, 0x7b, 0x45, 0xee, 0x6c, 0x7e, 0x37, 0x44, 0xa5, 0x63, 0xd2, 0x2f, 0xe3, 0x23, 0xbf, 0xe1,
	0x5e, 0xfb, 0x77, 0xdd, 0x3b, 0x4b, 0x18, 0x95, 0xcc, 0x3f, 0xa5, 0xd2, 0xee, 0x8c, 0x0d, 0xb7,
	0xf3, 0x83, 0x7b, 0x4b, 0x00, 0x13, 0xab, 0x4a, 0xf6, 0x25, 0xfe, 0x64, 0x42, 0x8f, 0x94, 0x5e,
	0xfe, 0x49, 0x27, 0x76, 0xa1, 0x3f, 0x57, 0xee, 0x2e, 0x8d, 0x58, 0x3d, 0x94, 0x5a, 0xc6, 0xa4,
	0xa7, 0xc3, 0x7f, 0x6b, 0x43, 0x6e, 0x42, 0x9f, 0xb0, 0x94, 0x67, 0xc9, 0x19, 0xfb, 0x8b, 0xd7,
	0xde, 0x85, 0xf6, 0x45, 0x18, 0xd7, 0x16, 0xad, 0x76, 0x52, 0x12, 0x26, 0x9a, 0x50, 0xbe, 0x9f,
	0x90, 0xe3, 0xea, 0x82, 0xaf, 0xfa, 0x7e, 0x42, 0x8e, 0x31, 0x51, 0xba, 0xda, 0x90, 0x0a, 0x91,
	0xda, 0x9d, 0x71, 0xeb, 0x27, 0x1b, 0x2a, 0x09, 0x13, 0x4d, 0xa0, 0x27, 0xd0, 0x63, 0x1f, 0x05,
	0x4f, 0x64, 0x6a, 0x77, 0xc7, 0x86, 0xfb, 0xbf, 0x77, 0xb7, 0xc8, 0x9d, 0x8d, 0xe6, 0x5f, 0x53,
	0xaa, 0x98, 0xd4, 0x1c, 0xda, 0x03, 0x10, 0x34, 0xa1, 0x11, 0x93, 0x2c, 0x49, 0xed, 0x9e, 0xae,
	0x5a, 0xf5, 0x77, 0x09, 0x60, 0xd2, 0xa0, 0xf1, 0x33, 0x18, 0xd4, 0xfe, 0xbe, 0xa3, 0x09, 0x5a,
	0x87, 0xd6, 0x05, 0xbb, 0x2a, 0x1d, 0x26, 0x2a, 0x44, 0x43, 0xe8, 0x5c, 0xd2, 0x45, 0xc6, 0xca,
	0xab, 0x45, 0xca, 0x04, 0x67, 0x60, 0xcd, 0xe2, 0xcb, 0x03, 0x1e, 0x9f, 0x87, 0x01, 0x9a, 0x42,
	0x5b, 0x75, 0xb2, 0x8d, 0x71, 0xcb, 0x1d, 0xec, 0xdc, 0xaf, 0xbd, 0xbe, 0x05, 0x26, 0x87, 0x54,
	0xd2, 0x59, 0x2c, 0x93, 0x2b, 0xa2, 0xc1, 0xd1, 0x0b, 0xb0, 0x6e, 0x97, 0x7e, 0xb5, 0xe5, 0x9e,
	0xf9, 0xd2, 0x78, 0xb4, 0x09, 0xdd, 0xf2, 0x04, 0xd1, 0x00, 0x7a, 0x07, 0x64, 0xb6, 0xff, 0x76,
	0x76, 0xb8, 0xfe, 0x9f, 0x37, 0xfc, 0x72, 0xb3, 0x6d, 0x7c, 0xbd, 0xd9, 0x36, 0xae, 0x6f, 0xb6,
	0x8d, 0xf7, 0xd5, 0xdb, 0x3c, 0xef, 0xea, 0x47, 0x73, 0xf7, 0x5b, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x73, 0xa7, 0x48, 0xed, 0xb9, 0x05, 0x00, 0x00,
}
