// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/loomnetwork/go-loom/builtin/types/address_mapper/address_mapper.proto

package address_mapper

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

// Bidirectional mapping between addresses (contracts or accounts) on two different blockchains.
// One of the address fields must contain a local address (same chain ID as the contract), while the
// other must contain a foreign address.
type AddressMapperMapping struct {
	// Address on a blockchain
	From *types.Address `protobuf:"bytes,1,opt,name=from" json:"from,omitempty"`
	// Corresponding address on another blockchain
	To                   *types.Address `protobuf:"bytes,2,opt,name=to" json:"to,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *AddressMapperMapping) Reset()         { *m = AddressMapperMapping{} }
func (m *AddressMapperMapping) String() string { return proto.CompactTextString(m) }
func (*AddressMapperMapping) ProtoMessage()    {}
func (*AddressMapperMapping) Descriptor() ([]byte, []int) {
	return fileDescriptor_address_mapper_cafe86bacccc8350, []int{0}
}
func (m *AddressMapperMapping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddressMapperMapping.Unmarshal(m, b)
}
func (m *AddressMapperMapping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddressMapperMapping.Marshal(b, m, deterministic)
}
func (dst *AddressMapperMapping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddressMapperMapping.Merge(dst, src)
}
func (m *AddressMapperMapping) XXX_Size() int {
	return xxx_messageInfo_AddressMapperMapping.Size(m)
}
func (m *AddressMapperMapping) XXX_DiscardUnknown() {
	xxx_messageInfo_AddressMapperMapping.DiscardUnknown(m)
}

var xxx_messageInfo_AddressMapperMapping proto.InternalMessageInfo

func (m *AddressMapperMapping) GetFrom() *types.Address {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *AddressMapperMapping) GetTo() *types.Address {
	if m != nil {
		return m.To
	}
	return nil
}

type AddressMapperInitRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddressMapperInitRequest) Reset()         { *m = AddressMapperInitRequest{} }
func (m *AddressMapperInitRequest) String() string { return proto.CompactTextString(m) }
func (*AddressMapperInitRequest) ProtoMessage()    {}
func (*AddressMapperInitRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_address_mapper_cafe86bacccc8350, []int{1}
}
func (m *AddressMapperInitRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddressMapperInitRequest.Unmarshal(m, b)
}
func (m *AddressMapperInitRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddressMapperInitRequest.Marshal(b, m, deterministic)
}
func (dst *AddressMapperInitRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddressMapperInitRequest.Merge(dst, src)
}
func (m *AddressMapperInitRequest) XXX_Size() int {
	return xxx_messageInfo_AddressMapperInitRequest.Size(m)
}
func (m *AddressMapperInitRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddressMapperInitRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddressMapperInitRequest proto.InternalMessageInfo

type AddressMapperAddIdentityMappingRequest struct {
	// Address on a blockchain
	From *types.Address `protobuf:"bytes,1,opt,name=from" json:"from,omitempty"`
	// Corresponding address on another blockchain
	To *types.Address `protobuf:"bytes,2,opt,name=to" json:"to,omitempty"`
	// Signature of a hash of the two addresses signed with the key that corresponds to the foreign address
	Signature            []byte   `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddressMapperAddIdentityMappingRequest) Reset() {
	*m = AddressMapperAddIdentityMappingRequest{}
}
func (m *AddressMapperAddIdentityMappingRequest) String() string { return proto.CompactTextString(m) }
func (*AddressMapperAddIdentityMappingRequest) ProtoMessage()    {}
func (*AddressMapperAddIdentityMappingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_address_mapper_cafe86bacccc8350, []int{2}
}
func (m *AddressMapperAddIdentityMappingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddressMapperAddIdentityMappingRequest.Unmarshal(m, b)
}
func (m *AddressMapperAddIdentityMappingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddressMapperAddIdentityMappingRequest.Marshal(b, m, deterministic)
}
func (dst *AddressMapperAddIdentityMappingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddressMapperAddIdentityMappingRequest.Merge(dst, src)
}
func (m *AddressMapperAddIdentityMappingRequest) XXX_Size() int {
	return xxx_messageInfo_AddressMapperAddIdentityMappingRequest.Size(m)
}
func (m *AddressMapperAddIdentityMappingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddressMapperAddIdentityMappingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddressMapperAddIdentityMappingRequest proto.InternalMessageInfo

func (m *AddressMapperAddIdentityMappingRequest) GetFrom() *types.Address {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *AddressMapperAddIdentityMappingRequest) GetTo() *types.Address {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *AddressMapperAddIdentityMappingRequest) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type AddressMapperRemoveMappingRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddressMapperRemoveMappingRequest) Reset()         { *m = AddressMapperRemoveMappingRequest{} }
func (m *AddressMapperRemoveMappingRequest) String() string { return proto.CompactTextString(m) }
func (*AddressMapperRemoveMappingRequest) ProtoMessage()    {}
func (*AddressMapperRemoveMappingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_address_mapper_cafe86bacccc8350, []int{3}
}
func (m *AddressMapperRemoveMappingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddressMapperRemoveMappingRequest.Unmarshal(m, b)
}
func (m *AddressMapperRemoveMappingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddressMapperRemoveMappingRequest.Marshal(b, m, deterministic)
}
func (dst *AddressMapperRemoveMappingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddressMapperRemoveMappingRequest.Merge(dst, src)
}
func (m *AddressMapperRemoveMappingRequest) XXX_Size() int {
	return xxx_messageInfo_AddressMapperRemoveMappingRequest.Size(m)
}
func (m *AddressMapperRemoveMappingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddressMapperRemoveMappingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddressMapperRemoveMappingRequest proto.InternalMessageInfo

type AddressMapperGetMappingRequest struct {
	From                 *types.Address `protobuf:"bytes,1,opt,name=from" json:"from,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *AddressMapperGetMappingRequest) Reset()         { *m = AddressMapperGetMappingRequest{} }
func (m *AddressMapperGetMappingRequest) String() string { return proto.CompactTextString(m) }
func (*AddressMapperGetMappingRequest) ProtoMessage()    {}
func (*AddressMapperGetMappingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_address_mapper_cafe86bacccc8350, []int{4}
}
func (m *AddressMapperGetMappingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddressMapperGetMappingRequest.Unmarshal(m, b)
}
func (m *AddressMapperGetMappingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddressMapperGetMappingRequest.Marshal(b, m, deterministic)
}
func (dst *AddressMapperGetMappingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddressMapperGetMappingRequest.Merge(dst, src)
}
func (m *AddressMapperGetMappingRequest) XXX_Size() int {
	return xxx_messageInfo_AddressMapperGetMappingRequest.Size(m)
}
func (m *AddressMapperGetMappingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddressMapperGetMappingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddressMapperGetMappingRequest proto.InternalMessageInfo

func (m *AddressMapperGetMappingRequest) GetFrom() *types.Address {
	if m != nil {
		return m.From
	}
	return nil
}

type AddressMapperGetMappingResponse struct {
	From                 *types.Address `protobuf:"bytes,1,opt,name=from" json:"from,omitempty"`
	To                   *types.Address `protobuf:"bytes,2,opt,name=to" json:"to,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *AddressMapperGetMappingResponse) Reset()         { *m = AddressMapperGetMappingResponse{} }
func (m *AddressMapperGetMappingResponse) String() string { return proto.CompactTextString(m) }
func (*AddressMapperGetMappingResponse) ProtoMessage()    {}
func (*AddressMapperGetMappingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_address_mapper_cafe86bacccc8350, []int{5}
}
func (m *AddressMapperGetMappingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddressMapperGetMappingResponse.Unmarshal(m, b)
}
func (m *AddressMapperGetMappingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddressMapperGetMappingResponse.Marshal(b, m, deterministic)
}
func (dst *AddressMapperGetMappingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddressMapperGetMappingResponse.Merge(dst, src)
}
func (m *AddressMapperGetMappingResponse) XXX_Size() int {
	return xxx_messageInfo_AddressMapperGetMappingResponse.Size(m)
}
func (m *AddressMapperGetMappingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddressMapperGetMappingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddressMapperGetMappingResponse proto.InternalMessageInfo

func (m *AddressMapperGetMappingResponse) GetFrom() *types.Address {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *AddressMapperGetMappingResponse) GetTo() *types.Address {
	if m != nil {
		return m.To
	}
	return nil
}

type AddressMapperListMappingRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddressMapperListMappingRequest) Reset()         { *m = AddressMapperListMappingRequest{} }
func (m *AddressMapperListMappingRequest) String() string { return proto.CompactTextString(m) }
func (*AddressMapperListMappingRequest) ProtoMessage()    {}
func (*AddressMapperListMappingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_address_mapper_cafe86bacccc8350, []int{6}
}
func (m *AddressMapperListMappingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddressMapperListMappingRequest.Unmarshal(m, b)
}
func (m *AddressMapperListMappingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddressMapperListMappingRequest.Marshal(b, m, deterministic)
}
func (dst *AddressMapperListMappingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddressMapperListMappingRequest.Merge(dst, src)
}
func (m *AddressMapperListMappingRequest) XXX_Size() int {
	return xxx_messageInfo_AddressMapperListMappingRequest.Size(m)
}
func (m *AddressMapperListMappingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddressMapperListMappingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddressMapperListMappingRequest proto.InternalMessageInfo

type AddressMapperListMappingResponse struct {
	AddressMapperGetMappingResponse []*AddressMapperGetMappingResponse `protobuf:"bytes,1,rep,name=AddressMapperGetMappingResponse" json:"AddressMapperGetMappingResponse,omitempty"`
	XXX_NoUnkeyedLiteral            struct{}                           `json:"-"`
	XXX_unrecognized                []byte                             `json:"-"`
	XXX_sizecache                   int32                              `json:"-"`
}

func (m *AddressMapperListMappingResponse) Reset()         { *m = AddressMapperListMappingResponse{} }
func (m *AddressMapperListMappingResponse) String() string { return proto.CompactTextString(m) }
func (*AddressMapperListMappingResponse) ProtoMessage()    {}
func (*AddressMapperListMappingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_address_mapper_cafe86bacccc8350, []int{7}
}
func (m *AddressMapperListMappingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddressMapperListMappingResponse.Unmarshal(m, b)
}
func (m *AddressMapperListMappingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddressMapperListMappingResponse.Marshal(b, m, deterministic)
}
func (dst *AddressMapperListMappingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddressMapperListMappingResponse.Merge(dst, src)
}
func (m *AddressMapperListMappingResponse) XXX_Size() int {
	return xxx_messageInfo_AddressMapperListMappingResponse.Size(m)
}
func (m *AddressMapperListMappingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddressMapperListMappingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddressMapperListMappingResponse proto.InternalMessageInfo

func (m *AddressMapperListMappingResponse) GetAddressMapperGetMappingResponse() []*AddressMapperGetMappingResponse {
	if m != nil {
		return m.AddressMapperGetMappingResponse
	}
	return nil
}

type AddressMapperHasMappingRequest struct {
	From                 *types.Address `protobuf:"bytes,1,opt,name=from" json:"from,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *AddressMapperHasMappingRequest) Reset()         { *m = AddressMapperHasMappingRequest{} }
func (m *AddressMapperHasMappingRequest) String() string { return proto.CompactTextString(m) }
func (*AddressMapperHasMappingRequest) ProtoMessage()    {}
func (*AddressMapperHasMappingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_address_mapper_cafe86bacccc8350, []int{8}
}
func (m *AddressMapperHasMappingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddressMapperHasMappingRequest.Unmarshal(m, b)
}
func (m *AddressMapperHasMappingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddressMapperHasMappingRequest.Marshal(b, m, deterministic)
}
func (dst *AddressMapperHasMappingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddressMapperHasMappingRequest.Merge(dst, src)
}
func (m *AddressMapperHasMappingRequest) XXX_Size() int {
	return xxx_messageInfo_AddressMapperHasMappingRequest.Size(m)
}
func (m *AddressMapperHasMappingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddressMapperHasMappingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddressMapperHasMappingRequest proto.InternalMessageInfo

func (m *AddressMapperHasMappingRequest) GetFrom() *types.Address {
	if m != nil {
		return m.From
	}
	return nil
}

type AddressMapperHasMappingResponse struct {
	HasMapping           bool     `protobuf:"varint,1,opt,name=has_mapping,json=hasMapping,proto3" json:"has_mapping,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddressMapperHasMappingResponse) Reset()         { *m = AddressMapperHasMappingResponse{} }
func (m *AddressMapperHasMappingResponse) String() string { return proto.CompactTextString(m) }
func (*AddressMapperHasMappingResponse) ProtoMessage()    {}
func (*AddressMapperHasMappingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_address_mapper_cafe86bacccc8350, []int{9}
}
func (m *AddressMapperHasMappingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddressMapperHasMappingResponse.Unmarshal(m, b)
}
func (m *AddressMapperHasMappingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddressMapperHasMappingResponse.Marshal(b, m, deterministic)
}
func (dst *AddressMapperHasMappingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddressMapperHasMappingResponse.Merge(dst, src)
}
func (m *AddressMapperHasMappingResponse) XXX_Size() int {
	return xxx_messageInfo_AddressMapperHasMappingResponse.Size(m)
}
func (m *AddressMapperHasMappingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddressMapperHasMappingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddressMapperHasMappingResponse proto.InternalMessageInfo

func (m *AddressMapperHasMappingResponse) GetHasMapping() bool {
	if m != nil {
		return m.HasMapping
	}
	return false
}

func init() {
	proto.RegisterType((*AddressMapperMapping)(nil), "AddressMapperMapping")
	proto.RegisterType((*AddressMapperInitRequest)(nil), "AddressMapperInitRequest")
	proto.RegisterType((*AddressMapperAddIdentityMappingRequest)(nil), "AddressMapperAddIdentityMappingRequest")
	proto.RegisterType((*AddressMapperRemoveMappingRequest)(nil), "AddressMapperRemoveMappingRequest")
	proto.RegisterType((*AddressMapperGetMappingRequest)(nil), "AddressMapperGetMappingRequest")
	proto.RegisterType((*AddressMapperGetMappingResponse)(nil), "AddressMapperGetMappingResponse")
	proto.RegisterType((*AddressMapperListMappingRequest)(nil), "AddressMapperListMappingRequest")
	proto.RegisterType((*AddressMapperListMappingResponse)(nil), "AddressMapperListMappingResponse")
	proto.RegisterType((*AddressMapperHasMappingRequest)(nil), "AddressMapperHasMappingRequest")
	proto.RegisterType((*AddressMapperHasMappingResponse)(nil), "AddressMapperHasMappingResponse")
}

func init() {
	proto.RegisterFile("github.com/loomnetwork/go-loom/builtin/types/address_mapper/address_mapper.proto", fileDescriptor_address_mapper_cafe86bacccc8350)
}

var fileDescriptor_address_mapper_cafe86bacccc8350 = []byte{
	// 322 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x41, 0x4b, 0xfb, 0x40,
	0x10, 0xc5, 0x49, 0xfb, 0xe7, 0x4f, 0x9d, 0x7a, 0x2a, 0x1e, 0x42, 0x29, 0x36, 0x5d, 0x41, 0x7a,
	0x31, 0x91, 0x7a, 0x17, 0xea, 0x45, 0x0b, 0x2a, 0xb2, 0x37, 0x4f, 0x92, 0x9a, 0x31, 0x5d, 0x6d,
	0x76, 0xe2, 0xee, 0x44, 0xe9, 0xc5, 0x8f, 0xe0, 0x67, 0x96, 0x26, 0x41, 0x4d, 0xa4, 0x84, 0xf6,
	0xb2, 0xb0, 0x33, 0x6f, 0x7e, 0xef, 0x65, 0xc8, 0xc2, 0x5d, 0xac, 0x78, 0x91, 0xcd, 0xfd, 0x47,
	0x4a, 0x82, 0x25, 0x51, 0xa2, 0x91, 0xdf, 0xc9, 0xbc, 0x04, 0x31, 0x9d, 0xac, 0xaf, 0xc1, 0x3c,
	0x53, 0x4b, 0x56, 0x3a, 0xe0, 0x55, 0x8a, 0x36, 0x08, 0xa3, 0xc8, 0xa0, 0xb5, 0x0f, 0x49, 0x98,
	0xa6, 0x68, 0x6a, 0x57, 0x3f, 0x35, 0xc4, 0xd4, 0x3f, 0x6d, 0x20, 0x16, 0xa4, 0xfc, 0x2c, 0x26,
	0xc4, 0x2d, 0x1c, 0x4c, 0x0b, 0xd2, 0x4d, 0x0e, 0x5a, 0x9f, 0x4a, 0xc7, 0xbd, 0x01, 0xfc, 0x7b,
	0x32, 0x94, 0xb8, 0x8e, 0xe7, 0x8c, 0xbb, 0x93, 0x8e, 0x5f, 0x8a, 0x64, 0x5e, 0xed, 0xb9, 0xd0,
	0x62, 0x72, 0x5b, 0xb5, 0x5e, 0x8b, 0x49, 0xf4, 0xc1, 0xad, 0xf0, 0x66, 0x5a, 0xb1, 0xc4, 0xd7,
	0x0c, 0x2d, 0x8b, 0x0f, 0x38, 0xae, 0xf4, 0xa6, 0x51, 0x34, 0x8b, 0x50, 0xb3, 0xe2, 0x55, 0x69,
	0x5b, 0x2a, 0x77, 0x75, 0xef, 0x0d, 0x60, 0xcf, 0xaa, 0x58, 0x87, 0x9c, 0x19, 0x74, 0xdb, 0x9e,
	0x33, 0xde, 0x97, 0x3f, 0x05, 0x71, 0x04, 0xa3, 0x8a, 0xbf, 0xc4, 0x84, 0xde, 0xb0, 0x6a, 0x2d,
	0xce, 0xe1, 0xb0, 0x22, 0xba, 0x44, 0xde, 0x26, 0x9c, 0xb8, 0x87, 0xe1, 0xc6, 0x79, 0x9b, 0x92,
	0xb6, 0xb8, 0xf3, 0x6e, 0x47, 0x35, 0xf4, 0xb5, 0xb2, 0xb5, 0x6c, 0xe2, 0xd3, 0x01, 0x6f, 0xb3,
	0xa6, 0xf4, 0x7f, 0x6e, 0x8c, 0xe8, 0x3a, 0x5e, 0x7b, 0xdc, 0x9d, 0x78, 0x7e, 0x83, 0x4e, 0x36,
	0x81, 0xfe, 0xac, 0xf3, 0x2a, 0xb4, 0x5b, 0xad, 0xf3, 0xa2, 0x96, 0xf5, 0xf7, 0x7c, 0xf9, 0x39,
	0x43, 0xe8, 0x2e, 0xc2, 0xe2, 0x21, 0x28, 0x1d, 0xe7, 0x9c, 0x8e, 0x84, 0xc5, 0xb7, 0x70, 0xfe,
	0x3f, 0xff, 0xd5, 0xcf, 0xbe, 0x02, 0x00, 0x00, 0xff, 0xff, 0x08, 0x74, 0x59, 0x59, 0x70, 0x03,
	0x00, 0x00,
}
