// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package user

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

type IdReq struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IdReq) Reset()         { *m = IdReq{} }
func (m *IdReq) String() string { return proto.CompactTextString(m) }
func (*IdReq) ProtoMessage()    {}
func (*IdReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *IdReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IdReq.Unmarshal(m, b)
}
func (m *IdReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IdReq.Marshal(b, m, deterministic)
}
func (m *IdReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdReq.Merge(m, src)
}
func (m *IdReq) XXX_Size() int {
	return xxx_messageInfo_IdReq.Size(m)
}
func (m *IdReq) XXX_DiscardUnknown() {
	xxx_messageInfo_IdReq.DiscardUnknown(m)
}

var xxx_messageInfo_IdReq proto.InternalMessageInfo

func (m *IdReq) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type UserInfoReplay struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Number               string   `protobuf:"bytes,3,opt,name=number,proto3" json:"number,omitempty"`
	Gender               string   `protobuf:"bytes,4,opt,name=gender,proto3" json:"gender,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfoReplay) Reset()         { *m = UserInfoReplay{} }
func (m *UserInfoReplay) String() string { return proto.CompactTextString(m) }
func (*UserInfoReplay) ProtoMessage()    {}
func (*UserInfoReplay) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *UserInfoReplay) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfoReplay.Unmarshal(m, b)
}
func (m *UserInfoReplay) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfoReplay.Marshal(b, m, deterministic)
}
func (m *UserInfoReplay) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfoReplay.Merge(m, src)
}
func (m *UserInfoReplay) XXX_Size() int {
	return xxx_messageInfo_UserInfoReplay.Size(m)
}
func (m *UserInfoReplay) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfoReplay.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfoReplay proto.InternalMessageInfo

func (m *UserInfoReplay) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserInfoReplay) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserInfoReplay) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

func (m *UserInfoReplay) GetGender() string {
	if m != nil {
		return m.Gender
	}
	return ""
}

func init() {
	proto.RegisterType((*IdReq)(nil), "user.IdReq")
	proto.RegisterType((*UserInfoReplay)(nil), "user.UserInfoReplay")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 169 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x2d, 0x4e, 0x2d,
	0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0xc4, 0xb9, 0x58, 0x3d, 0x53,
	0x82, 0x52, 0x0b, 0x85, 0xf8, 0xb8, 0x98, 0x32, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x98, 0x83,
	0x98, 0x32, 0x53, 0x94, 0x52, 0xb8, 0xf8, 0x42, 0x8b, 0x53, 0x8b, 0x3c, 0xf3, 0xd2, 0xf2, 0x83,
	0x52, 0x0b, 0x72, 0x12, 0x2b, 0xd1, 0x55, 0x08, 0x09, 0x71, 0xb1, 0xe4, 0x25, 0xe6, 0xa6, 0x4a,
	0x30, 0x29, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x42, 0x62, 0x5c, 0x6c, 0x79, 0xa5, 0xb9, 0x49,
	0xa9, 0x45, 0x12, 0xcc, 0x60, 0x51, 0x28, 0x0f, 0x24, 0x9e, 0x9e, 0x9a, 0x97, 0x92, 0x5a, 0x24,
	0xc1, 0x02, 0x11, 0x87, 0xf0, 0x8c, 0x4c, 0xb8, 0xc0, 0xce, 0x10, 0xd2, 0xe1, 0x62, 0x4f, 0x4f,
	0x2d, 0x01, 0x59, 0x28, 0xc4, 0xad, 0x07, 0x76, 0x24, 0xd8, 0x55, 0x52, 0x22, 0x10, 0x0e, 0xaa,
	0x4b, 0x9c, 0xd8, 0xa2, 0xc0, 0xba, 0x92, 0xd8, 0xc0, 0x3e, 0x31, 0x06, 0x04, 0x00, 0x00, 0xff,
	0xff, 0x89, 0x63, 0x85, 0x38, 0xd7, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserClient interface {
	GetUser(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*UserInfoReplay, error)
}

type userClient struct {
	cc *grpc.ClientConn
}

func NewUserClient(cc *grpc.ClientConn) UserClient {
	return &userClient{cc}
}

func (c *userClient) GetUser(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*UserInfoReplay, error) {
	out := new(UserInfoReplay)
	err := c.cc.Invoke(ctx, "/user.user/getUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
type UserServer interface {
	GetUser(context.Context, *IdReq) (*UserInfoReplay, error)
}

// UnimplementedUserServer can be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (*UnimplementedUserServer) GetUser(ctx context.Context, req *IdReq) (*UserInfoReplay, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.user/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUser(ctx, req.(*IdReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.user",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getUser",
			Handler:    _User_GetUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}