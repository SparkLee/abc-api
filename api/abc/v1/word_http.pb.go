// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.0
// - protoc             v5.27.2
// source: abc/v1/word.proto

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

const OperationWordServiceCreateWord = "/api.abc.v1.WordService/CreateWord"
const OperationWordServiceDeleteWord = "/api.abc.v1.WordService/DeleteWord"
const OperationWordServiceGetWord = "/api.abc.v1.WordService/GetWord"
const OperationWordServiceListWord = "/api.abc.v1.WordService/ListWord"
const OperationWordServiceUpdateWord = "/api.abc.v1.WordService/UpdateWord"

type WordServiceHTTPServer interface {
	CreateWord(context.Context, *CreateWordRequest) (*CreateWordReply, error)
	DeleteWord(context.Context, *DeleteWordRequest) (*DeleteWordReply, error)
	GetWord(context.Context, *GetWordRequest) (*GetWordReply, error)
	ListWord(context.Context, *ListWordRequest) (*ListWordReply, error)
	UpdateWord(context.Context, *UpdateWordRequest) (*UpdateWordReply, error)
}

func RegisterWordServiceHTTPServer(s *http.Server, srv WordServiceHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/words", _WordService_CreateWord0_HTTP_Handler(srv))
	r.PUT("/v1/words/{id}", _WordService_UpdateWord0_HTTP_Handler(srv))
	r.DELETE("/v1/words/{id}", _WordService_DeleteWord0_HTTP_Handler(srv))
	r.GET("/v1/words/{id}", _WordService_GetWord0_HTTP_Handler(srv))
	r.GET("/v1/words", _WordService_ListWord0_HTTP_Handler(srv))
}

func _WordService_CreateWord0_HTTP_Handler(srv WordServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateWordRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationWordServiceCreateWord)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateWord(ctx, req.(*CreateWordRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateWordReply)
		return ctx.Result(200, reply)
	}
}

func _WordService_UpdateWord0_HTTP_Handler(srv WordServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateWordRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationWordServiceUpdateWord)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateWord(ctx, req.(*UpdateWordRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateWordReply)
		return ctx.Result(200, reply)
	}
}

func _WordService_DeleteWord0_HTTP_Handler(srv WordServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteWordRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationWordServiceDeleteWord)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteWord(ctx, req.(*DeleteWordRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteWordReply)
		return ctx.Result(200, reply)
	}
}

func _WordService_GetWord0_HTTP_Handler(srv WordServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetWordRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationWordServiceGetWord)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetWord(ctx, req.(*GetWordRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetWordReply)
		return ctx.Result(200, reply)
	}
}

func _WordService_ListWord0_HTTP_Handler(srv WordServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListWordRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationWordServiceListWord)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListWord(ctx, req.(*ListWordRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListWordReply)
		return ctx.Result(200, reply)
	}
}

type WordServiceHTTPClient interface {
	CreateWord(ctx context.Context, req *CreateWordRequest, opts ...http.CallOption) (rsp *CreateWordReply, err error)
	DeleteWord(ctx context.Context, req *DeleteWordRequest, opts ...http.CallOption) (rsp *DeleteWordReply, err error)
	GetWord(ctx context.Context, req *GetWordRequest, opts ...http.CallOption) (rsp *GetWordReply, err error)
	ListWord(ctx context.Context, req *ListWordRequest, opts ...http.CallOption) (rsp *ListWordReply, err error)
	UpdateWord(ctx context.Context, req *UpdateWordRequest, opts ...http.CallOption) (rsp *UpdateWordReply, err error)
}

type WordServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewWordServiceHTTPClient(client *http.Client) WordServiceHTTPClient {
	return &WordServiceHTTPClientImpl{client}
}

func (c *WordServiceHTTPClientImpl) CreateWord(ctx context.Context, in *CreateWordRequest, opts ...http.CallOption) (*CreateWordReply, error) {
	var out CreateWordReply
	pattern := "/v1/words"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationWordServiceCreateWord))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *WordServiceHTTPClientImpl) DeleteWord(ctx context.Context, in *DeleteWordRequest, opts ...http.CallOption) (*DeleteWordReply, error) {
	var out DeleteWordReply
	pattern := "/v1/words/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationWordServiceDeleteWord))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *WordServiceHTTPClientImpl) GetWord(ctx context.Context, in *GetWordRequest, opts ...http.CallOption) (*GetWordReply, error) {
	var out GetWordReply
	pattern := "/v1/words/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationWordServiceGetWord))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *WordServiceHTTPClientImpl) ListWord(ctx context.Context, in *ListWordRequest, opts ...http.CallOption) (*ListWordReply, error) {
	var out ListWordReply
	pattern := "/v1/words"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationWordServiceListWord))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *WordServiceHTTPClientImpl) UpdateWord(ctx context.Context, in *UpdateWordRequest, opts ...http.CallOption) (*UpdateWordReply, error) {
	var out UpdateWordReply
	pattern := "/v1/words/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationWordServiceUpdateWord))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}