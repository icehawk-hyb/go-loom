// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/loomnetwork/go-loom/builtin/types/karma/karma.proto

package karma

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

type KarmaSourceTarget int32

const (
	KarmaSourceTarget_CALL   KarmaSourceTarget = 0
	KarmaSourceTarget_DEPLOY KarmaSourceTarget = 1
)

var KarmaSourceTarget_name = map[int32]string{
	0: "CALL",
	1: "DEPLOY",
}
var KarmaSourceTarget_value = map[string]int32{
	"CALL":   0,
	"DEPLOY": 1,
}

func (x KarmaSourceTarget) String() string {
	return proto.EnumName(KarmaSourceTarget_name, int32(x))
}
func (KarmaSourceTarget) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_karma_2da51ebadefe591e, []int{0}
}

type KarmaInitRequest struct {
	Oracle               *types.Address        `protobuf:"bytes,1,opt,name=Oracle" json:"Oracle,omitempty"`
	Sources              []*KarmaSourceReward  `protobuf:"bytes,2,rep,name=sources" json:"sources,omitempty"`
	Users                []*KarmaAddressSource `protobuf:"bytes,3,rep,name=users" json:"users,omitempty"`
	Upkeep               *KarmaUpkeepParams    `protobuf:"bytes,4,opt,name=upkeep" json:"upkeep,omitempty"`
	Config               *KarmaConfig          `protobuf:"bytes,5,opt,name=config" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *KarmaInitRequest) Reset()         { *m = KarmaInitRequest{} }
func (m *KarmaInitRequest) String() string { return proto.CompactTextString(m) }
func (*KarmaInitRequest) ProtoMessage()    {}
func (*KarmaInitRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_karma_2da51ebadefe591e, []int{0}
}
func (m *KarmaInitRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KarmaInitRequest.Unmarshal(m, b)
}
func (m *KarmaInitRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KarmaInitRequest.Marshal(b, m, deterministic)
}
func (dst *KarmaInitRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KarmaInitRequest.Merge(dst, src)
}
func (m *KarmaInitRequest) XXX_Size() int {
	return xxx_messageInfo_KarmaInitRequest.Size(m)
}
func (m *KarmaInitRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_KarmaInitRequest.DiscardUnknown(m)
}

var xxx_messageInfo_KarmaInitRequest proto.InternalMessageInfo

func (m *KarmaInitRequest) GetOracle() *types.Address {
	if m != nil {
		return m.Oracle
	}
	return nil
}

func (m *KarmaInitRequest) GetSources() []*KarmaSourceReward {
	if m != nil {
		return m.Sources
	}
	return nil
}

func (m *KarmaInitRequest) GetUsers() []*KarmaAddressSource {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *KarmaInitRequest) GetUpkeep() *KarmaUpkeepParams {
	if m != nil {
		return m.Upkeep
	}
	return nil
}

func (m *KarmaInitRequest) GetConfig() *KarmaConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

type KarmaConfig struct {
	MinKarmaToDeploy     int64    `protobuf:"varint,1,opt,name=min_karma_to_deploy,json=minKarmaToDeploy,proto3" json:"min_karma_to_deploy,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KarmaConfig) Reset()         { *m = KarmaConfig{} }
func (m *KarmaConfig) String() string { return proto.CompactTextString(m) }
func (*KarmaConfig) ProtoMessage()    {}
func (*KarmaConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_karma_2da51ebadefe591e, []int{1}
}
func (m *KarmaConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KarmaConfig.Unmarshal(m, b)
}
func (m *KarmaConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KarmaConfig.Marshal(b, m, deterministic)
}
func (dst *KarmaConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KarmaConfig.Merge(dst, src)
}
func (m *KarmaConfig) XXX_Size() int {
	return xxx_messageInfo_KarmaConfig.Size(m)
}
func (m *KarmaConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_KarmaConfig.DiscardUnknown(m)
}

var xxx_messageInfo_KarmaConfig proto.InternalMessageInfo

func (m *KarmaConfig) GetMinKarmaToDeploy() int64 {
	if m != nil {
		return m.MinKarmaToDeploy
	}
	return 0
}

type KarmaSources struct {
	Sources              []*KarmaSourceReward `protobuf:"bytes,1,rep,name=sources" json:"sources,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *KarmaSources) Reset()         { *m = KarmaSources{} }
func (m *KarmaSources) String() string { return proto.CompactTextString(m) }
func (*KarmaSources) ProtoMessage()    {}
func (*KarmaSources) Descriptor() ([]byte, []int) {
	return fileDescriptor_karma_2da51ebadefe591e, []int{2}
}
func (m *KarmaSources) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KarmaSources.Unmarshal(m, b)
}
func (m *KarmaSources) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KarmaSources.Marshal(b, m, deterministic)
}
func (dst *KarmaSources) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KarmaSources.Merge(dst, src)
}
func (m *KarmaSources) XXX_Size() int {
	return xxx_messageInfo_KarmaSources.Size(m)
}
func (m *KarmaSources) XXX_DiscardUnknown() {
	xxx_messageInfo_KarmaSources.DiscardUnknown(m)
}

var xxx_messageInfo_KarmaSources proto.InternalMessageInfo

func (m *KarmaSources) GetSources() []*KarmaSourceReward {
	if m != nil {
		return m.Sources
	}
	return nil
}

type KarmaNewOracle struct {
	NewOracle            *types.Address `protobuf:"bytes,1,opt,name=new_oracle,json=newOracle" json:"new_oracle,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *KarmaNewOracle) Reset()         { *m = KarmaNewOracle{} }
func (m *KarmaNewOracle) String() string { return proto.CompactTextString(m) }
func (*KarmaNewOracle) ProtoMessage()    {}
func (*KarmaNewOracle) Descriptor() ([]byte, []int) {
	return fileDescriptor_karma_2da51ebadefe591e, []int{3}
}
func (m *KarmaNewOracle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KarmaNewOracle.Unmarshal(m, b)
}
func (m *KarmaNewOracle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KarmaNewOracle.Marshal(b, m, deterministic)
}
func (dst *KarmaNewOracle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KarmaNewOracle.Merge(dst, src)
}
func (m *KarmaNewOracle) XXX_Size() int {
	return xxx_messageInfo_KarmaNewOracle.Size(m)
}
func (m *KarmaNewOracle) XXX_DiscardUnknown() {
	xxx_messageInfo_KarmaNewOracle.DiscardUnknown(m)
}

var xxx_messageInfo_KarmaNewOracle proto.InternalMessageInfo

func (m *KarmaNewOracle) GetNewOracle() *types.Address {
	if m != nil {
		return m.NewOracle
	}
	return nil
}

type KarmaUserTarget struct {
	User                 *types.Address    `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
	Target               KarmaSourceTarget `protobuf:"varint,2,opt,name=target,proto3,enum=karma.KarmaSourceTarget" json:"target,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *KarmaUserTarget) Reset()         { *m = KarmaUserTarget{} }
func (m *KarmaUserTarget) String() string { return proto.CompactTextString(m) }
func (*KarmaUserTarget) ProtoMessage()    {}
func (*KarmaUserTarget) Descriptor() ([]byte, []int) {
	return fileDescriptor_karma_2da51ebadefe591e, []int{4}
}
func (m *KarmaUserTarget) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KarmaUserTarget.Unmarshal(m, b)
}
func (m *KarmaUserTarget) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KarmaUserTarget.Marshal(b, m, deterministic)
}
func (dst *KarmaUserTarget) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KarmaUserTarget.Merge(dst, src)
}
func (m *KarmaUserTarget) XXX_Size() int {
	return xxx_messageInfo_KarmaUserTarget.Size(m)
}
func (m *KarmaUserTarget) XXX_DiscardUnknown() {
	xxx_messageInfo_KarmaUserTarget.DiscardUnknown(m)
}

var xxx_messageInfo_KarmaUserTarget proto.InternalMessageInfo

func (m *KarmaUserTarget) GetUser() *types.Address {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *KarmaUserTarget) GetTarget() KarmaSourceTarget {
	if m != nil {
		return m.Target
	}
	return KarmaSourceTarget_CALL
}

type KarmaUserAmount struct {
	User                 *types.Address `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
	Amount               *types.BigUInt `protobuf:"bytes,2,opt,name=amount" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *KarmaUserAmount) Reset()         { *m = KarmaUserAmount{} }
func (m *KarmaUserAmount) String() string { return proto.CompactTextString(m) }
func (*KarmaUserAmount) ProtoMessage()    {}
func (*KarmaUserAmount) Descriptor() ([]byte, []int) {
	return fileDescriptor_karma_2da51ebadefe591e, []int{5}
}
func (m *KarmaUserAmount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KarmaUserAmount.Unmarshal(m, b)
}
func (m *KarmaUserAmount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KarmaUserAmount.Marshal(b, m, deterministic)
}
func (dst *KarmaUserAmount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KarmaUserAmount.Merge(dst, src)
}
func (m *KarmaUserAmount) XXX_Size() int {
	return xxx_messageInfo_KarmaUserAmount.Size(m)
}
func (m *KarmaUserAmount) XXX_DiscardUnknown() {
	xxx_messageInfo_KarmaUserAmount.DiscardUnknown(m)
}

var xxx_messageInfo_KarmaUserAmount proto.InternalMessageInfo

func (m *KarmaUserAmount) GetUser() *types.Address {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *KarmaUserAmount) GetAmount() *types.BigUInt {
	if m != nil {
		return m.Amount
	}
	return nil
}

type KarmaSourceReward struct {
	Name                 string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Reward               int64             `protobuf:"varint,2,opt,name=reward,proto3" json:"reward,omitempty"`
	Target               KarmaSourceTarget `protobuf:"varint,3,opt,name=target,proto3,enum=karma.KarmaSourceTarget" json:"target,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *KarmaSourceReward) Reset()         { *m = KarmaSourceReward{} }
func (m *KarmaSourceReward) String() string { return proto.CompactTextString(m) }
func (*KarmaSourceReward) ProtoMessage()    {}
func (*KarmaSourceReward) Descriptor() ([]byte, []int) {
	return fileDescriptor_karma_2da51ebadefe591e, []int{6}
}
func (m *KarmaSourceReward) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KarmaSourceReward.Unmarshal(m, b)
}
func (m *KarmaSourceReward) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KarmaSourceReward.Marshal(b, m, deterministic)
}
func (dst *KarmaSourceReward) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KarmaSourceReward.Merge(dst, src)
}
func (m *KarmaSourceReward) XXX_Size() int {
	return xxx_messageInfo_KarmaSourceReward.Size(m)
}
func (m *KarmaSourceReward) XXX_DiscardUnknown() {
	xxx_messageInfo_KarmaSourceReward.DiscardUnknown(m)
}

var xxx_messageInfo_KarmaSourceReward proto.InternalMessageInfo

func (m *KarmaSourceReward) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *KarmaSourceReward) GetReward() int64 {
	if m != nil {
		return m.Reward
	}
	return 0
}

func (m *KarmaSourceReward) GetTarget() KarmaSourceTarget {
	if m != nil {
		return m.Target
	}
	return KarmaSourceTarget_CALL
}

type KarmaSource struct {
	Name                 string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Count                *types.BigUInt `protobuf:"bytes,2,opt,name=count" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *KarmaSource) Reset()         { *m = KarmaSource{} }
func (m *KarmaSource) String() string { return proto.CompactTextString(m) }
func (*KarmaSource) ProtoMessage()    {}
func (*KarmaSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_karma_2da51ebadefe591e, []int{7}
}
func (m *KarmaSource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KarmaSource.Unmarshal(m, b)
}
func (m *KarmaSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KarmaSource.Marshal(b, m, deterministic)
}
func (dst *KarmaSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KarmaSource.Merge(dst, src)
}
func (m *KarmaSource) XXX_Size() int {
	return xxx_messageInfo_KarmaSource.Size(m)
}
func (m *KarmaSource) XXX_DiscardUnknown() {
	xxx_messageInfo_KarmaSource.DiscardUnknown(m)
}

var xxx_messageInfo_KarmaSource proto.InternalMessageInfo

func (m *KarmaSource) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *KarmaSource) GetCount() *types.BigUInt {
	if m != nil {
		return m.Count
	}
	return nil
}

type KarmaUpkeepParams struct {
	Cost                 int64    `protobuf:"varint,1,opt,name=cost,proto3" json:"cost,omitempty"`
	Period               int64    `protobuf:"varint,3,opt,name=period,proto3" json:"period,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KarmaUpkeepParams) Reset()         { *m = KarmaUpkeepParams{} }
func (m *KarmaUpkeepParams) String() string { return proto.CompactTextString(m) }
func (*KarmaUpkeepParams) ProtoMessage()    {}
func (*KarmaUpkeepParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_karma_2da51ebadefe591e, []int{8}
}
func (m *KarmaUpkeepParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KarmaUpkeepParams.Unmarshal(m, b)
}
func (m *KarmaUpkeepParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KarmaUpkeepParams.Marshal(b, m, deterministic)
}
func (dst *KarmaUpkeepParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KarmaUpkeepParams.Merge(dst, src)
}
func (m *KarmaUpkeepParams) XXX_Size() int {
	return xxx_messageInfo_KarmaUpkeepParams.Size(m)
}
func (m *KarmaUpkeepParams) XXX_DiscardUnknown() {
	xxx_messageInfo_KarmaUpkeepParams.DiscardUnknown(m)
}

var xxx_messageInfo_KarmaUpkeepParams proto.InternalMessageInfo

func (m *KarmaUpkeepParams) GetCost() int64 {
	if m != nil {
		return m.Cost
	}
	return 0
}

func (m *KarmaUpkeepParams) GetPeriod() int64 {
	if m != nil {
		return m.Period
	}
	return 0
}

type KarmaAddressSource struct {
	User                 *types.Address `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
	Sources              []*KarmaSource `protobuf:"bytes,2,rep,name=sources" json:"sources,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *KarmaAddressSource) Reset()         { *m = KarmaAddressSource{} }
func (m *KarmaAddressSource) String() string { return proto.CompactTextString(m) }
func (*KarmaAddressSource) ProtoMessage()    {}
func (*KarmaAddressSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_karma_2da51ebadefe591e, []int{9}
}
func (m *KarmaAddressSource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KarmaAddressSource.Unmarshal(m, b)
}
func (m *KarmaAddressSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KarmaAddressSource.Marshal(b, m, deterministic)
}
func (dst *KarmaAddressSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KarmaAddressSource.Merge(dst, src)
}
func (m *KarmaAddressSource) XXX_Size() int {
	return xxx_messageInfo_KarmaAddressSource.Size(m)
}
func (m *KarmaAddressSource) XXX_DiscardUnknown() {
	xxx_messageInfo_KarmaAddressSource.DiscardUnknown(m)
}

var xxx_messageInfo_KarmaAddressSource proto.InternalMessageInfo

func (m *KarmaAddressSource) GetUser() *types.Address {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *KarmaAddressSource) GetSources() []*KarmaSource {
	if m != nil {
		return m.Sources
	}
	return nil
}

type KarmaState struct {
	SourceStates         []*KarmaSource `protobuf:"bytes,1,rep,name=source_states,json=sourceStates" json:"source_states,omitempty"`
	DeployKarmaTotal     *types.BigUInt `protobuf:"bytes,2,opt,name=deploy_karma_total,json=deployKarmaTotal" json:"deploy_karma_total,omitempty"`
	CallKarmaTotal       *types.BigUInt `protobuf:"bytes,3,opt,name=call_karma_total,json=callKarmaTotal" json:"call_karma_total,omitempty"`
	LastUpdateTime       int64          `protobuf:"varint,4,opt,name=last_update_time,json=lastUpdateTime,proto3" json:"last_update_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *KarmaState) Reset()         { *m = KarmaState{} }
func (m *KarmaState) String() string { return proto.CompactTextString(m) }
func (*KarmaState) ProtoMessage()    {}
func (*KarmaState) Descriptor() ([]byte, []int) {
	return fileDescriptor_karma_2da51ebadefe591e, []int{10}
}
func (m *KarmaState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KarmaState.Unmarshal(m, b)
}
func (m *KarmaState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KarmaState.Marshal(b, m, deterministic)
}
func (dst *KarmaState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KarmaState.Merge(dst, src)
}
func (m *KarmaState) XXX_Size() int {
	return xxx_messageInfo_KarmaState.Size(m)
}
func (m *KarmaState) XXX_DiscardUnknown() {
	xxx_messageInfo_KarmaState.DiscardUnknown(m)
}

var xxx_messageInfo_KarmaState proto.InternalMessageInfo

func (m *KarmaState) GetSourceStates() []*KarmaSource {
	if m != nil {
		return m.SourceStates
	}
	return nil
}

func (m *KarmaState) GetDeployKarmaTotal() *types.BigUInt {
	if m != nil {
		return m.DeployKarmaTotal
	}
	return nil
}

func (m *KarmaState) GetCallKarmaTotal() *types.BigUInt {
	if m != nil {
		return m.CallKarmaTotal
	}
	return nil
}

func (m *KarmaState) GetLastUpdateTime() int64 {
	if m != nil {
		return m.LastUpdateTime
	}
	return 0
}

type KarmaStateKeyUser struct {
	StateKeys            []string       `protobuf:"bytes,1,rep,name=state_keys,json=stateKeys" json:"state_keys,omitempty"`
	User                 *types.Address `protobuf:"bytes,2,opt,name=user" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *KarmaStateKeyUser) Reset()         { *m = KarmaStateKeyUser{} }
func (m *KarmaStateKeyUser) String() string { return proto.CompactTextString(m) }
func (*KarmaStateKeyUser) ProtoMessage()    {}
func (*KarmaStateKeyUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_karma_2da51ebadefe591e, []int{11}
}
func (m *KarmaStateKeyUser) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KarmaStateKeyUser.Unmarshal(m, b)
}
func (m *KarmaStateKeyUser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KarmaStateKeyUser.Marshal(b, m, deterministic)
}
func (dst *KarmaStateKeyUser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KarmaStateKeyUser.Merge(dst, src)
}
func (m *KarmaStateKeyUser) XXX_Size() int {
	return xxx_messageInfo_KarmaStateKeyUser.Size(m)
}
func (m *KarmaStateKeyUser) XXX_DiscardUnknown() {
	xxx_messageInfo_KarmaStateKeyUser.DiscardUnknown(m)
}

var xxx_messageInfo_KarmaStateKeyUser proto.InternalMessageInfo

func (m *KarmaStateKeyUser) GetStateKeys() []string {
	if m != nil {
		return m.StateKeys
	}
	return nil
}

func (m *KarmaStateKeyUser) GetUser() *types.Address {
	if m != nil {
		return m.User
	}
	return nil
}

type KarmaTotal struct {
	Count                *types.BigUInt `protobuf:"bytes,1,opt,name=count" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *KarmaTotal) Reset()         { *m = KarmaTotal{} }
func (m *KarmaTotal) String() string { return proto.CompactTextString(m) }
func (*KarmaTotal) ProtoMessage()    {}
func (*KarmaTotal) Descriptor() ([]byte, []int) {
	return fileDescriptor_karma_2da51ebadefe591e, []int{12}
}
func (m *KarmaTotal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KarmaTotal.Unmarshal(m, b)
}
func (m *KarmaTotal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KarmaTotal.Marshal(b, m, deterministic)
}
func (dst *KarmaTotal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KarmaTotal.Merge(dst, src)
}
func (m *KarmaTotal) XXX_Size() int {
	return xxx_messageInfo_KarmaTotal.Size(m)
}
func (m *KarmaTotal) XXX_DiscardUnknown() {
	xxx_messageInfo_KarmaTotal.DiscardUnknown(m)
}

var xxx_messageInfo_KarmaTotal proto.InternalMessageInfo

func (m *KarmaTotal) GetCount() *types.BigUInt {
	if m != nil {
		return m.Count
	}
	return nil
}

type ContractRecord struct {
	Owner                *types.Address `protobuf:"bytes,1,opt,name=owner" json:"owner,omitempty"`
	Address              *types.Address `protobuf:"bytes,2,opt,name=address" json:"address,omitempty"`
	CreationBlock        int64          `protobuf:"varint,3,opt,name=creation_block,json=creationBlock,proto3" json:"creation_block,omitempty"`
	Nonce                int64          `protobuf:"varint,4,opt,name=nonce,proto3" json:"nonce,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ContractRecord) Reset()         { *m = ContractRecord{} }
func (m *ContractRecord) String() string { return proto.CompactTextString(m) }
func (*ContractRecord) ProtoMessage()    {}
func (*ContractRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_karma_2da51ebadefe591e, []int{13}
}
func (m *ContractRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContractRecord.Unmarshal(m, b)
}
func (m *ContractRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContractRecord.Marshal(b, m, deterministic)
}
func (dst *ContractRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContractRecord.Merge(dst, src)
}
func (m *ContractRecord) XXX_Size() int {
	return xxx_messageInfo_ContractRecord.Size(m)
}
func (m *ContractRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_ContractRecord.DiscardUnknown(m)
}

var xxx_messageInfo_ContractRecord proto.InternalMessageInfo

func (m *ContractRecord) GetOwner() *types.Address {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *ContractRecord) GetAddress() *types.Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *ContractRecord) GetCreationBlock() int64 {
	if m != nil {
		return m.CreationBlock
	}
	return 0
}

func (m *ContractRecord) GetNonce() int64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

type GetConfigRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetConfigRequest) Reset()         { *m = GetConfigRequest{} }
func (m *GetConfigRequest) String() string { return proto.CompactTextString(m) }
func (*GetConfigRequest) ProtoMessage()    {}
func (*GetConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_karma_2da51ebadefe591e, []int{14}
}
func (m *GetConfigRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetConfigRequest.Unmarshal(m, b)
}
func (m *GetConfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetConfigRequest.Marshal(b, m, deterministic)
}
func (dst *GetConfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetConfigRequest.Merge(dst, src)
}
func (m *GetConfigRequest) XXX_Size() int {
	return xxx_messageInfo_GetConfigRequest.Size(m)
}
func (m *GetConfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetConfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetConfigRequest proto.InternalMessageInfo

type GetSourceRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSourceRequest) Reset()         { *m = GetSourceRequest{} }
func (m *GetSourceRequest) String() string { return proto.CompactTextString(m) }
func (*GetSourceRequest) ProtoMessage()    {}
func (*GetSourceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_karma_2da51ebadefe591e, []int{15}
}
func (m *GetSourceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSourceRequest.Unmarshal(m, b)
}
func (m *GetSourceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSourceRequest.Marshal(b, m, deterministic)
}
func (dst *GetSourceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSourceRequest.Merge(dst, src)
}
func (m *GetSourceRequest) XXX_Size() int {
	return xxx_messageInfo_GetSourceRequest.Size(m)
}
func (m *GetSourceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSourceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetSourceRequest proto.InternalMessageInfo

type AddKarmaRequest struct {
	User                 *types.Address `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
	KarmaSources         []*KarmaSource `protobuf:"bytes,2,rep,name=karma_sources,json=karmaSources" json:"karma_sources,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *AddKarmaRequest) Reset()         { *m = AddKarmaRequest{} }
func (m *AddKarmaRequest) String() string { return proto.CompactTextString(m) }
func (*AddKarmaRequest) ProtoMessage()    {}
func (*AddKarmaRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_karma_2da51ebadefe591e, []int{16}
}
func (m *AddKarmaRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddKarmaRequest.Unmarshal(m, b)
}
func (m *AddKarmaRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddKarmaRequest.Marshal(b, m, deterministic)
}
func (dst *AddKarmaRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddKarmaRequest.Merge(dst, src)
}
func (m *AddKarmaRequest) XXX_Size() int {
	return xxx_messageInfo_AddKarmaRequest.Size(m)
}
func (m *AddKarmaRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddKarmaRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddKarmaRequest proto.InternalMessageInfo

func (m *AddKarmaRequest) GetUser() *types.Address {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *AddKarmaRequest) GetKarmaSources() []*KarmaSource {
	if m != nil {
		return m.KarmaSources
	}
	return nil
}

func init() {
	proto.RegisterType((*KarmaInitRequest)(nil), "karma.KarmaInitRequest")
	proto.RegisterType((*KarmaConfig)(nil), "karma.KarmaConfig")
	proto.RegisterType((*KarmaSources)(nil), "karma.KarmaSources")
	proto.RegisterType((*KarmaNewOracle)(nil), "karma.KarmaNewOracle")
	proto.RegisterType((*KarmaUserTarget)(nil), "karma.KarmaUserTarget")
	proto.RegisterType((*KarmaUserAmount)(nil), "karma.KarmaUserAmount")
	proto.RegisterType((*KarmaSourceReward)(nil), "karma.KarmaSourceReward")
	proto.RegisterType((*KarmaSource)(nil), "karma.KarmaSource")
	proto.RegisterType((*KarmaUpkeepParams)(nil), "karma.KarmaUpkeepParams")
	proto.RegisterType((*KarmaAddressSource)(nil), "karma.KarmaAddressSource")
	proto.RegisterType((*KarmaState)(nil), "karma.KarmaState")
	proto.RegisterType((*KarmaStateKeyUser)(nil), "karma.KarmaStateKeyUser")
	proto.RegisterType((*KarmaTotal)(nil), "karma.KarmaTotal")
	proto.RegisterType((*ContractRecord)(nil), "karma.ContractRecord")
	proto.RegisterType((*GetConfigRequest)(nil), "karma.GetConfigRequest")
	proto.RegisterType((*GetSourceRequest)(nil), "karma.GetSourceRequest")
	proto.RegisterType((*AddKarmaRequest)(nil), "karma.AddKarmaRequest")
	proto.RegisterEnum("karma.KarmaSourceTarget", KarmaSourceTarget_name, KarmaSourceTarget_value)
}

func init() {
	proto.RegisterFile("github.com/loomnetwork/go-loom/builtin/types/karma/karma.proto", fileDescriptor_karma_2da51ebadefe591e)
}

var fileDescriptor_karma_2da51ebadefe591e = []byte{
	// 738 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0x51, 0x6f, 0xd3, 0x3a,
	0x18, 0xbd, 0x59, 0xda, 0x6e, 0xfd, 0xb6, 0x75, 0xb9, 0xbe, 0x57, 0xa8, 0x20, 0x40, 0x55, 0x24,
	0x44, 0x99, 0xb6, 0x76, 0x2a, 0x12, 0x13, 0x12, 0x02, 0x75, 0x1b, 0x42, 0xd3, 0x26, 0x36, 0xcc,
	0xf6, 0xc0, 0x53, 0x70, 0x13, 0xd3, 0x45, 0x4d, 0xec, 0xce, 0x76, 0x54, 0xf5, 0x67, 0xf0, 0xfb,
	0xf8, 0x21, 0xbc, 0xa2, 0xd8, 0x4e, 0x95, 0xad, 0x11, 0xdd, 0x4b, 0x15, 0x9f, 0xef, 0x1c, 0xfb,
	0xf3, 0xf1, 0xb1, 0x0b, 0xef, 0xc7, 0xb1, 0xba, 0xc9, 0x46, 0xbd, 0x90, 0xa7, 0xfd, 0x84, 0xf3,
	0x94, 0x51, 0x35, 0xe3, 0x62, 0xd2, 0x1f, 0xf3, 0xfd, 0x7c, 0xd8, 0x1f, 0x65, 0x71, 0xa2, 0x62,
	0xd6, 0x57, 0xf3, 0x29, 0x95, 0xfd, 0x09, 0x11, 0x29, 0x31, 0xbf, 0xbd, 0xa9, 0xe0, 0x8a, 0xa3,
	0xba, 0x1e, 0x3c, 0x39, 0x58, 0x31, 0x8d, 0x91, 0xeb, 0x5f, 0x23, 0xf4, 0x7f, 0x3b, 0xe0, 0x9d,
	0xe5, 0xda, 0x53, 0x16, 0x2b, 0x4c, 0x6f, 0x33, 0x2a, 0x15, 0xea, 0x40, 0xe3, 0x42, 0x90, 0x30,
	0xa1, 0x6d, 0xa7, 0xe3, 0x74, 0x37, 0x07, 0x1b, 0xbd, 0x61, 0x14, 0x09, 0x2a, 0x25, 0xb6, 0x38,
	0x1a, 0xc0, 0xba, 0xe4, 0x99, 0x08, 0xa9, 0x6c, 0xaf, 0x75, 0xdc, 0xee, 0xe6, 0xa0, 0xdd, 0x33,
	0xed, 0xe8, 0xb9, 0xbe, 0xea, 0x12, 0xa6, 0x33, 0x22, 0x22, 0x5c, 0x10, 0x51, 0x1f, 0xea, 0x99,
	0xa4, 0x42, 0xb6, 0x5d, 0xad, 0x78, 0x5c, 0x56, 0xd8, 0xf9, 0xad, 0xd0, 0xf0, 0xd0, 0x01, 0x34,
	0xb2, 0xe9, 0x84, 0xd2, 0x69, 0xbb, 0xa6, 0xdb, 0xb8, 0xb3, 0xc6, 0xb5, 0xae, 0x5c, 0x12, 0x41,
	0x52, 0x89, 0x2d, 0x0f, 0xed, 0x42, 0x23, 0xe4, 0xec, 0x47, 0x3c, 0x6e, 0xd7, 0xb5, 0x02, 0x95,
	0x15, 0xc7, 0xba, 0x82, 0x2d, 0xc3, 0x7f, 0x07, 0x9b, 0x25, 0x18, 0xed, 0xc3, 0x7f, 0x69, 0xcc,
	0x02, 0xcd, 0x0f, 0x14, 0x0f, 0x22, 0x3a, 0x4d, 0xf8, 0x5c, 0x1b, 0xe0, 0x62, 0x2f, 0x8d, 0x99,
	0x26, 0x5f, 0xf1, 0x13, 0x8d, 0xfb, 0x47, 0xb0, 0x55, 0xda, 0xaa, 0x2c, 0x1b, 0xe2, 0x3c, 0xd0,
	0x10, 0xff, 0x2d, 0xb4, 0x74, 0xf5, 0x33, 0x9d, 0x59, 0x5b, 0x5f, 0x02, 0x30, 0x3a, 0x0b, 0x78,
	0xb5, 0xf9, 0x4d, 0x56, 0x10, 0x7d, 0x02, 0x3b, 0xc6, 0x05, 0x49, 0xc5, 0x15, 0x11, 0x63, 0xaa,
	0xd0, 0x53, 0xa8, 0xe5, 0xb6, 0x2d, 0xa9, 0x34, 0x9a, 0x7b, 0xa9, 0x34, 0xaf, 0xbd, 0xd6, 0x71,
	0xba, 0xad, 0xaa, 0xf6, 0xcc, 0x3c, 0xd8, 0xf2, 0xfc, 0x2f, 0xa5, 0x25, 0x86, 0x29, 0xcf, 0xd8,
	0xaa, 0x25, 0x3a, 0xd0, 0x20, 0x9a, 0xa7, 0x97, 0xc8, 0xeb, 0x47, 0xf1, 0xf8, 0xfa, 0x94, 0x29,
	0x6c, 0x71, 0xff, 0x16, 0xfe, 0x5d, 0xb2, 0x03, 0x21, 0xa8, 0x31, 0x92, 0x9a, 0xdd, 0x36, 0xb1,
	0xfe, 0x46, 0x8f, 0xa0, 0x21, 0x74, 0x55, 0x4f, 0xe5, 0x62, 0x3b, 0x2a, 0xed, 0xc2, 0x7d, 0xe0,
	0x2e, 0x86, 0xf6, 0x94, 0x4d, 0xb1, 0x72, 0xb1, 0xe7, 0x50, 0x0f, 0x2b, 0xdb, 0x36, 0xb0, 0xff,
	0xc1, 0x76, 0x5d, 0x4e, 0x5c, 0x3e, 0x51, 0xc8, 0xa5, 0xb2, 0xf9, 0xd0, 0xdf, 0x79, 0xd7, 0x53,
	0x2a, 0x62, 0x1e, 0xe9, 0xee, 0x5c, 0x6c, 0x47, 0xfe, 0x77, 0x40, 0xcb, 0x21, 0x5f, 0x61, 0xe6,
	0xde, 0xfd, 0x0b, 0x86, 0x2a, 0xf2, 0xb4, 0x48, 0xd2, 0x2f, 0x07, 0xc0, 0x14, 0x14, 0x51, 0x14,
	0x1d, 0xc2, 0xb6, 0xa9, 0x04, 0x32, 0x1f, 0x17, 0x91, 0xac, 0x9a, 0x62, 0xcb, 0x10, 0xb5, 0x4e,
	0xa2, 0x37, 0x80, 0x4c, 0xee, 0x17, 0xf7, 0x40, 0x91, 0x64, 0xc9, 0x17, 0xcf, 0x70, 0xec, 0x85,
	0x50, 0x24, 0x41, 0x03, 0xf0, 0x42, 0x92, 0x24, 0x77, 0x54, 0xee, 0x3d, 0x55, 0x2b, 0x67, 0x94,
	0x34, 0x5d, 0xf0, 0x12, 0x22, 0x55, 0x90, 0x4d, 0x23, 0xa2, 0x68, 0xa0, 0xe2, 0x94, 0xea, 0x7b,
	0xee, 0xe2, 0x56, 0x8e, 0x5f, 0x6b, 0xf8, 0x2a, 0x4e, 0xa9, 0x7f, 0x59, 0xc4, 0x26, 0x6f, 0xf2,
	0x8c, 0xce, 0xf3, 0x44, 0xa2, 0x67, 0x00, 0x7a, 0x73, 0xc1, 0x84, 0xce, 0xcd, 0x06, 0x9b, 0xb8,
	0x29, 0x2d, 0x43, 0x2e, 0xdc, 0x5d, 0xab, 0x72, 0xd7, 0xdf, 0xb3, 0x76, 0x99, 0x4e, 0x16, 0x01,
	0x70, 0xaa, 0x03, 0xf0, 0xd3, 0x81, 0xd6, 0x31, 0x67, 0x4a, 0x90, 0x50, 0x61, 0x1a, 0x72, 0x11,
	0xe5, 0x12, 0x3e, 0x63, 0x15, 0xa7, 0x67, 0x60, 0xe4, 0xc3, 0x3a, 0x31, 0xc8, 0x52, 0x07, 0x45,
	0x01, 0xbd, 0x80, 0x56, 0x28, 0x28, 0x51, 0x31, 0x67, 0xc1, 0x28, 0xe1, 0xe1, 0xc4, 0xc6, 0x66,
	0xbb, 0x40, 0x8f, 0x72, 0x10, 0xfd, 0x0f, 0x75, 0xc6, 0x59, 0x58, 0x98, 0x63, 0x06, 0x3e, 0x02,
	0xef, 0x13, 0x55, 0xf6, 0x49, 0x33, 0xcf, 0xb6, 0xc5, 0x8a, 0xcb, 0x65, 0xb0, 0x1b, 0xd8, 0x19,
	0x46, 0x91, 0xde, 0x6c, 0xf1, 0xba, 0xff, 0x3d, 0x78, 0x87, 0xb0, 0x6d, 0x4e, 0x71, 0x75, 0xfc,
	0xb6, 0x26, 0xa5, 0x17, 0x70, 0xf7, 0xd5, 0x9d, 0xcb, 0x6d, 0x1f, 0xa5, 0x0d, 0xa8, 0x1d, 0x0f,
	0xcf, 0xcf, 0xbd, 0x7f, 0x10, 0x40, 0xe3, 0xe4, 0xe3, 0xe5, 0xf9, 0xc5, 0x37, 0xcf, 0x19, 0x35,
	0xf4, 0x7f, 0xcf, 0xeb, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x6b, 0xb3, 0xc9, 0x15, 0xf6, 0x06,
	0x00, 0x00,
}
