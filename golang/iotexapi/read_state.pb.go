// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/api/read_state.proto

package iotexapi

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

type ReadStakingDataMethod_Name int32

const (
	ReadStakingDataMethod_INVALID              ReadStakingDataMethod_Name = 0
	ReadStakingDataMethod_BUCKETS              ReadStakingDataMethod_Name = 1
	ReadStakingDataMethod_BUCKETS_BY_VOTER     ReadStakingDataMethod_Name = 2
	ReadStakingDataMethod_BUCKETS_BY_CANDIDATE ReadStakingDataMethod_Name = 3
	ReadStakingDataMethod_BUCKET_BY_INDEX      ReadStakingDataMethod_Name = 4
	ReadStakingDataMethod_CANDIDATES           ReadStakingDataMethod_Name = 5
	ReadStakingDataMethod_CANDIDATE_BY_NAME    ReadStakingDataMethod_Name = 6
)

var ReadStakingDataMethod_Name_name = map[int32]string{
	0: "INVALID",
	1: "BUCKETS",
	2: "BUCKETS_BY_VOTER",
	3: "BUCKETS_BY_CANDIDATE",
	4: "BUCKET_BY_INDEX",
	5: "CANDIDATES",
	6: "CANDIDATE_BY_NAME",
}

var ReadStakingDataMethod_Name_value = map[string]int32{
	"INVALID":              0,
	"BUCKETS":              1,
	"BUCKETS_BY_VOTER":     2,
	"BUCKETS_BY_CANDIDATE": 3,
	"BUCKET_BY_INDEX":      4,
	"CANDIDATES":           5,
	"CANDIDATE_BY_NAME":    6,
}

func (x ReadStakingDataMethod_Name) String() string {
	return proto.EnumName(ReadStakingDataMethod_Name_name, int32(x))
}

func (ReadStakingDataMethod_Name) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a08103f271b0c8be, []int{1, 0}
}

type PaginationParam struct {
	Offset               uint32   `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit                uint32   `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PaginationParam) Reset()         { *m = PaginationParam{} }
func (m *PaginationParam) String() string { return proto.CompactTextString(m) }
func (*PaginationParam) ProtoMessage()    {}
func (*PaginationParam) Descriptor() ([]byte, []int) {
	return fileDescriptor_a08103f271b0c8be, []int{0}
}

func (m *PaginationParam) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaginationParam.Unmarshal(m, b)
}
func (m *PaginationParam) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaginationParam.Marshal(b, m, deterministic)
}
func (m *PaginationParam) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaginationParam.Merge(m, src)
}
func (m *PaginationParam) XXX_Size() int {
	return xxx_messageInfo_PaginationParam.Size(m)
}
func (m *PaginationParam) XXX_DiscardUnknown() {
	xxx_messageInfo_PaginationParam.DiscardUnknown(m)
}

var xxx_messageInfo_PaginationParam proto.InternalMessageInfo

func (m *PaginationParam) GetOffset() uint32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *PaginationParam) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type ReadStakingDataMethod struct {
	Method               ReadStakingDataMethod_Name `protobuf:"varint,1,opt,name=method,proto3,enum=iotexapi.ReadStakingDataMethod_Name" json:"method,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *ReadStakingDataMethod) Reset()         { *m = ReadStakingDataMethod{} }
func (m *ReadStakingDataMethod) String() string { return proto.CompactTextString(m) }
func (*ReadStakingDataMethod) ProtoMessage()    {}
func (*ReadStakingDataMethod) Descriptor() ([]byte, []int) {
	return fileDescriptor_a08103f271b0c8be, []int{1}
}

func (m *ReadStakingDataMethod) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadStakingDataMethod.Unmarshal(m, b)
}
func (m *ReadStakingDataMethod) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadStakingDataMethod.Marshal(b, m, deterministic)
}
func (m *ReadStakingDataMethod) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadStakingDataMethod.Merge(m, src)
}
func (m *ReadStakingDataMethod) XXX_Size() int {
	return xxx_messageInfo_ReadStakingDataMethod.Size(m)
}
func (m *ReadStakingDataMethod) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadStakingDataMethod.DiscardUnknown(m)
}

var xxx_messageInfo_ReadStakingDataMethod proto.InternalMessageInfo

func (m *ReadStakingDataMethod) GetMethod() ReadStakingDataMethod_Name {
	if m != nil {
		return m.Method
	}
	return ReadStakingDataMethod_INVALID
}

type ReadStakingDataRequest struct {
	// Types that are valid to be assigned to Request:
	//	*ReadStakingDataRequest_Buckets
	//	*ReadStakingDataRequest_BucketsByVoter
	//	*ReadStakingDataRequest_BucketsByCandidate
	//	*ReadStakingDataRequest_BucketByIndex
	//	*ReadStakingDataRequest_Candidates_
	//	*ReadStakingDataRequest_CandidateByName_
	Request              isReadStakingDataRequest_Request `protobuf_oneof:"request"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *ReadStakingDataRequest) Reset()         { *m = ReadStakingDataRequest{} }
func (m *ReadStakingDataRequest) String() string { return proto.CompactTextString(m) }
func (*ReadStakingDataRequest) ProtoMessage()    {}
func (*ReadStakingDataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a08103f271b0c8be, []int{2}
}

func (m *ReadStakingDataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadStakingDataRequest.Unmarshal(m, b)
}
func (m *ReadStakingDataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadStakingDataRequest.Marshal(b, m, deterministic)
}
func (m *ReadStakingDataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadStakingDataRequest.Merge(m, src)
}
func (m *ReadStakingDataRequest) XXX_Size() int {
	return xxx_messageInfo_ReadStakingDataRequest.Size(m)
}
func (m *ReadStakingDataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadStakingDataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadStakingDataRequest proto.InternalMessageInfo

type isReadStakingDataRequest_Request interface {
	isReadStakingDataRequest_Request()
}

type ReadStakingDataRequest_Buckets struct {
	Buckets *ReadStakingDataRequest_VoteBuckets `protobuf:"bytes,1,opt,name=buckets,proto3,oneof"`
}

type ReadStakingDataRequest_BucketsByVoter struct {
	BucketsByVoter *ReadStakingDataRequest_VoteBucketsByVoter `protobuf:"bytes,2,opt,name=bucketsByVoter,proto3,oneof"`
}

type ReadStakingDataRequest_BucketsByCandidate struct {
	BucketsByCandidate *ReadStakingDataRequest_VoteBucketsByCandidate `protobuf:"bytes,3,opt,name=bucketsByCandidate,proto3,oneof"`
}

type ReadStakingDataRequest_BucketByIndex struct {
	BucketByIndex *ReadStakingDataRequest_VoteBucketByIndex `protobuf:"bytes,4,opt,name=bucketByIndex,proto3,oneof"`
}

type ReadStakingDataRequest_Candidates_ struct {
	Candidates *ReadStakingDataRequest_Candidates `protobuf:"bytes,5,opt,name=candidates,proto3,oneof"`
}

type ReadStakingDataRequest_CandidateByName_ struct {
	CandidateByName *ReadStakingDataRequest_CandidateByName `protobuf:"bytes,6,opt,name=candidateByName,proto3,oneof"`
}

func (*ReadStakingDataRequest_Buckets) isReadStakingDataRequest_Request() {}

func (*ReadStakingDataRequest_BucketsByVoter) isReadStakingDataRequest_Request() {}

func (*ReadStakingDataRequest_BucketsByCandidate) isReadStakingDataRequest_Request() {}

func (*ReadStakingDataRequest_BucketByIndex) isReadStakingDataRequest_Request() {}

func (*ReadStakingDataRequest_Candidates_) isReadStakingDataRequest_Request() {}

func (*ReadStakingDataRequest_CandidateByName_) isReadStakingDataRequest_Request() {}

func (m *ReadStakingDataRequest) GetRequest() isReadStakingDataRequest_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *ReadStakingDataRequest) GetBuckets() *ReadStakingDataRequest_VoteBuckets {
	if x, ok := m.GetRequest().(*ReadStakingDataRequest_Buckets); ok {
		return x.Buckets
	}
	return nil
}

func (m *ReadStakingDataRequest) GetBucketsByVoter() *ReadStakingDataRequest_VoteBucketsByVoter {
	if x, ok := m.GetRequest().(*ReadStakingDataRequest_BucketsByVoter); ok {
		return x.BucketsByVoter
	}
	return nil
}

func (m *ReadStakingDataRequest) GetBucketsByCandidate() *ReadStakingDataRequest_VoteBucketsByCandidate {
	if x, ok := m.GetRequest().(*ReadStakingDataRequest_BucketsByCandidate); ok {
		return x.BucketsByCandidate
	}
	return nil
}

func (m *ReadStakingDataRequest) GetBucketByIndex() *ReadStakingDataRequest_VoteBucketByIndex {
	if x, ok := m.GetRequest().(*ReadStakingDataRequest_BucketByIndex); ok {
		return x.BucketByIndex
	}
	return nil
}

func (m *ReadStakingDataRequest) GetCandidates() *ReadStakingDataRequest_Candidates {
	if x, ok := m.GetRequest().(*ReadStakingDataRequest_Candidates_); ok {
		return x.Candidates
	}
	return nil
}

func (m *ReadStakingDataRequest) GetCandidateByName() *ReadStakingDataRequest_CandidateByName {
	if x, ok := m.GetRequest().(*ReadStakingDataRequest_CandidateByName_); ok {
		return x.CandidateByName
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ReadStakingDataRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ReadStakingDataRequest_Buckets)(nil),
		(*ReadStakingDataRequest_BucketsByVoter)(nil),
		(*ReadStakingDataRequest_BucketsByCandidate)(nil),
		(*ReadStakingDataRequest_BucketByIndex)(nil),
		(*ReadStakingDataRequest_Candidates_)(nil),
		(*ReadStakingDataRequest_CandidateByName_)(nil),
	}
}

type ReadStakingDataRequest_VoteBuckets struct {
	Pagination           *PaginationParam `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ReadStakingDataRequest_VoteBuckets) Reset()         { *m = ReadStakingDataRequest_VoteBuckets{} }
func (m *ReadStakingDataRequest_VoteBuckets) String() string { return proto.CompactTextString(m) }
func (*ReadStakingDataRequest_VoteBuckets) ProtoMessage()    {}
func (*ReadStakingDataRequest_VoteBuckets) Descriptor() ([]byte, []int) {
	return fileDescriptor_a08103f271b0c8be, []int{2, 0}
}

func (m *ReadStakingDataRequest_VoteBuckets) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadStakingDataRequest_VoteBuckets.Unmarshal(m, b)
}
func (m *ReadStakingDataRequest_VoteBuckets) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadStakingDataRequest_VoteBuckets.Marshal(b, m, deterministic)
}
func (m *ReadStakingDataRequest_VoteBuckets) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadStakingDataRequest_VoteBuckets.Merge(m, src)
}
func (m *ReadStakingDataRequest_VoteBuckets) XXX_Size() int {
	return xxx_messageInfo_ReadStakingDataRequest_VoteBuckets.Size(m)
}
func (m *ReadStakingDataRequest_VoteBuckets) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadStakingDataRequest_VoteBuckets.DiscardUnknown(m)
}

var xxx_messageInfo_ReadStakingDataRequest_VoteBuckets proto.InternalMessageInfo

func (m *ReadStakingDataRequest_VoteBuckets) GetPagination() *PaginationParam {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type ReadStakingDataRequest_VoteBucketsByVoter struct {
	VoterAddress         string           `protobuf:"bytes,1,opt,name=voterAddress,proto3" json:"voterAddress,omitempty"`
	Pagination           *PaginationParam `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ReadStakingDataRequest_VoteBucketsByVoter) Reset() {
	*m = ReadStakingDataRequest_VoteBucketsByVoter{}
}
func (m *ReadStakingDataRequest_VoteBucketsByVoter) String() string { return proto.CompactTextString(m) }
func (*ReadStakingDataRequest_VoteBucketsByVoter) ProtoMessage()    {}
func (*ReadStakingDataRequest_VoteBucketsByVoter) Descriptor() ([]byte, []int) {
	return fileDescriptor_a08103f271b0c8be, []int{2, 1}
}

func (m *ReadStakingDataRequest_VoteBucketsByVoter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadStakingDataRequest_VoteBucketsByVoter.Unmarshal(m, b)
}
func (m *ReadStakingDataRequest_VoteBucketsByVoter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadStakingDataRequest_VoteBucketsByVoter.Marshal(b, m, deterministic)
}
func (m *ReadStakingDataRequest_VoteBucketsByVoter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadStakingDataRequest_VoteBucketsByVoter.Merge(m, src)
}
func (m *ReadStakingDataRequest_VoteBucketsByVoter) XXX_Size() int {
	return xxx_messageInfo_ReadStakingDataRequest_VoteBucketsByVoter.Size(m)
}
func (m *ReadStakingDataRequest_VoteBucketsByVoter) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadStakingDataRequest_VoteBucketsByVoter.DiscardUnknown(m)
}

var xxx_messageInfo_ReadStakingDataRequest_VoteBucketsByVoter proto.InternalMessageInfo

func (m *ReadStakingDataRequest_VoteBucketsByVoter) GetVoterAddress() string {
	if m != nil {
		return m.VoterAddress
	}
	return ""
}

func (m *ReadStakingDataRequest_VoteBucketsByVoter) GetPagination() *PaginationParam {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type ReadStakingDataRequest_VoteBucketsByCandidate struct {
	CandName             string           `protobuf:"bytes,1,opt,name=candName,proto3" json:"candName,omitempty"`
	Pagination           *PaginationParam `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ReadStakingDataRequest_VoteBucketsByCandidate) Reset() {
	*m = ReadStakingDataRequest_VoteBucketsByCandidate{}
}
func (m *ReadStakingDataRequest_VoteBucketsByCandidate) String() string {
	return proto.CompactTextString(m)
}
func (*ReadStakingDataRequest_VoteBucketsByCandidate) ProtoMessage() {}
func (*ReadStakingDataRequest_VoteBucketsByCandidate) Descriptor() ([]byte, []int) {
	return fileDescriptor_a08103f271b0c8be, []int{2, 2}
}

func (m *ReadStakingDataRequest_VoteBucketsByCandidate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadStakingDataRequest_VoteBucketsByCandidate.Unmarshal(m, b)
}
func (m *ReadStakingDataRequest_VoteBucketsByCandidate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadStakingDataRequest_VoteBucketsByCandidate.Marshal(b, m, deterministic)
}
func (m *ReadStakingDataRequest_VoteBucketsByCandidate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadStakingDataRequest_VoteBucketsByCandidate.Merge(m, src)
}
func (m *ReadStakingDataRequest_VoteBucketsByCandidate) XXX_Size() int {
	return xxx_messageInfo_ReadStakingDataRequest_VoteBucketsByCandidate.Size(m)
}
func (m *ReadStakingDataRequest_VoteBucketsByCandidate) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadStakingDataRequest_VoteBucketsByCandidate.DiscardUnknown(m)
}

var xxx_messageInfo_ReadStakingDataRequest_VoteBucketsByCandidate proto.InternalMessageInfo

func (m *ReadStakingDataRequest_VoteBucketsByCandidate) GetCandName() string {
	if m != nil {
		return m.CandName
	}
	return ""
}

func (m *ReadStakingDataRequest_VoteBucketsByCandidate) GetPagination() *PaginationParam {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type ReadStakingDataRequest_Candidates struct {
	Pagination           *PaginationParam `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ReadStakingDataRequest_Candidates) Reset()         { *m = ReadStakingDataRequest_Candidates{} }
func (m *ReadStakingDataRequest_Candidates) String() string { return proto.CompactTextString(m) }
func (*ReadStakingDataRequest_Candidates) ProtoMessage()    {}
func (*ReadStakingDataRequest_Candidates) Descriptor() ([]byte, []int) {
	return fileDescriptor_a08103f271b0c8be, []int{2, 3}
}

func (m *ReadStakingDataRequest_Candidates) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadStakingDataRequest_Candidates.Unmarshal(m, b)
}
func (m *ReadStakingDataRequest_Candidates) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadStakingDataRequest_Candidates.Marshal(b, m, deterministic)
}
func (m *ReadStakingDataRequest_Candidates) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadStakingDataRequest_Candidates.Merge(m, src)
}
func (m *ReadStakingDataRequest_Candidates) XXX_Size() int {
	return xxx_messageInfo_ReadStakingDataRequest_Candidates.Size(m)
}
func (m *ReadStakingDataRequest_Candidates) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadStakingDataRequest_Candidates.DiscardUnknown(m)
}

var xxx_messageInfo_ReadStakingDataRequest_Candidates proto.InternalMessageInfo

func (m *ReadStakingDataRequest_Candidates) GetPagination() *PaginationParam {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type ReadStakingDataRequest_CandidateByName struct {
	CandName             string   `protobuf:"bytes,1,opt,name=candName,proto3" json:"candName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadStakingDataRequest_CandidateByName) Reset() {
	*m = ReadStakingDataRequest_CandidateByName{}
}
func (m *ReadStakingDataRequest_CandidateByName) String() string { return proto.CompactTextString(m) }
func (*ReadStakingDataRequest_CandidateByName) ProtoMessage()    {}
func (*ReadStakingDataRequest_CandidateByName) Descriptor() ([]byte, []int) {
	return fileDescriptor_a08103f271b0c8be, []int{2, 4}
}

func (m *ReadStakingDataRequest_CandidateByName) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadStakingDataRequest_CandidateByName.Unmarshal(m, b)
}
func (m *ReadStakingDataRequest_CandidateByName) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadStakingDataRequest_CandidateByName.Marshal(b, m, deterministic)
}
func (m *ReadStakingDataRequest_CandidateByName) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadStakingDataRequest_CandidateByName.Merge(m, src)
}
func (m *ReadStakingDataRequest_CandidateByName) XXX_Size() int {
	return xxx_messageInfo_ReadStakingDataRequest_CandidateByName.Size(m)
}
func (m *ReadStakingDataRequest_CandidateByName) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadStakingDataRequest_CandidateByName.DiscardUnknown(m)
}

var xxx_messageInfo_ReadStakingDataRequest_CandidateByName proto.InternalMessageInfo

func (m *ReadStakingDataRequest_CandidateByName) GetCandName() string {
	if m != nil {
		return m.CandName
	}
	return ""
}

type ReadStakingDataRequest_VoteBucketByIndex struct {
	Index                uint64   `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadStakingDataRequest_VoteBucketByIndex) Reset() {
	*m = ReadStakingDataRequest_VoteBucketByIndex{}
}
func (m *ReadStakingDataRequest_VoteBucketByIndex) String() string { return proto.CompactTextString(m) }
func (*ReadStakingDataRequest_VoteBucketByIndex) ProtoMessage()    {}
func (*ReadStakingDataRequest_VoteBucketByIndex) Descriptor() ([]byte, []int) {
	return fileDescriptor_a08103f271b0c8be, []int{2, 5}
}

func (m *ReadStakingDataRequest_VoteBucketByIndex) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadStakingDataRequest_VoteBucketByIndex.Unmarshal(m, b)
}
func (m *ReadStakingDataRequest_VoteBucketByIndex) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadStakingDataRequest_VoteBucketByIndex.Marshal(b, m, deterministic)
}
func (m *ReadStakingDataRequest_VoteBucketByIndex) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadStakingDataRequest_VoteBucketByIndex.Merge(m, src)
}
func (m *ReadStakingDataRequest_VoteBucketByIndex) XXX_Size() int {
	return xxx_messageInfo_ReadStakingDataRequest_VoteBucketByIndex.Size(m)
}
func (m *ReadStakingDataRequest_VoteBucketByIndex) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadStakingDataRequest_VoteBucketByIndex.DiscardUnknown(m)
}

var xxx_messageInfo_ReadStakingDataRequest_VoteBucketByIndex proto.InternalMessageInfo

func (m *ReadStakingDataRequest_VoteBucketByIndex) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func init() {
	proto.RegisterEnum("iotexapi.ReadStakingDataMethod_Name", ReadStakingDataMethod_Name_name, ReadStakingDataMethod_Name_value)
	proto.RegisterType((*PaginationParam)(nil), "iotexapi.PaginationParam")
	proto.RegisterType((*ReadStakingDataMethod)(nil), "iotexapi.ReadStakingDataMethod")
	proto.RegisterType((*ReadStakingDataRequest)(nil), "iotexapi.ReadStakingDataRequest")
	proto.RegisterType((*ReadStakingDataRequest_VoteBuckets)(nil), "iotexapi.ReadStakingDataRequest.VoteBuckets")
	proto.RegisterType((*ReadStakingDataRequest_VoteBucketsByVoter)(nil), "iotexapi.ReadStakingDataRequest.VoteBucketsByVoter")
	proto.RegisterType((*ReadStakingDataRequest_VoteBucketsByCandidate)(nil), "iotexapi.ReadStakingDataRequest.VoteBucketsByCandidate")
	proto.RegisterType((*ReadStakingDataRequest_Candidates)(nil), "iotexapi.ReadStakingDataRequest.Candidates")
	proto.RegisterType((*ReadStakingDataRequest_CandidateByName)(nil), "iotexapi.ReadStakingDataRequest.CandidateByName")
	proto.RegisterType((*ReadStakingDataRequest_VoteBucketByIndex)(nil), "iotexapi.ReadStakingDataRequest.VoteBucketByIndex")
}

func init() { proto.RegisterFile("proto/api/read_state.proto", fileDescriptor_a08103f271b0c8be) }

var fileDescriptor_a08103f271b0c8be = []byte{
	// 572 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x5f, 0x6f, 0x12, 0x4f,
	0x14, 0x5d, 0x28, 0xff, 0x7a, 0xf9, 0x15, 0xb6, 0xf3, 0xa3, 0x64, 0xe5, 0xc9, 0x10, 0x1f, 0x34,
	0xda, 0xc5, 0x40, 0x8c, 0x31, 0x31, 0x31, 0x2c, 0x10, 0x97, 0x28, 0xd8, 0x0c, 0x48, 0xb4, 0xd1,
	0x90, 0x81, 0x9d, 0xd2, 0xb1, 0x65, 0x17, 0x77, 0x07, 0x53, 0xbe, 0x83, 0x6f, 0x7e, 0x45, 0x3f,
	0x88, 0x99, 0x99, 0x65, 0x4b, 0xb7, 0xf8, 0x87, 0xf4, 0x6d, 0xce, 0xdd, 0x7b, 0xce, 0xb9, 0x73,
	0xef, 0xdd, 0x81, 0xca, 0xc2, 0xf7, 0xb8, 0x57, 0x23, 0x0b, 0x56, 0xf3, 0x29, 0x71, 0xc6, 0x01,
	0x27, 0x9c, 0x9a, 0x32, 0x88, 0x72, 0xcc, 0xe3, 0xf4, 0x8a, 0x2c, 0x58, 0xf5, 0x15, 0x14, 0x4f,
	0xc8, 0x8c, 0xb9, 0x84, 0x33, 0xcf, 0x3d, 0x21, 0x3e, 0x99, 0xa3, 0x32, 0x64, 0xbc, 0xb3, 0xb3,
	0x80, 0x72, 0x23, 0x71, 0x3f, 0xf1, 0xf0, 0x00, 0x87, 0x08, 0x95, 0x20, 0x7d, 0xc9, 0xe6, 0x8c,
	0x1b, 0x49, 0x19, 0x56, 0xa0, 0xfa, 0x33, 0x01, 0x47, 0x98, 0x12, 0x67, 0xc0, 0xc9, 0x05, 0x73,
	0x67, 0x6d, 0xc2, 0x49, 0x8f, 0xf2, 0x73, 0xcf, 0x41, 0x2f, 0x21, 0x33, 0x97, 0x27, 0xa9, 0x53,
	0xa8, 0x3f, 0x30, 0xd7, 0xae, 0xe6, 0x56, 0x82, 0xd9, 0x27, 0x73, 0x8a, 0x43, 0x4e, 0xf5, 0x7b,
	0x02, 0x52, 0x22, 0x80, 0xf2, 0x90, 0xed, 0xf6, 0x47, 0xcd, 0xb7, 0xdd, 0xb6, 0xae, 0x09, 0x60,
	0xbd, 0x6f, 0xbd, 0xe9, 0x0c, 0x07, 0x7a, 0x02, 0x95, 0x40, 0x0f, 0xc1, 0xd8, 0xfa, 0x38, 0x1e,
	0xbd, 0x1b, 0x76, 0xb0, 0x9e, 0x44, 0x06, 0x94, 0x36, 0xa2, 0xad, 0x66, 0xbf, 0xdd, 0x6d, 0x37,
	0x87, 0x1d, 0x7d, 0x0f, 0xfd, 0x0f, 0x45, 0xf5, 0x45, 0x7c, 0xe8, 0xf6, 0xdb, 0x9d, 0x0f, 0x7a,
	0x0a, 0x15, 0x00, 0xa2, 0x9c, 0x81, 0x9e, 0x46, 0x47, 0x70, 0x18, 0x61, 0x91, 0xd7, 0x6f, 0xf6,
	0x3a, 0x7a, 0xa6, 0xfa, 0x23, 0x07, 0xe5, 0x58, 0xd5, 0x98, 0x7e, 0x5d, 0xd2, 0x80, 0x23, 0x1b,
	0xb2, 0x93, 0xe5, 0xf4, 0x82, 0xf2, 0x40, 0x5e, 0x34, 0x5f, 0x7f, 0xf2, 0xdb, 0x8b, 0x86, 0x14,
	0x73, 0xe4, 0x71, 0x6a, 0x29, 0x8e, 0xad, 0xe1, 0x35, 0x1d, 0x7d, 0x86, 0x42, 0x78, 0xb4, 0x56,
	0x22, 0xc5, 0x97, 0xad, 0xce, 0xd7, 0x1b, 0xbb, 0x08, 0x86, 0x54, 0x5b, 0xc3, 0x31, 0x31, 0xc4,
	0x00, 0x45, 0x91, 0x16, 0x71, 0x1d, 0xe6, 0x10, 0x4e, 0x8d, 0x3d, 0x69, 0xf1, 0x7c, 0x37, 0x8b,
	0x88, 0x6e, 0x6b, 0x78, 0x8b, 0x28, 0x3a, 0x85, 0x03, 0x15, 0xb5, 0x56, 0x5d, 0xd7, 0xa1, 0x57,
	0x46, 0x4a, 0xba, 0xd4, 0x77, 0x70, 0x09, 0x99, 0xb6, 0x86, 0x6f, 0x4a, 0xa1, 0x1e, 0xc0, 0x74,
	0x6d, 0x14, 0x18, 0x69, 0x29, 0xfc, 0xf8, 0xaf, 0xc2, 0x51, 0x6d, 0xa2, 0xe3, 0x1b, 0x02, 0xe8,
	0x13, 0x14, 0x23, 0x64, 0xad, 0xc4, 0xca, 0x19, 0x19, 0xa9, 0xf9, 0xf4, 0xdf, 0x35, 0x15, 0xcf,
	0xd6, 0x70, 0x5c, 0xaa, 0x62, 0x43, 0x7e, 0xa3, 0x71, 0xe8, 0x05, 0xc0, 0x22, 0xfa, 0xdd, 0xc2,
	0x75, 0xb9, 0x77, 0xed, 0x13, 0xfb, 0x15, 0xf1, 0x46, 0x72, 0x25, 0x00, 0x74, 0x7b, 0xca, 0xa8,
	0x0a, 0xff, 0x7d, 0x13, 0x87, 0xa6, 0xe3, 0xf8, 0x34, 0x50, 0x1b, 0xb8, 0x8f, 0x6f, 0xc4, 0x62,
	0xa6, 0xc9, 0x5d, 0x4c, 0x3d, 0x28, 0x6f, 0x9f, 0x3b, 0xaa, 0x40, 0x4e, 0xdc, 0x55, 0xf6, 0x4b,
	0x99, 0x46, 0xf8, 0x2e, 0x86, 0xaf, 0x01, 0xae, 0x27, 0x75, 0x97, 0x76, 0x1d, 0x43, 0x31, 0x36,
	0x9e, 0x3f, 0x95, 0x5c, 0x79, 0x04, 0x87, 0xb7, 0x56, 0x4f, 0xbc, 0x78, 0x4c, 0x6e, 0xaf, 0xc8,
	0x4e, 0x61, 0x05, 0xac, 0x7d, 0xc8, 0xfa, 0x6a, 0x01, 0xac, 0x67, 0xa7, 0x8d, 0x19, 0xe3, 0xe7,
	0xcb, 0x89, 0x39, 0xf5, 0xe6, 0x35, 0x59, 0xd7, 0xc2, 0xf7, 0xbe, 0xd0, 0x29, 0x57, 0xe0, 0x58,
	0xbd, 0xc1, 0x33, 0xef, 0x92, 0xb8, 0xb3, 0xda, 0xba, 0xee, 0x49, 0x46, 0x86, 0x1b, 0xbf, 0x02,
	0x00, 0x00, 0xff, 0xff, 0xfe, 0x6c, 0xa4, 0x11, 0xa3, 0x05, 0x00, 0x00,
}
