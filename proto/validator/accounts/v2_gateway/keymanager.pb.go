// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/validator/accounts/v2/keymanager.proto

package ethereum_validator_accounts_v2

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	v1alpha1 "github.com/prysmaticlabs/ethereumapis/eth/v1alpha1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type SignResponse_Status int32

const (
	SignResponse_UNKNOWN   SignResponse_Status = 0
	SignResponse_SUCCEEDED SignResponse_Status = 1
	SignResponse_DENIED    SignResponse_Status = 2
	SignResponse_FAILED    SignResponse_Status = 3
)

var SignResponse_Status_name = map[int32]string{
	0: "UNKNOWN",
	1: "SUCCEEDED",
	2: "DENIED",
	3: "FAILED",
}

var SignResponse_Status_value = map[string]int32{
	"UNKNOWN":   0,
	"SUCCEEDED": 1,
	"DENIED":    2,
	"FAILED":    3,
}

func (x SignResponse_Status) String() string {
	return proto.EnumName(SignResponse_Status_name, int32(x))
}

func (SignResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_795e98bd0a473d79, []int{2, 0}
}

type ListPublicKeysResponse struct {
	ValidatingPublicKeys [][]byte `protobuf:"bytes,2,rep,name=validating_public_keys,json=validatingPublicKeys,proto3" json:"validating_public_keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListPublicKeysResponse) Reset()         { *m = ListPublicKeysResponse{} }
func (m *ListPublicKeysResponse) String() string { return proto.CompactTextString(m) }
func (*ListPublicKeysResponse) ProtoMessage()    {}
func (*ListPublicKeysResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_795e98bd0a473d79, []int{0}
}

func (m *ListPublicKeysResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPublicKeysResponse.Unmarshal(m, b)
}
func (m *ListPublicKeysResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPublicKeysResponse.Marshal(b, m, deterministic)
}
func (m *ListPublicKeysResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPublicKeysResponse.Merge(m, src)
}
func (m *ListPublicKeysResponse) XXX_Size() int {
	return xxx_messageInfo_ListPublicKeysResponse.Size(m)
}
func (m *ListPublicKeysResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPublicKeysResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListPublicKeysResponse proto.InternalMessageInfo

func (m *ListPublicKeysResponse) GetValidatingPublicKeys() [][]byte {
	if m != nil {
		return m.ValidatingPublicKeys
	}
	return nil
}

type SignRequest struct {
	PublicKey       []byte `protobuf:"bytes,1,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	SigningRoot     []byte `protobuf:"bytes,2,opt,name=signing_root,json=signingRoot,proto3" json:"signing_root,omitempty"`
	SignatureDomain []byte `protobuf:"bytes,3,opt,name=signature_domain,json=signatureDomain,proto3" json:"signature_domain,omitempty"`
	// Types that are valid to be assigned to Object:
	//	*SignRequest_Block
	//	*SignRequest_AttestationData
	//	*SignRequest_AggregateAttestationAndProof
	//	*SignRequest_Exit
	//	*SignRequest_Slot
	//	*SignRequest_Epoch
	Object               isSignRequest_Object `protobuf_oneof:"object"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *SignRequest) Reset()         { *m = SignRequest{} }
func (m *SignRequest) String() string { return proto.CompactTextString(m) }
func (*SignRequest) ProtoMessage()    {}
func (*SignRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_795e98bd0a473d79, []int{1}
}

func (m *SignRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignRequest.Unmarshal(m, b)
}
func (m *SignRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignRequest.Marshal(b, m, deterministic)
}
func (m *SignRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignRequest.Merge(m, src)
}
func (m *SignRequest) XXX_Size() int {
	return xxx_messageInfo_SignRequest.Size(m)
}
func (m *SignRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignRequest proto.InternalMessageInfo

func (m *SignRequest) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *SignRequest) GetSigningRoot() []byte {
	if m != nil {
		return m.SigningRoot
	}
	return nil
}

func (m *SignRequest) GetSignatureDomain() []byte {
	if m != nil {
		return m.SignatureDomain
	}
	return nil
}

type isSignRequest_Object interface {
	isSignRequest_Object()
}

type SignRequest_Block struct {
	Block *v1alpha1.BeaconBlock `protobuf:"bytes,101,opt,name=block,proto3,oneof"`
}

type SignRequest_AttestationData struct {
	AttestationData *v1alpha1.AttestationData `protobuf:"bytes,102,opt,name=attestation_data,json=attestationData,proto3,oneof"`
}

type SignRequest_AggregateAttestationAndProof struct {
	AggregateAttestationAndProof *v1alpha1.AggregateAttestationAndProof `protobuf:"bytes,103,opt,name=aggregate_attestation_and_proof,json=aggregateAttestationAndProof,proto3,oneof"`
}

type SignRequest_Exit struct {
	Exit *v1alpha1.VoluntaryExit `protobuf:"bytes,104,opt,name=exit,proto3,oneof"`
}

type SignRequest_Slot struct {
	Slot uint64 `protobuf:"varint,105,opt,name=slot,proto3,oneof"`
}

type SignRequest_Epoch struct {
	Epoch uint64 `protobuf:"varint,106,opt,name=epoch,proto3,oneof"`
}

func (*SignRequest_Block) isSignRequest_Object() {}

func (*SignRequest_AttestationData) isSignRequest_Object() {}

func (*SignRequest_AggregateAttestationAndProof) isSignRequest_Object() {}

func (*SignRequest_Exit) isSignRequest_Object() {}

func (*SignRequest_Slot) isSignRequest_Object() {}

func (*SignRequest_Epoch) isSignRequest_Object() {}

func (m *SignRequest) GetObject() isSignRequest_Object {
	if m != nil {
		return m.Object
	}
	return nil
}

func (m *SignRequest) GetBlock() *v1alpha1.BeaconBlock {
	if x, ok := m.GetObject().(*SignRequest_Block); ok {
		return x.Block
	}
	return nil
}

func (m *SignRequest) GetAttestationData() *v1alpha1.AttestationData {
	if x, ok := m.GetObject().(*SignRequest_AttestationData); ok {
		return x.AttestationData
	}
	return nil
}

func (m *SignRequest) GetAggregateAttestationAndProof() *v1alpha1.AggregateAttestationAndProof {
	if x, ok := m.GetObject().(*SignRequest_AggregateAttestationAndProof); ok {
		return x.AggregateAttestationAndProof
	}
	return nil
}

func (m *SignRequest) GetExit() *v1alpha1.VoluntaryExit {
	if x, ok := m.GetObject().(*SignRequest_Exit); ok {
		return x.Exit
	}
	return nil
}

func (m *SignRequest) GetSlot() uint64 {
	if x, ok := m.GetObject().(*SignRequest_Slot); ok {
		return x.Slot
	}
	return 0
}

func (m *SignRequest) GetEpoch() uint64 {
	if x, ok := m.GetObject().(*SignRequest_Epoch); ok {
		return x.Epoch
	}
	return 0
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SignRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SignRequest_Block)(nil),
		(*SignRequest_AttestationData)(nil),
		(*SignRequest_AggregateAttestationAndProof)(nil),
		(*SignRequest_Exit)(nil),
		(*SignRequest_Slot)(nil),
		(*SignRequest_Epoch)(nil),
	}
}

type SignResponse struct {
	Signature            []byte              `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	Status               SignResponse_Status `protobuf:"varint,2,opt,name=status,proto3,enum=ethereum.validator.accounts.v2.SignResponse_Status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *SignResponse) Reset()         { *m = SignResponse{} }
func (m *SignResponse) String() string { return proto.CompactTextString(m) }
func (*SignResponse) ProtoMessage()    {}
func (*SignResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_795e98bd0a473d79, []int{2}
}

func (m *SignResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignResponse.Unmarshal(m, b)
}
func (m *SignResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignResponse.Marshal(b, m, deterministic)
}
func (m *SignResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignResponse.Merge(m, src)
}
func (m *SignResponse) XXX_Size() int {
	return xxx_messageInfo_SignResponse.Size(m)
}
func (m *SignResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SignResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SignResponse proto.InternalMessageInfo

func (m *SignResponse) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *SignResponse) GetStatus() SignResponse_Status {
	if m != nil {
		return m.Status
	}
	return SignResponse_UNKNOWN
}

func init() {
	proto.RegisterEnum("ethereum.validator.accounts.v2.SignResponse_Status", SignResponse_Status_name, SignResponse_Status_value)
	proto.RegisterType((*ListPublicKeysResponse)(nil), "ethereum.validator.accounts.v2.ListPublicKeysResponse")
	proto.RegisterType((*SignRequest)(nil), "ethereum.validator.accounts.v2.SignRequest")
	proto.RegisterType((*SignResponse)(nil), "ethereum.validator.accounts.v2.SignResponse")
}

func init() {
	proto.RegisterFile("proto/validator/accounts/v2/keymanager.proto", fileDescriptor_795e98bd0a473d79)
}

var fileDescriptor_795e98bd0a473d79 = []byte{
	// 653 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x5d, 0x6e, 0xd3, 0x40,
	0x10, 0x8e, 0x93, 0x34, 0xd0, 0x4d, 0xa0, 0xd1, 0xaa, 0x8a, 0xac, 0x10, 0xda, 0x60, 0x55, 0x28,
	0x88, 0xca, 0x56, 0x53, 0xc4, 0x43, 0xc5, 0x4b, 0x5b, 0x1b, 0xa5, 0x6a, 0x15, 0x2a, 0x47, 0x2d,
	0x8f, 0xd6, 0xc6, 0x99, 0x3a, 0x6e, 0x9d, 0x5d, 0x63, 0xaf, 0xa3, 0x46, 0xe2, 0x09, 0x2e, 0x80,
	0xc4, 0x21, 0x38, 0x02, 0x07, 0xe1, 0x0a, 0x5c, 0x03, 0x09, 0xed, 0xda, 0x71, 0x52, 0x29, 0xe5,
	0xe7, 0x2d, 0xf3, 0xcd, 0xf7, 0xcd, 0x8c, 0xbf, 0xd9, 0x09, 0xda, 0x0d, 0x23, 0xc6, 0x99, 0x31,
	0x25, 0x81, 0x3f, 0x22, 0x9c, 0x45, 0x06, 0x71, 0x5d, 0x96, 0x50, 0x1e, 0x1b, 0xd3, 0xae, 0x71,
	0x03, 0xb3, 0x09, 0xa1, 0xc4, 0x83, 0x48, 0x97, 0x34, 0xbc, 0x05, 0x7c, 0x0c, 0x11, 0x24, 0x13,
	0x3d, 0x17, 0xe8, 0x73, 0x81, 0x3e, 0xed, 0x36, 0x45, 0xde, 0x98, 0xee, 0x91, 0x20, 0x1c, 0x93,
	0x3d, 0x83, 0x70, 0x0e, 0x31, 0x27, 0xdc, 0x67, 0x34, 0xd5, 0x37, 0xb7, 0xef, 0xe4, 0x87, 0x40,
	0x5c, 0x46, 0x9d, 0x61, 0xc0, 0xdc, 0x9b, 0x8c, 0xd0, 0xf2, 0x18, 0xf3, 0x02, 0x30, 0x48, 0xe8,
	0x1b, 0x84, 0x52, 0x96, 0xaa, 0xe3, 0x2c, 0xfb, 0x24, 0xcb, 0xca, 0x68, 0x98, 0x5c, 0x19, 0x30,
	0x09, 0xf9, 0x2c, 0x4d, 0x6a, 0x7d, 0xd4, 0x38, 0xf3, 0x63, 0x7e, 0x9e, 0x0c, 0x03, 0xdf, 0x3d,
	0x85, 0x59, 0x6c, 0x43, 0x1c, 0x32, 0x1a, 0x03, 0x7e, 0x85, 0x1a, 0xd9, 0xb8, 0x3e, 0xf5, 0x9c,
	0x50, 0x12, 0x9c, 0x1b, 0x98, 0xc5, 0x6a, 0xb1, 0x5d, 0xea, 0xd4, 0xec, 0xcd, 0x45, 0x76, 0xa1,
	0xd6, 0x7e, 0x95, 0x50, 0x75, 0xe0, 0x7b, 0xd4, 0x86, 0x0f, 0x09, 0xc4, 0x1c, 0x3f, 0x45, 0x68,
	0x21, 0x55, 0x95, 0xb6, 0xd2, 0xa9, 0xd9, 0xeb, 0xe1, 0x9c, 0x8f, 0x9f, 0xa1, 0x5a, 0xec, 0x7b,
	0x54, 0x74, 0x88, 0x18, 0xe3, 0x6a, 0x51, 0x12, 0xaa, 0x19, 0x66, 0x33, 0xc6, 0xf1, 0x0b, 0x54,
	0x17, 0x21, 0xe1, 0x49, 0x04, 0xce, 0x88, 0x4d, 0x88, 0x4f, 0xd5, 0x92, 0xa4, 0x6d, 0xe4, 0xb8,
	0x29, 0x61, 0x7c, 0x80, 0xd6, 0xa4, 0x2d, 0x2a, 0xb4, 0x95, 0x4e, 0xb5, 0xab, 0xe9, 0xb9, 0xf1,
	0xc0, 0xc7, 0xfa, 0xdc, 0x41, 0xfd, 0x48, 0x3a, 0x78, 0x24, 0x98, 0xbd, 0x82, 0x9d, 0x4a, 0xf0,
	0x00, 0xd5, 0x97, 0x9c, 0x77, 0x46, 0x84, 0x13, 0xf5, 0x4a, 0x96, 0x79, 0x7e, 0x4f, 0x99, 0xc3,
	0x05, 0xdd, 0x24, 0x9c, 0xf4, 0x0a, 0xf6, 0x06, 0xb9, 0x0b, 0xe1, 0x8f, 0x68, 0x9b, 0x78, 0x5e,
	0x04, 0x1e, 0xe1, 0xe0, 0x2c, 0x97, 0x27, 0x74, 0xe4, 0x84, 0x11, 0x63, 0x57, 0xaa, 0x27, 0x7b,
	0xec, 0xdf, 0xd7, 0x63, 0xae, 0x5e, 0x6a, 0x76, 0x48, 0x47, 0xe7, 0x42, 0xda, 0x2b, 0xd8, 0x2d,
	0xf2, 0x87, 0x3c, 0x3e, 0x40, 0x65, 0xb8, 0xf5, 0xb9, 0x3a, 0x96, 0x2d, 0x76, 0xee, 0x69, 0x71,
	0xc9, 0x82, 0x84, 0x72, 0x12, 0xcd, 0xac, 0x5b, 0x9f, 0xf7, 0x0a, 0xb6, 0xd4, 0xe0, 0x4d, 0x54,
	0x8e, 0x03, 0xc6, 0x55, 0xbf, 0xad, 0x74, 0xca, 0x02, 0x15, 0x11, 0x6e, 0xa0, 0x35, 0x08, 0x99,
	0x3b, 0x56, 0xaf, 0x33, 0x38, 0x0d, 0x8f, 0x1e, 0xa2, 0x0a, 0x1b, 0x5e, 0x83, 0xcb, 0xb5, 0xef,
	0x0a, 0xaa, 0xa5, 0xfb, 0xcf, 0x9e, 0x51, 0x0b, 0xad, 0xe7, 0x6b, 0x9a, 0xef, 0x3f, 0x07, 0xf0,
	0x29, 0xaa, 0x88, 0xa9, 0x93, 0x58, 0x6e, 0xfe, 0xf1, 0xb2, 0x0f, 0x2b, 0x6f, 0x45, 0x5f, 0xae,
	0xad, 0x0f, 0xa4, 0xd4, 0xce, 0x4a, 0x68, 0x6f, 0x50, 0x25, 0x45, 0x70, 0x15, 0x3d, 0xb8, 0xe8,
	0x9f, 0xf6, 0xdf, 0xbd, 0xef, 0xd7, 0x0b, 0xf8, 0x11, 0x5a, 0x1f, 0x5c, 0x1c, 0x1f, 0x5b, 0x96,
	0x69, 0x99, 0x75, 0x05, 0x23, 0x54, 0x31, 0xad, 0xfe, 0x89, 0x65, 0xd6, 0x8b, 0xe2, 0xf7, 0xdb,
	0xc3, 0x93, 0x33, 0xcb, 0xac, 0x97, 0xba, 0xdf, 0x8a, 0xa8, 0x66, 0xc3, 0x84, 0x71, 0x10, 0x3d,
	0x20, 0xc2, 0x5f, 0x14, 0xa4, 0x8a, 0xdb, 0xb8, 0x5c, 0xf1, 0xce, 0x71, 0x43, 0x4f, 0xaf, 0x4a,
	0x9f, 0x5f, 0x95, 0x6e, 0x89, 0xab, 0x6a, 0xbe, 0xfe, 0xdb, 0x07, 0xac, 0xbe, 0x36, 0x6d, 0xe7,
	0xd3, 0x8f, 0x9f, 0x5f, 0x8b, 0x5b, 0xb8, 0x75, 0xe7, 0xaf, 0x24, 0x92, 0xf3, 0xe4, 0x10, 0xfe,
	0xac, 0xa0, 0xb2, 0x98, 0x0e, 0xbf, 0xfc, 0x37, 0x9f, 0xe4, 0x0d, 0x36, 0x77, 0xff, 0xc7, 0x54,
	0xad, 0x2d, 0x27, 0x69, 0x6a, 0xea, 0xaa, 0x49, 0xc4, 0xe6, 0x86, 0x15, 0xf9, 0xcd, 0xfb, 0xbf,
	0x03, 0x00, 0x00, 0xff, 0xff, 0x2d, 0x82, 0x15, 0x16, 0x06, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RemoteSignerClient is the client API for RemoteSigner service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RemoteSignerClient interface {
	ListValidatingPublicKeys(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ListPublicKeysResponse, error)
	Sign(ctx context.Context, in *SignRequest, opts ...grpc.CallOption) (*SignResponse, error)
}

type remoteSignerClient struct {
	cc grpc.ClientConnInterface
}

func NewRemoteSignerClient(cc grpc.ClientConnInterface) RemoteSignerClient {
	return &remoteSignerClient{cc}
}

func (c *remoteSignerClient) ListValidatingPublicKeys(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ListPublicKeysResponse, error) {
	out := new(ListPublicKeysResponse)
	err := c.cc.Invoke(ctx, "/ethereum.validator.accounts.v2.RemoteSigner/ListValidatingPublicKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteSignerClient) Sign(ctx context.Context, in *SignRequest, opts ...grpc.CallOption) (*SignResponse, error) {
	out := new(SignResponse)
	err := c.cc.Invoke(ctx, "/ethereum.validator.accounts.v2.RemoteSigner/Sign", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RemoteSignerServer is the server API for RemoteSigner service.
type RemoteSignerServer interface {
	ListValidatingPublicKeys(context.Context, *empty.Empty) (*ListPublicKeysResponse, error)
	Sign(context.Context, *SignRequest) (*SignResponse, error)
}

// UnimplementedRemoteSignerServer can be embedded to have forward compatible implementations.
type UnimplementedRemoteSignerServer struct {
}

func (*UnimplementedRemoteSignerServer) ListValidatingPublicKeys(ctx context.Context, req *empty.Empty) (*ListPublicKeysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListValidatingPublicKeys not implemented")
}
func (*UnimplementedRemoteSignerServer) Sign(ctx context.Context, req *SignRequest) (*SignResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sign not implemented")
}

func RegisterRemoteSignerServer(s *grpc.Server, srv RemoteSignerServer) {
	s.RegisterService(&_RemoteSigner_serviceDesc, srv)
}

func _RemoteSigner_ListValidatingPublicKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteSignerServer).ListValidatingPublicKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ethereum.validator.accounts.v2.RemoteSigner/ListValidatingPublicKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteSignerServer).ListValidatingPublicKeys(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteSigner_Sign_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteSignerServer).Sign(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ethereum.validator.accounts.v2.RemoteSigner/Sign",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteSignerServer).Sign(ctx, req.(*SignRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RemoteSigner_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ethereum.validator.accounts.v2.RemoteSigner",
	HandlerType: (*RemoteSignerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListValidatingPublicKeys",
			Handler:    _RemoteSigner_ListValidatingPublicKeys_Handler,
		},
		{
			MethodName: "Sign",
			Handler:    _RemoteSigner_Sign_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/validator/accounts/v2/keymanager.proto",
}
