// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/user/user.proto

/*
Package go_micro_srv_user is a generated protocol buffer package.

It is generated from these files:
	proto/user/user.proto

It has these top-level messages:
	Message
	Request
	Response
	Null
	User
	ID
	Username
	Mobile
	Email
	IDAndPassword
*/
package go_micro_srv_user

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

type STATUS int32

const (
	// 占位
	STATUS_UNKNOWN STATUS = 0
	// 正常
	STATUS_NORMAL STATUS = 1
	// 未激活
	STATUS_INACTIVE STATUS = 2
	// 禁用
	STATUS_DISABLED STATUS = 3
	// 删除
	STATUS_DELETED STATUS = 4
)

var STATUS_name = map[int32]string{
	0: "UNKNOWN",
	1: "NORMAL",
	2: "INACTIVE",
	3: "DISABLED",
	4: "DELETED",
}
var STATUS_value = map[string]int32{
	"UNKNOWN":  0,
	"NORMAL":   1,
	"INACTIVE": 2,
	"DISABLED": 3,
	"DELETED":  4,
}

func (x STATUS) String() string {
	return proto.EnumName(STATUS_name, int32(x))
}
func (STATUS) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Message struct {
	Say string `protobuf:"bytes,1,opt,name=say" json:"say,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Message) GetSay() string {
	if m != nil {
		return m.Say
	}
	return ""
}

type Request struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Response struct {
	Msg string `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type Null struct {
}

func (m *Null) Reset()                    { *m = Null{} }
func (m *Null) String() string            { return proto.CompactTextString(m) }
func (*Null) ProtoMessage()               {}
func (*Null) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type User struct {
	Id        int64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Email     string `protobuf:"bytes,2,opt,name=email" json:"email,omitempty"`
	Username  string `protobuf:"bytes,3,opt,name=username" json:"username,omitempty"`
	Nickname  string `protobuf:"bytes,4,opt,name=nickname" json:"nickname,omitempty"`
	Mobile    string `protobuf:"bytes,5,opt,name=mobile" json:"mobile,omitempty"`
	Password  string `protobuf:"bytes,6,opt,name=password" json:"password,omitempty"`
	Salt      string `protobuf:"bytes,7,opt,name=salt" json:"salt,omitempty"`
	Status    STATUS `protobuf:"varint,8,opt,name=status,enum=go.micro.srv.user.STATUS" json:"status,omitempty"`
	Avatar    string `protobuf:"bytes,9,opt,name=avatar" json:"avatar,omitempty"`
	CreatedAt int64  `protobuf:"varint,10,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	UpdatedAt int64  `protobuf:"varint,11,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *User) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *User) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetSalt() string {
	if m != nil {
		return m.Salt
	}
	return ""
}

func (m *User) GetStatus() STATUS {
	if m != nil {
		return m.Status
	}
	return STATUS_UNKNOWN
}

func (m *User) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *User) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *User) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

type ID struct {
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *ID) Reset()                    { *m = ID{} }
func (m *ID) String() string            { return proto.CompactTextString(m) }
func (*ID) ProtoMessage()               {}
func (*ID) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ID) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Username struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
}

func (m *Username) Reset()                    { *m = Username{} }
func (m *Username) String() string            { return proto.CompactTextString(m) }
func (*Username) ProtoMessage()               {}
func (*Username) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Username) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type Mobile struct {
	Mobile string `protobuf:"bytes,1,opt,name=mobile" json:"mobile,omitempty"`
}

func (m *Mobile) Reset()                    { *m = Mobile{} }
func (m *Mobile) String() string            { return proto.CompactTextString(m) }
func (*Mobile) ProtoMessage()               {}
func (*Mobile) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Mobile) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

type Email struct {
	Email string `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`
}

func (m *Email) Reset()                    { *m = Email{} }
func (m *Email) String() string            { return proto.CompactTextString(m) }
func (*Email) ProtoMessage()               {}
func (*Email) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Email) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type IDAndPassword struct {
	Id       int64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *IDAndPassword) Reset()                    { *m = IDAndPassword{} }
func (m *IDAndPassword) String() string            { return proto.CompactTextString(m) }
func (*IDAndPassword) ProtoMessage()               {}
func (*IDAndPassword) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *IDAndPassword) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *IDAndPassword) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func init() {
	proto.RegisterType((*Message)(nil), "go.micro.srv.user.Message")
	proto.RegisterType((*Request)(nil), "go.micro.srv.user.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.user.Response")
	proto.RegisterType((*Null)(nil), "go.micro.srv.user.Null")
	proto.RegisterType((*User)(nil), "go.micro.srv.user.User")
	proto.RegisterType((*ID)(nil), "go.micro.srv.user.ID")
	proto.RegisterType((*Username)(nil), "go.micro.srv.user.Username")
	proto.RegisterType((*Mobile)(nil), "go.micro.srv.user.Mobile")
	proto.RegisterType((*Email)(nil), "go.micro.srv.user.Email")
	proto.RegisterType((*IDAndPassword)(nil), "go.micro.srv.user.IDAndPassword")
	proto.RegisterEnum("go.micro.srv.user.STATUS", STATUS_name, STATUS_value)
}

func init() { proto.RegisterFile("proto/user/user.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 592 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0x5f, 0x6f, 0xd3, 0x3c,
	0x14, 0xc6, 0x97, 0x34, 0x4d, 0xb3, 0xb3, 0xad, 0xca, 0x6b, 0x6d, 0x2f, 0xa1, 0x63, 0x52, 0xe5,
	0x0b, 0x34, 0x71, 0x11, 0xc4, 0xb8, 0x03, 0x2e, 0x96, 0x36, 0x41, 0xca, 0xd6, 0x66, 0x28, 0x6d,
	0xe1, 0x12, 0x79, 0x8d, 0x29, 0xd1, 0x92, 0xa6, 0xc4, 0x69, 0x51, 0x3f, 0x18, 0x9f, 0x8a, 0x2f,
	0x81, 0x6c, 0xa7, 0xa5, 0x1b, 0xc9, 0x05, 0x7f, 0x6e, 0x2a, 0x9f, 0x3c, 0xe7, 0x39, 0xb6, 0x7f,
	0xe7, 0xb8, 0x70, 0xb2, 0xc8, 0xb3, 0x22, 0x7b, 0xbe, 0x64, 0x34, 0x17, 0x3f, 0xb6, 0x88, 0xd1,
	0x7f, 0xb3, 0xcc, 0x4e, 0xe3, 0x69, 0x9e, 0xd9, 0x2c, 0x5f, 0xd9, 0x5c, 0xc0, 0xa7, 0xd0, 0x1a,
	0x52, 0xc6, 0xc8, 0x8c, 0x22, 0x13, 0x1a, 0x8c, 0xac, 0x2d, 0xa5, 0xab, 0x9c, 0xef, 0x87, 0x7c,
	0x89, 0xcf, 0xa0, 0x15, 0xd2, 0x2f, 0x4b, 0xca, 0x0a, 0x84, 0x40, 0x9b, 0x93, 0x94, 0x96, 0xaa,
	0x58, 0xe3, 0x27, 0x60, 0x84, 0x94, 0x2d, 0xb2, 0x39, 0x13, 0xe6, 0x94, 0xcd, 0x36, 0xe6, 0x94,
	0xcd, 0xb0, 0x0e, 0x5a, 0xb0, 0x4c, 0x12, 0xfc, 0x4d, 0x05, 0x6d, 0xc2, 0x68, 0x8e, 0xda, 0xa0,
	0xc6, 0x91, 0xc8, 0x68, 0x84, 0x6a, 0x1c, 0xa1, 0x63, 0x68, 0xd2, 0x94, 0xc4, 0x89, 0xa5, 0x0a,
	0x93, 0x0c, 0x50, 0x07, 0x0c, 0x7e, 0x30, 0xb1, 0x59, 0x43, 0x08, 0xdb, 0x98, 0x6b, 0xf3, 0x78,
	0x7a, 0x27, 0x34, 0x4d, 0x6a, 0x9b, 0x18, 0xfd, 0x0f, 0x7a, 0x9a, 0xdd, 0xc6, 0x09, 0xb5, 0x9a,
	0x42, 0x29, 0x23, 0xee, 0x59, 0x10, 0xc6, 0xbe, 0x66, 0x79, 0x64, 0xe9, 0xd2, 0xb3, 0x89, 0xf9,
	0xa5, 0x18, 0x49, 0x0a, 0xab, 0x25, 0x2f, 0xc5, 0xd7, 0xe8, 0x05, 0xe8, 0xac, 0x20, 0xc5, 0x92,
	0x59, 0x46, 0x57, 0x39, 0x6f, 0x5f, 0x3c, 0xb6, 0x7f, 0x81, 0x66, 0x8f, 0xc6, 0xce, 0x78, 0x32,
	0x0a, 0xcb, 0x44, 0xbe, 0x35, 0x59, 0x91, 0x82, 0xe4, 0xd6, 0xbe, 0xdc, 0x5a, 0x46, 0xe8, 0x0c,
	0x60, 0x9a, 0x53, 0x52, 0xd0, 0xe8, 0x23, 0x29, 0x2c, 0x10, 0x17, 0xdf, 0x2f, 0xbf, 0x38, 0x05,
	0x97, 0x97, 0x8b, 0x68, 0x23, 0x1f, 0x48, 0xb9, 0xfc, 0xe2, 0x14, 0xf8, 0x18, 0x54, 0xdf, 0x7d,
	0x08, 0x0d, 0x3f, 0x05, 0x63, 0xb2, 0x83, 0x63, 0x8b, 0x4a, 0xb9, 0x8f, 0x0a, 0x77, 0x41, 0x1f,
	0x4a, 0x00, 0x3f, 0xc1, 0x28, 0xbb, 0x60, 0xf0, 0x19, 0x34, 0x3d, 0x41, 0x7c, 0xdb, 0x07, 0x65,
	0xa7, 0x0f, 0xf8, 0x35, 0x1c, 0xf9, 0xae, 0x33, 0x8f, 0xde, 0x6d, 0x60, 0x3d, 0x6c, 0xdf, 0x2e,
	0x58, 0xf5, 0x3e, 0xd8, 0x67, 0x57, 0xa0, 0x4b, 0x46, 0xe8, 0x00, 0x5a, 0x93, 0xe0, 0x3a, 0xb8,
	0xf9, 0x10, 0x98, 0x7b, 0x08, 0x40, 0x0f, 0x6e, 0xc2, 0xa1, 0x33, 0x30, 0x15, 0x74, 0x08, 0x86,
	0x1f, 0x38, 0xfd, 0xb1, 0xff, 0xde, 0x33, 0x55, 0x1e, 0xb9, 0xfe, 0xc8, 0xe9, 0x0d, 0x3c, 0xd7,
	0x6c, 0x70, 0x93, 0xeb, 0x0d, 0xbc, 0xb1, 0xe7, 0x9a, 0xda, 0xc5, 0xf7, 0x26, 0xb4, 0x46, 0x34,
	0x5f, 0xc5, 0x53, 0x8a, 0x2e, 0xa1, 0x31, 0x22, 0x6b, 0xd4, 0xa9, 0xe8, 0x49, 0x39, 0xa8, 0x9d,
	0xd3, 0x4a, 0x4d, 0x4e, 0x29, 0xde, 0x43, 0x6f, 0x40, 0xef, 0x8b, 0x0e, 0xa0, 0x47, 0x15, 0x89,
	0x1c, 0x6d, 0xa7, 0x4e, 0xc0, 0x7b, 0xc8, 0x01, 0x6d, 0x10, 0xb3, 0xe2, 0x2f, 0x0f, 0xe0, 0xd2,
	0x84, 0xfe, 0xee, 0x01, 0xc4, 0x53, 0x12, 0xee, 0x89, 0x98, 0x90, 0x3f, 0x3a, 0xfe, 0x2b, 0x30,
	0xfc, 0xf9, 0xa7, 0xac, 0xb7, 0xf6, 0x5d, 0x74, 0x52, 0x91, 0xe6, 0xbb, 0xb5, 0x6e, 0xf4, 0x16,
	0xda, 0xd2, 0xbb, 0x1d, 0xbf, 0xd3, 0x9a, 0x54, 0x2e, 0xd6, 0xd7, 0xe9, 0xc1, 0xa1, 0xac, 0x53,
	0x8e, 0x67, 0xd5, 0xfb, 0x92, 0x52, 0x7d, 0x8d, 0x4b, 0x38, 0x90, 0x35, 0xe4, 0x00, 0x5b, 0x15,
	0x79, 0x42, 0xa9, 0xaf, 0x70, 0x0d, 0x6d, 0xc9, 0x71, 0x3b, 0xde, 0xdd, 0x4a, 0x1e, 0x3b, 0x0f,
	0xa0, 0xbe, 0xd8, 0x15, 0x1c, 0xf5, 0x3f, 0xd3, 0xe9, 0xdd, 0x3f, 0xa8, 0x75, 0xab, 0x8b, 0x7f,
	0xea, 0x97, 0x3f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x58, 0xff, 0x47, 0x22, 0xc2, 0x05, 0x00, 0x00,
}
