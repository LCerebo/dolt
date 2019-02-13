// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dolt/services/remotesapi/v1alpha1/credentials.proto

package remotesapi

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type WhoAmIRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WhoAmIRequest) Reset()         { *m = WhoAmIRequest{} }
func (m *WhoAmIRequest) String() string { return proto.CompactTextString(m) }
func (*WhoAmIRequest) ProtoMessage()    {}
func (*WhoAmIRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_24ed1d4c5faa311b, []int{0}
}

func (m *WhoAmIRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WhoAmIRequest.Unmarshal(m, b)
}
func (m *WhoAmIRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WhoAmIRequest.Marshal(b, m, deterministic)
}
func (m *WhoAmIRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WhoAmIRequest.Merge(m, src)
}
func (m *WhoAmIRequest) XXX_Size() int {
	return xxx_messageInfo_WhoAmIRequest.Size(m)
}
func (m *WhoAmIRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WhoAmIRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WhoAmIRequest proto.InternalMessageInfo

type WhoAmIResponse struct {
	// Ex: "bheni"
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	// Ex: "Brian Hendriks"
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Ex: "brian@liquidata.co"
	EmailAddress         string   `protobuf:"bytes,3,opt,name=email_address,json=emailAddress,proto3" json:"email_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WhoAmIResponse) Reset()         { *m = WhoAmIResponse{} }
func (m *WhoAmIResponse) String() string { return proto.CompactTextString(m) }
func (*WhoAmIResponse) ProtoMessage()    {}
func (*WhoAmIResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_24ed1d4c5faa311b, []int{1}
}

func (m *WhoAmIResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WhoAmIResponse.Unmarshal(m, b)
}
func (m *WhoAmIResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WhoAmIResponse.Marshal(b, m, deterministic)
}
func (m *WhoAmIResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WhoAmIResponse.Merge(m, src)
}
func (m *WhoAmIResponse) XXX_Size() int {
	return xxx_messageInfo_WhoAmIResponse.Size(m)
}
func (m *WhoAmIResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WhoAmIResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WhoAmIResponse proto.InternalMessageInfo

func (m *WhoAmIResponse) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *WhoAmIResponse) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *WhoAmIResponse) GetEmailAddress() string {
	if m != nil {
		return m.EmailAddress
	}
	return ""
}

func init() {
	proto.RegisterType((*WhoAmIRequest)(nil), "dolt.services.remotesapi.v1alpha1.WhoAmIRequest")
	proto.RegisterType((*WhoAmIResponse)(nil), "dolt.services.remotesapi.v1alpha1.WhoAmIResponse")
}

func init() {
	proto.RegisterFile("dolt/services/remotesapi/v1alpha1/credentials.proto", fileDescriptor_24ed1d4c5faa311b)
}

var fileDescriptor_24ed1d4c5faa311b = []byte{
	// 237 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0xb1, 0x4b, 0x43, 0x31,
	0x10, 0xc6, 0x79, 0x0a, 0x45, 0xcf, 0x56, 0x21, 0x53, 0xe9, 0x64, 0xeb, 0xe2, 0x20, 0x79, 0x3e,
	0x3b, 0x3a, 0x55, 0x27, 0x17, 0x87, 0x3a, 0x08, 0x2e, 0x8f, 0xb3, 0x39, 0x68, 0x20, 0x79, 0x89,
	0xb9, 0xb4, 0xe0, 0xec, 0x3f, 0x2e, 0x2f, 0x31, 0x3e, 0x1c, 0x44, 0x1c, 0xef, 0x77, 0xf7, 0x83,
	0xef, 0x3b, 0x58, 0x2a, 0x67, 0x62, 0xcd, 0x14, 0xf6, 0x7a, 0x43, 0x5c, 0x07, 0xb2, 0x2e, 0x12,
	0xa3, 0xd7, 0xf5, 0xbe, 0x41, 0xe3, 0xb7, 0xd8, 0xd4, 0x9b, 0x40, 0x8a, 0xba, 0xa8, 0xd1, 0xb0,
	0xf4, 0xc1, 0x45, 0x27, 0xe6, 0xbd, 0x24, 0x8b, 0x24, 0x07, 0x49, 0x16, 0x69, 0x71, 0x06, 0x93,
	0xe7, 0xad, 0x5b, 0xd9, 0x87, 0x35, 0xbd, 0xed, 0x88, 0xe3, 0x22, 0xc2, 0x69, 0x01, 0xec, 0x5d,
	0xc7, 0x24, 0x66, 0x70, 0xb4, 0x63, 0x0a, 0x1d, 0x5a, 0x9a, 0x56, 0xe7, 0xd5, 0xe5, 0xf1, 0xfa,
	0x7b, 0x16, 0x73, 0x18, 0x2b, 0xcd, 0xde, 0xe0, 0x7b, 0x9b, 0xf6, 0x07, 0x69, 0x7f, 0xf2, 0xc5,
	0x1e, 0xfb, 0x93, 0x0b, 0x98, 0x90, 0x45, 0x6d, 0x5a, 0x54, 0x2a, 0x10, 0xf3, 0xf4, 0x30, 0xdd,
	0x8c, 0x13, 0x5c, 0x65, 0x76, 0xf3, 0x51, 0x81, 0xb8, 0x1f, 0xf2, 0x3f, 0xe5, 0xc8, 0xc2, 0xc2,
	0x28, 0x87, 0x11, 0xd7, 0xf2, 0xcf, 0x2e, 0xf2, 0x47, 0x91, 0x59, 0xf3, 0x0f, 0x23, 0x37, 0xbd,
	0x93, 0x2f, 0x57, 0xbf, 0xbd, 0xb9, 0x2d, 0xce, 0xed, 0xc0, 0x5e, 0x47, 0xe9, 0xcd, 0xcb, 0xcf,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x74, 0x1c, 0x7d, 0xef, 0x9d, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CredentialsServiceClient is the client API for CredentialsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CredentialsServiceClient interface {
	WhoAmI(ctx context.Context, in *WhoAmIRequest, opts ...grpc.CallOption) (*WhoAmIResponse, error)
}

type credentialsServiceClient struct {
	cc *grpc.ClientConn
}

func NewCredentialsServiceClient(cc *grpc.ClientConn) CredentialsServiceClient {
	return &credentialsServiceClient{cc}
}

func (c *credentialsServiceClient) WhoAmI(ctx context.Context, in *WhoAmIRequest, opts ...grpc.CallOption) (*WhoAmIResponse, error) {
	out := new(WhoAmIResponse)
	err := c.cc.Invoke(ctx, "/dolt.services.remotesapi.v1alpha1.CredentialsService/WhoAmI", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CredentialsServiceServer is the server API for CredentialsService service.
type CredentialsServiceServer interface {
	WhoAmI(context.Context, *WhoAmIRequest) (*WhoAmIResponse, error)
}

func RegisterCredentialsServiceServer(s *grpc.Server, srv CredentialsServiceServer) {
	s.RegisterService(&_CredentialsService_serviceDesc, srv)
}

func _CredentialsService_WhoAmI_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WhoAmIRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CredentialsServiceServer).WhoAmI(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dolt.services.remotesapi.v1alpha1.CredentialsService/WhoAmI",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CredentialsServiceServer).WhoAmI(ctx, req.(*WhoAmIRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CredentialsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dolt.services.remotesapi.v1alpha1.CredentialsService",
	HandlerType: (*CredentialsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "WhoAmI",
			Handler:    _CredentialsService_WhoAmI_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dolt/services/remotesapi/v1alpha1/credentials.proto",
}
