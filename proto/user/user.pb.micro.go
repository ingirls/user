// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/user/user.proto

package go_micro_srv_user

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Service service

type Service interface {
	Say(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Create(ctx context.Context, in *User, opts ...client.CallOption) (*User, error)
	List(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Delete(ctx context.Context, in *User, opts ...client.CallOption) (*Null, error)
	Update(ctx context.Context, in *User, opts ...client.CallOption) (*User, error)
	InfoByID(ctx context.Context, in *ID, opts ...client.CallOption) (*User, error)
	InfoByUsername(ctx context.Context, in *Username, opts ...client.CallOption) (*User, error)
	InfoByMobile(ctx context.Context, in *Mobile, opts ...client.CallOption) (*User, error)
	InfoByEmail(ctx context.Context, in *Email, opts ...client.CallOption) (*User, error)
	UpdatePassword(ctx context.Context, in *IDAndPassword, opts ...client.CallOption) (*User, error)
	CheckPassword(ctx context.Context, in *IDAndPassword, opts ...client.CallOption) (*User, error)
}

type service struct {
	c    client.Client
	name string
}

func NewService(name string, c client.Client) Service {
	return &service{
		c:    c,
		name: name,
	}
}

func (c *service) Say(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Service.Say", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *service) Create(ctx context.Context, in *User, opts ...client.CallOption) (*User, error) {
	req := c.c.NewRequest(c.name, "Service.Create", in)
	out := new(User)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *service) List(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Service.List", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *service) Delete(ctx context.Context, in *User, opts ...client.CallOption) (*Null, error) {
	req := c.c.NewRequest(c.name, "Service.Delete", in)
	out := new(Null)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *service) Update(ctx context.Context, in *User, opts ...client.CallOption) (*User, error) {
	req := c.c.NewRequest(c.name, "Service.Update", in)
	out := new(User)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *service) InfoByID(ctx context.Context, in *ID, opts ...client.CallOption) (*User, error) {
	req := c.c.NewRequest(c.name, "Service.InfoByID", in)
	out := new(User)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *service) InfoByUsername(ctx context.Context, in *Username, opts ...client.CallOption) (*User, error) {
	req := c.c.NewRequest(c.name, "Service.InfoByUsername", in)
	out := new(User)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *service) InfoByMobile(ctx context.Context, in *Mobile, opts ...client.CallOption) (*User, error) {
	req := c.c.NewRequest(c.name, "Service.InfoByMobile", in)
	out := new(User)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *service) InfoByEmail(ctx context.Context, in *Email, opts ...client.CallOption) (*User, error) {
	req := c.c.NewRequest(c.name, "Service.InfoByEmail", in)
	out := new(User)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *service) UpdatePassword(ctx context.Context, in *IDAndPassword, opts ...client.CallOption) (*User, error) {
	req := c.c.NewRequest(c.name, "Service.UpdatePassword", in)
	out := new(User)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *service) CheckPassword(ctx context.Context, in *IDAndPassword, opts ...client.CallOption) (*User, error) {
	req := c.c.NewRequest(c.name, "Service.CheckPassword", in)
	out := new(User)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Service service

type ServiceHandler interface {
	Say(context.Context, *Request, *Response) error
	Create(context.Context, *User, *User) error
	List(context.Context, *Request, *Response) error
	Delete(context.Context, *User, *Null) error
	Update(context.Context, *User, *User) error
	InfoByID(context.Context, *ID, *User) error
	InfoByUsername(context.Context, *Username, *User) error
	InfoByMobile(context.Context, *Mobile, *User) error
	InfoByEmail(context.Context, *Email, *User) error
	UpdatePassword(context.Context, *IDAndPassword, *User) error
	CheckPassword(context.Context, *IDAndPassword, *User) error
}

func RegisterServiceHandler(s server.Server, hdlr ServiceHandler, opts ...server.HandlerOption) error {
	type service interface {
		Say(ctx context.Context, in *Request, out *Response) error
		Create(ctx context.Context, in *User, out *User) error
		List(ctx context.Context, in *Request, out *Response) error
		Delete(ctx context.Context, in *User, out *Null) error
		Update(ctx context.Context, in *User, out *User) error
		InfoByID(ctx context.Context, in *ID, out *User) error
		InfoByUsername(ctx context.Context, in *Username, out *User) error
		InfoByMobile(ctx context.Context, in *Mobile, out *User) error
		InfoByEmail(ctx context.Context, in *Email, out *User) error
		UpdatePassword(ctx context.Context, in *IDAndPassword, out *User) error
		CheckPassword(ctx context.Context, in *IDAndPassword, out *User) error
	}
	type Service struct {
		service
	}
	h := &serviceHandler{hdlr}
	return s.Handle(s.NewHandler(&Service{h}, opts...))
}

type serviceHandler struct {
	ServiceHandler
}

func (h *serviceHandler) Say(ctx context.Context, in *Request, out *Response) error {
	return h.ServiceHandler.Say(ctx, in, out)
}

func (h *serviceHandler) Create(ctx context.Context, in *User, out *User) error {
	return h.ServiceHandler.Create(ctx, in, out)
}

func (h *serviceHandler) List(ctx context.Context, in *Request, out *Response) error {
	return h.ServiceHandler.List(ctx, in, out)
}

func (h *serviceHandler) Delete(ctx context.Context, in *User, out *Null) error {
	return h.ServiceHandler.Delete(ctx, in, out)
}

func (h *serviceHandler) Update(ctx context.Context, in *User, out *User) error {
	return h.ServiceHandler.Update(ctx, in, out)
}

func (h *serviceHandler) InfoByID(ctx context.Context, in *ID, out *User) error {
	return h.ServiceHandler.InfoByID(ctx, in, out)
}

func (h *serviceHandler) InfoByUsername(ctx context.Context, in *Username, out *User) error {
	return h.ServiceHandler.InfoByUsername(ctx, in, out)
}

func (h *serviceHandler) InfoByMobile(ctx context.Context, in *Mobile, out *User) error {
	return h.ServiceHandler.InfoByMobile(ctx, in, out)
}

func (h *serviceHandler) InfoByEmail(ctx context.Context, in *Email, out *User) error {
	return h.ServiceHandler.InfoByEmail(ctx, in, out)
}

func (h *serviceHandler) UpdatePassword(ctx context.Context, in *IDAndPassword, out *User) error {
	return h.ServiceHandler.UpdatePassword(ctx, in, out)
}

func (h *serviceHandler) CheckPassword(ctx context.Context, in *IDAndPassword, out *User) error {
	return h.ServiceHandler.CheckPassword(ctx, in, out)
}
