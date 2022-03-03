// Code generated by protoc-gen-go. DO NOT EDIT.
// source: agent/proto/bot.proto

package go_micro_bot

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

type HelpRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelpRequest) Reset()         { *m = HelpRequest{} }
func (m *HelpRequest) String() string { return proto.CompactTextString(m) }
func (*HelpRequest) ProtoMessage()    {}
func (*HelpRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_79b974b8c77805fa, []int{0}
}

func (m *HelpRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelpRequest.Unmarshal(m, b)
}
func (m *HelpRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelpRequest.Marshal(b, m, deterministic)
}
func (m *HelpRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelpRequest.Merge(m, src)
}
func (m *HelpRequest) XXX_Size() int {
	return xxx_messageInfo_HelpRequest.Size(m)
}
func (m *HelpRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelpRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelpRequest proto.InternalMessageInfo

type HelpResponse struct {
	Usage                string   `protobuf:"bytes,1,opt,name=usage,proto3" json:"usage,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelpResponse) Reset()         { *m = HelpResponse{} }
func (m *HelpResponse) String() string { return proto.CompactTextString(m) }
func (*HelpResponse) ProtoMessage()    {}
func (*HelpResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_79b974b8c77805fa, []int{1}
}

func (m *HelpResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelpResponse.Unmarshal(m, b)
}
func (m *HelpResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelpResponse.Marshal(b, m, deterministic)
}
func (m *HelpResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelpResponse.Merge(m, src)
}
func (m *HelpResponse) XXX_Size() int {
	return xxx_messageInfo_HelpResponse.Size(m)
}
func (m *HelpResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HelpResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HelpResponse proto.InternalMessageInfo

func (m *HelpResponse) GetUsage() string {
	if m != nil {
		return m.Usage
	}
	return ""
}

func (m *HelpResponse) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type ExecRequest struct {
	Args                 []string `protobuf:"bytes,1,rep,name=args,proto3" json:"args,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExecRequest) Reset()         { *m = ExecRequest{} }
func (m *ExecRequest) String() string { return proto.CompactTextString(m) }
func (*ExecRequest) ProtoMessage()    {}
func (*ExecRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_79b974b8c77805fa, []int{2}
}

func (m *ExecRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecRequest.Unmarshal(m, b)
}
func (m *ExecRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecRequest.Marshal(b, m, deterministic)
}
func (m *ExecRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecRequest.Merge(m, src)
}
func (m *ExecRequest) XXX_Size() int {
	return xxx_messageInfo_ExecRequest.Size(m)
}
func (m *ExecRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExecRequest proto.InternalMessageInfo

func (m *ExecRequest) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

type ExecResponse struct {
	Result               []byte   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExecResponse) Reset()         { *m = ExecResponse{} }
func (m *ExecResponse) String() string { return proto.CompactTextString(m) }
func (*ExecResponse) ProtoMessage()    {}
func (*ExecResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_79b974b8c77805fa, []int{3}
}

func (m *ExecResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecResponse.Unmarshal(m, b)
}
func (m *ExecResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecResponse.Marshal(b, m, deterministic)
}
func (m *ExecResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecResponse.Merge(m, src)
}
func (m *ExecResponse) XXX_Size() int {
	return xxx_messageInfo_ExecResponse.Size(m)
}
func (m *ExecResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExecResponse proto.InternalMessageInfo

func (m *ExecResponse) GetResult() []byte {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *ExecResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*HelpRequest)(nil), "go.micro.bot.HelpRequest")
	proto.RegisterType((*HelpResponse)(nil), "go.micro.bot.HelpResponse")
	proto.RegisterType((*ExecRequest)(nil), "go.micro.bot.ExecRequest")
	proto.RegisterType((*ExecResponse)(nil), "go.micro.bot.ExecResponse")
}

func init() { proto.RegisterFile("agent/proto/bot.proto", fileDescriptor_79b974b8c77805fa) }

var fileDescriptor_79b974b8c77805fa = []byte{
	// 234 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x3f, 0x4f, 0xc3, 0x30,
	0x10, 0xc5, 0x1b, 0x28, 0x45, 0xbd, 0x84, 0xc5, 0x02, 0x14, 0x3a, 0x05, 0x4f, 0x9d, 0x5c, 0x09,
	0x56, 0x24, 0x06, 0x04, 0x62, 0xce, 0x37, 0x48, 0xd2, 0x53, 0x14, 0xa9, 0xf1, 0x99, 0xb3, 0x23,
	0xf1, 0x1d, 0xf8, 0xd2, 0xc8, 0x7f, 0x06, 0xab, 0xea, 0x76, 0xcf, 0x67, 0xbd, 0xf7, 0x7b, 0x07,
	0x0f, 0xdd, 0x88, 0xda, 0x1d, 0x0c, 0x93, 0xa3, 0x43, 0x4f, 0x4e, 0x85, 0x49, 0x54, 0x23, 0xa9,
	0x79, 0x1a, 0x98, 0x54, 0x4f, 0x4e, 0xde, 0x41, 0xf9, 0x8d, 0x27, 0xd3, 0xe2, 0xcf, 0x82, 0xd6,
	0xc9, 0x2f, 0xa8, 0xa2, 0xb4, 0x86, 0xb4, 0x45, 0x71, 0x0f, 0x37, 0x8b, 0xed, 0x46, 0xac, 0x8b,
	0xa6, 0xd8, 0x6f, 0xdb, 0x28, 0x44, 0x03, 0xe5, 0x11, 0xed, 0xc0, 0x93, 0x71, 0x13, 0xe9, 0xfa,
	0x2a, 0xec, 0xf2, 0x27, 0xf9, 0x0c, 0xe5, 0xe7, 0x2f, 0x0e, 0xc9, 0x56, 0x08, 0x58, 0x77, 0x3c,
	0xda, 0xba, 0x68, 0xae, 0xf7, 0xdb, 0x36, 0xcc, 0xf2, 0x0d, 0xaa, 0xf8, 0x25, 0x45, 0x3d, 0xc2,
	0x86, 0xd1, 0x2e, 0x27, 0x17, 0xb2, 0xaa, 0x36, 0x29, 0x8f, 0x80, 0xcc, 0xc4, 0x29, 0x26, 0x8a,
	0x97, 0xbf, 0x02, 0x6e, 0x3f, 0x68, 0x9e, 0x3b, 0x7d, 0x14, 0xef, 0xb0, 0xf6, 0xd0, 0xe2, 0x49,
	0xe5, 0xd5, 0x54, 0xd6, 0x6b, 0xb7, 0xbb, 0xb4, 0x8a, 0xc1, 0x72, 0xe5, 0x0d, 0x3c, 0xca, 0xb9,
	0x41, 0xd6, 0xe0, 0xdc, 0x20, 0x27, 0x97, 0xab, 0x7e, 0x13, 0x4e, 0xfb, 0xfa, 0x1f, 0x00, 0x00,
	0xff, 0xff, 0xe8, 0x08, 0x5e, 0xad, 0x73, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CommandClient is the client API for Command service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CommandClient interface {
	Help(ctx context.Context, in *HelpRequest, opts ...grpc.CallOption) (*HelpResponse, error)
	Exec(ctx context.Context, in *ExecRequest, opts ...grpc.CallOption) (*ExecResponse, error)
}

type commandClient struct {
	cc *grpc.ClientConn
}

func NewCommandClient(cc *grpc.ClientConn) CommandClient {
	return &commandClient{cc}
}

func (c *commandClient) Help(ctx context.Context, in *HelpRequest, opts ...grpc.CallOption) (*HelpResponse, error) {
	out := new(HelpResponse)
	err := c.cc.Invoke(ctx, "/go.micro.bot.Command/Help", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commandClient) Exec(ctx context.Context, in *ExecRequest, opts ...grpc.CallOption) (*ExecResponse, error) {
	out := new(ExecResponse)
	err := c.cc.Invoke(ctx, "/go.micro.bot.Command/Exec", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommandServer is the server API for Command service.
type CommandServer interface {
	Help(context.Context, *HelpRequest) (*HelpResponse, error)
	Exec(context.Context, *ExecRequest) (*ExecResponse, error)
}

// UnimplementedCommandServer can be embedded to have forward compatible implementations.
type UnimplementedCommandServer struct {
}

func (*UnimplementedCommandServer) Help(ctx context.Context, req *HelpRequest) (*HelpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Help not implemented")
}
func (*UnimplementedCommandServer) Exec(ctx context.Context, req *ExecRequest) (*ExecResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Exec not implemented")
}

func RegisterCommandServer(s *grpc.Server, srv CommandServer) {
	s.RegisterService(&_Command_serviceDesc, srv)
}

func _Command_Help_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandServer).Help(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.bot.Command/Help",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandServer).Help(ctx, req.(*HelpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Command_Exec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandServer).Exec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.bot.Command/Exec",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandServer).Exec(ctx, req.(*ExecRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Command_serviceDesc = grpc.ServiceDesc{
	ServiceName: "go.micro.bot.Command",
	HandlerType: (*CommandServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Help",
			Handler:    _Command_Help_Handler,
		},
		{
			MethodName: "Exec",
			Handler:    _Command_Exec_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "agent/proto/bot.proto",
}