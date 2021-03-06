// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/loomnetwork/go-loom/vm/vm.proto

package vm

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import types "github.com/loomnetwork/go-loom/types"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type VMType int32

const (
	VMType_PLUGIN VMType = 0
	VMType_EVM    VMType = 1
)

var VMType_name = map[int32]string{
	0: "PLUGIN",
	1: "EVM",
}
var VMType_value = map[string]int32{
	"PLUGIN": 0,
	"EVM":    1,
}

func (x VMType) String() string {
	return proto.EnumName(VMType_name, int32(x))
}
func (VMType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_vm_bcdc7a57d7080a9e, []int{0}
}

type MessageTx struct {
	To                   *types.Address `protobuf:"bytes,1,opt,name=to" json:"to,omitempty"`
	From                 *types.Address `protobuf:"bytes,2,opt,name=from" json:"from,omitempty"`
	Data                 []byte         `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *MessageTx) Reset()         { *m = MessageTx{} }
func (m *MessageTx) String() string { return proto.CompactTextString(m) }
func (*MessageTx) ProtoMessage()    {}
func (*MessageTx) Descriptor() ([]byte, []int) {
	return fileDescriptor_vm_bcdc7a57d7080a9e, []int{0}
}
func (m *MessageTx) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageTx.Unmarshal(m, b)
}
func (m *MessageTx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageTx.Marshal(b, m, deterministic)
}
func (dst *MessageTx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageTx.Merge(dst, src)
}
func (m *MessageTx) XXX_Size() int {
	return xxx_messageInfo_MessageTx.Size(m)
}
func (m *MessageTx) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageTx.DiscardUnknown(m)
}

var xxx_messageInfo_MessageTx proto.InternalMessageInfo

func (m *MessageTx) GetTo() *types.Address {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *MessageTx) GetFrom() *types.Address {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *MessageTx) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type DeployTx struct {
	VmType               VMType         `protobuf:"varint,1,opt,name=vm_type,json=vmType,proto3,enum=VMType" json:"vm_type,omitempty"`
	Code                 []byte         `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Name                 string         `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Value                *types.BigUInt `protobuf:"bytes,4,opt,name=value" json:"value,omitempty"`
	ContractVersion      string         `protobuf:"bytes,5,opt,name=contract_version,json=contractVersion,proto3" json:"contract_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *DeployTx) Reset()         { *m = DeployTx{} }
func (m *DeployTx) String() string { return proto.CompactTextString(m) }
func (*DeployTx) ProtoMessage()    {}
func (*DeployTx) Descriptor() ([]byte, []int) {
	return fileDescriptor_vm_bcdc7a57d7080a9e, []int{1}
}
func (m *DeployTx) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeployTx.Unmarshal(m, b)
}
func (m *DeployTx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeployTx.Marshal(b, m, deterministic)
}
func (dst *DeployTx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeployTx.Merge(dst, src)
}
func (m *DeployTx) XXX_Size() int {
	return xxx_messageInfo_DeployTx.Size(m)
}
func (m *DeployTx) XXX_DiscardUnknown() {
	xxx_messageInfo_DeployTx.DiscardUnknown(m)
}

var xxx_messageInfo_DeployTx proto.InternalMessageInfo

func (m *DeployTx) GetVmType() VMType {
	if m != nil {
		return m.VmType
	}
	return VMType_PLUGIN
}

func (m *DeployTx) GetCode() []byte {
	if m != nil {
		return m.Code
	}
	return nil
}

func (m *DeployTx) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DeployTx) GetValue() *types.BigUInt {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *DeployTx) GetContractVersion() string {
	if m != nil {
		return m.ContractVersion
	}
	return ""
}

type CallTx struct {
	VmType               VMType         `protobuf:"varint,1,opt,name=vm_type,json=vmType,proto3,enum=VMType" json:"vm_type,omitempty"`
	Input                []byte         `protobuf:"bytes,2,opt,name=input,proto3" json:"input,omitempty"`
	Value                *types.BigUInt `protobuf:"bytes,3,opt,name=value" json:"value,omitempty"`
	ContractVersion      string         `protobuf:"bytes,4,opt,name=contract_version,json=contractVersion,proto3" json:"contract_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CallTx) Reset()         { *m = CallTx{} }
func (m *CallTx) String() string { return proto.CompactTextString(m) }
func (*CallTx) ProtoMessage()    {}
func (*CallTx) Descriptor() ([]byte, []int) {
	return fileDescriptor_vm_bcdc7a57d7080a9e, []int{2}
}
func (m *CallTx) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CallTx.Unmarshal(m, b)
}
func (m *CallTx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CallTx.Marshal(b, m, deterministic)
}
func (dst *CallTx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallTx.Merge(dst, src)
}
func (m *CallTx) XXX_Size() int {
	return xxx_messageInfo_CallTx.Size(m)
}
func (m *CallTx) XXX_DiscardUnknown() {
	xxx_messageInfo_CallTx.DiscardUnknown(m)
}

var xxx_messageInfo_CallTx proto.InternalMessageInfo

func (m *CallTx) GetVmType() VMType {
	if m != nil {
		return m.VmType
	}
	return VMType_PLUGIN
}

func (m *CallTx) GetInput() []byte {
	if m != nil {
		return m.Input
	}
	return nil
}

func (m *CallTx) GetValue() *types.BigUInt {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *CallTx) GetContractVersion() string {
	if m != nil {
		return m.ContractVersion
	}
	return ""
}

type DeployResponse struct {
	Contract             *types.Address `protobuf:"bytes,1,opt,name=contract" json:"contract,omitempty"`
	Output               []byte         `protobuf:"bytes,2,opt,name=output,proto3" json:"output,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *DeployResponse) Reset()         { *m = DeployResponse{} }
func (m *DeployResponse) String() string { return proto.CompactTextString(m) }
func (*DeployResponse) ProtoMessage()    {}
func (*DeployResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_vm_bcdc7a57d7080a9e, []int{3}
}
func (m *DeployResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeployResponse.Unmarshal(m, b)
}
func (m *DeployResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeployResponse.Marshal(b, m, deterministic)
}
func (dst *DeployResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeployResponse.Merge(dst, src)
}
func (m *DeployResponse) XXX_Size() int {
	return xxx_messageInfo_DeployResponse.Size(m)
}
func (m *DeployResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeployResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeployResponse proto.InternalMessageInfo

func (m *DeployResponse) GetContract() *types.Address {
	if m != nil {
		return m.Contract
	}
	return nil
}

func (m *DeployResponse) GetOutput() []byte {
	if m != nil {
		return m.Output
	}
	return nil
}

type DeployResponseData struct {
	TxHash               []byte   `protobuf:"bytes,1,opt,name=tx_hash,json=txHash,proto3" json:"tx_hash,omitempty"`
	Bytecode             []byte   `protobuf:"bytes,2,opt,name=bytecode,proto3" json:"bytecode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeployResponseData) Reset()         { *m = DeployResponseData{} }
func (m *DeployResponseData) String() string { return proto.CompactTextString(m) }
func (*DeployResponseData) ProtoMessage()    {}
func (*DeployResponseData) Descriptor() ([]byte, []int) {
	return fileDescriptor_vm_bcdc7a57d7080a9e, []int{4}
}
func (m *DeployResponseData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeployResponseData.Unmarshal(m, b)
}
func (m *DeployResponseData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeployResponseData.Marshal(b, m, deterministic)
}
func (dst *DeployResponseData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeployResponseData.Merge(dst, src)
}
func (m *DeployResponseData) XXX_Size() int {
	return xxx_messageInfo_DeployResponseData.Size(m)
}
func (m *DeployResponseData) XXX_DiscardUnknown() {
	xxx_messageInfo_DeployResponseData.DiscardUnknown(m)
}

var xxx_messageInfo_DeployResponseData proto.InternalMessageInfo

func (m *DeployResponseData) GetTxHash() []byte {
	if m != nil {
		return m.TxHash
	}
	return nil
}

func (m *DeployResponseData) GetBytecode() []byte {
	if m != nil {
		return m.Bytecode
	}
	return nil
}

type TxHashList struct {
	TxHash               [][]byte `protobuf:"bytes,1,rep,name=tx_hash,json=txHash" json:"tx_hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TxHashList) Reset()         { *m = TxHashList{} }
func (m *TxHashList) String() string { return proto.CompactTextString(m) }
func (*TxHashList) ProtoMessage()    {}
func (*TxHashList) Descriptor() ([]byte, []int) {
	return fileDescriptor_vm_bcdc7a57d7080a9e, []int{5}
}
func (m *TxHashList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TxHashList.Unmarshal(m, b)
}
func (m *TxHashList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TxHashList.Marshal(b, m, deterministic)
}
func (dst *TxHashList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxHashList.Merge(dst, src)
}
func (m *TxHashList) XXX_Size() int {
	return xxx_messageInfo_TxHashList.Size(m)
}
func (m *TxHashList) XXX_DiscardUnknown() {
	xxx_messageInfo_TxHashList.DiscardUnknown(m)
}

var xxx_messageInfo_TxHashList proto.InternalMessageInfo

func (m *TxHashList) GetTxHash() [][]byte {
	if m != nil {
		return m.TxHash
	}
	return nil
}

func init() {
	proto.RegisterType((*MessageTx)(nil), "MessageTx")
	proto.RegisterType((*DeployTx)(nil), "DeployTx")
	proto.RegisterType((*CallTx)(nil), "CallTx")
	proto.RegisterType((*DeployResponse)(nil), "DeployResponse")
	proto.RegisterType((*DeployResponseData)(nil), "DeployResponseData")
	proto.RegisterType((*TxHashList)(nil), "TxHashList")
	proto.RegisterEnum("VMType", VMType_name, VMType_value)
}

func init() {
	proto.RegisterFile("github.com/loomnetwork/go-loom/vm/vm.proto", fileDescriptor_vm_bcdc7a57d7080a9e)
}

var fileDescriptor_vm_bcdc7a57d7080a9e = []byte{
	// 400 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xe1, 0x8a, 0xd3, 0x40,
	0x10, 0x36, 0x6d, 0x9a, 0xe4, 0xc6, 0x72, 0x96, 0x45, 0x34, 0x1c, 0x2a, 0x25, 0x28, 0x9c, 0x07,
	0xa6, 0x72, 0x3e, 0x81, 0x7a, 0xa2, 0x85, 0xeb, 0x21, 0x4b, 0xae, 0xfe, 0x2c, 0xdb, 0x64, 0x4d,
	0x82, 0x49, 0x26, 0x64, 0x27, 0x31, 0x79, 0x06, 0x1f, 0xc2, 0x57, 0x95, 0x6c, 0x7a, 0x57, 0x4f,
	0x05, 0x0b, 0x21, 0x99, 0xef, 0x9b, 0x99, 0x6f, 0x3e, 0x66, 0x02, 0x67, 0x71, 0x4a, 0x49, 0xbd,
	0xf5, 0x43, 0xcc, 0x17, 0x19, 0x62, 0x5e, 0x48, 0xfa, 0x8e, 0xd5, 0xb7, 0x45, 0x8c, 0xaf, 0x7a,
	0xb8, 0x68, 0xfa, 0xc7, 0x2f, 0x2b, 0x24, 0x3c, 0x79, 0xfd, 0x9f, 0x5a, 0xea, 0x4a, 0xa9, 0x86,
	0xf7, 0xd0, 0xe1, 0x7d, 0x81, 0xa3, 0x95, 0x54, 0x4a, 0xc4, 0x32, 0x68, 0x99, 0x0b, 0x23, 0x42,
	0xd7, 0x98, 0x1b, 0xa7, 0xf7, 0xcf, 0x1d, 0xff, 0x6d, 0x14, 0x55, 0x52, 0x29, 0x3e, 0x22, 0x64,
	0x4f, 0xc0, 0xfc, 0x5a, 0x61, 0xee, 0x8e, 0xfe, 0xc8, 0x69, 0x96, 0x31, 0x30, 0x23, 0x41, 0xc2,
	0x1d, 0xcf, 0x8d, 0xd3, 0x29, 0xd7, 0xb1, 0xf7, 0xd3, 0x00, 0xe7, 0x42, 0x96, 0x19, 0x76, 0x41,
	0xcb, 0xe6, 0x60, 0x37, 0xf9, 0xa6, 0x9f, 0xab, 0xd5, 0x8f, 0xcf, 0x6d, 0x7f, 0xbd, 0x0a, 0xba,
	0x52, 0x72, 0xab, 0xc9, 0xfb, 0x6f, 0x2f, 0x11, 0x62, 0x24, 0xf5, 0x80, 0x29, 0xd7, 0x71, 0xcf,
	0x15, 0x22, 0x97, 0x5a, 0xf6, 0x88, 0xeb, 0x98, 0x3d, 0x83, 0x49, 0x23, 0xb2, 0x5a, 0xba, 0xe6,
	0xce, 0xc9, 0xbb, 0x34, 0xbe, 0x5e, 0x16, 0xc4, 0x07, 0x9a, 0xbd, 0x84, 0x59, 0x88, 0x05, 0x55,
	0x22, 0xa4, 0x4d, 0x23, 0x2b, 0x95, 0x62, 0xe1, 0x4e, 0x74, 0xff, 0x83, 0x1b, 0x7e, 0x3d, 0xd0,
	0xde, 0x0f, 0x03, 0xac, 0xf7, 0x22, 0xcb, 0x0e, 0xf2, 0xf7, 0x10, 0x26, 0x69, 0x51, 0xd6, 0xb4,
	0x33, 0x38, 0x80, 0xbd, 0x9b, 0xf1, 0xe1, 0x6e, 0xcc, 0x7f, 0xbb, 0xb9, 0x82, 0xe3, 0x61, 0x5d,
	0x5c, 0xaa, 0x12, 0x0b, 0x25, 0xd9, 0x73, 0x70, 0x6e, 0x8a, 0xfe, 0xba, 0xc9, 0x6d, 0x86, 0x3d,
	0x02, 0x0b, 0x6b, 0xda, 0x3b, 0xdb, 0x21, 0x6f, 0x09, 0xec, 0xae, 0xde, 0x85, 0x20, 0xc1, 0x1e,
	0x83, 0x4d, 0xed, 0x26, 0x11, 0x2a, 0xd1, 0x92, 0x53, 0x6e, 0x51, 0xfb, 0x49, 0xa8, 0x84, 0x9d,
	0x80, 0xb3, 0xed, 0x48, 0xfe, 0x76, 0x83, 0x5b, 0xec, 0xbd, 0x00, 0x08, 0x74, 0xd5, 0x65, 0xaa,
	0xe8, 0xae, 0xc4, 0x78, 0x2f, 0x71, 0xf6, 0x14, 0xac, 0x61, 0x69, 0x0c, 0xc0, 0xfa, 0x7c, 0x79,
	0xfd, 0x71, 0x79, 0x35, 0xbb, 0xc7, 0x6c, 0x18, 0x7f, 0x58, 0xaf, 0x66, 0xc6, 0xd6, 0xd2, 0x3f,
	0xdc, 0x9b, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc3, 0x7e, 0x33, 0x9b, 0xd0, 0x02, 0x00, 0x00,
}
