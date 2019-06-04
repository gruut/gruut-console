// Code generated by protoc-gen-go. DO NOT EDIT.
// source: admin_service.proto

package grpc_admin

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type ReqStart_Mode int32

const (
	ReqStart_DEFAULT ReqStart_Mode = 0
	ReqStart_MONITOR ReqStart_Mode = 1
)

var ReqStart_Mode_name = map[int32]string{
	0: "DEFAULT",
	1: "MONITOR",
}

var ReqStart_Mode_value = map[string]int32{
	"DEFAULT": 0,
	"MONITOR": 1,
}

func (x ReqStart_Mode) String() string {
	return proto.EnumName(ReqStart_Mode_name, int32(x))
}

func (ReqStart_Mode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f9bc00259dc38d22, []int{4, 0}
}

//TODO : 각 메시지별  내부 스펙들은 변경 될 수 있음.
type ReqSetupKey struct {
	SetupPort            string   `protobuf:"bytes,1,opt,name=setup_port,json=setupPort,proto3" json:"setup_port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqSetupKey) Reset()         { *m = ReqSetupKey{} }
func (m *ReqSetupKey) String() string { return proto.CompactTextString(m) }
func (*ReqSetupKey) ProtoMessage()    {}
func (*ReqSetupKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9bc00259dc38d22, []int{0}
}

func (m *ReqSetupKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqSetupKey.Unmarshal(m, b)
}
func (m *ReqSetupKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqSetupKey.Marshal(b, m, deterministic)
}
func (m *ReqSetupKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqSetupKey.Merge(m, src)
}
func (m *ReqSetupKey) XXX_Size() int {
	return xxx_messageInfo_ReqSetupKey.Size(m)
}
func (m *ReqSetupKey) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqSetupKey.DiscardUnknown(m)
}

var xxx_messageInfo_ReqSetupKey proto.InternalMessageInfo

func (m *ReqSetupKey) GetSetupPort() string {
	if m != nil {
		return m.SetupPort
	}
	return ""
}

type ResSetupKey struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Info                 string   `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResSetupKey) Reset()         { *m = ResSetupKey{} }
func (m *ResSetupKey) String() string { return proto.CompactTextString(m) }
func (*ResSetupKey) ProtoMessage()    {}
func (*ResSetupKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9bc00259dc38d22, []int{1}
}

func (m *ResSetupKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResSetupKey.Unmarshal(m, b)
}
func (m *ResSetupKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResSetupKey.Marshal(b, m, deterministic)
}
func (m *ResSetupKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResSetupKey.Merge(m, src)
}
func (m *ResSetupKey) XXX_Size() int {
	return xxx_messageInfo_ResSetupKey.Size(m)
}
func (m *ResSetupKey) XXX_DiscardUnknown() {
	xxx_messageInfo_ResSetupKey.DiscardUnknown(m)
}

var xxx_messageInfo_ResSetupKey proto.InternalMessageInfo

func (m *ResSetupKey) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *ResSetupKey) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

type ReqLogin struct {
	Password             string   `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqLogin) Reset()         { *m = ReqLogin{} }
func (m *ReqLogin) String() string { return proto.CompactTextString(m) }
func (*ReqLogin) ProtoMessage()    {}
func (*ReqLogin) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9bc00259dc38d22, []int{2}
}

func (m *ReqLogin) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqLogin.Unmarshal(m, b)
}
func (m *ReqLogin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqLogin.Marshal(b, m, deterministic)
}
func (m *ReqLogin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqLogin.Merge(m, src)
}
func (m *ReqLogin) XXX_Size() int {
	return xxx_messageInfo_ReqLogin.Size(m)
}
func (m *ReqLogin) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqLogin.DiscardUnknown(m)
}

var xxx_messageInfo_ReqLogin proto.InternalMessageInfo

func (m *ReqLogin) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type ResLogin struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Info                 string   `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResLogin) Reset()         { *m = ResLogin{} }
func (m *ResLogin) String() string { return proto.CompactTextString(m) }
func (*ResLogin) ProtoMessage()    {}
func (*ResLogin) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9bc00259dc38d22, []int{3}
}

func (m *ResLogin) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResLogin.Unmarshal(m, b)
}
func (m *ResLogin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResLogin.Marshal(b, m, deterministic)
}
func (m *ResLogin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResLogin.Merge(m, src)
}
func (m *ResLogin) XXX_Size() int {
	return xxx_messageInfo_ResLogin.Size(m)
}
func (m *ResLogin) XXX_DiscardUnknown() {
	xxx_messageInfo_ResLogin.DiscardUnknown(m)
}

var xxx_messageInfo_ResLogin proto.InternalMessageInfo

func (m *ResLogin) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *ResLogin) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

type ReqStart struct {
	Mode                 ReqStart_Mode `protobuf:"varint,1,opt,name=mode,proto3,enum=grpc_admin.ReqStart_Mode" json:"mode,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ReqStart) Reset()         { *m = ReqStart{} }
func (m *ReqStart) String() string { return proto.CompactTextString(m) }
func (*ReqStart) ProtoMessage()    {}
func (*ReqStart) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9bc00259dc38d22, []int{4}
}

func (m *ReqStart) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqStart.Unmarshal(m, b)
}
func (m *ReqStart) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqStart.Marshal(b, m, deterministic)
}
func (m *ReqStart) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqStart.Merge(m, src)
}
func (m *ReqStart) XXX_Size() int {
	return xxx_messageInfo_ReqStart.Size(m)
}
func (m *ReqStart) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqStart.DiscardUnknown(m)
}

var xxx_messageInfo_ReqStart proto.InternalMessageInfo

func (m *ReqStart) GetMode() ReqStart_Mode {
	if m != nil {
		return m.Mode
	}
	return ReqStart_DEFAULT
}

type ResStart struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Info                 string   `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResStart) Reset()         { *m = ResStart{} }
func (m *ResStart) String() string { return proto.CompactTextString(m) }
func (*ResStart) ProtoMessage()    {}
func (*ResStart) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9bc00259dc38d22, []int{5}
}

func (m *ResStart) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResStart.Unmarshal(m, b)
}
func (m *ResStart) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResStart.Marshal(b, m, deterministic)
}
func (m *ResStart) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResStart.Merge(m, src)
}
func (m *ResStart) XXX_Size() int {
	return xxx_messageInfo_ResStart.Size(m)
}
func (m *ResStart) XXX_DiscardUnknown() {
	xxx_messageInfo_ResStart.DiscardUnknown(m)
}

var xxx_messageInfo_ResStart proto.InternalMessageInfo

func (m *ResStart) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *ResStart) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

type ReqStatus struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqStatus) Reset()         { *m = ReqStatus{} }
func (m *ReqStatus) String() string { return proto.CompactTextString(m) }
func (*ReqStatus) ProtoMessage()    {}
func (*ReqStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9bc00259dc38d22, []int{6}
}

func (m *ReqStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqStatus.Unmarshal(m, b)
}
func (m *ReqStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqStatus.Marshal(b, m, deterministic)
}
func (m *ReqStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqStatus.Merge(m, src)
}
func (m *ReqStatus) XXX_Size() int {
	return xxx_messageInfo_ReqStatus.Size(m)
}
func (m *ReqStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ReqStatus proto.InternalMessageInfo

type ResStatus struct {
	Alive                bool     `protobuf:"varint,1,opt,name=alive,proto3" json:"alive,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResStatus) Reset()         { *m = ResStatus{} }
func (m *ResStatus) String() string { return proto.CompactTextString(m) }
func (*ResStatus) ProtoMessage()    {}
func (*ResStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9bc00259dc38d22, []int{7}
}

func (m *ResStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResStatus.Unmarshal(m, b)
}
func (m *ResStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResStatus.Marshal(b, m, deterministic)
}
func (m *ResStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResStatus.Merge(m, src)
}
func (m *ResStatus) XXX_Size() int {
	return xxx_messageInfo_ResStatus.Size(m)
}
func (m *ResStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ResStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ResStatus proto.InternalMessageInfo

func (m *ResStatus) GetAlive() bool {
	if m != nil {
		return m.Alive
	}
	return false
}

type ReqLoadWorld struct {
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqLoadWorld) Reset()         { *m = ReqLoadWorld{} }
func (m *ReqLoadWorld) String() string { return proto.CompactTextString(m) }
func (*ReqLoadWorld) ProtoMessage()    {}
func (*ReqLoadWorld) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9bc00259dc38d22, []int{8}
}

func (m *ReqLoadWorld) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqLoadWorld.Unmarshal(m, b)
}
func (m *ReqLoadWorld) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqLoadWorld.Marshal(b, m, deterministic)
}
func (m *ReqLoadWorld) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqLoadWorld.Merge(m, src)
}
func (m *ReqLoadWorld) XXX_Size() int {
	return xxx_messageInfo_ReqLoadWorld.Size(m)
}
func (m *ReqLoadWorld) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqLoadWorld.DiscardUnknown(m)
}

var xxx_messageInfo_ReqLoadWorld proto.InternalMessageInfo

func (m *ReqLoadWorld) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type ResLoadWorld struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Info                 string   `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResLoadWorld) Reset()         { *m = ResLoadWorld{} }
func (m *ResLoadWorld) String() string { return proto.CompactTextString(m) }
func (*ResLoadWorld) ProtoMessage()    {}
func (*ResLoadWorld) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9bc00259dc38d22, []int{9}
}

func (m *ResLoadWorld) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResLoadWorld.Unmarshal(m, b)
}
func (m *ResLoadWorld) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResLoadWorld.Marshal(b, m, deterministic)
}
func (m *ResLoadWorld) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResLoadWorld.Merge(m, src)
}
func (m *ResLoadWorld) XXX_Size() int {
	return xxx_messageInfo_ResLoadWorld.Size(m)
}
func (m *ResLoadWorld) XXX_DiscardUnknown() {
	xxx_messageInfo_ResLoadWorld.DiscardUnknown(m)
}

var xxx_messageInfo_ResLoadWorld proto.InternalMessageInfo

func (m *ResLoadWorld) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *ResLoadWorld) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

type ReqLoadChain struct {
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqLoadChain) Reset()         { *m = ReqLoadChain{} }
func (m *ReqLoadChain) String() string { return proto.CompactTextString(m) }
func (*ReqLoadChain) ProtoMessage()    {}
func (*ReqLoadChain) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9bc00259dc38d22, []int{10}
}

func (m *ReqLoadChain) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqLoadChain.Unmarshal(m, b)
}
func (m *ReqLoadChain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqLoadChain.Marshal(b, m, deterministic)
}
func (m *ReqLoadChain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqLoadChain.Merge(m, src)
}
func (m *ReqLoadChain) XXX_Size() int {
	return xxx_messageInfo_ReqLoadChain.Size(m)
}
func (m *ReqLoadChain) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqLoadChain.DiscardUnknown(m)
}

var xxx_messageInfo_ReqLoadChain proto.InternalMessageInfo

func (m *ReqLoadChain) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type ResLoadChain struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Info                 string   `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResLoadChain) Reset()         { *m = ResLoadChain{} }
func (m *ResLoadChain) String() string { return proto.CompactTextString(m) }
func (*ResLoadChain) ProtoMessage()    {}
func (*ResLoadChain) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9bc00259dc38d22, []int{11}
}

func (m *ResLoadChain) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResLoadChain.Unmarshal(m, b)
}
func (m *ResLoadChain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResLoadChain.Marshal(b, m, deterministic)
}
func (m *ResLoadChain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResLoadChain.Merge(m, src)
}
func (m *ResLoadChain) XXX_Size() int {
	return xxx_messageInfo_ResLoadChain.Size(m)
}
func (m *ResLoadChain) XXX_DiscardUnknown() {
	xxx_messageInfo_ResLoadChain.DiscardUnknown(m)
}

var xxx_messageInfo_ResLoadChain proto.InternalMessageInfo

func (m *ResLoadChain) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *ResLoadChain) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

func init() {
	proto.RegisterEnum("grpc_admin.ReqStart_Mode", ReqStart_Mode_name, ReqStart_Mode_value)
	proto.RegisterType((*ReqSetupKey)(nil), "grpc_admin.ReqSetupKey")
	proto.RegisterType((*ResSetupKey)(nil), "grpc_admin.ResSetupKey")
	proto.RegisterType((*ReqLogin)(nil), "grpc_admin.ReqLogin")
	proto.RegisterType((*ResLogin)(nil), "grpc_admin.ResLogin")
	proto.RegisterType((*ReqStart)(nil), "grpc_admin.ReqStart")
	proto.RegisterType((*ResStart)(nil), "grpc_admin.ResStart")
	proto.RegisterType((*ReqStatus)(nil), "grpc_admin.ReqStatus")
	proto.RegisterType((*ResStatus)(nil), "grpc_admin.ResStatus")
	proto.RegisterType((*ReqLoadWorld)(nil), "grpc_admin.ReqLoadWorld")
	proto.RegisterType((*ResLoadWorld)(nil), "grpc_admin.ResLoadWorld")
	proto.RegisterType((*ReqLoadChain)(nil), "grpc_admin.ReqLoadChain")
	proto.RegisterType((*ResLoadChain)(nil), "grpc_admin.ResLoadChain")
}

func init() { proto.RegisterFile("admin_service.proto", fileDescriptor_f9bc00259dc38d22) }

var fileDescriptor_f9bc00259dc38d22 = []byte{
	// 425 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x4b, 0xcb, 0xd3, 0x40,
	0x14, 0x4d, 0xb5, 0xb5, 0xc9, 0x8d, 0x48, 0x1d, 0x5b, 0x8c, 0x01, 0xa1, 0xce, 0x42, 0x5c, 0x68,
	0x16, 0x15, 0x41, 0xf0, 0x01, 0xa5, 0x3e, 0x10, 0x5b, 0x2b, 0x69, 0xc5, 0x85, 0x8b, 0x12, 0x93,
	0xb1, 0x0d, 0xb6, 0x99, 0x74, 0x66, 0x52, 0xf1, 0x0f, 0xf8, 0xbb, 0x65, 0x1e, 0x4d, 0x3e, 0x92,
	0xc2, 0x47, 0x76, 0x73, 0xef, 0x39, 0xf7, 0x9c, 0x99, 0x9b, 0x13, 0xb8, 0x17, 0x25, 0x87, 0x34,
	0xdb, 0x70, 0xc2, 0x4e, 0x69, 0x4c, 0x82, 0x9c, 0x51, 0x41, 0x11, 0x6c, 0x59, 0x1e, 0x6f, 0x14,
	0x82, 0x9f, 0x82, 0x1b, 0x92, 0xe3, 0x8a, 0x88, 0x22, 0xff, 0x4c, 0xfe, 0xa2, 0x87, 0x00, 0x5c,
	0x9e, 0x37, 0x39, 0x65, 0xc2, 0xeb, 0x8c, 0x3b, 0x4f, 0x9c, 0xd0, 0x51, 0x9d, 0xaf, 0x94, 0x09,
	0xfc, 0x4a, 0xb2, 0x79, 0xc9, 0xf6, 0xa0, 0xcf, 0x8b, 0x38, 0x26, 0x9c, 0x2b, 0xaa, 0x1d, 0x9e,
	0x4b, 0x84, 0xa0, 0x9b, 0x66, 0xbf, 0xa8, 0x77, 0x43, 0x29, 0xa8, 0x33, 0x7e, 0x0c, 0x76, 0x48,
	0x8e, 0x73, 0xba, 0x4d, 0x33, 0xe4, 0x83, 0x9d, 0x47, 0x9c, 0xff, 0xa1, 0x2c, 0x31, 0x2e, 0x65,
	0x8d, 0x5f, 0x4a, 0x1e, 0xd7, 0xbc, 0x76, 0x0e, 0x3f, 0x94, 0xc3, 0x4a, 0x44, 0x4c, 0xa0, 0x67,
	0xd0, 0x3d, 0xd0, 0x84, 0xa8, 0xb1, 0x3b, 0x93, 0x07, 0x41, 0xf5, 0xe6, 0xe0, 0xcc, 0x09, 0x16,
	0x34, 0x21, 0xa1, 0xa2, 0xe1, 0x31, 0x74, 0x65, 0x85, 0x5c, 0xe8, 0xbf, 0x7b, 0xff, 0x61, 0xfa,
	0x6d, 0xbe, 0x1e, 0x58, 0xb2, 0x58, 0x2c, 0xbf, 0x7c, 0x5a, 0x2f, 0xc3, 0x41, 0xc7, 0x5c, 0x4b,
	0x8b, 0xb7, 0xbb, 0x96, 0x0b, 0x8e, 0xb6, 0x14, 0x05, 0xc7, 0x8f, 0x64, 0xc1, 0x75, 0x81, 0x86,
	0xd0, 0x8b, 0xf6, 0xe9, 0x89, 0x18, 0x15, 0x5d, 0x60, 0x0c, 0xb7, 0xd5, 0xa2, 0xa2, 0xe4, 0x3b,
	0x65, 0xfb, 0x44, 0x6a, 0xe6, 0x91, 0xd8, 0x99, 0x45, 0xa9, 0x33, 0x7e, 0x2d, 0x39, 0xbc, 0xe2,
	0xb4, 0xbb, 0x51, 0xe5, 0x30, 0xdb, 0x45, 0x69, 0x76, 0x8d, 0x83, 0xe6, 0xb4, 0x72, 0x98, 0xfc,
	0xbb, 0x09, 0x77, 0x3f, 0xb2, 0xa2, 0x10, 0x53, 0xb9, 0xf2, 0x95, 0xce, 0x1f, 0x7a, 0x0b, 0x76,
	0x19, 0x9e, 0xfb, 0xf5, 0x4f, 0x62, 0x00, 0xbf, 0x06, 0x94, 0x71, 0xc3, 0x16, 0x7a, 0x01, 0x3d,
	0x9d, 0x8b, 0x61, 0x6d, 0x58, 0x75, 0xfd, 0x5a, 0x57, 0x67, 0x48, 0x8f, 0xe9, 0xef, 0x36, 0xbc,
	0x14, 0x83, 0xc6, 0x98, 0xea, 0x62, 0x0b, 0x4d, 0xc1, 0xb9, 0xb2, 0xe0, 0x86, 0xa3, 0x41, 0x7c,
	0xaf, 0xe1, 0x6a, 0x90, 0x4a, 0xc2, 0x6c, 0xf0, 0x82, 0x84, 0x42, 0x2e, 0x4a, 0x28, 0x04, 0x5b,
	0xe8, 0x0d, 0xb8, 0xb3, 0x1d, 0x89, 0x7f, 0x9b, 0xc8, 0x8c, 0x9a, 0x4f, 0x10, 0x05, 0xf7, 0x47,
	0xcd, 0x37, 0xc8, 0xb4, 0x59, 0x3f, 0x6f, 0xa9, 0x7f, 0xfe, 0xf9, 0xff, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x79, 0x8d, 0x1e, 0x09, 0x0a, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GruutAdminServiceClient is the client API for GruutAdminService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GruutAdminServiceClient interface {
	SetupKey(ctx context.Context, in *ReqSetupKey, opts ...grpc.CallOption) (*ResSetupKey, error)
	Login(ctx context.Context, in *ReqLogin, opts ...grpc.CallOption) (*ResLogin, error)
	Start(ctx context.Context, in *ReqStart, opts ...grpc.CallOption) (*ResStart, error)
	LoadWorld(ctx context.Context, in *ReqLoadWorld, opts ...grpc.CallOption) (*ResLoadWorld, error)
	LoadChain(ctx context.Context, in *ReqLoadChain, opts ...grpc.CallOption) (*ResLoadChain, error)
	CheckStatus(ctx context.Context, in *ReqStatus, opts ...grpc.CallOption) (*ResStatus, error)
}

type gruutAdminServiceClient struct {
	cc *grpc.ClientConn
}

func NewGruutAdminServiceClient(cc *grpc.ClientConn) GruutAdminServiceClient {
	return &gruutAdminServiceClient{cc}
}

func (c *gruutAdminServiceClient) SetupKey(ctx context.Context, in *ReqSetupKey, opts ...grpc.CallOption) (*ResSetupKey, error) {
	out := new(ResSetupKey)
	err := c.cc.Invoke(ctx, "/grpc_admin.GruutAdminService/SetupKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gruutAdminServiceClient) Login(ctx context.Context, in *ReqLogin, opts ...grpc.CallOption) (*ResLogin, error) {
	out := new(ResLogin)
	err := c.cc.Invoke(ctx, "/grpc_admin.GruutAdminService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gruutAdminServiceClient) Start(ctx context.Context, in *ReqStart, opts ...grpc.CallOption) (*ResStart, error) {
	out := new(ResStart)
	err := c.cc.Invoke(ctx, "/grpc_admin.GruutAdminService/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gruutAdminServiceClient) LoadWorld(ctx context.Context, in *ReqLoadWorld, opts ...grpc.CallOption) (*ResLoadWorld, error) {
	out := new(ResLoadWorld)
	err := c.cc.Invoke(ctx, "/grpc_admin.GruutAdminService/LoadWorld", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gruutAdminServiceClient) LoadChain(ctx context.Context, in *ReqLoadChain, opts ...grpc.CallOption) (*ResLoadChain, error) {
	out := new(ResLoadChain)
	err := c.cc.Invoke(ctx, "/grpc_admin.GruutAdminService/LoadChain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gruutAdminServiceClient) CheckStatus(ctx context.Context, in *ReqStatus, opts ...grpc.CallOption) (*ResStatus, error) {
	out := new(ResStatus)
	err := c.cc.Invoke(ctx, "/grpc_admin.GruutAdminService/CheckStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GruutAdminServiceServer is the server API for GruutAdminService service.
type GruutAdminServiceServer interface {
	SetupKey(context.Context, *ReqSetupKey) (*ResSetupKey, error)
	Login(context.Context, *ReqLogin) (*ResLogin, error)
	Start(context.Context, *ReqStart) (*ResStart, error)
	LoadWorld(context.Context, *ReqLoadWorld) (*ResLoadWorld, error)
	LoadChain(context.Context, *ReqLoadChain) (*ResLoadChain, error)
	CheckStatus(context.Context, *ReqStatus) (*ResStatus, error)
}

// UnimplementedGruutAdminServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGruutAdminServiceServer struct {
}

func (*UnimplementedGruutAdminServiceServer) SetupKey(ctx context.Context, req *ReqSetupKey) (*ResSetupKey, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetupKey not implemented")
}
func (*UnimplementedGruutAdminServiceServer) Login(ctx context.Context, req *ReqLogin) (*ResLogin, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (*UnimplementedGruutAdminServiceServer) Start(ctx context.Context, req *ReqStart) (*ResStart, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (*UnimplementedGruutAdminServiceServer) LoadWorld(ctx context.Context, req *ReqLoadWorld) (*ResLoadWorld, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadWorld not implemented")
}
func (*UnimplementedGruutAdminServiceServer) LoadChain(ctx context.Context, req *ReqLoadChain) (*ResLoadChain, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadChain not implemented")
}
func (*UnimplementedGruutAdminServiceServer) CheckStatus(ctx context.Context, req *ReqStatus) (*ResStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckStatus not implemented")
}

func RegisterGruutAdminServiceServer(s *grpc.Server, srv GruutAdminServiceServer) {
	s.RegisterService(&_GruutAdminService_serviceDesc, srv)
}

func _GruutAdminService_SetupKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqSetupKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GruutAdminServiceServer).SetupKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_admin.GruutAdminService/SetupKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GruutAdminServiceServer).SetupKey(ctx, req.(*ReqSetupKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _GruutAdminService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqLogin)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GruutAdminServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_admin.GruutAdminService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GruutAdminServiceServer).Login(ctx, req.(*ReqLogin))
	}
	return interceptor(ctx, in, info, handler)
}

func _GruutAdminService_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqStart)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GruutAdminServiceServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_admin.GruutAdminService/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GruutAdminServiceServer).Start(ctx, req.(*ReqStart))
	}
	return interceptor(ctx, in, info, handler)
}

func _GruutAdminService_LoadWorld_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqLoadWorld)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GruutAdminServiceServer).LoadWorld(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_admin.GruutAdminService/LoadWorld",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GruutAdminServiceServer).LoadWorld(ctx, req.(*ReqLoadWorld))
	}
	return interceptor(ctx, in, info, handler)
}

func _GruutAdminService_LoadChain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqLoadChain)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GruutAdminServiceServer).LoadChain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_admin.GruutAdminService/LoadChain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GruutAdminServiceServer).LoadChain(ctx, req.(*ReqLoadChain))
	}
	return interceptor(ctx, in, info, handler)
}

func _GruutAdminService_CheckStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqStatus)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GruutAdminServiceServer).CheckStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_admin.GruutAdminService/CheckStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GruutAdminServiceServer).CheckStatus(ctx, req.(*ReqStatus))
	}
	return interceptor(ctx, in, info, handler)
}

var _GruutAdminService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc_admin.GruutAdminService",
	HandlerType: (*GruutAdminServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetupKey",
			Handler:    _GruutAdminService_SetupKey_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _GruutAdminService_Login_Handler,
		},
		{
			MethodName: "Start",
			Handler:    _GruutAdminService_Start_Handler,
		},
		{
			MethodName: "LoadWorld",
			Handler:    _GruutAdminService_LoadWorld_Handler,
		},
		{
			MethodName: "LoadChain",
			Handler:    _GruutAdminService_LoadChain_Handler,
		},
		{
			MethodName: "CheckStatus",
			Handler:    _GruutAdminService_CheckStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admin_service.proto",
}
