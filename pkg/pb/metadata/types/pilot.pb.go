// Code generated by protoc-gen-go. DO NOT EDIT.
// source: metadata/types/pilot.proto

package pbtypes // import "openpitrix.io/openpitrix/pkg/pb/metadata/types"

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

type PilotConfig struct {
	Id                     string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	Host                   string   `protobuf:"bytes,2,opt,name=host,proto3" json:"host"`
	ListenPort             int32    `protobuf:"varint,4,opt,name=listen_port,json=listenPort,proto3" json:"listen_port"`
	LogLevel               string   `protobuf:"bytes,5,opt,name=log_level,json=logLevel,proto3" json:"log_level"`
	ForFrontgateListenPort int32    `protobuf:"varint,6,opt,name=for_frontgate_listen_port,json=forFrontgateListenPort,proto3" json:"for_frontgate_listen_port"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *PilotConfig) Reset()         { *m = PilotConfig{} }
func (m *PilotConfig) String() string { return proto.CompactTextString(m) }
func (*PilotConfig) ProtoMessage()    {}
func (*PilotConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_pilot_59c0ea2cea906a34, []int{0}
}
func (m *PilotConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PilotConfig.Unmarshal(m, b)
}
func (m *PilotConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PilotConfig.Marshal(b, m, deterministic)
}
func (dst *PilotConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PilotConfig.Merge(dst, src)
}
func (m *PilotConfig) XXX_Size() int {
	return xxx_messageInfo_PilotConfig.Size(m)
}
func (m *PilotConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_PilotConfig.DiscardUnknown(m)
}

var xxx_messageInfo_PilotConfig proto.InternalMessageInfo

func (m *PilotConfig) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PilotConfig) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *PilotConfig) GetListenPort() int32 {
	if m != nil {
		return m.ListenPort
	}
	return 0
}

func (m *PilotConfig) GetLogLevel() string {
	if m != nil {
		return m.LogLevel
	}
	return ""
}

func (m *PilotConfig) GetForFrontgateListenPort() int32 {
	if m != nil {
		return m.ForFrontgateListenPort
	}
	return 0
}

type PilotClientTLSConfig struct {
	CaCrtData            string   `protobuf:"bytes,1,opt,name=ca_crt_data,json=caCrtData,proto3" json:"ca_crt_data"`
	ClientCrtData        string   `protobuf:"bytes,2,opt,name=client_crt_data,json=clientCrtData,proto3" json:"client_crt_data"`
	ClientKeyData        string   `protobuf:"bytes,3,opt,name=client_key_data,json=clientKeyData,proto3" json:"client_key_data"`
	PilotServerName      string   `protobuf:"bytes,4,opt,name=pilot_server_name,json=pilotServerName,proto3" json:"pilot_server_name"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PilotClientTLSConfig) Reset()         { *m = PilotClientTLSConfig{} }
func (m *PilotClientTLSConfig) String() string { return proto.CompactTextString(m) }
func (*PilotClientTLSConfig) ProtoMessage()    {}
func (*PilotClientTLSConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_pilot_59c0ea2cea906a34, []int{1}
}
func (m *PilotClientTLSConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PilotClientTLSConfig.Unmarshal(m, b)
}
func (m *PilotClientTLSConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PilotClientTLSConfig.Marshal(b, m, deterministic)
}
func (dst *PilotClientTLSConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PilotClientTLSConfig.Merge(dst, src)
}
func (m *PilotClientTLSConfig) XXX_Size() int {
	return xxx_messageInfo_PilotClientTLSConfig.Size(m)
}
func (m *PilotClientTLSConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_PilotClientTLSConfig.DiscardUnknown(m)
}

var xxx_messageInfo_PilotClientTLSConfig proto.InternalMessageInfo

func (m *PilotClientTLSConfig) GetCaCrtData() string {
	if m != nil {
		return m.CaCrtData
	}
	return ""
}

func (m *PilotClientTLSConfig) GetClientCrtData() string {
	if m != nil {
		return m.ClientCrtData
	}
	return ""
}

func (m *PilotClientTLSConfig) GetClientKeyData() string {
	if m != nil {
		return m.ClientKeyData
	}
	return ""
}

func (m *PilotClientTLSConfig) GetPilotServerName() string {
	if m != nil {
		return m.PilotServerName
	}
	return ""
}

type PilotEndpoint struct {
	PilotId              string   `protobuf:"bytes,1,opt,name=pilot_id,json=pilotId,proto3" json:"pilot_id"`
	PilotHost            string   `protobuf:"bytes,2,opt,name=pilot_host,json=pilotHost,proto3" json:"pilot_host"`
	PilotPort            int32    `protobuf:"varint,3,opt,name=pilot_port,json=pilotPort,proto3" json:"pilot_port"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PilotEndpoint) Reset()         { *m = PilotEndpoint{} }
func (m *PilotEndpoint) String() string { return proto.CompactTextString(m) }
func (*PilotEndpoint) ProtoMessage()    {}
func (*PilotEndpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_pilot_59c0ea2cea906a34, []int{2}
}
func (m *PilotEndpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PilotEndpoint.Unmarshal(m, b)
}
func (m *PilotEndpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PilotEndpoint.Marshal(b, m, deterministic)
}
func (dst *PilotEndpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PilotEndpoint.Merge(dst, src)
}
func (m *PilotEndpoint) XXX_Size() int {
	return xxx_messageInfo_PilotEndpoint.Size(m)
}
func (m *PilotEndpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_PilotEndpoint.DiscardUnknown(m)
}

var xxx_messageInfo_PilotEndpoint proto.InternalMessageInfo

func (m *PilotEndpoint) GetPilotId() string {
	if m != nil {
		return m.PilotId
	}
	return ""
}

func (m *PilotEndpoint) GetPilotHost() string {
	if m != nil {
		return m.PilotHost
	}
	return ""
}

func (m *PilotEndpoint) GetPilotPort() int32 {
	if m != nil {
		return m.PilotPort
	}
	return 0
}

func init() {
	proto.RegisterType((*PilotConfig)(nil), "metadata.types.PilotConfig")
	proto.RegisterType((*PilotClientTLSConfig)(nil), "metadata.types.PilotClientTLSConfig")
	proto.RegisterType((*PilotEndpoint)(nil), "metadata.types.PilotEndpoint")
}

func init() { proto.RegisterFile("metadata/types/pilot.proto", fileDescriptor_pilot_59c0ea2cea906a34) }

var fileDescriptor_pilot_59c0ea2cea906a34 = []byte{
	// 363 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0x4f, 0x6b, 0xdb, 0x40,
	0x10, 0xc5, 0x91, 0xff, 0xd5, 0x1a, 0x63, 0x9b, 0x2e, 0xa5, 0xc8, 0x2d, 0x6d, 0x8d, 0x0f, 0xc5,
	0xf4, 0x60, 0x1d, 0x0a, 0xa5, 0xa5, 0xb7, 0xba, 0x2d, 0x09, 0x31, 0xc1, 0xd8, 0x39, 0xe5, 0xb2,
	0xac, 0xa5, 0x95, 0xbc, 0x58, 0xde, 0x5d, 0x56, 0x83, 0x89, 0xbf, 0x52, 0x8e, 0xf9, 0x84, 0x41,
	0xb3, 0xfe, 0x17, 0x72, 0x92, 0xe6, 0xbd, 0x37, 0xc3, 0xec, 0x8f, 0x81, 0x0f, 0x5b, 0x89, 0x22,
	0x15, 0x28, 0x62, 0xdc, 0x5b, 0x59, 0xc6, 0x56, 0x15, 0x06, 0x27, 0xd6, 0x19, 0x34, 0xac, 0x77,
	0xf4, 0x26, 0xe4, 0x8d, 0x1e, 0x03, 0xe8, 0xcc, 0x2b, 0x7f, 0x6a, 0x74, 0xa6, 0x72, 0xd6, 0x83,
	0x9a, 0x4a, 0xa3, 0x60, 0x18, 0x8c, 0xc3, 0x45, 0x4d, 0xa5, 0x8c, 0x41, 0x63, 0x6d, 0x4a, 0x8c,
	0x6a, 0xa4, 0xd0, 0x3f, 0xfb, 0x02, 0x9d, 0x42, 0x95, 0x28, 0x35, 0xb7, 0xc6, 0x61, 0xd4, 0x18,
	0x06, 0xe3, 0xe6, 0x02, 0xbc, 0x34, 0x37, 0x0e, 0xd9, 0x47, 0x08, 0x0b, 0x93, 0xf3, 0x42, 0xee,
	0x64, 0x11, 0x35, 0xa9, 0xb3, 0x5d, 0x98, 0x7c, 0x56, 0xd5, 0xec, 0x17, 0x0c, 0x32, 0xe3, 0x78,
	0xe6, 0x8c, 0xc6, 0x5c, 0xa0, 0xe4, 0x97, 0xb3, 0x5a, 0x34, 0xeb, 0x7d, 0x66, 0xdc, 0xff, 0xa3,
	0x3f, 0x3b, 0xcd, 0x1d, 0x3d, 0x05, 0xf0, 0xce, 0x2f, 0x5b, 0x28, 0xa9, 0xf1, 0x6e, 0xb6, 0x3c,
	0x6c, 0xfd, 0x19, 0x3a, 0x89, 0xe0, 0x89, 0x43, 0x5e, 0x3d, 0xed, 0xb0, 0x7e, 0x98, 0x88, 0xa9,
	0xc3, 0xbf, 0x02, 0x05, 0xfb, 0x0a, 0xfd, 0x84, 0x5a, 0xce, 0x19, 0xff, 0xa0, 0xae, 0x97, 0x5f,
	0xe7, 0x36, 0x72, 0xef, 0x73, 0xf5, 0xcb, 0xdc, 0x8d, 0xdc, 0x53, 0xee, 0x1b, 0xbc, 0x25, 0xa8,
	0xbc, 0x94, 0x6e, 0x27, 0x1d, 0xd7, 0x62, 0x2b, 0x89, 0x43, 0xb8, 0xe8, 0x93, 0xb1, 0x24, 0xfd,
	0x56, 0x6c, 0xe5, 0x68, 0x0d, 0x5d, 0xda, 0xf9, 0x9f, 0x4e, 0xad, 0x51, 0x1a, 0xd9, 0x00, 0xda,
	0xbe, 0xf9, 0x04, 0xfa, 0x0d, 0xd5, 0xd7, 0x29, 0xfb, 0x04, 0xe0, 0xad, 0x0b, 0xe6, 0x21, 0x29,
	0x57, 0x15, 0xf8, 0x93, 0x4d, 0xac, 0xea, 0xc4, 0xca, 0xdb, 0x15, 0x9e, 0x3f, 0x3f, 0xef, 0x7f,
	0x18, 0x2b, 0xb5, 0x55, 0xe8, 0xd4, 0xc3, 0x44, 0x99, 0xf8, 0x5c, 0xc5, 0x76, 0x93, 0xc7, 0x76,
	0x15, 0xbf, 0xbc, 0x8c, 0xdf, 0x76, 0x45, 0xdf, 0x55, 0x8b, 0x8e, 0xe3, 0xfb, 0x73, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x47, 0x9a, 0x62, 0x1a, 0x3a, 0x02, 0x00, 0x00,
}
