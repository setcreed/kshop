// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.1.0

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type ShopInterfaceHTTPServer interface {
	Register(context.Context, *RegisterReq) (*RegisterReply, error)
}

func RegisterShopInterfaceHTTPServer(s *http.Server, srv ShopInterfaceHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/register", _ShopInterface_Register0_HTTP_Handler(srv))
}

func _ShopInterface_Register0_HTTP_Handler(srv ShopInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RegisterReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/shop.interface.v1.ShopInterface/Register")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Register(ctx, req.(*RegisterReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RegisterReply)
		return ctx.Result(200, reply)
	}
}

type ShopInterfaceHTTPClient interface {
	Register(ctx context.Context, req *RegisterReq, opts ...http.CallOption) (rsp *RegisterReply, err error)
}

type ShopInterfaceHTTPClientImpl struct {
	cc *http.Client
}

func NewShopInterfaceHTTPClient(client *http.Client) ShopInterfaceHTTPClient {
	return &ShopInterfaceHTTPClientImpl{client}
}

func (c *ShopInterfaceHTTPClientImpl) Register(ctx context.Context, in *RegisterReq, opts ...http.CallOption) (*RegisterReply, error) {
	var out RegisterReply
	pattern := "/v1/register"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/shop.interface.v1.ShopInterface/Register"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}